package experiment

import (
	crand "crypto/rand"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"fdxt-poc/internal/compare"
	"fdxt-poc/internal/crypto"
	"fdxt-poc/internal/demodata"
	"fdxt-poc/internal/fdxt"
	"fdxt-poc/internal/grpcwire"
	"fdxt-poc/internal/queryparser"
)

// Config controls the proof-of-concept experiment.
type Config struct {
	Runs         int
	Warmups      int
	CoverBudget  int
	OutputDir    string
	IncludePlots bool
	AllPairs     bool
}

type Record struct {
	Timestamp          string  `json:"timestamp"`
	Variant            string  `json:"variant"`
	Workload           string  `json:"workload"`
	Query              string  `json:"query"`
	N                  int     `json:"n"`
	RunID              int     `json:"run_id"`
	ResultCount        int     `json:"result_count"`
	GroundTruthCount   int     `json:"ground_truth_count"`
	ExactMatch         bool    `json:"exact_match"`
	FalsePositives     int     `json:"false_positives"`
	FalseNegatives     int     `json:"false_negatives"`
	SetupTimeNS        int64   `json:"setup_time_ns"`
	UpdateTimeNS       int64   `json:"update_time_ns"`
	TrapdoorTimeNS     int64   `json:"trapdoor_time_ns"`
	SearchTimeNS       int64   `json:"search_time_ns"`
	RecoveryTimeNS     int64   `json:"recovery_time_ns"`
	TotalTimeNS        int64   `json:"total_time_ns"`
	CommunicationBytes int     `json:"communication_bytes"`
	RequestBytes       int     `json:"request_bytes"`
	ResponseBytes      int     `json:"response_bytes"`
	StorageBytes       int64   `json:"storage_bytes"`
	GenuineVolume      int     `json:"genuine_volume"`
	ObservedVolume     int     `json:"observed_volume"`
	CoverVolume        int     `json:"cover_volume"`
	ResidualExposure   float64 `json:"residual_volume_exposure"`
	C1Overhead         float64 `json:"c1_overhead"`
}

type Summary struct {
	Variant              string  `json:"variant"`
	Workload             string  `json:"workload"`
	Runs                 int     `json:"runs"`
	MeanSetupMS          float64 `json:"mean_setup_ms"`
	MeanUpdateMS         float64 `json:"mean_update_ms"`
	MeanTrapdoorMS       float64 `json:"mean_trapdoor_ms"`
	MeanSearchMS         float64 `json:"mean_search_ms"`
	MeanLatencyMS        float64 `json:"mean_latency_ms"`
	MedianLatencyMS      float64 `json:"median_latency_ms"`
	StddevLatencyMS      float64 `json:"stddev_latency_ms"`
	MinLatencyMS         float64 `json:"min_latency_ms"`
	MaxLatencyMS         float64 `json:"max_latency_ms"`
	MeanCommBytes        float64 `json:"mean_communication_bytes"`
	MeanStorageBytes     float64 `json:"mean_storage_bytes"`
	MeanObservedVol      float64 `json:"mean_observed_volume"`
	MeanCoverVol         float64 `json:"mean_cover_volume"`
	MeanExposure         float64 `json:"mean_residual_volume_exposure"`
	MeanC1Overhead       float64 `json:"mean_c1_overhead"`
	MeanResultCount      float64 `json:"mean_result_count"`
	MeanGroundTruthCount float64 `json:"mean_ground_truth_count"`
	AllCorrect           bool    `json:"all_correct"`
}

type QueryWorkload struct {
	Name  string
	Query string
	N     int
}

// c1PolicyForVariant maps experiment variant names to C1 policies.
func c1PolicyForVariant(variant string) fdxt.C1Policy {
	switch variant {
	case "c1_sqrt":
		return fdxt.C1PolicySqrt
	case "c1_log":
		return fdxt.C1PolicyLog
	case "c1_hybrid":
		return fdxt.C1PolicyHybrid
	case "c1_power2":
		return fdxt.C1PolicyPower2
	default:
		return fdxt.C1PolicyPower2
	}
}

func isC1Variant(variant string) bool {
	return strings.HasPrefix(variant, "c1_")
}

