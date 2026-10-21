# Experiment Summary

All measurements are from actual execution. The dataset is the deterministic 50-document local proof-of-concept. `residual_volume_exposure` is a response-volume cover proxy (genuine FDXT candidate volume divided by observed response volume), not a formal security metric.

## C1 Policies

| Variant | C1 policy |
|---|---|
| corrected_fdxt | No C1 |
| c1_power2 | Next power of two |
| c1_sqrt | V + ceil(sqrt(V)) |
| c1_log | V + ceil(log2(V+1)) |
| c1_hybrid | V + ceil(sqrt(V)) + ceil(log2(V+1)) |

## Table 1 — Correctness

| Variant | Workload | Runs | Mean result count | Ground-truth count | All correct |
|---|---|---:|---:|---:|---|
| c1_hybrid | high_high | 10 | 12.00 | 12.00 | true |
| c1_hybrid | high_medium | 10 | 9.00 | 9.00 | true |
| c1_hybrid | low_low | 10 | 5.00 | 5.00 | true |
| c1_hybrid | medium_low | 10 | 6.00 | 6.00 | true |
| c1_hybrid | medium_medium | 10 | 10.00 | 10.00 | true |
| c1_hybrid | n1_regression | 10 | 21.00 | 21.00 | true |
| c1_hybrid | n3_regression | 10 | 2.00 | 2.00 | true |
| c1_hybrid | zero_result | 10 | 0.00 | 0.00 | true |
| c1_log | high_high | 10 | 12.00 | 12.00 | true |
| c1_log | high_medium | 10 | 9.00 | 9.00 | true |
| c1_log | low_low | 10 | 5.00 | 5.00 | true |
| c1_log | medium_low | 10 | 6.00 | 6.00 | true |
| c1_log | medium_medium | 10 | 10.00 | 10.00 | true |
| c1_log | n1_regression | 10 | 21.00 | 21.00 | true |
| c1_log | n3_regression | 10 | 2.00 | 2.00 | true |
| c1_log | zero_result | 10 | 0.00 | 0.00 | true |
| c1_power2 | high_high | 10 | 12.00 | 12.00 | true |
| c1_power2 | high_medium | 10 | 9.00 | 9.00 | true |
| c1_power2 | low_low | 10 | 5.00 | 5.00 | true |
| c1_power2 | medium_low | 10 | 6.00 | 6.00 | true |
| c1_power2 | medium_medium | 10 | 10.00 | 10.00 | true |
| c1_power2 | n1_regression | 10 | 21.00 | 21.00 | true |
| c1_power2 | n3_regression | 10 | 2.00 | 2.00 | true |
| c1_power2 | zero_result | 10 | 0.00 | 0.00 | true |
| c1_sqrt | high_high | 10 | 12.00 | 12.00 | true |
| c1_sqrt | high_medium | 10 | 9.00 | 9.00 | true |
| c1_sqrt | low_low | 10 | 5.00 | 5.00 | true |
| c1_sqrt | medium_low | 10 | 6.00 | 6.00 | true |
| c1_sqrt | medium_medium | 10 | 10.00 | 10.00 | true |
| c1_sqrt | n1_regression | 10 | 21.00 | 21.00 | true |
| c1_sqrt | n3_regression | 10 | 2.00 | 2.00 | true |
| c1_sqrt | zero_result | 10 | 0.00 | 0.00 | true |
| corrected_fdxt | high_high | 10 | 12.00 | 12.00 | true |
| corrected_fdxt | high_medium | 10 | 9.00 | 9.00 | true |
| corrected_fdxt | low_low | 10 | 5.00 | 5.00 | true |
| corrected_fdxt | medium_low | 10 | 6.00 | 6.00 | true |
| corrected_fdxt | medium_medium | 10 | 10.00 | 10.00 | true |
| corrected_fdxt | n1_regression | 10 | 21.00 | 21.00 | true |
| corrected_fdxt | n3_regression | 10 | 2.00 | 2.00 | true |
| corrected_fdxt | zero_result | 10 | 0.00 | 0.00 | true |

## Table 2 — Performance

