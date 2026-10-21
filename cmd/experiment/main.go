package main

import (
	"flag"
	"fmt"
	"os"

	"fdxt-poc/internal/experiment"
)

func main() {
	runs := flag.Int("runs", 10, "number of measured runs per workload and variant")
	warmups := flag.Int("warmups", 1, "warm-up runs per workload and variant")
	budget := flag.Int("c1-cover-budget", 32, "maximum adaptive C1 cover entries")
	out := flag.String("out", "results", "output directory")
	noPlots := flag.Bool("no-plots", false, "skip SVG plot generation")
	allPairs := flag.Bool("all-pairs", true, "run every unique 2-keyword pair in the dataset (default: true)")
	flag.Parse()
	if *runs <= 0 || *warmups < 0 || *budget < 0 {
		fmt.Fprintln(os.Stderr, "runs must be > 0; warmups and c1-cover-budget must be >= 0")
		os.Exit(2)
	}

	fmt.Printf("FDXT experiment: corrected baseline vs corrected FDXT + C1\n")
	fmt.Printf("Dataset: deterministic 50-document local dataset\nRuns: %d  Warmups: %d  C1 cover budget: %d  All pairs: %t\n\n", *runs, *warmups, *budget, *allPairs)

	records, err := experiment.Run(experiment.Config{Runs: *runs, Warmups: *warmups, CoverBudget: *budget, OutputDir: *out, IncludePlots: !*noPlots, AllPairs: *allPairs})
	if err != nil {
		fmt.Fprintln(os.Stderr, "experiment:", err)
		os.Exit(1)
	}
	summaries := experiment.Summarize(records)
	if err := experiment.WriteOutputs(records, summaries, *out, !*noPlots); err != nil {
		fmt.Fprintln(os.Stderr, "write outputs:", err)
		os.Exit(1)
	}

	bad := 0
	for _, r := range records {
		if !r.ExactMatch {
			bad++
		}
	}
	fmt.Printf("Measured records: %d\nIncorrect records: %d\n", len(records), bad)
	fmt.Println("Outputs:")
	fmt.Printf("  %s/raw_runs.csv\n", *out)
	fmt.Printf("  %s/raw_runs.json\n", *out)
	fmt.Printf("  %s/summary.csv\n", *out)
	fmt.Printf("  %s/correctness.csv\n", *out)
	fmt.Printf("  %s/workloads.csv\n", *out)
	fmt.Printf("  %s/summary_tables.md\n", *out)
	fmt.Printf("  %s/leakage_analysis.csv\n", *out)
	if !*noPlots {
		fmt.Printf("  %s/*.svg\n", *out)
	}
	if bad != 0 {
		fmt.Fprintln(os.Stderr, "experiment completed with correctness failures")
		os.Exit(1)
	}
}
