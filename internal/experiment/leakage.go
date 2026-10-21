package experiment

import (
	"math"
	"math/rand"
	"sort"
)

// LeakageResult evaluates how much an observer of response volume can infer
// about the genuine candidate volume. The attacker sees only observed volume.
// Predictions use deterministic 5-fold cross-validation so a query is never
// used to train its own volume-to-volume inference rule.
type LeakageResult struct {
	Variant                     string  `json:"variant"`
	Queries                     int     `json:"queries"`
	Folds                       int     `json:"folds"`
	InferenceAccuracy           float64 `json:"inference_accuracy"`
	MeanAbsoluteError           float64 `json:"mean_absolute_error"`
	RMSE                        float64 `json:"rmse"`
	PearsonCorrelation          float64 `json:"pearson_correlation"`
	MutualInformationBits       float64 `json:"mutual_information_bits"`
	NormalizedMutualInformation float64 `json:"normalized_mutual_information"`
	CollisionRate               float64 `json:"collision_rate"`
	UniqueTrueVolumes           int     `json:"unique_true_volumes"`
	UniqueObservedVolumes       int     `json:"unique_observed_volumes"`
	AccuracyCI95Low             float64 `json:"accuracy_ci95_low"`
	AccuracyCI95High            float64 `json:"accuracy_ci95_high"`
	MAECI95Low                  float64 `json:"mae_ci95_low"`
	MAECI95High                 float64 `json:"mae_ci95_high"`
	RMSECI95Low                 float64 `json:"rmse_ci95_low"`
	RMSECI95High                float64 `json:"rmse_ci95_high"`
	MICI95Low                   float64 `json:"mi_ci95_low"`
	MICI95High                  float64 `json:"mi_ci95_high"`
	CollisionCI95Low            float64 `json:"collision_ci95_low"`
	CollisionCI95High           float64 `json:"collision_ci95_high"`
}

type leakagePoint struct {
	query string
	trueV int
	obsV  int
	idx   int
}

func AnalyzeLeakage(records []Record) []LeakageResult {
	byVariant := make(map[string]map[string]leakagePoint)
	for _, r := range records {
		if r.N != 2 {
			continue
		}
		if _, ok := byVariant[r.Variant]; !ok {
			byVariant[r.Variant] = make(map[string]leakagePoint)
		}
		// Repeated benchmark runs have the same logical volumes. Keep one
		// point per query so timing repetition does not overweight a query.
		key := r.Query
		if _, ok := byVariant[r.Variant][key]; !ok {
			byVariant[r.Variant][key] = leakagePoint{query: r.Query, trueV: r.GenuineVolume, obsV: r.ObservedVolume}
		}
	}

	variants := make([]string, 0, len(byVariant))
	for v := range byVariant {
		variants = append(variants, v)
	}
	sort.Strings(variants)
	out := make([]LeakageResult, 0, len(variants))
	for _, v := range variants {
		points := make([]leakagePoint, 0, len(byVariant[v]))
		for _, p := range byVariant[v] {
			points = append(points, p)
		}
		sort.Slice(points, func(i, j int) bool { return points[i].query < points[j].query })
		for i := range points {
			points[i].idx = i
		}
		out = append(out, analyzeVariant(v, points))
	}
	return out
}

func analyzeVariant(variant string, points []leakagePoint) LeakageResult {
	const folds = 5
	n := len(points)
	if n == 0 {
		return LeakageResult{Variant: variant, Folds: folds}
	}

	predictions := make([]float64, n)
	globalVals := make([]int, n)
	for i, p := range points {
		globalVals[i] = p.trueV
	}
	globalMedian := float64(medianInts(globalVals))

	for fold := 0; fold < folds; fold++ {
		train := make([]leakagePoint, 0, n-n/folds)
		test := make([]int, 0, n/folds+1)
		for i, p := range points {
			if i%folds == fold {
				test = append(test, i)
			} else {
				train = append(train, p)
			}
		}
		cond := make(map[int][]int)
		for _, p := range train {
			cond[p.obsV] = append(cond[p.obsV], p.trueV)
		}
		for _, i := range test {
			vals := cond[points[i].obsV]
			if len(vals) == 0 {
				predictions[i] = globalMedian
			} else {
				predictions[i] = float64(medianInts(vals))
			}
		}
	}

	var mae, mse float64
	correct := 0
	xs := make([]float64, n)
	ys := make([]float64, n)
	for i, p := range points {
		pred := predictions[i]
		err := pred - float64(p.trueV)
		mae += math.Abs(err)
		mse += err * err
		if math.Abs(err) < 1e-12 {
			correct++
		}
		xs[i] = float64(p.trueV)
		ys[i] = float64(p.obsV)
	}

	uniqueTrue := make(map[int]struct{})
	uniqueObs := make(map[int]struct{})
	for _, p := range points {
		uniqueTrue[p.trueV] = struct{}{}
		uniqueObs[p.obsV] = struct{}{}
	}

	acc := float64(correct) / float64(n)
	maeV := mae / float64(n)
	rmseV := math.Sqrt(mse / float64(n))
	miV := mutualInformation(xs, ys)
	collV := collisionRate(points)
	accLo, accHi := bootstrapMetric(points, metricAccuracy, 2000, 20260918)
	maeLo, maeHi := bootstrapMetric(points, metricMAE, 2000, 20260919)
	rmseLo, rmseHi := bootstrapMetric(points, metricRMSE, 2000, 20260920)
	miLo, miHi := bootstrapMetric(points, metricMI, 2000, 20260921)
	collLo, collHi := bootstrapMetric(points, metricCollision, 2000, 20260922)
	return LeakageResult{
		Variant:                     variant,
		Queries:                     n,
		Folds:                       folds,
		InferenceAccuracy:           acc,
		MeanAbsoluteError:           maeV,
		RMSE:                        rmseV,
		PearsonCorrelation:          pearson(xs, ys),
		MutualInformationBits:       miV,
		NormalizedMutualInformation: normalizedMI(xs, ys),
		CollisionRate:               collV,
		UniqueTrueVolumes:           len(uniqueTrue),
		UniqueObservedVolumes:       len(uniqueObs),
		AccuracyCI95Low:             accLo, AccuracyCI95High: accHi,
		MAECI95Low: maeLo, MAECI95High: maeHi,
		RMSECI95Low: rmseLo, RMSECI95High: rmseHi,
		MICI95Low: miLo, MICI95High: miHi,
		CollisionCI95Low: collLo, CollisionCI95High: collHi,
	}
}

