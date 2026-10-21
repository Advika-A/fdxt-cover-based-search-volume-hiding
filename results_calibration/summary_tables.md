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
| c1_hybrid | high_high | 0.006 | 1853.793 | 24.606 | 24.427 | 50.212 | 50.212 | 0.000 |
| c1_hybrid | high_medium | 0.006 | 1853.793 | 22.200 | 22.880 | 46.342 | 46.342 | 0.000 |
| c1_hybrid | low_low | 0.006 | 1853.793 | 9.609 | 9.411 | 19.397 | 19.397 | 0.000 |
| c1_hybrid | medium_low | 0.006 | 1853.793 | 10.208 | 10.403 | 21.346 | 21.346 | 0.000 |
| c1_hybrid | medium_medium | 0.006 | 1853.793 | 18.177 | 18.065 | 37.220 | 37.220 | 0.000 |
| c1_hybrid | n1_regression | 0.006 | 1853.793 | 0.102 | 0.006 | 0.500 | 0.500 | 0.000 |
| c1_hybrid | n3_regression | 0.006 | 1853.793 | 27.593 | 26.046 | 53.896 | 53.896 | 0.000 |
| c1_hybrid | zero_result | 0.006 | 1853.793 | 9.744 | 9.485 | 19.696 | 19.696 | 0.000 |
| c1_log | high_high | 0.005 | 1876.476 | 20.802 | 20.416 | 42.067 | 42.067 | 0.000 |
| c1_log | high_medium | 0.005 | 1876.476 | 17.930 | 19.197 | 37.844 | 37.844 | 0.000 |
| c1_log | low_low | 0.005 | 1876.476 | 7.586 | 7.463 | 15.545 | 15.545 | 0.000 |
| c1_log | medium_low | 0.005 | 1876.476 | 7.445 | 7.022 | 14.724 | 14.724 | 0.000 |
| c1_log | medium_medium | 0.005 | 1876.476 | 14.426 | 13.869 | 28.969 | 28.969 | 0.000 |
| c1_log | n1_regression | 0.005 | 1876.476 | 0.107 | 0.006 | 0.478 | 0.478 | 0.000 |
| c1_log | n3_regression | 0.005 | 1876.476 | 31.208 | 31.989 | 63.741 | 63.741 | 0.000 |
| c1_log | zero_result | 0.005 | 1876.476 | 7.103 | 7.058 | 14.294 | 14.294 | 0.000 |
| c1_power2 | high_high | 0.005 | 1862.515 | 25.761 | 24.869 | 51.816 | 51.816 | 0.000 |
| c1_power2 | high_medium | 0.005 | 1862.515 | 14.510 | 14.376 | 29.549 | 29.549 | 0.000 |
| c1_power2 | low_low | 0.005 | 1862.515 | 7.447 | 6.652 | 14.297 | 14.297 | 0.000 |
| c1_power2 | medium_low | 0.005 | 1862.515 | 6.314 | 6.290 | 12.753 | 12.753 | 0.000 |
| c1_power2 | medium_medium | 0.005 | 1862.515 | 16.605 | 14.524 | 31.885 | 31.885 | 0.000 |
| c1_power2 | n1_regression | 0.005 | 1862.515 | 0.106 | 0.009 | 0.484 | 0.484 | 0.000 |
| c1_power2 | n3_regression | 0.005 | 1862.515 | 26.329 | 30.359 | 57.156 | 57.156 | 0.000 |
| c1_power2 | zero_result | 0.005 | 1862.515 | 6.635 | 7.153 | 14.364 | 14.364 | 0.000 |
| c1_sqrt | high_high | 0.006 | 1801.790 | 20.277 | 20.216 | 41.260 | 41.260 | 0.000 |
| c1_sqrt | high_medium | 0.006 | 1801.790 | 17.454 | 18.124 | 36.352 | 36.352 | 0.000 |
| c1_sqrt | low_low | 0.006 | 1801.790 | 7.290 | 7.037 | 14.628 | 14.628 | 0.000 |
| c1_sqrt | medium_low | 0.006 | 1801.790 | 7.107 | 7.041 | 14.377 | 14.377 | 0.000 |
| c1_sqrt | medium_medium | 0.006 | 1801.790 | 12.568 | 12.493 | 25.471 | 25.471 | 0.000 |
| c1_sqrt | n1_regression | 0.006 | 1801.790 | 0.121 | 0.099 | 0.757 | 0.757 | 0.000 |
| c1_sqrt | n3_regression | 0.006 | 1801.790 | 26.809 | 25.308 | 52.365 | 52.365 | 0.000 |
| c1_sqrt | zero_result | 0.006 | 1801.790 | 7.138 | 7.007 | 14.346 | 14.346 | 0.000 |
| corrected_fdxt | high_high | 7.622 | 1820.086 | 16.402 | 16.323 | 33.100 | 33.100 | 0.000 |
| corrected_fdxt | high_medium | 7.622 | 1820.086 | 13.712 | 12.430 | 30.510 | 30.510 | 0.000 |
| corrected_fdxt | low_low | 7.622 | 1820.086 | 4.668 | 4.849 | 9.564 | 9.564 | 0.000 |
| corrected_fdxt | medium_low | 7.622 | 1820.086 | 5.159 | 4.888 | 10.096 | 10.096 | 0.000 |
| corrected_fdxt | medium_medium | 7.622 | 1820.086 | 10.999 | 10.187 | 21.350 | 21.350 | 0.000 |
| corrected_fdxt | n1_regression | 7.622 | 1820.086 | 0.098 | 0.005 | 0.492 | 0.492 | 0.000 |
| corrected_fdxt | n3_regression | 7.622 | 1820.086 | 26.275 | 25.351 | 51.852 | 51.852 | 0.000 |
| corrected_fdxt | zero_result | 7.622 | 1820.086 | 4.887 | 4.685 | 9.619 | 9.619 | 0.000 |

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