func Run(cfg Config) ([]Record, error) {
	if cfg.Runs <= 0 {
		cfg.Runs = 1
	}
	if cfg.Warmups < 0 {
		cfg.Warmups = 0
	}
	if cfg.CoverBudget < 0 {
		return nil, fmt.Errorf("cover budget must be >= 0")
	}
	if cfg.OutputDir == "" {
		cfg.OutputDir = "results"
	}

	var workloads []QueryWorkload
	var err error

	if cfg.AllPairs {
		workloads, err = selectAllPairs()
	} else {
		workloads, err = selectWorkloads()
	}

	if err != nil {
		return nil, err
	}

	// Five experiment variants:
	//   corrected_fdxt = baseline without C1
	//   c1_power2      = existing power-of-two C1 policy
	//   c1_sqrt       = square-root adaptive policy
	//   c1_log        = logarithmic adaptive policy
	//   c1_hybrid      = square-root + logarithmic adaptive policy
	variants := []string{
		"corrected_fdxt",
		"c1_power2",
		"c1_sqrt",
		"c1_log",
		"c1_hybrid",
	}

	all := make([]Record, 0, (cfg.Runs+cfg.Warmups)*len(workloads)*len(variants))

	for run := 1; run <= cfg.Runs+cfg.Warmups; run++ {
		for _, variant := range variants {
			client, server, setupNS, updateNS := freshIndex()
			storageBytes := serverStorageBytes(server)

			for _, w := range workloads {
				q, err := queryparser.Parse(w.Query)
				if err != nil {
					return nil, err
				}

				startTotal := time.Now()

				var parts fdxt.SearchRequestParts

				if isC1Variant(variant) && len(q) == 2 {
					parts, err = client.PrepareSearchC1(q, fdxt.C1Config{
						MaxCover: cfg.CoverBudget,
						Policy:   c1PolicyForVariant(variant),
					})
					if err != nil {
						return nil, fmt.Errorf(
							"variant %s, workload %s: %w",
							variant,
							w.Name,
							err,
						)
					}
				} else {
					parts = client.PrepareSearchBaseline(q)
				}

				trapdoorNS := time.Since(startTotal).Nanoseconds()

				startSearch := time.Now()

				var res []fdxt.SearchOutput

				if isC1Variant(variant) && len(q) == 2 {
					res = server.ProcessQuery(parts)
				} else {
					res = server.ProcessSearch(
						parts.TKL,
						parts.Rows,
						parts.RestKeywords,
					)
				}

				searchNS := time.Since(startSearch).Nanoseconds()

				startRecovery := time.Now()
				got := client.RecoverSearch(q, res)
				recoveryNS := time.Since(startRecovery).Nanoseconds()

				totalNS := time.Since(startTotal).Nanoseconds()

				ground := demodata.GroundTruth(q)
				cmp := compare.Compare(got, ground)

				requestBytes, responseBytes :=
					grpcwire.SearchWireBytes(parts, res)

				// Genuine volume V is always the unpadded candidate volume.
				genuineVolume := len(parts.Rows)

				if isC1Variant(variant) && len(q) == 2 {
					baselineParts := client.PrepareSearchBaseline(q)
					genuineVolume = len(baselineParts.Rows)
				}

				observedVolume := len(res)

				cover := observedVolume - genuineVolume
				if cover < 0 {
					cover = 0
				}

				exposure := 1.0
				if observedVolume > 0 {
					exposure = float64(genuineVolume) /
						float64(observedVolume)
				}

				c1Overhead := 0.0
				if isC1Variant(variant) && genuineVolume > 0 {
					c1Overhead = float64(cover) /
						float64(genuineVolume)
				}

				// Warmup runs are executed but not recorded.
				if run <= cfg.Warmups {
					continue
				}

				recordRun := run - cfg.Warmups

				all = append(all, Record{
					Timestamp:        time.Now().UTC().Format(time.RFC3339Nano),
					Variant:          variant,
					Workload:         w.Name,
					Query:            w.Query,
					N:                len(q),
					RunID:            recordRun,
					ResultCount:      cmp.ReturnedCount,
					GroundTruthCount: cmp.GroundTruth,
					ExactMatch:       cmp.ExactMatch,
					FalsePositives:   len(cmp.FalsePositives),
					FalseNegatives:   len(cmp.FalseNegatives),

					SetupTimeNS:    setupNS,
					UpdateTimeNS:   updateNS,
					TrapdoorTimeNS: trapdoorNS,
					SearchTimeNS:   searchNS,
					RecoveryTimeNS: recoveryNS,
					TotalTimeNS:    totalNS,

					CommunicationBytes: requestBytes + responseBytes,
					RequestBytes:       requestBytes,
					ResponseBytes:      responseBytes,
					StorageBytes:       storageBytes,

					GenuineVolume:    genuineVolume,
					ObservedVolume:   observedVolume,
					CoverVolume:      cover,
					ResidualExposure: exposure,
					C1Overhead:       c1Overhead,
				})
			}
		}
	}

	return all, nil
}

func freshIndex() (*fdxt.Client, *fdxt.Server, int64, int64) {
	setupStart := time.Now()

	p := crypto.Init()
	client := fdxt.NewClient(p)
	server := fdxt.NewServer(p)

	setupNS := time.Since(setupStart).Nanoseconds()

	updateStart := time.Now()

	pool, err := client.BuildC1Pool(demodata.Keywords(), 64)
	if err != nil {
		panic(err)
	}

	type event struct {
		cover       *fdxt.C1PoolEntry
		keyword, id string
	}

	var events []event

	for _, d := range demodata.GetAllDocuments() {
		for _, kw := range d.Keywords {
			events = append(events, event{
				keyword: kw,
				id:      d.ID,
			})
		}
	}

	for i := range pool {
		ce := pool[i]
		events = append(events, event{cover: &ce})
	}

	for i := len(events) - 1; i > 0; i-- {
		n, err := crand.Int(
			crand.Reader,
			big.NewInt(int64(i+1)),
		)
		if err != nil {
			panic(err)
		}

		j := int(n.Int64())
		events[i], events[j] = events[j], events[i]
	}

	for _, e := range events {
		if e.cover != nil {
			if err := client.InstallC1Entry(server, *e.cover); err != nil {
				panic(err)
			}
		} else {
			client.Update(server, fdxt.OpAdd, e.keyword, e.id)
		}
	}

	client.C1Pool = pool

	return client, server, setupNS, time.Since(updateStart).Nanoseconds()
}

func serverStorageBytes(s *fdxt.Server) int64 {
	path := filepath.Join(
		os.TempDir(),
		fmt.Sprintf("fdxt-c1-state-%d.json", time.Now().UnixNano()),
	)

	if err := s.SaveState(path); err != nil {
		return 0
	}

	st, err := os.Stat(path)
	_ = os.Remove(path)

	if err != nil {
		return 0
	}

	return st.Size()
}