| Variant | Workload | Setup ms | Update ms | Trapdoor ms | Search ms | Total ms | Median ms | Stddev ms |
|---|---|---:|---:|---:|---:|---:|---:|---:|
| c1_hybrid | high_high | 0.008 | 2987.784 | 43.923 | 39.883 | 87.111 | 91.064 | 14.733 |
| c1_hybrid | high_medium | 0.008 | 2987.784 | 31.157 | 33.524 | 66.984 | 57.247 | 23.915 |
| c1_hybrid | low_low | 0.008 | 2987.784 | 14.486 | 16.421 | 31.459 | 28.604 | 13.466 |
| c1_hybrid | medium_low | 0.008 | 2987.784 | 14.532 | 14.231 | 29.533 | 29.999 | 7.539 |
| c1_hybrid | medium_medium | 0.008 | 2987.784 | 22.547 | 24.605 | 48.793 | 50.540 | 13.598 |
| c1_hybrid | n1_regression | 0.008 | 2987.784 | 0.139 | 0.016 | 1.047 | 0.860 | 0.750 |
| c1_hybrid | n3_regression | 0.008 | 2987.784 | 43.303 | 44.741 | 88.720 | 87.481 | 32.704 |
| c1_hybrid | zero_result | 0.008 | 2987.784 | 13.107 | 12.814 | 26.936 | 24.215 | 7.000 |
| c1_log | high_high | 0.006 | 3085.541 | 38.816 | 37.344 | 78.809 | 74.978 | 24.821 |
| c1_log | high_medium | 0.006 | 3085.541 | 28.386 | 25.537 | 56.562 | 51.235 | 18.863 |
| c1_log | low_low | 0.006 | 3085.541 | 10.103 | 10.616 | 21.016 | 18.784 | 5.864 |
| c1_log | medium_low | 0.006 | 3085.541 | 10.800 | 12.103 | 23.395 | 22.893 | 6.567 |
| c1_log | medium_medium | 0.006 | 3085.541 | 19.003 | 18.653 | 38.397 | 35.559 | 11.830 |
| c1_log | n1_regression | 0.006 | 3085.541 | 0.124 | 0.011 | 0.676 | 0.587 | 0.171 |
| c1_log | n3_regression | 0.006 | 3085.541 | 37.700 | 35.934 | 74.104 | 73.922 | 15.594 |
| c1_log | zero_result | 0.006 | 3085.541 | 9.969 | 10.892 | 21.288 | 22.036 | 6.090 |
| c1_power2 | high_high | 0.008 | 3133.308 | 52.915 | 45.703 | 101.317 | 96.742 | 27.063 |
| c1_power2 | high_medium | 0.008 | 3133.308 | 23.459 | 23.104 | 47.170 | 44.525 | 13.661 |
| c1_power2 | low_low | 0.008 | 3133.308 | 14.515 | 13.830 | 28.679 | 28.662 | 12.198 |
| c1_power2 | medium_low | 0.008 | 3133.308 | 13.299 | 12.823 | 27.763 | 28.797 | 8.991 |
| c1_power2 | medium_medium | 0.008 | 3133.308 | 22.505 | 23.295 | 46.877 | 46.091 | 15.963 |
| c1_power2 | n1_regression | 0.008 | 3133.308 | 0.237 | 0.017 | 1.158 | 0.832 | 0.692 |
| c1_power2 | n3_regression | 0.008 | 3133.308 | 43.019 | 38.557 | 81.999 | 72.327 | 22.323 |
| c1_power2 | zero_result | 0.008 | 3133.308 | 12.949 | 12.297 | 25.635 | 29.554 | 8.945 |
| c1_sqrt | high_high | 0.011 | 5500.552 | 39.925 | 37.713 | 80.893 | 83.158 | 19.000 |
| c1_sqrt | high_medium | 0.011 | 5500.552 | 29.872 | 25.352 | 56.714 | 60.304 | 15.736 |
| c1_sqrt | low_low | 0.011 | 5500.552 | 11.963 | 11.846 | 24.310 | 17.797 | 10.807 |
| c1_sqrt | medium_low | 0.011 | 5500.552 | 11.984 | 12.295 | 24.679 | 24.832 | 7.586 |
| c1_sqrt | medium_medium | 0.011 | 5500.552 | 23.326 | 19.171 | 43.365 | 43.564 | 10.874 |
| c1_sqrt | n1_regression | 0.011 | 5500.552 | 0.253 | 0.023 | 1.402 | 1.140 | 0.888 |
| c1_sqrt | n3_regression | 0.011 | 5500.552 | 49.090 | 45.703 | 95.450 | 102.584 | 31.132 |
| c1_sqrt | zero_result | 0.011 | 5500.552 | 12.220 | 11.038 | 23.817 | 20.592 | 8.011 |
| corrected_fdxt | high_high | 0.006 | 3887.967 | 34.643 | 29.746 | 65.458 | 65.515 | 17.832 |
| corrected_fdxt | high_medium | 0.006 | 3887.967 | 22.520 | 22.743 | 45.760 | 44.786 | 14.272 |
| corrected_fdxt | low_low | 0.006 | 3887.967 | 7.763 | 7.113 | 14.957 | 12.738 | 5.421 |
| corrected_fdxt | medium_low | 0.006 | 3887.967 | 6.954 | 6.818 | 13.863 | 12.360 | 3.525 |
| corrected_fdxt | medium_medium | 0.006 | 3887.967 | 14.530 | 14.557 | 29.300 | 28.808 | 7.395 |
| corrected_fdxt | n1_regression | 0.006 | 3887.967 | 0.116 | 0.012 | 0.794 | 0.613 | 0.355 |
| corrected_fdxt | n3_regression | 0.006 | 3887.967 | 38.151 | 40.220 | 78.769 | 77.950 | 18.227 |
| corrected_fdxt | zero_result | 0.006 | 3887.967 | 7.262 | 8.137 | 15.465 | 12.256 | 6.007 |

