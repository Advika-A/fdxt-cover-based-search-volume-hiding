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
| c1_hybrid | high_high | 0.008 | 2328.764 | 42.339 | 29.638 | 73.497 | 73.497 | 0.000 |
| c1_hybrid | high_medium | 0.008 | 2328.764 | 37.932 | 27.117 | 66.101 | 66.101 | 0.000 |
| c1_hybrid | low_low | 0.008 | 2328.764 | 13.175 | 12.780 | 26.884 | 26.884 | 0.000 |
| c1_hybrid | medium_low | 0.008 | 2328.764 | 18.870 | 12.700 | 32.273 | 32.273 | 0.000 |
| c1_hybrid | medium_medium | 0.008 | 2328.764 | 22.260 | 19.002 | 42.122 | 42.122 | 0.000 |
| c1_hybrid | n1_regression | 0.008 | 2328.764 | 0.141 | 0.007 | 0.606 | 0.606 | 0.000 |
| c1_hybrid | n3_regression | 0.008 | 2328.764 | 38.239 | 44.726 | 83.406 | 83.406 | 0.000 |
| c1_hybrid | zero_result | 0.008 | 2328.764 | 15.639 | 17.884 | 34.611 | 34.611 | 0.000 |
| c1_log | high_high | 0.008 | 2247.044 | 24.915 | 25.123 | 51.272 | 51.272 | 0.000 |
| c1_log | high_medium | 0.008 | 2247.044 | 35.508 | 32.527 | 71.372 | 71.372 | 0.000 |
| c1_log | low_low | 0.008 | 2247.044 | 14.133 | 14.487 | 28.883 | 28.883 | 0.000 |
| c1_log | medium_low | 0.008 | 2247.044 | 13.342 | 13.571 | 27.333 | 27.333 | 0.000 |
| c1_log | medium_medium | 0.008 | 2247.044 | 15.881 | 15.576 | 32.001 | 32.001 | 0.000 |
| c1_log | n1_regression | 0.008 | 2247.044 | 0.261 | 0.008 | 0.962 | 0.962 | 0.000 |
| c1_log | n3_regression | 0.008 | 2247.044 | 33.845 | 32.491 | 66.669 | 66.669 | 0.000 |
| c1_log | zero_result | 0.008 | 2247.044 | 26.379 | 22.537 | 49.570 | 49.570 | 0.000 |
| c1_power2 | high_high | 0.009 | 2370.902 | 38.955 | 39.376 | 80.156 | 80.156 | 0.000 |
| c1_power2 | high_medium | 0.009 | 2370.902 | 17.935 | 17.919 | 36.140 | 36.140 | 0.000 |
| c1_power2 | low_low | 0.009 | 2370.902 | 7.593 | 7.520 | 15.351 | 15.351 | 0.000 |
| c1_power2 | medium_low | 0.009 | 2370.902 | 7.975 | 7.488 | 15.641 | 15.641 | 0.000 |
| c1_power2 | medium_medium | 0.009 | 2370.902 | 15.678 | 15.579 | 31.759 | 31.759 | 0.000 |
| c1_power2 | n1_regression | 0.009 | 2370.902 | 0.136 | 0.008 | 0.701 | 0.701 | 0.000 |
| c1_power2 | n3_regression | 0.009 | 2370.902 | 43.243 | 40.654 | 84.210 | 84.210 | 0.000 |
| c1_power2 | zero_result | 0.009 | 2370.902 | 11.695 | 11.634 | 23.768 | 23.768 | 0.000 |
| c1_sqrt | high_high | 0.012 | 2594.226 | 25.017 | 25.365 | 51.143 | 51.143 | 0.000 |
| c1_sqrt | high_medium | 0.012 | 2594.226 | 25.980 | 23.542 | 50.574 | 50.574 | 0.000 |
| c1_sqrt | low_low | 0.012 | 2594.226 | 11.043 | 14.757 | 27.001 | 27.001 | 0.000 |
| c1_sqrt | medium_low | 0.012 | 2594.226 | 8.708 | 8.720 | 17.742 | 17.742 | 0.000 |
| c1_sqrt | medium_medium | 0.012 | 2594.226 | 16.685 | 15.194 | 32.402 | 32.402 | 0.000 |
| c1_sqrt | n1_regression | 0.012 | 2594.226 | 0.122 | 0.007 | 1.058 | 1.058 | 0.000 |
| c1_sqrt | n3_regression | 0.012 | 2594.226 | 44.944 | 50.513 | 96.100 | 96.100 | 0.000 |
| c1_sqrt | zero_result | 0.012 | 2594.226 | 10.778 | 9.531 | 20.659 | 20.659 | 0.000 |
| corrected_fdxt | high_high | 5.112 | 2393.588 | 20.259 | 20.369 | 41.459 | 41.459 | 0.000 |
| corrected_fdxt | high_medium | 5.112 | 2393.588 | 23.535 | 21.276 | 45.484 | 45.484 | 0.000 |
| corrected_fdxt | low_low | 5.112 | 2393.588 | 7.131 | 6.093 | 13.282 | 13.282 | 0.000 |
| corrected_fdxt | medium_low | 5.112 | 2393.588 | 6.424 | 7.228 | 13.774 | 13.774 | 0.000 |
| corrected_fdxt | medium_medium | 5.112 | 2393.588 | 16.745 | 14.078 | 31.008 | 31.008 | 0.000 |
| corrected_fdxt | n1_regression | 5.112 | 2393.588 | 0.153 | 0.008 | 0.771 | 0.771 | 0.000 |
| corrected_fdxt | n3_regression | 5.112 | 2393.588 | 31.080 | 30.168 | 61.545 | 61.545 | 0.000 |
| corrected_fdxt | zero_result | 5.112 | 2393.588 | 7.095 | 7.153 | 14.321 | 14.321 | 0.000 |