func selectAllPairs() ([]QueryWorkload, error) {
	keywords := demodata.Keywords()

	if len(keywords) < 2 {
		return nil, fmt.Errorf("need at least two keywords for all-pairs experiment")
	}

	out := make(
		[]QueryWorkload,
		0,
		len(keywords)*(len(keywords)-1)/2,
	)

	idx := 1

	for i := 0; i < len(keywords); i++ {
		for j := i + 1; j < len(keywords); j++ {
			q := keywords[i] + " AND " + keywords[j]

			out = append(out, QueryWorkload{
				Name:  fmt.Sprintf("pair_%03d", idx),
				Query: q,
				N:     2,
			})

			idx++
		}
	}

	return out, nil
}

func selectWorkloads() ([]QueryWorkload, error) {
	freq := demodata.KeywordFrequencies()
	keywords := demodata.Keywords()

	category := func(n int) string {
		switch {
		case n >= 20:
			return "high"
		case n >= 8:
			return "medium"
		default:
			return "low"
		}
	}

	type pair struct {
		a, b  string
		inter int
	}

	var pairs []pair

	for i, a := range keywords {
		for j := i + 1; j < len(keywords); j++ {
			b := keywords[j]

			set := make(map[string]struct{})

			for _, id := range demodata.GetDocumentsForKeyword(a) {
				set[id] = struct{}{}
			}

			inter := 0

			for _, id := range demodata.GetDocumentsForKeyword(b) {
				if _, ok := set[id]; ok {
					inter++
				}
			}

			pairs = append(pairs, pair{
				a:     a,
				b:     b,
				inter: inter,
			})
		}
	}

	find := func(c1, c2 string, nonzero bool) (pair, bool) {
		cand := make([]pair, 0)

		for _, p := range pairs {
			if category(freq[p.a]) != c1 ||
				category(freq[p.b]) != c2 {

				if category(freq[p.a]) != c2 ||
					category(freq[p.b]) != c1 {
					continue
				}
			}

			if (p.inter > 0) != nonzero {
				continue
			}

			cand = append(cand, p)
		}

		if len(cand) == 0 {
			return pair{}, false
		}

		sort.Slice(cand, func(i, j int) bool {
			if cand[i].inter != cand[j].inter {
				return cand[i].inter > cand[j].inter
			}

			q1 := cand[i].a + " AND " + cand[i].b
			q2 := cand[j].a + " AND " + cand[j].b

			return q1 < q2
		})

		return cand[0], true
	}

	wanted := []struct {
		name, c1, c2 string
		nonzero      bool
	}{
		{"high_high", "high", "high", true},
		{"high_medium", "high", "medium", true},
		{"medium_medium", "medium", "medium", true},
		{"medium_low", "medium", "low", true},
		{"low_low", "low", "low", true},
		{"zero_result", "low", "low", false},
	}

	out := make([]QueryWorkload, 0, len(wanted)+2)

	for _, w := range wanted {
		p, ok := find(w.c1, w.c2, w.nonzero)

		if !ok {
			return nil, fmt.Errorf(
				"could not find workload category %s",
				w.name,
			)
		}

		out = append(out, QueryWorkload{
			Name:  w.name,
			Query: p.a + " AND " + p.b,
			N:     2,
		})
	}

	// Correctness-only regression workloads; C1 is not applied to these.
	out = append(out, QueryWorkload{
		Name:  "n1_regression",
		Query: keywords[0],
		N:     1,
	})

	var n3 string

	for _, q := range demodata.QuerySet {
		if strings.Count(q, " AND ") == 2 {
			n3 = q
			break
		}
	}

	if n3 == "" {
		return nil, fmt.Errorf("no n=3 regression query available")
	}

	out = append(out, QueryWorkload{
		Name:  "n3_regression",
		Query: n3,
		N:     3,
	})

	return out, nil
}

func Summarize(records []Record) []Summary {
	groups := make(map[string][]Record)

	for _, r := range records {
		key := r.Variant + "\x00" + r.Workload
		groups[key] = append(groups[key], r)
	}

	keys := make([]string, 0, len(groups))

	for k := range groups {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	out := make([]Summary, 0, len(keys))

	for _, k := range keys {
		rs := groups[k]

		lat := make([]float64, len(rs))

		var setup,
			update,
			trapdoor,
			search,
			comm,
			storage,
			obs,
			cover,
			exposure,
			c1o,
			resultCount,
			gtCount float64

		correct := true

		for i, r := range rs {
			lat[i] = float64(r.TotalTimeNS) / 1e6

			resultCount += float64(r.ResultCount)
			gtCount += float64(r.GroundTruthCount)

			setup += float64(r.SetupTimeNS) / 1e6
			update += float64(r.UpdateTimeNS) / 1e6
			trapdoor += float64(r.TrapdoorTimeNS) / 1e6
			search += float64(r.SearchTimeNS) / 1e6

			comm += float64(r.CommunicationBytes)
			storage += float64(r.StorageBytes)
			obs += float64(r.ObservedVolume)
			cover += float64(r.CoverVolume)
			exposure += r.ResidualExposure
			c1o += r.C1Overhead

			correct = correct && r.ExactMatch
		}

		sort.Float64s(lat)

		meanLatency := mean(lat)
		med := median(lat)
		sd := stddev(lat, meanLatency)

		parts := strings.Split(k, "\x00")

		out = append(out, Summary{
			Variant:              parts[0],
			Workload:             parts[1],
			Runs:                 len(rs),
			MeanSetupMS:          setup / float64(len(rs)),
			MeanUpdateMS:         update / float64(len(rs)),
			MeanTrapdoorMS:       trapdoor / float64(len(rs)),
			MeanSearchMS:         search / float64(len(rs)),
			MeanLatencyMS:        meanLatency,
			MedianLatencyMS:      med,
			StddevLatencyMS:      sd,
			MinLatencyMS:         lat[0],
			MaxLatencyMS:         lat[len(lat)-1],
			MeanCommBytes:        comm / float64(len(rs)),
			MeanStorageBytes:     storage / float64(len(rs)),
			MeanObservedVol:      obs / float64(len(rs)),
			MeanCoverVol:         cover / float64(len(rs)),
			MeanExposure:         exposure / float64(len(rs)),
			MeanC1Overhead:       c1o / float64(len(rs)),
			MeanResultCount:      resultCount / float64(len(rs)),
			MeanGroundTruthCount: gtCount / float64(len(rs)),
			AllCorrect:           correct,
		})
	}

	return out
}

func mean(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}

	s := 0.0

	for _, x := range v {
		s += x
	}

	return s / float64(len(v))
}