## Table 3 — Communication and storage

| Variant | Workload | Mean communication bytes | Mean storage bytes |
|---|---|---:|---:|
| c1_hybrid | high_high | 11375.0 | 586970.0 |
| c1_hybrid | high_medium | 9495.0 | 586970.0 |
| c1_hybrid | low_low | 3831.0 | 586970.0 |
| c1_hybrid | medium_low | 4391.0 | 586970.0 |
| c1_hybrid | medium_medium | 6480.0 | 586970.0 |
| c1_hybrid | n1_regression | 2278.0 | 586970.0 |
| c1_hybrid | n3_regression | 8045.0 | 586970.0 |
| c1_hybrid | zero_result | 3940.0 | 586970.0 |
| c1_log | high_high | 9780.0 | 586970.0 |
| c1_log | high_medium | 8203.0 | 586970.0 |
| c1_log | low_low | 2877.0 | 586970.0 |
| c1_log | medium_low | 3437.0 | 586970.0 |
| c1_log | medium_medium | 5188.0 | 586970.0 |
| c1_log | n1_regression | 2278.0 | 586970.0 |
| c1_log | n3_regression | 8045.0 | 586970.0 |
| c1_log | zero_result | 2977.0 | 586970.0 |
| c1_power2 | high_high | 11694.0 | 586970.0 |
| c1_power2 | high_medium | 6588.0 | 586970.0 |
| c1_power2 | low_low | 2559.0 | 586970.0 |
| c1_power2 | medium_low | 3119.0 | 586970.0 |
| c1_power2 | medium_medium | 5188.0 | 586970.0 |
| c1_power2 | n1_regression | 2278.0 | 586970.0 |
| c1_power2 | n3_regression | 8045.0 | 586970.0 |
| c1_power2 | zero_result | 2656.0 | 586970.0 |
| c1_sqrt | high_high | 9780.0 | 586970.0 |
| c1_sqrt | high_medium | 7880.0 | 586970.0 |
| c1_sqrt | low_low | 2877.0 | 586970.0 |
| c1_sqrt | medium_low | 3437.0 | 586970.0 |
| c1_sqrt | medium_medium | 5188.0 | 586970.0 |
| c1_sqrt | n1_regression | 2278.0 | 586970.0 |
| c1_sqrt | n3_regression | 8045.0 | 586970.0 |
| c1_sqrt | zero_result | 2977.0 | 586970.0 |
| corrected_fdxt | high_high | 8185.0 | 586970.0 |
| corrected_fdxt | high_medium | 6588.0 | 586970.0 |
| corrected_fdxt | low_low | 1923.0 | 586970.0 |
| corrected_fdxt | medium_low | 2483.0 | 586970.0 |
| corrected_fdxt | medium_medium | 3896.0 | 586970.0 |
| corrected_fdxt | n1_regression | 2278.0 | 586970.0 |
| corrected_fdxt | n3_regression | 8045.0 | 586970.0 |
| corrected_fdxt | zero_result | 2014.0 | 586970.0 |