func medianInts(v []int) int {
	if len(v) == 0 {
		return 0
	}
	x := append([]int(nil), v...)
	sort.Ints(x)
	if len(x)%2 == 1 {
		return x[len(x)/2]
	}
	return int(math.Round(float64(x[len(x)/2-1]+x[len(x)/2]) / 2.0))
}

func pearson(x, y []float64) float64 {
	if len(x) < 2 {
		return 0
	}
	mx, my := mean(x), mean(y)
	var num, dx, dy float64
	for i := range x {
		a, b := x[i]-mx, y[i]-my
		num += a * b
		dx += a * a
		dy += b * b
	}
	if dx == 0 || dy == 0 {
		return 1
	}
	return num / math.Sqrt(dx*dy)
}

func mutualInformation(x, y []float64) float64 {
	if len(x) == 0 {
		return 0
	}
	joint := make(map[[2]int]int)
	cx := make(map[int]int)
	cy := make(map[int]int)
	for i := range x {
		a, b := int(x[i]), int(y[i])
		joint[[2]int{a, b}]++
		cx[a]++
		cy[b]++
	}
	n := float64(len(x))
	mi := 0.0
	for k, c := range joint {
		pxy := float64(c) / n
		px := float64(cx[k[0]]) / n
		py := float64(cy[k[1]]) / n
		mi += pxy * math.Log2(pxy/(px*py))
	}
	return mi
}

func entropy(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}
	c := make(map[int]int)
	for _, x := range v {
		c[int(x)]++
	}
	n := float64(len(v))
	h := 0.0
	for _, k := range c {
		p := float64(k) / n
		h -= p * math.Log2(p)
	}
	return h
}

func normalizedMI(x, y []float64) float64 {
	mi := mutualInformation(x, y)
	denom := math.Sqrt(entropy(x) * entropy(y))
	if denom == 0 {
		return 0
	}
	return mi / denom
}

// Collision rate is the fraction of unordered query pairs whose genuine
// volumes differ but whose observed volumes collide. It measures direct
// volume indistinguishability under the stated observation model.
func collisionRate(points []leakagePoint) float64 {
	if len(points) < 2 {
		return 0
	}
	total, collisions := 0, 0
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i].trueV == points[j].trueV {
				continue
			}
			total++
			if points[i].obsV == points[j].obsV {
				collisions++
			}
		}
	}
	if total == 0 {
		return 0
	}
	return float64(collisions) / float64(total)
}

func metricAccuracy(p []leakagePoint) float64 {
	pred := crossValidatedPredictions(p)
	correct := 0
	for i := range p {
		if math.Abs(pred[i]-float64(p[i].trueV)) < 1e-12 {
			correct++
		}
	}
	return float64(correct) / float64(len(p))
}

func metricMAE(p []leakagePoint) float64 {
	pred := crossValidatedPredictions(p)
	var s float64
	for i := range p {
		s += math.Abs(pred[i] - float64(p[i].trueV))
	}
	return s / float64(len(p))
}

func metricRMSE(p []leakagePoint) float64 {
	pred := crossValidatedPredictions(p)
	var s float64
	for i := range p {
		d := pred[i] - float64(p[i].trueV)
		s += d * d
	}
	return math.Sqrt(s / float64(len(p)))
}

func metricMI(p []leakagePoint) float64 {
	x := make([]float64, len(p))
	y := make([]float64, len(p))
	for i := range p {
		x[i] = float64(p[i].trueV)
		y[i] = float64(p[i].obsV)
	}
	return mutualInformation(x, y)
}

func metricCollision(p []leakagePoint) float64 { return collisionRate(p) }

func crossValidatedPredictions(points []leakagePoint) []float64 {
	const folds = 5
	n := len(points)
	out := make([]float64, n)
	global := make([]int, n)
	for i, p := range points {
		global[i] = p.trueV
	}
	gm := float64(medianInts(global))
	for fold := 0; fold < folds; fold++ {
		cond := make(map[int][]int)
		for i, p := range points {
			if i%folds != fold {
				cond[p.obsV] = append(cond[p.obsV], p.trueV)
			}
		}
		for i, p := range points {
			if i%folds == fold {
				if vals := cond[p.obsV]; len(vals) > 0 {
					out[i] = float64(medianInts(vals))
				} else {
					out[i] = gm
				}
			}
		}
	}
	return out
}

func bootstrapMetric(points []leakagePoint, metric func([]leakagePoint) float64, reps int, seed int64) (float64, float64) {
	if len(points) < 2 {
		v := metric(points)
		return v, v
	}
	r := rand.New(rand.NewSource(seed))
	vals := make([]float64, reps)
	sample := make([]leakagePoint, len(points))
	for b := 0; b < reps; b++ {
		for i := range sample {
			sample[i] = points[r.Intn(len(points))]
		}
		vals[b] = metric(sample)
	}
	sort.Float64s(vals)
	lo := vals[int(0.025*float64(reps))]
	hi := vals[int(0.975*float64(reps))-1]
	return lo, hi
}