## Table 3 — Communication and storage

| Variant | Workload | Mean communication bytes | Mean storage bytes |
|---|---|---:|---:|
| c1_hybrid | high_high | 11375.0 | 589674.0 |
| c1_hybrid | high_medium | 9495.0 | 592014.0 |
| c1_hybrid | low_low | 3831.0 | 594718.0 |
| c1_hybrid | medium_low | 4391.0 | 594094.0 |
| c1_hybrid | medium_medium | 6480.0 | 593054.0 |
| c1_hybrid | n1_regression | 2278.0 | 595394.0 |
| c1_hybrid | n3_regression | 8045.0 | 596642.0 |
| c1_hybrid | zero_result | 3940.0 | 595394.0 |
| c1_log | high_high | 9780.0 | 589414.0 |
| c1_log | high_medium | 8203.0 | 591546.0 |
| c1_log | low_low | 2877.0 | 593730.0 |
| c1_log | medium_low | 3437.0 | 593262.0 |
| c1_log | medium_medium | 5188.0 | 592378.0 |
| c1_log | n1_regression | 2278.0 | 594250.0 |
| c1_log | n3_regression | 8045.0 | 595498.0 |
| c1_log | zero_result | 2977.0 | 594250.0 |
| c1_power2 | high_high | 11694.0 | 589726.0 |
| c1_power2 | high_medium | 6588.0 | 591598.0 |
| c1_power2 | low_low | 2559.0 | 593678.0 |
| c1_power2 | medium_low | 3119.0 | 593262.0 |
| c1_power2 | medium_medium | 5188.0 | 592430.0 |
| c1_power2 | n1_regression | 2278.0 | 594146.0 |
| c1_power2 | n3_regression | 8045.0 | 595394.0 |
| c1_power2 | zero_result | 2656.0 | 594146.0 |
| c1_sqrt | high_high | 9780.0 | 589414.0 |
| c1_sqrt | high_medium | 7880.0 | 591494.0 |
| c1_sqrt | low_low | 2877.0 | 593678.0 |
| c1_sqrt | medium_low | 3437.0 | 593210.0 |
| c1_sqrt | medium_medium | 5188.0 | 592326.0 |
| c1_sqrt | n1_regression | 2278.0 | 594198.0 |
| c1_sqrt | n3_regression | 8045.0 | 595446.0 |
| c1_sqrt | zero_result | 2977.0 | 594198.0 |
| corrected_fdxt | high_high | 8185.0 | 589154.0 |
| corrected_fdxt | high_medium | 6588.0 | 591026.0 |
| corrected_fdxt | low_low | 1923.0 | 592690.0 |
| corrected_fdxt | medium_low | 2483.0 | 592378.0 |
| corrected_fdxt | medium_medium | 3896.0 | 591650.0 |
| corrected_fdxt | n1_regression | 2278.0 | 593054.0 |
| corrected_fdxt | n3_regression | 8045.0 | 594302.0 |
| corrected_fdxt | zero_result | 2014.0 | 593054.0 |

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