## Table 4 — Volume / leakage-proxy measurements

| Variant | Workload | Mean observed volume | Mean cover volume | Exposure proxy |
|---|---|---:|---:|---:|
| c1_hybrid | high_high | 31.00 | 10.00 | 0.6774 |
| c1_hybrid | high_medium | 25.00 | 9.00 | 0.6400 |
| c1_hybrid | low_low | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | medium_low | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | medium_medium | 20.00 | 8.00 | 0.6000 |
| c1_hybrid | n1_regression | 21.00 | 0.00 | 1.0000 |
| c1_hybrid | n3_regression | 16.00 | 0.00 | 1.0000 |
| c1_hybrid | zero_result | 12.00 | 6.00 | 0.5000 |
| c1_log | high_high | 26.00 | 5.00 | 0.8077 |
| c1_log | high_medium | 21.00 | 5.00 | 0.7619 |
| c1_log | low_low | 9.00 | 3.00 | 0.6667 |
| c1_log | medium_low | 9.00 | 3.00 | 0.6667 |
| c1_log | medium_medium | 16.00 | 4.00 | 0.7500 |
| c1_log | n1_regression | 21.00 | 0.00 | 1.0000 |
| c1_log | n3_regression | 16.00 | 0.00 | 1.0000 |
| c1_log | zero_result | 9.00 | 3.00 | 0.6667 |
| c1_power2 | high_high | 32.00 | 11.00 | 0.6562 |
| c1_power2 | high_medium | 16.00 | 0.00 | 1.0000 |
| c1_power2 | low_low | 8.00 | 2.00 | 0.7500 |
| c1_power2 | medium_low | 8.00 | 2.00 | 0.7500 |
| c1_power2 | medium_medium | 16.00 | 4.00 | 0.7500 |
| c1_power2 | n1_regression | 21.00 | 0.00 | 1.0000 |
| c1_power2 | n3_regression | 16.00 | 0.00 | 1.0000 |
| c1_power2 | zero_result | 8.00 | 2.00 | 0.7500 |
| c1_sqrt | high_high | 26.00 | 5.00 | 0.8077 |
| c1_sqrt | high_medium | 20.00 | 4.00 | 0.8000 |
| c1_sqrt | low_low | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | medium_low | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | medium_medium | 16.00 | 4.00 | 0.7500 |
| c1_sqrt | n1_regression | 21.00 | 0.00 | 1.0000 |
| c1_sqrt | n3_regression | 16.00 | 0.00 | 1.0000 |
| c1_sqrt | zero_result | 9.00 | 3.00 | 0.6667 |
| corrected_fdxt | high_high | 21.00 | 0.00 | 1.0000 |
| corrected_fdxt | high_medium | 16.00 | 0.00 | 1.0000 |
| corrected_fdxt | low_low | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | medium_low | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | medium_medium | 12.00 | 0.00 | 1.0000 |
| corrected_fdxt | n1_regression | 21.00 | 0.00 | 1.0000 |
| corrected_fdxt | n3_regression | 16.00 | 0.00 | 1.0000 |
| corrected_fdxt | zero_result | 6.00 | 0.00 | 1.0000 |

## Table 5 — C1 overhead