func median(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}

	if len(v)%2 == 1 {
		return v[len(v)/2]
	}

	return (v[len(v)/2-1] + v[len(v)/2]) / 2
}

func stddev(v []float64, m float64) float64 {
	if len(v) <= 1 {
		return 0
	}

	s := 0.0

	for _, x := range v {
		d := x - m
		s += d * d
	}

	return math.Sqrt(s / float64(len(v)-1))
}

func WriteOutputs(
	records []Record,
	summaries []Summary,
	dir string,
	plots bool,
) error {
	leakage := AnalyzeLeakage(records)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if err := writeJSON(
		filepath.Join(dir, "raw_runs.json"),
		records,
	); err != nil {
		return err
	}

	if err := writeCSV(
		filepath.Join(dir, "raw_runs.csv"),
		records,
	); err != nil {
		return err
	}

	if err := writeSummaryCSV(
		filepath.Join(dir, "summary.csv"),
		summaries,
	); err != nil {
		return err
	}

	if err := writeCorrectnessCSV(
		filepath.Join(dir, "correctness.csv"),
		records,
	); err != nil {
		return err
	}

	if err := writeWorkloadsCSV(
		filepath.Join(dir, "workloads.csv"),
		records,
	); err != nil {
		return err
	}

	if err := writeLeakageCSV(
		filepath.Join(dir, "leakage_analysis.csv"),
		leakage,
	); err != nil {
		return err
	}

	if err := writeLeakagePointsCSV(
		filepath.Join(dir, "leakage_points.csv"),
		records,
	); err != nil {
		return err
	}

	if err := writeMarkdown(
		filepath.Join(dir, "summary_tables.md"),
		summaries,
		leakage,
	); err != nil {
		return err
	}

	if plots {
		if err := writePlots(dir, summaries); err != nil {
			return err
		}

		return writeLeakagePlots(dir, records)
	}

	return nil
}

func writeJSON(path string, v any) error {
	b, e := json.MarshalIndent(v, "", "  ")

	if e != nil {
		return e
	}

	return os.WriteFile(path, b, 0644)
}

func writeCSV(path string, records []Record) error {
	f, e := os.Create(path)

	if e != nil {
		return e
	}

	defer f.Close()

	w := csv.NewWriter(f)

	h := []string{
		"timestamp",
		"variant",
		"workload",
		"query",
		"n",
		"run_id",
		"result_count",
		"ground_truth_count",
		"exact_match",
		"false_positives",
		"false_negatives",
		"setup_time_ns",
		"update_time_ns",
		"trapdoor_time_ns",
		"search_time_ns",
		"recovery_time_ns",
		"total_time_ns",
		"communication_bytes",
		"request_bytes",
		"response_bytes",
		"storage_bytes",
		"genuine_volume",
		"observed_volume",
		"cover_volume",
		"residual_volume_exposure",
		"c1_overhead",
	}

	if e = w.Write(h); e != nil {
		return e
	}

	for _, r := range records {
		row := []string{
			r.Timestamp,
			r.Variant,
			r.Workload,
			r.Query,
			strconv.Itoa(r.N),
			strconv.Itoa(r.RunID),
			strconv.Itoa(r.ResultCount),
			strconv.Itoa(r.GroundTruthCount),
			strconv.FormatBool(r.ExactMatch),
			strconv.Itoa(r.FalsePositives),
			strconv.Itoa(r.FalseNegatives),
			strconv.FormatInt(r.SetupTimeNS, 10),
			strconv.FormatInt(r.UpdateTimeNS, 10),
			strconv.FormatInt(r.TrapdoorTimeNS, 10),
			strconv.FormatInt(r.SearchTimeNS, 10),
			strconv.FormatInt(r.RecoveryTimeNS, 10),
			strconv.FormatInt(r.TotalTimeNS, 10),
			strconv.Itoa(r.CommunicationBytes),
			strconv.Itoa(r.RequestBytes),
			strconv.Itoa(r.ResponseBytes),
			strconv.FormatInt(r.StorageBytes, 10),
			strconv.Itoa(r.GenuineVolume),
			strconv.Itoa(r.ObservedVolume),
			strconv.Itoa(r.CoverVolume),
			strconv.FormatFloat(r.ResidualExposure, 'f', 6, 64),
			strconv.FormatFloat(r.C1Overhead, 'f', 6, 64),
		}

		if e = w.Write(row); e != nil {
			return e
		}
	}

	w.Flush()

	return w.Error()
}

