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
| c1_hybrid | high_high | 1 | 12.00 | 12.00 | true |
| c1_hybrid | high_medium | 1 | 9.00 | 9.00 | true |
| c1_hybrid | low_low | 1 | 5.00 | 5.00 | true |
| c1_hybrid | medium_low | 1 | 6.00 | 6.00 | true |
| c1_hybrid | medium_medium | 1 | 10.00 | 10.00 | true |
| c1_hybrid | n1_regression | 1 | 21.00 | 21.00 | true |
| c1_hybrid | n3_regression | 1 | 2.00 | 2.00 | true |
| c1_hybrid | zero_result | 1 | 0.00 | 0.00 | true |
| c1_log | high_high | 1 | 12.00 | 12.00 | true |
| c1_log | high_medium | 1 | 9.00 | 9.00 | true |
| c1_log | low_low | 1 | 5.00 | 5.00 | true |
| c1_log | medium_low | 1 | 6.00 | 6.00 | true |
| c1_log | medium_medium | 1 | 10.00 | 10.00 | true |
| c1_log | n1_regression | 1 | 21.00 | 21.00 | true |
| c1_log | n3_regression | 1 | 2.00 | 2.00 | true |
| c1_log | zero_result | 1 | 0.00 | 0.00 | true |
| c1_power2 | high_high | 1 | 12.00 | 12.00 | true |
| c1_power2 | high_medium | 1 | 9.00 | 9.00 | true |
| c1_power2 | low_low | 1 | 5.00 | 5.00 | true |
| c1_power2 | medium_low | 1 | 6.00 | 6.00 | true |
| c1_power2 | medium_medium | 1 | 10.00 | 10.00 | true |
| c1_power2 | n1_regression | 1 | 21.00 | 21.00 | true |
| c1_power2 | n3_regression | 1 | 2.00 | 2.00 | true |
| c1_power2 | zero_result | 1 | 0.00 | 0.00 | true |
| c1_sqrt | high_high | 1 | 12.00 | 12.00 | true |
| c1_sqrt | high_medium | 1 | 9.00 | 9.00 | true |
| c1_sqrt | low_low | 1 | 5.00 | 5.00 | true |
| c1_sqrt | medium_low | 1 | 6.00 | 6.00 | true |
| c1_sqrt | medium_medium | 1 | 10.00 | 10.00 | true |
| c1_sqrt | n1_regression | 1 | 21.00 | 21.00 | true |
| c1_sqrt | n3_regression | 1 | 2.00 | 2.00 | true |
| c1_sqrt | zero_result | 1 | 0.00 | 0.00 | true |
| corrected_fdxt | high_high | 1 | 12.00 | 12.00 | true |
| corrected_fdxt | high_medium | 1 | 9.00 | 9.00 | true |
| corrected_fdxt | low_low | 1 | 5.00 | 5.00 | true |
| corrected_fdxt | medium_low | 1 | 6.00 | 6.00 | true |
| corrected_fdxt | medium_medium | 1 | 10.00 | 10.00 | true |
| corrected_fdxt | n1_regression | 1 | 21.00 | 21.00 | true |
| corrected_fdxt | n3_regression | 1 | 2.00 | 2.00 | true |
| corrected_fdxt | zero_result | 1 | 0.00 | 0.00 | true |

## Table 2 — Performance