| Variant | Workload | Mean C1 overhead |
|---|---|---:|
| c1_hybrid | high_high | 0.4762 |
| c1_hybrid | high_medium | 0.5625 |
| c1_hybrid | low_low | 1.0000 |
| c1_hybrid | medium_low | 1.0000 |
| c1_hybrid | medium_medium | 0.6667 |
| c1_hybrid | n1_regression | 0.0000 |
| c1_hybrid | n3_regression | 0.0000 |
| c1_hybrid | zero_result | 1.0000 |
| c1_log | high_high | 0.2381 |
| c1_log | high_medium | 0.3125 |
| c1_log | low_low | 0.5000 |
| c1_log | medium_low | 0.5000 |
| c1_log | medium_medium | 0.3333 |
| c1_log | n1_regression | 0.0000 |
| c1_log | n3_regression | 0.0000 |
| c1_log | zero_result | 0.5000 |
| c1_power2 | high_high | 0.5238 |
| c1_power2 | high_medium | 0.0000 |
| c1_power2 | low_low | 0.3333 |
| c1_power2 | medium_low | 0.3333 |
| c1_power2 | medium_medium | 0.3333 |
| c1_power2 | n1_regression | 0.0000 |
| c1_power2 | n3_regression | 0.0000 |
| c1_power2 | zero_result | 0.3333 |
| c1_sqrt | high_high | 0.2381 |
| c1_sqrt | high_medium | 0.2500 |
| c1_sqrt | low_low | 0.5000 |
| c1_sqrt | medium_low | 0.5000 |
| c1_sqrt | medium_medium | 0.3333 |
| c1_sqrt | n1_regression | 0.0000 |
| c1_sqrt | n3_regression | 0.0000 |
| c1_sqrt | zero_result | 0.5000 |
| corrected_fdxt | high_high | 0.0000 |
| corrected_fdxt | high_medium | 0.0000 |
| corrected_fdxt | low_low | 0.0000 |
| corrected_fdxt | medium_low | 0.0000 |
| corrected_fdxt | medium_medium | 0.0000 |
| corrected_fdxt | n1_regression | 0.0000 |
| corrected_fdxt | n3_regression | 0.0000 |
| corrected_fdxt | zero_result | 0.0000 |

## Table 6 — Response-volume leakage analysis

The leakage model gives the observer only the encrypted response volume and asks it to infer the genuine candidate volume. Results use deterministic 5-fold cross-validation; repeated timing runs are collapsed to one point per query. Lower inference accuracy, correlation, mutual information, and normalized mutual information, and higher MAE/RMSE/collision rate indicate less direct volume information under this model. This is an experimental leakage measurement, not a formal cryptographic proof.

| Variant | Queries | Inference accuracy (95% CI) | MAE (95% CI) | RMSE (95% CI) | Pearson r | MI (bits, 95% CI) | NMI | Collision rate (95% CI) | Unique observed volumes |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| c1_hybrid | 6 | 0.5000 (0.5000–1.0000) | 3.6667 (0.0000–4.1667) | 5.8023 (0.0000–7.3598) | 0.9998 | 1.7925 (0.6500–1.9183) | 1.0000 | 0.0000 (0.0000–0.0000) | 4 |
| c1_log | 6 | 0.5000 (0.5000–1.0000) | 3.6667 (0.0000–4.1667) | 5.8023 (0.0000–7.3598) | 0.9994 | 1.7925 (0.6500–1.9183) | 1.0000 | 0.0000 (0.0000–0.0000) | 4 |
| c1_power2 | 6 | 0.5000 (0.0000–1.0000) | 3.3333 (0.0000–4.1667) | 5.4160 (0.0000–7.3598) | 0.9493 | 1.4591 (0.6500–1.5850) | 0.9022 | 0.0833 (0.0000–0.3333) | 3 |
| c1_sqrt | 6 | 0.5000 (0.5000–1.0000) | 3.6667 (0.0000–4.1667) | 5.8023 (0.0000–7.3598) | 0.9997 | 1.7925 (0.6500–1.9183) | 1.0000 | 0.0000 (0.0000–0.0000) | 4 |
| corrected_fdxt | 6 | 0.5000 (0.5000–1.0000) | 3.6667 (0.0000–4.1667) | 5.8023 (0.0000–7.3598) | 1.0000 | 1.7925 (0.6500–1.9183) | 1.0000 | 0.0000 (0.0000–0.0000) | 4 |