func writeSummaryCSV(path string, ss []Summary) error {
	f, e := os.Create(path)

	if e != nil {
		return e
	}

	defer f.Close()

	w := csv.NewWriter(f)

	_ = w.Write([]string{
		"variant",
		"workload",
		"runs",
		"mean_setup_ms",
		"mean_update_ms",
		"mean_trapdoor_ms",
		"mean_search_ms",
		"mean_latency_ms",
		"median_latency_ms",
		"stddev_latency_ms",
		"min_latency_ms",
		"max_latency_ms",
		"mean_communication_bytes",
		"mean_storage_bytes",
		"mean_observed_volume",
		"mean_cover_volume",
		"mean_residual_volume_exposure",
		"mean_c1_overhead",
		"mean_result_count",
		"mean_ground_truth_count",
		"all_correct",
	})

	for _, s := range ss {
		_ = w.Write([]string{
			s.Variant,
			s.Workload,
			strconv.Itoa(s.Runs),
			fmt.Sprintf("%.6f", s.MeanSetupMS),
			fmt.Sprintf("%.6f", s.MeanUpdateMS),
			fmt.Sprintf("%.6f", s.MeanTrapdoorMS),
			fmt.Sprintf("%.6f", s.MeanSearchMS),
			fmt.Sprintf("%.6f", s.MeanLatencyMS),
			fmt.Sprintf("%.6f", s.MedianLatencyMS),
			fmt.Sprintf("%.6f", s.StddevLatencyMS),
			fmt.Sprintf("%.6f", s.MinLatencyMS),
			fmt.Sprintf("%.6f", s.MaxLatencyMS),
			fmt.Sprintf("%.3f", s.MeanCommBytes),
			fmt.Sprintf("%.3f", s.MeanStorageBytes),
			fmt.Sprintf("%.3f", s.MeanObservedVol),
			fmt.Sprintf("%.3f", s.MeanCoverVol),
			fmt.Sprintf("%.6f", s.MeanExposure),
			fmt.Sprintf("%.6f", s.MeanC1Overhead),
			fmt.Sprintf("%.3f", s.MeanResultCount),
			fmt.Sprintf("%.3f", s.MeanGroundTruthCount),
			strconv.FormatBool(s.AllCorrect),
		})
	}

	w.Flush()

	return w.Error()
}

func writeWorkloadsCSV(path string, records []Record) error {
	type item struct {
		name, query string
		n, gt       int
	}

	m := make(map[string]item)

	for _, r := range records {
		key := r.Workload + "\x00" + r.Query

		if _, ok := m[key]; !ok {
			m[key] = item{
				name:  r.Workload,
				query: r.Query,
				n:     r.N,
				gt:    r.GroundTruthCount,
			}
		}
	}

	items := make([]item, 0, len(m))

	for _, v := range m {
		items = append(items, v)
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].name < items[j].name
	})

	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)

	if err := w.Write([]string{
		"workload",
		"query",
		"n",
		"ground_truth_count",
	}); err != nil {
		return err
	}

	for _, v := range items {
		if err := w.Write([]string{
			v.name,
			v.query,
			strconv.Itoa(v.n),
			strconv.Itoa(v.gt),
		}); err != nil {
			return err
		}
	}

	w.Flush()

	return w.Error()
}

func writeCorrectnessCSV(path string, rs []Record) error {
	f, e := os.Create(path)

	if e != nil {
		return e
	}

	defer f.Close()

	w := csv.NewWriter(f)

	_ = w.Write([]string{
		"variant",
		"workload",
		"query",
		"n",
		"run_id",
		"result_count",
		"ground_truth_count",
		"exact_match",
		"false_positives",
		"false_negatives",
	})

	for _, r := range rs {
		_ = w.Write([]string{
			r.Variant,
			r.Workload,
			r.Query,
			strconv.Itoa(r.N),
			strconv.Itoa(r.RunID),
			strconv.Itoa(r.ResultCount),
			strconv.Itoa(r.GroundTruthCount),
			strconv.FormatBool(r.ExactMatch),
			strconv.Itoa(r.FalsePositives),
			strconv.Itoa(r.FalseNegatives),
		})
	}

	w.Flush()

	return w.Error()
}

func writeLeakagePointsCSV(path string, records []Record) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)

	if err := w.Write([]string{
		"variant",
		"query",
		"genuine_volume",
		"observed_volume",
	}); err != nil {
		return err
	}

	seen := make(map[string]bool)

	for _, r := range records {
		if r.N != 2 {
			continue
		}

		key := r.Variant + "\x00" + r.Query

		if seen[key] {
			continue
		}

		seen[key] = true

		if err := w.Write([]string{
			r.Variant,
			r.Query,
			strconv.Itoa(r.GenuineVolume),
			strconv.Itoa(r.ObservedVolume),
		}); err != nil {
			return err
		}
	}

	w.Flush()

	return w.Error()
}

func writeLeakagePlots(dir string, records []Record) error {
	points := make(map[string][]leakagePoint)
	seen := make(map[string]bool)

	for _, r := range records {
		if r.N != 2 {
			continue
		}

		key := r.Variant + "\x00" + r.Query

		if seen[key] {
			continue
		}

		seen[key] = true

		points[r.Variant] = append(
			points[r.Variant],
			leakagePoint{
				query: r.Query,
				trueV: r.GenuineVolume,
				obsV:  r.ObservedVolume,
			},
		)
	}

	if err := writeLeakageScatterSVG(
		filepath.Join(dir, "leakage_volume_mapping.svg"),
		points,
	); err != nil {
		return err
	}

	return writeLeakageMetricsSVG(
		filepath.Join(dir, "leakage_metrics.svg"),
		AnalyzeLeakage(records),
	)
}