| Variant | Workload | Setup ms | Update ms | Trapdoor ms | Search ms | Total ms | Median ms | Stddev ms |
|---|---|---:|---:|---:|---:|---:|---:|---:|
| c1_hybrid | high_high | 0.004 | 1836.860 | 24.499 | 24.328 | 50.118 | 50.118 | 0.000 |
| c1_hybrid | high_medium | 0.004 | 1836.860 | 19.817 | 19.778 | 40.452 | 40.452 | 0.000 |
| c1_hybrid | low_low | 0.004 | 1836.860 | 9.407 | 9.363 | 26.728 | 26.728 | 0.000 |
| c1_hybrid | medium_low | 0.004 | 1836.860 | 9.504 | 9.430 | 19.327 | 19.327 | 0.000 |
| c1_hybrid | medium_medium | 0.004 | 1836.860 | 15.933 | 15.627 | 32.329 | 32.329 | 0.000 |
| c1_hybrid | n1_regression | 0.004 | 1836.860 | 0.090 | 0.018 | 0.610 | 0.610 | 0.000 |
| c1_hybrid | n3_regression | 0.004 | 1836.860 | 28.860 | 30.294 | 59.394 | 59.394 | 0.000 |
| c1_hybrid | zero_result | 0.004 | 1836.860 | 9.909 | 16.802 | 27.384 | 27.384 | 0.000 |
| c1_log | high_high | 0.005 | 1816.787 | 25.456 | 20.765 | 47.207 | 47.207 | 0.000 |
| c1_log | high_medium | 0.005 | 1816.787 | 16.837 | 16.374 | 33.824 | 33.824 | 0.000 |
| c1_log | low_low | 0.005 | 1816.787 | 7.127 | 7.034 | 14.301 | 14.301 | 0.000 |
| c1_log | medium_low | 0.005 | 1816.787 | 7.196 | 7.010 | 14.404 | 14.404 | 0.000 |
| c1_log | medium_medium | 0.005 | 1816.787 | 12.613 | 12.495 | 25.492 | 25.492 | 0.000 |
| c1_log | n1_regression | 0.005 | 1816.787 | 0.087 | 0.009 | 0.566 | 0.566 | 0.000 |
| c1_log | n3_regression | 0.005 | 1816.787 | 25.364 | 25.008 | 50.666 | 50.666 | 0.000 |
| c1_log | zero_result | 0.005 | 1816.787 | 7.125 | 7.245 | 14.590 | 14.590 | 0.000 |
| c1_power2 | high_high | 0.004 | 1832.552 | 31.178 | 31.123 | 63.895 | 63.895 | 0.000 |
| c1_power2 | high_medium | 0.004 | 1832.552 | 12.862 | 12.544 | 25.641 | 25.641 | 0.000 |
| c1_power2 | low_low | 0.004 | 1832.552 | 6.329 | 6.201 | 12.645 | 12.645 | 0.000 |
| c1_power2 | medium_low | 0.004 | 1832.552 | 6.400 | 6.202 | 12.726 | 12.726 | 0.000 |
| c1_power2 | medium_medium | 0.004 | 1832.552 | 12.807 | 12.481 | 25.692 | 25.692 | 0.000 |
| c1_power2 | n1_regression | 0.004 | 1832.552 | 0.122 | 0.010 | 0.530 | 0.530 | 0.000 |
| c1_power2 | n3_regression | 0.004 | 1832.552 | 25.421 | 25.590 | 51.256 | 51.256 | 0.000 |
| c1_power2 | zero_result | 0.004 | 1832.552 | 6.341 | 6.242 | 12.686 | 12.686 | 0.000 |
| c1_sqrt | high_high | 0.005 | 1856.341 | 21.731 | 20.543 | 43.436 | 43.436 | 0.000 |
| c1_sqrt | high_medium | 0.005 | 1856.341 | 18.700 | 16.909 | 36.290 | 36.290 | 0.000 |
| c1_sqrt | low_low | 0.005 | 1856.341 | 7.114 | 7.063 | 14.488 | 14.488 | 0.000 |
| c1_sqrt | medium_low | 0.005 | 1856.341 | 7.262 | 7.292 | 14.852 | 14.852 | 0.000 |
| c1_sqrt | medium_medium | 0.005 | 1856.341 | 13.106 | 12.582 | 26.332 | 26.332 | 0.000 |
| c1_sqrt | n1_regression | 0.005 | 1856.341 | 0.098 | 0.011 | 0.549 | 0.549 | 0.000 |
| c1_sqrt | n3_regression | 0.005 | 1856.341 | 36.831 | 25.377 | 62.450 | 62.450 | 0.000 |
| c1_sqrt | zero_result | 0.005 | 1856.341 | 7.014 | 7.079 | 14.419 | 14.419 | 0.000 |
| corrected_fdxt | high_high | 6.914 | 1827.712 | 17.332 | 17.156 | 34.896 | 34.896 | 0.000 |
| corrected_fdxt | high_medium | 6.914 | 1827.712 | 12.768 | 12.471 | 25.501 | 25.501 | 0.000 |
| corrected_fdxt | low_low | 6.914 | 1827.712 | 4.816 | 4.735 | 9.818 | 9.818 | 0.000 |
| corrected_fdxt | medium_low | 6.914 | 1827.712 | 4.739 | 4.697 | 9.479 | 9.479 | 0.000 |
| corrected_fdxt | medium_medium | 6.914 | 1827.712 | 9.437 | 9.374 | 19.028 | 19.028 | 0.000 |
| corrected_fdxt | n1_regression | 6.914 | 1827.712 | 0.140 | 0.010 | 0.528 | 0.528 | 0.000 |
| corrected_fdxt | n3_regression | 6.914 | 1827.712 | 36.888 | 29.820 | 67.233 | 67.233 | 0.000 |
| corrected_fdxt | zero_result | 6.914 | 1827.712 | 10.120 | 8.652 | 18.831 | 18.831 | 0.000 |

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