func writeLeakageScatterSVG(
	path string,
	byVariant map[string][]leakagePoint,
) error {
	w, h := 1200, 700
	left, top := 100, 70
	cw, ch := 950, 520

	maxV := 1

	for _, ps := range byVariant {
		for _, p := range ps {
			if p.trueV > maxV {
				maxV = p.trueV
			}

			if p.obsV > maxV {
				maxV = p.obsV
			}
		}
	}

	var b strings.Builder

	fmt.Fprintf(
		&b,
		"<svg xmlns='http://www.w3.org/2000/svg' width='%d' height='%d'>"+
			"<rect width='100%%' height='100%%' fill='white'/>"+
			"<text x='%d' y='35' font-size='24'>True vs observed response volume</text>",
		w,
		h,
		left,
	)

	fmt.Fprintf(
		&b,
		"<line x1='%d' y1='%d' x2='%d' y2='%d' stroke='black'/>"+
			"<line x1='%d' y1='%d' x2='%d' y2='%d' stroke='black'/>",
		left,
		top,
		left,
		top+ch,
		left,
		top+ch,
		left+cw,
		top+ch,
	)

	fmt.Fprintf(
		&b,
		"<line x1='%d' y1='%d' x2='%d' y2='%d' "+
			"stroke='#999' stroke-dasharray='6,6'/>",
		left,
		top+ch,
		left+cw,
		top,
	)

	variants := []string{
		"corrected_fdxt",
		"c1_power2",
		"c1_sqrt",
		"c1_log",
		"c1_hybrid",
	}

	colors := map[string]string{
		"corrected_fdxt": "#333333",
		"c1_power2":      "#555555",
		"c1_sqrt":        "#777777",
		"c1_log":         "#999999",
		"c1_hybrid":      "#111111",
	}

	for _, variant := range variants {
		for _, p := range byVariant[variant] {
			x := float64(left) +
				float64(p.trueV)/float64(maxV)*float64(cw)

			y := float64(top+ch) -
				float64(p.obsV)/float64(maxV)*float64(ch)

			fmt.Fprintf(
				&b,
				"<circle cx='%.1f' cy='%.1f' r='4' fill='%s' opacity='0.45'/>",
				x,
				y,
				colors[variant],
			)
		}
	}

	fmt.Fprintf(
		&b,
		"<text x='%d' y='%d' font-size='13'>"+
			"x = genuine candidate volume; "+
			"y = observed encrypted response volume</text>",
		left,
		top+ch+50,
	)

	b.WriteString("</svg>")

	return os.WriteFile(path, []byte(b.String()), 0644)
}

func writeLeakageMetricsSVG(
	path string,
	ls []LeakageResult,
) error {
	// Keep a stable ordering for the five experiment variants.
	variants := []string{
		"corrected_fdxt",
		"c1_power2",
		"c1_sqrt",
		"c1_log",
		"c1_hybrid",
	}

	byVariant := make(map[string]LeakageResult)

	for _, l := range ls {
		byVariant[l.Variant] = l
	}

	type metric struct {
		name  string
		value func(LeakageResult) float64
	}

	metrics := []metric{
		{
			"Inference accuracy",
			func(l LeakageResult) float64 {
				return l.InferenceAccuracy
			},
		},
		{
			"Normalized MI",
			func(l LeakageResult) float64 {
				return l.NormalizedMutualInformation
			},
		},
		{
			"Pearson r",
			func(l LeakageResult) float64 {
				return l.PearsonCorrelation
			},
		},
		{
			"Collision rate",
			func(l LeakageResult) float64 {
				return l.CollisionRate
			},
		},
	}

	w, h := 1400, 750
	left, top := 100, 70
	cw, ch := 1150, 520

	var b strings.Builder

	fmt.Fprintf(
		&b,
		"<svg xmlns='http://www.w3.org/2000/svg' width='%d' height='%d'>"+
			"<rect width='100%%' height='100%%' fill='white'/>"+
			"<text x='%d' y='35' font-size='24'>Leakage-analysis metrics</text>",
		w,
		h,
		left,
	)

	gap := float64(cw) / float64(len(metrics))
	barW := gap / float64(len(variants)+1)

	// These metrics are all normalized to the [0,1] range except Pearson,
	// which may be negative. For the visualization we use a zero baseline
	// and a fixed vertical range of [-1, 1].
	for i, m := range metrics {
		x0 := float64(left) + float64(i)*gap

		fmt.Fprintf(
			&b,
			"<text x='%.1f' y='%d' font-size='13'>%s</text>",
			x0,
			top-15,
			m.name,
		)

		// Draw axes for each metric.
		fmt.Fprintf(
			&b,
			"<line x1='%.1f' y1='%d' x2='%.1f' y2='%d' stroke='black'/>",
			x0,
			top,
			x0,
			top+ch,
		)

		for j, variant := range variants {
			l, ok := byVariant[variant]
			if !ok {
				continue
			}

			v := m.value(l)

			// All plotted leakage metrics are normally [0,1].
			// Clamp only for visualization.
			if v < 0 {
				v = 0
			}
			if v > 1 {
				v = 1
			}

			x := x0 + float64(j+1)*barW
			bh := v * float64(ch)
			y := float64(top+ch) - bh

			fmt.Fprintf(
				&b,
				"<rect x='%.1f' y='%.1f' width='%.1f' height='%.1f' "+
					"fill='#555'/>",
				x,
				y,
				barW*0.75,
				bh,
			)

			fmt.Fprintf(
				&b,
				"<text x='%.1f' y='%.1f' font-size='10'>%.3f</text>",
				x,
				y-5,
				v,
			)
		}

		fmt.Fprintf(
			&b,
			"<text x='%.1f' y='%d' font-size='11'>0</text>",
			x0,
			top+ch+20,
		)

		fmt.Fprintf(
			&b,
			"<text x='%.1f' y='%d' font-size='11'>1</text>",
			x0,
			top+10,
		)
	}

	// Legend.
	legendY := top + ch + 65

	for i, variant := range variants {
		x := left + i*230

		fmt.Fprintf(
			&b,
			"<rect x='%d' y='%d' width='14' height='14' fill='#555'/>"+
				"<text x='%d' y='%d' font-size='11'>%s</text>",
			x,
			legendY,
			x+20,
			legendY+12,
			variant,
		)
	}

	b.WriteString("</svg>")

	return os.WriteFile(path, []byte(b.String()), 0644)
}

func writeMarkdown(
	path string,
	ss []Summary,
	leakage []LeakageResult,
) error {
	var b strings.Builder

	b.WriteString("# Experiment Summary\n\n")

	b.WriteString(
		"All measurements are from actual execution. " +
			"The dataset is the deterministic 50-document local " +
			"proof-of-concept. `residual_volume_exposure` is a " +
			"response-volume cover proxy (genuine FDXT candidate " +
			"volume divided by observed response volume), not a " +
			"formal security metric.\n\n",
	)

	b.WriteString("## C1 Policies\n\n")

	b.WriteString(
		"| Variant | C1 policy |\n" +
			"|---|---|\n" +
			"| corrected_fdxt | No C1 |\n" +
			"| c1_power2 | Next power of two |\n" +
			"| c1_sqrt | V + ceil(sqrt(V)) |\n" +
			"| c1_log | V + ceil(log2(V+1)) |\n" +
			"| c1_hybrid | V + ceil(sqrt(V)) + ceil(log2(V+1)) |\n\n",
	)

	b.WriteString("## Table 1 — Correctness\n\n")

	b.WriteString(
		"| Variant | Workload | Runs | Mean result count | " +
			"Ground-truth count | All correct |\n" +
			"|---|---|---:|---:|---:|---|\n",
	)

	for _, s := range ss {
		b.WriteString(fmt.Sprintf(
			"| %s | %s | %d | %.2f | %.2f | %v |\n",
			s.Variant,
			s.Workload,
			s.Runs,
			s.MeanResultCount,
			s.MeanGroundTruthCount,
			s.AllCorrect,
		))
	}

	b.WriteString("\n## Table 2 — Performance\n\n")

	b.WriteString(
		"| Variant | Workload | Setup ms | Update ms | " +
			"Trapdoor ms | Search ms | Total ms | Median ms | Stddev ms |\n" +
			"|---|---|---:|---:|---:|---:|---:|---:|---:|\n",
	)

	for _, s := range ss {
		b.WriteString(fmt.Sprintf(
			"| %s | %s | %.3f | %.3f | %.3f | %.3f | %.3f | %.3f | %.3f |\n",
			s.Variant,
			s.Workload,
			s.MeanSetupMS,
			s.MeanUpdateMS,
			s.MeanTrapdoorMS,
			s.MeanSearchMS,
			s.MeanLatencyMS,
			s.MedianLatencyMS,
			s.StddevLatencyMS,
		))
	}

	b.WriteString("\n## Table 3 — Communication and storage\n\n")

	b.WriteString(
		"| Variant | Workload | Mean communication bytes | " +
			"Mean storage bytes |\n" +
			"|---|---|---:|---:|\n",
	)

	for _, s := range ss {
		b.WriteString(fmt.Sprintf(
			"| %s | %s | %.1f | %.1f |\n",
			s.Variant,
			s.Workload,
			s.MeanCommBytes,
			s.MeanStorageBytes,
		))
	}

	b.WriteString("\n## Table 4 — Volume / leakage-proxy measurements\n\n")

	b.WriteString(
		"| Variant | Workload | Mean observed volume | " +
			"Mean cover volume | Exposure proxy |\n" +
			"|---|---|---:|---:|---:|\n",
	)

	for _, s := range ss {
		b.WriteString(fmt.Sprintf(
			"| %s | %s | %.2f | %.2f | %.4f |\n",
			s.Variant,
			s.Workload,
			s.MeanObservedVol,
			s.MeanCoverVol,
			s.MeanExposure,
		))
	}

	b.WriteString("\n## Table 5 — C1 overhead\n\n")

	b.WriteString(
		"| Variant | Workload | Mean C1 overhead |\n" +
			"|---|---|---:|\n",
	)

	for _, s := range ss {
		b.WriteString(fmt.Sprintf(
			"| %s | %s | %.4f |\n",
			s.Variant,
			s.Workload,
			s.MeanC1Overhead,
		))
	}

	b.WriteString("\n## Table 6 — Response-volume leakage analysis\n\n")

	b.WriteString(
		"The leakage model gives the observer only the encrypted " +
			"response volume and asks it to infer the genuine candidate " +
			"volume. Results use deterministic 5-fold cross-validation; " +
			"repeated timing runs are collapsed to one point per query. " +
			"Lower inference accuracy, correlation, mutual information, " +
			"and normalized mutual information, and higher MAE/RMSE/" +
			"collision rate indicate less direct volume information " +
			"under this model. This is an experimental leakage " +
			"measurement, not a formal cryptographic proof.\n\n",
	)

	b.WriteString(
		"| Variant | Queries | Inference accuracy (95% CI) | " +
			"MAE (95% CI) | RMSE (95% CI) | Pearson r | " +
			"MI (bits, 95% CI) | NMI | Collision rate (95% CI) | " +
			"Unique observed volumes |\n" +
			"|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n",
	)

	for _, l := range leakage {
		b.WriteString(fmt.Sprintf(
			"| %s | %d | %.4f (%.4f–%.4f) | "+
				"%.4f (%.4f–%.4f) | %.4f (%.4f–%.4f) | "+
				"%.4f | %.4f (%.4f–%.4f) | %.4f | "+
				"%.4f (%.4f–%.4f) | %d |\n",
			l.Variant,
			l.Queries,
			l.InferenceAccuracy,
			l.AccuracyCI95Low,
			l.AccuracyCI95High,
			l.MeanAbsoluteError,
			l.MAECI95Low,
			l.MAECI95High,
			l.RMSE,
			l.RMSECI95Low,
			l.RMSECI95High,
			l.PearsonCorrelation,
			l.MutualInformationBits,
			l.MICI95Low,
			l.MICI95High,
			l.NormalizedMutualInformation,
			l.CollisionRate,
			l.CollisionCI95Low,
			l.CollisionCI95High,
			l.UniqueObservedVolumes,
		))
	}

	return os.WriteFile(path, []byte(b.String()), 0644)
}

func writeLeakageCSV(
	path string,
	ls []LeakageResult,
) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)

	if err := w.Write([]string{
		"variant",
		"queries",
		"folds",
		"inference_accuracy",
		"mean_absolute_error",
		"rmse",
		"pearson_correlation",
		"mutual_information_bits",
		"normalized_mutual_information",
		"collision_rate",
		"unique_true_volumes",
		"unique_observed_volumes",
	}); err != nil {
		return err
	}

	for _, l := range ls {
		if err := w.Write([]string{
			l.Variant,
			strconv.Itoa(l.Queries),
			strconv.Itoa(l.Folds),
			fmt.Sprintf("%.8f", l.InferenceAccuracy),
			fmt.Sprintf("%.8f", l.MeanAbsoluteError),
			fmt.Sprintf("%.8f", l.RMSE),
			fmt.Sprintf("%.8f", l.PearsonCorrelation),
			fmt.Sprintf("%.8f", l.MutualInformationBits),
			fmt.Sprintf("%.8f", l.NormalizedMutualInformation),
			fmt.Sprintf("%.8f", l.CollisionRate),
			strconv.Itoa(l.UniqueTrueVolumes),
			strconv.Itoa(l.UniqueObservedVolumes),
		}); err != nil {
			return err
		}
	}

	w.Flush()

	return w.Error()
}

// writePlots emits dependency-free SVG plots so the one-command
// experiment works on a minimal Kali Go installation.
func writePlots(
	dir string,
	ss []Summary,
) error {
	metrics := []struct {
		name  string
		value func(Summary) float64
	}{
		{
			"search_latency_ms",
			func(s Summary) float64 {
				return s.MeanLatencyMS
			},
		},
		{
			"communication_bytes",
			func(s Summary) float64 {
				return s.MeanCommBytes
			},
		},
		{
			"observed_volume",
			func(s Summary) float64 {
				return s.MeanObservedVol
			},
		},
		{
			"cover_volume",
			func(s Summary) float64 {
				return s.MeanCoverVol
			},
		},
		{
			"residual_volume_exposure",
			func(s Summary) float64 {
				return s.MeanExposure
			},
		},
		{
			"c1_overhead",
			func(s Summary) float64 {
				return s.MeanC1Overhead
			},
		},
	}

	for _, m := range metrics {
		if err := writeBarSVG(
			filepath.Join(dir, m.name+".svg"),
			ss,
			m.name,
			m.value,
		); err != nil {
			return err
		}
	}

	return nil
}

func writeBarSVG(
	path string,
	ss []Summary,
	title string,
	value func(Summary) float64,
) error {
	w, h := 1200, 650
	left, top := 180, 70
	chartW, chartH := 900, 500

	maxV := 0.0

	for _, s := range ss {
		v := value(s)

		if v > maxV {
			maxV = v
		}
	}

	if maxV == 0 {
		maxV = 1
	}

	if len(ss) == 0 {
		maxV = 1
	}

	barW := float64(chartW) /
		float64(len(ss)*2+1)

	var b strings.Builder

	fmt.Fprintf(
		&b,
		"<svg xmlns='http://www.w3.org/2000/svg' "+
			"width='%d' height='%d' viewBox='0 0 %d %d'>",
		w,
		h,
		w,
		h,
	)

	fmt.Fprintf(
		&b,
		"<rect width='100%%' height='100%%' fill='white'/>"+
			"<text x='%d' y='35' font-size='24'>%s</text>",
		left,
		title,
	)

	fmt.Fprintf(
		&b,
		"<line x1='%d' y1='%d' x2='%d' y2='%d' stroke='black'/>"+
			"<line x1='%d' y1='%d' x2='%d' y2='%d' stroke='black'/>",
		left,
		top,
		left,
		top+chartH,
		left,
		top+chartH,
		left+chartW,
		top+chartH,
	)

	for i, s := range ss {
		v := value(s)

		x := float64(left) +
			float64(2*i+1)*barW

		bh := v / maxV * float64(chartH)
		y := float64(top+chartH) - bh

		fmt.Fprintf(
			&b,
			"<rect x='%.1f' y='%.1f' width='%.1f' "+
				"height='%.1f' fill='#666666'/>",
			x,
			y,
			barW,
			bh,
		)

		fmt.Fprintf(
			&b,
			"<text x='%.1f' y='%d' font-size='11' "+
				"transform='rotate(-45 %.1f %d)'>%s/%s</text>",
			x+barW/2,
			top+chartH+35,
			x+barW/2,
			top+chartH+35,
			s.Variant,
			s.Workload,
		)

		fmt.Fprintf(
			&b,
			"<text x='%.1f' y='%.1f' font-size='11'>%.2f</text>",
			x,
			y-5,
			v,
		)
	}

	b.WriteString("</svg>")

	return os.WriteFile(
		path,
		[]byte(b.String()),
		0644,
	)
}
