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
| c1_hybrid | pair_001 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_002 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_003 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_004 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_005 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_006 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_007 | 3 | 4.00 | 4.00 | true |
| c1_hybrid | pair_008 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_009 | 3 | 12.00 | 12.00 | true |
| c1_hybrid | pair_010 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_011 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_012 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_013 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_014 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_015 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_016 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_017 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_018 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_019 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_020 | 3 | 9.00 | 9.00 | true |
| c1_hybrid | pair_021 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_022 | 3 | 10.00 | 10.00 | true |
| c1_hybrid | pair_023 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_024 | 3 | 4.00 | 4.00 | true |
| c1_hybrid | pair_025 | 3 | 7.00 | 7.00 | true |
| c1_hybrid | pair_026 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_027 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_028 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_029 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_030 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_031 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_032 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_033 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_034 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_035 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_036 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_037 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_038 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_039 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_040 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_041 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_042 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_043 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_044 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_045 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_046 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_047 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_048 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_049 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_050 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_051 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_052 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_053 | 3 | 4.00 | 4.00 | true |
| c1_hybrid | pair_054 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_055 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_056 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_057 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_058 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_059 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_060 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_061 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_062 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_063 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_064 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_065 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_066 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_067 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_068 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_069 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_070 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_071 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_072 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_073 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_074 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_075 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_076 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_077 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_078 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_079 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_080 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_081 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_082 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_083 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_084 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_085 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_086 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_087 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_088 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_089 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_090 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_091 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_092 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_093 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_094 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_095 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_096 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_097 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_098 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_099 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_100 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_101 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_102 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_103 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_104 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_105 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_106 | 3 | 4.00 | 4.00 | true |
| c1_hybrid | pair_107 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_108 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_109 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_110 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_111 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_112 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_113 | 3 | 4.00 | 4.00 | true |
| c1_hybrid | pair_114 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_115 | 3 | 4.00 | 4.00 | true |
| c1_hybrid | pair_116 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_117 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_118 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_119 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_120 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_121 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_122 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_123 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_124 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_125 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_126 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_127 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_128 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_129 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_130 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_131 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_132 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_133 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_134 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_135 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_136 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_137 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_138 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_139 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_140 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_141 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_142 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_143 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_144 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_145 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_146 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_147 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_148 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_149 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_150 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_151 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_152 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_153 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_154 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_155 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_156 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_157 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_158 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_159 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_160 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_161 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_162 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_163 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_164 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_165 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_166 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_167 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_168 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_169 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_170 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_171 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_172 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_173 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_174 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_175 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_176 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_177 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_178 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_179 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_180 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_181 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_182 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_183 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_184 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_185 | 3 | 4.00 | 4.00 | true |
| c1_hybrid | pair_186 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_187 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_188 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_189 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_190 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_191 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_192 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_193 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_194 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_195 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_196 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_197 | 3 | 9.00 | 9.00 | true |
| c1_hybrid | pair_198 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_199 | 3 | 10.00 | 10.00 | true |
| c1_hybrid | pair_200 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_201 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_202 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_203 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_204 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_205 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_206 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_207 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_208 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_209 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_210 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_211 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_212 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_213 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_214 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_215 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_216 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_217 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_218 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_219 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_220 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_221 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_222 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_223 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_224 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_225 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_226 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_227 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_228 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_229 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_230 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_231 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_232 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_233 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_234 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_235 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_236 | 3 | 7.00 | 7.00 | true |
| c1_hybrid | pair_237 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_238 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_239 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_240 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_241 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_242 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_243 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_244 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_245 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_246 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_247 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_248 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_249 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_250 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_251 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_252 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_253 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_254 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_255 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_256 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_257 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_258 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_259 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_260 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_261 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_262 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_263 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_264 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_265 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_266 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_267 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_268 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_269 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_270 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_271 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_272 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_273 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_274 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_275 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_276 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_277 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_278 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_279 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_280 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_281 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_282 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_283 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_284 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_285 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_286 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_287 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_288 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_289 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_290 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_291 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_292 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_293 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_294 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_295 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_296 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_297 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_298 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_299 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_300 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_301 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_302 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_303 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_304 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_305 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_306 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_307 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_308 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_309 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_310 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_311 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_312 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_313 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_314 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_315 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_316 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_317 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_318 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_319 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_320 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_321 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_322 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_323 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_324 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_325 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_326 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_327 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_328 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_329 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_330 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_331 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_332 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_333 | 3 | 4.00 | 4.00 | true |
| c1_hybrid | pair_334 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_335 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_336 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_337 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_338 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_339 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_340 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_341 | 3 | 4.00 | 4.00 | true |
| c1_hybrid | pair_342 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_343 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_344 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_345 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_346 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_347 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_348 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_349 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_350 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_351 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_352 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_353 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_354 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_355 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_356 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_357 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_358 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_359 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_360 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_361 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_362 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_363 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_364 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_365 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_366 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_367 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_368 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_369 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_370 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_371 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_372 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_373 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_374 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_375 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_376 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_377 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_378 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_379 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_380 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_381 | 3 | 4.00 | 4.00 | true |
| c1_hybrid | pair_382 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_383 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_384 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_385 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_386 | 3 | 6.00 | 6.00 | true |
| c1_hybrid | pair_387 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_388 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_389 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_390 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_391 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_392 | 3 | 10.00 | 10.00 | true |
| c1_hybrid | pair_393 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_394 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_395 | 3 | 11.00 | 11.00 | true |
| c1_hybrid | pair_396 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_397 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_398 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_399 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_400 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_401 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_402 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_403 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_404 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_405 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_406 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_407 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_408 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_409 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_410 | 3 | 9.00 | 9.00 | true |
| c1_hybrid | pair_411 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_412 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_413 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_414 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_415 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_416 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_417 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_418 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_419 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_420 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_421 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_422 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_423 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_424 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_425 | 3 | 5.00 | 5.00 | true |
| c1_hybrid | pair_426 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_427 | 3 | 3.00 | 3.00 | true |
| c1_hybrid | pair_428 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_429 | 3 | 2.00 | 2.00 | true |
| c1_hybrid | pair_430 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_431 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_432 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_433 | 3 | 0.00 | 0.00 | true |
| c1_hybrid | pair_434 | 3 | 1.00 | 1.00 | true |
| c1_hybrid | pair_435 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_001 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_002 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_003 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_004 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_005 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_006 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_007 | 3 | 4.00 | 4.00 | true |
| c1_log | pair_008 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_009 | 3 | 12.00 | 12.00 | true |
| c1_log | pair_010 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_011 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_012 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_013 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_014 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_015 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_016 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_017 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_018 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_019 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_020 | 3 | 9.00 | 9.00 | true |
| c1_log | pair_021 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_022 | 3 | 10.00 | 10.00 | true |
| c1_log | pair_023 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_024 | 3 | 4.00 | 4.00 | true |
| c1_log | pair_025 | 3 | 7.00 | 7.00 | true |
| c1_log | pair_026 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_027 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_028 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_029 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_030 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_031 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_032 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_033 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_034 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_035 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_036 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_037 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_038 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_039 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_040 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_041 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_042 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_043 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_044 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_045 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_046 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_047 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_048 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_049 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_050 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_051 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_052 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_053 | 3 | 4.00 | 4.00 | true |
| c1_log | pair_054 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_055 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_056 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_057 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_058 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_059 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_060 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_061 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_062 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_063 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_064 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_065 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_066 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_067 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_068 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_069 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_070 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_071 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_072 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_073 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_074 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_075 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_076 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_077 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_078 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_079 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_080 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_081 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_082 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_083 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_084 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_085 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_086 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_087 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_088 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_089 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_090 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_091 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_092 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_093 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_094 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_095 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_096 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_097 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_098 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_099 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_100 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_101 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_102 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_103 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_104 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_105 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_106 | 3 | 4.00 | 4.00 | true |
| c1_log | pair_107 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_108 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_109 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_110 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_111 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_112 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_113 | 3 | 4.00 | 4.00 | true |
| c1_log | pair_114 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_115 | 3 | 4.00 | 4.00 | true |
| c1_log | pair_116 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_117 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_118 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_119 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_120 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_121 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_122 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_123 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_124 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_125 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_126 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_127 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_128 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_129 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_130 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_131 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_132 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_133 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_134 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_135 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_136 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_137 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_138 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_139 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_140 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_141 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_142 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_143 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_144 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_145 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_146 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_147 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_148 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_149 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_150 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_151 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_152 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_153 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_154 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_155 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_156 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_157 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_158 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_159 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_160 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_161 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_162 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_163 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_164 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_165 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_166 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_167 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_168 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_169 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_170 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_171 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_172 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_173 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_174 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_175 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_176 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_177 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_178 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_179 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_180 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_181 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_182 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_183 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_184 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_185 | 3 | 4.00 | 4.00 | true |
| c1_log | pair_186 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_187 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_188 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_189 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_190 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_191 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_192 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_193 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_194 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_195 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_196 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_197 | 3 | 9.00 | 9.00 | true |
| c1_log | pair_198 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_199 | 3 | 10.00 | 10.00 | true |
| c1_log | pair_200 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_201 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_202 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_203 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_204 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_205 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_206 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_207 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_208 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_209 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_210 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_211 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_212 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_213 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_214 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_215 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_216 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_217 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_218 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_219 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_220 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_221 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_222 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_223 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_224 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_225 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_226 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_227 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_228 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_229 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_230 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_231 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_232 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_233 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_234 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_235 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_236 | 3 | 7.00 | 7.00 | true |
| c1_log | pair_237 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_238 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_239 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_240 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_241 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_242 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_243 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_244 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_245 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_246 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_247 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_248 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_249 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_250 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_251 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_252 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_253 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_254 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_255 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_256 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_257 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_258 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_259 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_260 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_261 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_262 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_263 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_264 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_265 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_266 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_267 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_268 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_269 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_270 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_271 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_272 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_273 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_274 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_275 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_276 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_277 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_278 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_279 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_280 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_281 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_282 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_283 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_284 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_285 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_286 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_287 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_288 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_289 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_290 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_291 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_292 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_293 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_294 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_295 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_296 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_297 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_298 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_299 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_300 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_301 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_302 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_303 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_304 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_305 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_306 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_307 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_308 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_309 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_310 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_311 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_312 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_313 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_314 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_315 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_316 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_317 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_318 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_319 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_320 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_321 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_322 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_323 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_324 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_325 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_326 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_327 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_328 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_329 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_330 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_331 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_332 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_333 | 3 | 4.00 | 4.00 | true |
| c1_log | pair_334 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_335 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_336 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_337 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_338 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_339 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_340 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_341 | 3 | 4.00 | 4.00 | true |
| c1_log | pair_342 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_343 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_344 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_345 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_346 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_347 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_348 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_349 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_350 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_351 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_352 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_353 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_354 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_355 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_356 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_357 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_358 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_359 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_360 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_361 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_362 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_363 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_364 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_365 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_366 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_367 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_368 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_369 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_370 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_371 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_372 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_373 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_374 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_375 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_376 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_377 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_378 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_379 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_380 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_381 | 3 | 4.00 | 4.00 | true |
| c1_log | pair_382 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_383 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_384 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_385 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_386 | 3 | 6.00 | 6.00 | true |
| c1_log | pair_387 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_388 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_389 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_390 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_391 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_392 | 3 | 10.00 | 10.00 | true |
| c1_log | pair_393 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_394 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_395 | 3 | 11.00 | 11.00 | true |
| c1_log | pair_396 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_397 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_398 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_399 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_400 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_401 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_402 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_403 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_404 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_405 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_406 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_407 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_408 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_409 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_410 | 3 | 9.00 | 9.00 | true |
| c1_log | pair_411 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_412 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_413 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_414 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_415 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_416 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_417 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_418 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_419 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_420 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_421 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_422 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_423 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_424 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_425 | 3 | 5.00 | 5.00 | true |
| c1_log | pair_426 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_427 | 3 | 3.00 | 3.00 | true |
| c1_log | pair_428 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_429 | 3 | 2.00 | 2.00 | true |
| c1_log | pair_430 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_431 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_432 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_433 | 3 | 0.00 | 0.00 | true |
| c1_log | pair_434 | 3 | 1.00 | 1.00 | true |
| c1_log | pair_435 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_001 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_002 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_003 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_004 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_005 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_006 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_007 | 3 | 4.00 | 4.00 | true |
| c1_power2 | pair_008 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_009 | 3 | 12.00 | 12.00 | true |
| c1_power2 | pair_010 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_011 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_012 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_013 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_014 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_015 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_016 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_017 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_018 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_019 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_020 | 3 | 9.00 | 9.00 | true |
| c1_power2 | pair_021 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_022 | 3 | 10.00 | 10.00 | true |
| c1_power2 | pair_023 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_024 | 3 | 4.00 | 4.00 | true |
| c1_power2 | pair_025 | 3 | 7.00 | 7.00 | true |
| c1_power2 | pair_026 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_027 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_028 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_029 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_030 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_031 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_032 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_033 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_034 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_035 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_036 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_037 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_038 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_039 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_040 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_041 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_042 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_043 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_044 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_045 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_046 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_047 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_048 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_049 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_050 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_051 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_052 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_053 | 3 | 4.00 | 4.00 | true |
| c1_power2 | pair_054 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_055 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_056 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_057 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_058 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_059 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_060 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_061 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_062 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_063 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_064 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_065 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_066 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_067 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_068 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_069 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_070 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_071 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_072 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_073 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_074 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_075 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_076 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_077 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_078 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_079 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_080 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_081 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_082 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_083 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_084 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_085 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_086 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_087 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_088 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_089 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_090 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_091 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_092 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_093 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_094 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_095 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_096 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_097 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_098 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_099 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_100 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_101 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_102 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_103 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_104 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_105 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_106 | 3 | 4.00 | 4.00 | true |
| c1_power2 | pair_107 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_108 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_109 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_110 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_111 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_112 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_113 | 3 | 4.00 | 4.00 | true |
| c1_power2 | pair_114 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_115 | 3 | 4.00 | 4.00 | true |
| c1_power2 | pair_116 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_117 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_118 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_119 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_120 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_121 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_122 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_123 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_124 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_125 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_126 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_127 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_128 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_129 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_130 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_131 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_132 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_133 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_134 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_135 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_136 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_137 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_138 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_139 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_140 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_141 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_142 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_143 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_144 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_145 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_146 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_147 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_148 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_149 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_150 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_151 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_152 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_153 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_154 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_155 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_156 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_157 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_158 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_159 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_160 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_161 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_162 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_163 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_164 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_165 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_166 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_167 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_168 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_169 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_170 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_171 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_172 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_173 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_174 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_175 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_176 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_177 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_178 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_179 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_180 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_181 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_182 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_183 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_184 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_185 | 3 | 4.00 | 4.00 | true |
| c1_power2 | pair_186 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_187 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_188 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_189 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_190 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_191 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_192 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_193 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_194 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_195 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_196 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_197 | 3 | 9.00 | 9.00 | true |
| c1_power2 | pair_198 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_199 | 3 | 10.00 | 10.00 | true |
| c1_power2 | pair_200 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_201 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_202 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_203 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_204 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_205 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_206 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_207 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_208 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_209 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_210 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_211 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_212 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_213 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_214 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_215 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_216 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_217 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_218 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_219 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_220 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_221 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_222 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_223 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_224 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_225 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_226 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_227 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_228 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_229 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_230 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_231 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_232 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_233 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_234 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_235 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_236 | 3 | 7.00 | 7.00 | true |
| c1_power2 | pair_237 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_238 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_239 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_240 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_241 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_242 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_243 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_244 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_245 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_246 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_247 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_248 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_249 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_250 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_251 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_252 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_253 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_254 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_255 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_256 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_257 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_258 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_259 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_260 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_261 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_262 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_263 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_264 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_265 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_266 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_267 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_268 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_269 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_270 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_271 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_272 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_273 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_274 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_275 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_276 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_277 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_278 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_279 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_280 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_281 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_282 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_283 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_284 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_285 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_286 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_287 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_288 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_289 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_290 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_291 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_292 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_293 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_294 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_295 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_296 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_297 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_298 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_299 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_300 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_301 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_302 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_303 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_304 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_305 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_306 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_307 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_308 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_309 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_310 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_311 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_312 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_313 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_314 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_315 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_316 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_317 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_318 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_319 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_320 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_321 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_322 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_323 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_324 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_325 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_326 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_327 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_328 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_329 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_330 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_331 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_332 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_333 | 3 | 4.00 | 4.00 | true |
| c1_power2 | pair_334 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_335 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_336 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_337 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_338 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_339 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_340 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_341 | 3 | 4.00 | 4.00 | true |
| c1_power2 | pair_342 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_343 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_344 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_345 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_346 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_347 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_348 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_349 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_350 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_351 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_352 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_353 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_354 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_355 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_356 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_357 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_358 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_359 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_360 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_361 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_362 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_363 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_364 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_365 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_366 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_367 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_368 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_369 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_370 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_371 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_372 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_373 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_374 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_375 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_376 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_377 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_378 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_379 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_380 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_381 | 3 | 4.00 | 4.00 | true |
| c1_power2 | pair_382 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_383 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_384 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_385 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_386 | 3 | 6.00 | 6.00 | true |
| c1_power2 | pair_387 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_388 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_389 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_390 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_391 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_392 | 3 | 10.00 | 10.00 | true |
| c1_power2 | pair_393 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_394 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_395 | 3 | 11.00 | 11.00 | true |
| c1_power2 | pair_396 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_397 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_398 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_399 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_400 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_401 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_402 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_403 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_404 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_405 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_406 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_407 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_408 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_409 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_410 | 3 | 9.00 | 9.00 | true |
| c1_power2 | pair_411 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_412 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_413 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_414 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_415 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_416 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_417 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_418 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_419 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_420 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_421 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_422 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_423 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_424 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_425 | 3 | 5.00 | 5.00 | true |
| c1_power2 | pair_426 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_427 | 3 | 3.00 | 3.00 | true |
| c1_power2 | pair_428 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_429 | 3 | 2.00 | 2.00 | true |
| c1_power2 | pair_430 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_431 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_432 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_433 | 3 | 0.00 | 0.00 | true |
| c1_power2 | pair_434 | 3 | 1.00 | 1.00 | true |
| c1_power2 | pair_435 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_001 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_002 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_003 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_004 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_005 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_006 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_007 | 3 | 4.00 | 4.00 | true |
| c1_sqrt | pair_008 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_009 | 3 | 12.00 | 12.00 | true |
| c1_sqrt | pair_010 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_011 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_012 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_013 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_014 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_015 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_016 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_017 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_018 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_019 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_020 | 3 | 9.00 | 9.00 | true |
| c1_sqrt | pair_021 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_022 | 3 | 10.00 | 10.00 | true |
| c1_sqrt | pair_023 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_024 | 3 | 4.00 | 4.00 | true |
| c1_sqrt | pair_025 | 3 | 7.00 | 7.00 | true |
| c1_sqrt | pair_026 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_027 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_028 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_029 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_030 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_031 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_032 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_033 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_034 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_035 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_036 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_037 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_038 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_039 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_040 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_041 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_042 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_043 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_044 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_045 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_046 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_047 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_048 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_049 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_050 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_051 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_052 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_053 | 3 | 4.00 | 4.00 | true |
| c1_sqrt | pair_054 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_055 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_056 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_057 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_058 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_059 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_060 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_061 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_062 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_063 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_064 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_065 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_066 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_067 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_068 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_069 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_070 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_071 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_072 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_073 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_074 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_075 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_076 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_077 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_078 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_079 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_080 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_081 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_082 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_083 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_084 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_085 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_086 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_087 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_088 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_089 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_090 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_091 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_092 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_093 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_094 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_095 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_096 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_097 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_098 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_099 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_100 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_101 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_102 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_103 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_104 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_105 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_106 | 3 | 4.00 | 4.00 | true |
| c1_sqrt | pair_107 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_108 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_109 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_110 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_111 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_112 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_113 | 3 | 4.00 | 4.00 | true |
| c1_sqrt | pair_114 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_115 | 3 | 4.00 | 4.00 | true |
| c1_sqrt | pair_116 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_117 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_118 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_119 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_120 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_121 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_122 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_123 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_124 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_125 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_126 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_127 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_128 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_129 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_130 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_131 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_132 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_133 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_134 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_135 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_136 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_137 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_138 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_139 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_140 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_141 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_142 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_143 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_144 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_145 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_146 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_147 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_148 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_149 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_150 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_151 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_152 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_153 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_154 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_155 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_156 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_157 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_158 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_159 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_160 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_161 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_162 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_163 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_164 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_165 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_166 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_167 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_168 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_169 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_170 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_171 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_172 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_173 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_174 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_175 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_176 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_177 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_178 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_179 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_180 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_181 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_182 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_183 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_184 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_185 | 3 | 4.00 | 4.00 | true |
| c1_sqrt | pair_186 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_187 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_188 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_189 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_190 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_191 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_192 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_193 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_194 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_195 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_196 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_197 | 3 | 9.00 | 9.00 | true |
| c1_sqrt | pair_198 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_199 | 3 | 10.00 | 10.00 | true |
| c1_sqrt | pair_200 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_201 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_202 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_203 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_204 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_205 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_206 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_207 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_208 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_209 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_210 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_211 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_212 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_213 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_214 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_215 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_216 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_217 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_218 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_219 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_220 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_221 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_222 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_223 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_224 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_225 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_226 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_227 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_228 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_229 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_230 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_231 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_232 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_233 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_234 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_235 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_236 | 3 | 7.00 | 7.00 | true |
| c1_sqrt | pair_237 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_238 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_239 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_240 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_241 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_242 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_243 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_244 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_245 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_246 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_247 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_248 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_249 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_250 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_251 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_252 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_253 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_254 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_255 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_256 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_257 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_258 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_259 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_260 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_261 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_262 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_263 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_264 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_265 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_266 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_267 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_268 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_269 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_270 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_271 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_272 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_273 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_274 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_275 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_276 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_277 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_278 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_279 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_280 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_281 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_282 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_283 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_284 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_285 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_286 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_287 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_288 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_289 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_290 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_291 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_292 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_293 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_294 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_295 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_296 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_297 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_298 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_299 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_300 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_301 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_302 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_303 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_304 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_305 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_306 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_307 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_308 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_309 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_310 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_311 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_312 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_313 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_314 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_315 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_316 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_317 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_318 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_319 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_320 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_321 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_322 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_323 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_324 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_325 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_326 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_327 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_328 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_329 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_330 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_331 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_332 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_333 | 3 | 4.00 | 4.00 | true |
| c1_sqrt | pair_334 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_335 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_336 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_337 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_338 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_339 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_340 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_341 | 3 | 4.00 | 4.00 | true |
| c1_sqrt | pair_342 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_343 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_344 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_345 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_346 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_347 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_348 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_349 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_350 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_351 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_352 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_353 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_354 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_355 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_356 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_357 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_358 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_359 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_360 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_361 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_362 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_363 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_364 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_365 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_366 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_367 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_368 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_369 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_370 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_371 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_372 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_373 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_374 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_375 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_376 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_377 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_378 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_379 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_380 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_381 | 3 | 4.00 | 4.00 | true |
| c1_sqrt | pair_382 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_383 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_384 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_385 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_386 | 3 | 6.00 | 6.00 | true |
| c1_sqrt | pair_387 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_388 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_389 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_390 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_391 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_392 | 3 | 10.00 | 10.00 | true |
| c1_sqrt | pair_393 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_394 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_395 | 3 | 11.00 | 11.00 | true |
| c1_sqrt | pair_396 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_397 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_398 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_399 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_400 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_401 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_402 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_403 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_404 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_405 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_406 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_407 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_408 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_409 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_410 | 3 | 9.00 | 9.00 | true |
| c1_sqrt | pair_411 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_412 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_413 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_414 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_415 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_416 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_417 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_418 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_419 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_420 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_421 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_422 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_423 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_424 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_425 | 3 | 5.00 | 5.00 | true |
| c1_sqrt | pair_426 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_427 | 3 | 3.00 | 3.00 | true |
| c1_sqrt | pair_428 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_429 | 3 | 2.00 | 2.00 | true |
| c1_sqrt | pair_430 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_431 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_432 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_433 | 3 | 0.00 | 0.00 | true |
| c1_sqrt | pair_434 | 3 | 1.00 | 1.00 | true |
| c1_sqrt | pair_435 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_001 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_002 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_003 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_004 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_005 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_006 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_007 | 3 | 4.00 | 4.00 | true |
| corrected_fdxt | pair_008 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_009 | 3 | 12.00 | 12.00 | true |
| corrected_fdxt | pair_010 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_011 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_012 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_013 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_014 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_015 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_016 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_017 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_018 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_019 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_020 | 3 | 9.00 | 9.00 | true |
| corrected_fdxt | pair_021 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_022 | 3 | 10.00 | 10.00 | true |
| corrected_fdxt | pair_023 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_024 | 3 | 4.00 | 4.00 | true |
| corrected_fdxt | pair_025 | 3 | 7.00 | 7.00 | true |
| corrected_fdxt | pair_026 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_027 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_028 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_029 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_030 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_031 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_032 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_033 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_034 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_035 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_036 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_037 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_038 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_039 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_040 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_041 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_042 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_043 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_044 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_045 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_046 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_047 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_048 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_049 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_050 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_051 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_052 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_053 | 3 | 4.00 | 4.00 | true |
| corrected_fdxt | pair_054 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_055 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_056 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_057 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_058 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_059 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_060 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_061 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_062 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_063 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_064 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_065 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_066 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_067 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_068 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_069 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_070 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_071 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_072 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_073 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_074 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_075 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_076 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_077 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_078 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_079 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_080 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_081 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_082 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_083 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_084 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_085 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_086 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_087 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_088 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_089 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_090 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_091 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_092 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_093 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_094 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_095 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_096 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_097 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_098 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_099 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_100 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_101 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_102 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_103 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_104 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_105 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_106 | 3 | 4.00 | 4.00 | true |
| corrected_fdxt | pair_107 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_108 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_109 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_110 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_111 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_112 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_113 | 3 | 4.00 | 4.00 | true |
| corrected_fdxt | pair_114 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_115 | 3 | 4.00 | 4.00 | true |
| corrected_fdxt | pair_116 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_117 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_118 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_119 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_120 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_121 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_122 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_123 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_124 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_125 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_126 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_127 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_128 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_129 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_130 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_131 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_132 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_133 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_134 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_135 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_136 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_137 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_138 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_139 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_140 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_141 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_142 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_143 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_144 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_145 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_146 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_147 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_148 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_149 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_150 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_151 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_152 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_153 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_154 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_155 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_156 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_157 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_158 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_159 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_160 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_161 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_162 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_163 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_164 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_165 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_166 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_167 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_168 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_169 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_170 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_171 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_172 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_173 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_174 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_175 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_176 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_177 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_178 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_179 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_180 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_181 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_182 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_183 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_184 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_185 | 3 | 4.00 | 4.00 | true |
| corrected_fdxt | pair_186 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_187 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_188 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_189 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_190 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_191 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_192 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_193 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_194 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_195 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_196 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_197 | 3 | 9.00 | 9.00 | true |
| corrected_fdxt | pair_198 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_199 | 3 | 10.00 | 10.00 | true |
| corrected_fdxt | pair_200 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_201 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_202 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_203 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_204 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_205 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_206 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_207 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_208 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_209 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_210 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_211 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_212 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_213 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_214 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_215 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_216 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_217 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_218 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_219 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_220 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_221 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_222 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_223 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_224 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_225 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_226 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_227 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_228 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_229 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_230 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_231 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_232 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_233 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_234 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_235 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_236 | 3 | 7.00 | 7.00 | true |
| corrected_fdxt | pair_237 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_238 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_239 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_240 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_241 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_242 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_243 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_244 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_245 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_246 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_247 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_248 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_249 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_250 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_251 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_252 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_253 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_254 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_255 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_256 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_257 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_258 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_259 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_260 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_261 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_262 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_263 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_264 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_265 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_266 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_267 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_268 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_269 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_270 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_271 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_272 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_273 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_274 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_275 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_276 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_277 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_278 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_279 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_280 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_281 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_282 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_283 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_284 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_285 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_286 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_287 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_288 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_289 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_290 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_291 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_292 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_293 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_294 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_295 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_296 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_297 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_298 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_299 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_300 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_301 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_302 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_303 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_304 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_305 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_306 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_307 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_308 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_309 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_310 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_311 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_312 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_313 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_314 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_315 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_316 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_317 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_318 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_319 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_320 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_321 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_322 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_323 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_324 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_325 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_326 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_327 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_328 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_329 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_330 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_331 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_332 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_333 | 3 | 4.00 | 4.00 | true |
| corrected_fdxt | pair_334 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_335 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_336 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_337 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_338 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_339 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_340 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_341 | 3 | 4.00 | 4.00 | true |
| corrected_fdxt | pair_342 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_343 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_344 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_345 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_346 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_347 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_348 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_349 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_350 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_351 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_352 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_353 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_354 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_355 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_356 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_357 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_358 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_359 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_360 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_361 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_362 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_363 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_364 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_365 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_366 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_367 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_368 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_369 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_370 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_371 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_372 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_373 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_374 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_375 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_376 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_377 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_378 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_379 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_380 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_381 | 3 | 4.00 | 4.00 | true |
| corrected_fdxt | pair_382 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_383 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_384 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_385 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_386 | 3 | 6.00 | 6.00 | true |
| corrected_fdxt | pair_387 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_388 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_389 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_390 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_391 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_392 | 3 | 10.00 | 10.00 | true |
| corrected_fdxt | pair_393 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_394 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_395 | 3 | 11.00 | 11.00 | true |
| corrected_fdxt | pair_396 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_397 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_398 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_399 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_400 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_401 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_402 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_403 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_404 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_405 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_406 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_407 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_408 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_409 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_410 | 3 | 9.00 | 9.00 | true |
| corrected_fdxt | pair_411 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_412 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_413 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_414 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_415 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_416 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_417 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_418 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_419 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_420 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_421 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_422 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_423 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_424 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_425 | 3 | 5.00 | 5.00 | true |
| corrected_fdxt | pair_426 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_427 | 3 | 3.00 | 3.00 | true |
| corrected_fdxt | pair_428 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_429 | 3 | 2.00 | 2.00 | true |
| corrected_fdxt | pair_430 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_431 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_432 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_433 | 3 | 0.00 | 0.00 | true |
| corrected_fdxt | pair_434 | 3 | 1.00 | 1.00 | true |
| corrected_fdxt | pair_435 | 3 | 0.00 | 0.00 | true |

## Table 2 — Performance

| Variant | Workload | Setup ms | Update ms | Trapdoor ms | Search ms | Total ms | Median ms | Stddev ms |
|---|---|---:|---:|---:|---:|---:|---:|---:|
| c1_hybrid | pair_001 | 0.003 | 2085.589 | 12.419 | 14.206 | 27.263 | 26.724 | 2.927 |
| c1_hybrid | pair_002 | 0.003 | 2085.589 | 10.961 | 10.489 | 21.977 | 19.864 | 4.331 |
| c1_hybrid | pair_003 | 0.003 | 2085.589 | 12.774 | 10.601 | 23.881 | 21.101 | 5.146 |
| c1_hybrid | pair_004 | 0.003 | 2085.589 | 12.427 | 12.567 | 25.587 | 24.070 | 2.704 |
| c1_hybrid | pair_005 | 0.003 | 2085.589 | 5.876 | 5.751 | 17.871 | 11.414 | 11.252 |
| c1_hybrid | pair_006 | 0.003 | 2085.589 | 13.138 | 11.311 | 24.951 | 26.721 | 5.180 |
| c1_hybrid | pair_007 | 0.003 | 2085.589 | 22.590 | 22.390 | 45.852 | 42.753 | 7.929 |
| c1_hybrid | pair_008 | 0.003 | 2085.589 | 11.142 | 14.075 | 26.036 | 25.265 | 5.591 |
| c1_hybrid | pair_009 | 0.003 | 2085.589 | 32.889 | 25.414 | 59.654 | 56.922 | 11.028 |
| c1_hybrid | pair_010 | 0.003 | 2085.589 | 11.025 | 10.373 | 21.854 | 21.030 | 1.787 |
| c1_hybrid | pair_011 | 0.003 | 2085.589 | 8.749 | 8.823 | 18.033 | 17.988 | 0.469 |
| c1_hybrid | pair_012 | 0.003 | 2085.589 | 8.993 | 8.853 | 19.287 | 19.393 | 1.206 |
| c1_hybrid | pair_013 | 0.003 | 2085.589 | 24.165 | 9.926 | 34.576 | 24.187 | 20.894 |
| c1_hybrid | pair_014 | 0.003 | 2085.589 | 19.653 | 10.595 | 30.831 | 22.681 | 14.187 |
| c1_hybrid | pair_015 | 0.003 | 2085.589 | 13.347 | 17.781 | 31.606 | 36.931 | 10.712 |
| c1_hybrid | pair_016 | 0.003 | 2085.589 | 5.617 | 5.730 | 11.582 | 11.240 | 0.645 |
| c1_hybrid | pair_017 | 0.003 | 2085.589 | 5.624 | 5.631 | 11.533 | 11.328 | 0.533 |
| c1_hybrid | pair_018 | 0.003 | 2085.589 | 24.543 | 17.993 | 43.284 | 43.876 | 12.126 |
| c1_hybrid | pair_019 | 0.003 | 2085.589 | 11.069 | 10.550 | 22.002 | 22.030 | 0.959 |
| c1_hybrid | pair_020 | 0.003 | 2085.589 | 32.784 | 35.115 | 69.852 | 54.735 | 28.449 |
| c1_hybrid | pair_021 | 0.003 | 2085.589 | 15.791 | 12.809 | 29.108 | 30.413 | 7.825 |
| c1_hybrid | pair_022 | 0.003 | 2085.589 | 29.244 | 30.151 | 63.666 | 65.229 | 8.033 |
| c1_hybrid | pair_023 | 0.003 | 2085.589 | 10.348 | 10.801 | 22.180 | 22.093 | 4.612 |
| c1_hybrid | pair_024 | 0.003 | 2085.589 | 17.137 | 22.272 | 40.431 | 40.591 | 8.047 |
| c1_hybrid | pair_025 | 0.003 | 2085.589 | 30.528 | 32.862 | 82.065 | 71.519 | 34.069 |
| c1_hybrid | pair_026 | 0.003 | 2085.589 | 11.375 | 9.694 | 21.554 | 23.169 | 3.414 |
| c1_hybrid | pair_027 | 0.003 | 2085.589 | 9.770 | 10.005 | 20.245 | 18.035 | 4.372 |
| c1_hybrid | pair_028 | 0.003 | 2085.589 | 6.095 | 6.004 | 12.325 | 11.462 | 1.777 |
| c1_hybrid | pair_029 | 0.003 | 2085.589 | 11.159 | 10.838 | 22.556 | 19.765 | 5.358 |
| c1_hybrid | pair_030 | 0.003 | 2085.589 | 11.291 | 10.171 | 21.898 | 22.958 | 2.164 |
| c1_hybrid | pair_031 | 0.003 | 2085.589 | 10.473 | 10.440 | 21.343 | 20.732 | 1.063 |
| c1_hybrid | pair_032 | 0.003 | 2085.589 | 12.035 | 11.795 | 27.250 | 24.765 | 4.622 |
| c1_hybrid | pair_033 | 0.003 | 2085.589 | 7.584 | 6.218 | 14.057 | 15.233 | 2.087 |
| c1_hybrid | pair_034 | 0.003 | 2085.589 | 9.632 | 9.556 | 19.558 | 19.548 | 0.266 |
| c1_hybrid | pair_035 | 0.003 | 2085.589 | 11.928 | 11.686 | 24.178 | 24.166 | 0.066 |
| c1_hybrid | pair_036 | 0.003 | 2085.589 | 10.327 | 11.294 | 22.123 | 21.193 | 1.677 |
| c1_hybrid | pair_037 | 0.003 | 2085.589 | 12.373 | 12.109 | 24.958 | 25.071 | 0.668 |
| c1_hybrid | pair_038 | 0.003 | 2085.589 | 10.555 | 10.367 | 21.350 | 20.936 | 0.734 |
| c1_hybrid | pair_039 | 0.003 | 2085.589 | 8.973 | 8.804 | 18.207 | 18.257 | 0.089 |
| c1_hybrid | pair_040 | 0.003 | 2085.589 | 8.983 | 9.673 | 19.208 | 18.339 | 1.635 |
| c1_hybrid | pair_041 | 0.003 | 2085.589 | 10.133 | 11.433 | 30.382 | 22.310 | 16.586 |
| c1_hybrid | pair_042 | 0.003 | 2085.589 | 11.791 | 10.500 | 23.107 | 20.640 | 5.546 |
| c1_hybrid | pair_043 | 0.003 | 2085.589 | 15.501 | 10.865 | 26.877 | 27.560 | 7.161 |
| c1_hybrid | pair_044 | 0.003 | 2085.589 | 10.869 | 9.425 | 20.882 | 19.821 | 10.039 |
| c1_hybrid | pair_045 | 0.003 | 2085.589 | 6.579 | 5.989 | 12.833 | 12.389 | 1.838 |
| c1_hybrid | pair_046 | 0.003 | 2085.589 | 22.765 | 13.216 | 37.206 | 30.648 | 16.173 |
| c1_hybrid | pair_047 | 0.003 | 2085.589 | 13.973 | 13.432 | 27.896 | 22.745 | 10.631 |
| c1_hybrid | pair_048 | 0.003 | 2085.589 | 13.718 | 13.966 | 28.583 | 24.137 | 7.870 |
| c1_hybrid | pair_049 | 0.003 | 2085.589 | 10.189 | 10.211 | 20.872 | 20.809 | 0.326 |
| c1_hybrid | pair_050 | 0.003 | 2085.589 | 11.905 | 11.718 | 24.224 | 24.241 | 0.056 |
| c1_hybrid | pair_051 | 0.003 | 2085.589 | 8.656 | 8.646 | 23.661 | 17.796 | 10.569 |
| c1_hybrid | pair_052 | 0.003 | 2085.589 | 15.160 | 13.828 | 29.597 | 24.002 | 9.858 |
| c1_hybrid | pair_053 | 0.003 | 2085.589 | 12.270 | 12.417 | 25.295 | 24.294 | 1.889 |
| c1_hybrid | pair_054 | 0.003 | 2085.589 | 8.718 | 8.556 | 17.655 | 17.597 | 0.178 |
| c1_hybrid | pair_055 | 0.003 | 2085.589 | 8.637 | 8.759 | 17.800 | 17.892 | 0.235 |
| c1_hybrid | pair_056 | 0.003 | 2085.589 | 9.560 | 5.428 | 15.233 | 11.158 | 7.081 |
| c1_hybrid | pair_057 | 0.003 | 2085.589 | 9.391 | 9.464 | 19.216 | 19.176 | 0.100 |
| c1_hybrid | pair_058 | 0.003 | 2085.589 | 9.582 | 9.318 | 19.294 | 19.294 | 0.289 |
| c1_hybrid | pair_059 | 0.003 | 2085.589 | 9.407 | 9.346 | 24.001 | 19.198 | 8.552 |
| c1_hybrid | pair_060 | 0.003 | 2085.589 | 5.605 | 6.951 | 12.790 | 11.202 | 2.770 |
| c1_hybrid | pair_061 | 0.003 | 2085.589 | 10.322 | 9.993 | 20.870 | 19.192 | 3.073 |
| c1_hybrid | pair_062 | 0.003 | 2085.589 | 9.929 | 9.861 | 20.202 | 19.248 | 1.886 |
| c1_hybrid | pair_063 | 0.003 | 2085.589 | 10.767 | 13.041 | 24.121 | 21.289 | 5.015 |
| c1_hybrid | pair_064 | 0.003 | 2085.589 | 9.628 | 12.476 | 22.483 | 19.669 | 5.400 |
| c1_hybrid | pair_065 | 0.003 | 2085.589 | 9.524 | 9.326 | 19.351 | 19.128 | 0.423 |
| c1_hybrid | pair_066 | 0.003 | 2085.589 | 8.770 | 8.966 | 18.201 | 17.678 | 1.004 |
| c1_hybrid | pair_067 | 0.003 | 2085.589 | 9.517 | 9.064 | 18.921 | 17.854 | 2.213 |
| c1_hybrid | pair_068 | 0.003 | 2085.589 | 10.592 | 10.716 | 21.713 | 20.055 | 3.558 |
| c1_hybrid | pair_069 | 0.003 | 2085.589 | 10.257 | 10.317 | 21.192 | 20.663 | 2.329 |
| c1_hybrid | pair_070 | 0.003 | 2085.589 | 10.405 | 11.056 | 22.031 | 21.746 | 1.660 |
| c1_hybrid | pair_071 | 0.003 | 2085.589 | 5.923 | 5.687 | 11.973 | 11.694 | 0.803 |
| c1_hybrid | pair_072 | 0.003 | 2085.589 | 6.286 | 6.704 | 14.426 | 15.239 | 2.699 |
| c1_hybrid | pair_073 | 0.003 | 2085.589 | 13.766 | 10.056 | 24.369 | 19.640 | 8.649 |
| c1_hybrid | pair_074 | 0.003 | 2085.589 | 12.168 | 10.614 | 23.162 | 22.736 | 4.271 |
| c1_hybrid | pair_075 | 0.003 | 2085.589 | 9.962 | 9.511 | 19.862 | 19.976 | 0.380 |
| c1_hybrid | pair_076 | 0.003 | 2085.589 | 13.515 | 12.720 | 26.597 | 19.460 | 12.381 |
| c1_hybrid | pair_077 | 0.003 | 2085.589 | 14.515 | 9.572 | 24.415 | 19.204 | 9.037 |
| c1_hybrid | pair_078 | 0.003 | 2085.589 | 8.779 | 10.607 | 23.393 | 17.915 | 9.785 |
| c1_hybrid | pair_079 | 0.003 | 2085.589 | 13.039 | 10.008 | 23.669 | 22.830 | 4.863 |
| c1_hybrid | pair_080 | 0.003 | 2085.589 | 9.746 | 9.392 | 19.476 | 19.424 | 0.158 |
| c1_hybrid | pair_081 | 0.003 | 2085.589 | 8.662 | 8.662 | 17.667 | 17.700 | 0.188 |
| c1_hybrid | pair_082 | 0.003 | 2085.589 | 8.601 | 8.654 | 17.669 | 17.671 | 0.178 |
| c1_hybrid | pair_083 | 0.003 | 2085.589 | 5.506 | 5.922 | 12.405 | 12.681 | 1.188 |
| c1_hybrid | pair_084 | 0.003 | 2085.589 | 10.734 | 9.440 | 20.612 | 19.503 | 2.018 |
| c1_hybrid | pair_085 | 0.003 | 2085.589 | 10.471 | 10.158 | 21.035 | 21.189 | 0.325 |
| c1_hybrid | pair_086 | 0.003 | 2085.589 | 5.428 | 5.457 | 11.098 | 11.086 | 0.089 |
| c1_hybrid | pair_087 | 0.003 | 2085.589 | 9.498 | 10.126 | 20.084 | 19.897 | 0.988 |
| c1_hybrid | pair_088 | 0.003 | 2085.589 | 10.721 | 10.968 | 22.136 | 21.831 | 1.616 |
| c1_hybrid | pair_089 | 0.003 | 2085.589 | 10.794 | 10.190 | 21.407 | 21.380 | 0.737 |
| c1_hybrid | pair_090 | 0.003 | 2085.589 | 10.401 | 10.103 | 20.908 | 20.761 | 0.323 |
| c1_hybrid | pair_091 | 0.003 | 2085.589 | 10.195 | 10.108 | 20.714 | 20.751 | 0.142 |
| c1_hybrid | pair_092 | 0.003 | 2085.589 | 8.783 | 8.572 | 17.714 | 17.590 | 0.228 |
| c1_hybrid | pair_093 | 0.003 | 2085.589 | 8.712 | 8.567 | 17.707 | 17.650 | 0.318 |
| c1_hybrid | pair_094 | 0.003 | 2085.589 | 9.515 | 9.368 | 20.356 | 19.354 | 1.737 |
| c1_hybrid | pair_095 | 0.003 | 2085.589 | 11.131 | 9.971 | 21.680 | 21.157 | 2.372 |
| c1_hybrid | pair_096 | 0.003 | 2085.589 | 9.904 | 10.699 | 20.997 | 19.150 | 3.233 |
| c1_hybrid | pair_097 | 0.003 | 2085.589 | 6.256 | 6.082 | 12.575 | 11.170 | 2.476 |
| c1_hybrid | pair_098 | 0.003 | 2085.589 | 7.128 | 5.640 | 13.037 | 12.247 | 2.377 |
| c1_hybrid | pair_099 | 0.003 | 2085.589 | 13.337 | 10.903 | 27.969 | 20.888 | 12.268 |
| c1_hybrid | pair_100 | 0.003 | 2085.589 | 13.184 | 10.899 | 24.666 | 20.680 | 6.918 |
| c1_hybrid | pair_101 | 0.003 | 2085.589 | 10.502 | 10.579 | 21.607 | 20.668 | 1.757 |
| c1_hybrid | pair_102 | 0.003 | 2085.589 | 10.456 | 13.177 | 24.109 | 20.698 | 6.070 |
| c1_hybrid | pair_103 | 0.003 | 2085.589 | 10.282 | 10.188 | 20.901 | 20.733 | 0.407 |
| c1_hybrid | pair_104 | 0.003 | 2085.589 | 8.634 | 8.558 | 17.616 | 17.624 | 0.072 |
| c1_hybrid | pair_105 | 0.003 | 2085.589 | 10.295 | 10.244 | 20.973 | 20.747 | 0.459 |
| c1_hybrid | pair_106 | 0.003 | 2085.589 | 13.045 | 10.646 | 24.375 | 20.849 | 6.312 |
| c1_hybrid | pair_107 | 0.003 | 2085.589 | 11.766 | 8.632 | 20.811 | 18.141 | 4.819 |
| c1_hybrid | pair_108 | 0.003 | 2085.589 | 8.801 | 8.668 | 17.835 | 17.649 | 0.429 |
| c1_hybrid | pair_109 | 0.003 | 2085.589 | 5.511 | 5.498 | 11.270 | 11.199 | 0.158 |
| c1_hybrid | pair_110 | 0.003 | 2085.589 | 9.447 | 10.552 | 20.411 | 19.199 | 2.128 |
| c1_hybrid | pair_111 | 0.003 | 2085.589 | 6.303 | 5.451 | 12.111 | 11.394 | 1.460 |
| c1_hybrid | pair_112 | 0.003 | 2085.589 | 11.629 | 9.355 | 21.495 | 21.452 | 2.316 |
| c1_hybrid | pair_113 | 0.003 | 2085.589 | 14.892 | 12.227 | 28.213 | 25.821 | 5.489 |
| c1_hybrid | pair_114 | 0.003 | 2085.589 | 10.466 | 14.022 | 25.119 | 21.152 | 7.042 |
| c1_hybrid | pair_115 | 0.003 | 2085.589 | 14.616 | 13.033 | 28.441 | 24.221 | 7.612 |
| c1_hybrid | pair_116 | 0.003 | 2085.589 | 10.313 | 10.504 | 21.501 | 20.937 | 1.123 |
| c1_hybrid | pair_117 | 0.003 | 2085.589 | 12.176 | 8.667 | 23.952 | 17.908 | 10.693 |
| c1_hybrid | pair_118 | 0.003 | 2085.589 | 12.356 | 12.806 | 25.711 | 18.042 | 13.632 |
| c1_hybrid | pair_119 | 0.003 | 2085.589 | 9.974 | 11.354 | 21.957 | 19.305 | 4.678 |
| c1_hybrid | pair_120 | 0.003 | 2085.589 | 9.508 | 9.396 | 19.254 | 19.271 | 0.076 |
| c1_hybrid | pair_121 | 0.003 | 2085.589 | 9.527 | 17.610 | 27.597 | 19.729 | 13.686 |
| c1_hybrid | pair_122 | 0.003 | 2085.589 | 5.494 | 5.486 | 11.242 | 11.304 | 0.123 |
| c1_hybrid | pair_123 | 0.003 | 2085.589 | 5.507 | 5.467 | 11.186 | 11.151 | 0.091 |
| c1_hybrid | pair_124 | 0.003 | 2085.589 | 11.823 | 11.797 | 24.063 | 24.201 | 0.296 |
| c1_hybrid | pair_125 | 0.003 | 2085.589 | 10.277 | 10.148 | 20.789 | 20.871 | 0.145 |
| c1_hybrid | pair_126 | 0.003 | 2085.589 | 13.013 | 12.131 | 25.661 | 24.165 | 2.898 |
| c1_hybrid | pair_127 | 0.003 | 2085.589 | 10.223 | 10.183 | 20.892 | 20.826 | 0.149 |
| c1_hybrid | pair_128 | 0.003 | 2085.589 | 11.773 | 11.660 | 23.983 | 24.054 | 0.233 |
| c1_hybrid | pair_129 | 0.003 | 2085.589 | 8.687 | 8.793 | 17.867 | 17.978 | 0.397 |
| c1_hybrid | pair_130 | 0.003 | 2085.589 | 11.855 | 11.716 | 24.052 | 23.961 | 0.234 |
| c1_hybrid | pair_131 | 0.003 | 2085.589 | 11.919 | 11.711 | 24.093 | 23.888 | 0.357 |
| c1_hybrid | pair_132 | 0.003 | 2085.589 | 8.753 | 8.622 | 17.832 | 17.959 | 0.267 |
| c1_hybrid | pair_133 | 0.003 | 2085.589 | 8.645 | 8.910 | 18.028 | 18.169 | 0.360 |
| c1_hybrid | pair_134 | 0.003 | 2085.589 | 5.473 | 5.514 | 11.237 | 11.289 | 0.105 |
| c1_hybrid | pair_135 | 0.003 | 2085.589 | 15.783 | 10.222 | 26.354 | 19.332 | 12.495 |
| c1_hybrid | pair_136 | 0.003 | 2085.589 | 7.057 | 6.025 | 13.462 | 11.089 | 4.152 |
| c1_hybrid | pair_137 | 0.003 | 2085.589 | 6.077 | 7.089 | 13.390 | 11.383 | 3.632 |
| c1_hybrid | pair_138 | 0.003 | 2085.589 | 6.036 | 6.691 | 14.691 | 14.790 | 1.094 |
| c1_hybrid | pair_139 | 0.003 | 2085.589 | 8.515 | 6.219 | 15.361 | 14.264 | 3.556 |
| c1_hybrid | pair_140 | 0.003 | 2085.589 | 5.481 | 5.584 | 11.278 | 11.130 | 0.298 |
| c1_hybrid | pair_141 | 0.003 | 2085.589 | 5.559 | 5.540 | 11.287 | 11.166 | 0.253 |
| c1_hybrid | pair_142 | 0.003 | 2085.589 | 5.983 | 6.846 | 13.314 | 11.544 | 3.430 |
| c1_hybrid | pair_143 | 0.003 | 2085.589 | 6.491 | 5.579 | 12.313 | 11.979 | 1.368 |
| c1_hybrid | pair_144 | 0.003 | 2085.589 | 5.510 | 5.604 | 11.380 | 11.381 | 0.277 |
| c1_hybrid | pair_145 | 0.003 | 2085.589 | 5.536 | 7.577 | 13.548 | 11.308 | 3.885 |
| c1_hybrid | pair_146 | 0.003 | 2085.589 | 5.958 | 5.485 | 11.678 | 11.220 | 0.965 |
| c1_hybrid | pair_147 | 0.003 | 2085.589 | 5.486 | 5.502 | 11.219 | 11.142 | 0.244 |
| c1_hybrid | pair_148 | 0.003 | 2085.589 | 5.559 | 5.478 | 11.269 | 11.263 | 0.037 |
| c1_hybrid | pair_149 | 0.003 | 2085.589 | 5.528 | 5.429 | 11.183 | 11.215 | 0.096 |
| c1_hybrid | pair_150 | 0.003 | 2085.589 | 5.595 | 5.561 | 11.488 | 11.241 | 0.489 |
| c1_hybrid | pair_151 | 0.003 | 2085.589 | 5.492 | 6.451 | 12.394 | 11.302 | 2.182 |
| c1_hybrid | pair_152 | 0.003 | 2085.589 | 6.296 | 6.233 | 12.869 | 11.610 | 2.530 |
| c1_hybrid | pair_153 | 0.003 | 2085.589 | 7.182 | 7.468 | 14.879 | 11.383 | 6.191 |
| c1_hybrid | pair_154 | 0.003 | 2085.589 | 7.884 | 6.901 | 15.038 | 11.310 | 6.580 |
| c1_hybrid | pair_155 | 0.003 | 2085.589 | 6.562 | 6.700 | 13.563 | 13.797 | 2.197 |
| c1_hybrid | pair_156 | 0.003 | 2085.589 | 5.697 | 5.550 | 11.515 | 11.443 | 0.205 |
| c1_hybrid | pair_157 | 0.003 | 2085.589 | 5.552 | 5.527 | 11.256 | 11.225 | 0.235 |
| c1_hybrid | pair_158 | 0.003 | 2085.589 | 5.501 | 5.466 | 11.250 | 11.209 | 0.128 |
| c1_hybrid | pair_159 | 0.003 | 2085.589 | 5.515 | 5.448 | 11.252 | 11.244 | 0.131 |
| c1_hybrid | pair_160 | 0.003 | 2085.589 | 9.597 | 9.439 | 19.450 | 19.467 | 0.098 |
| c1_hybrid | pair_161 | 0.003 | 2085.589 | 9.394 | 9.534 | 19.425 | 19.348 | 0.244 |
| c1_hybrid | pair_162 | 0.003 | 2085.589 | 9.968 | 9.369 | 22.640 | 21.099 | 3.341 |
| c1_hybrid | pair_163 | 0.003 | 2085.589 | 14.812 | 9.844 | 25.149 | 19.528 | 10.021 |
| c1_hybrid | pair_164 | 0.003 | 2085.589 | 8.740 | 8.751 | 17.869 | 17.944 | 0.387 |
| c1_hybrid | pair_165 | 0.003 | 2085.589 | 8.642 | 11.715 | 20.795 | 17.871 | 5.299 |
| c1_hybrid | pair_166 | 0.003 | 2085.589 | 9.498 | 9.922 | 19.840 | 20.015 | 0.753 |
| c1_hybrid | pair_167 | 0.003 | 2085.589 | 9.630 | 9.426 | 19.453 | 19.387 | 0.302 |
| c1_hybrid | pair_168 | 0.003 | 2085.589 | 9.468 | 9.362 | 19.248 | 19.180 | 0.184 |
| c1_hybrid | pair_169 | 0.003 | 2085.589 | 5.519 | 5.491 | 11.272 | 11.227 | 0.109 |
| c1_hybrid | pair_170 | 0.003 | 2085.589 | 5.607 | 5.519 | 11.372 | 11.319 | 0.136 |
| c1_hybrid | pair_171 | 0.003 | 2085.589 | 9.568 | 9.328 | 19.335 | 19.385 | 0.165 |
| c1_hybrid | pair_172 | 0.003 | 2085.589 | 9.539 | 9.452 | 19.335 | 19.362 | 0.195 |
| c1_hybrid | pair_173 | 0.003 | 2085.589 | 11.188 | 9.362 | 20.986 | 19.434 | 2.721 |
| c1_hybrid | pair_174 | 0.003 | 2085.589 | 13.748 | 12.726 | 26.869 | 26.302 | 1.686 |
| c1_hybrid | pair_175 | 0.003 | 2085.589 | 11.026 | 11.092 | 22.599 | 19.590 | 5.381 |
| c1_hybrid | pair_176 | 0.003 | 2085.589 | 10.081 | 10.394 | 20.892 | 18.440 | 4.690 |
| c1_hybrid | pair_177 | 0.003 | 2085.589 | 11.051 | 9.613 | 21.127 | 19.374 | 3.050 |
| c1_hybrid | pair_178 | 0.003 | 2085.589 | 9.620 | 9.550 | 19.588 | 19.640 | 0.257 |
| c1_hybrid | pair_179 | 0.003 | 2085.589 | 10.438 | 9.512 | 20.767 | 18.494 | 4.564 |
| c1_hybrid | pair_180 | 0.003 | 2085.589 | 145977.858 | 22.005 | 146001.934 | 20.264 | 252849.481 |
| c1_hybrid | pair_181 | 0.003 | 2085.589 | 13.915 | 1234969.630 | 1234984.393 | 41.146 | 2139010.052 |
| c1_hybrid | pair_182 | 0.003 | 2085.589 | 345.256 | 95.947 | 461.461 | 37.611 | 750.262 |
| c1_hybrid | pair_183 | 0.003 | 2085.589 | 49.990 | 28.646 | 79.468 | 23.699 | 99.358 |
| c1_hybrid | pair_184 | 0.003 | 2085.589 | 52.528 | 63.055 | 117.103 | 42.005 | 132.031 |
| c1_hybrid | pair_185 | 0.003 | 2085.589 | 26.087 | 23.779 | 50.602 | 20.879 | 51.645 |
| c1_hybrid | pair_186 | 0.003 | 2085.589 | 17.008 | 14.940 | 32.541 | 17.570 | 25.946 |
| c1_hybrid | pair_187 | 0.003 | 2085.589 | 17.098 | 21.488 | 39.306 | 17.712 | 37.484 |
| c1_hybrid | pair_188 | 0.003 | 2085.589 | 41.866 | 26.623 | 69.139 | 19.694 | 85.761 |
| c1_hybrid | pair_189 | 0.003 | 2085.589 | 13.237 | 13.270 | 27.149 | 19.314 | 13.612 |
| c1_hybrid | pair_190 | 0.003 | 2085.589 | 13.298 | 12.871 | 26.749 | 19.097 | 13.337 |
| c1_hybrid | pair_191 | 0.003 | 2085.589 | 7.815 | 7.681 | 15.723 | 11.264 | 7.991 |
| c1_hybrid | pair_192 | 0.003 | 2085.589 | 7.636 | 7.687 | 15.741 | 11.330 | 7.736 |
| c1_hybrid | pair_193 | 0.003 | 2085.589 | 23.232 | 20.900 | 45.209 | 30.603 | 25.569 |
| c1_hybrid | pair_194 | 0.003 | 2085.589 | 14.567 | 14.188 | 29.501 | 21.064 | 14.907 |
| c1_hybrid | pair_195 | 0.003 | 2085.589 | 29.153 | 63.527 | 100.357 | 44.782 | 100.488 |
| c1_hybrid | pair_196 | 0.003 | 2085.589 | 14.853 | 14.705 | 30.715 | 23.138 | 15.268 |
| c1_hybrid | pair_197 | 0.003 | 2085.589 | 46.260 | 27.175 | 74.967 | 40.290 | 60.225 |
| c1_hybrid | pair_198 | 0.003 | 2085.589 | 12.005 | 11.954 | 24.598 | 17.498 | 12.383 |
| c1_hybrid | pair_199 | 0.003 | 2085.589 | 26.891 | 58.831 | 86.836 | 32.094 | 95.197 |
| c1_hybrid | pair_200 | 0.003 | 2085.589 | 27.091 | 26.999 | 55.578 | 40.238 | 26.895 |
| c1_hybrid | pair_201 | 0.003 | 2085.589 | 11.968 | 11.998 | 24.437 | 17.930 | 11.662 |
| c1_hybrid | pair_202 | 0.003 | 2085.589 | 18.226 | 12.038 | 31.003 | 17.948 | 22.828 |
| c1_hybrid | pair_203 | 0.003 | 2085.589 | 7.934 | 7.783 | 16.230 | 11.245 | 8.781 |
| c1_hybrid | pair_204 | 0.003 | 2085.589 | 13.559 | 13.748 | 27.996 | 19.233 | 15.242 |
| c1_hybrid | pair_205 | 0.003 | 2085.589 | 15.042 | 16.382 | 35.991 | 27.849 | 20.506 |
| c1_hybrid | pair_206 | 0.003 | 2085.589 | 19.360 | 15.458 | 35.501 | 21.037 | 25.285 |
| c1_hybrid | pair_207 | 0.003 | 2085.589 | 12.403 | 17.992 | 31.226 | 17.569 | 23.769 |
| c1_hybrid | pair_208 | 0.003 | 2085.589 | 15.160 | 13.660 | 29.455 | 17.589 | 20.567 |
| c1_hybrid | pair_209 | 0.003 | 2085.589 | 24.718 | 19.630 | 45.246 | 19.301 | 44.963 |
| c1_hybrid | pair_210 | 0.003 | 2085.589 | 14.160 | 12.196 | 28.850 | 24.630 | 8.437 |
| c1_hybrid | pair_211 | 0.003 | 2085.589 | 19.040 | 20.203 | 40.358 | 19.817 | 36.162 |
| c1_hybrid | pair_212 | 0.003 | 2085.589 | 20.574 | 19.266 | 40.582 | 11.235 | 50.853 |
| c1_hybrid | pair_213 | 0.003 | 2085.589 | 14.687 | 24.747 | 39.911 | 11.257 | 49.777 |
| c1_hybrid | pair_214 | 0.003 | 2085.589 | 13.726 | 13.921 | 28.248 | 20.702 | 13.141 |
| c1_hybrid | pair_215 | 0.003 | 2085.589 | 14.044 | 13.867 | 28.602 | 20.778 | 13.558 |
| c1_hybrid | pair_216 | 0.003 | 2085.589 | 14.007 | 13.369 | 28.041 | 20.851 | 12.611 |
| c1_hybrid | pair_217 | 0.003 | 2085.589 | 13.987 | 13.547 | 28.145 | 21.320 | 12.123 |
| c1_hybrid | pair_218 | 0.003 | 2085.589 | 14.117 | 13.978 | 28.723 | 21.238 | 13.324 |
| c1_hybrid | pair_219 | 0.003 | 2085.589 | 11.253 | 11.422 | 23.286 | 17.694 | 9.800 |
| c1_hybrid | pair_220 | 0.003 | 2085.589 | 14.017 | 13.474 | 28.232 | 20.880 | 12.891 |
| c1_hybrid | pair_221 | 0.003 | 2085.589 | 13.329 | 13.344 | 27.324 | 20.783 | 11.476 |
| c1_hybrid | pair_222 | 0.003 | 2085.589 | 11.031 | 10.477 | 22.099 | 17.651 | 7.726 |
| c1_hybrid | pair_223 | 0.003 | 2085.589 | 10.946 | 10.851 | 22.436 | 17.766 | 8.253 |
| c1_hybrid | pair_224 | 0.003 | 2085.589 | 8.346 | 9.906 | 20.690 | 24.347 | 8.330 |
| c1_hybrid | pair_225 | 0.003 | 2085.589 | 22.880 | 17.140 | 41.433 | 48.013 | 19.908 |
| c1_hybrid | pair_226 | 0.003 | 2085.589 | 14.436 | 14.143 | 29.213 | 21.801 | 13.790 |
| c1_hybrid | pair_227 | 0.003 | 2085.589 | 12.228 | 21.624 | 38.279 | 18.513 | 35.025 |
| c1_hybrid | pair_228 | 0.003 | 2085.589 | 23.079 | 14.288 | 37.938 | 17.613 | 35.246 |
| c1_hybrid | pair_229 | 0.003 | 2085.589 | 19.224 | 23.760 | 48.412 | 23.330 | 47.078 |
| c1_hybrid | pair_230 | 0.003 | 2085.589 | 27.890 | 22.361 | 51.671 | 22.965 | 53.111 |
| c1_hybrid | pair_231 | 0.003 | 2085.589 | 15.766 | 13.973 | 30.359 | 19.366 | 19.192 |
| c1_hybrid | pair_232 | 0.003 | 2085.589 | 8.713 | 9.657 | 18.953 | 13.852 | 11.096 |
| c1_hybrid | pair_233 | 0.003 | 2085.589 | 9.646 | 8.053 | 18.170 | 11.472 | 11.732 |
| c1_hybrid | pair_234 | 0.003 | 2085.589 | 25.725 | 25.190 | 52.013 | 31.230 | 36.380 |
| c1_hybrid | pair_235 | 0.003 | 2085.589 | 18.092 | 15.900 | 35.079 | 22.297 | 23.489 |
| c1_hybrid | pair_236 | 0.003 | 2085.589 | 38.249 | 37.100 | 78.070 | 50.212 | 48.376 |
| c1_hybrid | pair_237 | 0.003 | 2085.589 | 18.447 | 18.148 | 37.307 | 21.030 | 28.265 |
| c1_hybrid | pair_238 | 0.003 | 2085.589 | 41.030 | 61.709 | 105.501 | 48.653 | 98.655 |
| c1_hybrid | pair_239 | 0.003 | 2085.589 | 22.381 | 19.760 | 43.432 | 18.336 | 44.006 |
| c1_hybrid | pair_240 | 0.003 | 2085.589 | 34.253 | 30.047 | 66.511 | 34.657 | 57.134 |
| c1_hybrid | pair_241 | 0.003 | 2085.589 | 45.156 | 57.138 | 104.441 | 50.684 | 93.490 |
| c1_hybrid | pair_242 | 0.003 | 2085.589 | 15.209 | 12.502 | 28.399 | 26.332 | 11.665 |
| c1_hybrid | pair_243 | 0.003 | 2085.589 | 12.395 | 12.294 | 25.227 | 17.864 | 12.977 |
| c1_hybrid | pair_244 | 0.003 | 2085.589 | 10.646 | 11.518 | 22.745 | 11.382 | 19.892 |
| c1_hybrid | pair_245 | 0.003 | 2085.589 | 22.319 | 17.634 | 42.140 | 19.125 | 39.952 |
| c1_hybrid | pair_246 | 0.003 | 2085.589 | 16.689 | 19.154 | 37.092 | 17.712 | 33.843 |
| c1_hybrid | pair_247 | 0.003 | 2085.589 | 16.971 | 13.178 | 30.896 | 17.410 | 23.385 |
| c1_hybrid | pair_248 | 0.003 | 2085.589 | 14.903 | 13.805 | 29.432 | 21.005 | 15.999 |
| c1_hybrid | pair_249 | 0.003 | 2085.589 | 13.586 | 14.740 | 28.820 | 19.093 | 16.972 |
| c1_hybrid | pair_250 | 0.003 | 2085.589 | 13.819 | 13.504 | 28.123 | 19.105 | 15.645 |
| c1_hybrid | pair_251 | 0.003 | 2085.589 | 8.439 | 9.462 | 18.620 | 11.222 | 12.877 |
| c1_hybrid | pair_252 | 0.003 | 2085.589 | 8.590 | 8.281 | 17.449 | 11.157 | 10.926 |
| c1_hybrid | pair_253 | 0.003 | 2085.589 | 15.637 | 14.831 | 31.202 | 20.850 | 17.952 |
| c1_hybrid | pair_254 | 0.003 | 2085.589 | 16.380 | 15.634 | 32.706 | 20.873 | 20.606 |
| c1_hybrid | pair_255 | 0.003 | 2085.589 | 16.510 | 16.005 | 33.914 | 24.188 | 18.174 |
| c1_hybrid | pair_256 | 0.003 | 2085.589 | 18.133 | 16.864 | 36.038 | 28.241 | 19.577 |
| c1_hybrid | pair_257 | 0.003 | 2085.589 | 13.851 | 13.311 | 27.873 | 20.975 | 12.231 |
| c1_hybrid | pair_258 | 0.003 | 2085.589 | 12.053 | 12.122 | 24.852 | 17.912 | 12.220 |
| c1_hybrid | pair_259 | 0.003 | 2085.589 | 17.050 | 15.417 | 33.112 | 20.607 | 21.801 |
| c1_hybrid | pair_260 | 0.003 | 2085.589 | 16.727 | 16.744 | 34.290 | 20.722 | 23.561 |
| c1_hybrid | pair_261 | 0.003 | 2085.589 | 13.694 | 21.717 | 36.121 | 17.604 | 32.085 |
| c1_hybrid | pair_262 | 0.003 | 2085.589 | 28.540 | 20.672 | 51.313 | 56.563 | 31.372 |
| c1_hybrid | pair_263 | 0.003 | 2085.589 | 13.491 | 7.897 | 21.773 | 12.734 | 16.921 |
| c1_hybrid | pair_264 | 0.003 | 2085.589 | 12.952 | 12.622 | 26.179 | 19.523 | 11.934 |
| c1_hybrid | pair_265 | 0.003 | 2085.589 | 17.247 | 11.806 | 29.760 | 18.520 | 20.300 |
| c1_hybrid | pair_266 | 0.003 | 2085.589 | 17.106 | 16.489 | 34.193 | 19.598 | 26.948 |
| c1_hybrid | pair_267 | 0.003 | 2085.589 | 39.688 | 15.187 | 55.495 | 20.048 | 63.341 |
| c1_hybrid | pair_268 | 0.003 | 2085.589 | 15.407 | 23.231 | 39.591 | 19.705 | 36.110 |
| c1_hybrid | pair_269 | 0.003 | 2085.589 | 9.307 | 12.352 | 22.888 | 12.681 | 18.796 |
| c1_hybrid | pair_270 | 0.003 | 2085.589 | 7.910 | 7.664 | 16.032 | 12.338 | 7.514 |
| c1_hybrid | pair_271 | 0.003 | 2085.589 | 12.397 | 13.303 | 26.424 | 19.873 | 13.286 |
| c1_hybrid | pair_272 | 0.003 | 2085.589 | 14.771 | 21.811 | 38.693 | 18.339 | 35.959 |
| c1_hybrid | pair_273 | 0.003 | 2085.589 | 22.858 | 25.879 | 49.427 | 18.497 | 54.405 |
| c1_hybrid | pair_274 | 0.003 | 2085.589 | 28.628 | 17.412 | 47.759 | 17.577 | 52.312 |
| c1_hybrid | pair_275 | 0.003 | 2085.589 | 12.692 | 11.898 | 28.608 | 27.624 | 11.705 |
| c1_hybrid | pair_276 | 0.003 | 2085.589 | 12.200 | 11.801 | 25.580 | 17.515 | 14.020 |
| c1_hybrid | pair_277 | 0.003 | 2085.589 | 12.247 | 11.786 | 24.969 | 17.537 | 12.875 |
| c1_hybrid | pair_278 | 0.003 | 2085.589 | 11.636 | 11.225 | 26.226 | 17.736 | 14.916 |
| c1_hybrid | pair_279 | 0.003 | 2085.589 | 22.370 | 16.343 | 39.511 | 17.634 | 38.148 |
| c1_hybrid | pair_280 | 0.003 | 2085.589 | 16.504 | 15.910 | 33.146 | 17.488 | 27.158 |
| c1_hybrid | pair_281 | 0.003 | 2085.589 | 8.910 | 9.442 | 18.888 | 12.373 | 12.356 |
| c1_hybrid | pair_282 | 0.003 | 2085.589 | 22.720 | 19.371 | 43.770 | 18.988 | 44.203 |
| c1_hybrid | pair_283 | 0.003 | 2085.589 | 17.099 | 19.457 | 37.278 | 17.500 | 34.299 |
| c1_hybrid | pair_284 | 0.003 | 2085.589 | 17.740 | 18.811 | 37.523 | 17.524 | 34.655 |
| c1_hybrid | pair_285 | 0.003 | 2085.589 | 22.807 | 24.117 | 49.228 | 19.592 | 53.112 |
| c1_hybrid | pair_286 | 0.003 | 2085.589 | 12.446 | 10.745 | 23.770 | 11.719 | 21.482 |
| c1_hybrid | pair_287 | 0.003 | 2085.589 | 11.089 | 15.203 | 26.896 | 11.330 | 27.130 |
| c1_hybrid | pair_288 | 0.003 | 2085.589 | 21.225 | 18.560 | 40.653 | 17.662 | 39.960 |
| c1_hybrid | pair_289 | 0.003 | 2085.589 | 13.776 | 13.296 | 27.722 | 17.466 | 17.774 |
| c1_hybrid | pair_290 | 0.003 | 2085.589 | 13.096 | 15.094 | 28.733 | 17.524 | 19.453 |
| c1_hybrid | pair_291 | 0.003 | 2085.589 | 12.897 | 13.803 | 27.449 | 17.512 | 17.338 |
| c1_hybrid | pair_292 | 0.003 | 2085.589 | 14.105 | 13.450 | 28.209 | 17.500 | 18.569 |
| c1_hybrid | pair_293 | 0.003 | 2085.589 | 14.049 | 14.353 | 29.022 | 17.602 | 19.881 |
| c1_hybrid | pair_294 | 0.003 | 2085.589 | 13.206 | 24.169 | 39.939 | 17.608 | 38.852 |
| c1_hybrid | pair_295 | 0.003 | 2085.589 | 29.532 | 21.059 | 51.335 | 17.525 | 58.580 |
| c1_hybrid | pair_296 | 0.003 | 2085.589 | 23.911 | 10.851 | 35.222 | 23.497 | 25.685 |
| c1_hybrid | pair_297 | 0.003 | 2085.589 | 21.189 | 20.447 | 42.306 | 17.946 | 42.546 |
| c1_hybrid | pair_298 | 0.003 | 2085.589 | 17.565 | 11.209 | 29.172 | 11.905 | 30.586 |
| c1_hybrid | pair_299 | 0.003 | 2085.589 | 31.411 | 11.700 | 54.084 | 17.765 | 63.084 |
| c1_hybrid | pair_300 | 0.003 | 2085.589 | 17.863 | 18.828 | 37.540 | 19.490 | 31.666 |
| c1_hybrid | pair_301 | 0.003 | 2085.589 | 13.166 | 12.963 | 26.829 | 19.075 | 13.452 |
| c1_hybrid | pair_302 | 0.003 | 2085.589 | 7.554 | 7.435 | 15.336 | 11.263 | 7.146 |
| c1_hybrid | pair_303 | 0.003 | 2085.589 | 7.243 | 7.085 | 14.684 | 11.098 | 6.295 |
| c1_hybrid | pair_304 | 0.003 | 2085.589 | 12.560 | 12.394 | 25.638 | 19.243 | 11.221 |
| c1_hybrid | pair_305 | 0.003 | 2085.589 | 12.848 | 13.142 | 26.503 | 19.430 | 12.762 |
| c1_hybrid | pair_306 | 0.003 | 2085.589 | 12.059 | 12.161 | 24.787 | 19.074 | 9.905 |
| c1_hybrid | pair_307 | 0.003 | 2085.589 | 12.384 | 12.331 | 25.216 | 19.155 | 10.694 |
| c1_hybrid | pair_308 | 0.003 | 2085.589 | 12.077 | 11.454 | 24.213 | 19.414 | 8.622 |
| c1_hybrid | pair_309 | 0.003 | 2085.589 | 11.501 | 11.192 | 23.267 | 17.671 | 9.738 |
| c1_hybrid | pair_310 | 0.003 | 2085.589 | 12.308 | 12.152 | 25.158 | 19.132 | 10.536 |
| c1_hybrid | pair_311 | 0.003 | 2085.589 | 19.976 | 16.958 | 38.048 | 19.197 | 32.716 |
| c1_hybrid | pair_312 | 0.003 | 2085.589 | 20.373 | 16.513 | 37.377 | 17.620 | 34.293 |
| c1_hybrid | pair_313 | 0.003 | 2085.589 | 19.437 | 12.873 | 32.947 | 17.445 | 26.900 |
| c1_hybrid | pair_314 | 0.003 | 2085.589 | 7.319 | 8.562 | 16.593 | 11.091 | 9.559 |
| c1_hybrid | pair_315 | 0.003 | 2085.589 | 13.494 | 12.875 | 27.101 | 19.133 | 13.832 |
| c1_hybrid | pair_316 | 0.003 | 2085.589 | 13.081 | 18.414 | 32.310 | 38.611 | 11.404 |
| c1_hybrid | pair_317 | 0.003 | 2085.589 | 7.460 | 8.530 | 16.449 | 11.218 | 9.126 |
| c1_hybrid | pair_318 | 0.003 | 2085.589 | 8.007 | 7.902 | 16.783 | 11.274 | 9.593 |
| c1_hybrid | pair_319 | 0.003 | 2085.589 | 15.057 | 16.139 | 31.931 | 19.415 | 21.949 |
| c1_hybrid | pair_320 | 0.003 | 2085.589 | 14.170 | 13.626 | 28.388 | 19.302 | 15.796 |
| c1_hybrid | pair_321 | 0.003 | 2085.589 | 11.726 | 13.533 | 26.029 | 19.074 | 12.172 |
| c1_hybrid | pair_322 | 0.003 | 2085.589 | 14.021 | 14.957 | 30.249 | 20.468 | 18.207 |
| c1_hybrid | pair_323 | 0.003 | 2085.589 | 14.752 | 19.576 | 35.251 | 19.283 | 27.813 |
| c1_hybrid | pair_324 | 0.003 | 2085.589 | 13.452 | 12.824 | 27.169 | 17.671 | 16.470 |
| c1_hybrid | pair_325 | 0.003 | 2085.589 | 14.862 | 14.179 | 29.921 | 19.566 | 18.194 |
| c1_hybrid | pair_326 | 0.003 | 2085.589 | 16.786 | 14.733 | 32.560 | 19.231 | 23.153 |
| c1_hybrid | pair_327 | 0.003 | 2085.589 | 13.767 | 14.116 | 28.508 | 17.765 | 18.811 |
| c1_hybrid | pair_328 | 0.003 | 2085.589 | 13.947 | 30.705 | 45.348 | 17.873 | 47.852 |
| c1_hybrid | pair_329 | 0.003 | 2085.589 | 13.547 | 13.747 | 28.066 | 11.147 | 29.313 |
| c1_hybrid | pair_330 | 0.003 | 2085.589 | 23.960 | 23.520 | 48.338 | 19.197 | 50.629 |
| c1_hybrid | pair_331 | 0.003 | 2085.589 | 13.701 | 14.648 | 29.392 | 11.547 | 31.135 |
| c1_hybrid | pair_332 | 0.003 | 2085.589 | 14.888 | 15.305 | 30.537 | 13.036 | 31.956 |
| c1_hybrid | pair_333 | 0.003 | 2085.589 | 22.281 | 22.936 | 46.258 | 19.202 | 46.881 |
| c1_hybrid | pair_334 | 0.003 | 2085.589 | 20.650 | 20.352 | 46.246 | 29.585 | 38.184 |
| c1_hybrid | pair_335 | 0.003 | 2085.589 | 23.819 | 29.849 | 56.049 | 51.617 | 38.769 |
| c1_hybrid | pair_336 | 0.003 | 2085.589 | 22.365 | 22.373 | 46.098 | 21.150 | 44.763 |
| c1_hybrid | pair_337 | 0.003 | 2085.589 | 22.996 | 31.611 | 55.320 | 22.476 | 59.691 |
| c1_hybrid | pair_338 | 0.003 | 2085.589 | 18.867 | 20.213 | 40.415 | 21.286 | 36.423 |
| c1_hybrid | pair_339 | 0.003 | 2085.589 | 22.441 | 20.770 | 43.848 | 20.101 | 41.937 |
| c1_hybrid | pair_340 | 0.003 | 2085.589 | 20.469 | 17.101 | 38.170 | 19.615 | 32.432 |
| c1_hybrid | pair_341 | 0.003 | 2085.589 | 14.039 | 18.434 | 34.109 | 17.539 | 28.782 |
| c1_hybrid | pair_342 | 0.003 | 2085.589 | 12.201 | 11.654 | 24.998 | 19.046 | 11.673 |
| c1_hybrid | pair_343 | 0.003 | 2085.589 | 17.555 | 13.004 | 30.869 | 39.066 | 17.080 |
| c1_hybrid | pair_344 | 0.003 | 2085.589 | 21.396 | 28.755 | 50.930 | 21.657 | 52.916 |
| c1_hybrid | pair_345 | 0.003 | 2085.589 | 15.599 | 16.572 | 33.903 | 13.386 | 37.407 |
| c1_hybrid | pair_346 | 0.003 | 2085.589 | 23.510 | 15.700 | 40.198 | 17.346 | 44.989 |
| c1_hybrid | pair_347 | 0.003 | 2085.589 | 17.348 | 28.356 | 46.170 | 13.237 | 58.720 |
| c1_hybrid | pair_348 | 0.003 | 2085.589 | 7.741 | 15.935 | 24.398 | 12.867 | 21.599 |
| c1_hybrid | pair_349 | 0.003 | 2085.589 | 9.738 | 7.424 | 17.553 | 12.352 | 9.774 |
| c1_hybrid | pair_350 | 0.003 | 2085.589 | 6.459 | 6.868 | 13.618 | 12.405 | 3.352 |
| c1_hybrid | pair_351 | 0.003 | 2085.589 | 6.129 | 10.470 | 16.827 | 12.303 | 8.848 |
| c1_hybrid | pair_352 | 0.003 | 2085.589 | 5.996 | 5.929 | 12.245 | 11.781 | 1.365 |
| c1_hybrid | pair_353 | 0.003 | 2085.589 | 5.838 | 5.793 | 11.897 | 11.429 | 1.099 |
| c1_hybrid | pair_354 | 0.003 | 2085.589 | 7.573 | 7.646 | 15.541 | 11.348 | 7.405 |
| c1_hybrid | pair_355 | 0.003 | 2085.589 | 6.595 | 6.590 | 13.552 | 11.397 | 4.011 |
| c1_hybrid | pair_356 | 0.003 | 2085.589 | 7.839 | 7.990 | 17.281 | 11.648 | 10.190 |
| c1_hybrid | pair_357 | 0.003 | 2085.589 | 8.077 | 6.830 | 15.917 | 13.133 | 6.658 |
| c1_hybrid | pair_358 | 0.003 | 2085.589 | 8.140 | 9.433 | 18.062 | 13.426 | 9.994 |
| c1_hybrid | pair_359 | 0.003 | 2085.589 | 7.570 | 9.309 | 17.125 | 14.305 | 7.835 |
| c1_hybrid | pair_360 | 0.003 | 2085.589 | 9.077 | 8.183 | 17.627 | 14.454 | 8.435 |
| c1_hybrid | pair_361 | 0.003 | 2085.589 | 8.463 | 8.454 | 17.360 | 12.710 | 9.410 |
| c1_hybrid | pair_362 | 0.003 | 2085.589 | 8.616 | 8.572 | 17.570 | 11.906 | 10.567 |
| c1_hybrid | pair_363 | 0.003 | 2085.589 | 8.126 | 8.571 | 17.159 | 13.211 | 8.633 |
| c1_hybrid | pair_364 | 0.003 | 2085.589 | 9.013 | 8.823 | 18.207 | 12.573 | 11.027 |
| c1_hybrid | pair_365 | 0.003 | 2085.589 | 8.301 | 7.886 | 16.644 | 12.211 | 8.621 |
| c1_hybrid | pair_366 | 0.003 | 2085.589 | 7.622 | 8.519 | 16.510 | 12.817 | 7.920 |
| c1_hybrid | pair_367 | 0.003 | 2085.589 | 8.529 | 8.210 | 17.148 | 12.807 | 8.953 |
| c1_hybrid | pair_368 | 0.003 | 2085.589 | 9.081 | 8.898 | 18.424 | 13.350 | 10.538 |
| c1_hybrid | pair_369 | 0.003 | 2085.589 | 8.477 | 8.818 | 17.697 | 11.791 | 10.670 |
| c1_hybrid | pair_370 | 0.003 | 2085.589 | 16.381 | 15.915 | 33.152 | 23.110 | 19.627 |
| c1_hybrid | pair_371 | 0.003 | 2085.589 | 25.864 | 23.642 | 50.761 | 35.341 | 29.099 |
| c1_hybrid | pair_372 | 0.003 | 2085.589 | 26.910 | 20.561 | 48.542 | 20.720 | 48.325 |
| c1_hybrid | pair_373 | 0.003 | 2085.589 | 24.532 | 23.415 | 48.785 | 30.513 | 32.018 |
| c1_hybrid | pair_374 | 0.003 | 2085.589 | 13.903 | 12.952 | 27.342 | 17.551 | 16.965 |
| c1_hybrid | pair_375 | 0.003 | 2085.589 | 22.651 | 21.935 | 45.572 | 30.491 | 26.348 |
| c1_hybrid | pair_376 | 0.003 | 2085.589 | 24.929 | 24.752 | 50.860 | 30.503 | 35.628 |
| c1_hybrid | pair_377 | 0.003 | 2085.589 | 13.284 | 13.141 | 30.418 | 27.404 | 14.535 |
| c1_hybrid | pair_378 | 0.003 | 2085.589 | 12.023 | 12.587 | 25.085 | 19.712 | 11.120 |
| c1_hybrid | pair_379 | 0.003 | 2085.589 | 8.123 | 7.648 | 16.310 | 13.515 | 6.927 |
| c1_hybrid | pair_380 | 0.003 | 2085.589 | 14.094 | 15.695 | 30.815 | 19.436 | 19.930 |
| c1_hybrid | pair_381 | 0.003 | 2085.589 | 14.076 | 13.855 | 28.762 | 20.849 | 13.782 |
| c1_hybrid | pair_382 | 0.003 | 2085.589 | 16.003 | 15.971 | 32.522 | 20.926 | 20.289 |
| c1_hybrid | pair_383 | 0.003 | 2085.589 | 13.909 | 13.387 | 27.718 | 20.636 | 12.360 |
| c1_hybrid | pair_384 | 0.003 | 2085.589 | 13.342 | 12.730 | 26.824 | 17.565 | 16.131 |
| c1_hybrid | pair_385 | 0.003 | 2085.589 | 16.002 | 15.257 | 31.972 | 20.699 | 19.544 |
| c1_hybrid | pair_386 | 0.003 | 2085.589 | 14.238 | 14.383 | 30.189 | 23.502 | 14.162 |
| c1_hybrid | pair_387 | 0.003 | 2085.589 | 15.722 | 16.304 | 32.612 | 17.685 | 25.988 |
| c1_hybrid | pair_388 | 0.003 | 2085.589 | 11.702 | 13.421 | 26.075 | 17.657 | 14.644 |
| c1_hybrid | pair_389 | 0.003 | 2085.589 | 9.057 | 9.334 | 18.725 | 11.318 | 12.866 |
| c1_hybrid | pair_390 | 0.003 | 2085.589 | 17.057 | 15.103 | 32.804 | 19.176 | 23.682 |
| c1_hybrid | pair_391 | 0.003 | 2085.589 | 14.900 | 14.833 | 30.422 | 20.858 | 16.816 |
| c1_hybrid | pair_392 | 0.003 | 2085.589 | 35.674 | 39.903 | 78.206 | 47.852 | 52.583 |
| c1_hybrid | pair_393 | 0.003 | 2085.589 | 12.835 | 14.205 | 27.753 | 22.027 | 13.925 |
| c1_hybrid | pair_394 | 0.003 | 2085.589 | 26.265 | 24.003 | 51.355 | 34.776 | 29.164 |
| c1_hybrid | pair_395 | 0.003 | 2085.589 | 36.720 | 39.046 | 78.088 | 62.793 | 38.564 |
| c1_hybrid | pair_396 | 0.003 | 2085.589 | 14.572 | 13.695 | 28.818 | 23.310 | 14.862 |
| c1_hybrid | pair_397 | 0.003 | 2085.589 | 12.792 | 14.740 | 28.239 | 24.501 | 12.842 |
| c1_hybrid | pair_398 | 0.003 | 2085.589 | 8.907 | 8.383 | 17.727 | 12.192 | 10.345 |
| c1_hybrid | pair_399 | 0.003 | 2085.589 | 14.022 | 17.580 | 32.303 | 20.752 | 21.110 |
| c1_hybrid | pair_400 | 0.003 | 2085.589 | 18.718 | 15.213 | 42.629 | 44.774 | 18.395 |
| c1_hybrid | pair_401 | 0.003 | 2085.589 | 12.551 | 11.967 | 25.051 | 21.586 | 9.253 |
| c1_hybrid | pair_402 | 0.003 | 2085.589 | 15.656 | 13.363 | 29.496 | 21.253 | 14.675 |
| c1_hybrid | pair_403 | 0.003 | 2085.589 | 13.621 | 13.937 | 28.221 | 21.478 | 12.231 |
| c1_hybrid | pair_404 | 0.003 | 2085.589 | 12.847 | 13.500 | 27.014 | 23.025 | 12.089 |
| c1_hybrid | pair_405 | 0.003 | 2085.589 | 13.204 | 13.269 | 26.940 | 19.722 | 13.898 |
| c1_hybrid | pair_406 | 0.003 | 2085.589 | 7.914 | 7.881 | 16.096 | 12.377 | 7.332 |
| c1_hybrid | pair_407 | 0.003 | 2085.589 | 15.758 | 14.064 | 30.710 | 26.153 | 13.409 |
| c1_hybrid | pair_408 | 0.003 | 2085.589 | 13.856 | 13.637 | 28.220 | 25.141 | 11.964 |
| c1_hybrid | pair_409 | 0.003 | 2085.589 | 24.223 | 24.146 | 50.003 | 40.546 | 19.091 |
| c1_hybrid | pair_410 | 0.003 | 2085.589 | 40.231 | 41.835 | 83.917 | 58.225 | 53.156 |
| c1_hybrid | pair_411 | 0.003 | 2085.589 | 13.814 | 12.791 | 27.206 | 19.651 | 14.877 |
| c1_hybrid | pair_412 | 0.003 | 2085.589 | 11.680 | 12.970 | 25.299 | 19.315 | 11.723 |
| c1_hybrid | pair_413 | 0.003 | 2085.589 | 8.120 | 8.405 | 16.855 | 11.575 | 9.482 |
| c1_hybrid | pair_414 | 0.003 | 2085.589 | 12.849 | 16.568 | 30.069 | 19.621 | 18.579 |
| c1_hybrid | pair_415 | 0.003 | 2085.589 | 15.128 | 17.748 | 37.417 | 29.204 | 24.974 |
| c1_hybrid | pair_416 | 0.003 | 2085.589 | 11.534 | 10.435 | 22.449 | 17.836 | 8.183 |
| c1_hybrid | pair_417 | 0.003 | 2085.589 | 10.740 | 11.261 | 23.109 | 18.058 | 9.126 |
| c1_hybrid | pair_418 | 0.003 | 2085.589 | 14.237 | 12.105 | 27.017 | 17.820 | 16.094 |
| c1_hybrid | pair_419 | 0.003 | 2085.589 | 7.586 | 8.020 | 16.062 | 11.317 | 8.331 |
| c1_hybrid | pair_420 | 0.003 | 2085.589 | 11.972 | 11.848 | 24.308 | 17.833 | 11.411 |
| c1_hybrid | pair_421 | 0.003 | 2085.589 | 22.609 | 23.916 | 48.114 | 32.447 | 27.622 |
| c1_hybrid | pair_422 | 0.003 | 2085.589 | 14.384 | 11.378 | 26.390 | 17.833 | 15.093 |
| c1_hybrid | pair_423 | 0.003 | 2085.589 | 11.254 | 14.298 | 26.287 | 20.627 | 12.577 |
| c1_hybrid | pair_424 | 0.003 | 2085.589 | 7.535 | 9.031 | 16.986 | 11.858 | 9.457 |
| c1_hybrid | pair_425 | 0.003 | 2085.589 | 13.782 | 12.990 | 27.333 | 19.348 | 13.878 |
| c1_hybrid | pair_426 | 0.003 | 2085.589 | 11.862 | 11.129 | 24.299 | 19.630 | 9.928 |
| c1_hybrid | pair_427 | 0.003 | 2085.589 | 15.318 | 11.922 | 27.877 | 30.168 | 9.402 |
| c1_hybrid | pair_428 | 0.003 | 2085.589 | 7.922 | 7.721 | 16.164 | 11.944 | 7.839 |
| c1_hybrid | pair_429 | 0.003 | 2085.589 | 13.432 | 13.157 | 27.393 | 19.373 | 14.156 |
| c1_hybrid | pair_430 | 0.003 | 2085.589 | 15.927 | 11.073 | 27.574 | 30.224 | 8.843 |
| c1_hybrid | pair_431 | 0.003 | 2085.589 | 8.171 | 7.530 | 16.059 | 11.281 | 8.408 |
| c1_hybrid | pair_432 | 0.003 | 2085.589 | 12.075 | 11.421 | 23.970 | 17.722 | 10.888 |
| c1_hybrid | pair_433 | 0.003 | 2085.589 | 7.308 | 8.248 | 15.898 | 11.405 | 8.064 |
| c1_hybrid | pair_434 | 0.003 | 2085.589 | 11.251 | 11.349 | 23.089 | 17.604 | 9.505 |
| c1_hybrid | pair_435 | 0.003 | 2085.589 | 7.336 | 7.221 | 14.926 | 11.226 | 6.471 |
| c1_log | pair_001 | 0.003 | 1912.533 | 14.700 | 12.283 | 27.435 | 25.603 | 3.351 |
| c1_log | pair_002 | 0.003 | 1912.533 | 7.889 | 8.481 | 16.838 | 16.422 | 2.483 |
| c1_log | pair_003 | 0.003 | 1912.533 | 8.580 | 8.204 | 17.151 | 16.428 | 1.292 |
| c1_log | pair_004 | 0.003 | 1912.533 | 9.615 | 9.645 | 19.716 | 19.372 | 0.684 |
| c1_log | pair_005 | 0.003 | 1912.533 | 4.019 | 4.134 | 8.263 | 8.056 | 0.406 |
| c1_log | pair_006 | 0.003 | 1912.533 | 7.032 | 7.155 | 14.425 | 14.476 | 0.093 |
| c1_log | pair_007 | 0.003 | 1912.533 | 16.588 | 16.805 | 38.557 | 35.081 | 7.351 |
| c1_log | pair_008 | 0.003 | 1912.533 | 8.465 | 8.297 | 18.067 | 18.542 | 1.961 |
| c1_log | pair_009 | 0.003 | 1912.533 | 21.816 | 20.489 | 43.126 | 43.213 | 1.019 |
| c1_log | pair_010 | 0.003 | 1912.533 | 7.893 | 8.105 | 16.230 | 15.943 | 0.515 |
| c1_log | pair_011 | 0.003 | 1912.533 | 6.677 | 6.402 | 13.373 | 13.324 | 0.336 |
| c1_log | pair_012 | 0.003 | 1912.533 | 6.353 | 6.354 | 12.896 | 12.791 | 0.226 |
| c1_log | pair_013 | 0.003 | 1912.533 | 7.168 | 11.145 | 18.534 | 14.694 | 6.873 |
| c1_log | pair_014 | 0.003 | 1912.533 | 7.326 | 8.216 | 15.831 | 15.170 | 1.763 |
| c1_log | pair_015 | 0.003 | 1912.533 | 7.236 | 7.017 | 14.489 | 14.454 | 0.104 |
| c1_log | pair_016 | 0.003 | 1912.533 | 3.944 | 3.898 | 7.971 | 7.966 | 0.067 |
| c1_log | pair_017 | 0.003 | 1912.533 | 3.940 | 3.888 | 7.967 | 8.007 | 0.112 |
| c1_log | pair_018 | 0.003 | 1912.533 | 11.834 | 11.747 | 24.037 | 23.995 | 0.125 |
| c1_log | pair_019 | 0.003 | 1912.533 | 8.179 | 7.857 | 16.327 | 16.082 | 0.576 |
| c1_log | pair_020 | 0.003 | 1912.533 | 20.478 | 21.179 | 42.669 | 41.954 | 1.652 |
| c1_log | pair_021 | 0.003 | 1912.533 | 7.950 | 7.800 | 16.063 | 16.107 | 0.120 |
| c1_log | pair_022 | 0.003 | 1912.533 | 19.709 | 19.566 | 40.139 | 40.176 | 0.239 |
| c1_log | pair_023 | 0.003 | 1912.533 | 6.330 | 6.288 | 12.872 | 12.967 | 0.187 |
| c1_log | pair_024 | 0.003 | 1912.533 | 12.703 | 12.566 | 25.740 | 25.609 | 0.345 |
| c1_log | pair_025 | 0.003 | 1912.533 | 28.226 | 20.885 | 53.774 | 53.179 | 12.310 |
| c1_log | pair_026 | 0.003 | 1912.533 | 8.209 | 8.682 | 17.203 | 13.026 | 7.518 |
| c1_log | pair_027 | 0.003 | 1912.533 | 7.648 | 7.126 | 14.990 | 13.028 | 3.764 |
| c1_log | pair_028 | 0.003 | 1912.533 | 4.724 | 4.072 | 8.924 | 8.648 | 0.718 |
| c1_log | pair_029 | 0.003 | 1912.533 | 7.525 | 7.303 | 15.163 | 14.396 | 1.390 |
| c1_log | pair_030 | 0.003 | 1912.533 | 7.514 | 7.284 | 15.042 | 14.811 | 0.862 |
| c1_log | pair_031 | 0.003 | 1912.533 | 12.836 | 9.518 | 22.872 | 16.065 | 11.885 |
| c1_log | pair_032 | 0.003 | 1912.533 | 10.235 | 10.119 | 20.745 | 19.418 | 2.408 |
| c1_log | pair_033 | 0.003 | 1912.533 | 4.273 | 3.949 | 8.387 | 8.098 | 0.507 |
| c1_log | pair_034 | 0.003 | 1912.533 | 7.725 | 7.294 | 15.256 | 14.501 | 1.470 |
| c1_log | pair_035 | 0.003 | 1912.533 | 9.749 | 9.676 | 19.797 | 19.205 | 1.063 |
| c1_log | pair_036 | 0.003 | 1912.533 | 11.271 | 10.446 | 22.175 | 15.884 | 10.946 |
| c1_log | pair_037 | 0.003 | 1912.533 | 14.930 | 22.469 | 38.055 | 19.927 | 32.124 |
| c1_log | pair_038 | 0.003 | 1912.533 | 16.419 | 15.428 | 32.597 | 16.113 | 28.638 |
| c1_log | pair_039 | 0.003 | 1912.533 | 9.956 | 9.177 | 19.468 | 13.517 | 10.405 |
| c1_log | pair_040 | 0.003 | 1912.533 | 8.110 | 9.320 | 17.833 | 16.673 | 5.798 |
| c1_log | pair_041 | 0.003 | 1912.533 | 9.548 | 7.833 | 17.724 | 16.648 | 4.080 |
| c1_log | pair_042 | 0.003 | 1912.533 | 8.904 | 11.810 | 21.115 | 15.912 | 10.395 |
| c1_log | pair_043 | 0.003 | 1912.533 | 8.484 | 10.333 | 19.221 | 16.246 | 6.891 |
| c1_log | pair_044 | 0.003 | 1912.533 | 4.551 | 4.962 | 9.772 | 8.940 | 2.286 |
| c1_log | pair_045 | 0.003 | 1912.533 | 7.095 | 6.204 | 13.520 | 8.938 | 8.809 |
| c1_log | pair_046 | 0.003 | 1912.533 | 11.664 | 11.375 | 23.389 | 20.482 | 5.482 |
| c1_log | pair_047 | 0.003 | 1912.533 | 9.089 | 10.864 | 20.281 | 18.044 | 5.736 |
| c1_log | pair_048 | 0.003 | 1912.533 | 11.128 | 12.200 | 25.954 | 26.136 | 6.752 |
| c1_log | pair_049 | 0.003 | 1912.533 | 9.953 | 9.238 | 19.411 | 17.652 | 4.707 |
| c1_log | pair_050 | 0.003 | 1912.533 | 12.021 | 11.688 | 24.157 | 19.341 | 8.368 |
| c1_log | pair_051 | 0.003 | 1912.533 | 8.343 | 7.413 | 15.959 | 13.888 | 4.590 |
| c1_log | pair_052 | 0.003 | 1912.533 | 16.927 | 30.578 | 48.765 | 19.133 | 51.331 |
| c1_log | pair_053 | 0.003 | 1912.533 | 14.247 | 11.342 | 25.997 | 26.312 | 6.720 |
| c1_log | pair_054 | 0.003 | 1912.533 | 6.865 | 11.929 | 19.054 | 13.164 | 10.303 |
| c1_log | pair_055 | 0.003 | 1912.533 | 9.072 | 8.946 | 18.296 | 13.293 | 8.989 |
| c1_log | pair_056 | 0.003 | 1912.533 | 5.785 | 5.760 | 11.738 | 8.304 | 6.162 |
| c1_log | pair_057 | 0.003 | 1912.533 | 20.241 | 16.161 | 36.846 | 39.500 | 21.423 |
| c1_log | pair_058 | 0.003 | 1912.533 | 10.370 | 10.263 | 20.862 | 14.460 | 11.342 |
| c1_log | pair_059 | 0.003 | 1912.533 | 9.959 | 10.161 | 20.400 | 14.577 | 10.434 |
| c1_log | pair_060 | 0.003 | 1912.533 | 5.814 | 5.528 | 11.597 | 7.994 | 6.335 |
| c1_log | pair_061 | 0.003 | 1912.533 | 10.296 | 8.868 | 19.423 | 14.478 | 8.638 |
| c1_log | pair_062 | 0.003 | 1912.533 | 10.284 | 10.212 | 20.979 | 14.820 | 10.987 |
| c1_log | pair_063 | 0.003 | 1912.533 | 10.088 | 9.551 | 19.918 | 14.368 | 9.815 |
| c1_log | pair_064 | 0.003 | 1912.533 | 9.096 | 9.060 | 18.408 | 14.449 | 7.109 |
| c1_log | pair_065 | 0.003 | 1912.533 | 7.868 | 7.769 | 15.861 | 14.528 | 2.483 |
| c1_log | pair_066 | 0.003 | 1912.533 | 9.718 | 10.543 | 20.715 | 12.835 | 13.764 |
| c1_log | pair_067 | 0.003 | 1912.533 | 8.520 | 7.738 | 16.521 | 12.921 | 6.413 |
| c1_log | pair_068 | 0.003 | 1912.533 | 7.707 | 8.065 | 16.022 | 14.673 | 2.804 |
| c1_log | pair_069 | 0.003 | 1912.533 | 8.248 | 7.457 | 15.965 | 14.497 | 2.714 |
| c1_log | pair_070 | 0.003 | 1912.533 | 7.814 | 7.758 | 15.796 | 14.473 | 2.480 |
| c1_log | pair_071 | 0.003 | 1912.533 | 4.242 | 4.279 | 8.666 | 8.193 | 0.963 |
| c1_log | pair_072 | 0.003 | 1912.533 | 4.159 | 4.271 | 8.552 | 8.514 | 0.605 |
| c1_log | pair_073 | 0.003 | 1912.533 | 7.498 | 8.618 | 16.348 | 14.296 | 3.570 |
| c1_log | pair_074 | 0.003 | 1912.533 | 7.101 | 7.088 | 14.449 | 14.541 | 0.174 |
| c1_log | pair_075 | 0.003 | 1912.533 | 7.078 | 7.165 | 14.578 | 14.429 | 0.310 |
| c1_log | pair_076 | 0.003 | 1912.533 | 7.143 | 7.069 | 17.489 | 14.551 | 5.211 |
| c1_log | pair_077 | 0.003 | 1912.533 | 7.548 | 8.753 | 16.685 | 16.037 | 2.911 |
| c1_log | pair_078 | 0.003 | 1912.533 | 8.642 | 8.919 | 17.871 | 16.199 | 6.186 |
| c1_log | pair_079 | 0.003 | 1912.533 | 8.147 | 9.141 | 17.521 | 14.855 | 5.204 |
| c1_log | pair_080 | 0.003 | 1912.533 | 7.236 | 7.148 | 14.598 | 14.730 | 0.312 |
| c1_log | pair_081 | 0.003 | 1912.533 | 6.486 | 6.470 | 13.336 | 12.925 | 0.785 |
| c1_log | pair_082 | 0.003 | 1912.533 | 6.487 | 9.827 | 20.711 | 13.305 | 13.295 |
| c1_log | pair_083 | 0.003 | 1912.533 | 6.934 | 5.365 | 12.448 | 8.073 | 7.649 |
| c1_log | pair_084 | 0.003 | 1912.533 | 7.663 | 7.083 | 14.963 | 14.223 | 1.370 |
| c1_log | pair_085 | 0.003 | 1912.533 | 8.788 | 8.134 | 17.240 | 16.150 | 2.022 |
| c1_log | pair_086 | 0.003 | 1912.533 | 4.294 | 4.859 | 9.536 | 9.879 | 1.562 |
| c1_log | pair_087 | 0.003 | 1912.533 | 8.658 | 10.854 | 19.895 | 15.712 | 8.509 |
| c1_log | pair_088 | 0.003 | 1912.533 | 8.938 | 10.166 | 19.433 | 16.079 | 6.096 |
| c1_log | pair_089 | 0.003 | 1912.533 | 8.013 | 7.800 | 16.062 | 16.095 | 0.060 |
| c1_log | pair_090 | 0.003 | 1912.533 | 8.465 | 7.810 | 16.533 | 16.388 | 0.828 |
| c1_log | pair_091 | 0.003 | 1912.533 | 7.922 | 8.116 | 16.312 | 16.030 | 0.567 |
| c1_log | pair_092 | 0.003 | 1912.533 | 6.509 | 6.452 | 13.163 | 12.716 | 0.924 |
| c1_log | pair_093 | 0.003 | 1912.533 | 6.308 | 6.320 | 12.852 | 12.833 | 0.086 |
| c1_log | pair_094 | 0.003 | 1912.533 | 7.051 | 7.207 | 14.536 | 14.399 | 0.451 |
| c1_log | pair_095 | 0.003 | 1912.533 | 7.249 | 7.044 | 14.476 | 14.298 | 0.345 |
| c1_log | pair_096 | 0.003 | 1912.533 | 7.095 | 7.012 | 14.299 | 14.283 | 0.071 |
| c1_log | pair_097 | 0.003 | 1912.533 | 4.221 | 5.911 | 10.295 | 8.066 | 3.947 |
| c1_log | pair_098 | 0.003 | 1912.533 | 4.035 | 3.862 | 8.027 | 7.968 | 0.173 |
| c1_log | pair_099 | 0.003 | 1912.533 | 7.904 | 7.848 | 15.966 | 15.868 | 0.272 |
| c1_log | pair_100 | 0.003 | 1912.533 | 7.844 | 7.925 | 15.978 | 15.920 | 0.105 |
| c1_log | pair_101 | 0.003 | 1912.533 | 8.144 | 7.873 | 16.292 | 16.396 | 0.516 |
| c1_log | pair_102 | 0.003 | 1912.533 | 7.919 | 7.878 | 16.066 | 16.036 | 0.082 |
| c1_log | pair_103 | 0.003 | 1912.533 | 7.790 | 7.964 | 16.007 | 15.972 | 0.296 |
| c1_log | pair_104 | 0.003 | 1912.533 | 6.297 | 6.509 | 13.074 | 13.061 | 0.255 |
| c1_log | pair_105 | 0.003 | 1912.533 | 9.128 | 8.930 | 18.324 | 16.064 | 4.070 |
| c1_log | pair_106 | 0.003 | 1912.533 | 9.163 | 11.355 | 23.191 | 18.045 | 10.720 |
| c1_log | pair_107 | 0.003 | 1912.533 | 6.311 | 6.414 | 13.030 | 12.890 | 0.274 |
| c1_log | pair_108 | 0.003 | 1912.533 | 6.472 | 6.382 | 13.052 | 12.923 | 0.262 |
| c1_log | pair_109 | 0.003 | 1912.533 | 6.673 | 4.132 | 10.948 | 8.660 | 4.679 |
| c1_log | pair_110 | 0.003 | 1912.533 | 7.565 | 7.184 | 14.962 | 14.471 | 1.020 |
| c1_log | pair_111 | 0.003 | 1912.533 | 4.128 | 4.017 | 9.009 | 7.955 | 1.945 |
| c1_log | pair_112 | 0.003 | 1912.533 | 11.294 | 8.350 | 19.897 | 14.303 | 9.798 |
| c1_log | pair_113 | 0.003 | 1912.533 | 11.682 | 11.195 | 24.219 | 21.787 | 6.696 |
| c1_log | pair_114 | 0.003 | 1912.533 | 8.260 | 9.248 | 18.104 | 16.117 | 3.522 |
| c1_log | pair_115 | 0.003 | 1912.533 | 9.535 | 9.400 | 19.297 | 19.125 | 0.316 |
| c1_log | pair_116 | 0.003 | 1912.533 | 9.945 | 7.819 | 18.019 | 15.989 | 3.750 |
| c1_log | pair_117 | 0.003 | 1912.533 | 6.793 | 7.332 | 14.391 | 12.711 | 3.012 |
| c1_log | pair_118 | 0.003 | 1912.533 | 6.260 | 6.271 | 12.775 | 12.772 | 0.022 |
| c1_log | pair_119 | 0.003 | 1912.533 | 7.060 | 7.050 | 14.325 | 14.289 | 0.069 |
| c1_log | pair_120 | 0.003 | 1912.533 | 7.219 | 7.339 | 14.819 | 14.478 | 0.750 |
| c1_log | pair_121 | 0.003 | 1912.533 | 7.270 | 7.162 | 14.656 | 14.501 | 0.369 |
| c1_log | pair_122 | 0.003 | 1912.533 | 3.939 | 3.955 | 8.008 | 8.043 | 0.202 |
| c1_log | pair_123 | 0.003 | 1912.533 | 4.050 | 3.903 | 8.091 | 8.123 | 0.178 |
| c1_log | pair_124 | 0.003 | 1912.533 | 9.543 | 9.370 | 19.180 | 19.174 | 0.106 |
| c1_log | pair_125 | 0.003 | 1912.533 | 7.880 | 7.983 | 16.072 | 16.061 | 0.045 |
| c1_log | pair_126 | 0.003 | 1912.533 | 9.461 | 9.437 | 19.219 | 19.170 | 0.198 |
| c1_log | pair_127 | 0.003 | 1912.533 | 7.925 | 7.877 | 16.018 | 16.071 | 0.100 |
| c1_log | pair_128 | 0.003 | 1912.533 | 9.550 | 9.460 | 19.323 | 19.213 | 0.237 |
| c1_log | pair_129 | 0.003 | 1912.533 | 6.379 | 6.347 | 12.922 | 12.891 | 0.100 |
| c1_log | pair_130 | 0.003 | 1912.533 | 9.392 | 9.459 | 19.173 | 19.200 | 0.106 |
| c1_log | pair_131 | 0.003 | 1912.533 | 16.398 | 9.895 | 26.684 | 19.762 | 12.446 |
| c1_log | pair_132 | 0.003 | 1912.533 | 6.301 | 6.251 | 12.725 | 12.727 | 0.017 |
| c1_log | pair_133 | 0.003 | 1912.533 | 6.392 | 6.239 | 12.843 | 12.834 | 0.216 |
| c1_log | pair_134 | 0.003 | 1912.533 | 4.927 | 5.275 | 10.360 | 8.037 | 4.144 |
| c1_log | pair_135 | 0.003 | 1912.533 | 7.353 | 7.386 | 163.220 | 15.953 | 256.259 |
| c1_log | pair_136 | 0.003 | 1912.533 | 5.810 | 4.663 | 10.771 | 8.653 | 4.038 |
| c1_log | pair_137 | 0.003 | 1912.533 | 4.058 | 4.087 | 8.275 | 8.378 | 0.301 |
| c1_log | pair_138 | 0.003 | 1912.533 | 4.164 | 3.918 | 8.211 | 8.297 | 0.237 |
| c1_log | pair_139 | 0.003 | 1912.533 | 4.199 | 4.090 | 8.416 | 8.025 | 0.687 |
| c1_log | pair_140 | 0.003 | 1912.533 | 3.913 | 4.538 | 8.540 | 7.914 | 1.179 |
| c1_log | pair_141 | 0.003 | 1912.533 | 4.419 | 3.854 | 8.367 | 8.028 | 0.727 |
| c1_log | pair_142 | 0.003 | 1912.533 | 4.375 | 4.228 | 13.054 | 10.195 | 6.952 |
| c1_log | pair_143 | 0.003 | 1912.533 | 4.621 | 4.498 | 9.278 | 8.482 | 1.892 |
| c1_log | pair_144 | 0.003 | 1912.533 | 4.566 | 4.708 | 9.535 | 9.024 | 1.846 |
| c1_log | pair_145 | 0.003 | 1912.533 | 4.135 | 4.448 | 8.709 | 8.555 | 0.847 |
| c1_log | pair_146 | 0.003 | 1912.533 | 4.236 | 4.605 | 9.065 | 8.092 | 1.887 |
| c1_log | pair_147 | 0.003 | 1912.533 | 4.255 | 3.900 | 8.306 | 7.977 | 0.666 |
| c1_log | pair_148 | 0.003 | 1912.533 | 6.817 | 4.299 | 11.251 | 7.999 | 5.642 |
| c1_log | pair_149 | 0.003 | 1912.533 | 5.032 | 4.552 | 15.277 | 13.264 | 8.544 |
| c1_log | pair_150 | 0.003 | 1912.533 | 7.411 | 6.603 | 14.166 | 7.860 | 10.948 |
| c1_log | pair_151 | 0.003 | 1912.533 | 4.139 | 4.061 | 8.426 | 8.049 | 0.722 |
| c1_log | pair_152 | 0.003 | 1912.533 | 4.030 | 3.902 | 8.102 | 8.002 | 0.238 |
| c1_log | pair_153 | 0.003 | 1912.533 | 4.615 | 4.020 | 8.852 | 7.896 | 1.679 |
| c1_log | pair_154 | 0.003 | 1912.533 | 4.010 | 3.996 | 8.113 | 7.982 | 0.276 |
| c1_log | pair_155 | 0.003 | 1912.533 | 3.948 | 3.956 | 8.047 | 7.948 | 0.206 |
| c1_log | pair_156 | 0.003 | 1912.533 | 4.398 | 4.002 | 8.537 | 8.204 | 0.814 |
| c1_log | pair_157 | 0.003 | 1912.533 | 4.070 | 3.995 | 8.194 | 7.976 | 0.430 |
| c1_log | pair_158 | 0.003 | 1912.533 | 4.029 | 4.013 | 8.246 | 8.071 | 0.536 |
| c1_log | pair_159 | 0.003 | 1912.533 | 4.053 | 3.929 | 8.123 | 7.988 | 0.285 |
| c1_log | pair_160 | 0.003 | 1912.533 | 7.172 | 7.597 | 15.130 | 15.118 | 0.845 |
| c1_log | pair_161 | 0.003 | 1912.533 | 7.053 | 7.218 | 14.519 | 14.458 | 0.281 |
| c1_log | pair_162 | 0.003 | 1912.533 | 7.125 | 7.109 | 14.523 | 14.592 | 0.203 |
| c1_log | pair_163 | 0.003 | 1912.533 | 7.103 | 7.056 | 14.420 | 14.415 | 0.069 |
| c1_log | pair_164 | 0.003 | 1912.533 | 6.306 | 6.249 | 12.773 | 12.756 | 0.180 |
| c1_log | pair_165 | 0.003 | 1912.533 | 6.242 | 6.257 | 12.683 | 12.649 | 0.102 |
| c1_log | pair_166 | 0.003 | 1912.533 | 7.072 | 7.009 | 14.274 | 14.272 | 0.067 |
| c1_log | pair_167 | 0.003 | 1912.533 | 7.165 | 7.074 | 14.464 | 14.492 | 0.219 |
| c1_log | pair_168 | 0.003 | 1912.533 | 7.272 | 7.048 | 14.624 | 14.585 | 0.156 |
| c1_log | pair_169 | 0.003 | 1912.533 | 4.704 | 4.769 | 9.628 | 8.040 | 2.887 |
| c1_log | pair_170 | 0.003 | 1912.533 | 6.553 | 5.908 | 12.619 | 11.728 | 5.070 |
| c1_log | pair_171 | 0.003 | 1912.533 | 7.130 | 7.018 | 14.355 | 14.361 | 0.122 |
| c1_log | pair_172 | 0.003 | 1912.533 | 7.116 | 6.971 | 16.210 | 14.361 | 3.226 |
| c1_log | pair_173 | 0.003 | 1912.533 | 7.049 | 7.025 | 14.307 | 14.326 | 0.065 |
| c1_log | pair_174 | 0.003 | 1912.533 | 7.042 | 7.025 | 14.291 | 14.313 | 0.102 |
| c1_log | pair_175 | 0.003 | 1912.533 | 7.126 | 7.077 | 14.440 | 14.202 | 0.415 |
| c1_log | pair_176 | 0.003 | 1912.533 | 6.279 | 6.239 | 12.675 | 12.686 | 0.058 |
| c1_log | pair_177 | 0.003 | 1912.533 | 7.099 | 6.999 | 14.347 | 14.350 | 0.144 |
| c1_log | pair_178 | 0.003 | 1912.533 | 7.026 | 7.012 | 14.248 | 14.240 | 0.022 |
| c1_log | pair_179 | 0.003 | 1912.533 | 6.344 | 6.267 | 18.118 | 13.078 | 8.999 |
| c1_log | pair_180 | 0.003 | 1912.533 | 7.297 | 6.382 | 13.880 | 12.924 | 1.889 |
| c1_log | pair_181 | 0.003 | 1912.533 | 4.600 | 4.310 | 9.022 | 7.934 | 1.969 |
| c1_log | pair_182 | 0.003 | 1912.533 | 7.813 | 8.214 | 16.340 | 14.208 | 3.744 |
| c1_log | pair_183 | 0.003 | 1912.533 | 9.250 | 8.999 | 18.940 | 17.123 | 4.385 |
| c1_log | pair_184 | 0.003 | 1912.533 | 24.303 | 23.459 | 49.559 | 48.702 | 16.920 |
| c1_log | pair_185 | 0.003 | 1912.533 | 9.459 | 8.642 | 18.619 | 18.978 | 2.686 |
| c1_log | pair_186 | 0.003 | 1912.533 | 12.186 | 7.964 | 20.506 | 13.898 | 12.480 |
| c1_log | pair_187 | 0.003 | 1912.533 | 6.860 | 7.195 | 14.399 | 14.051 | 1.831 |
| c1_log | pair_188 | 0.003 | 1912.533 | 7.091 | 7.570 | 14.879 | 14.654 | 0.839 |
| c1_log | pair_189 | 0.003 | 1912.533 | 8.660 | 8.816 | 17.744 | 14.563 | 5.590 |
| c1_log | pair_190 | 0.003 | 1912.533 | 8.418 | 8.814 | 17.448 | 14.306 | 5.520 |
| c1_log | pair_191 | 0.003 | 1912.533 | 4.674 | 5.000 | 9.824 | 8.153 | 3.126 |
| c1_log | pair_192 | 0.003 | 1912.533 | 4.652 | 4.339 | 9.111 | 7.864 | 2.162 |
| c1_log | pair_193 | 0.003 | 1912.533 | 13.717 | 12.004 | 26.132 | 24.304 | 3.362 |
| c1_log | pair_194 | 0.003 | 1912.533 | 7.921 | 7.894 | 16.103 | 16.095 | 0.235 |
| c1_log | pair_195 | 0.003 | 1912.533 | 18.219 | 22.001 | 40.813 | 39.261 | 8.569 |
| c1_log | pair_196 | 0.003 | 1912.533 | 7.822 | 7.801 | 15.921 | 15.921 | 0.050 |
| c1_log | pair_197 | 0.003 | 1912.533 | 16.860 | 16.624 | 34.661 | 34.935 | 1.409 |
| c1_log | pair_198 | 0.003 | 1912.533 | 6.461 | 7.203 | 13.925 | 12.864 | 2.011 |
| c1_log | pair_199 | 0.003 | 1912.533 | 13.024 | 12.861 | 26.278 | 25.962 | 0.837 |
| c1_log | pair_200 | 0.003 | 1912.533 | 16.603 | 16.710 | 33.953 | 33.585 | 0.645 |
| c1_log | pair_201 | 0.003 | 1912.533 | 10.549 | 6.835 | 17.651 | 12.877 | 8.484 |
| c1_log | pair_202 | 0.003 | 1912.533 | 6.711 | 6.286 | 13.189 | 12.718 | 0.832 |
| c1_log | pair_203 | 0.003 | 1912.533 | 3.944 | 3.923 | 8.012 | 8.015 | 0.097 |
| c1_log | pair_204 | 0.003 | 1912.533 | 7.212 | 7.007 | 15.363 | 14.771 | 1.424 |
| c1_log | pair_205 | 0.003 | 1912.533 | 8.079 | 8.145 | 16.558 | 16.064 | 1.096 |
| c1_log | pair_206 | 0.003 | 1912.533 | 10.920 | 10.061 | 21.335 | 20.229 | 5.604 |
| c1_log | pair_207 | 0.003 | 1912.533 | 6.837 | 6.844 | 13.877 | 12.831 | 1.829 |
| c1_log | pair_208 | 0.003 | 1912.533 | 6.797 | 6.758 | 13.725 | 13.077 | 1.337 |
| c1_log | pair_209 | 0.003 | 1912.533 | 7.439 | 7.378 | 15.064 | 14.433 | 1.342 |
| c1_log | pair_210 | 0.003 | 1912.533 | 7.118 | 8.983 | 16.479 | 16.053 | 2.323 |
| c1_log | pair_211 | 0.003 | 1912.533 | 8.639 | 7.373 | 16.219 | 15.774 | 2.166 |
| c1_log | pair_212 | 0.003 | 1912.533 | 5.710 | 5.314 | 11.178 | 7.967 | 5.573 |
| c1_log | pair_213 | 0.003 | 1912.533 | 3.945 | 3.947 | 8.029 | 8.034 | 0.035 |
| c1_log | pair_214 | 0.003 | 1912.533 | 8.173 | 8.561 | 17.001 | 17.162 | 0.812 |
| c1_log | pair_215 | 0.003 | 1912.533 | 7.874 | 7.910 | 15.991 | 15.970 | 0.069 |
| c1_log | pair_216 | 0.003 | 1912.533 | 7.856 | 7.895 | 15.995 | 15.975 | 0.179 |
| c1_log | pair_217 | 0.003 | 1912.533 | 7.866 | 7.937 | 16.050 | 15.950 | 0.182 |
| c1_log | pair_218 | 0.003 | 1912.533 | 7.971 | 7.843 | 16.049 | 16.045 | 0.221 |
| c1_log | pair_219 | 0.003 | 1912.533 | 6.357 | 6.218 | 12.795 | 12.804 | 0.167 |
| c1_log | pair_220 | 0.003 | 1912.533 | 7.966 | 7.916 | 16.136 | 16.213 | 0.186 |
| c1_log | pair_221 | 0.003 | 1912.533 | 7.851 | 7.830 | 15.857 | 15.861 | 0.114 |
| c1_log | pair_222 | 0.003 | 1912.533 | 6.257 | 6.270 | 12.731 | 12.741 | 0.056 |
| c1_log | pair_223 | 0.003 | 1912.533 | 14.107 | 7.518 | 21.877 | 12.810 | 15.783 |
| c1_log | pair_224 | 0.003 | 1912.533 | 4.294 | 8.665 | 13.194 | 13.644 | 5.010 |
| c1_log | pair_225 | 0.003 | 1912.533 | 7.300 | 7.018 | 14.581 | 14.376 | 0.431 |
| c1_log | pair_226 | 0.003 | 1912.533 | 8.101 | 7.775 | 16.119 | 15.937 | 0.410 |
| c1_log | pair_227 | 0.003 | 1912.533 | 6.432 | 6.354 | 12.942 | 12.893 | 0.123 |
| c1_log | pair_228 | 0.003 | 1912.533 | 6.304 | 6.263 | 12.807 | 12.810 | 0.053 |
| c1_log | pair_229 | 0.003 | 1912.533 | 7.104 | 7.083 | 14.650 | 14.419 | 0.444 |
| c1_log | pair_230 | 0.003 | 1912.533 | 7.204 | 7.675 | 15.091 | 14.264 | 1.484 |
| c1_log | pair_231 | 0.003 | 1912.533 | 7.657 | 7.608 | 15.503 | 14.308 | 2.162 |
| c1_log | pair_232 | 0.003 | 1912.533 | 4.203 | 4.056 | 8.394 | 8.204 | 0.542 |
| c1_log | pair_233 | 0.003 | 1912.533 | 4.053 | 4.000 | 8.209 | 8.223 | 0.157 |
| c1_log | pair_234 | 0.003 | 1912.533 | 12.937 | 13.273 | 28.529 | 29.079 | 4.157 |
| c1_log | pair_235 | 0.003 | 1912.533 | 10.628 | 13.399 | 24.562 | 27.340 | 7.396 |
| c1_log | pair_236 | 0.003 | 1912.533 | 30.104 | 23.798 | 54.833 | 59.276 | 11.486 |
| c1_log | pair_237 | 0.003 | 1912.533 | 8.336 | 7.868 | 16.477 | 16.509 | 0.344 |
| c1_log | pair_238 | 0.003 | 1912.533 | 19.916 | 21.579 | 42.204 | 40.709 | 3.437 |
| c1_log | pair_239 | 0.003 | 1912.533 | 6.292 | 6.243 | 12.802 | 12.717 | 0.148 |
| c1_log | pair_240 | 0.003 | 1912.533 | 12.691 | 12.801 | 25.923 | 26.084 | 0.383 |
| c1_log | pair_241 | 0.003 | 1912.533 | 20.627 | 21.322 | 42.958 | 41.688 | 2.201 |
| c1_log | pair_242 | 0.003 | 1912.533 | 6.277 | 6.311 | 12.813 | 12.796 | 0.054 |
| c1_log | pair_243 | 0.003 | 1912.533 | 6.435 | 6.363 | 13.002 | 12.862 | 0.303 |
| c1_log | pair_244 | 0.003 | 1912.533 | 3.938 | 3.959 | 8.052 | 8.053 | 0.010 |
| c1_log | pair_245 | 0.003 | 1912.533 | 7.104 | 7.050 | 14.593 | 14.736 | 0.318 |
| c1_log | pair_246 | 0.003 | 1912.533 | 7.903 | 7.721 | 15.822 | 13.092 | 4.983 |
| c1_log | pair_247 | 0.003 | 1912.533 | 6.326 | 6.210 | 12.752 | 12.789 | 0.198 |
| c1_log | pair_248 | 0.003 | 1912.533 | 7.162 | 7.014 | 14.395 | 14.358 | 0.091 |
| c1_log | pair_249 | 0.003 | 1912.533 | 7.134 | 7.044 | 14.392 | 14.387 | 0.249 |
| c1_log | pair_250 | 0.003 | 1912.533 | 7.109 | 7.015 | 17.156 | 14.241 | 5.095 |
| c1_log | pair_251 | 0.003 | 1912.533 | 3.935 | 4.025 | 8.128 | 8.099 | 0.082 |
| c1_log | pair_252 | 0.003 | 1912.533 | 5.379 | 4.180 | 9.720 | 8.093 | 2.953 |
| c1_log | pair_253 | 0.003 | 1912.533 | 7.915 | 7.874 | 16.055 | 15.975 | 0.209 |
| c1_log | pair_254 | 0.003 | 1912.533 | 11.110 | 8.629 | 19.952 | 15.861 | 7.148 |
| c1_log | pair_255 | 0.003 | 1912.533 | 7.832 | 7.815 | 16.688 | 16.026 | 1.354 |
| c1_log | pair_256 | 0.003 | 1912.533 | 8.536 | 8.518 | 17.310 | 16.042 | 2.431 |
| c1_log | pair_257 | 0.003 | 1912.533 | 11.620 | 12.108 | 23.991 | 20.588 | 10.391 |
| c1_log | pair_258 | 0.003 | 1912.533 | 6.802 | 6.226 | 13.295 | 13.084 | 0.718 |
| c1_log | pair_259 | 0.003 | 1912.533 | 10.578 | 8.282 | 19.224 | 15.947 | 5.799 |
| c1_log | pair_260 | 0.003 | 1912.533 | 12.480 | 8.651 | 21.401 | 15.930 | 9.571 |
| c1_log | pair_261 | 0.003 | 1912.533 | 7.059 | 6.257 | 13.549 | 12.802 | 1.420 |
| c1_log | pair_262 | 0.003 | 1912.533 | 6.316 | 6.251 | 12.795 | 12.727 | 0.145 |
| c1_log | pair_263 | 0.003 | 1912.533 | 4.060 | 3.909 | 8.112 | 7.981 | 0.245 |
| c1_log | pair_264 | 0.003 | 1912.533 | 7.294 | 10.830 | 18.418 | 14.610 | 6.918 |
| c1_log | pair_265 | 0.003 | 1912.533 | 6.598 | 6.258 | 13.137 | 12.990 | 0.479 |
| c1_log | pair_266 | 0.003 | 1912.533 | 6.294 | 9.495 | 16.011 | 12.666 | 5.815 |
| c1_log | pair_267 | 0.003 | 1912.533 | 6.318 | 8.357 | 14.934 | 12.651 | 4.145 |
| c1_log | pair_268 | 0.003 | 1912.533 | 7.821 | 6.206 | 14.266 | 12.618 | 2.899 |
| c1_log | pair_269 | 0.003 | 1912.533 | 3.947 | 3.943 | 8.052 | 8.047 | 0.131 |
| c1_log | pair_270 | 0.003 | 1912.533 | 3.914 | 3.908 | 7.943 | 7.946 | 0.041 |
| c1_log | pair_271 | 0.003 | 1912.533 | 9.328 | 9.723 | 19.393 | 13.009 | 11.293 |
| c1_log | pair_272 | 0.003 | 1912.533 | 6.733 | 6.293 | 13.282 | 13.054 | 0.664 |
| c1_log | pair_273 | 0.003 | 1912.533 | 9.633 | 8.989 | 18.967 | 12.719 | 10.900 |
| c1_log | pair_274 | 0.003 | 1912.533 | 6.882 | 6.620 | 13.734 | 12.978 | 1.657 |
| c1_log | pair_275 | 0.003 | 1912.533 | 6.251 | 6.251 | 12.713 | 12.722 | 0.129 |
| c1_log | pair_276 | 0.003 | 1912.533 | 6.379 | 7.854 | 14.610 | 13.045 | 3.056 |
| c1_log | pair_277 | 0.003 | 1912.533 | 9.142 | 6.754 | 16.119 | 12.759 | 5.857 |
| c1_log | pair_278 | 0.003 | 1912.533 | 7.792 | 8.224 | 16.218 | 12.728 | 6.148 |
| c1_log | pair_279 | 0.003 | 1912.533 | 6.291 | 6.198 | 12.691 | 12.695 | 0.018 |
| c1_log | pair_280 | 0.003 | 1912.533 | 6.360 | 6.244 | 12.845 | 12.885 | 0.078 |
| c1_log | pair_281 | 0.003 | 1912.533 | 3.963 | 3.943 | 8.929 | 8.125 | 1.462 |
| c1_log | pair_282 | 0.003 | 1912.533 | 17.100 | 6.834 | 24.247 | 12.810 | 19.814 |
| c1_log | pair_283 | 0.003 | 1912.533 | 6.462 | 6.926 | 13.676 | 12.887 | 1.537 |
| c1_log | pair_284 | 0.003 | 1912.533 | 7.153 | 6.883 | 14.386 | 13.237 | 2.506 |
| c1_log | pair_285 | 0.003 | 1912.533 | 8.725 | 9.079 | 18.426 | 13.016 | 9.655 |
| c1_log | pair_286 | 0.003 | 1912.533 | 4.340 | 4.722 | 9.249 | 8.122 | 2.207 |
| c1_log | pair_287 | 0.003 | 1912.533 | 4.914 | 3.892 | 8.982 | 7.951 | 1.813 |
| c1_log | pair_288 | 0.003 | 1912.533 | 6.409 | 6.280 | 13.566 | 13.315 | 0.963 |
| c1_log | pair_289 | 0.003 | 1912.533 | 8.736 | 6.407 | 15.353 | 12.974 | 4.162 |
| c1_log | pair_290 | 0.003 | 1912.533 | 6.318 | 6.446 | 12.971 | 12.860 | 0.332 |
| c1_log | pair_291 | 0.003 | 1912.533 | 6.281 | 6.304 | 12.765 | 12.694 | 0.170 |
| c1_log | pair_292 | 0.003 | 1912.533 | 6.271 | 6.342 | 12.823 | 12.793 | 0.183 |
| c1_log | pair_293 | 0.003 | 1912.533 | 11.455 | 6.617 | 18.333 | 13.299 | 9.235 |
| c1_log | pair_294 | 0.003 | 1912.533 | 7.485 | 6.344 | 13.968 | 12.609 | 2.374 |
| c1_log | pair_295 | 0.003 | 1912.533 | 6.443 | 6.363 | 12.997 | 13.152 | 0.338 |
| c1_log | pair_296 | 0.003 | 1912.533 | 6.311 | 6.202 | 12.702 | 12.669 | 0.064 |
| c1_log | pair_297 | 0.003 | 1912.533 | 6.264 | 6.224 | 12.707 | 12.732 | 0.050 |
| c1_log | pair_298 | 0.003 | 1912.533 | 3.896 | 3.878 | 7.903 | 7.926 | 0.057 |
| c1_log | pair_299 | 0.003 | 1912.533 | 6.264 | 6.150 | 12.613 | 12.610 | 0.088 |
| c1_log | pair_300 | 0.003 | 1912.533 | 7.020 | 6.977 | 14.217 | 14.230 | 0.038 |
| c1_log | pair_301 | 0.003 | 1912.533 | 9.685 | 8.414 | 18.282 | 14.722 | 6.419 |
| c1_log | pair_302 | 0.003 | 1912.533 | 4.179 | 3.920 | 8.240 | 8.134 | 0.309 |
| c1_log | pair_303 | 0.003 | 1912.533 | 3.957 | 3.900 | 7.968 | 7.967 | 0.171 |
| c1_log | pair_304 | 0.003 | 1912.533 | 7.137 | 6.979 | 14.338 | 14.365 | 0.137 |
| c1_log | pair_305 | 0.003 | 1912.533 | 7.044 | 7.042 | 14.309 | 14.322 | 0.115 |
| c1_log | pair_306 | 0.003 | 1912.533 | 7.558 | 8.080 | 16.267 | 15.651 | 2.529 |
| c1_log | pair_307 | 0.003 | 1912.533 | 8.893 | 11.200 | 20.536 | 18.285 | 7.702 |
| c1_log | pair_308 | 0.003 | 1912.533 | 7.262 | 7.296 | 14.831 | 14.707 | 0.476 |
| c1_log | pair_309 | 0.003 | 1912.533 | 9.232 | 6.507 | 15.933 | 12.720 | 5.594 |
| c1_log | pair_310 | 0.003 | 1912.533 | 8.283 | 7.838 | 16.316 | 14.589 | 3.308 |
| c1_log | pair_311 | 0.003 | 1912.533 | 9.499 | 7.069 | 16.807 | 14.618 | 4.119 |
| c1_log | pair_312 | 0.003 | 1912.533 | 10.011 | 7.292 | 17.561 | 14.245 | 7.043 |
| c1_log | pair_313 | 0.003 | 1912.533 | 8.376 | 7.156 | 15.844 | 15.666 | 3.166 |
| c1_log | pair_314 | 0.003 | 1912.533 | 4.179 | 4.294 | 8.649 | 8.028 | 1.092 |
| c1_log | pair_315 | 0.003 | 1912.533 | 9.475 | 10.555 | 21.251 | 22.446 | 6.678 |
| c1_log | pair_316 | 0.003 | 1912.533 | 9.784 | 9.978 | 20.232 | 21.468 | 5.353 |
| c1_log | pair_317 | 0.003 | 1912.533 | 4.008 | 4.152 | 8.282 | 8.429 | 0.339 |
| c1_log | pair_318 | 0.003 | 1912.533 | 3.976 | 5.051 | 9.199 | 8.304 | 1.876 |
| c1_log | pair_319 | 0.003 | 1912.533 | 8.604 | 7.145 | 15.954 | 14.735 | 2.528 |
| c1_log | pair_320 | 0.003 | 1912.533 | 7.221 | 7.262 | 14.730 | 14.758 | 0.416 |
| c1_log | pair_321 | 0.003 | 1912.533 | 9.358 | 12.718 | 22.336 | 15.723 | 12.783 |
| c1_log | pair_322 | 0.003 | 1912.533 | 11.589 | 7.777 | 19.615 | 16.415 | 7.504 |
| c1_log | pair_323 | 0.003 | 1912.533 | 7.384 | 8.310 | 15.880 | 15.507 | 1.816 |
| c1_log | pair_324 | 0.003 | 1912.533 | 6.333 | 6.293 | 12.839 | 12.870 | 0.133 |
| c1_log | pair_325 | 0.003 | 1912.533 | 7.072 | 7.067 | 14.332 | 14.328 | 0.095 |
| c1_log | pair_326 | 0.003 | 1912.533 | 7.076 | 7.893 | 15.251 | 14.273 | 1.708 |
| c1_log | pair_327 | 0.003 | 1912.533 | 8.281 | 6.361 | 14.899 | 12.941 | 3.674 |
| c1_log | pair_328 | 0.003 | 1912.533 | 6.281 | 6.435 | 12.949 | 12.949 | 0.112 |
| c1_log | pair_329 | 0.003 | 1912.533 | 4.004 | 4.045 | 8.167 | 8.233 | 0.225 |
| c1_log | pair_330 | 0.003 | 1912.533 | 7.087 | 7.160 | 14.482 | 14.408 | 0.255 |
| c1_log | pair_331 | 0.003 | 1912.533 | 4.582 | 3.980 | 8.693 | 8.063 | 1.215 |
| c1_log | pair_332 | 0.003 | 1912.533 | 4.080 | 3.994 | 8.206 | 8.256 | 0.253 |
| c1_log | pair_333 | 0.003 | 1912.533 | 7.155 | 7.014 | 14.540 | 14.431 | 0.460 |
| c1_log | pair_334 | 0.003 | 1912.533 | 10.004 | 8.992 | 19.223 | 15.058 | 7.738 |
| c1_log | pair_335 | 0.003 | 1912.533 | 7.228 | 8.549 | 16.009 | 15.075 | 2.198 |
| c1_log | pair_336 | 0.003 | 1912.533 | 7.288 | 7.134 | 14.661 | 14.706 | 0.150 |
| c1_log | pair_337 | 0.003 | 1912.533 | 7.094 | 7.058 | 14.380 | 14.391 | 0.233 |
| c1_log | pair_338 | 0.003 | 1912.533 | 6.257 | 6.299 | 12.754 | 12.800 | 0.113 |
| c1_log | pair_339 | 0.003 | 1912.533 | 9.229 | 7.008 | 16.455 | 14.326 | 3.703 |
| c1_log | pair_340 | 0.003 | 1912.533 | 8.856 | 9.347 | 18.464 | 14.361 | 7.181 |
| c1_log | pair_341 | 0.003 | 1912.533 | 6.275 | 6.324 | 12.784 | 12.792 | 0.028 |
| c1_log | pair_342 | 0.003 | 1912.533 | 6.286 | 6.308 | 12.808 | 12.792 | 0.066 |
| c1_log | pair_343 | 0.003 | 1912.533 | 3.975 | 3.966 | 8.104 | 8.064 | 0.098 |
| c1_log | pair_344 | 0.003 | 1912.533 | 8.051 | 6.993 | 15.245 | 14.391 | 1.738 |
| c1_log | pair_345 | 0.003 | 1912.533 | 3.922 | 3.921 | 7.990 | 8.008 | 0.043 |
| c1_log | pair_346 | 0.003 | 1912.533 | 5.036 | 3.924 | 9.068 | 8.115 | 1.811 |
| c1_log | pair_347 | 0.003 | 1912.533 | 3.976 | 4.070 | 8.197 | 8.109 | 0.302 |
| c1_log | pair_348 | 0.003 | 1912.533 | 4.124 | 4.062 | 8.265 | 8.168 | 0.476 |
| c1_log | pair_349 | 0.003 | 1912.533 | 4.005 | 4.005 | 8.134 | 8.172 | 0.097 |
| c1_log | pair_350 | 0.003 | 1912.533 | 3.929 | 3.911 | 7.991 | 7.968 | 0.076 |
| c1_log | pair_351 | 0.003 | 1912.533 | 3.978 | 3.890 | 7.973 | 7.990 | 0.093 |
| c1_log | pair_352 | 0.003 | 1912.533 | 5.729 | 5.957 | 12.337 | 8.236 | 7.395 |
| c1_log | pair_353 | 0.003 | 1912.533 | 3.915 | 4.012 | 10.631 | 7.951 | 4.711 |
| c1_log | pair_354 | 0.003 | 1912.533 | 5.315 | 5.160 | 24.109 | 15.791 | 21.591 |
| c1_log | pair_355 | 0.003 | 1912.533 | 4.731 | 5.214 | 10.274 | 10.358 | 2.291 |
| c1_log | pair_356 | 0.003 | 1912.533 | 7.062 | 5.419 | 12.825 | 10.177 | 6.699 |
| c1_log | pair_357 | 0.003 | 1912.533 | 4.065 | 3.897 | 8.147 | 8.197 | 0.090 |
| c1_log | pair_358 | 0.003 | 1912.533 | 3.971 | 3.927 | 8.044 | 8.030 | 0.108 |
| c1_log | pair_359 | 0.003 | 1912.533 | 5.098 | 6.206 | 11.432 | 8.015 | 6.054 |
| c1_log | pair_360 | 0.003 | 1912.533 | 4.325 | 4.185 | 8.656 | 8.871 | 0.674 |
| c1_log | pair_361 | 0.003 | 1912.533 | 7.150 | 5.466 | 12.732 | 8.916 | 7.631 |
| c1_log | pair_362 | 0.003 | 1912.533 | 5.179 | 4.902 | 10.212 | 8.788 | 3.324 |
| c1_log | pair_363 | 0.003 | 1912.533 | 5.190 | 5.425 | 10.800 | 8.924 | 4.070 |
| c1_log | pair_364 | 0.003 | 1912.533 | 5.074 | 4.753 | 9.999 | 8.888 | 2.746 |
| c1_log | pair_365 | 0.003 | 1912.533 | 5.257 | 5.546 | 11.008 | 8.804 | 4.589 |
| c1_log | pair_366 | 0.003 | 1912.533 | 4.622 | 4.308 | 9.110 | 8.987 | 1.204 |
| c1_log | pair_367 | 0.003 | 1912.533 | 4.354 | 4.171 | 8.662 | 8.740 | 0.485 |
| c1_log | pair_368 | 0.003 | 1912.533 | 4.095 | 4.492 | 8.692 | 8.102 | 1.097 |
| c1_log | pair_369 | 0.003 | 1912.533 | 5.240 | 4.778 | 10.159 | 8.112 | 3.707 |
| c1_log | pair_370 | 0.003 | 1912.533 | 7.902 | 8.019 | 16.156 | 15.935 | 0.482 |
| c1_log | pair_371 | 0.003 | 1912.533 | 14.383 | 13.526 | 28.437 | 24.118 | 7.920 |
| c1_log | pair_372 | 0.003 | 1912.533 | 8.561 | 8.794 | 17.607 | 15.924 | 2.981 |
| c1_log | pair_373 | 0.003 | 1912.533 | 12.481 | 12.491 | 25.353 | 23.701 | 2.909 |
| c1_log | pair_374 | 0.003 | 1912.533 | 6.320 | 6.431 | 13.005 | 12.802 | 0.428 |
| c1_log | pair_375 | 0.003 | 1912.533 | 11.951 | 12.738 | 25.117 | 24.210 | 1.768 |
| c1_log | pair_376 | 0.003 | 1912.533 | 14.378 | 12.189 | 29.195 | 30.741 | 4.269 |
| c1_log | pair_377 | 0.003 | 1912.533 | 9.100 | 6.404 | 15.762 | 14.396 | 3.909 |
| c1_log | pair_378 | 0.003 | 1912.533 | 6.382 | 6.236 | 12.868 | 12.942 | 0.141 |
| c1_log | pair_379 | 0.003 | 1912.533 | 4.020 | 3.904 | 8.025 | 8.005 | 0.055 |
| c1_log | pair_380 | 0.003 | 1912.533 | 9.245 | 7.427 | 16.903 | 14.483 | 4.355 |
| c1_log | pair_381 | 0.003 | 1912.533 | 9.438 | 7.796 | 17.446 | 15.895 | 2.732 |
| c1_log | pair_382 | 0.003 | 1912.533 | 7.892 | 7.823 | 15.980 | 15.877 | 0.197 |
| c1_log | pair_383 | 0.003 | 1912.533 | 7.840 | 7.996 | 16.119 | 16.139 | 0.231 |
| c1_log | pair_384 | 0.003 | 1912.533 | 6.256 | 6.388 | 12.889 | 12.974 | 0.243 |
| c1_log | pair_385 | 0.003 | 1912.533 | 7.845 | 10.215 | 18.409 | 15.939 | 4.490 |
| c1_log | pair_386 | 0.003 | 1912.533 | 7.855 | 7.868 | 17.411 | 16.204 | 2.337 |
| c1_log | pair_387 | 0.003 | 1912.533 | 6.305 | 6.417 | 12.976 | 12.736 | 0.572 |
| c1_log | pair_388 | 0.003 | 1912.533 | 7.672 | 8.758 | 16.607 | 13.281 | 6.305 |
| c1_log | pair_389 | 0.003 | 1912.533 | 4.124 | 4.003 | 8.257 | 8.218 | 0.321 |
| c1_log | pair_390 | 0.003 | 1912.533 | 11.021 | 8.433 | 19.674 | 14.269 | 9.461 |
| c1_log | pair_391 | 0.003 | 1912.533 | 10.060 | 9.184 | 19.530 | 16.178 | 5.882 |
| c1_log | pair_392 | 0.003 | 1912.533 | 24.069 | 19.998 | 44.868 | 40.187 | 8.419 |
| c1_log | pair_393 | 0.003 | 1912.533 | 10.677 | 11.080 | 21.980 | 13.040 | 15.789 |
| c1_log | pair_394 | 0.003 | 1912.533 | 12.620 | 12.573 | 25.611 | 25.260 | 0.638 |
| c1_log | pair_395 | 0.003 | 1912.533 | 20.359 | 20.392 | 41.581 | 41.345 | 0.569 |
| c1_log | pair_396 | 0.003 | 1912.533 | 6.577 | 6.999 | 13.808 | 12.747 | 1.841 |
| c1_log | pair_397 | 0.003 | 1912.533 | 6.280 | 6.246 | 12.686 | 12.623 | 0.129 |
| c1_log | pair_398 | 0.003 | 1912.533 | 3.967 | 3.922 | 8.060 | 8.044 | 0.035 |
| c1_log | pair_399 | 0.003 | 1912.533 | 7.035 | 7.114 | 14.376 | 14.271 | 0.258 |
| c1_log | pair_400 | 0.003 | 1912.533 | 11.598 | 8.395 | 20.244 | 15.966 | 7.585 |
| c1_log | pair_401 | 0.003 | 1912.533 | 6.261 | 6.354 | 12.843 | 12.827 | 0.190 |
| c1_log | pair_402 | 0.003 | 1912.533 | 7.858 | 7.770 | 17.288 | 16.094 | 2.271 |
| c1_log | pair_403 | 0.003 | 1912.533 | 7.868 | 7.751 | 15.832 | 15.877 | 0.114 |
| c1_log | pair_404 | 0.003 | 1912.533 | 7.463 | 7.183 | 14.947 | 15.548 | 1.510 |
| c1_log | pair_405 | 0.003 | 1912.533 | 6.272 | 6.313 | 12.843 | 12.823 | 0.041 |
| c1_log | pair_406 | 0.003 | 1912.533 | 5.733 | 4.352 | 10.249 | 8.150 | 3.865 |
| c1_log | pair_407 | 0.003 | 1912.533 | 7.087 | 7.715 | 15.091 | 14.383 | 1.359 |
| c1_log | pair_408 | 0.003 | 1912.533 | 6.363 | 6.305 | 12.894 | 12.810 | 0.199 |
| c1_log | pair_409 | 0.003 | 1912.533 | 13.003 | 12.668 | 26.121 | 25.886 | 0.557 |
| c1_log | pair_410 | 0.003 | 1912.533 | 25.453 | 19.587 | 45.768 | 40.659 | 9.628 |
| c1_log | pair_411 | 0.003 | 1912.533 | 6.297 | 6.315 | 12.806 | 12.752 | 0.120 |
| c1_log | pair_412 | 0.003 | 1912.533 | 8.796 | 6.246 | 15.306 | 12.736 | 4.472 |
| c1_log | pair_413 | 0.003 | 1912.533 | 3.950 | 3.969 | 8.064 | 8.007 | 0.105 |
| c1_log | pair_414 | 0.003 | 1912.533 | 7.050 | 8.331 | 15.722 | 14.365 | 2.518 |
| c1_log | pair_415 | 0.003 | 1912.533 | 6.345 | 6.340 | 12.899 | 12.768 | 0.356 |
| c1_log | pair_416 | 0.003 | 1912.533 | 6.250 | 6.357 | 12.813 | 12.747 | 0.399 |
| c1_log | pair_417 | 0.003 | 1912.533 | 6.340 | 6.343 | 12.863 | 12.613 | 0.437 |
| c1_log | pair_418 | 0.003 | 1912.533 | 6.231 | 6.340 | 12.820 | 12.717 | 0.209 |
| c1_log | pair_419 | 0.003 | 1912.533 | 3.928 | 3.892 | 7.934 | 7.952 | 0.116 |
| c1_log | pair_420 | 0.003 | 1912.533 | 6.345 | 6.272 | 12.826 | 12.773 | 0.126 |
| c1_log | pair_421 | 0.003 | 1912.533 | 19.685 | 12.549 | 32.677 | 35.605 | 6.436 |
| c1_log | pair_422 | 0.003 | 1912.533 | 6.317 | 6.319 | 12.842 | 12.753 | 0.259 |
| c1_log | pair_423 | 0.003 | 1912.533 | 6.322 | 6.341 | 12.873 | 12.740 | 0.296 |
| c1_log | pair_424 | 0.003 | 1912.533 | 3.937 | 3.873 | 7.989 | 7.983 | 0.038 |
| c1_log | pair_425 | 0.003 | 1912.533 | 8.014 | 7.015 | 15.682 | 15.697 | 1.246 |
| c1_log | pair_426 | 0.003 | 1912.533 | 9.081 | 6.358 | 15.665 | 15.651 | 2.955 |
| c1_log | pair_427 | 0.003 | 1912.533 | 7.830 | 6.819 | 14.882 | 13.073 | 3.413 |
| c1_log | pair_428 | 0.003 | 1912.533 | 4.009 | 3.947 | 8.102 | 8.092 | 0.080 |
| c1_log | pair_429 | 0.003 | 1912.533 | 7.160 | 7.069 | 14.494 | 14.416 | 0.177 |
| c1_log | pair_430 | 0.003 | 1912.533 | 6.313 | 6.244 | 12.810 | 12.799 | 0.047 |
| c1_log | pair_431 | 0.003 | 1912.533 | 3.929 | 3.929 | 8.056 | 8.121 | 0.145 |
| c1_log | pair_432 | 0.003 | 1912.533 | 6.238 | 6.250 | 12.725 | 12.680 | 0.081 |
| c1_log | pair_433 | 0.003 | 1912.533 | 3.910 | 3.924 | 7.951 | 7.918 | 0.066 |
| c1_log | pair_434 | 0.003 | 1912.533 | 6.280 | 6.293 | 12.797 | 12.765 | 0.057 |
| c1_log | pair_435 | 0.003 | 1912.533 | 3.935 | 3.916 | 8.132 | 8.108 | 0.246 |
| c1_power2 | pair_001 | 0.004 | 1869.593 | 6.750 | 6.459 | 13.292 | 13.486 | 0.536 |
| c1_power2 | pair_002 | 0.004 | 1869.593 | 7.116 | 6.304 | 13.672 | 13.841 | 0.752 |
| c1_power2 | pair_003 | 0.004 | 1869.593 | 6.335 | 6.313 | 12.772 | 12.836 | 0.206 |
| c1_power2 | pair_004 | 0.004 | 1869.593 | 6.876 | 6.766 | 13.716 | 13.043 | 1.628 |
| c1_power2 | pair_005 | 0.004 | 1869.593 | 3.173 | 3.128 | 6.355 | 6.366 | 0.052 |
| c1_power2 | pair_006 | 0.004 | 1869.593 | 6.321 | 6.283 | 12.742 | 12.791 | 0.093 |
| c1_power2 | pair_007 | 0.004 | 1869.593 | 12.738 | 12.527 | 25.493 | 25.502 | 0.175 |
| c1_power2 | pair_008 | 0.004 | 1869.593 | 6.297 | 6.432 | 12.872 | 12.656 | 0.447 |
| c1_power2 | pair_009 | 0.004 | 1869.593 | 25.215 | 24.918 | 51.494 | 51.276 | 0.469 |
| c1_power2 | pair_010 | 0.004 | 1869.593 | 6.359 | 6.216 | 12.705 | 12.692 | 0.113 |
| c1_power2 | pair_011 | 0.004 | 1869.593 | 6.313 | 6.250 | 12.769 | 12.660 | 0.213 |
| c1_power2 | pair_012 | 0.004 | 1869.593 | 6.424 | 6.223 | 12.858 | 12.678 | 0.318 |
| c1_power2 | pair_013 | 0.004 | 1869.593 | 6.329 | 6.230 | 12.695 | 12.722 | 0.050 |
| c1_power2 | pair_014 | 0.004 | 1869.593 | 6.388 | 6.207 | 12.789 | 12.772 | 0.124 |
| c1_power2 | pair_015 | 0.004 | 1869.593 | 6.279 | 6.297 | 12.716 | 12.788 | 0.124 |
| c1_power2 | pair_016 | 0.004 | 1869.593 | 3.149 | 3.084 | 6.314 | 6.294 | 0.078 |
| c1_power2 | pair_017 | 0.004 | 1869.593 | 3.133 | 3.145 | 6.343 | 6.364 | 0.049 |
| c1_power2 | pair_018 | 0.004 | 1869.593 | 17.113 | 14.040 | 31.589 | 25.345 | 10.928 |
| c1_power2 | pair_019 | 0.004 | 1869.593 | 6.287 | 6.307 | 12.712 | 12.728 | 0.139 |
| c1_power2 | pair_020 | 0.004 | 1869.593 | 26.010 | 31.665 | 59.186 | 51.841 | 13.101 |
| c1_power2 | pair_021 | 0.004 | 1869.593 | 6.587 | 6.556 | 13.283 | 12.769 | 0.914 |
| c1_power2 | pair_022 | 0.004 | 1869.593 | 26.142 | 26.468 | 54.028 | 51.790 | 4.207 |
| c1_power2 | pair_023 | 0.004 | 1869.593 | 6.683 | 6.614 | 13.556 | 12.901 | 1.358 |
| c1_power2 | pair_024 | 0.004 | 1869.593 | 13.018 | 12.886 | 26.365 | 25.357 | 1.790 |
| c1_power2 | pair_025 | 0.004 | 1869.593 | 26.086 | 24.996 | 52.462 | 51.648 | 1.650 |
| c1_power2 | pair_026 | 0.004 | 1869.593 | 6.745 | 6.269 | 13.238 | 12.749 | 0.937 |
| c1_power2 | pair_027 | 0.004 | 1869.593 | 6.453 | 6.574 | 13.243 | 12.730 | 0.991 |
| c1_power2 | pair_028 | 0.004 | 1869.593 | 3.352 | 3.314 | 6.713 | 6.412 | 0.554 |
| c1_power2 | pair_029 | 0.004 | 1869.593 | 6.710 | 6.567 | 13.460 | 12.859 | 1.181 |
| c1_power2 | pair_030 | 0.004 | 1869.593 | 6.674 | 6.701 | 13.558 | 12.663 | 1.572 |
| c1_power2 | pair_031 | 0.004 | 1869.593 | 6.401 | 6.289 | 12.785 | 12.649 | 0.261 |
| c1_power2 | pair_032 | 0.004 | 1869.593 | 6.763 | 6.638 | 13.476 | 13.390 | 0.995 |
| c1_power2 | pair_033 | 0.004 | 1869.593 | 3.263 | 3.119 | 6.468 | 6.431 | 0.111 |
| c1_power2 | pair_034 | 0.004 | 1869.593 | 6.358 | 6.221 | 12.728 | 12.785 | 0.125 |
| c1_power2 | pair_035 | 0.004 | 1869.593 | 6.213 | 6.199 | 12.489 | 12.477 | 0.026 |
| c1_power2 | pair_036 | 0.004 | 1869.593 | 6.285 | 6.265 | 12.673 | 12.698 | 0.093 |
| c1_power2 | pair_037 | 0.004 | 1869.593 | 6.301 | 6.204 | 12.579 | 12.576 | 0.117 |
| c1_power2 | pair_038 | 0.004 | 1869.593 | 6.229 | 6.265 | 12.614 | 12.589 | 0.075 |
| c1_power2 | pair_039 | 0.004 | 1869.593 | 6.257 | 6.277 | 12.791 | 12.784 | 0.046 |
| c1_power2 | pair_040 | 0.004 | 1869.593 | 6.329 | 6.268 | 12.785 | 12.742 | 0.131 |
| c1_power2 | pair_041 | 0.004 | 1869.593 | 6.248 | 10.947 | 17.347 | 12.600 | 8.250 |
| c1_power2 | pair_042 | 0.004 | 1869.593 | 6.285 | 6.276 | 12.723 | 12.639 | 0.149 |
| c1_power2 | pair_043 | 0.004 | 1869.593 | 6.419 | 6.677 | 13.222 | 12.670 | 1.035 |
| c1_power2 | pair_044 | 0.004 | 1869.593 | 3.185 | 3.137 | 6.411 | 6.451 | 0.097 |
| c1_power2 | pair_045 | 0.004 | 1869.593 | 3.148 | 3.101 | 6.333 | 6.336 | 0.013 |
| c1_power2 | pair_046 | 0.004 | 1869.593 | 6.324 | 6.188 | 12.579 | 12.541 | 0.104 |
| c1_power2 | pair_047 | 0.004 | 1869.593 | 6.429 | 6.424 | 12.985 | 13.014 | 0.380 |
| c1_power2 | pair_048 | 0.004 | 1869.593 | 6.378 | 6.216 | 12.663 | 12.654 | 0.121 |
| c1_power2 | pair_049 | 0.004 | 1869.593 | 6.264 | 6.200 | 12.593 | 12.586 | 0.048 |
| c1_power2 | pair_050 | 0.004 | 1869.593 | 6.247 | 6.220 | 12.536 | 12.561 | 0.059 |
| c1_power2 | pair_051 | 0.004 | 1869.593 | 6.250 | 6.250 | 12.670 | 12.697 | 0.111 |
| c1_power2 | pair_052 | 0.004 | 1869.593 | 6.227 | 6.311 | 12.615 | 12.648 | 0.095 |
| c1_power2 | pair_053 | 0.004 | 1869.593 | 6.224 | 6.263 | 12.557 | 12.517 | 0.070 |
| c1_power2 | pair_054 | 0.004 | 1869.593 | 6.358 | 6.229 | 12.841 | 12.746 | 0.199 |
| c1_power2 | pair_055 | 0.004 | 1869.593 | 6.319 | 6.311 | 12.869 | 12.872 | 0.041 |
| c1_power2 | pair_056 | 0.004 | 1869.593 | 3.104 | 3.118 | 6.302 | 6.316 | 0.068 |
| c1_power2 | pair_057 | 0.004 | 1869.593 | 6.248 | 6.213 | 12.617 | 12.624 | 0.075 |
| c1_power2 | pair_058 | 0.004 | 1869.593 | 6.242 | 6.234 | 12.631 | 12.637 | 0.068 |
| c1_power2 | pair_059 | 0.004 | 1869.593 | 6.295 | 6.356 | 12.790 | 12.672 | 0.230 |
| c1_power2 | pair_060 | 0.004 | 1869.593 | 3.158 | 3.102 | 6.359 | 6.338 | 0.045 |
| c1_power2 | pair_061 | 0.004 | 1869.593 | 6.455 | 7.102 | 13.819 | 13.043 | 1.631 |
| c1_power2 | pair_062 | 0.004 | 1869.593 | 6.388 | 6.204 | 16.247 | 12.927 | 5.984 |
| c1_power2 | pair_063 | 0.004 | 1869.593 | 6.360 | 6.237 | 12.758 | 12.771 | 0.083 |
| c1_power2 | pair_064 | 0.004 | 1869.593 | 7.634 | 6.262 | 14.064 | 12.885 | 2.207 |
| c1_power2 | pair_065 | 0.004 | 1869.593 | 6.722 | 6.313 | 13.214 | 13.038 | 0.490 |
| c1_power2 | pair_066 | 0.004 | 1869.593 | 6.285 | 7.012 | 13.581 | 12.751 | 1.491 |
| c1_power2 | pair_067 | 0.004 | 1869.593 | 6.412 | 6.260 | 12.844 | 12.631 | 0.424 |
| c1_power2 | pair_068 | 0.004 | 1869.593 | 6.275 | 6.302 | 12.763 | 12.815 | 0.125 |
| c1_power2 | pair_069 | 0.004 | 1869.593 | 6.288 | 6.261 | 12.739 | 12.699 | 0.236 |
| c1_power2 | pair_070 | 0.004 | 1869.593 | 6.675 | 6.410 | 13.223 | 12.982 | 0.513 |
| c1_power2 | pair_071 | 0.004 | 1869.593 | 3.189 | 3.118 | 6.368 | 6.371 | 0.033 |
| c1_power2 | pair_072 | 0.004 | 1869.593 | 3.176 | 3.557 | 6.854 | 6.488 | 0.728 |
| c1_power2 | pair_073 | 0.004 | 1869.593 | 7.037 | 6.647 | 13.854 | 12.976 | 1.685 |
| c1_power2 | pair_074 | 0.004 | 1869.593 | 6.393 | 6.260 | 12.812 | 12.815 | 0.089 |
| c1_power2 | pair_075 | 0.004 | 1869.593 | 6.744 | 6.558 | 13.504 | 12.969 | 1.154 |
| c1_power2 | pair_076 | 0.004 | 1869.593 | 6.471 | 6.530 | 13.150 | 12.905 | 0.570 |
| c1_power2 | pair_077 | 0.004 | 1869.593 | 6.305 | 6.278 | 15.149 | 12.844 | 4.220 |
| c1_power2 | pair_078 | 0.004 | 1869.593 | 9.110 | 10.115 | 19.439 | 12.956 | 11.496 |
| c1_power2 | pair_079 | 0.004 | 1869.593 | 7.994 | 8.783 | 17.107 | 12.672 | 7.700 |
| c1_power2 | pair_080 | 0.004 | 1869.593 | 9.781 | 9.166 | 19.338 | 12.770 | 11.498 |
| c1_power2 | pair_081 | 0.004 | 1869.593 | 6.761 | 6.205 | 13.194 | 12.795 | 0.866 |
| c1_power2 | pair_082 | 0.004 | 1869.593 | 6.263 | 6.308 | 12.796 | 12.769 | 0.305 |
| c1_power2 | pair_083 | 0.004 | 1869.593 | 3.186 | 3.160 | 6.401 | 6.369 | 0.152 |
| c1_power2 | pair_084 | 0.004 | 1869.593 | 6.474 | 6.382 | 13.022 | 13.169 | 0.282 |
| c1_power2 | pair_085 | 0.004 | 1869.593 | 6.388 | 6.258 | 12.770 | 12.763 | 0.134 |
| c1_power2 | pair_086 | 0.004 | 1869.593 | 3.215 | 3.115 | 6.392 | 6.384 | 0.026 |
| c1_power2 | pair_087 | 0.004 | 1869.593 | 6.328 | 6.339 | 12.847 | 12.919 | 0.202 |
| c1_power2 | pair_088 | 0.004 | 1869.593 | 6.588 | 6.399 | 13.101 | 12.817 | 0.511 |
| c1_power2 | pair_089 | 0.004 | 1869.593 | 6.385 | 6.493 | 13.021 | 12.812 | 0.423 |
| c1_power2 | pair_090 | 0.004 | 1869.593 | 6.350 | 6.315 | 12.794 | 12.796 | 0.156 |
| c1_power2 | pair_091 | 0.004 | 1869.593 | 6.404 | 6.251 | 12.786 | 12.801 | 0.084 |
| c1_power2 | pair_092 | 0.004 | 1869.593 | 6.502 | 6.370 | 13.070 | 12.936 | 0.255 |
| c1_power2 | pair_093 | 0.004 | 1869.593 | 6.333 | 6.277 | 12.795 | 12.739 | 0.181 |
| c1_power2 | pair_094 | 0.004 | 1869.593 | 6.422 | 6.383 | 12.957 | 12.986 | 0.116 |
| c1_power2 | pair_095 | 0.004 | 1869.593 | 6.500 | 7.543 | 14.237 | 12.911 | 2.479 |
| c1_power2 | pair_096 | 0.004 | 1869.593 | 6.293 | 6.275 | 12.723 | 12.697 | 0.066 |
| c1_power2 | pair_097 | 0.004 | 1869.593 | 4.669 | 4.787 | 9.531 | 6.297 | 5.611 |
| c1_power2 | pair_098 | 0.004 | 1869.593 | 3.218 | 3.303 | 6.607 | 6.370 | 0.426 |
| c1_power2 | pair_099 | 0.004 | 1869.593 | 6.433 | 6.471 | 13.023 | 12.873 | 0.588 |
| c1_power2 | pair_100 | 0.004 | 1869.593 | 8.428 | 14.913 | 23.577 | 14.087 | 17.708 |
| c1_power2 | pair_101 | 0.004 | 1869.593 | 6.448 | 6.422 | 12.990 | 13.157 | 0.344 |
| c1_power2 | pair_102 | 0.004 | 1869.593 | 7.613 | 7.012 | 14.756 | 14.396 | 1.689 |
| c1_power2 | pair_103 | 0.004 | 1869.593 | 6.936 | 6.888 | 14.066 | 14.444 | 1.199 |
| c1_power2 | pair_104 | 0.004 | 1869.593 | 6.618 | 6.463 | 13.371 | 12.756 | 1.076 |
| c1_power2 | pair_105 | 0.004 | 1869.593 | 6.545 | 6.463 | 13.131 | 12.792 | 0.595 |
| c1_power2 | pair_106 | 0.004 | 1869.593 | 6.303 | 6.773 | 13.208 | 12.680 | 0.983 |
| c1_power2 | pair_107 | 0.004 | 1869.593 | 6.241 | 6.532 | 12.991 | 12.702 | 0.611 |
| c1_power2 | pair_108 | 0.004 | 1869.593 | 6.481 | 6.506 | 13.210 | 13.387 | 0.368 |
| c1_power2 | pair_109 | 0.004 | 1869.593 | 3.145 | 3.116 | 6.329 | 6.306 | 0.061 |
| c1_power2 | pair_110 | 0.004 | 1869.593 | 6.276 | 6.220 | 12.656 | 12.649 | 0.066 |
| c1_power2 | pair_111 | 0.004 | 1869.593 | 3.158 | 3.094 | 6.339 | 6.341 | 0.011 |
| c1_power2 | pair_112 | 0.004 | 1869.593 | 6.432 | 6.242 | 12.836 | 12.736 | 0.282 |
| c1_power2 | pair_113 | 0.004 | 1869.593 | 6.243 | 6.240 | 12.553 | 12.543 | 0.040 |
| c1_power2 | pair_114 | 0.004 | 1869.593 | 6.424 | 6.271 | 12.810 | 12.613 | 0.380 |
| c1_power2 | pair_115 | 0.004 | 1869.593 | 6.581 | 6.343 | 12.998 | 12.692 | 0.705 |
| c1_power2 | pair_116 | 0.004 | 1869.593 | 6.394 | 6.354 | 12.887 | 12.676 | 0.465 |
| c1_power2 | pair_117 | 0.004 | 1869.593 | 6.421 | 6.394 | 13.074 | 12.850 | 0.518 |
| c1_power2 | pair_118 | 0.004 | 1869.593 | 6.402 | 6.498 | 13.115 | 13.047 | 0.423 |
| c1_power2 | pair_119 | 0.004 | 1869.593 | 6.429 | 6.538 | 13.255 | 12.830 | 0.847 |
| c1_power2 | pair_120 | 0.004 | 1869.593 | 7.095 | 6.369 | 13.636 | 12.732 | 1.728 |
| c1_power2 | pair_121 | 0.004 | 1869.593 | 6.430 | 7.883 | 14.474 | 13.350 | 2.679 |
| c1_power2 | pair_122 | 0.004 | 1869.593 | 3.255 | 3.239 | 6.542 | 6.617 | 0.184 |
| c1_power2 | pair_123 | 0.004 | 1869.593 | 3.215 | 3.218 | 6.475 | 6.472 | 0.184 |
| c1_power2 | pair_124 | 0.004 | 1869.593 | 6.329 | 6.298 | 12.697 | 12.715 | 0.272 |
| c1_power2 | pair_125 | 0.004 | 1869.593 | 8.809 | 6.725 | 15.674 | 13.490 | 4.506 |
| c1_power2 | pair_126 | 0.004 | 1869.593 | 7.574 | 6.803 | 14.452 | 12.517 | 3.394 |
| c1_power2 | pair_127 | 0.004 | 1869.593 | 6.451 | 6.221 | 12.787 | 12.847 | 0.243 |
| c1_power2 | pair_128 | 0.004 | 1869.593 | 6.226 | 6.201 | 12.501 | 12.546 | 0.106 |
| c1_power2 | pair_129 | 0.004 | 1869.593 | 6.268 | 6.233 | 12.750 | 12.752 | 0.059 |
| c1_power2 | pair_130 | 0.004 | 1869.593 | 6.211 | 6.255 | 12.538 | 12.482 | 0.180 |
| c1_power2 | pair_131 | 0.004 | 1869.593 | 6.228 | 6.218 | 12.515 | 12.493 | 0.050 |
| c1_power2 | pair_132 | 0.004 | 1869.593 | 6.260 | 6.279 | 12.728 | 12.833 | 0.216 |
| c1_power2 | pair_133 | 0.004 | 1869.593 | 6.330 | 6.270 | 14.013 | 12.969 | 1.971 |
| c1_power2 | pair_134 | 0.004 | 1869.593 | 3.370 | 3.331 | 6.792 | 6.340 | 0.805 |
| c1_power2 | pair_135 | 0.004 | 1869.593 | 6.745 | 6.832 | 13.923 | 12.864 | 1.921 |
| c1_power2 | pair_136 | 0.004 | 1869.593 | 3.107 | 3.070 | 6.213 | 6.201 | 0.040 |
| c1_power2 | pair_137 | 0.004 | 1869.593 | 3.153 | 3.096 | 6.346 | 6.324 | 0.077 |
| c1_power2 | pair_138 | 0.004 | 1869.593 | 3.133 | 3.082 | 6.277 | 6.265 | 0.027 |
| c1_power2 | pair_139 | 0.004 | 1869.593 | 3.176 | 3.123 | 6.366 | 6.365 | 0.016 |
| c1_power2 | pair_140 | 0.004 | 1869.593 | 3.149 | 3.083 | 6.297 | 6.293 | 0.031 |
| c1_power2 | pair_141 | 0.004 | 1869.593 | 3.138 | 3.079 | 6.315 | 6.347 | 0.063 |
| c1_power2 | pair_142 | 0.004 | 1869.593 | 3.127 | 3.091 | 6.292 | 6.324 | 0.069 |
| c1_power2 | pair_143 | 0.004 | 1869.593 | 3.147 | 3.114 | 6.314 | 6.290 | 0.076 |
| c1_power2 | pair_144 | 0.004 | 1869.593 | 3.100 | 3.064 | 6.256 | 6.255 | 0.058 |
| c1_power2 | pair_145 | 0.004 | 1869.593 | 3.101 | 3.093 | 6.267 | 6.308 | 0.071 |
| c1_power2 | pair_146 | 0.004 | 1869.593 | 3.126 | 3.093 | 6.270 | 6.283 | 0.030 |
| c1_power2 | pair_147 | 0.004 | 1869.593 | 3.160 | 3.063 | 6.283 | 6.293 | 0.021 |
| c1_power2 | pair_148 | 0.004 | 1869.593 | 3.124 | 3.102 | 6.283 | 6.261 | 0.087 |
| c1_power2 | pair_149 | 0.004 | 1869.593 | 3.324 | 3.138 | 6.548 | 6.430 | 0.352 |
| c1_power2 | pair_150 | 0.004 | 1869.593 | 3.128 | 3.135 | 6.355 | 6.432 | 0.166 |
| c1_power2 | pair_151 | 0.004 | 1869.593 | 3.161 | 3.081 | 6.322 | 6.335 | 0.081 |
| c1_power2 | pair_152 | 0.004 | 1869.593 | 3.137 | 3.160 | 6.348 | 6.295 | 0.098 |
| c1_power2 | pair_153 | 0.004 | 1869.593 | 3.151 | 3.075 | 6.301 | 6.289 | 0.057 |
| c1_power2 | pair_154 | 0.004 | 1869.593 | 3.172 | 3.135 | 6.364 | 6.361 | 0.052 |
| c1_power2 | pair_155 | 0.004 | 1869.593 | 3.123 | 3.116 | 6.290 | 6.289 | 0.007 |
| c1_power2 | pair_156 | 0.004 | 1869.593 | 3.145 | 3.131 | 6.354 | 6.382 | 0.075 |
| c1_power2 | pair_157 | 0.004 | 1869.593 | 3.146 | 3.148 | 6.360 | 6.361 | 0.055 |
| c1_power2 | pair_158 | 0.004 | 1869.593 | 3.193 | 3.103 | 6.352 | 6.365 | 0.025 |
| c1_power2 | pair_159 | 0.004 | 1869.593 | 3.138 | 3.104 | 6.327 | 6.341 | 0.061 |
| c1_power2 | pair_160 | 0.004 | 1869.593 | 6.257 | 6.257 | 12.666 | 12.660 | 0.021 |
| c1_power2 | pair_161 | 0.004 | 1869.593 | 6.344 | 6.199 | 12.716 | 12.688 | 0.184 |
| c1_power2 | pair_162 | 0.004 | 1869.593 | 6.296 | 6.179 | 12.630 | 12.646 | 0.039 |
| c1_power2 | pair_163 | 0.004 | 1869.593 | 6.274 | 6.194 | 12.663 | 12.671 | 0.029 |
| c1_power2 | pair_164 | 0.004 | 1869.593 | 6.258 | 6.311 | 12.785 | 12.608 | 0.320 |
| c1_power2 | pair_165 | 0.004 | 1869.593 | 6.323 | 6.181 | 12.721 | 12.719 | 0.010 |
| c1_power2 | pair_166 | 0.004 | 1869.593 | 9.753 | 10.601 | 20.783 | 24.044 | 7.092 |
| c1_power2 | pair_167 | 0.004 | 1869.593 | 9.355 | 6.469 | 16.053 | 12.843 | 5.778 |
| c1_power2 | pair_168 | 0.004 | 1869.593 | 7.494 | 6.215 | 13.864 | 12.767 | 2.082 |
| c1_power2 | pair_169 | 0.004 | 1869.593 | 3.148 | 3.116 | 6.338 | 6.329 | 0.069 |
| c1_power2 | pair_170 | 0.004 | 1869.593 | 3.131 | 3.125 | 6.323 | 6.331 | 0.051 |
| c1_power2 | pair_171 | 0.004 | 1869.593 | 6.226 | 6.233 | 12.631 | 12.631 | 0.105 |
| c1_power2 | pair_172 | 0.004 | 1869.593 | 6.283 | 6.240 | 12.719 | 12.670 | 0.120 |
| c1_power2 | pair_173 | 0.004 | 1869.593 | 6.275 | 6.261 | 12.710 | 12.686 | 0.091 |
| c1_power2 | pair_174 | 0.004 | 1869.593 | 6.244 | 6.253 | 12.649 | 12.686 | 0.100 |
| c1_power2 | pair_175 | 0.004 | 1869.593 | 6.239 | 6.474 | 12.925 | 12.710 | 0.513 |
| c1_power2 | pair_176 | 0.004 | 1869.593 | 6.252 | 6.262 | 12.734 | 12.719 | 0.054 |
| c1_power2 | pair_177 | 0.004 | 1869.593 | 6.448 | 6.243 | 12.867 | 12.685 | 0.378 |
| c1_power2 | pair_178 | 0.004 | 1869.593 | 6.362 | 6.233 | 12.732 | 12.581 | 0.291 |
| c1_power2 | pair_179 | 0.004 | 1869.593 | 6.260 | 6.186 | 12.666 | 12.627 | 0.118 |
| c1_power2 | pair_180 | 0.004 | 1869.593 | 6.188 | 6.270 | 14.572 | 12.697 | 3.344 |
| c1_power2 | pair_181 | 0.004 | 1869.593 | 3.105 | 3.116 | 6.282 | 6.313 | 0.067 |
| c1_power2 | pair_182 | 0.004 | 1869.593 | 6.351 | 6.181 | 12.705 | 12.710 | 0.047 |
| c1_power2 | pair_183 | 0.004 | 1869.593 | 6.323 | 6.195 | 12.630 | 12.670 | 0.076 |
| c1_power2 | pair_184 | 0.004 | 1869.593 | 12.738 | 12.632 | 25.610 | 25.416 | 0.540 |
| c1_power2 | pair_185 | 0.004 | 1869.593 | 6.258 | 6.297 | 12.643 | 12.623 | 0.063 |
| c1_power2 | pair_186 | 0.004 | 1869.593 | 6.261 | 6.274 | 12.756 | 12.760 | 0.019 |
| c1_power2 | pair_187 | 0.004 | 1869.593 | 6.240 | 6.280 | 12.726 | 12.655 | 0.184 |
| c1_power2 | pair_188 | 0.004 | 1869.593 | 6.288 | 6.331 | 12.818 | 12.831 | 0.117 |
| c1_power2 | pair_189 | 0.004 | 1869.593 | 6.186 | 6.209 | 12.600 | 12.578 | 0.074 |
| c1_power2 | pair_190 | 0.004 | 1869.593 | 6.349 | 6.403 | 12.926 | 12.747 | 0.421 |
| c1_power2 | pair_191 | 0.004 | 1869.593 | 3.857 | 3.620 | 7.581 | 6.581 | 2.035 |
| c1_power2 | pair_192 | 0.004 | 1869.593 | 3.241 | 3.198 | 6.509 | 6.415 | 0.306 |
| c1_power2 | pair_193 | 0.004 | 1869.593 | 15.313 | 13.447 | 29.236 | 26.025 | 6.226 |
| c1_power2 | pair_194 | 0.004 | 1869.593 | 6.340 | 6.269 | 12.748 | 12.577 | 0.308 |
| c1_power2 | pair_195 | 0.004 | 1869.593 | 12.590 | 12.492 | 25.309 | 25.202 | 0.285 |
| c1_power2 | pair_196 | 0.004 | 1869.593 | 6.340 | 6.330 | 12.770 | 12.643 | 0.314 |
| c1_power2 | pair_197 | 0.004 | 1869.593 | 12.510 | 12.428 | 25.179 | 25.183 | 0.027 |
| c1_power2 | pair_198 | 0.004 | 1869.593 | 13.439 | 7.337 | 21.010 | 12.748 | 14.352 |
| c1_power2 | pair_199 | 0.004 | 1869.593 | 12.515 | 12.422 | 30.769 | 25.517 | 9.147 |
| c1_power2 | pair_200 | 0.004 | 1869.593 | 12.986 | 12.907 | 26.153 | 25.736 | 1.420 |
| c1_power2 | pair_201 | 0.004 | 1869.593 | 6.465 | 6.362 | 13.027 | 12.912 | 0.494 |
| c1_power2 | pair_202 | 0.004 | 1869.593 | 6.304 | 6.241 | 12.741 | 12.797 | 0.114 |
| c1_power2 | pair_203 | 0.004 | 1869.593 | 3.176 | 3.126 | 6.346 | 6.321 | 0.060 |
| c1_power2 | pair_204 | 0.004 | 1869.593 | 6.264 | 6.214 | 12.631 | 12.606 | 0.051 |
| c1_power2 | pair_205 | 0.004 | 1869.593 | 6.316 | 6.250 | 12.655 | 12.680 | 0.060 |
| c1_power2 | pair_206 | 0.004 | 1869.593 | 6.313 | 6.223 | 12.652 | 12.633 | 0.046 |
| c1_power2 | pair_207 | 0.004 | 1869.593 | 6.277 | 6.206 | 12.714 | 12.688 | 0.102 |
| c1_power2 | pair_208 | 0.004 | 1869.593 | 6.294 | 6.226 | 12.670 | 12.671 | 0.020 |
| c1_power2 | pair_209 | 0.004 | 1869.593 | 6.309 | 6.341 | 12.795 | 12.797 | 0.154 |
| c1_power2 | pair_210 | 0.004 | 1869.593 | 6.256 | 6.206 | 12.628 | 12.670 | 0.098 |
| c1_power2 | pair_211 | 0.004 | 1869.593 | 6.490 | 6.241 | 13.027 | 13.148 | 0.245 |
| c1_power2 | pair_212 | 0.004 | 1869.593 | 3.135 | 3.113 | 6.325 | 6.291 | 0.091 |
| c1_power2 | pair_213 | 0.004 | 1869.593 | 3.131 | 3.116 | 6.296 | 6.297 | 0.021 |
| c1_power2 | pair_214 | 0.004 | 1869.593 | 6.283 | 6.217 | 12.629 | 12.584 | 0.082 |
| c1_power2 | pair_215 | 0.004 | 1869.593 | 6.282 | 6.225 | 12.614 | 12.578 | 0.079 |
| c1_power2 | pair_216 | 0.004 | 1869.593 | 6.256 | 6.223 | 12.613 | 12.663 | 0.121 |
| c1_power2 | pair_217 | 0.004 | 1869.593 | 6.453 | 6.397 | 13.007 | 12.731 | 0.526 |
| c1_power2 | pair_218 | 0.004 | 1869.593 | 9.157 | 8.459 | 17.738 | 12.619 | 8.945 |
| c1_power2 | pair_219 | 0.004 | 1869.593 | 6.693 | 6.516 | 13.438 | 12.727 | 1.289 |
| c1_power2 | pair_220 | 0.004 | 1869.593 | 6.720 | 6.636 | 13.510 | 12.720 | 1.563 |
| c1_power2 | pair_221 | 0.004 | 1869.593 | 6.496 | 6.397 | 13.011 | 12.604 | 0.714 |
| c1_power2 | pair_222 | 0.004 | 1869.593 | 6.306 | 6.228 | 12.743 | 12.657 | 0.189 |
| c1_power2 | pair_223 | 0.004 | 1869.593 | 6.280 | 6.243 | 12.781 | 12.749 | 0.124 |
| c1_power2 | pair_224 | 0.004 | 1869.593 | 3.129 | 3.140 | 6.354 | 6.347 | 0.056 |
| c1_power2 | pair_225 | 0.004 | 1869.593 | 6.249 | 6.212 | 12.606 | 12.638 | 0.115 |
| c1_power2 | pair_226 | 0.004 | 1869.593 | 6.217 | 6.214 | 12.549 | 12.564 | 0.029 |
| c1_power2 | pair_227 | 0.004 | 1869.593 | 6.270 | 6.256 | 12.794 | 12.820 | 0.111 |
| c1_power2 | pair_228 | 0.004 | 1869.593 | 6.282 | 6.339 | 12.844 | 12.827 | 0.069 |
| c1_power2 | pair_229 | 0.004 | 1869.593 | 6.248 | 6.324 | 13.060 | 12.804 | 0.562 |
| c1_power2 | pair_230 | 0.004 | 1869.593 | 6.249 | 6.424 | 12.854 | 12.699 | 0.284 |
| c1_power2 | pair_231 | 0.004 | 1869.593 | 6.283 | 6.237 | 12.677 | 12.653 | 0.086 |
| c1_power2 | pair_232 | 0.004 | 1869.593 | 3.103 | 3.129 | 6.326 | 6.308 | 0.035 |
| c1_power2 | pair_233 | 0.004 | 1869.593 | 3.115 | 3.116 | 6.294 | 6.309 | 0.049 |
| c1_power2 | pair_234 | 0.004 | 1869.593 | 12.494 | 12.405 | 26.862 | 25.358 | 2.670 |
| c1_power2 | pair_235 | 0.004 | 1869.593 | 6.375 | 6.363 | 12.850 | 12.884 | 0.231 |
| c1_power2 | pair_236 | 0.004 | 1869.593 | 25.159 | 26.213 | 53.167 | 51.232 | 3.560 |
| c1_power2 | pair_237 | 0.004 | 1869.593 | 6.532 | 6.395 | 13.065 | 13.034 | 0.276 |
| c1_power2 | pair_238 | 0.004 | 1869.593 | 25.598 | 25.600 | 55.996 | 55.253 | 5.052 |
| c1_power2 | pair_239 | 0.004 | 1869.593 | 6.401 | 6.363 | 12.973 | 12.867 | 0.267 |
| c1_power2 | pair_240 | 0.004 | 1869.593 | 17.200 | 16.947 | 34.622 | 26.539 | 14.917 |
| c1_power2 | pair_241 | 0.004 | 1869.593 | 28.121 | 24.951 | 54.389 | 51.268 | 5.449 |
| c1_power2 | pair_242 | 0.004 | 1869.593 | 6.290 | 6.249 | 12.794 | 12.773 | 0.116 |
| c1_power2 | pair_243 | 0.004 | 1869.593 | 6.235 | 6.246 | 12.736 | 12.766 | 0.087 |
| c1_power2 | pair_244 | 0.004 | 1869.593 | 3.172 | 3.108 | 6.335 | 6.299 | 0.127 |
| c1_power2 | pair_245 | 0.004 | 1869.593 | 6.296 | 6.218 | 13.027 | 12.738 | 0.587 |
| c1_power2 | pair_246 | 0.004 | 1869.593 | 10.436 | 11.012 | 21.852 | 12.955 | 15.593 |
| c1_power2 | pair_247 | 0.004 | 1869.593 | 11.924 | 8.580 | 21.137 | 12.709 | 14.698 |
| c1_power2 | pair_248 | 0.004 | 1869.593 | 6.303 | 6.497 | 12.978 | 12.732 | 0.558 |
| c1_power2 | pair_249 | 0.004 | 1869.593 | 7.008 | 6.843 | 14.214 | 12.885 | 2.531 |
| c1_power2 | pair_250 | 0.004 | 1869.593 | 7.189 | 7.034 | 14.503 | 15.117 | 1.566 |
| c1_power2 | pair_251 | 0.004 | 1869.593 | 3.210 | 3.578 | 6.889 | 6.763 | 0.673 |
| c1_power2 | pair_252 | 0.004 | 1869.593 | 3.196 | 3.659 | 6.990 | 6.617 | 0.936 |
| c1_power2 | pair_253 | 0.004 | 1869.593 | 7.177 | 6.698 | 13.988 | 14.192 | 1.397 |
| c1_power2 | pair_254 | 0.004 | 1869.593 | 6.885 | 6.690 | 13.683 | 13.147 | 1.371 |
| c1_power2 | pair_255 | 0.004 | 1869.593 | 7.195 | 6.603 | 14.036 | 13.068 | 2.137 |
| c1_power2 | pair_256 | 0.004 | 1869.593 | 6.285 | 6.254 | 12.671 | 12.564 | 0.187 |
| c1_power2 | pair_257 | 0.004 | 1869.593 | 6.260 | 6.252 | 12.629 | 12.648 | 0.033 |
| c1_power2 | pair_258 | 0.004 | 1869.593 | 6.322 | 6.230 | 12.785 | 12.777 | 0.136 |
| c1_power2 | pair_259 | 0.004 | 1869.593 | 6.296 | 6.201 | 12.629 | 12.652 | 0.047 |
| c1_power2 | pair_260 | 0.004 | 1869.593 | 6.412 | 6.478 | 14.173 | 12.814 | 2.410 |
| c1_power2 | pair_261 | 0.004 | 1869.593 | 6.374 | 8.642 | 15.234 | 12.812 | 4.324 |
| c1_power2 | pair_262 | 0.004 | 1869.593 | 6.822 | 6.276 | 13.283 | 12.794 | 0.893 |
| c1_power2 | pair_263 | 0.004 | 1869.593 | 3.120 | 3.134 | 6.328 | 6.333 | 0.018 |
| c1_power2 | pair_264 | 0.004 | 1869.593 | 6.276 | 6.220 | 12.643 | 12.626 | 0.107 |
| c1_power2 | pair_265 | 0.004 | 1869.593 | 6.288 | 6.232 | 13.497 | 12.775 | 1.327 |
| c1_power2 | pair_266 | 0.004 | 1869.593 | 6.399 | 6.211 | 12.813 | 12.773 | 0.114 |
| c1_power2 | pair_267 | 0.004 | 1869.593 | 6.314 | 6.238 | 12.767 | 12.787 | 0.085 |
| c1_power2 | pair_268 | 0.004 | 1869.593 | 6.266 | 6.207 | 12.675 | 12.589 | 0.183 |
| c1_power2 | pair_269 | 0.004 | 1869.593 | 3.159 | 3.101 | 6.310 | 6.322 | 0.092 |
| c1_power2 | pair_270 | 0.004 | 1869.593 | 3.152 | 3.179 | 6.392 | 6.409 | 0.042 |
| c1_power2 | pair_271 | 0.004 | 1869.593 | 6.326 | 6.270 | 12.829 | 12.796 | 0.112 |
| c1_power2 | pair_272 | 0.004 | 1869.593 | 6.300 | 6.289 | 12.743 | 12.673 | 0.166 |
| c1_power2 | pair_273 | 0.004 | 1869.593 | 6.230 | 6.207 | 12.676 | 12.685 | 0.090 |
| c1_power2 | pair_274 | 0.004 | 1869.593 | 6.277 | 6.175 | 12.671 | 12.644 | 0.102 |
| c1_power2 | pair_275 | 0.004 | 1869.593 | 6.240 | 6.208 | 12.677 | 12.677 | 0.020 |
| c1_power2 | pair_276 | 0.004 | 1869.593 | 6.244 | 6.224 | 12.694 | 12.642 | 0.100 |
| c1_power2 | pair_277 | 0.004 | 1869.593 | 6.258 | 6.237 | 12.700 | 12.732 | 0.076 |
| c1_power2 | pair_278 | 0.004 | 1869.593 | 6.258 | 6.197 | 12.708 | 12.669 | 0.088 |
| c1_power2 | pair_279 | 0.004 | 1869.593 | 6.272 | 6.177 | 12.634 | 12.590 | 0.117 |
| c1_power2 | pair_280 | 0.004 | 1869.593 | 6.277 | 6.257 | 12.732 | 12.741 | 0.032 |
| c1_power2 | pair_281 | 0.004 | 1869.593 | 4.054 | 4.477 | 8.599 | 6.314 | 3.993 |
| c1_power2 | pair_282 | 0.004 | 1869.593 | 11.283 | 7.262 | 18.728 | 14.269 | 9.103 |
| c1_power2 | pair_283 | 0.004 | 1869.593 | 6.345 | 6.243 | 12.811 | 12.830 | 0.068 |
| c1_power2 | pair_284 | 0.004 | 1869.593 | 6.383 | 6.314 | 12.837 | 12.697 | 0.374 |
| c1_power2 | pair_285 | 0.004 | 1869.593 | 6.331 | 6.338 | 12.869 | 12.874 | 0.159 |
| c1_power2 | pair_286 | 0.004 | 1869.593 | 3.154 | 3.126 | 6.363 | 6.343 | 0.054 |
| c1_power2 | pair_287 | 0.004 | 1869.593 | 3.159 | 3.092 | 6.326 | 6.321 | 0.087 |
| c1_power2 | pair_288 | 0.004 | 1869.593 | 6.296 | 6.221 | 12.731 | 12.663 | 0.119 |
| c1_power2 | pair_289 | 0.004 | 1869.593 | 6.264 | 6.262 | 12.717 | 12.692 | 0.062 |
| c1_power2 | pair_290 | 0.004 | 1869.593 | 6.333 | 6.223 | 12.715 | 12.711 | 0.016 |
| c1_power2 | pair_291 | 0.004 | 1869.593 | 6.303 | 6.212 | 12.716 | 12.681 | 0.071 |
| c1_power2 | pair_292 | 0.004 | 1869.593 | 6.338 | 6.271 | 12.815 | 12.792 | 0.147 |
| c1_power2 | pair_293 | 0.004 | 1869.593 | 6.310 | 6.192 | 12.693 | 12.688 | 0.152 |
| c1_power2 | pair_294 | 0.004 | 1869.593 | 6.285 | 6.261 | 12.733 | 12.693 | 0.190 |
| c1_power2 | pair_295 | 0.004 | 1869.593 | 6.351 | 6.228 | 12.781 | 12.802 | 0.171 |
| c1_power2 | pair_296 | 0.004 | 1869.593 | 6.312 | 7.963 | 14.465 | 12.717 | 3.030 |
| c1_power2 | pair_297 | 0.004 | 1869.593 | 6.299 | 6.232 | 12.769 | 12.788 | 0.081 |
| c1_power2 | pair_298 | 0.004 | 1869.593 | 3.145 | 3.148 | 6.488 | 6.433 | 0.134 |
| c1_power2 | pair_299 | 0.004 | 1869.593 | 6.604 | 7.924 | 14.767 | 12.842 | 3.578 |
| c1_power2 | pair_300 | 0.004 | 1869.593 | 6.262 | 6.267 | 12.649 | 12.654 | 0.092 |
| c1_power2 | pair_301 | 0.004 | 1869.593 | 6.318 | 6.311 | 12.787 | 12.692 | 0.220 |
| c1_power2 | pair_302 | 0.004 | 1869.593 | 3.151 | 3.100 | 6.331 | 6.339 | 0.041 |
| c1_power2 | pair_303 | 0.004 | 1869.593 | 3.214 | 3.123 | 6.415 | 6.433 | 0.044 |
| c1_power2 | pair_304 | 0.004 | 1869.593 | 6.295 | 6.295 | 12.775 | 12.729 | 0.098 |
| c1_power2 | pair_305 | 0.004 | 1869.593 | 6.334 | 6.331 | 12.830 | 12.856 | 0.240 |
| c1_power2 | pair_306 | 0.004 | 1869.593 | 6.336 | 6.283 | 12.819 | 12.795 | 0.120 |
| c1_power2 | pair_307 | 0.004 | 1869.593 | 6.291 | 6.242 | 12.652 | 12.689 | 0.076 |
| c1_power2 | pair_308 | 0.004 | 1869.593 | 6.243 | 6.285 | 12.718 | 12.649 | 0.132 |
| c1_power2 | pair_309 | 0.004 | 1869.593 | 6.277 | 6.278 | 12.757 | 12.675 | 0.170 |
| c1_power2 | pair_310 | 0.004 | 1869.593 | 6.341 | 6.262 | 12.830 | 12.730 | 0.193 |
| c1_power2 | pair_311 | 0.004 | 1869.593 | 6.526 | 6.252 | 12.918 | 12.816 | 0.236 |
| c1_power2 | pair_312 | 0.004 | 1869.593 | 6.290 | 6.433 | 12.908 | 12.695 | 0.445 |
| c1_power2 | pair_313 | 0.004 | 1869.593 | 6.301 | 6.285 | 12.743 | 12.662 | 0.183 |
| c1_power2 | pair_314 | 0.004 | 1869.593 | 3.135 | 3.113 | 6.324 | 6.326 | 0.044 |
| c1_power2 | pair_315 | 0.004 | 1869.593 | 12.041 | 6.315 | 18.513 | 12.903 | 9.969 |
| c1_power2 | pair_316 | 0.004 | 1869.593 | 6.273 | 6.256 | 12.668 | 12.638 | 0.114 |
| c1_power2 | pair_317 | 0.004 | 1869.593 | 3.146 | 3.154 | 6.386 | 6.325 | 0.142 |
| c1_power2 | pair_318 | 0.004 | 1869.593 | 3.166 | 3.111 | 6.328 | 6.316 | 0.078 |
| c1_power2 | pair_319 | 0.004 | 1869.593 | 6.263 | 6.281 | 12.701 | 12.693 | 0.074 |
| c1_power2 | pair_320 | 0.004 | 1869.593 | 6.381 | 6.234 | 13.441 | 13.133 | 1.093 |
| c1_power2 | pair_321 | 0.004 | 1869.593 | 6.347 | 8.272 | 14.797 | 12.795 | 3.648 |
| c1_power2 | pair_322 | 0.004 | 1869.593 | 6.251 | 6.307 | 12.708 | 12.716 | 0.133 |
| c1_power2 | pair_323 | 0.004 | 1869.593 | 6.414 | 8.906 | 15.522 | 16.585 | 2.297 |
| c1_power2 | pair_324 | 0.004 | 1869.593 | 6.557 | 6.267 | 13.023 | 12.877 | 0.334 |
| c1_power2 | pair_325 | 0.004 | 1869.593 | 6.315 | 6.219 | 12.688 | 12.696 | 0.023 |
| c1_power2 | pair_326 | 0.004 | 1869.593 | 6.328 | 6.197 | 12.675 | 12.669 | 0.036 |
| c1_power2 | pair_327 | 0.004 | 1869.593 | 6.273 | 6.168 | 12.690 | 12.685 | 0.071 |
| c1_power2 | pair_328 | 0.004 | 1869.593 | 6.277 | 6.243 | 12.722 | 12.735 | 0.079 |
| c1_power2 | pair_329 | 0.004 | 1869.593 | 3.147 | 3.134 | 6.330 | 6.307 | 0.052 |
| c1_power2 | pair_330 | 0.004 | 1869.593 | 6.291 | 6.238 | 12.710 | 12.706 | 0.010 |
| c1_power2 | pair_331 | 0.004 | 1869.593 | 3.157 | 3.080 | 6.279 | 6.258 | 0.040 |
| c1_power2 | pair_332 | 0.004 | 1869.593 | 3.185 | 3.093 | 6.373 | 6.405 | 0.104 |
| c1_power2 | pair_333 | 0.004 | 1869.593 | 6.248 | 6.262 | 12.697 | 12.711 | 0.113 |
| c1_power2 | pair_334 | 0.004 | 1869.593 | 6.258 | 6.268 | 12.684 | 12.686 | 0.081 |
| c1_power2 | pair_335 | 0.004 | 1869.593 | 6.258 | 6.231 | 16.668 | 12.741 | 6.918 |
| c1_power2 | pair_336 | 0.004 | 1869.593 | 6.382 | 6.403 | 13.103 | 12.836 | 0.522 |
| c1_power2 | pair_337 | 0.004 | 1869.593 | 6.298 | 6.880 | 13.345 | 12.867 | 1.094 |
| c1_power2 | pair_338 | 0.004 | 1869.593 | 6.348 | 6.238 | 12.817 | 12.830 | 0.058 |
| c1_power2 | pair_339 | 0.004 | 1869.593 | 6.343 | 6.217 | 12.725 | 12.738 | 0.070 |
| c1_power2 | pair_340 | 0.004 | 1869.593 | 6.380 | 6.228 | 12.759 | 12.786 | 0.086 |
| c1_power2 | pair_341 | 0.004 | 1869.593 | 6.315 | 6.169 | 12.725 | 12.730 | 0.028 |
| c1_power2 | pair_342 | 0.004 | 1869.593 | 6.320 | 6.280 | 12.836 | 12.812 | 0.083 |
| c1_power2 | pair_343 | 0.004 | 1869.593 | 3.159 | 3.149 | 6.381 | 6.395 | 0.036 |
| c1_power2 | pair_344 | 0.004 | 1869.593 | 6.286 | 6.274 | 12.689 | 12.655 | 0.169 |
| c1_power2 | pair_345 | 0.004 | 1869.593 | 3.283 | 3.131 | 6.488 | 6.479 | 0.151 |
| c1_power2 | pair_346 | 0.004 | 1869.593 | 3.872 | 3.133 | 7.101 | 6.387 | 1.292 |
| c1_power2 | pair_347 | 0.004 | 1869.593 | 3.557 | 3.242 | 6.890 | 6.454 | 0.965 |
| c1_power2 | pair_348 | 0.004 | 1869.593 | 3.120 | 3.120 | 6.300 | 6.279 | 0.039 |
| c1_power2 | pair_349 | 0.004 | 1869.593 | 3.199 | 3.097 | 6.364 | 6.370 | 0.032 |
| c1_power2 | pair_350 | 0.004 | 1869.593 | 3.181 | 3.079 | 6.327 | 6.324 | 0.011 |
| c1_power2 | pair_351 | 0.004 | 1869.593 | 3.164 | 3.083 | 6.317 | 6.296 | 0.058 |
| c1_power2 | pair_352 | 0.004 | 1869.593 | 3.175 | 3.114 | 6.379 | 6.335 | 0.102 |
| c1_power2 | pair_353 | 0.004 | 1869.593 | 3.245 | 3.425 | 6.764 | 6.400 | 0.671 |
| c1_power2 | pair_354 | 0.004 | 1869.593 | 3.164 | 3.105 | 6.331 | 6.331 | 0.022 |
| c1_power2 | pair_355 | 0.004 | 1869.593 | 3.173 | 3.118 | 6.357 | 6.339 | 0.040 |
| c1_power2 | pair_356 | 0.004 | 1869.593 | 3.170 | 3.088 | 6.336 | 6.357 | 0.045 |
| c1_power2 | pair_357 | 0.004 | 1869.593 | 3.150 | 3.116 | 6.355 | 6.353 | 0.035 |
| c1_power2 | pair_358 | 0.004 | 1869.593 | 3.136 | 3.108 | 6.292 | 6.303 | 0.032 |
| c1_power2 | pair_359 | 0.004 | 1869.593 | 3.155 | 3.153 | 9.726 | 6.375 | 5.896 |
| c1_power2 | pair_360 | 0.004 | 1869.593 | 3.215 | 3.366 | 6.662 | 6.316 | 0.621 |
| c1_power2 | pair_361 | 0.004 | 1869.593 | 3.194 | 3.133 | 6.390 | 6.311 | 0.159 |
| c1_power2 | pair_362 | 0.004 | 1869.593 | 3.155 | 3.151 | 6.392 | 6.330 | 0.160 |
| c1_power2 | pair_363 | 0.004 | 1869.593 | 3.172 | 3.087 | 6.343 | 6.314 | 0.052 |
| c1_power2 | pair_364 | 0.004 | 1869.593 | 3.172 | 3.109 | 6.324 | 6.336 | 0.055 |
| c1_power2 | pair_365 | 0.004 | 1869.593 | 3.151 | 3.139 | 6.353 | 6.363 | 0.085 |
| c1_power2 | pair_366 | 0.004 | 1869.593 | 3.131 | 3.141 | 6.309 | 6.302 | 0.069 |
| c1_power2 | pair_367 | 0.004 | 1869.593 | 3.153 | 5.550 | 8.799 | 6.322 | 4.297 |
| c1_power2 | pair_368 | 0.004 | 1869.593 | 3.228 | 3.126 | 6.432 | 6.411 | 0.045 |
| c1_power2 | pair_369 | 0.004 | 1869.593 | 3.167 | 3.112 | 6.344 | 6.359 | 0.046 |
| c1_power2 | pair_370 | 0.004 | 1869.593 | 6.299 | 6.210 | 12.638 | 12.635 | 0.062 |
| c1_power2 | pair_371 | 0.004 | 1869.593 | 12.576 | 12.503 | 26.322 | 25.653 | 1.297 |
| c1_power2 | pair_372 | 0.004 | 1869.593 | 9.409 | 6.246 | 15.795 | 12.663 | 5.451 |
| c1_power2 | pair_373 | 0.004 | 1869.593 | 12.589 | 12.847 | 25.866 | 25.568 | 0.797 |
| c1_power2 | pair_374 | 0.004 | 1869.593 | 6.289 | 6.240 | 12.721 | 12.744 | 0.104 |
| c1_power2 | pair_375 | 0.004 | 1869.593 | 12.475 | 12.424 | 25.335 | 25.388 | 0.184 |
| c1_power2 | pair_376 | 0.004 | 1869.593 | 16.423 | 12.467 | 29.368 | 25.728 | 6.519 |
| c1_power2 | pair_377 | 0.004 | 1869.593 | 6.282 | 6.258 | 12.719 | 12.677 | 0.086 |
| c1_power2 | pair_378 | 0.004 | 1869.593 | 6.283 | 6.241 | 12.694 | 12.650 | 0.091 |
| c1_power2 | pair_379 | 0.004 | 1869.593 | 3.130 | 3.167 | 6.351 | 6.336 | 0.028 |
| c1_power2 | pair_380 | 0.004 | 1869.593 | 6.255 | 6.292 | 12.728 | 12.758 | 0.144 |
| c1_power2 | pair_381 | 0.004 | 1869.593 | 6.553 | 7.066 | 13.766 | 12.799 | 1.817 |
| c1_power2 | pair_382 | 0.004 | 1869.593 | 6.312 | 6.240 | 12.688 | 12.696 | 0.122 |
| c1_power2 | pair_383 | 0.004 | 1869.593 | 6.269 | 6.219 | 12.627 | 12.666 | 0.139 |
| c1_power2 | pair_384 | 0.004 | 1869.593 | 6.300 | 6.284 | 12.798 | 12.714 | 0.157 |
| c1_power2 | pair_385 | 0.004 | 1869.593 | 6.299 | 6.373 | 12.793 | 12.875 | 0.159 |
| c1_power2 | pair_386 | 0.004 | 1869.593 | 6.291 | 6.310 | 12.738 | 12.737 | 0.187 |
| c1_power2 | pair_387 | 0.004 | 1869.593 | 6.289 | 6.223 | 12.753 | 12.692 | 0.117 |
| c1_power2 | pair_388 | 0.004 | 1869.593 | 6.361 | 6.514 | 13.085 | 12.937 | 0.302 |
| c1_power2 | pair_389 | 0.004 | 1869.593 | 3.190 | 3.104 | 6.349 | 6.347 | 0.048 |
| c1_power2 | pair_390 | 0.004 | 1869.593 | 8.877 | 6.647 | 15.673 | 12.722 | 5.269 |
| c1_power2 | pair_391 | 0.004 | 1869.593 | 8.914 | 9.332 | 18.638 | 13.153 | 9.904 |
| c1_power2 | pair_392 | 0.004 | 1869.593 | 25.095 | 24.876 | 54.423 | 51.491 | 5.674 |
| c1_power2 | pair_393 | 0.004 | 1869.593 | 6.283 | 6.179 | 12.707 | 12.733 | 0.073 |
| c1_power2 | pair_394 | 0.004 | 1869.593 | 12.511 | 12.618 | 25.534 | 25.535 | 0.181 |
| c1_power2 | pair_395 | 0.004 | 1869.593 | 25.121 | 24.931 | 52.833 | 51.183 | 2.928 |
| c1_power2 | pair_396 | 0.004 | 1869.593 | 10.639 | 8.577 | 19.600 | 15.232 | 9.628 |
| c1_power2 | pair_397 | 0.004 | 1869.593 | 6.851 | 6.801 | 13.832 | 13.464 | 1.347 |
| c1_power2 | pair_398 | 0.004 | 1869.593 | 3.491 | 3.438 | 7.000 | 7.163 | 0.602 |
| c1_power2 | pair_399 | 0.004 | 1869.593 | 6.907 | 6.782 | 13.983 | 14.307 | 1.135 |
| c1_power2 | pair_400 | 0.004 | 1869.593 | 6.639 | 6.650 | 13.464 | 12.728 | 1.377 |
| c1_power2 | pair_401 | 0.004 | 1869.593 | 6.474 | 6.630 | 13.487 | 13.582 | 0.764 |
| c1_power2 | pair_402 | 0.004 | 1869.593 | 6.307 | 6.672 | 13.103 | 12.837 | 0.705 |
| c1_power2 | pair_403 | 0.004 | 1869.593 | 6.695 | 7.189 | 14.039 | 13.696 | 1.188 |
| c1_power2 | pair_404 | 0.004 | 1869.593 | 6.643 | 6.518 | 13.360 | 12.804 | 1.133 |
| c1_power2 | pair_405 | 0.004 | 1869.593 | 7.420 | 6.918 | 14.592 | 12.884 | 3.184 |
| c1_power2 | pair_406 | 0.004 | 1869.593 | 3.467 | 3.370 | 6.926 | 6.780 | 0.674 |
| c1_power2 | pair_407 | 0.004 | 1869.593 | 7.195 | 6.877 | 14.192 | 12.639 | 2.818 |
| c1_power2 | pair_408 | 0.004 | 1869.593 | 6.807 | 6.935 | 14.069 | 12.776 | 2.287 |
| c1_power2 | pair_409 | 0.004 | 1869.593 | 12.732 | 12.676 | 25.801 | 25.392 | 0.795 |
| c1_power2 | pair_410 | 0.004 | 1869.593 | 25.206 | 26.802 | 53.506 | 51.811 | 3.396 |
| c1_power2 | pair_411 | 0.004 | 1869.593 | 14.819 | 7.326 | 22.504 | 16.769 | 12.182 |
| c1_power2 | pair_412 | 0.004 | 1869.593 | 6.409 | 6.469 | 13.106 | 13.043 | 0.195 |
| c1_power2 | pair_413 | 0.004 | 1869.593 | 3.399 | 3.268 | 6.753 | 6.396 | 0.666 |
| c1_power2 | pair_414 | 0.004 | 1869.593 | 6.628 | 6.782 | 13.570 | 13.562 | 0.703 |
| c1_power2 | pair_415 | 0.004 | 1869.593 | 6.326 | 6.783 | 13.345 | 13.423 | 0.428 |
| c1_power2 | pair_416 | 0.004 | 1869.593 | 6.396 | 6.302 | 12.900 | 12.964 | 0.167 |
| c1_power2 | pair_417 | 0.004 | 1869.593 | 6.296 | 6.272 | 12.745 | 12.760 | 0.131 |
| c1_power2 | pair_418 | 0.004 | 1869.593 | 6.293 | 6.327 | 12.868 | 12.917 | 0.133 |
| c1_power2 | pair_419 | 0.004 | 1869.593 | 3.239 | 3.179 | 6.498 | 6.509 | 0.064 |
| c1_power2 | pair_420 | 0.004 | 1869.593 | 6.370 | 6.522 | 14.848 | 13.302 | 3.205 |
| c1_power2 | pair_421 | 0.004 | 1869.593 | 12.639 | 12.548 | 26.464 | 26.211 | 1.258 |
| c1_power2 | pair_422 | 0.004 | 1869.593 | 6.304 | 6.202 | 12.693 | 12.671 | 0.155 |
| c1_power2 | pair_423 | 0.004 | 1869.593 | 6.319 | 6.273 | 12.840 | 12.763 | 0.134 |
| c1_power2 | pair_424 | 0.004 | 1869.593 | 3.118 | 3.140 | 6.313 | 6.293 | 0.053 |
| c1_power2 | pair_425 | 0.004 | 1869.593 | 6.261 | 6.315 | 12.732 | 12.754 | 0.129 |
| c1_power2 | pair_426 | 0.004 | 1869.593 | 6.324 | 6.305 | 12.862 | 12.870 | 0.059 |
| c1_power2 | pair_427 | 0.004 | 1869.593 | 6.287 | 6.279 | 12.754 | 12.781 | 0.063 |
| c1_power2 | pair_428 | 0.004 | 1869.593 | 3.141 | 3.130 | 6.361 | 6.342 | 0.049 |
| c1_power2 | pair_429 | 0.004 | 1869.593 | 6.268 | 6.299 | 12.714 | 12.711 | 0.062 |
| c1_power2 | pair_430 | 0.004 | 1869.593 | 6.299 | 6.241 | 12.713 | 12.742 | 0.074 |
| c1_power2 | pair_431 | 0.004 | 1869.593 | 3.162 | 3.145 | 6.391 | 6.360 | 0.057 |
| c1_power2 | pair_432 | 0.004 | 1869.593 | 6.313 | 6.203 | 12.693 | 12.646 | 0.131 |
| c1_power2 | pair_433 | 0.004 | 1869.593 | 3.122 | 3.139 | 6.349 | 6.342 | 0.067 |
| c1_power2 | pair_434 | 0.004 | 1869.593 | 6.287 | 6.302 | 12.808 | 12.794 | 0.125 |
| c1_power2 | pair_435 | 0.004 | 1869.593 | 3.168 | 3.107 | 6.338 | 6.306 | 0.113 |
| c1_sqrt | pair_001 | 0.004 | 1877.073 | 8.873 | 8.702 | 17.795 | 17.797 | 0.026 |
| c1_sqrt | pair_002 | 0.004 | 1877.073 | 7.132 | 7.103 | 14.443 | 14.461 | 0.097 |
| c1_sqrt | pair_003 | 0.004 | 1877.073 | 7.854 | 7.889 | 16.003 | 16.056 | 0.121 |
| c1_sqrt | pair_004 | 0.004 | 1877.073 | 8.796 | 8.660 | 17.689 | 17.655 | 0.076 |
| c1_sqrt | pair_005 | 0.004 | 1877.073 | 4.068 | 7.824 | 12.061 | 8.157 | 6.832 |
| c1_sqrt | pair_006 | 0.004 | 1877.073 | 7.605 | 7.684 | 15.479 | 14.655 | 1.652 |
| c1_sqrt | pair_007 | 0.004 | 1877.073 | 16.527 | 16.848 | 33.937 | 32.388 | 3.047 |
| c1_sqrt | pair_008 | 0.004 | 1877.073 | 11.133 | 8.220 | 19.717 | 17.469 | 5.151 |
| c1_sqrt | pair_009 | 0.004 | 1877.073 | 21.953 | 20.593 | 43.354 | 43.893 | 1.398 |
| c1_sqrt | pair_010 | 0.004 | 1877.073 | 8.010 | 7.848 | 16.141 | 16.099 | 0.121 |
| c1_sqrt | pair_011 | 0.004 | 1877.073 | 6.352 | 6.261 | 12.795 | 12.765 | 0.109 |
| c1_sqrt | pair_012 | 0.004 | 1877.073 | 6.396 | 6.346 | 12.960 | 12.955 | 0.048 |
| c1_sqrt | pair_013 | 0.004 | 1877.073 | 7.148 | 7.151 | 14.566 | 14.616 | 0.120 |
| c1_sqrt | pair_014 | 0.004 | 1877.073 | 7.239 | 7.080 | 14.524 | 14.411 | 0.281 |
| c1_sqrt | pair_015 | 0.004 | 1877.073 | 7.127 | 7.043 | 14.909 | 14.387 | 0.944 |
| c1_sqrt | pair_016 | 0.004 | 1877.073 | 6.425 | 4.151 | 10.733 | 9.185 | 3.685 |
| c1_sqrt | pair_017 | 0.004 | 1877.073 | 4.244 | 4.282 | 8.638 | 8.026 | 1.086 |
| c1_sqrt | pair_018 | 0.004 | 1877.073 | 12.536 | 12.666 | 25.622 | 23.970 | 2.878 |
| c1_sqrt | pair_019 | 0.004 | 1877.073 | 8.091 | 8.324 | 16.638 | 15.932 | 1.295 |
| c1_sqrt | pair_020 | 0.004 | 1877.073 | 22.437 | 26.068 | 49.429 | 41.615 | 13.583 |
| c1_sqrt | pair_021 | 0.004 | 1877.073 | 9.039 | 8.511 | 17.804 | 16.861 | 2.528 |
| c1_sqrt | pair_022 | 0.004 | 1877.073 | 22.501 | 23.426 | 46.904 | 45.691 | 7.710 |
| c1_sqrt | pair_023 | 0.004 | 1877.073 | 6.811 | 7.097 | 14.191 | 14.385 | 1.350 |
| c1_sqrt | pair_024 | 0.004 | 1877.073 | 13.516 | 12.962 | 27.016 | 27.244 | 1.560 |
| c1_sqrt | pair_025 | 0.004 | 1877.073 | 21.360 | 20.820 | 45.732 | 43.806 | 3.871 |
| c1_sqrt | pair_026 | 0.004 | 1877.073 | 6.345 | 6.677 | 13.302 | 13.015 | 0.603 |
| c1_sqrt | pair_027 | 0.004 | 1877.073 | 11.602 | 7.606 | 19.464 | 13.005 | 11.312 |
| c1_sqrt | pair_028 | 0.004 | 1877.073 | 4.276 | 4.274 | 8.799 | 8.151 | 1.428 |
| c1_sqrt | pair_029 | 0.004 | 1877.073 | 7.501 | 7.669 | 15.482 | 14.653 | 1.650 |
| c1_sqrt | pair_030 | 0.004 | 1877.073 | 7.225 | 9.549 | 16.987 | 14.449 | 4.469 |
| c1_sqrt | pair_031 | 0.004 | 1877.073 | 7.973 | 8.049 | 16.301 | 16.219 | 0.510 |
| c1_sqrt | pair_032 | 0.004 | 1877.073 | 9.241 | 8.969 | 18.613 | 18.336 | 1.413 |
| c1_sqrt | pair_033 | 0.004 | 1877.073 | 3.936 | 3.995 | 8.098 | 8.190 | 0.207 |
| c1_sqrt | pair_034 | 0.004 | 1877.073 | 7.228 | 7.306 | 14.823 | 14.432 | 0.880 |
| c1_sqrt | pair_035 | 0.004 | 1877.073 | 8.718 | 8.597 | 17.634 | 17.635 | 0.147 |
| c1_sqrt | pair_036 | 0.004 | 1877.073 | 8.591 | 8.553 | 17.364 | 15.998 | 2.491 |
| c1_sqrt | pair_037 | 0.004 | 1877.073 | 8.636 | 8.656 | 17.581 | 17.656 | 0.252 |
| c1_sqrt | pair_038 | 0.004 | 1877.073 | 7.949 | 7.803 | 15.984 | 16.036 | 0.152 |
| c1_sqrt | pair_039 | 0.004 | 1877.073 | 6.356 | 6.343 | 12.992 | 13.035 | 0.126 |
| c1_sqrt | pair_040 | 0.004 | 1877.073 | 6.764 | 6.465 | 13.420 | 13.048 | 0.869 |
| c1_sqrt | pair_041 | 0.004 | 1877.073 | 7.800 | 7.068 | 15.083 | 14.380 | 1.274 |
| c1_sqrt | pair_042 | 0.004 | 1877.073 | 7.088 | 7.169 | 14.494 | 14.417 | 0.178 |
| c1_sqrt | pair_043 | 0.004 | 1877.073 | 7.123 | 7.030 | 14.379 | 14.353 | 0.079 |
| c1_sqrt | pair_044 | 0.004 | 1877.073 | 3.936 | 3.928 | 8.003 | 8.006 | 0.083 |
| c1_sqrt | pair_045 | 0.004 | 1877.073 | 3.945 | 3.905 | 7.958 | 7.988 | 0.140 |
| c1_sqrt | pair_046 | 0.004 | 1877.073 | 8.688 | 8.668 | 17.642 | 17.676 | 0.119 |
| c1_sqrt | pair_047 | 0.004 | 1877.073 | 7.905 | 8.019 | 16.211 | 16.118 | 0.265 |
| c1_sqrt | pair_048 | 0.004 | 1877.073 | 16.772 | 9.272 | 26.319 | 18.334 | 14.443 |
| c1_sqrt | pair_049 | 0.004 | 1877.073 | 8.276 | 8.300 | 17.372 | 17.508 | 1.233 |
| c1_sqrt | pair_050 | 0.004 | 1877.073 | 8.766 | 8.646 | 17.671 | 17.613 | 0.169 |
| c1_sqrt | pair_051 | 0.004 | 1877.073 | 6.268 | 6.311 | 12.796 | 12.811 | 0.167 |
| c1_sqrt | pair_052 | 0.004 | 1877.073 | 8.753 | 8.544 | 17.626 | 17.653 | 0.119 |
| c1_sqrt | pair_053 | 0.004 | 1877.073 | 8.689 | 8.583 | 17.562 | 17.541 | 0.045 |
| c1_sqrt | pair_054 | 0.004 | 1877.073 | 6.304 | 6.285 | 12.784 | 12.775 | 0.016 |
| c1_sqrt | pair_055 | 0.004 | 1877.073 | 6.326 | 6.234 | 12.786 | 12.778 | 0.143 |
| c1_sqrt | pair_056 | 0.004 | 1877.073 | 7.802 | 3.909 | 11.848 | 7.932 | 6.795 |
| c1_sqrt | pair_057 | 0.004 | 1877.073 | 7.134 | 7.020 | 14.376 | 14.398 | 0.154 |
| c1_sqrt | pair_058 | 0.004 | 1877.073 | 7.094 | 7.052 | 14.383 | 14.340 | 0.096 |
| c1_sqrt | pair_059 | 0.004 | 1877.073 | 7.043 | 7.077 | 14.356 | 14.348 | 0.211 |
| c1_sqrt | pair_060 | 0.004 | 1877.073 | 3.946 | 3.918 | 7.959 | 7.963 | 0.007 |
| c1_sqrt | pair_061 | 0.004 | 1877.073 | 7.170 | 6.994 | 14.377 | 14.356 | 0.216 |
| c1_sqrt | pair_062 | 0.004 | 1877.073 | 7.042 | 7.038 | 14.339 | 14.399 | 0.118 |
| c1_sqrt | pair_063 | 0.004 | 1877.073 | 7.033 | 7.033 | 14.301 | 14.297 | 0.157 |
| c1_sqrt | pair_064 | 0.004 | 1877.073 | 7.202 | 7.019 | 14.443 | 14.267 | 0.361 |
| c1_sqrt | pair_065 | 0.004 | 1877.073 | 7.101 | 6.984 | 20.300 | 14.271 | 10.465 |
| c1_sqrt | pair_066 | 0.004 | 1877.073 | 7.175 | 6.283 | 13.707 | 12.730 | 1.726 |
| c1_sqrt | pair_067 | 0.004 | 1877.073 | 6.858 | 6.275 | 13.353 | 13.179 | 0.683 |
| c1_sqrt | pair_068 | 0.004 | 1877.073 | 7.138 | 7.006 | 14.411 | 14.431 | 0.166 |
| c1_sqrt | pair_069 | 0.004 | 1877.073 | 7.084 | 7.015 | 14.295 | 14.349 | 0.127 |
| c1_sqrt | pair_070 | 0.004 | 1877.073 | 7.077 | 7.103 | 14.442 | 14.467 | 0.071 |
| c1_sqrt | pair_071 | 0.004 | 1877.073 | 3.934 | 3.891 | 7.944 | 7.925 | 0.049 |
| c1_sqrt | pair_072 | 0.004 | 1877.073 | 3.914 | 3.905 | 8.002 | 7.978 | 0.128 |
| c1_sqrt | pair_073 | 0.004 | 1877.073 | 7.084 | 7.010 | 14.292 | 14.316 | 0.103 |
| c1_sqrt | pair_074 | 0.004 | 1877.073 | 7.096 | 7.062 | 14.365 | 14.330 | 0.093 |
| c1_sqrt | pair_075 | 0.004 | 1877.073 | 7.098 | 7.039 | 14.350 | 14.390 | 0.136 |
| c1_sqrt | pair_076 | 0.004 | 1877.073 | 7.130 | 8.788 | 16.146 | 14.394 | 3.088 |
| c1_sqrt | pair_077 | 0.004 | 1877.073 | 7.852 | 7.019 | 15.105 | 15.468 | 0.760 |
| c1_sqrt | pair_078 | 0.004 | 1877.073 | 6.241 | 6.269 | 12.765 | 12.678 | 0.162 |
| c1_sqrt | pair_079 | 0.004 | 1877.073 | 7.031 | 7.020 | 14.315 | 14.311 | 0.020 |
| c1_sqrt | pair_080 | 0.004 | 1877.073 | 7.018 | 7.018 | 14.263 | 14.235 | 0.055 |
| c1_sqrt | pair_081 | 0.004 | 1877.073 | 6.260 | 6.211 | 95.838 | 12.738 | 144.017 |
| c1_sqrt | pair_082 | 0.004 | 1877.073 | 9.563 | 9.661 | 19.466 | 16.200 | 8.765 |
| c1_sqrt | pair_083 | 0.004 | 1877.073 | 5.396 | 5.139 | 10.696 | 10.233 | 3.006 |
| c1_sqrt | pair_084 | 0.004 | 1877.073 | 8.158 | 8.999 | 17.500 | 17.561 | 3.108 |
| c1_sqrt | pair_085 | 0.004 | 1877.073 | 8.363 | 8.047 | 16.714 | 16.603 | 0.922 |
| c1_sqrt | pair_086 | 0.004 | 1877.073 | 4.453 | 4.085 | 8.647 | 7.981 | 1.214 |
| c1_sqrt | pair_087 | 0.004 | 1877.073 | 7.365 | 7.210 | 14.904 | 14.302 | 1.087 |
| c1_sqrt | pair_088 | 0.004 | 1877.073 | 8.010 | 8.238 | 16.510 | 15.928 | 1.154 |
| c1_sqrt | pair_089 | 0.004 | 1877.073 | 8.340 | 7.913 | 16.800 | 16.410 | 1.181 |
| c1_sqrt | pair_090 | 0.004 | 1877.073 | 8.171 | 8.004 | 16.566 | 16.030 | 1.032 |
| c1_sqrt | pair_091 | 0.004 | 1877.073 | 7.938 | 8.008 | 16.171 | 16.146 | 0.331 |
| c1_sqrt | pair_092 | 0.004 | 1877.073 | 6.818 | 6.787 | 14.002 | 13.247 | 1.736 |
| c1_sqrt | pair_093 | 0.004 | 1877.073 | 6.414 | 6.772 | 17.141 | 14.305 | 6.191 |
| c1_sqrt | pair_094 | 0.004 | 1877.073 | 8.383 | 7.169 | 15.798 | 15.757 | 1.471 |
| c1_sqrt | pair_095 | 0.004 | 1877.073 | 7.412 | 7.324 | 14.962 | 14.516 | 0.863 |
| c1_sqrt | pair_096 | 0.004 | 1877.073 | 7.208 | 7.273 | 14.717 | 14.665 | 0.377 |
| c1_sqrt | pair_097 | 0.004 | 1877.073 | 4.027 | 4.163 | 8.351 | 8.133 | 0.382 |
| c1_sqrt | pair_098 | 0.004 | 1877.073 | 4.119 | 3.958 | 8.235 | 8.050 | 0.413 |
| c1_sqrt | pair_099 | 0.004 | 1877.073 | 8.521 | 8.284 | 17.121 | 17.576 | 0.834 |
| c1_sqrt | pair_100 | 0.004 | 1877.073 | 10.988 | 10.332 | 21.518 | 16.891 | 8.583 |
| c1_sqrt | pair_101 | 0.004 | 1877.073 | 7.915 | 7.987 | 16.189 | 15.928 | 0.582 |
| c1_sqrt | pair_102 | 0.004 | 1877.073 | 7.843 | 7.759 | 15.856 | 15.865 | 0.034 |
| c1_sqrt | pair_103 | 0.004 | 1877.073 | 7.862 | 7.837 | 15.931 | 15.981 | 0.113 |
| c1_sqrt | pair_104 | 0.004 | 1877.073 | 6.247 | 6.223 | 12.635 | 12.618 | 0.036 |
| c1_sqrt | pair_105 | 0.004 | 1877.073 | 7.859 | 7.826 | 15.900 | 15.906 | 0.040 |
| c1_sqrt | pair_106 | 0.004 | 1877.073 | 7.839 | 7.830 | 15.898 | 15.881 | 0.140 |
| c1_sqrt | pair_107 | 0.004 | 1877.073 | 6.442 | 9.490 | 16.143 | 13.190 | 5.551 |
| c1_sqrt | pair_108 | 0.004 | 1877.073 | 6.263 | 6.259 | 12.752 | 12.743 | 0.095 |
| c1_sqrt | pair_109 | 0.004 | 1877.073 | 3.924 | 3.920 | 7.932 | 7.906 | 0.083 |
| c1_sqrt | pair_110 | 0.004 | 1877.073 | 7.146 | 7.715 | 15.142 | 14.393 | 1.373 |
| c1_sqrt | pair_111 | 0.004 | 1877.073 | 5.133 | 4.065 | 9.315 | 8.235 | 2.213 |
| c1_sqrt | pair_112 | 0.004 | 1877.073 | 7.046 | 7.041 | 14.500 | 14.301 | 0.354 |
| c1_sqrt | pair_113 | 0.004 | 1877.073 | 15.528 | 8.964 | 24.725 | 18.301 | 11.756 |
| c1_sqrt | pair_114 | 0.004 | 1877.073 | 7.865 | 7.795 | 15.865 | 15.841 | 0.084 |
| c1_sqrt | pair_115 | 0.004 | 1877.073 | 8.661 | 8.596 | 17.471 | 17.471 | 0.149 |
| c1_sqrt | pair_116 | 0.004 | 1877.073 | 7.846 | 7.785 | 15.887 | 15.873 | 0.080 |
| c1_sqrt | pair_117 | 0.004 | 1877.073 | 6.332 | 6.250 | 12.799 | 12.878 | 0.151 |
| c1_sqrt | pair_118 | 0.004 | 1877.073 | 6.330 | 6.321 | 12.859 | 12.859 | 0.142 |
| c1_sqrt | pair_119 | 0.004 | 1877.073 | 7.018 | 7.048 | 14.335 | 14.311 | 0.055 |
| c1_sqrt | pair_120 | 0.004 | 1877.073 | 7.069 | 6.986 | 14.278 | 14.285 | 0.052 |
| c1_sqrt | pair_121 | 0.004 | 1877.073 | 7.203 | 7.057 | 14.489 | 14.298 | 0.391 |
| c1_sqrt | pair_122 | 0.004 | 1877.073 | 3.950 | 3.921 | 8.025 | 8.022 | 0.017 |
| c1_sqrt | pair_123 | 0.004 | 1877.073 | 3.956 | 3.914 | 7.997 | 7.998 | 0.010 |
| c1_sqrt | pair_124 | 0.004 | 1877.073 | 8.658 | 9.157 | 18.596 | 17.701 | 1.662 |
| c1_sqrt | pair_125 | 0.004 | 1877.073 | 7.942 | 7.832 | 16.051 | 15.991 | 0.140 |
| c1_sqrt | pair_126 | 0.004 | 1877.073 | 8.756 | 8.577 | 17.553 | 17.545 | 0.085 |
| c1_sqrt | pair_127 | 0.004 | 1877.073 | 7.831 | 7.851 | 15.901 | 15.866 | 0.193 |
| c1_sqrt | pair_128 | 0.004 | 1877.073 | 11.364 | 9.060 | 20.693 | 17.586 | 5.470 |
| c1_sqrt | pair_129 | 0.004 | 1877.073 | 6.331 | 6.182 | 12.693 | 12.685 | 0.020 |
| c1_sqrt | pair_130 | 0.004 | 1877.073 | 8.628 | 8.836 | 17.736 | 17.705 | 0.356 |
| c1_sqrt | pair_131 | 0.004 | 1877.073 | 8.691 | 8.654 | 17.604 | 17.436 | 0.309 |
| c1_sqrt | pair_132 | 0.004 | 1877.073 | 6.259 | 6.196 | 12.970 | 12.799 | 0.444 |
| c1_sqrt | pair_133 | 0.004 | 1877.073 | 9.874 | 6.578 | 16.667 | 12.701 | 6.929 |
| c1_sqrt | pair_134 | 0.004 | 1877.073 | 4.120 | 4.087 | 8.361 | 7.943 | 0.733 |
| c1_sqrt | pair_135 | 0.004 | 1877.073 | 7.206 | 7.411 | 14.950 | 14.339 | 1.060 |
| c1_sqrt | pair_136 | 0.004 | 1877.073 | 3.890 | 3.886 | 7.927 | 7.945 | 0.039 |
| c1_sqrt | pair_137 | 0.004 | 1877.073 | 3.951 | 4.326 | 8.397 | 7.975 | 0.821 |
| c1_sqrt | pair_138 | 0.004 | 1877.073 | 3.955 | 4.325 | 8.904 | 9.114 | 0.931 |
| c1_sqrt | pair_139 | 0.004 | 1877.073 | 3.913 | 4.203 | 8.305 | 7.974 | 0.602 |
| c1_sqrt | pair_140 | 0.004 | 1877.073 | 4.114 | 4.135 | 8.373 | 7.908 | 0.916 |
| c1_sqrt | pair_141 | 0.004 | 1877.073 | 3.949 | 3.901 | 8.004 | 7.949 | 0.107 |
| c1_sqrt | pair_142 | 0.004 | 1877.073 | 3.905 | 3.873 | 7.921 | 7.932 | 0.018 |
| c1_sqrt | pair_143 | 0.004 | 1877.073 | 3.935 | 3.891 | 7.945 | 7.998 | 0.104 |
| c1_sqrt | pair_144 | 0.004 | 1877.073 | 3.934 | 3.907 | 7.969 | 7.960 | 0.033 |
| c1_sqrt | pair_145 | 0.004 | 1877.073 | 3.954 | 3.883 | 7.981 | 7.975 | 0.015 |
| c1_sqrt | pair_146 | 0.004 | 1877.073 | 3.913 | 3.889 | 7.894 | 7.872 | 0.047 |
| c1_sqrt | pair_147 | 0.004 | 1877.073 | 3.977 | 3.877 | 7.987 | 7.990 | 0.015 |
| c1_sqrt | pair_148 | 0.004 | 1877.073 | 3.918 | 3.878 | 7.919 | 7.897 | 0.038 |
| c1_sqrt | pair_149 | 0.004 | 1877.073 | 3.965 | 3.953 | 8.053 | 8.004 | 0.146 |
| c1_sqrt | pair_150 | 0.004 | 1877.073 | 3.934 | 3.889 | 7.970 | 7.972 | 0.045 |
| c1_sqrt | pair_151 | 0.004 | 1877.073 | 3.955 | 3.852 | 7.953 | 7.951 | 0.008 |
| c1_sqrt | pair_152 | 0.004 | 1877.073 | 3.972 | 3.911 | 8.013 | 8.033 | 0.035 |
| c1_sqrt | pair_153 | 0.004 | 1877.073 | 3.917 | 3.847 | 7.894 | 7.921 | 0.050 |
| c1_sqrt | pair_154 | 0.004 | 1877.073 | 4.023 | 3.886 | 8.042 | 8.005 | 0.104 |
| c1_sqrt | pair_155 | 0.004 | 1877.073 | 3.940 | 3.885 | 7.970 | 7.983 | 0.032 |
| c1_sqrt | pair_156 | 0.004 | 1877.073 | 3.957 | 3.894 | 7.989 | 7.967 | 0.055 |
| c1_sqrt | pair_157 | 0.004 | 1877.073 | 3.956 | 3.922 | 7.995 | 7.990 | 0.089 |
| c1_sqrt | pair_158 | 0.004 | 1877.073 | 3.927 | 3.875 | 7.939 | 7.928 | 0.042 |
| c1_sqrt | pair_159 | 0.004 | 1877.073 | 3.944 | 3.885 | 7.915 | 7.896 | 0.038 |
| c1_sqrt | pair_160 | 0.004 | 1877.073 | 8.185 | 8.081 | 16.457 | 14.677 | 3.345 |
| c1_sqrt | pair_161 | 0.004 | 1877.073 | 7.056 | 6.995 | 14.227 | 14.202 | 0.116 |
| c1_sqrt | pair_162 | 0.004 | 1877.073 | 7.048 | 7.080 | 14.389 | 14.378 | 0.097 |
| c1_sqrt | pair_163 | 0.004 | 1877.073 | 7.075 | 7.002 | 14.305 | 14.309 | 0.111 |
| c1_sqrt | pair_164 | 0.004 | 1877.073 | 6.241 | 6.176 | 16.350 | 12.769 | 6.378 |
| c1_sqrt | pair_165 | 0.004 | 1877.073 | 6.256 | 6.253 | 12.698 | 12.678 | 0.125 |
| c1_sqrt | pair_166 | 0.004 | 1877.073 | 7.069 | 7.000 | 14.279 | 14.345 | 0.126 |
| c1_sqrt | pair_167 | 0.004 | 1877.073 | 10.978 | 7.370 | 18.614 | 14.437 | 7.382 |
| c1_sqrt | pair_168 | 0.004 | 1877.073 | 7.187 | 7.502 | 14.930 | 14.358 | 1.075 |
| c1_sqrt | pair_169 | 0.004 | 1877.073 | 4.099 | 4.153 | 8.491 | 7.994 | 0.890 |
| c1_sqrt | pair_170 | 0.004 | 1877.073 | 4.037 | 4.301 | 8.480 | 8.314 | 0.624 |
| c1_sqrt | pair_171 | 0.004 | 1877.073 | 7.325 | 6.999 | 14.546 | 14.331 | 0.480 |
| c1_sqrt | pair_172 | 0.004 | 1877.073 | 7.026 | 7.023 | 14.299 | 14.292 | 0.063 |
| c1_sqrt | pair_173 | 0.004 | 1877.073 | 7.093 | 7.092 | 14.820 | 14.854 | 0.565 |
| c1_sqrt | pair_174 | 0.004 | 1877.073 | 7.499 | 7.479 | 15.235 | 14.510 | 1.441 |
| c1_sqrt | pair_175 | 0.004 | 1877.073 | 7.671 | 7.306 | 15.202 | 14.318 | 1.552 |
| c1_sqrt | pair_176 | 0.004 | 1877.073 | 6.907 | 6.768 | 13.991 | 12.983 | 1.893 |
| c1_sqrt | pair_177 | 0.004 | 1877.073 | 7.670 | 7.203 | 15.053 | 14.515 | 1.169 |
| c1_sqrt | pair_178 | 0.004 | 1877.073 | 7.220 | 7.213 | 14.777 | 14.393 | 0.668 |
| c1_sqrt | pair_179 | 0.004 | 1877.073 | 6.594 | 6.463 | 13.281 | 12.724 | 0.981 |
| c1_sqrt | pair_180 | 0.004 | 1877.073 | 6.265 | 6.221 | 12.709 | 12.690 | 0.034 |
| c1_sqrt | pair_181 | 0.004 | 1877.073 | 3.919 | 3.901 | 7.992 | 7.992 | 0.005 |
| c1_sqrt | pair_182 | 0.004 | 1877.073 | 7.037 | 7.051 | 14.309 | 14.220 | 0.164 |
| c1_sqrt | pair_183 | 0.004 | 1877.073 | 7.872 | 7.868 | 16.040 | 16.017 | 0.192 |
| c1_sqrt | pair_184 | 0.004 | 1877.073 | 15.628 | 15.634 | 31.823 | 31.681 | 0.279 |
| c1_sqrt | pair_185 | 0.004 | 1877.073 | 7.832 | 7.800 | 15.940 | 15.942 | 0.006 |
| c1_sqrt | pair_186 | 0.004 | 1877.073 | 6.229 | 6.226 | 12.674 | 12.619 | 0.146 |
| c1_sqrt | pair_187 | 0.004 | 1877.073 | 6.228 | 6.289 | 12.696 | 12.627 | 0.119 |
| c1_sqrt | pair_188 | 0.004 | 1877.073 | 7.023 | 7.018 | 14.252 | 14.287 | 0.083 |
| c1_sqrt | pair_189 | 0.004 | 1877.073 | 7.066 | 7.017 | 14.300 | 14.275 | 0.057 |
| c1_sqrt | pair_190 | 0.004 | 1877.073 | 7.101 | 7.113 | 14.480 | 14.373 | 0.213 |
| c1_sqrt | pair_191 | 0.004 | 1877.073 | 5.022 | 6.495 | 11.784 | 8.006 | 6.581 |
| c1_sqrt | pair_192 | 0.004 | 1877.073 | 5.521 | 5.380 | 11.067 | 8.977 | 4.555 |
| c1_sqrt | pair_193 | 0.004 | 1877.073 | 18.938 | 12.077 | 31.386 | 32.379 | 6.654 |
| c1_sqrt | pair_194 | 0.004 | 1877.073 | 8.112 | 7.882 | 16.269 | 15.985 | 0.556 |
| c1_sqrt | pair_195 | 0.004 | 1877.073 | 16.348 | 16.062 | 36.678 | 35.582 | 5.440 |
| c1_sqrt | pair_196 | 0.004 | 1877.073 | 8.110 | 7.774 | 16.090 | 15.925 | 0.484 |
| c1_sqrt | pair_197 | 0.004 | 1877.073 | 15.655 | 15.585 | 31.804 | 31.767 | 0.190 |
| c1_sqrt | pair_198 | 0.004 | 1877.073 | 6.357 | 6.247 | 12.833 | 12.829 | 0.154 |
| c1_sqrt | pair_199 | 0.004 | 1877.073 | 12.535 | 12.463 | 25.437 | 25.435 | 0.057 |
| c1_sqrt | pair_200 | 0.004 | 1877.073 | 15.597 | 15.605 | 31.716 | 31.757 | 0.080 |
| c1_sqrt | pair_201 | 0.004 | 1877.073 | 6.300 | 6.226 | 12.710 | 12.676 | 0.222 |
| c1_sqrt | pair_202 | 0.004 | 1877.073 | 6.312 | 6.232 | 12.755 | 12.757 | 0.137 |
| c1_sqrt | pair_203 | 0.004 | 1877.073 | 4.026 | 3.912 | 8.053 | 7.965 | 0.161 |
| c1_sqrt | pair_204 | 0.004 | 1877.073 | 7.053 | 7.006 | 14.305 | 14.286 | 0.058 |
| c1_sqrt | pair_205 | 0.004 | 1877.073 | 7.934 | 8.927 | 17.118 | 15.974 | 2.043 |
| c1_sqrt | pair_206 | 0.004 | 1877.073 | 7.843 | 9.049 | 17.164 | 16.061 | 2.015 |
| c1_sqrt | pair_207 | 0.004 | 1877.073 | 6.528 | 6.378 | 13.098 | 13.079 | 0.332 |
| c1_sqrt | pair_208 | 0.004 | 1877.073 | 6.431 | 6.344 | 12.996 | 12.935 | 0.226 |
| c1_sqrt | pair_209 | 0.004 | 1877.073 | 7.150 | 7.000 | 14.359 | 14.334 | 0.119 |
| c1_sqrt | pair_210 | 0.004 | 1877.073 | 7.054 | 7.021 | 14.305 | 14.263 | 0.092 |
| c1_sqrt | pair_211 | 0.004 | 1877.073 | 7.084 | 7.060 | 14.339 | 14.319 | 0.148 |
| c1_sqrt | pair_212 | 0.004 | 1877.073 | 3.963 | 3.873 | 7.969 | 7.963 | 0.081 |
| c1_sqrt | pair_213 | 0.004 | 1877.073 | 3.951 | 3.873 | 7.938 | 7.918 | 0.055 |
| c1_sqrt | pair_214 | 0.004 | 1877.073 | 8.851 | 9.650 | 18.733 | 16.029 | 4.824 |
| c1_sqrt | pair_215 | 0.004 | 1877.073 | 7.917 | 7.813 | 15.970 | 15.955 | 0.107 |
| c1_sqrt | pair_216 | 0.004 | 1877.073 | 7.826 | 7.783 | 15.848 | 15.821 | 0.119 |
| c1_sqrt | pair_217 | 0.004 | 1877.073 | 7.845 | 7.789 | 17.407 | 16.028 | 2.453 |
| c1_sqrt | pair_218 | 0.004 | 1877.073 | 8.249 | 8.668 | 17.296 | 16.047 | 2.232 |
| c1_sqrt | pair_219 | 0.004 | 1877.073 | 6.349 | 6.793 | 13.360 | 12.818 | 1.011 |
| c1_sqrt | pair_220 | 0.004 | 1877.073 | 8.087 | 7.866 | 16.183 | 16.045 | 0.386 |
| c1_sqrt | pair_221 | 0.004 | 1877.073 | 7.880 | 7.841 | 15.953 | 15.884 | 0.131 |
| c1_sqrt | pair_222 | 0.004 | 1877.073 | 6.319 | 6.297 | 12.839 | 12.810 | 0.145 |
| c1_sqrt | pair_223 | 0.004 | 1877.073 | 6.349 | 6.266 | 12.789 | 12.775 | 0.235 |
| c1_sqrt | pair_224 | 0.004 | 1877.073 | 4.097 | 3.932 | 10.149 | 8.540 | 3.311 |
| c1_sqrt | pair_225 | 0.004 | 1877.073 | 11.283 | 7.443 | 19.053 | 14.435 | 8.002 |
| c1_sqrt | pair_226 | 0.004 | 1877.073 | 8.047 | 8.072 | 16.331 | 15.990 | 0.605 |
| c1_sqrt | pair_227 | 0.004 | 1877.073 | 6.370 | 6.221 | 12.773 | 12.781 | 0.138 |
| c1_sqrt | pair_228 | 0.004 | 1877.073 | 6.259 | 6.310 | 12.773 | 12.737 | 0.151 |
| c1_sqrt | pair_229 | 0.004 | 1877.073 | 7.236 | 7.043 | 14.481 | 14.392 | 0.188 |
| c1_sqrt | pair_230 | 0.004 | 1877.073 | 7.111 | 7.056 | 14.368 | 14.384 | 0.108 |
| c1_sqrt | pair_231 | 0.004 | 1877.073 | 7.133 | 7.025 | 14.390 | 14.443 | 0.215 |
| c1_sqrt | pair_232 | 0.004 | 1877.073 | 3.946 | 3.916 | 7.972 | 7.953 | 0.094 |
| c1_sqrt | pair_233 | 0.004 | 1877.073 | 3.993 | 3.915 | 8.021 | 7.993 | 0.079 |
| c1_sqrt | pair_234 | 0.004 | 1877.073 | 11.806 | 11.759 | 23.911 | 23.943 | 0.119 |
| c1_sqrt | pair_235 | 0.004 | 1877.073 | 7.887 | 7.819 | 15.931 | 15.915 | 0.088 |
| c1_sqrt | pair_236 | 0.004 | 1877.073 | 20.467 | 27.761 | 49.209 | 50.293 | 7.433 |
| c1_sqrt | pair_237 | 0.004 | 1877.073 | 7.980 | 8.293 | 16.591 | 15.939 | 1.233 |
| c1_sqrt | pair_238 | 0.004 | 1877.073 | 20.394 | 19.682 | 43.208 | 41.758 | 4.146 |
| c1_sqrt | pair_239 | 0.004 | 1877.073 | 6.253 | 6.226 | 12.712 | 12.718 | 0.095 |
| c1_sqrt | pair_240 | 0.004 | 1877.073 | 12.601 | 12.455 | 25.429 | 25.457 | 0.177 |
| c1_sqrt | pair_241 | 0.004 | 1877.073 | 20.260 | 20.224 | 41.261 | 41.175 | 0.216 |
| c1_sqrt | pair_242 | 0.004 | 1877.073 | 6.267 | 6.186 | 12.718 | 12.704 | 0.086 |
| c1_sqrt | pair_243 | 0.004 | 1877.073 | 6.332 | 6.196 | 12.773 | 12.736 | 0.119 |
| c1_sqrt | pair_244 | 0.004 | 1877.073 | 3.891 | 3.904 | 7.905 | 7.890 | 0.076 |
| c1_sqrt | pair_245 | 0.004 | 1877.073 | 7.794 | 7.013 | 14.991 | 14.248 | 1.325 |
| c1_sqrt | pair_246 | 0.004 | 1877.073 | 6.309 | 6.379 | 12.930 | 12.871 | 0.147 |
| c1_sqrt | pair_247 | 0.004 | 1877.073 | 6.295 | 6.232 | 12.698 | 12.653 | 0.080 |
| c1_sqrt | pair_248 | 0.004 | 1877.073 | 7.055 | 7.024 | 14.303 | 14.274 | 0.083 |
| c1_sqrt | pair_249 | 0.004 | 1877.073 | 7.104 | 6.986 | 14.286 | 14.292 | 0.034 |
| c1_sqrt | pair_250 | 0.004 | 1877.073 | 7.085 | 11.710 | 19.076 | 14.289 | 8.391 |
| c1_sqrt | pair_251 | 0.004 | 1877.073 | 4.030 | 3.912 | 8.069 | 7.957 | 0.201 |
| c1_sqrt | pair_252 | 0.004 | 1877.073 | 3.964 | 3.918 | 7.995 | 7.981 | 0.043 |
| c1_sqrt | pair_253 | 0.004 | 1877.073 | 8.068 | 9.241 | 17.582 | 16.180 | 2.748 |
| c1_sqrt | pair_254 | 0.004 | 1877.073 | 9.249 | 8.824 | 18.331 | 16.117 | 4.074 |
| c1_sqrt | pair_255 | 0.004 | 1877.073 | 9.253 | 8.912 | 18.449 | 16.000 | 4.296 |
| c1_sqrt | pair_256 | 0.004 | 1877.073 | 8.045 | 7.798 | 16.105 | 15.978 | 0.465 |
| c1_sqrt | pair_257 | 0.004 | 1877.073 | 7.860 | 7.774 | 15.837 | 15.793 | 0.116 |
| c1_sqrt | pair_258 | 0.004 | 1877.073 | 6.222 | 6.240 | 12.614 | 12.657 | 0.155 |
| c1_sqrt | pair_259 | 0.004 | 1877.073 | 7.831 | 7.837 | 15.904 | 15.850 | 0.229 |
| c1_sqrt | pair_260 | 0.004 | 1877.073 | 7.839 | 7.895 | 15.906 | 15.870 | 0.146 |
| c1_sqrt | pair_261 | 0.004 | 1877.073 | 6.266 | 6.222 | 12.705 | 12.699 | 0.102 |
| c1_sqrt | pair_262 | 0.004 | 1877.073 | 6.261 | 6.237 | 12.761 | 12.773 | 0.070 |
| c1_sqrt | pair_263 | 0.004 | 1877.073 | 3.915 | 3.957 | 8.006 | 8.004 | 0.060 |
| c1_sqrt | pair_264 | 0.004 | 1877.073 | 7.037 | 7.187 | 14.390 | 14.272 | 0.324 |
| c1_sqrt | pair_265 | 0.004 | 1877.073 | 6.323 | 6.612 | 15.405 | 13.952 | 3.574 |
| c1_sqrt | pair_266 | 0.004 | 1877.073 | 11.521 | 8.337 | 20.145 | 12.974 | 12.636 |
| c1_sqrt | pair_267 | 0.004 | 1877.073 | 6.693 | 6.581 | 13.554 | 12.859 | 1.314 |
| c1_sqrt | pair_268 | 0.004 | 1877.073 | 9.642 | 10.415 | 20.254 | 13.992 | 11.760 |
| c1_sqrt | pair_269 | 0.004 | 1877.073 | 3.978 | 3.969 | 8.072 | 7.954 | 0.243 |
| c1_sqrt | pair_270 | 0.004 | 1877.073 | 4.019 | 3.955 | 8.226 | 8.048 | 0.358 |
| c1_sqrt | pair_271 | 0.004 | 1877.073 | 6.339 | 6.522 | 13.092 | 12.946 | 0.350 |
| c1_sqrt | pair_272 | 0.004 | 1877.073 | 6.415 | 6.325 | 13.241 | 12.751 | 0.889 |
| c1_sqrt | pair_273 | 0.004 | 1877.073 | 6.334 | 6.685 | 13.220 | 13.003 | 0.596 |
| c1_sqrt | pair_274 | 0.004 | 1877.073 | 6.295 | 6.212 | 12.727 | 12.766 | 0.120 |
| c1_sqrt | pair_275 | 0.004 | 1877.073 | 6.272 | 6.227 | 12.736 | 12.635 | 0.229 |
| c1_sqrt | pair_276 | 0.004 | 1877.073 | 6.366 | 6.272 | 12.811 | 12.778 | 0.102 |
| c1_sqrt | pair_277 | 0.004 | 1877.073 | 6.349 | 6.505 | 13.116 | 12.914 | 0.622 |
| c1_sqrt | pair_278 | 0.004 | 1877.073 | 6.275 | 6.282 | 12.756 | 12.755 | 0.153 |
| c1_sqrt | pair_279 | 0.004 | 1877.073 | 6.337 | 6.232 | 12.790 | 12.797 | 0.219 |
| c1_sqrt | pair_280 | 0.004 | 1877.073 | 6.294 | 6.228 | 12.734 | 12.632 | 0.210 |
| c1_sqrt | pair_281 | 0.004 | 1877.073 | 3.984 | 3.900 | 7.999 | 8.051 | 0.116 |
| c1_sqrt | pair_282 | 0.004 | 1877.073 | 6.255 | 6.259 | 12.702 | 12.757 | 0.148 |
| c1_sqrt | pair_283 | 0.004 | 1877.073 | 6.266 | 6.273 | 12.727 | 12.733 | 0.173 |
| c1_sqrt | pair_284 | 0.004 | 1877.073 | 6.293 | 8.618 | 15.077 | 12.747 | 4.042 |
| c1_sqrt | pair_285 | 0.004 | 1877.073 | 10.328 | 6.331 | 16.883 | 18.239 | 3.604 |
| c1_sqrt | pair_286 | 0.004 | 1877.073 | 3.928 | 3.925 | 7.990 | 7.979 | 0.052 |
| c1_sqrt | pair_287 | 0.004 | 1877.073 | 3.901 | 3.904 | 7.976 | 7.995 | 0.112 |
| c1_sqrt | pair_288 | 0.004 | 1877.073 | 6.344 | 6.300 | 12.876 | 12.791 | 0.199 |
| c1_sqrt | pair_289 | 0.004 | 1877.073 | 6.397 | 6.222 | 12.859 | 12.737 | 0.248 |
| c1_sqrt | pair_290 | 0.004 | 1877.073 | 6.409 | 6.258 | 12.875 | 12.962 | 0.287 |
| c1_sqrt | pair_291 | 0.004 | 1877.073 | 6.318 | 6.200 | 12.739 | 12.735 | 0.069 |
| c1_sqrt | pair_292 | 0.004 | 1877.073 | 6.289 | 6.203 | 12.684 | 12.675 | 0.147 |
| c1_sqrt | pair_293 | 0.004 | 1877.073 | 6.290 | 6.200 | 12.729 | 12.752 | 0.059 |
| c1_sqrt | pair_294 | 0.004 | 1877.073 | 6.299 | 6.232 | 12.716 | 12.757 | 0.077 |
| c1_sqrt | pair_295 | 0.004 | 1877.073 | 6.337 | 6.245 | 12.783 | 12.838 | 0.126 |
| c1_sqrt | pair_296 | 0.004 | 1877.073 | 6.976 | 6.868 | 14.038 | 12.788 | 2.221 |
| c1_sqrt | pair_297 | 0.004 | 1877.073 | 6.986 | 6.749 | 13.987 | 12.796 | 2.096 |
| c1_sqrt | pair_298 | 0.004 | 1877.073 | 4.183 | 3.940 | 8.249 | 7.910 | 0.616 |
| c1_sqrt | pair_299 | 0.004 | 1877.073 | 6.254 | 6.282 | 12.743 | 12.739 | 0.081 |
| c1_sqrt | pair_300 | 0.004 | 1877.073 | 7.333 | 7.360 | 15.188 | 15.027 | 1.123 |
| c1_sqrt | pair_301 | 0.004 | 1877.073 | 7.367 | 7.335 | 14.881 | 14.776 | 0.486 |
| c1_sqrt | pair_302 | 0.004 | 1877.073 | 4.137 | 4.053 | 8.786 | 9.082 | 0.670 |
| c1_sqrt | pair_303 | 0.004 | 1877.073 | 5.973 | 4.271 | 10.378 | 10.023 | 2.404 |
| c1_sqrt | pair_304 | 0.004 | 1877.073 | 7.777 | 7.510 | 15.466 | 14.465 | 2.046 |
| c1_sqrt | pair_305 | 0.004 | 1877.073 | 7.553 | 7.553 | 15.351 | 14.436 | 1.651 |
| c1_sqrt | pair_306 | 0.004 | 1877.073 | 7.140 | 7.114 | 14.532 | 14.247 | 0.502 |
| c1_sqrt | pair_307 | 0.004 | 1877.073 | 7.405 | 7.513 | 15.253 | 14.432 | 1.457 |
| c1_sqrt | pair_308 | 0.004 | 1877.073 | 7.722 | 9.703 | 17.646 | 16.083 | 4.384 |
| c1_sqrt | pair_309 | 0.004 | 1877.073 | 6.586 | 6.233 | 13.040 | 12.816 | 0.486 |
| c1_sqrt | pair_310 | 0.004 | 1877.073 | 7.087 | 7.079 | 14.424 | 14.365 | 0.143 |
| c1_sqrt | pair_311 | 0.004 | 1877.073 | 7.062 | 7.033 | 14.308 | 14.280 | 0.092 |
| c1_sqrt | pair_312 | 0.004 | 1877.073 | 6.225 | 6.249 | 12.662 | 12.665 | 0.083 |
| c1_sqrt | pair_313 | 0.004 | 1877.073 | 6.299 | 6.343 | 12.842 | 12.799 | 0.131 |
| c1_sqrt | pair_314 | 0.004 | 1877.073 | 3.976 | 3.901 | 7.991 | 8.012 | 0.040 |
| c1_sqrt | pair_315 | 0.004 | 1877.073 | 7.027 | 7.037 | 14.286 | 14.187 | 0.183 |
| c1_sqrt | pair_316 | 0.004 | 1877.073 | 7.024 | 7.001 | 14.221 | 14.205 | 0.138 |
| c1_sqrt | pair_317 | 0.004 | 1877.073 | 3.922 | 3.899 | 7.964 | 7.984 | 0.054 |
| c1_sqrt | pair_318 | 0.004 | 1877.073 | 3.950 | 3.904 | 7.980 | 7.980 | 0.110 |
| c1_sqrt | pair_319 | 0.004 | 1877.073 | 7.073 | 6.988 | 14.254 | 14.238 | 0.121 |
| c1_sqrt | pair_320 | 0.004 | 1877.073 | 7.100 | 6.999 | 14.284 | 14.339 | 0.137 |
| c1_sqrt | pair_321 | 0.004 | 1877.073 | 7.036 | 7.050 | 14.300 | 14.234 | 0.135 |
| c1_sqrt | pair_322 | 0.004 | 1877.073 | 7.034 | 7.022 | 14.278 | 14.254 | 0.053 |
| c1_sqrt | pair_323 | 0.004 | 1877.073 | 7.281 | 7.123 | 14.622 | 14.354 | 0.480 |
| c1_sqrt | pair_324 | 0.004 | 1877.073 | 6.345 | 6.324 | 12.865 | 12.954 | 0.209 |
| c1_sqrt | pair_325 | 0.004 | 1877.073 | 7.038 | 7.026 | 14.381 | 14.325 | 0.207 |
| c1_sqrt | pair_326 | 0.004 | 1877.073 | 7.960 | 7.032 | 15.256 | 14.602 | 1.397 |
| c1_sqrt | pair_327 | 0.004 | 1877.073 | 6.303 | 6.243 | 12.736 | 12.735 | 0.103 |
| c1_sqrt | pair_328 | 0.004 | 1877.073 | 6.535 | 6.607 | 13.391 | 12.888 | 1.058 |
| c1_sqrt | pair_329 | 0.004 | 1877.073 | 3.947 | 3.952 | 8.046 | 8.015 | 0.092 |
| c1_sqrt | pair_330 | 0.004 | 1877.073 | 7.294 | 7.315 | 14.837 | 14.421 | 0.903 |
| c1_sqrt | pair_331 | 0.004 | 1877.073 | 4.075 | 3.908 | 8.299 | 8.022 | 0.547 |
| c1_sqrt | pair_332 | 0.004 | 1877.073 | 3.931 | 3.954 | 7.998 | 8.010 | 0.116 |
| c1_sqrt | pair_333 | 0.004 | 1877.073 | 7.077 | 6.986 | 14.374 | 14.305 | 0.272 |
| c1_sqrt | pair_334 | 0.004 | 1877.073 | 7.126 | 7.237 | 14.545 | 14.411 | 0.380 |
| c1_sqrt | pair_335 | 0.004 | 1877.073 | 7.087 | 7.197 | 14.557 | 14.583 | 0.071 |
| c1_sqrt | pair_336 | 0.004 | 1877.073 | 7.347 | 7.002 | 14.544 | 14.677 | 0.255 |
| c1_sqrt | pair_337 | 0.004 | 1877.073 | 7.098 | 7.011 | 14.325 | 14.307 | 0.108 |
| c1_sqrt | pair_338 | 0.004 | 1877.073 | 6.347 | 6.229 | 12.818 | 12.771 | 0.081 |
| c1_sqrt | pair_339 | 0.004 | 1877.073 | 7.098 | 7.090 | 14.407 | 14.366 | 0.111 |
| c1_sqrt | pair_340 | 0.004 | 1877.073 | 7.135 | 7.074 | 14.415 | 14.276 | 0.262 |
| c1_sqrt | pair_341 | 0.004 | 1877.073 | 6.284 | 7.503 | 13.956 | 12.822 | 2.100 |
| c1_sqrt | pair_342 | 0.004 | 1877.073 | 12.518 | 9.617 | 22.408 | 22.009 | 10.066 |
| c1_sqrt | pair_343 | 0.004 | 1877.073 | 3.988 | 3.948 | 8.057 | 8.139 | 0.172 |
| c1_sqrt | pair_344 | 0.004 | 1877.073 | 7.033 | 7.061 | 14.281 | 14.289 | 0.039 |
| c1_sqrt | pair_345 | 0.004 | 1877.073 | 4.162 | 4.140 | 8.440 | 7.949 | 0.872 |
| c1_sqrt | pair_346 | 0.004 | 1877.073 | 3.948 | 3.909 | 7.963 | 7.937 | 0.100 |
| c1_sqrt | pair_347 | 0.004 | 1877.073 | 3.972 | 4.047 | 8.142 | 8.091 | 0.131 |
| c1_sqrt | pair_348 | 0.004 | 1877.073 | 3.914 | 3.884 | 7.932 | 7.937 | 0.060 |
| c1_sqrt | pair_349 | 0.004 | 1877.073 | 3.939 | 4.105 | 8.171 | 7.997 | 0.416 |
| c1_sqrt | pair_350 | 0.004 | 1877.073 | 3.916 | 3.905 | 7.965 | 7.919 | 0.112 |
| c1_sqrt | pair_351 | 0.004 | 1877.073 | 3.953 | 3.886 | 7.952 | 7.915 | 0.076 |
| c1_sqrt | pair_352 | 0.004 | 1877.073 | 3.997 | 3.907 | 8.053 | 8.043 | 0.078 |
| c1_sqrt | pair_353 | 0.004 | 1877.073 | 3.909 | 3.897 | 7.959 | 7.941 | 0.084 |
| c1_sqrt | pair_354 | 0.004 | 1877.073 | 3.982 | 3.875 | 7.924 | 7.921 | 0.041 |
| c1_sqrt | pair_355 | 0.004 | 1877.073 | 3.986 | 3.996 | 8.061 | 8.015 | 0.154 |
| c1_sqrt | pair_356 | 0.004 | 1877.073 | 4.012 | 3.896 | 8.046 | 8.082 | 0.091 |
| c1_sqrt | pair_357 | 0.004 | 1877.073 | 3.940 | 3.895 | 7.925 | 7.963 | 0.075 |
| c1_sqrt | pair_358 | 0.004 | 1877.073 | 3.974 | 3.928 | 8.046 | 8.013 | 0.132 |
| c1_sqrt | pair_359 | 0.004 | 1877.073 | 3.926 | 3.873 | 7.936 | 7.893 | 0.097 |
| c1_sqrt | pair_360 | 0.004 | 1877.073 | 3.962 | 3.875 | 8.089 | 8.039 | 0.193 |
| c1_sqrt | pair_361 | 0.004 | 1877.073 | 5.288 | 4.869 | 10.300 | 8.048 | 3.972 |
| c1_sqrt | pair_362 | 0.004 | 1877.073 | 4.614 | 4.044 | 8.785 | 7.989 | 1.406 |
| c1_sqrt | pair_363 | 0.004 | 1877.073 | 4.098 | 3.948 | 8.183 | 7.982 | 0.361 |
| c1_sqrt | pair_364 | 0.004 | 1877.073 | 4.169 | 4.133 | 8.443 | 8.075 | 0.771 |
| c1_sqrt | pair_365 | 0.004 | 1877.073 | 4.089 | 3.914 | 8.124 | 7.936 | 0.345 |
| c1_sqrt | pair_366 | 0.004 | 1877.073 | 3.947 | 3.948 | 7.982 | 8.012 | 0.115 |
| c1_sqrt | pair_367 | 0.004 | 1877.073 | 3.970 | 3.870 | 8.019 | 8.016 | 0.053 |
| c1_sqrt | pair_368 | 0.004 | 1877.073 | 3.965 | 3.865 | 7.923 | 7.851 | 0.175 |
| c1_sqrt | pair_369 | 0.004 | 1877.073 | 3.927 | 3.898 | 7.918 | 7.904 | 0.036 |
| c1_sqrt | pair_370 | 0.004 | 1877.073 | 7.825 | 7.801 | 15.872 | 15.912 | 0.099 |
| c1_sqrt | pair_371 | 0.004 | 1877.073 | 12.856 | 13.600 | 26.872 | 24.069 | 5.268 |
| c1_sqrt | pair_372 | 0.004 | 1877.073 | 7.935 | 7.790 | 15.948 | 15.968 | 0.125 |
| c1_sqrt | pair_373 | 0.004 | 1877.073 | 14.888 | 12.499 | 27.789 | 24.180 | 6.623 |
| c1_sqrt | pair_374 | 0.004 | 1877.073 | 6.442 | 6.283 | 12.895 | 12.899 | 0.079 |
| c1_sqrt | pair_375 | 0.004 | 1877.073 | 11.825 | 11.886 | 24.103 | 24.220 | 0.323 |
| c1_sqrt | pair_376 | 0.004 | 1877.073 | 11.850 | 11.795 | 24.066 | 24.140 | 0.255 |
| c1_sqrt | pair_377 | 0.004 | 1877.073 | 6.315 | 6.317 | 12.838 | 12.830 | 0.120 |
| c1_sqrt | pair_378 | 0.004 | 1877.073 | 6.277 | 6.216 | 12.692 | 12.722 | 0.058 |
| c1_sqrt | pair_379 | 0.004 | 1877.073 | 3.965 | 3.939 | 7.990 | 7.994 | 0.038 |
| c1_sqrt | pair_380 | 0.004 | 1877.073 | 7.071 | 7.060 | 14.372 | 14.311 | 0.242 |
| c1_sqrt | pair_381 | 0.004 | 1877.073 | 7.949 | 7.837 | 16.004 | 15.992 | 0.195 |
| c1_sqrt | pair_382 | 0.004 | 1877.073 | 7.853 | 7.800 | 15.906 | 15.923 | 0.104 |
| c1_sqrt | pair_383 | 0.004 | 1877.073 | 7.898 | 7.798 | 15.939 | 15.987 | 0.088 |
| c1_sqrt | pair_384 | 0.004 | 1877.073 | 6.409 | 6.526 | 13.187 | 13.069 | 0.484 |
| c1_sqrt | pair_385 | 0.004 | 1877.073 | 17.794 | 12.840 | 31.057 | 16.143 | 26.180 |
| c1_sqrt | pair_386 | 0.004 | 1877.073 | 10.120 | 10.182 | 20.592 | 15.916 | 8.161 |
| c1_sqrt | pair_387 | 0.004 | 1877.073 | 6.518 | 6.496 | 13.223 | 12.893 | 0.810 |
| c1_sqrt | pair_388 | 0.004 | 1877.073 | 6.763 | 6.892 | 13.990 | 13.127 | 1.707 |
| c1_sqrt | pair_389 | 0.004 | 1877.073 | 4.649 | 4.222 | 9.006 | 8.839 | 1.175 |
| c1_sqrt | pair_390 | 0.004 | 1877.073 | 7.403 | 7.266 | 14.879 | 14.436 | 0.910 |
| c1_sqrt | pair_391 | 0.004 | 1877.073 | 8.196 | 7.922 | 16.374 | 15.989 | 0.674 |
| c1_sqrt | pair_392 | 0.004 | 1877.073 | 21.109 | 20.738 | 43.335 | 40.196 | 5.593 |
| c1_sqrt | pair_393 | 0.004 | 1877.073 | 6.530 | 6.333 | 13.029 | 13.074 | 0.259 |
| c1_sqrt | pair_394 | 0.004 | 1877.073 | 13.000 | 12.919 | 26.312 | 26.055 | 0.614 |
| c1_sqrt | pair_395 | 0.004 | 1877.073 | 21.007 | 21.684 | 44.169 | 44.512 | 0.595 |
| c1_sqrt | pair_396 | 0.004 | 1877.073 | 7.474 | 9.179 | 16.914 | 14.708 | 4.194 |
| c1_sqrt | pair_397 | 0.004 | 1877.073 | 10.590 | 6.773 | 17.563 | 15.725 | 5.739 |
| c1_sqrt | pair_398 | 0.004 | 1877.073 | 4.247 | 4.414 | 8.820 | 8.022 | 1.412 |
| c1_sqrt | pair_399 | 0.004 | 1877.073 | 7.190 | 7.209 | 14.629 | 14.340 | 0.592 |
| c1_sqrt | pair_400 | 0.004 | 1877.073 | 8.085 | 7.908 | 16.211 | 15.969 | 0.577 |
| c1_sqrt | pair_401 | 0.004 | 1877.073 | 6.534 | 6.345 | 13.122 | 12.767 | 0.671 |
| c1_sqrt | pair_402 | 0.004 | 1877.073 | 8.393 | 9.628 | 18.321 | 18.421 | 2.324 |
| c1_sqrt | pair_403 | 0.004 | 1877.073 | 8.192 | 8.885 | 17.404 | 15.923 | 2.664 |
| c1_sqrt | pair_404 | 0.004 | 1877.073 | 6.633 | 6.586 | 13.427 | 13.100 | 0.870 |
| c1_sqrt | pair_405 | 0.004 | 1877.073 | 6.640 | 6.502 | 13.398 | 13.022 | 0.766 |
| c1_sqrt | pair_406 | 0.004 | 1877.073 | 3.939 | 3.897 | 7.963 | 7.962 | 0.046 |
| c1_sqrt | pair_407 | 0.004 | 1877.073 | 7.005 | 7.142 | 14.355 | 14.362 | 0.046 |
| c1_sqrt | pair_408 | 0.004 | 1877.073 | 6.337 | 6.315 | 12.884 | 12.794 | 0.181 |
| c1_sqrt | pair_409 | 0.004 | 1877.073 | 12.536 | 12.443 | 25.401 | 25.303 | 0.187 |
| c1_sqrt | pair_410 | 0.004 | 1877.073 | 19.634 | 19.722 | 40.236 | 40.277 | 0.327 |
| c1_sqrt | pair_411 | 0.004 | 1877.073 | 6.366 | 6.228 | 12.784 | 12.879 | 0.168 |
| c1_sqrt | pair_412 | 0.004 | 1877.073 | 6.324 | 6.177 | 12.707 | 12.709 | 0.112 |
| c1_sqrt | pair_413 | 0.004 | 1877.073 | 3.943 | 3.923 | 8.002 | 7.997 | 0.052 |
| c1_sqrt | pair_414 | 0.004 | 1877.073 | 7.153 | 7.189 | 14.570 | 14.414 | 0.276 |
| c1_sqrt | pair_415 | 0.004 | 1877.073 | 6.300 | 6.249 | 12.728 | 12.752 | 0.042 |
| c1_sqrt | pair_416 | 0.004 | 1877.073 | 6.483 | 6.417 | 13.161 | 12.798 | 0.697 |
| c1_sqrt | pair_417 | 0.004 | 1877.073 | 6.494 | 6.258 | 12.958 | 12.751 | 0.378 |
| c1_sqrt | pair_418 | 0.004 | 1877.073 | 6.669 | 6.250 | 13.725 | 13.895 | 0.902 |
| c1_sqrt | pair_419 | 0.004 | 1877.073 | 6.448 | 4.098 | 10.796 | 8.253 | 4.692 |
| c1_sqrt | pair_420 | 0.004 | 1877.073 | 6.322 | 6.238 | 12.743 | 12.780 | 0.065 |
| c1_sqrt | pair_421 | 0.004 | 1877.073 | 12.573 | 12.711 | 25.685 | 25.435 | 0.579 |
| c1_sqrt | pair_422 | 0.004 | 1877.073 | 6.356 | 6.319 | 12.921 | 12.859 | 0.325 |
| c1_sqrt | pair_423 | 0.004 | 1877.073 | 9.151 | 6.555 | 15.950 | 12.849 | 5.437 |
| c1_sqrt | pair_424 | 0.004 | 1877.073 | 3.919 | 3.888 | 9.842 | 7.904 | 3.365 |
| c1_sqrt | pair_425 | 0.004 | 1877.073 | 7.126 | 8.038 | 15.421 | 14.507 | 1.905 |
| c1_sqrt | pair_426 | 0.004 | 1877.073 | 6.343 | 6.434 | 12.996 | 12.767 | 0.428 |
| c1_sqrt | pair_427 | 0.004 | 1877.073 | 6.309 | 6.240 | 12.758 | 12.679 | 0.207 |
| c1_sqrt | pair_428 | 0.004 | 1877.073 | 3.914 | 3.887 | 7.921 | 7.951 | 0.127 |
| c1_sqrt | pair_429 | 0.004 | 1877.073 | 7.158 | 7.249 | 14.675 | 14.498 | 0.566 |
| c1_sqrt | pair_430 | 0.004 | 1877.073 | 6.701 | 6.204 | 13.139 | 12.761 | 0.671 |
| c1_sqrt | pair_431 | 0.004 | 1877.073 | 4.128 | 3.906 | 8.173 | 8.246 | 0.178 |
| c1_sqrt | pair_432 | 0.004 | 1877.073 | 6.291 | 6.271 | 12.752 | 12.789 | 0.070 |
| c1_sqrt | pair_433 | 0.004 | 1877.073 | 3.905 | 3.947 | 8.009 | 7.942 | 0.182 |
| c1_sqrt | pair_434 | 0.004 | 1877.073 | 6.345 | 6.176 | 12.747 | 12.643 | 0.220 |
| c1_sqrt | pair_435 | 0.004 | 1877.073 | 3.901 | 3.920 | 7.954 | 7.979 | 0.044 |
| corrected_fdxt | pair_001 | 0.004 | 1943.314 | 6.355 | 6.286 | 12.712 | 12.718 | 0.035 |
| corrected_fdxt | pair_002 | 0.004 | 1943.314 | 4.746 | 4.649 | 9.439 | 9.433 | 0.034 |
| corrected_fdxt | pair_003 | 0.004 | 1943.314 | 5.588 | 6.196 | 11.847 | 11.049 | 1.420 |
| corrected_fdxt | pair_004 | 0.004 | 1943.314 | 6.302 | 6.240 | 12.612 | 12.611 | 0.037 |
| corrected_fdxt | pair_005 | 0.004 | 1943.314 | 2.354 | 2.342 | 4.712 | 4.705 | 0.013 |
| corrected_fdxt | pair_006 | 0.004 | 1943.314 | 4.784 | 4.722 | 9.550 | 9.559 | 0.028 |
| corrected_fdxt | pair_007 | 0.004 | 1943.314 | 12.522 | 12.491 | 27.502 | 25.295 | 3.872 |
| corrected_fdxt | pair_008 | 0.004 | 1943.314 | 7.247 | 5.450 | 12.757 | 10.983 | 3.085 |
| corrected_fdxt | pair_009 | 0.004 | 1943.314 | 16.605 | 16.415 | 33.415 | 33.296 | 0.237 |
| corrected_fdxt | pair_010 | 0.004 | 1943.314 | 5.473 | 5.433 | 10.961 | 11.003 | 0.074 |
| corrected_fdxt | pair_011 | 0.004 | 1943.314 | 3.904 | 3.893 | 7.831 | 7.813 | 0.036 |
| corrected_fdxt | pair_012 | 0.004 | 1943.314 | 3.946 | 3.873 | 7.852 | 7.878 | 0.068 |
| corrected_fdxt | pair_013 | 0.004 | 1943.314 | 4.697 | 4.675 | 9.415 | 9.422 | 0.036 |
| corrected_fdxt | pair_014 | 0.004 | 1943.314 | 4.864 | 4.717 | 9.627 | 9.418 | 0.380 |
| corrected_fdxt | pair_015 | 0.004 | 1943.314 | 4.957 | 4.683 | 9.686 | 9.761 | 0.268 |
| corrected_fdxt | pair_016 | 0.004 | 1943.314 | 2.358 | 2.331 | 4.705 | 4.690 | 0.048 |
| corrected_fdxt | pair_017 | 0.004 | 1943.314 | 2.357 | 2.340 | 4.714 | 4.694 | 0.054 |
| corrected_fdxt | pair_018 | 0.004 | 1943.314 | 8.644 | 8.543 | 17.314 | 17.342 | 0.057 |
| corrected_fdxt | pair_019 | 0.004 | 1943.314 | 5.481 | 5.447 | 10.987 | 10.983 | 0.028 |
| corrected_fdxt | pair_020 | 0.004 | 1943.314 | 16.414 | 16.415 | 33.222 | 33.249 | 0.168 |
| corrected_fdxt | pair_021 | 0.004 | 1943.314 | 5.488 | 5.426 | 10.975 | 10.960 | 0.065 |
| corrected_fdxt | pair_022 | 0.004 | 1943.314 | 15.730 | 15.511 | 31.606 | 31.618 | 0.066 |
| corrected_fdxt | pair_023 | 0.004 | 1943.314 | 3.963 | 3.874 | 7.876 | 7.898 | 0.113 |
| corrected_fdxt | pair_024 | 0.004 | 1943.314 | 9.447 | 9.369 | 18.964 | 18.942 | 0.050 |
| corrected_fdxt | pair_025 | 0.004 | 1943.314 | 22.376 | 17.184 | 40.115 | 33.334 | 11.808 |
| corrected_fdxt | pair_026 | 0.004 | 1943.314 | 4.075 | 4.153 | 8.272 | 7.944 | 0.600 |
| corrected_fdxt | pair_027 | 0.004 | 1943.314 | 4.151 | 4.124 | 8.308 | 7.836 | 0.842 |
| corrected_fdxt | pair_028 | 0.004 | 1943.314 | 2.622 | 2.590 | 5.234 | 5.118 | 0.555 |
| corrected_fdxt | pair_029 | 0.004 | 1943.314 | 4.839 | 4.795 | 9.682 | 9.529 | 0.438 |
| corrected_fdxt | pair_030 | 0.004 | 1943.314 | 4.941 | 4.750 | 9.736 | 9.400 | 0.585 |
| corrected_fdxt | pair_031 | 0.004 | 1943.314 | 5.633 | 5.642 | 11.332 | 10.979 | 0.693 |
| corrected_fdxt | pair_032 | 0.004 | 1943.314 | 6.517 | 6.653 | 13.263 | 13.379 | 0.620 |
| corrected_fdxt | pair_033 | 0.004 | 1943.314 | 2.526 | 2.423 | 4.965 | 4.915 | 0.317 |
| corrected_fdxt | pair_034 | 0.004 | 1943.314 | 4.834 | 4.883 | 9.765 | 9.462 | 0.564 |
| corrected_fdxt | pair_035 | 0.004 | 1943.314 | 6.493 | 6.471 | 13.041 | 12.565 | 0.899 |
| corrected_fdxt | pair_036 | 0.004 | 1943.314 | 5.674 | 5.740 | 11.477 | 11.005 | 0.875 |
| corrected_fdxt | pair_037 | 0.004 | 1943.314 | 6.638 | 6.698 | 13.450 | 12.643 | 1.547 |
| corrected_fdxt | pair_038 | 0.004 | 1943.314 | 5.620 | 6.619 | 12.299 | 11.004 | 2.264 |
| corrected_fdxt | pair_039 | 0.004 | 1943.314 | 4.006 | 4.062 | 8.104 | 7.822 | 0.506 |
| corrected_fdxt | pair_040 | 0.004 | 1943.314 | 4.307 | 3.994 | 8.338 | 7.885 | 0.836 |
| corrected_fdxt | pair_041 | 0.004 | 1943.314 | 4.899 | 4.957 | 9.909 | 9.528 | 0.766 |
| corrected_fdxt | pair_042 | 0.004 | 1943.314 | 4.890 | 4.753 | 9.690 | 9.434 | 0.468 |
| corrected_fdxt | pair_043 | 0.004 | 1943.314 | 4.776 | 4.911 | 9.735 | 9.429 | 0.597 |
| corrected_fdxt | pair_044 | 0.004 | 1943.314 | 2.410 | 2.449 | 4.879 | 4.779 | 0.206 |
| corrected_fdxt | pair_045 | 0.004 | 1943.314 | 2.612 | 2.575 | 5.214 | 4.702 | 0.907 |
| corrected_fdxt | pair_046 | 0.004 | 1943.314 | 6.428 | 6.471 | 13.614 | 13.704 | 0.988 |
| corrected_fdxt | pair_047 | 0.004 | 1943.314 | 6.021 | 6.946 | 13.034 | 12.310 | 2.475 |
| corrected_fdxt | pair_048 | 0.004 | 1943.314 | 10.128 | 6.771 | 16.975 | 14.229 | 6.048 |
| corrected_fdxt | pair_049 | 0.004 | 1943.314 | 5.749 | 14.202 | 20.014 | 12.261 | 14.444 |
| corrected_fdxt | pair_050 | 0.004 | 1943.314 | 7.109 | 7.132 | 14.382 | 13.688 | 2.250 |
| corrected_fdxt | pair_051 | 0.004 | 1943.314 | 4.334 | 4.231 | 8.601 | 8.267 | 1.012 |
| corrected_fdxt | pair_052 | 0.004 | 1943.314 | 6.372 | 7.478 | 13.923 | 13.486 | 1.660 |
| corrected_fdxt | pair_053 | 0.004 | 1943.314 | 6.556 | 6.686 | 13.319 | 12.698 | 1.094 |
| corrected_fdxt | pair_054 | 0.004 | 1943.314 | 3.937 | 3.974 | 7.946 | 8.013 | 0.138 |
| corrected_fdxt | pair_055 | 0.004 | 1943.314 | 4.171 | 3.879 | 8.092 | 7.984 | 0.354 |
| corrected_fdxt | pair_056 | 0.004 | 1943.314 | 2.366 | 2.342 | 4.725 | 4.730 | 0.014 |
| corrected_fdxt | pair_057 | 0.004 | 1943.314 | 4.829 | 4.878 | 9.757 | 9.555 | 0.382 |
| corrected_fdxt | pair_058 | 0.004 | 1943.314 | 4.684 | 4.659 | 9.392 | 9.392 | 0.037 |
| corrected_fdxt | pair_059 | 0.004 | 1943.314 | 4.796 | 4.670 | 9.516 | 9.497 | 0.034 |
| corrected_fdxt | pair_060 | 0.004 | 1943.314 | 2.388 | 2.354 | 4.760 | 4.747 | 0.041 |
| corrected_fdxt | pair_061 | 0.004 | 1943.314 | 4.678 | 4.673 | 9.399 | 9.408 | 0.037 |
| corrected_fdxt | pair_062 | 0.004 | 1943.314 | 4.687 | 4.637 | 9.369 | 9.377 | 0.030 |
| corrected_fdxt | pair_063 | 0.004 | 1943.314 | 4.698 | 4.671 | 9.415 | 9.414 | 0.009 |
| corrected_fdxt | pair_064 | 0.004 | 1943.314 | 4.709 | 4.643 | 9.398 | 9.394 | 0.072 |
| corrected_fdxt | pair_065 | 0.004 | 1943.314 | 4.673 | 4.654 | 9.474 | 9.406 | 0.175 |
| corrected_fdxt | pair_066 | 0.004 | 1943.314 | 6.234 | 4.073 | 10.340 | 7.881 | 4.356 |
| corrected_fdxt | pair_067 | 0.004 | 1943.314 | 3.915 | 3.875 | 7.824 | 7.789 | 0.071 |
| corrected_fdxt | pair_068 | 0.004 | 1943.314 | 4.709 | 4.685 | 9.441 | 9.441 | 0.055 |
| corrected_fdxt | pair_069 | 0.004 | 1943.314 | 4.682 | 4.722 | 9.449 | 9.477 | 0.115 |
| corrected_fdxt | pair_070 | 0.004 | 1943.314 | 4.720 | 4.694 | 9.462 | 9.502 | 0.101 |
| corrected_fdxt | pair_071 | 0.004 | 1943.314 | 2.358 | 2.356 | 4.732 | 4.713 | 0.074 |
| corrected_fdxt | pair_072 | 0.004 | 1943.314 | 2.343 | 2.360 | 4.719 | 4.720 | 0.028 |
| corrected_fdxt | pair_073 | 0.004 | 1943.314 | 4.683 | 4.694 | 9.421 | 9.414 | 0.030 |
| corrected_fdxt | pair_074 | 0.004 | 1943.314 | 4.722 | 4.661 | 9.431 | 9.457 | 0.087 |
| corrected_fdxt | pair_075 | 0.004 | 1943.314 | 5.201 | 4.644 | 9.901 | 9.411 | 0.938 |
| corrected_fdxt | pair_076 | 0.004 | 1943.314 | 4.722 | 4.634 | 9.404 | 9.396 | 0.070 |
| corrected_fdxt | pair_077 | 0.004 | 1943.314 | 4.732 | 4.709 | 9.501 | 9.483 | 0.097 |
| corrected_fdxt | pair_078 | 0.004 | 1943.314 | 3.894 | 3.849 | 7.776 | 7.784 | 0.073 |
| corrected_fdxt | pair_079 | 0.004 | 1943.314 | 4.680 | 4.957 | 9.705 | 9.405 | 0.569 |
| corrected_fdxt | pair_080 | 0.004 | 1943.314 | 4.898 | 4.812 | 9.766 | 9.562 | 0.476 |
| corrected_fdxt | pair_081 | 0.004 | 1943.314 | 4.073 | 4.190 | 8.309 | 8.005 | 0.648 |
| corrected_fdxt | pair_082 | 0.004 | 1943.314 | 4.149 | 4.047 | 8.268 | 7.936 | 0.672 |
| corrected_fdxt | pair_083 | 0.004 | 1943.314 | 2.463 | 2.398 | 4.879 | 4.739 | 0.246 |
| corrected_fdxt | pair_084 | 0.004 | 1943.314 | 4.973 | 4.772 | 9.792 | 9.507 | 0.580 |
| corrected_fdxt | pair_085 | 0.004 | 1943.314 | 5.824 | 5.706 | 11.593 | 11.177 | 0.848 |
| corrected_fdxt | pair_086 | 0.004 | 1943.314 | 2.435 | 2.717 | 5.174 | 4.771 | 0.757 |
| corrected_fdxt | pair_087 | 0.004 | 1943.314 | 5.430 | 4.955 | 10.436 | 10.032 | 1.253 |
| corrected_fdxt | pair_088 | 0.004 | 1943.314 | 5.718 | 5.655 | 11.433 | 11.060 | 0.713 |
| corrected_fdxt | pair_089 | 0.004 | 1943.314 | 5.951 | 6.024 | 12.034 | 11.058 | 1.823 |
| corrected_fdxt | pair_090 | 0.004 | 1943.314 | 5.721 | 5.440 | 11.217 | 11.090 | 0.266 |
| corrected_fdxt | pair_091 | 0.004 | 1943.314 | 5.482 | 5.471 | 11.006 | 10.956 | 0.105 |
| corrected_fdxt | pair_092 | 0.004 | 1943.314 | 3.948 | 3.980 | 7.964 | 7.924 | 0.082 |
| corrected_fdxt | pair_093 | 0.004 | 1943.314 | 3.911 | 3.904 | 7.852 | 7.853 | 0.027 |
| corrected_fdxt | pair_094 | 0.004 | 1943.314 | 4.706 | 4.719 | 9.469 | 9.469 | 0.049 |
| corrected_fdxt | pair_095 | 0.004 | 1943.314 | 4.723 | 4.689 | 9.457 | 9.439 | 0.108 |
| corrected_fdxt | pair_096 | 0.004 | 1943.314 | 4.701 | 4.701 | 9.446 | 9.475 | 0.075 |
| corrected_fdxt | pair_097 | 0.004 | 1943.314 | 2.324 | 2.928 | 5.269 | 4.696 | 1.035 |
| corrected_fdxt | pair_098 | 0.004 | 1943.314 | 3.398 | 3.614 | 7.031 | 4.892 | 3.837 |
| corrected_fdxt | pair_099 | 0.004 | 1943.314 | 8.068 | 5.497 | 13.623 | 11.087 | 4.434 |
| corrected_fdxt | pair_100 | 0.004 | 1943.314 | 5.536 | 5.472 | 11.062 | 11.047 | 0.078 |
| corrected_fdxt | pair_101 | 0.004 | 1943.314 | 5.467 | 5.449 | 10.975 | 10.952 | 0.077 |
| corrected_fdxt | pair_102 | 0.004 | 1943.314 | 5.521 | 5.424 | 11.002 | 11.025 | 0.069 |
| corrected_fdxt | pair_103 | 0.004 | 1943.314 | 5.516 | 5.432 | 11.002 | 10.949 | 0.118 |
| corrected_fdxt | pair_104 | 0.004 | 1943.314 | 3.914 | 3.868 | 7.817 | 7.775 | 0.095 |
| corrected_fdxt | pair_105 | 0.004 | 1943.314 | 5.508 | 5.500 | 11.081 | 11.065 | 0.036 |
| corrected_fdxt | pair_106 | 0.004 | 1943.314 | 5.605 | 5.495 | 11.163 | 11.105 | 0.281 |
| corrected_fdxt | pair_107 | 0.004 | 1943.314 | 3.936 | 3.888 | 7.856 | 7.897 | 0.074 |
| corrected_fdxt | pair_108 | 0.004 | 1943.314 | 3.910 | 3.890 | 7.836 | 7.832 | 0.072 |
| corrected_fdxt | pair_109 | 0.004 | 1943.314 | 2.380 | 2.331 | 4.726 | 4.675 | 0.089 |
| corrected_fdxt | pair_110 | 0.004 | 1943.314 | 4.681 | 4.680 | 9.415 | 9.416 | 0.142 |
| corrected_fdxt | pair_111 | 0.004 | 1943.314 | 2.347 | 2.350 | 4.713 | 4.682 | 0.077 |
| corrected_fdxt | pair_112 | 0.004 | 1943.314 | 4.687 | 4.683 | 9.419 | 9.431 | 0.062 |
| corrected_fdxt | pair_113 | 0.004 | 1943.314 | 6.276 | 6.251 | 12.619 | 12.599 | 0.101 |
| corrected_fdxt | pair_114 | 0.004 | 1943.314 | 5.495 | 5.555 | 11.109 | 11.060 | 0.291 |
| corrected_fdxt | pair_115 | 0.004 | 1943.314 | 6.332 | 6.254 | 12.765 | 12.822 | 0.274 |
| corrected_fdxt | pair_116 | 0.004 | 1943.314 | 10.504 | 5.586 | 16.150 | 11.506 | 8.573 |
| corrected_fdxt | pair_117 | 0.004 | 1943.314 | 4.029 | 3.950 | 8.016 | 7.938 | 0.182 |
| corrected_fdxt | pair_118 | 0.004 | 1943.314 | 4.034 | 3.872 | 7.942 | 7.881 | 0.244 |
| corrected_fdxt | pair_119 | 0.004 | 1943.314 | 4.752 | 4.645 | 9.442 | 9.438 | 0.084 |
| corrected_fdxt | pair_120 | 0.004 | 1943.314 | 4.851 | 4.627 | 9.521 | 9.461 | 0.179 |
| corrected_fdxt | pair_121 | 0.004 | 1943.314 | 4.733 | 4.711 | 9.492 | 9.535 | 0.098 |
| corrected_fdxt | pair_122 | 0.004 | 1943.314 | 2.367 | 2.388 | 4.772 | 4.782 | 0.030 |
| corrected_fdxt | pair_123 | 0.004 | 1943.314 | 2.346 | 2.329 | 4.691 | 4.707 | 0.032 |
| corrected_fdxt | pair_124 | 0.004 | 1943.314 | 6.305 | 6.226 | 12.605 | 12.568 | 0.149 |
| corrected_fdxt | pair_125 | 0.004 | 1943.314 | 5.514 | 5.455 | 11.025 | 10.993 | 0.064 |
| corrected_fdxt | pair_126 | 0.004 | 1943.314 | 6.280 | 6.268 | 12.622 | 12.694 | 0.132 |
| corrected_fdxt | pair_127 | 0.004 | 1943.314 | 5.487 | 5.442 | 10.984 | 10.980 | 0.084 |
| corrected_fdxt | pair_128 | 0.004 | 1943.314 | 6.250 | 6.211 | 12.532 | 12.562 | 0.104 |
| corrected_fdxt | pair_129 | 0.004 | 1943.314 | 3.925 | 3.894 | 7.852 | 7.828 | 0.135 |
| corrected_fdxt | pair_130 | 0.004 | 1943.314 | 6.327 | 6.213 | 12.613 | 12.514 | 0.192 |
| corrected_fdxt | pair_131 | 0.004 | 1943.314 | 6.266 | 6.246 | 12.591 | 12.635 | 0.083 |
| corrected_fdxt | pair_132 | 0.004 | 1943.314 | 3.900 | 3.957 | 7.894 | 7.918 | 0.079 |
| corrected_fdxt | pair_133 | 0.004 | 1943.314 | 3.914 | 3.926 | 7.874 | 7.875 | 0.050 |
| corrected_fdxt | pair_134 | 0.004 | 1943.314 | 2.376 | 2.337 | 4.728 | 4.696 | 0.059 |
| corrected_fdxt | pair_135 | 0.004 | 1943.314 | 4.717 | 4.687 | 9.447 | 9.438 | 0.085 |
| corrected_fdxt | pair_136 | 0.004 | 1943.314 | 2.405 | 2.329 | 4.755 | 4.761 | 0.048 |
| corrected_fdxt | pair_137 | 0.004 | 1943.314 | 2.379 | 2.356 | 4.752 | 4.667 | 0.150 |
| corrected_fdxt | pair_138 | 0.004 | 1943.314 | 2.372 | 2.337 | 4.726 | 4.770 | 0.093 |
| corrected_fdxt | pair_139 | 0.004 | 1943.314 | 2.438 | 2.462 | 4.919 | 4.802 | 0.290 |
| corrected_fdxt | pair_140 | 0.004 | 1943.314 | 2.341 | 2.356 | 4.712 | 4.699 | 0.043 |
| corrected_fdxt | pair_141 | 0.004 | 1943.314 | 2.370 | 2.363 | 4.749 | 4.753 | 0.014 |
| corrected_fdxt | pair_142 | 0.004 | 1943.314 | 2.373 | 2.339 | 4.727 | 4.707 | 0.047 |
| corrected_fdxt | pair_143 | 0.004 | 1943.314 | 2.433 | 2.583 | 5.040 | 4.807 | 0.536 |
| corrected_fdxt | pair_144 | 0.004 | 1943.314 | 2.721 | 2.730 | 5.473 | 4.738 | 1.294 |
| corrected_fdxt | pair_145 | 0.004 | 1943.314 | 2.732 | 2.822 | 5.578 | 4.793 | 1.483 |
| corrected_fdxt | pair_146 | 0.004 | 1943.314 | 2.838 | 3.434 | 6.972 | 7.628 | 1.983 |
| corrected_fdxt | pair_147 | 0.004 | 1943.314 | 3.799 | 6.141 | 9.970 | 7.189 | 7.108 |
| corrected_fdxt | pair_148 | 0.004 | 1943.314 | 2.913 | 2.906 | 5.847 | 5.118 | 1.618 |
| corrected_fdxt | pair_149 | 0.004 | 1943.314 | 3.098 | 2.846 | 5.975 | 5.609 | 1.507 |
| corrected_fdxt | pair_150 | 0.004 | 1943.314 | 2.824 | 2.799 | 5.647 | 4.842 | 1.537 |
| corrected_fdxt | pair_151 | 0.004 | 1943.314 | 2.756 | 2.726 | 5.502 | 4.868 | 1.336 |
| corrected_fdxt | pair_152 | 0.004 | 1943.314 | 2.426 | 2.510 | 4.962 | 4.760 | 0.411 |
| corrected_fdxt | pair_153 | 0.004 | 1943.314 | 2.878 | 2.820 | 5.717 | 4.752 | 1.736 |
| corrected_fdxt | pair_154 | 0.004 | 1943.314 | 2.427 | 2.377 | 4.821 | 4.702 | 0.213 |
| corrected_fdxt | pair_155 | 0.004 | 1943.314 | 2.393 | 2.352 | 4.761 | 4.756 | 0.025 |
| corrected_fdxt | pair_156 | 0.004 | 1943.314 | 2.349 | 2.337 | 4.701 | 4.710 | 0.033 |
| corrected_fdxt | pair_157 | 0.004 | 1943.314 | 2.339 | 2.334 | 4.692 | 4.670 | 0.045 |
| corrected_fdxt | pair_158 | 0.004 | 1943.314 | 2.365 | 2.427 | 4.808 | 4.811 | 0.119 |
| corrected_fdxt | pair_159 | 0.004 | 1943.314 | 2.390 | 2.350 | 4.756 | 4.758 | 0.005 |
| corrected_fdxt | pair_160 | 0.004 | 1943.314 | 4.731 | 4.705 | 9.485 | 9.464 | 0.101 |
| corrected_fdxt | pair_161 | 0.004 | 1943.314 | 4.778 | 8.828 | 13.655 | 9.498 | 7.232 |
| corrected_fdxt | pair_162 | 0.004 | 1943.314 | 9.810 | 6.226 | 16.088 | 9.521 | 11.423 |
| corrected_fdxt | pair_163 | 0.004 | 1943.314 | 5.405 | 5.017 | 10.480 | 9.561 | 1.780 |
| corrected_fdxt | pair_164 | 0.004 | 1943.314 | 4.299 | 3.939 | 8.277 | 7.923 | 0.665 |
| corrected_fdxt | pair_165 | 0.004 | 1943.314 | 4.015 | 4.077 | 8.131 | 7.853 | 0.550 |
| corrected_fdxt | pair_166 | 0.004 | 1943.314 | 4.852 | 5.030 | 9.930 | 9.487 | 0.816 |
| corrected_fdxt | pair_167 | 0.004 | 1943.314 | 4.995 | 4.818 | 9.860 | 9.432 | 0.853 |
| corrected_fdxt | pair_168 | 0.004 | 1943.314 | 4.786 | 4.930 | 9.767 | 9.621 | 0.412 |
| corrected_fdxt | pair_169 | 0.004 | 1943.314 | 2.460 | 2.434 | 4.910 | 4.856 | 0.270 |
| corrected_fdxt | pair_170 | 0.004 | 1943.314 | 2.378 | 2.461 | 4.859 | 4.718 | 0.310 |
| corrected_fdxt | pair_171 | 0.004 | 1943.314 | 4.864 | 4.858 | 9.774 | 9.492 | 0.575 |
| corrected_fdxt | pair_172 | 0.004 | 1943.314 | 4.858 | 4.822 | 9.728 | 9.484 | 0.487 |
| corrected_fdxt | pair_173 | 0.004 | 1943.314 | 4.917 | 4.856 | 9.821 | 9.564 | 0.584 |
| corrected_fdxt | pair_174 | 0.004 | 1943.314 | 4.856 | 4.887 | 9.789 | 9.599 | 0.513 |
| corrected_fdxt | pair_175 | 0.004 | 1943.314 | 4.873 | 4.820 | 9.738 | 9.540 | 0.532 |
| corrected_fdxt | pair_176 | 0.004 | 1943.314 | 4.081 | 4.016 | 8.132 | 7.888 | 0.435 |
| corrected_fdxt | pair_177 | 0.004 | 1943.314 | 4.914 | 4.916 | 9.894 | 9.563 | 0.690 |
| corrected_fdxt | pair_178 | 0.004 | 1943.314 | 4.961 | 4.934 | 9.957 | 10.173 | 0.375 |
| corrected_fdxt | pair_179 | 0.004 | 1943.314 | 4.174 | 4.103 | 8.313 | 8.074 | 0.498 |
| corrected_fdxt | pair_180 | 0.004 | 1943.314 | 4.121 | 4.038 | 8.204 | 7.993 | 0.560 |
| corrected_fdxt | pair_181 | 0.004 | 1943.314 | 2.499 | 8.593 | 11.147 | 5.142 | 10.840 |
| corrected_fdxt | pair_182 | 0.004 | 1943.314 | 8.307 | 4.887 | 13.250 | 10.384 | 5.877 |
| corrected_fdxt | pair_183 | 0.004 | 1943.314 | 5.643 | 5.708 | 11.416 | 11.120 | 0.612 |
| corrected_fdxt | pair_184 | 0.004 | 1943.314 | 13.452 | 12.865 | 26.651 | 25.252 | 2.648 |
| corrected_fdxt | pair_185 | 0.004 | 1943.314 | 5.485 | 5.886 | 11.435 | 11.018 | 0.804 |
| corrected_fdxt | pair_186 | 0.004 | 1943.314 | 4.145 | 3.861 | 8.044 | 7.853 | 0.369 |
| corrected_fdxt | pair_187 | 0.004 | 1943.314 | 3.886 | 3.901 | 7.821 | 7.823 | 0.030 |
| corrected_fdxt | pair_188 | 0.004 | 1943.314 | 4.776 | 4.642 | 9.467 | 9.467 | 0.031 |
| corrected_fdxt | pair_189 | 0.004 | 1943.314 | 4.714 | 4.716 | 9.491 | 9.455 | 0.111 |
| corrected_fdxt | pair_190 | 0.004 | 1943.314 | 4.733 | 5.060 | 9.855 | 9.484 | 0.673 |
| corrected_fdxt | pair_191 | 0.004 | 1943.314 | 2.410 | 2.333 | 4.760 | 4.721 | 0.070 |
| corrected_fdxt | pair_192 | 0.004 | 1943.314 | 2.418 | 2.344 | 4.783 | 4.772 | 0.025 |
| corrected_fdxt | pair_193 | 0.004 | 1943.314 | 8.593 | 8.642 | 17.360 | 17.272 | 0.179 |
| corrected_fdxt | pair_194 | 0.004 | 1943.314 | 5.548 | 5.509 | 11.119 | 10.987 | 0.303 |
| corrected_fdxt | pair_195 | 0.004 | 1943.314 | 12.611 | 12.548 | 25.399 | 25.402 | 0.227 |
| corrected_fdxt | pair_196 | 0.004 | 1943.314 | 5.580 | 5.741 | 11.380 | 11.191 | 0.532 |
| corrected_fdxt | pair_197 | 0.004 | 1943.314 | 12.475 | 12.559 | 25.266 | 25.290 | 0.143 |
| corrected_fdxt | pair_198 | 0.004 | 1943.314 | 3.887 | 3.890 | 7.811 | 7.816 | 0.019 |
| corrected_fdxt | pair_199 | 0.004 | 1943.314 | 9.341 | 9.437 | 18.919 | 18.934 | 0.036 |
| corrected_fdxt | pair_200 | 0.004 | 1943.314 | 12.862 | 13.247 | 26.352 | 25.852 | 1.553 |
| corrected_fdxt | pair_201 | 0.004 | 1943.314 | 3.983 | 3.955 | 7.971 | 7.863 | 0.197 |
| corrected_fdxt | pair_202 | 0.004 | 1943.314 | 3.942 | 3.915 | 7.890 | 7.819 | 0.157 |
| corrected_fdxt | pair_203 | 0.004 | 1943.314 | 2.378 | 2.357 | 4.750 | 4.712 | 0.110 |
| corrected_fdxt | pair_204 | 0.004 | 1943.314 | 4.825 | 4.700 | 9.582 | 9.667 | 0.156 |
| corrected_fdxt | pair_205 | 0.004 | 1943.314 | 5.543 | 9.199 | 14.803 | 11.256 | 6.331 |
| corrected_fdxt | pair_206 | 0.004 | 1943.314 | 9.616 | 8.074 | 17.749 | 11.176 | 11.585 |
| corrected_fdxt | pair_207 | 0.004 | 1943.314 | 4.288 | 3.930 | 8.255 | 8.138 | 0.516 |
| corrected_fdxt | pair_208 | 0.004 | 1943.314 | 4.002 | 7.652 | 11.689 | 7.875 | 6.627 |
| corrected_fdxt | pair_209 | 0.004 | 1943.314 | 10.241 | 5.060 | 15.352 | 9.378 | 10.376 |
| corrected_fdxt | pair_210 | 0.004 | 1943.314 | 4.941 | 5.003 | 9.991 | 9.439 | 1.095 |
| corrected_fdxt | pair_211 | 0.004 | 1943.314 | 4.899 | 5.070 | 10.020 | 9.481 | 1.007 |
| corrected_fdxt | pair_212 | 0.004 | 1943.314 | 2.467 | 2.496 | 4.981 | 4.743 | 0.458 |
| corrected_fdxt | pair_213 | 0.004 | 1943.314 | 2.512 | 2.466 | 4.996 | 4.716 | 0.528 |
| corrected_fdxt | pair_214 | 0.004 | 1943.314 | 5.882 | 5.849 | 11.791 | 11.108 | 1.300 |
| corrected_fdxt | pair_215 | 0.004 | 1943.314 | 6.075 | 5.899 | 12.039 | 11.851 | 1.164 |
| corrected_fdxt | pair_216 | 0.004 | 1943.314 | 5.967 | 5.777 | 11.808 | 11.201 | 1.188 |
| corrected_fdxt | pair_217 | 0.004 | 1943.314 | 5.803 | 5.872 | 11.738 | 10.965 | 1.357 |
| corrected_fdxt | pair_218 | 0.004 | 1943.314 | 5.818 | 5.837 | 11.711 | 11.046 | 1.168 |
| corrected_fdxt | pair_219 | 0.004 | 1943.314 | 4.208 | 4.237 | 8.483 | 8.046 | 1.037 |
| corrected_fdxt | pair_220 | 0.004 | 1943.314 | 5.901 | 5.980 | 11.940 | 10.971 | 1.679 |
| corrected_fdxt | pair_221 | 0.004 | 1943.314 | 5.825 | 5.774 | 11.663 | 11.042 | 1.126 |
| corrected_fdxt | pair_222 | 0.004 | 1943.314 | 4.081 | 4.142 | 8.261 | 7.771 | 0.853 |
| corrected_fdxt | pair_223 | 0.004 | 1943.314 | 4.174 | 4.243 | 8.461 | 7.835 | 1.102 |
| corrected_fdxt | pair_224 | 0.004 | 1943.314 | 2.585 | 2.510 | 5.113 | 4.700 | 0.719 |
| corrected_fdxt | pair_225 | 0.004 | 1943.314 | 5.057 | 5.263 | 10.373 | 9.457 | 1.679 |
| corrected_fdxt | pair_226 | 0.004 | 1943.314 | 5.955 | 5.919 | 11.945 | 10.988 | 1.701 |
| corrected_fdxt | pair_227 | 0.004 | 1943.314 | 4.237 | 4.303 | 8.575 | 7.856 | 1.260 |
| corrected_fdxt | pair_228 | 0.004 | 1943.314 | 4.288 | 4.632 | 8.957 | 9.103 | 0.917 |
| corrected_fdxt | pair_229 | 0.004 | 1943.314 | 5.833 | 4.902 | 10.788 | 10.124 | 1.888 |
| corrected_fdxt | pair_230 | 0.004 | 1943.314 | 4.892 | 5.031 | 9.970 | 9.383 | 1.027 |
| corrected_fdxt | pair_231 | 0.004 | 1943.314 | 4.892 | 4.922 | 9.871 | 9.493 | 0.772 |
| corrected_fdxt | pair_232 | 0.004 | 1943.314 | 2.459 | 2.447 | 4.925 | 4.747 | 0.366 |
| corrected_fdxt | pair_233 | 0.004 | 1943.314 | 2.507 | 2.403 | 4.934 | 4.714 | 0.399 |
| corrected_fdxt | pair_234 | 0.004 | 1943.314 | 9.215 | 9.071 | 18.432 | 17.294 | 2.012 |
| corrected_fdxt | pair_235 | 0.004 | 1943.314 | 7.094 | 6.291 | 13.451 | 11.385 | 3.825 |
| corrected_fdxt | pair_236 | 0.004 | 1943.314 | 17.548 | 18.402 | 36.367 | 36.799 | 2.316 |
| corrected_fdxt | pair_237 | 0.004 | 1943.314 | 5.598 | 5.653 | 11.308 | 11.212 | 0.335 |
| corrected_fdxt | pair_238 | 0.004 | 1943.314 | 16.178 | 16.319 | 32.963 | 32.010 | 1.905 |
| corrected_fdxt | pair_239 | 0.004 | 1943.314 | 3.968 | 4.221 | 8.230 | 8.212 | 0.450 |
| corrected_fdxt | pair_240 | 0.004 | 1943.314 | 10.853 | 10.860 | 21.858 | 21.297 | 2.874 |
| corrected_fdxt | pair_241 | 0.004 | 1943.314 | 16.631 | 16.413 | 33.443 | 33.282 | 0.365 |
| corrected_fdxt | pair_242 | 0.004 | 1943.314 | 3.969 | 3.891 | 7.894 | 7.884 | 0.020 |
| corrected_fdxt | pair_243 | 0.004 | 1943.314 | 3.928 | 3.895 | 7.861 | 7.852 | 0.024 |
| corrected_fdxt | pair_244 | 0.004 | 1943.314 | 3.758 | 2.474 | 6.249 | 4.782 | 2.656 |
| corrected_fdxt | pair_245 | 0.004 | 1943.314 | 6.061 | 6.831 | 12.945 | 9.501 | 6.064 |
| corrected_fdxt | pair_246 | 0.004 | 1943.314 | 4.090 | 4.502 | 8.628 | 7.868 | 1.407 |
| corrected_fdxt | pair_247 | 0.004 | 1943.314 | 3.955 | 4.189 | 8.179 | 7.866 | 0.574 |
| corrected_fdxt | pair_248 | 0.004 | 1943.314 | 4.931 | 5.044 | 10.020 | 9.425 | 1.152 |
| corrected_fdxt | pair_249 | 0.004 | 1943.314 | 4.847 | 5.053 | 9.951 | 9.495 | 0.814 |
| corrected_fdxt | pair_250 | 0.004 | 1943.314 | 4.896 | 4.967 | 9.908 | 9.467 | 0.833 |
| corrected_fdxt | pair_251 | 0.004 | 1943.314 | 2.385 | 2.503 | 4.911 | 4.732 | 0.372 |
| corrected_fdxt | pair_252 | 0.004 | 1943.314 | 2.507 | 2.385 | 4.910 | 4.753 | 0.336 |
| corrected_fdxt | pair_253 | 0.004 | 1943.314 | 5.772 | 5.600 | 11.427 | 11.033 | 0.721 |
| corrected_fdxt | pair_254 | 0.004 | 1943.314 | 5.670 | 5.562 | 11.292 | 11.025 | 0.524 |
| corrected_fdxt | pair_255 | 0.004 | 1943.314 | 7.843 | 11.295 | 19.418 | 12.072 | 13.596 |
| corrected_fdxt | pair_256 | 0.004 | 1943.314 | 6.489 | 5.890 | 12.439 | 11.977 | 1.699 |
| corrected_fdxt | pair_257 | 0.004 | 1943.314 | 6.109 | 5.877 | 12.052 | 12.281 | 0.938 |
| corrected_fdxt | pair_258 | 0.004 | 1943.314 | 4.315 | 4.265 | 8.617 | 8.470 | 0.503 |
| corrected_fdxt | pair_259 | 0.004 | 1943.314 | 5.927 | 5.791 | 11.779 | 11.949 | 0.695 |
| corrected_fdxt | pair_260 | 0.004 | 1943.314 | 5.998 | 5.783 | 11.846 | 11.982 | 0.833 |
| corrected_fdxt | pair_261 | 0.004 | 1943.314 | 4.250 | 4.100 | 8.390 | 8.600 | 0.476 |
| corrected_fdxt | pair_262 | 0.004 | 1943.314 | 4.210 | 4.191 | 8.442 | 8.454 | 0.530 |
| corrected_fdxt | pair_263 | 0.004 | 1943.314 | 2.492 | 2.513 | 5.023 | 5.125 | 0.314 |
| corrected_fdxt | pair_264 | 0.004 | 1943.314 | 5.201 | 5.388 | 10.645 | 10.316 | 1.454 |
| corrected_fdxt | pair_265 | 0.004 | 1943.314 | 4.336 | 4.125 | 8.507 | 8.750 | 0.624 |
| corrected_fdxt | pair_266 | 0.004 | 1943.314 | 4.112 | 4.145 | 8.293 | 8.387 | 0.470 |
| corrected_fdxt | pair_267 | 0.004 | 1943.314 | 4.186 | 4.244 | 8.466 | 8.648 | 0.619 |
| corrected_fdxt | pair_268 | 0.004 | 1943.314 | 4.151 | 4.207 | 8.395 | 8.644 | 0.476 |
| corrected_fdxt | pair_269 | 0.004 | 1943.314 | 2.542 | 2.534 | 5.094 | 5.220 | 0.295 |
| corrected_fdxt | pair_270 | 0.004 | 1943.314 | 2.465 | 2.483 | 4.965 | 5.030 | 0.254 |
| corrected_fdxt | pair_271 | 0.004 | 1943.314 | 4.157 | 4.166 | 8.357 | 8.537 | 0.403 |
| corrected_fdxt | pair_272 | 0.004 | 1943.314 | 4.083 | 4.155 | 8.270 | 8.091 | 0.623 |
| corrected_fdxt | pair_273 | 0.004 | 1943.314 | 4.107 | 4.005 | 8.147 | 7.804 | 0.594 |
| corrected_fdxt | pair_274 | 0.004 | 1943.314 | 4.018 | 4.022 | 8.076 | 7.870 | 0.441 |
| corrected_fdxt | pair_275 | 0.004 | 1943.314 | 4.061 | 4.070 | 8.165 | 7.949 | 0.523 |
| corrected_fdxt | pair_276 | 0.004 | 1943.314 | 4.023 | 4.008 | 8.070 | 7.874 | 0.394 |
| corrected_fdxt | pair_277 | 0.004 | 1943.314 | 4.085 | 4.124 | 8.243 | 7.833 | 0.757 |
| corrected_fdxt | pair_278 | 0.004 | 1943.314 | 4.136 | 4.093 | 8.269 | 7.947 | 0.580 |
| corrected_fdxt | pair_279 | 0.004 | 1943.314 | 4.216 | 4.020 | 8.276 | 7.920 | 0.749 |
| corrected_fdxt | pair_280 | 0.004 | 1943.314 | 4.125 | 4.117 | 8.278 | 8.094 | 0.548 |
| corrected_fdxt | pair_281 | 0.004 | 1943.314 | 2.533 | 2.497 | 5.047 | 5.146 | 0.288 |
| corrected_fdxt | pair_282 | 0.004 | 1943.314 | 4.105 | 3.996 | 8.142 | 8.067 | 0.372 |
| corrected_fdxt | pair_283 | 0.004 | 1943.314 | 4.111 | 4.035 | 8.180 | 7.895 | 0.500 |
| corrected_fdxt | pair_284 | 0.004 | 1943.314 | 4.002 | 3.995 | 8.142 | 8.157 | 0.413 |
| corrected_fdxt | pair_285 | 0.004 | 1943.314 | 9.073 | 3.993 | 13.101 | 8.519 | 8.518 |
| corrected_fdxt | pair_286 | 0.004 | 1943.314 | 2.491 | 2.395 | 4.903 | 4.698 | 0.399 |
| corrected_fdxt | pair_287 | 0.004 | 1943.314 | 2.344 | 2.431 | 4.793 | 4.793 | 0.021 |
| corrected_fdxt | pair_288 | 0.004 | 1943.314 | 3.959 | 4.003 | 7.995 | 7.832 | 0.285 |
| corrected_fdxt | pair_289 | 0.004 | 1943.314 | 4.073 | 3.998 | 8.105 | 8.100 | 0.312 |
| corrected_fdxt | pair_290 | 0.004 | 1943.314 | 4.112 | 4.033 | 8.179 | 7.865 | 0.632 |
| corrected_fdxt | pair_291 | 0.004 | 1943.314 | 4.079 | 4.219 | 8.344 | 7.775 | 1.030 |
| corrected_fdxt | pair_292 | 0.004 | 1943.314 | 4.919 | 5.098 | 10.066 | 7.998 | 3.685 |
| corrected_fdxt | pair_293 | 0.004 | 1943.314 | 4.533 | 4.101 | 8.677 | 7.803 | 1.523 |
| corrected_fdxt | pair_294 | 0.004 | 1943.314 | 4.458 | 4.376 | 8.875 | 7.979 | 1.648 |
| corrected_fdxt | pair_295 | 0.004 | 1943.314 | 4.442 | 4.984 | 9.474 | 7.877 | 2.834 |
| corrected_fdxt | pair_296 | 0.004 | 1943.314 | 4.972 | 4.355 | 9.372 | 7.907 | 2.572 |
| corrected_fdxt | pair_297 | 0.004 | 1943.314 | 5.073 | 4.632 | 9.740 | 7.832 | 3.355 |
| corrected_fdxt | pair_298 | 0.004 | 1943.314 | 2.458 | 2.866 | 5.350 | 4.759 | 1.090 |
| corrected_fdxt | pair_299 | 0.004 | 1943.314 | 4.474 | 4.527 | 9.043 | 7.840 | 2.114 |
| corrected_fdxt | pair_300 | 0.004 | 1943.314 | 5.943 | 5.729 | 11.810 | 9.352 | 4.293 |
| corrected_fdxt | pair_301 | 0.004 | 1943.314 | 4.769 | 5.029 | 9.847 | 9.563 | 0.682 |
| corrected_fdxt | pair_302 | 0.004 | 1943.314 | 2.310 | 2.341 | 4.667 | 4.668 | 0.018 |
| corrected_fdxt | pair_303 | 0.004 | 1943.314 | 3.098 | 2.394 | 5.510 | 4.762 | 1.297 |
| corrected_fdxt | pair_304 | 0.004 | 1943.314 | 5.069 | 4.668 | 9.783 | 9.393 | 0.732 |
| corrected_fdxt | pair_305 | 0.004 | 1943.314 | 6.337 | 6.340 | 12.724 | 9.551 | 5.563 |
| corrected_fdxt | pair_306 | 0.004 | 1943.314 | 5.931 | 4.634 | 10.609 | 9.384 | 2.180 |
| corrected_fdxt | pair_307 | 0.004 | 1943.314 | 4.887 | 4.907 | 9.842 | 9.754 | 0.232 |
| corrected_fdxt | pair_308 | 0.004 | 1943.314 | 4.685 | 4.697 | 9.426 | 9.393 | 0.059 |
| corrected_fdxt | pair_309 | 0.004 | 1943.314 | 3.982 | 3.883 | 7.902 | 7.875 | 0.092 |
| corrected_fdxt | pair_310 | 0.004 | 1943.314 | 4.646 | 4.635 | 9.323 | 9.356 | 0.068 |
| corrected_fdxt | pair_311 | 0.004 | 1943.314 | 4.800 | 4.772 | 9.614 | 9.373 | 0.467 |
| corrected_fdxt | pair_312 | 0.004 | 1943.314 | 3.912 | 3.906 | 7.851 | 7.831 | 0.049 |
| corrected_fdxt | pair_313 | 0.004 | 1943.314 | 3.936 | 3.842 | 7.814 | 7.786 | 0.057 |
| corrected_fdxt | pair_314 | 0.004 | 1943.314 | 2.384 | 2.349 | 4.749 | 4.755 | 0.011 |
| corrected_fdxt | pair_315 | 0.004 | 1943.314 | 4.791 | 4.740 | 9.680 | 9.365 | 0.595 |
| corrected_fdxt | pair_316 | 0.004 | 1943.314 | 6.043 | 5.365 | 11.453 | 9.391 | 3.614 |
| corrected_fdxt | pair_317 | 0.004 | 1943.314 | 2.444 | 2.496 | 4.963 | 5.004 | 0.265 |
| corrected_fdxt | pair_318 | 0.004 | 1943.314 | 2.435 | 2.400 | 4.852 | 4.717 | 0.265 |
| corrected_fdxt | pair_319 | 0.004 | 1943.314 | 4.868 | 4.685 | 9.598 | 9.494 | 0.380 |
| corrected_fdxt | pair_320 | 0.004 | 1943.314 | 4.698 | 4.680 | 9.424 | 9.421 | 0.111 |
| corrected_fdxt | pair_321 | 0.004 | 1943.314 | 4.717 | 4.671 | 9.431 | 9.430 | 0.054 |
| corrected_fdxt | pair_322 | 0.004 | 1943.314 | 4.788 | 4.650 | 9.481 | 9.487 | 0.152 |
| corrected_fdxt | pair_323 | 0.004 | 1943.314 | 4.715 | 4.696 | 9.460 | 9.455 | 0.049 |
| corrected_fdxt | pair_324 | 0.004 | 1943.314 | 4.000 | 3.866 | 7.900 | 7.797 | 0.196 |
| corrected_fdxt | pair_325 | 0.004 | 1943.314 | 4.744 | 4.665 | 9.453 | 9.425 | 0.173 |
| corrected_fdxt | pair_326 | 0.004 | 1943.314 | 4.680 | 4.671 | 9.393 | 9.343 | 0.092 |
| corrected_fdxt | pair_327 | 0.004 | 1943.314 | 3.897 | 3.897 | 7.827 | 7.848 | 0.058 |
| corrected_fdxt | pair_328 | 0.004 | 1943.314 | 3.919 | 3.857 | 7.809 | 7.770 | 0.071 |
| corrected_fdxt | pair_329 | 0.004 | 1943.314 | 2.343 | 2.324 | 4.682 | 4.701 | 0.051 |
| corrected_fdxt | pair_330 | 0.004 | 1943.314 | 4.738 | 4.656 | 9.438 | 9.436 | 0.011 |
| corrected_fdxt | pair_331 | 0.004 | 1943.314 | 2.335 | 2.352 | 4.703 | 4.709 | 0.014 |
| corrected_fdxt | pair_332 | 0.004 | 1943.314 | 2.382 | 2.398 | 4.797 | 4.786 | 0.147 |
| corrected_fdxt | pair_333 | 0.004 | 1943.314 | 4.718 | 4.781 | 9.546 | 9.402 | 0.332 |
| corrected_fdxt | pair_334 | 0.004 | 1943.314 | 4.733 | 4.725 | 9.503 | 9.432 | 0.133 |
| corrected_fdxt | pair_335 | 0.004 | 1943.314 | 4.731 | 4.819 | 9.597 | 9.470 | 0.274 |
| corrected_fdxt | pair_336 | 0.004 | 1943.314 | 4.985 | 5.209 | 10.244 | 9.383 | 1.520 |
| corrected_fdxt | pair_337 | 0.004 | 1943.314 | 5.321 | 5.010 | 10.386 | 9.420 | 1.821 |
| corrected_fdxt | pair_338 | 0.004 | 1943.314 | 4.153 | 4.009 | 8.195 | 7.822 | 0.666 |
| corrected_fdxt | pair_339 | 0.004 | 1943.314 | 4.906 | 4.936 | 9.889 | 9.400 | 0.890 |
| corrected_fdxt | pair_340 | 0.004 | 1943.314 | 5.091 | 5.519 | 10.673 | 9.461 | 2.189 |
| corrected_fdxt | pair_341 | 0.004 | 1943.314 | 4.220 | 3.986 | 8.256 | 7.829 | 0.787 |
| corrected_fdxt | pair_342 | 0.004 | 1943.314 | 4.397 | 4.125 | 8.560 | 7.894 | 1.276 |
| corrected_fdxt | pair_343 | 0.004 | 1943.314 | 2.425 | 2.344 | 4.785 | 4.757 | 0.070 |
| corrected_fdxt | pair_344 | 0.004 | 1943.314 | 4.657 | 4.788 | 9.490 | 9.489 | 0.124 |
| corrected_fdxt | pair_345 | 0.004 | 1943.314 | 2.485 | 3.293 | 5.794 | 4.735 | 1.850 |
| corrected_fdxt | pair_346 | 0.004 | 1943.314 | 2.501 | 2.478 | 5.001 | 4.725 | 0.525 |
| corrected_fdxt | pair_347 | 0.004 | 1943.314 | 2.394 | 2.350 | 4.762 | 4.725 | 0.071 |
| corrected_fdxt | pair_348 | 0.004 | 1943.314 | 2.394 | 2.320 | 4.731 | 4.676 | 0.140 |
| corrected_fdxt | pair_349 | 0.004 | 1943.314 | 2.336 | 2.356 | 4.708 | 4.694 | 0.045 |
| corrected_fdxt | pair_350 | 0.004 | 1943.314 | 2.346 | 2.368 | 4.729 | 4.747 | 0.058 |
| corrected_fdxt | pair_351 | 0.004 | 1943.314 | 2.346 | 2.401 | 4.763 | 4.765 | 0.075 |
| corrected_fdxt | pair_352 | 0.004 | 1943.314 | 2.356 | 2.371 | 4.743 | 4.770 | 0.050 |
| corrected_fdxt | pair_353 | 0.004 | 1943.314 | 2.353 | 2.369 | 4.739 | 4.771 | 0.075 |
| corrected_fdxt | pair_354 | 0.004 | 1943.314 | 2.384 | 2.349 | 4.750 | 4.780 | 0.067 |
| corrected_fdxt | pair_355 | 0.004 | 1943.314 | 2.384 | 2.353 | 4.752 | 4.775 | 0.071 |
| corrected_fdxt | pair_356 | 0.004 | 1943.314 | 2.349 | 2.417 | 4.783 | 4.817 | 0.060 |
| corrected_fdxt | pair_357 | 0.004 | 1943.314 | 2.388 | 2.438 | 4.861 | 4.746 | 0.206 |
| corrected_fdxt | pair_358 | 0.004 | 1943.314 | 2.612 | 2.383 | 5.053 | 4.783 | 0.551 |
| corrected_fdxt | pair_359 | 0.004 | 1943.314 | 2.434 | 4.313 | 6.768 | 4.748 | 3.535 |
| corrected_fdxt | pair_360 | 0.004 | 1943.314 | 2.582 | 2.435 | 5.038 | 4.790 | 0.540 |
| corrected_fdxt | pair_361 | 0.004 | 1943.314 | 2.422 | 2.408 | 4.847 | 4.704 | 0.277 |
| corrected_fdxt | pair_362 | 0.004 | 1943.314 | 4.017 | 3.939 | 7.973 | 4.669 | 5.737 |
| corrected_fdxt | pair_363 | 0.004 | 1943.314 | 2.346 | 3.324 | 5.687 | 4.678 | 1.751 |
| corrected_fdxt | pair_364 | 0.004 | 1943.314 | 2.528 | 2.810 | 5.357 | 4.743 | 1.095 |
| corrected_fdxt | pair_365 | 0.004 | 1943.314 | 2.502 | 3.252 | 5.774 | 5.686 | 1.129 |
| corrected_fdxt | pair_366 | 0.004 | 1943.314 | 4.427 | 2.336 | 6.781 | 5.085 | 3.316 |
| corrected_fdxt | pair_367 | 0.004 | 1943.314 | 2.451 | 2.422 | 4.892 | 4.738 | 0.308 |
| corrected_fdxt | pair_368 | 0.004 | 1943.314 | 2.358 | 2.506 | 4.884 | 4.709 | 0.341 |
| corrected_fdxt | pair_369 | 0.004 | 1943.314 | 2.461 | 2.336 | 4.815 | 4.791 | 0.183 |
| corrected_fdxt | pair_370 | 0.004 | 1943.314 | 5.838 | 5.645 | 11.542 | 11.099 | 0.928 |
| corrected_fdxt | pair_371 | 0.004 | 1943.314 | 8.913 | 8.946 | 18.049 | 17.400 | 1.281 |
| corrected_fdxt | pair_372 | 0.004 | 1943.314 | 5.556 | 5.464 | 11.076 | 10.948 | 0.248 |
| corrected_fdxt | pair_373 | 0.004 | 1943.314 | 8.684 | 8.838 | 17.644 | 17.276 | 0.778 |
| corrected_fdxt | pair_374 | 0.004 | 1943.314 | 4.057 | 4.062 | 8.154 | 7.821 | 0.614 |
| corrected_fdxt | pair_375 | 0.004 | 1943.314 | 12.190 | 8.877 | 21.187 | 18.938 | 5.442 |
| corrected_fdxt | pair_376 | 0.004 | 1943.314 | 8.772 | 8.649 | 17.540 | 17.239 | 0.554 |
| corrected_fdxt | pair_377 | 0.004 | 1943.314 | 3.876 | 3.860 | 7.773 | 7.757 | 0.037 |
| corrected_fdxt | pair_378 | 0.004 | 1943.314 | 3.966 | 3.929 | 7.928 | 7.914 | 0.131 |
| corrected_fdxt | pair_379 | 0.004 | 1943.314 | 2.367 | 2.434 | 4.817 | 4.742 | 0.198 |
| corrected_fdxt | pair_380 | 0.004 | 1943.314 | 4.657 | 4.692 | 9.394 | 9.400 | 0.018 |
| corrected_fdxt | pair_381 | 0.004 | 1943.314 | 5.472 | 5.411 | 10.947 | 10.964 | 0.036 |
| corrected_fdxt | pair_382 | 0.004 | 1943.314 | 5.485 | 5.426 | 10.970 | 10.947 | 0.100 |
| corrected_fdxt | pair_383 | 0.004 | 1943.314 | 5.460 | 5.481 | 11.003 | 11.007 | 0.013 |
| corrected_fdxt | pair_384 | 0.004 | 1943.314 | 3.972 | 3.846 | 7.851 | 7.860 | 0.061 |
| corrected_fdxt | pair_385 | 0.004 | 1943.314 | 5.492 | 5.504 | 11.059 | 11.020 | 0.155 |
| corrected_fdxt | pair_386 | 0.004 | 1943.314 | 5.461 | 5.447 | 10.970 | 10.940 | 0.071 |
| corrected_fdxt | pair_387 | 0.004 | 1943.314 | 4.061 | 3.896 | 7.992 | 7.816 | 0.310 |
| corrected_fdxt | pair_388 | 0.004 | 1943.314 | 3.938 | 3.890 | 7.860 | 7.846 | 0.031 |
| corrected_fdxt | pair_389 | 0.004 | 1943.314 | 2.365 | 2.336 | 4.716 | 4.694 | 0.068 |
| corrected_fdxt | pair_390 | 0.004 | 1943.314 | 4.728 | 4.695 | 9.468 | 9.484 | 0.129 |
| corrected_fdxt | pair_391 | 0.004 | 1943.314 | 5.496 | 5.660 | 11.214 | 11.039 | 0.333 |
| corrected_fdxt | pair_392 | 0.004 | 1943.314 | 16.391 | 16.632 | 33.406 | 31.582 | 3.174 |
| corrected_fdxt | pair_393 | 0.004 | 1943.314 | 4.018 | 3.962 | 8.017 | 7.908 | 0.269 |
| corrected_fdxt | pair_394 | 0.004 | 1943.314 | 10.363 | 9.601 | 20.209 | 19.903 | 1.087 |
| corrected_fdxt | pair_395 | 0.004 | 1943.314 | 25.077 | 18.707 | 44.269 | 36.292 | 16.674 |
| corrected_fdxt | pair_396 | 0.004 | 1943.314 | 4.053 | 4.691 | 8.792 | 8.685 | 1.053 |
| corrected_fdxt | pair_397 | 0.004 | 1943.314 | 4.252 | 4.520 | 8.824 | 8.119 | 1.552 |
| corrected_fdxt | pair_398 | 0.004 | 1943.314 | 2.659 | 2.424 | 5.102 | 4.872 | 0.531 |
| corrected_fdxt | pair_399 | 0.004 | 1943.314 | 4.908 | 4.733 | 9.691 | 9.533 | 0.412 |
| corrected_fdxt | pair_400 | 0.004 | 1943.314 | 5.621 | 5.543 | 11.220 | 11.165 | 0.228 |
| corrected_fdxt | pair_401 | 0.004 | 1943.314 | 3.941 | 4.120 | 8.211 | 8.230 | 0.086 |
| corrected_fdxt | pair_402 | 0.004 | 1943.314 | 9.340 | 5.553 | 14.952 | 11.462 | 6.484 |
| corrected_fdxt | pair_403 | 0.004 | 1943.314 | 6.090 | 5.572 | 11.733 | 11.143 | 1.165 |
| corrected_fdxt | pair_404 | 0.004 | 1943.314 | 3.976 | 3.897 | 7.909 | 7.913 | 0.140 |
| corrected_fdxt | pair_405 | 0.004 | 1943.314 | 3.873 | 3.923 | 7.829 | 7.829 | 0.070 |
| corrected_fdxt | pair_406 | 0.004 | 1943.314 | 2.376 | 2.342 | 4.734 | 4.688 | 0.082 |
| corrected_fdxt | pair_407 | 0.004 | 1943.314 | 4.673 | 4.660 | 9.380 | 9.367 | 0.068 |
| corrected_fdxt | pair_408 | 0.004 | 1943.314 | 3.906 | 3.866 | 7.805 | 7.812 | 0.047 |
| corrected_fdxt | pair_409 | 0.004 | 1943.314 | 9.358 | 9.341 | 18.840 | 18.797 | 0.111 |
| corrected_fdxt | pair_410 | 0.004 | 1943.314 | 15.591 | 15.512 | 31.454 | 31.402 | 0.138 |
| corrected_fdxt | pair_411 | 0.004 | 1943.314 | 3.906 | 3.847 | 7.785 | 7.784 | 0.029 |
| corrected_fdxt | pair_412 | 0.004 | 1943.314 | 3.946 | 3.898 | 7.877 | 7.887 | 0.063 |
| corrected_fdxt | pair_413 | 0.004 | 1943.314 | 2.333 | 2.303 | 4.652 | 4.667 | 0.031 |
| corrected_fdxt | pair_414 | 0.004 | 1943.314 | 4.690 | 5.559 | 10.294 | 9.416 | 1.535 |
| corrected_fdxt | pair_415 | 0.004 | 1943.314 | 4.246 | 4.006 | 8.298 | 8.259 | 0.477 |
| corrected_fdxt | pair_416 | 0.004 | 1943.314 | 3.884 | 4.041 | 7.960 | 7.884 | 0.221 |
| corrected_fdxt | pair_417 | 0.004 | 1943.314 | 3.880 | 3.822 | 7.734 | 7.738 | 0.016 |
| corrected_fdxt | pair_418 | 0.004 | 1943.314 | 4.161 | 4.158 | 8.354 | 7.879 | 0.891 |
| corrected_fdxt | pair_419 | 0.004 | 1943.314 | 2.358 | 2.316 | 4.690 | 4.639 | 0.112 |
| corrected_fdxt | pair_420 | 0.004 | 1943.314 | 3.876 | 3.847 | 7.778 | 7.769 | 0.099 |
| corrected_fdxt | pair_421 | 0.004 | 1943.314 | 9.595 | 9.365 | 19.109 | 18.892 | 0.498 |
| corrected_fdxt | pair_422 | 0.004 | 1943.314 | 3.899 | 3.845 | 7.781 | 7.785 | 0.023 |
| corrected_fdxt | pair_423 | 0.004 | 1943.314 | 3.890 | 3.876 | 7.799 | 7.848 | 0.085 |
| corrected_fdxt | pair_424 | 0.004 | 1943.314 | 2.367 | 2.366 | 4.748 | 4.711 | 0.113 |
| corrected_fdxt | pair_425 | 0.004 | 1943.314 | 4.665 | 4.673 | 9.385 | 9.378 | 0.071 |
| corrected_fdxt | pair_426 | 0.004 | 1943.314 | 4.034 | 3.894 | 7.963 | 7.897 | 0.164 |
| corrected_fdxt | pair_427 | 0.004 | 1943.314 | 3.926 | 3.855 | 7.814 | 7.784 | 0.057 |
| corrected_fdxt | pair_428 | 0.004 | 1943.314 | 2.324 | 2.345 | 4.685 | 4.675 | 0.043 |
| corrected_fdxt | pair_429 | 0.004 | 1943.314 | 4.655 | 4.642 | 9.340 | 9.338 | 0.033 |
| corrected_fdxt | pair_430 | 0.004 | 1943.314 | 3.913 | 3.900 | 7.846 | 7.853 | 0.022 |
| corrected_fdxt | pair_431 | 0.004 | 1943.314 | 2.319 | 2.334 | 4.668 | 4.646 | 0.046 |
| corrected_fdxt | pair_432 | 0.004 | 1943.314 | 3.891 | 3.891 | 7.816 | 7.764 | 0.091 |
| corrected_fdxt | pair_433 | 0.004 | 1943.314 | 2.353 | 2.339 | 4.708 | 4.727 | 0.042 |
| corrected_fdxt | pair_434 | 0.004 | 1943.314 | 3.917 | 3.877 | 7.826 | 7.813 | 0.036 |
| corrected_fdxt | pair_435 | 0.004 | 1943.314 | 2.350 | 2.312 | 4.677 | 4.659 | 0.054 |

## Table 3 — Communication and storage

| Variant | Workload | Mean communication bytes | Mean storage bytes |
|---|---|---:|---:|
| c1_hybrid | pair_001 | 6303.0 | 586970.0 |
| c1_hybrid | pair_002 | 3870.0 | 586970.0 |
| c1_hybrid | pair_003 | 4191.0 | 586970.0 |
| c1_hybrid | pair_004 | 4833.0 | 586970.0 |
| c1_hybrid | pair_005 | 2265.0 | 586970.0 |
| c1_hybrid | pair_006 | 3870.0 | 586970.0 |
| c1_hybrid | pair_007 | 8043.0 | 586970.0 |
| c1_hybrid | pair_008 | 4191.0 | 586970.0 |
| c1_hybrid | pair_009 | 9905.0 | 586970.0 |
| c1_hybrid | pair_010 | 4191.0 | 586970.0 |
| c1_hybrid | pair_011 | 3549.0 | 586970.0 |
| c1_hybrid | pair_012 | 3549.0 | 586970.0 |
| c1_hybrid | pair_013 | 3870.0 | 586970.0 |
| c1_hybrid | pair_014 | 3870.0 | 586970.0 |
| c1_hybrid | pair_015 | 3870.0 | 586970.0 |
| c1_hybrid | pair_016 | 2265.0 | 586970.0 |
| c1_hybrid | pair_017 | 2265.0 | 586970.0 |
| c1_hybrid | pair_018 | 6117.0 | 586970.0 |
| c1_hybrid | pair_019 | 4191.0 | 586970.0 |
| c1_hybrid | pair_020 | 10211.0 | 586970.0 |
| c1_hybrid | pair_021 | 4191.0 | 586970.0 |
| c1_hybrid | pair_022 | 9648.0 | 586970.0 |
| c1_hybrid | pair_023 | 3549.0 | 586970.0 |
| c1_hybrid | pair_024 | 6438.0 | 586970.0 |
| c1_hybrid | pair_025 | 9969.0 | 586970.0 |
| c1_hybrid | pair_026 | 3549.0 | 586970.0 |
| c1_hybrid | pair_027 | 3549.0 | 586970.0 |
| c1_hybrid | pair_028 | 2265.0 | 586970.0 |
| c1_hybrid | pair_029 | 3870.0 | 586970.0 |
| c1_hybrid | pair_030 | 3411.0 | 586970.0 |
| c1_hybrid | pair_031 | 3659.0 | 586970.0 |
| c1_hybrid | pair_032 | 4401.0 | 586970.0 |
| c1_hybrid | pair_033 | 2031.0 | 586970.0 |
| c1_hybrid | pair_034 | 3411.0 | 586970.0 |
| c1_hybrid | pair_035 | 4305.0 | 586970.0 |
| c1_hybrid | pair_036 | 3659.0 | 586970.0 |
| c1_hybrid | pair_037 | 4241.0 | 586970.0 |
| c1_hybrid | pair_038 | 3659.0 | 586970.0 |
| c1_hybrid | pair_039 | 3163.0 | 586970.0 |
| c1_hybrid | pair_040 | 3163.0 | 586970.0 |
| c1_hybrid | pair_041 | 3411.0 | 586970.0 |
| c1_hybrid | pair_042 | 3411.0 | 586970.0 |
| c1_hybrid | pair_043 | 3411.0 | 586970.0 |
| c1_hybrid | pair_044 | 2031.0 | 586970.0 |
| c1_hybrid | pair_045 | 2031.0 | 586970.0 |
| c1_hybrid | pair_046 | 4289.0 | 586970.0 |
| c1_hybrid | pair_047 | 3659.0 | 586970.0 |
| c1_hybrid | pair_048 | 4289.0 | 586970.0 |
| c1_hybrid | pair_049 | 3659.0 | 586970.0 |
| c1_hybrid | pair_050 | 4305.0 | 586970.0 |
| c1_hybrid | pair_051 | 3163.0 | 586970.0 |
| c1_hybrid | pair_052 | 4305.0 | 586970.0 |
| c1_hybrid | pair_053 | 4273.0 | 586970.0 |
| c1_hybrid | pair_054 | 3163.0 | 586970.0 |
| c1_hybrid | pair_055 | 3163.0 | 586970.0 |
| c1_hybrid | pair_056 | 2031.0 | 586970.0 |
| c1_hybrid | pair_057 | 3411.0 | 586970.0 |
| c1_hybrid | pair_058 | 3450.0 | 586970.0 |
| c1_hybrid | pair_059 | 3554.0 | 586970.0 |
| c1_hybrid | pair_060 | 2095.0 | 586970.0 |
| c1_hybrid | pair_061 | 3437.0 | 586970.0 |
| c1_hybrid | pair_062 | 3476.0 | 586970.0 |
| c1_hybrid | pair_063 | 3502.0 | 586970.0 |
| c1_hybrid | pair_064 | 3424.0 | 586970.0 |
| c1_hybrid | pair_065 | 3476.0 | 586970.0 |
| c1_hybrid | pair_066 | 3259.0 | 586970.0 |
| c1_hybrid | pair_067 | 3259.0 | 586970.0 |
| c1_hybrid | pair_068 | 3476.0 | 586970.0 |
| c1_hybrid | pair_069 | 3476.0 | 586970.0 |
| c1_hybrid | pair_070 | 3489.0 | 586970.0 |
| c1_hybrid | pair_071 | 2095.0 | 586970.0 |
| c1_hybrid | pair_072 | 2095.0 | 586970.0 |
| c1_hybrid | pair_073 | 3463.0 | 586970.0 |
| c1_hybrid | pair_074 | 3502.0 | 586970.0 |
| c1_hybrid | pair_075 | 3463.0 | 586970.0 |
| c1_hybrid | pair_076 | 3476.0 | 586970.0 |
| c1_hybrid | pair_077 | 3476.0 | 586970.0 |
| c1_hybrid | pair_078 | 3259.0 | 586970.0 |
| c1_hybrid | pair_079 | 3476.0 | 586970.0 |
| c1_hybrid | pair_080 | 3450.0 | 586970.0 |
| c1_hybrid | pair_081 | 3259.0 | 586970.0 |
| c1_hybrid | pair_082 | 3259.0 | 586970.0 |
| c1_hybrid | pair_083 | 2095.0 | 586970.0 |
| c1_hybrid | pair_084 | 3411.0 | 586970.0 |
| c1_hybrid | pair_085 | 3813.0 | 586970.0 |
| c1_hybrid | pair_086 | 2055.0 | 586970.0 |
| c1_hybrid | pair_087 | 3450.0 | 586970.0 |
| c1_hybrid | pair_088 | 3729.0 | 586970.0 |
| c1_hybrid | pair_089 | 3757.0 | 586970.0 |
| c1_hybrid | pair_090 | 3673.0 | 586970.0 |
| c1_hybrid | pair_091 | 3729.0 | 586970.0 |
| c1_hybrid | pair_092 | 3199.0 | 586970.0 |
| c1_hybrid | pair_093 | 3199.0 | 586970.0 |
| c1_hybrid | pair_094 | 3450.0 | 586970.0 |
| c1_hybrid | pair_095 | 3450.0 | 586970.0 |
| c1_hybrid | pair_096 | 3450.0 | 586970.0 |
| c1_hybrid | pair_097 | 2055.0 | 586970.0 |
| c1_hybrid | pair_098 | 2055.0 | 586970.0 |
| c1_hybrid | pair_099 | 3715.0 | 586970.0 |
| c1_hybrid | pair_100 | 3757.0 | 586970.0 |
| c1_hybrid | pair_101 | 3715.0 | 586970.0 |
| c1_hybrid | pair_102 | 3729.0 | 586970.0 |
| c1_hybrid | pair_103 | 3729.0 | 586970.0 |
| c1_hybrid | pair_104 | 3199.0 | 586970.0 |
| c1_hybrid | pair_105 | 3729.0 | 586970.0 |
| c1_hybrid | pair_106 | 3701.0 | 586970.0 |
| c1_hybrid | pair_107 | 3199.0 | 586970.0 |
| c1_hybrid | pair_108 | 3199.0 | 586970.0 |
| c1_hybrid | pair_109 | 2055.0 | 586970.0 |
| c1_hybrid | pair_110 | 3450.0 | 586970.0 |
| c1_hybrid | pair_111 | 2119.0 | 586970.0 |
| c1_hybrid | pair_112 | 3554.0 | 586970.0 |
| c1_hybrid | pair_113 | 4305.0 | 586970.0 |
| c1_hybrid | pair_114 | 3813.0 | 586970.0 |
| c1_hybrid | pair_115 | 4241.0 | 586970.0 |
| c1_hybrid | pair_116 | 3813.0 | 586970.0 |
| c1_hybrid | pair_117 | 3295.0 | 586970.0 |
| c1_hybrid | pair_118 | 3295.0 | 586970.0 |
| c1_hybrid | pair_119 | 3554.0 | 586970.0 |
| c1_hybrid | pair_120 | 3554.0 | 586970.0 |
| c1_hybrid | pair_121 | 3554.0 | 586970.0 |
| c1_hybrid | pair_122 | 2119.0 | 586970.0 |
| c1_hybrid | pair_123 | 2119.0 | 586970.0 |
| c1_hybrid | pair_124 | 4289.0 | 586970.0 |
| c1_hybrid | pair_125 | 3813.0 | 586970.0 |
| c1_hybrid | pair_126 | 4289.0 | 586970.0 |
| c1_hybrid | pair_127 | 3813.0 | 586970.0 |
| c1_hybrid | pair_128 | 4305.0 | 586970.0 |
| c1_hybrid | pair_129 | 3295.0 | 586970.0 |
| c1_hybrid | pair_130 | 4305.0 | 586970.0 |
| c1_hybrid | pair_131 | 4273.0 | 586970.0 |
| c1_hybrid | pair_132 | 3295.0 | 586970.0 |
| c1_hybrid | pair_133 | 3295.0 | 586970.0 |
| c1_hybrid | pair_134 | 2119.0 | 586970.0 |
| c1_hybrid | pair_135 | 3554.0 | 586970.0 |
| c1_hybrid | pair_136 | 2047.0 | 586970.0 |
| c1_hybrid | pair_137 | 2071.0 | 586970.0 |
| c1_hybrid | pair_138 | 2087.0 | 586970.0 |
| c1_hybrid | pair_139 | 2039.0 | 586970.0 |
| c1_hybrid | pair_140 | 2071.0 | 586970.0 |
| c1_hybrid | pair_141 | 2047.0 | 586970.0 |
| c1_hybrid | pair_142 | 2087.0 | 586970.0 |
| c1_hybrid | pair_143 | 2071.0 | 586970.0 |
| c1_hybrid | pair_144 | 2071.0 | 586970.0 |
| c1_hybrid | pair_145 | 2079.0 | 586970.0 |
| c1_hybrid | pair_146 | 2071.0 | 586970.0 |
| c1_hybrid | pair_147 | 2063.0 | 586970.0 |
| c1_hybrid | pair_148 | 2063.0 | 586970.0 |
| c1_hybrid | pair_149 | 2087.0 | 586970.0 |
| c1_hybrid | pair_150 | 2063.0 | 586970.0 |
| c1_hybrid | pair_151 | 2071.0 | 586970.0 |
| c1_hybrid | pair_152 | 2071.0 | 586970.0 |
| c1_hybrid | pair_153 | 2055.0 | 586970.0 |
| c1_hybrid | pair_154 | 2071.0 | 586970.0 |
| c1_hybrid | pair_155 | 2055.0 | 586970.0 |
| c1_hybrid | pair_156 | 2055.0 | 586970.0 |
| c1_hybrid | pair_157 | 2063.0 | 586970.0 |
| c1_hybrid | pair_158 | 2111.0 | 586970.0 |
| c1_hybrid | pair_159 | 2031.0 | 586970.0 |
| c1_hybrid | pair_160 | 3476.0 | 586970.0 |
| c1_hybrid | pair_161 | 3502.0 | 586970.0 |
| c1_hybrid | pair_162 | 3424.0 | 586970.0 |
| c1_hybrid | pair_163 | 3476.0 | 586970.0 |
| c1_hybrid | pair_164 | 3187.0 | 586970.0 |
| c1_hybrid | pair_165 | 3187.0 | 586970.0 |
| c1_hybrid | pair_166 | 3476.0 | 586970.0 |
| c1_hybrid | pair_167 | 3476.0 | 586970.0 |
| c1_hybrid | pair_168 | 3489.0 | 586970.0 |
| c1_hybrid | pair_169 | 2047.0 | 586970.0 |
| c1_hybrid | pair_170 | 2047.0 | 586970.0 |
| c1_hybrid | pair_171 | 3463.0 | 586970.0 |
| c1_hybrid | pair_172 | 3502.0 | 586970.0 |
| c1_hybrid | pair_173 | 3463.0 | 586970.0 |
| c1_hybrid | pair_174 | 3476.0 | 586970.0 |
| c1_hybrid | pair_175 | 3476.0 | 586970.0 |
| c1_hybrid | pair_176 | 3187.0 | 586970.0 |
| c1_hybrid | pair_177 | 3476.0 | 586970.0 |
| c1_hybrid | pair_178 | 3450.0 | 586970.0 |
| c1_hybrid | pair_179 | 3187.0 | 586970.0 |
| c1_hybrid | pair_180 | 3187.0 | 586970.0 |
| c1_hybrid | pair_181 | 2047.0 | 586970.0 |
| c1_hybrid | pair_182 | 3411.0 | 586970.0 |
| c1_hybrid | pair_183 | 3729.0 | 586970.0 |
| c1_hybrid | pair_184 | 6871.0 | 586970.0 |
| c1_hybrid | pair_185 | 3729.0 | 586970.0 |
| c1_hybrid | pair_186 | 3223.0 | 586970.0 |
| c1_hybrid | pair_187 | 3223.0 | 586970.0 |
| c1_hybrid | pair_188 | 3476.0 | 586970.0 |
| c1_hybrid | pair_189 | 3476.0 | 586970.0 |
| c1_hybrid | pair_190 | 3476.0 | 586970.0 |
| c1_hybrid | pair_191 | 2071.0 | 586970.0 |
| c1_hybrid | pair_192 | 2071.0 | 586970.0 |
| c1_hybrid | pair_193 | 5387.0 | 586970.0 |
| c1_hybrid | pair_194 | 3729.0 | 586970.0 |
| c1_hybrid | pair_195 | 6949.0 | 586970.0 |
| c1_hybrid | pair_196 | 3729.0 | 586970.0 |
| c1_hybrid | pair_197 | 6975.0 | 586970.0 |
| c1_hybrid | pair_198 | 3223.0 | 586970.0 |
| c1_hybrid | pair_199 | 5640.0 | 586970.0 |
| c1_hybrid | pair_200 | 6923.0 | 586970.0 |
| c1_hybrid | pair_201 | 3223.0 | 586970.0 |
| c1_hybrid | pair_202 | 3223.0 | 586970.0 |
| c1_hybrid | pair_203 | 2071.0 | 586970.0 |
| c1_hybrid | pair_204 | 3476.0 | 586970.0 |
| c1_hybrid | pair_205 | 3673.0 | 586970.0 |
| c1_hybrid | pair_206 | 3729.0 | 586970.0 |
| c1_hybrid | pair_207 | 3247.0 | 586970.0 |
| c1_hybrid | pair_208 | 3247.0 | 586970.0 |
| c1_hybrid | pair_209 | 3502.0 | 586970.0 |
| c1_hybrid | pair_210 | 3502.0 | 586970.0 |
| c1_hybrid | pair_211 | 3502.0 | 586970.0 |
| c1_hybrid | pair_212 | 2087.0 | 586970.0 |
| c1_hybrid | pair_213 | 2087.0 | 586970.0 |
| c1_hybrid | pair_214 | 3715.0 | 586970.0 |
| c1_hybrid | pair_215 | 3757.0 | 586970.0 |
| c1_hybrid | pair_216 | 3715.0 | 586970.0 |
| c1_hybrid | pair_217 | 3729.0 | 586970.0 |
| c1_hybrid | pair_218 | 3729.0 | 586970.0 |
| c1_hybrid | pair_219 | 3247.0 | 586970.0 |
| c1_hybrid | pair_220 | 3729.0 | 586970.0 |
| c1_hybrid | pair_221 | 3701.0 | 586970.0 |
| c1_hybrid | pair_222 | 3247.0 | 586970.0 |
| c1_hybrid | pair_223 | 3247.0 | 586970.0 |
| c1_hybrid | pair_224 | 2087.0 | 586970.0 |
| c1_hybrid | pair_225 | 3502.0 | 586970.0 |
| c1_hybrid | pair_226 | 3673.0 | 586970.0 |
| c1_hybrid | pair_227 | 3175.0 | 586970.0 |
| c1_hybrid | pair_228 | 3175.0 | 586970.0 |
| c1_hybrid | pair_229 | 3424.0 | 586970.0 |
| c1_hybrid | pair_230 | 3424.0 | 586970.0 |
| c1_hybrid | pair_231 | 3424.0 | 586970.0 |
| c1_hybrid | pair_232 | 2039.0 | 586970.0 |
| c1_hybrid | pair_233 | 2039.0 | 586970.0 |
| c1_hybrid | pair_234 | 5307.0 | 586970.0 |
| c1_hybrid | pair_235 | 3673.0 | 586970.0 |
| c1_hybrid | pair_236 | 8531.0 | 586970.0 |
| c1_hybrid | pair_237 | 3673.0 | 586970.0 |
| c1_hybrid | pair_238 | 8186.0 | 586970.0 |
| c1_hybrid | pair_239 | 3175.0 | 586970.0 |
| c1_hybrid | pair_240 | 5556.0 | 586970.0 |
| c1_hybrid | pair_241 | 8499.0 | 586970.0 |
| c1_hybrid | pair_242 | 3175.0 | 586970.0 |
| c1_hybrid | pair_243 | 3175.0 | 586970.0 |
| c1_hybrid | pair_244 | 2039.0 | 586970.0 |
| c1_hybrid | pair_245 | 3424.0 | 586970.0 |
| c1_hybrid | pair_246 | 3223.0 | 586970.0 |
| c1_hybrid | pair_247 | 3223.0 | 586970.0 |
| c1_hybrid | pair_248 | 3476.0 | 586970.0 |
| c1_hybrid | pair_249 | 3476.0 | 586970.0 |
| c1_hybrid | pair_250 | 3476.0 | 586970.0 |
| c1_hybrid | pair_251 | 2071.0 | 586970.0 |
| c1_hybrid | pair_252 | 2071.0 | 586970.0 |
| c1_hybrid | pair_253 | 3715.0 | 586970.0 |
| c1_hybrid | pair_254 | 3757.0 | 586970.0 |
| c1_hybrid | pair_255 | 3715.0 | 586970.0 |
| c1_hybrid | pair_256 | 3729.0 | 586970.0 |
| c1_hybrid | pair_257 | 3729.0 | 586970.0 |
| c1_hybrid | pair_258 | 3223.0 | 586970.0 |
| c1_hybrid | pair_259 | 3729.0 | 586970.0 |
| c1_hybrid | pair_260 | 3701.0 | 586970.0 |
| c1_hybrid | pair_261 | 3223.0 | 586970.0 |
| c1_hybrid | pair_262 | 3223.0 | 586970.0 |
| c1_hybrid | pair_263 | 2071.0 | 586970.0 |
| c1_hybrid | pair_264 | 3476.0 | 586970.0 |
| c1_hybrid | pair_265 | 3247.0 | 586970.0 |
| c1_hybrid | pair_266 | 3223.0 | 586970.0 |
| c1_hybrid | pair_267 | 3223.0 | 586970.0 |
| c1_hybrid | pair_268 | 3235.0 | 586970.0 |
| c1_hybrid | pair_269 | 2047.0 | 586970.0 |
| c1_hybrid | pair_270 | 2047.0 | 586970.0 |
| c1_hybrid | pair_271 | 3211.0 | 586970.0 |
| c1_hybrid | pair_272 | 3247.0 | 586970.0 |
| c1_hybrid | pair_273 | 3211.0 | 586970.0 |
| c1_hybrid | pair_274 | 3223.0 | 586970.0 |
| c1_hybrid | pair_275 | 3223.0 | 586970.0 |
| c1_hybrid | pair_276 | 3199.0 | 586970.0 |
| c1_hybrid | pair_277 | 3223.0 | 586970.0 |
| c1_hybrid | pair_278 | 3199.0 | 586970.0 |
| c1_hybrid | pair_279 | 3199.0 | 586970.0 |
| c1_hybrid | pair_280 | 3211.0 | 586970.0 |
| c1_hybrid | pair_281 | 2047.0 | 586970.0 |
| c1_hybrid | pair_282 | 3163.0 | 586970.0 |
| c1_hybrid | pair_283 | 3223.0 | 586970.0 |
| c1_hybrid | pair_284 | 3223.0 | 586970.0 |
| c1_hybrid | pair_285 | 3235.0 | 586970.0 |
| c1_hybrid | pair_286 | 2087.0 | 586970.0 |
| c1_hybrid | pair_287 | 2087.0 | 586970.0 |
| c1_hybrid | pair_288 | 3211.0 | 586970.0 |
| c1_hybrid | pair_289 | 3247.0 | 586970.0 |
| c1_hybrid | pair_290 | 3211.0 | 586970.0 |
| c1_hybrid | pair_291 | 3223.0 | 586970.0 |
| c1_hybrid | pair_292 | 3223.0 | 586970.0 |
| c1_hybrid | pair_293 | 3199.0 | 586970.0 |
| c1_hybrid | pair_294 | 3223.0 | 586970.0 |
| c1_hybrid | pair_295 | 3199.0 | 586970.0 |
| c1_hybrid | pair_296 | 3199.0 | 586970.0 |
| c1_hybrid | pair_297 | 3211.0 | 586970.0 |
| c1_hybrid | pair_298 | 2087.0 | 586970.0 |
| c1_hybrid | pair_299 | 3163.0 | 586970.0 |
| c1_hybrid | pair_300 | 3476.0 | 586970.0 |
| c1_hybrid | pair_301 | 3489.0 | 586970.0 |
| c1_hybrid | pair_302 | 2071.0 | 586970.0 |
| c1_hybrid | pair_303 | 2071.0 | 586970.0 |
| c1_hybrid | pair_304 | 3463.0 | 586970.0 |
| c1_hybrid | pair_305 | 3502.0 | 586970.0 |
| c1_hybrid | pair_306 | 3463.0 | 586970.0 |
| c1_hybrid | pair_307 | 3476.0 | 586970.0 |
| c1_hybrid | pair_308 | 3476.0 | 586970.0 |
| c1_hybrid | pair_309 | 3223.0 | 586970.0 |
| c1_hybrid | pair_310 | 3476.0 | 586970.0 |
| c1_hybrid | pair_311 | 3450.0 | 586970.0 |
| c1_hybrid | pair_312 | 3223.0 | 586970.0 |
| c1_hybrid | pair_313 | 3223.0 | 586970.0 |
| c1_hybrid | pair_314 | 2071.0 | 586970.0 |
| c1_hybrid | pair_315 | 3411.0 | 586970.0 |
| c1_hybrid | pair_316 | 3489.0 | 586970.0 |
| c1_hybrid | pair_317 | 2071.0 | 586970.0 |
| c1_hybrid | pair_318 | 2071.0 | 586970.0 |
| c1_hybrid | pair_319 | 3463.0 | 586970.0 |
| c1_hybrid | pair_320 | 3502.0 | 586970.0 |
| c1_hybrid | pair_321 | 3463.0 | 586970.0 |
| c1_hybrid | pair_322 | 3476.0 | 586970.0 |
| c1_hybrid | pair_323 | 3476.0 | 586970.0 |
| c1_hybrid | pair_324 | 3223.0 | 586970.0 |
| c1_hybrid | pair_325 | 3476.0 | 586970.0 |
| c1_hybrid | pair_326 | 3450.0 | 586970.0 |
| c1_hybrid | pair_327 | 3223.0 | 586970.0 |
| c1_hybrid | pair_328 | 3223.0 | 586970.0 |
| c1_hybrid | pair_329 | 2071.0 | 586970.0 |
| c1_hybrid | pair_330 | 3411.0 | 586970.0 |
| c1_hybrid | pair_331 | 2079.0 | 586970.0 |
| c1_hybrid | pair_332 | 2079.0 | 586970.0 |
| c1_hybrid | pair_333 | 3463.0 | 586970.0 |
| c1_hybrid | pair_334 | 3502.0 | 586970.0 |
| c1_hybrid | pair_335 | 3463.0 | 586970.0 |
| c1_hybrid | pair_336 | 3476.0 | 586970.0 |
| c1_hybrid | pair_337 | 3476.0 | 586970.0 |
| c1_hybrid | pair_338 | 3235.0 | 586970.0 |
| c1_hybrid | pair_339 | 3476.0 | 586970.0 |
| c1_hybrid | pair_340 | 3450.0 | 586970.0 |
| c1_hybrid | pair_341 | 3235.0 | 586970.0 |
| c1_hybrid | pair_342 | 3235.0 | 586970.0 |
| c1_hybrid | pair_343 | 2079.0 | 586970.0 |
| c1_hybrid | pair_344 | 3411.0 | 586970.0 |
| c1_hybrid | pair_345 | 2063.0 | 586970.0 |
| c1_hybrid | pair_346 | 2063.0 | 586970.0 |
| c1_hybrid | pair_347 | 2087.0 | 586970.0 |
| c1_hybrid | pair_348 | 2063.0 | 586970.0 |
| c1_hybrid | pair_349 | 2071.0 | 586970.0 |
| c1_hybrid | pair_350 | 2071.0 | 586970.0 |
| c1_hybrid | pair_351 | 2055.0 | 586970.0 |
| c1_hybrid | pair_352 | 2071.0 | 586970.0 |
| c1_hybrid | pair_353 | 2055.0 | 586970.0 |
| c1_hybrid | pair_354 | 2055.0 | 586970.0 |
| c1_hybrid | pair_355 | 2063.0 | 586970.0 |
| c1_hybrid | pair_356 | 2111.0 | 586970.0 |
| c1_hybrid | pair_357 | 2031.0 | 586970.0 |
| c1_hybrid | pair_358 | 2063.0 | 586970.0 |
| c1_hybrid | pair_359 | 2087.0 | 586970.0 |
| c1_hybrid | pair_360 | 2063.0 | 586970.0 |
| c1_hybrid | pair_361 | 2071.0 | 586970.0 |
| c1_hybrid | pair_362 | 2071.0 | 586970.0 |
| c1_hybrid | pair_363 | 2055.0 | 586970.0 |
| c1_hybrid | pair_364 | 2071.0 | 586970.0 |
| c1_hybrid | pair_365 | 2055.0 | 586970.0 |
| c1_hybrid | pair_366 | 2055.0 | 586970.0 |
| c1_hybrid | pair_367 | 2063.0 | 586970.0 |
| c1_hybrid | pair_368 | 2111.0 | 586970.0 |
| c1_hybrid | pair_369 | 2031.0 | 586970.0 |
| c1_hybrid | pair_370 | 3715.0 | 586970.0 |
| c1_hybrid | pair_371 | 5367.0 | 586970.0 |
| c1_hybrid | pair_372 | 3715.0 | 586970.0 |
| c1_hybrid | pair_373 | 5387.0 | 586970.0 |
| c1_hybrid | pair_374 | 3211.0 | 586970.0 |
| c1_hybrid | pair_375 | 5387.0 | 586970.0 |
| c1_hybrid | pair_376 | 5347.0 | 586970.0 |
| c1_hybrid | pair_377 | 3211.0 | 586970.0 |
| c1_hybrid | pair_378 | 3211.0 | 586970.0 |
| c1_hybrid | pair_379 | 2063.0 | 586970.0 |
| c1_hybrid | pair_380 | 3463.0 | 586970.0 |
| c1_hybrid | pair_381 | 3715.0 | 586970.0 |
| c1_hybrid | pair_382 | 3729.0 | 586970.0 |
| c1_hybrid | pair_383 | 3729.0 | 586970.0 |
| c1_hybrid | pair_384 | 3247.0 | 586970.0 |
| c1_hybrid | pair_385 | 3729.0 | 586970.0 |
| c1_hybrid | pair_386 | 3701.0 | 586970.0 |
| c1_hybrid | pair_387 | 3247.0 | 586970.0 |
| c1_hybrid | pair_388 | 3247.0 | 586970.0 |
| c1_hybrid | pair_389 | 2087.0 | 586970.0 |
| c1_hybrid | pair_390 | 3502.0 | 586970.0 |
| c1_hybrid | pair_391 | 3715.0 | 586970.0 |
| c1_hybrid | pair_392 | 8279.0 | 586970.0 |
| c1_hybrid | pair_393 | 3211.0 | 586970.0 |
| c1_hybrid | pair_394 | 5619.0 | 586970.0 |
| c1_hybrid | pair_395 | 8531.0 | 586970.0 |
| c1_hybrid | pair_396 | 3211.0 | 586970.0 |
| c1_hybrid | pair_397 | 3211.0 | 586970.0 |
| c1_hybrid | pair_398 | 2063.0 | 586970.0 |
| c1_hybrid | pair_399 | 3463.0 | 586970.0 |
| c1_hybrid | pair_400 | 3729.0 | 586970.0 |
| c1_hybrid | pair_401 | 3223.0 | 586970.0 |
| c1_hybrid | pair_402 | 3729.0 | 586970.0 |
| c1_hybrid | pair_403 | 3701.0 | 586970.0 |
| c1_hybrid | pair_404 | 3223.0 | 586970.0 |
| c1_hybrid | pair_405 | 3223.0 | 586970.0 |
| c1_hybrid | pair_406 | 2071.0 | 586970.0 |
| c1_hybrid | pair_407 | 3476.0 | 586970.0 |
| c1_hybrid | pair_408 | 3223.0 | 586970.0 |
| c1_hybrid | pair_409 | 5640.0 | 586970.0 |
| c1_hybrid | pair_410 | 8248.0 | 586970.0 |
| c1_hybrid | pair_411 | 3223.0 | 586970.0 |
| c1_hybrid | pair_412 | 3223.0 | 586970.0 |
| c1_hybrid | pair_413 | 2071.0 | 586970.0 |
| c1_hybrid | pair_414 | 3476.0 | 586970.0 |
| c1_hybrid | pair_415 | 3223.0 | 586970.0 |
| c1_hybrid | pair_416 | 3199.0 | 586970.0 |
| c1_hybrid | pair_417 | 3199.0 | 586970.0 |
| c1_hybrid | pair_418 | 3211.0 | 586970.0 |
| c1_hybrid | pair_419 | 2055.0 | 586970.0 |
| c1_hybrid | pair_420 | 3163.0 | 586970.0 |
| c1_hybrid | pair_421 | 5598.0 | 586970.0 |
| c1_hybrid | pair_422 | 3223.0 | 586970.0 |
| c1_hybrid | pair_423 | 3223.0 | 586970.0 |
| c1_hybrid | pair_424 | 2071.0 | 586970.0 |
| c1_hybrid | pair_425 | 3476.0 | 586970.0 |
| c1_hybrid | pair_426 | 3199.0 | 586970.0 |
| c1_hybrid | pair_427 | 3199.0 | 586970.0 |
| c1_hybrid | pair_428 | 2055.0 | 586970.0 |
| c1_hybrid | pair_429 | 3450.0 | 586970.0 |
| c1_hybrid | pair_430 | 3211.0 | 586970.0 |
| c1_hybrid | pair_431 | 2055.0 | 586970.0 |
| c1_hybrid | pair_432 | 3163.0 | 586970.0 |
| c1_hybrid | pair_433 | 2063.0 | 586970.0 |
| c1_hybrid | pair_434 | 3163.0 | 586970.0 |
| c1_hybrid | pair_435 | 2031.0 | 586970.0 |
| c1_log | pair_001 | 5340.0 | 586970.0 |
| c1_log | pair_002 | 2907.0 | 586970.0 |
| c1_log | pair_003 | 3228.0 | 586970.0 |
| c1_log | pair_004 | 3870.0 | 586970.0 |
| c1_log | pair_005 | 1623.0 | 586970.0 |
| c1_log | pair_006 | 2907.0 | 586970.0 |
| c1_log | pair_007 | 6759.0 | 586970.0 |
| c1_log | pair_008 | 3228.0 | 586970.0 |
| c1_log | pair_009 | 8310.0 | 586970.0 |
| c1_log | pair_010 | 3228.0 | 586970.0 |
| c1_log | pair_011 | 2586.0 | 586970.0 |
| c1_log | pair_012 | 2586.0 | 586970.0 |
| c1_log | pair_013 | 2907.0 | 586970.0 |
| c1_log | pair_014 | 2907.0 | 586970.0 |
| c1_log | pair_015 | 2907.0 | 586970.0 |
| c1_log | pair_016 | 1623.0 | 586970.0 |
| c1_log | pair_017 | 1623.0 | 586970.0 |
| c1_log | pair_018 | 4833.0 | 586970.0 |
| c1_log | pair_019 | 3228.0 | 586970.0 |
| c1_log | pair_020 | 8601.0 | 586970.0 |
| c1_log | pair_021 | 3228.0 | 586970.0 |
| c1_log | pair_022 | 8043.0 | 586970.0 |
| c1_log | pair_023 | 2586.0 | 586970.0 |
| c1_log | pair_024 | 5154.0 | 586970.0 |
| c1_log | pair_025 | 8364.0 | 586970.0 |
| c1_log | pair_026 | 2586.0 | 586970.0 |
| c1_log | pair_027 | 2586.0 | 586970.0 |
| c1_log | pair_028 | 1623.0 | 586970.0 |
| c1_log | pair_029 | 2907.0 | 586970.0 |
| c1_log | pair_030 | 2457.0 | 586970.0 |
| c1_log | pair_031 | 2705.0 | 586970.0 |
| c1_log | pair_032 | 3414.0 | 586970.0 |
| c1_log | pair_033 | 1395.0 | 586970.0 |
| c1_log | pair_034 | 2457.0 | 586970.0 |
| c1_log | pair_035 | 3336.0 | 586970.0 |
| c1_log | pair_036 | 2705.0 | 586970.0 |
| c1_log | pair_037 | 3284.0 | 586970.0 |
| c1_log | pair_038 | 2705.0 | 586970.0 |
| c1_log | pair_039 | 2209.0 | 586970.0 |
| c1_log | pair_040 | 2209.0 | 586970.0 |
| c1_log | pair_041 | 2457.0 | 586970.0 |
| c1_log | pair_042 | 2457.0 | 586970.0 |
| c1_log | pair_043 | 2457.0 | 586970.0 |
| c1_log | pair_044 | 1395.0 | 586970.0 |
| c1_log | pair_045 | 1395.0 | 586970.0 |
| c1_log | pair_046 | 3323.0 | 586970.0 |
| c1_log | pair_047 | 2705.0 | 586970.0 |
| c1_log | pair_048 | 3323.0 | 586970.0 |
| c1_log | pair_049 | 2705.0 | 586970.0 |
| c1_log | pair_050 | 3336.0 | 586970.0 |
| c1_log | pair_051 | 2209.0 | 586970.0 |
| c1_log | pair_052 | 3336.0 | 586970.0 |
| c1_log | pair_053 | 3310.0 | 586970.0 |
| c1_log | pair_054 | 2209.0 | 586970.0 |
| c1_log | pair_055 | 2209.0 | 586970.0 |
| c1_log | pair_056 | 1395.0 | 586970.0 |
| c1_log | pair_057 | 2457.0 | 586970.0 |
| c1_log | pair_058 | 2487.0 | 586970.0 |
| c1_log | pair_059 | 2567.0 | 586970.0 |
| c1_log | pair_060 | 1443.0 | 586970.0 |
| c1_log | pair_061 | 2477.0 | 586970.0 |
| c1_log | pair_062 | 2507.0 | 586970.0 |
| c1_log | pair_063 | 2527.0 | 586970.0 |
| c1_log | pair_064 | 2467.0 | 586970.0 |
| c1_log | pair_065 | 2507.0 | 586970.0 |
| c1_log | pair_066 | 2281.0 | 586970.0 |
| c1_log | pair_067 | 2281.0 | 586970.0 |
| c1_log | pair_068 | 2507.0 | 586970.0 |
| c1_log | pair_069 | 2507.0 | 586970.0 |
| c1_log | pair_070 | 2517.0 | 586970.0 |
| c1_log | pair_071 | 1443.0 | 586970.0 |
| c1_log | pair_072 | 1443.0 | 586970.0 |
| c1_log | pair_073 | 2497.0 | 586970.0 |
| c1_log | pair_074 | 2527.0 | 586970.0 |
| c1_log | pair_075 | 2497.0 | 586970.0 |
| c1_log | pair_076 | 2507.0 | 586970.0 |
| c1_log | pair_077 | 2507.0 | 586970.0 |
| c1_log | pair_078 | 2281.0 | 586970.0 |
| c1_log | pair_079 | 2507.0 | 586970.0 |
| c1_log | pair_080 | 2487.0 | 586970.0 |
| c1_log | pair_081 | 2281.0 | 586970.0 |
| c1_log | pair_082 | 2281.0 | 586970.0 |
| c1_log | pair_083 | 1443.0 | 586970.0 |
| c1_log | pair_084 | 2457.0 | 586970.0 |
| c1_log | pair_085 | 2826.0 | 586970.0 |
| c1_log | pair_086 | 1413.0 | 586970.0 |
| c1_log | pair_087 | 2487.0 | 586970.0 |
| c1_log | pair_088 | 2760.0 | 586970.0 |
| c1_log | pair_089 | 2782.0 | 586970.0 |
| c1_log | pair_090 | 2716.0 | 586970.0 |
| c1_log | pair_091 | 2760.0 | 586970.0 |
| c1_log | pair_092 | 2236.0 | 586970.0 |
| c1_log | pair_093 | 2236.0 | 586970.0 |
| c1_log | pair_094 | 2487.0 | 586970.0 |
| c1_log | pair_095 | 2487.0 | 586970.0 |
| c1_log | pair_096 | 2487.0 | 586970.0 |
| c1_log | pair_097 | 1413.0 | 586970.0 |
| c1_log | pair_098 | 1413.0 | 586970.0 |
| c1_log | pair_099 | 2749.0 | 586970.0 |
| c1_log | pair_100 | 2782.0 | 586970.0 |
| c1_log | pair_101 | 2749.0 | 586970.0 |
| c1_log | pair_102 | 2760.0 | 586970.0 |
| c1_log | pair_103 | 2760.0 | 586970.0 |
| c1_log | pair_104 | 2236.0 | 586970.0 |
| c1_log | pair_105 | 2760.0 | 586970.0 |
| c1_log | pair_106 | 2738.0 | 586970.0 |
| c1_log | pair_107 | 2236.0 | 586970.0 |
| c1_log | pair_108 | 2236.0 | 586970.0 |
| c1_log | pair_109 | 1413.0 | 586970.0 |
| c1_log | pair_110 | 2487.0 | 586970.0 |
| c1_log | pair_111 | 1461.0 | 586970.0 |
| c1_log | pair_112 | 2567.0 | 586970.0 |
| c1_log | pair_113 | 3336.0 | 586970.0 |
| c1_log | pair_114 | 2826.0 | 586970.0 |
| c1_log | pair_115 | 3284.0 | 586970.0 |
| c1_log | pair_116 | 2826.0 | 586970.0 |
| c1_log | pair_117 | 2308.0 | 586970.0 |
| c1_log | pair_118 | 2308.0 | 586970.0 |
| c1_log | pair_119 | 2567.0 | 586970.0 |
| c1_log | pair_120 | 2567.0 | 586970.0 |
| c1_log | pair_121 | 2567.0 | 586970.0 |
| c1_log | pair_122 | 1461.0 | 586970.0 |
| c1_log | pair_123 | 1461.0 | 586970.0 |
| c1_log | pair_124 | 3323.0 | 586970.0 |
| c1_log | pair_125 | 2826.0 | 586970.0 |
| c1_log | pair_126 | 3323.0 | 586970.0 |
| c1_log | pair_127 | 2826.0 | 586970.0 |
| c1_log | pair_128 | 3336.0 | 586970.0 |
| c1_log | pair_129 | 2308.0 | 586970.0 |
| c1_log | pair_130 | 3336.0 | 586970.0 |
| c1_log | pair_131 | 3310.0 | 586970.0 |
| c1_log | pair_132 | 2308.0 | 586970.0 |
| c1_log | pair_133 | 2308.0 | 586970.0 |
| c1_log | pair_134 | 1461.0 | 586970.0 |
| c1_log | pair_135 | 2567.0 | 586970.0 |
| c1_log | pair_136 | 1407.0 | 586970.0 |
| c1_log | pair_137 | 1425.0 | 586970.0 |
| c1_log | pair_138 | 1437.0 | 586970.0 |
| c1_log | pair_139 | 1401.0 | 586970.0 |
| c1_log | pair_140 | 1425.0 | 586970.0 |
| c1_log | pair_141 | 1407.0 | 586970.0 |
| c1_log | pair_142 | 1437.0 | 586970.0 |
| c1_log | pair_143 | 1425.0 | 586970.0 |
| c1_log | pair_144 | 1425.0 | 586970.0 |
| c1_log | pair_145 | 1431.0 | 586970.0 |
| c1_log | pair_146 | 1425.0 | 586970.0 |
| c1_log | pair_147 | 1419.0 | 586970.0 |
| c1_log | pair_148 | 1419.0 | 586970.0 |
| c1_log | pair_149 | 1437.0 | 586970.0 |
| c1_log | pair_150 | 1419.0 | 586970.0 |
| c1_log | pair_151 | 1425.0 | 586970.0 |
| c1_log | pair_152 | 1425.0 | 586970.0 |
| c1_log | pair_153 | 1413.0 | 586970.0 |
| c1_log | pair_154 | 1425.0 | 586970.0 |
| c1_log | pair_155 | 1413.0 | 586970.0 |
| c1_log | pair_156 | 1413.0 | 586970.0 |
| c1_log | pair_157 | 1419.0 | 586970.0 |
| c1_log | pair_158 | 1455.0 | 586970.0 |
| c1_log | pair_159 | 1395.0 | 586970.0 |
| c1_log | pair_160 | 2507.0 | 586970.0 |
| c1_log | pair_161 | 2527.0 | 586970.0 |
| c1_log | pair_162 | 2467.0 | 586970.0 |
| c1_log | pair_163 | 2507.0 | 586970.0 |
| c1_log | pair_164 | 2227.0 | 586970.0 |
| c1_log | pair_165 | 2227.0 | 586970.0 |
| c1_log | pair_166 | 2507.0 | 586970.0 |
| c1_log | pair_167 | 2507.0 | 586970.0 |
| c1_log | pair_168 | 2517.0 | 586970.0 |
| c1_log | pair_169 | 1407.0 | 586970.0 |
| c1_log | pair_170 | 1407.0 | 586970.0 |
| c1_log | pair_171 | 2497.0 | 586970.0 |
| c1_log | pair_172 | 2527.0 | 586970.0 |
| c1_log | pair_173 | 2497.0 | 586970.0 |
| c1_log | pair_174 | 2507.0 | 586970.0 |
| c1_log | pair_175 | 2507.0 | 586970.0 |
| c1_log | pair_176 | 2227.0 | 586970.0 |
| c1_log | pair_177 | 2507.0 | 586970.0 |
| c1_log | pair_178 | 2487.0 | 586970.0 |
| c1_log | pair_179 | 2227.0 | 586970.0 |
| c1_log | pair_180 | 2227.0 | 586970.0 |
| c1_log | pair_181 | 1407.0 | 586970.0 |
| c1_log | pair_182 | 2457.0 | 586970.0 |
| c1_log | pair_183 | 2760.0 | 586970.0 |
| c1_log | pair_184 | 5595.0 | 586970.0 |
| c1_log | pair_185 | 2760.0 | 586970.0 |
| c1_log | pair_186 | 2254.0 | 586970.0 |
| c1_log | pair_187 | 2254.0 | 586970.0 |
| c1_log | pair_188 | 2507.0 | 586970.0 |
| c1_log | pair_189 | 2507.0 | 586970.0 |
| c1_log | pair_190 | 2507.0 | 586970.0 |
| c1_log | pair_191 | 1425.0 | 586970.0 |
| c1_log | pair_192 | 1425.0 | 586970.0 |
| c1_log | pair_193 | 4095.0 | 586970.0 |
| c1_log | pair_194 | 2760.0 | 586970.0 |
| c1_log | pair_195 | 5661.0 | 586970.0 |
| c1_log | pair_196 | 2760.0 | 586970.0 |
| c1_log | pair_197 | 5683.0 | 586970.0 |
| c1_log | pair_198 | 2254.0 | 586970.0 |
| c1_log | pair_199 | 4348.0 | 586970.0 |
| c1_log | pair_200 | 5639.0 | 586970.0 |
| c1_log | pair_201 | 2254.0 | 586970.0 |
| c1_log | pair_202 | 2254.0 | 586970.0 |
| c1_log | pair_203 | 1425.0 | 586970.0 |
| c1_log | pair_204 | 2507.0 | 586970.0 |
| c1_log | pair_205 | 2716.0 | 586970.0 |
| c1_log | pair_206 | 2760.0 | 586970.0 |
| c1_log | pair_207 | 2272.0 | 586970.0 |
| c1_log | pair_208 | 2272.0 | 586970.0 |
| c1_log | pair_209 | 2527.0 | 586970.0 |
| c1_log | pair_210 | 2527.0 | 586970.0 |
| c1_log | pair_211 | 2527.0 | 586970.0 |
| c1_log | pair_212 | 1437.0 | 586970.0 |
| c1_log | pair_213 | 1437.0 | 586970.0 |
| c1_log | pair_214 | 2749.0 | 586970.0 |
| c1_log | pair_215 | 2782.0 | 586970.0 |
| c1_log | pair_216 | 2749.0 | 586970.0 |
| c1_log | pair_217 | 2760.0 | 586970.0 |
| c1_log | pair_218 | 2760.0 | 586970.0 |
| c1_log | pair_219 | 2272.0 | 586970.0 |
| c1_log | pair_220 | 2760.0 | 586970.0 |
| c1_log | pair_221 | 2738.0 | 586970.0 |
| c1_log | pair_222 | 2272.0 | 586970.0 |
| c1_log | pair_223 | 2272.0 | 586970.0 |
| c1_log | pair_224 | 1437.0 | 586970.0 |
| c1_log | pair_225 | 2527.0 | 586970.0 |
| c1_log | pair_226 | 2716.0 | 586970.0 |
| c1_log | pair_227 | 2218.0 | 586970.0 |
| c1_log | pair_228 | 2218.0 | 586970.0 |
| c1_log | pair_229 | 2467.0 | 586970.0 |
| c1_log | pair_230 | 2467.0 | 586970.0 |
| c1_log | pair_231 | 2467.0 | 586970.0 |
| c1_log | pair_232 | 1401.0 | 586970.0 |
| c1_log | pair_233 | 1401.0 | 586970.0 |
| c1_log | pair_234 | 4031.0 | 586970.0 |
| c1_log | pair_235 | 2716.0 | 586970.0 |
| c1_log | pair_236 | 6921.0 | 586970.0 |
| c1_log | pair_237 | 2716.0 | 586970.0 |
| c1_log | pair_238 | 6591.0 | 586970.0 |
| c1_log | pair_239 | 2218.0 | 586970.0 |
| c1_log | pair_240 | 4280.0 | 586970.0 |
| c1_log | pair_241 | 6894.0 | 586970.0 |
| c1_log | pair_242 | 2218.0 | 586970.0 |
| c1_log | pair_243 | 2218.0 | 586970.0 |
| c1_log | pair_244 | 1401.0 | 586970.0 |
| c1_log | pair_245 | 2467.0 | 586970.0 |
| c1_log | pair_246 | 2254.0 | 586970.0 |
| c1_log | pair_247 | 2254.0 | 586970.0 |
| c1_log | pair_248 | 2507.0 | 586970.0 |
| c1_log | pair_249 | 2507.0 | 586970.0 |
| c1_log | pair_250 | 2507.0 | 586970.0 |
| c1_log | pair_251 | 1425.0 | 586970.0 |
| c1_log | pair_252 | 1425.0 | 586970.0 |
| c1_log | pair_253 | 2749.0 | 586970.0 |
| c1_log | pair_254 | 2782.0 | 586970.0 |
| c1_log | pair_255 | 2749.0 | 586970.0 |
| c1_log | pair_256 | 2760.0 | 586970.0 |
| c1_log | pair_257 | 2760.0 | 586970.0 |
| c1_log | pair_258 | 2254.0 | 586970.0 |
| c1_log | pair_259 | 2760.0 | 586970.0 |
| c1_log | pair_260 | 2738.0 | 586970.0 |
| c1_log | pair_261 | 2254.0 | 586970.0 |
| c1_log | pair_262 | 2254.0 | 586970.0 |
| c1_log | pair_263 | 1425.0 | 586970.0 |
| c1_log | pair_264 | 2507.0 | 586970.0 |
| c1_log | pair_265 | 2272.0 | 586970.0 |
| c1_log | pair_266 | 2254.0 | 586970.0 |
| c1_log | pair_267 | 2254.0 | 586970.0 |
| c1_log | pair_268 | 2263.0 | 586970.0 |
| c1_log | pair_269 | 1407.0 | 586970.0 |
| c1_log | pair_270 | 1407.0 | 586970.0 |
| c1_log | pair_271 | 2245.0 | 586970.0 |
| c1_log | pair_272 | 2272.0 | 586970.0 |
| c1_log | pair_273 | 2245.0 | 586970.0 |
| c1_log | pair_274 | 2254.0 | 586970.0 |
| c1_log | pair_275 | 2254.0 | 586970.0 |
| c1_log | pair_276 | 2236.0 | 586970.0 |
| c1_log | pair_277 | 2254.0 | 586970.0 |
| c1_log | pair_278 | 2236.0 | 586970.0 |
| c1_log | pair_279 | 2236.0 | 586970.0 |
| c1_log | pair_280 | 2245.0 | 586970.0 |
| c1_log | pair_281 | 1407.0 | 586970.0 |
| c1_log | pair_282 | 2209.0 | 586970.0 |
| c1_log | pair_283 | 2254.0 | 586970.0 |
| c1_log | pair_284 | 2254.0 | 586970.0 |
| c1_log | pair_285 | 2263.0 | 586970.0 |
| c1_log | pair_286 | 1437.0 | 586970.0 |
| c1_log | pair_287 | 1437.0 | 586970.0 |
| c1_log | pair_288 | 2245.0 | 586970.0 |
| c1_log | pair_289 | 2272.0 | 586970.0 |
| c1_log | pair_290 | 2245.0 | 586970.0 |
| c1_log | pair_291 | 2254.0 | 586970.0 |
| c1_log | pair_292 | 2254.0 | 586970.0 |
| c1_log | pair_293 | 2236.0 | 586970.0 |
| c1_log | pair_294 | 2254.0 | 586970.0 |
| c1_log | pair_295 | 2236.0 | 586970.0 |
| c1_log | pair_296 | 2236.0 | 586970.0 |
| c1_log | pair_297 | 2245.0 | 586970.0 |
| c1_log | pair_298 | 1437.0 | 586970.0 |
| c1_log | pair_299 | 2209.0 | 586970.0 |
| c1_log | pair_300 | 2507.0 | 586970.0 |
| c1_log | pair_301 | 2517.0 | 586970.0 |
| c1_log | pair_302 | 1425.0 | 586970.0 |
| c1_log | pair_303 | 1425.0 | 586970.0 |
| c1_log | pair_304 | 2497.0 | 586970.0 |
| c1_log | pair_305 | 2527.0 | 586970.0 |
| c1_log | pair_306 | 2497.0 | 586970.0 |
| c1_log | pair_307 | 2507.0 | 586970.0 |
| c1_log | pair_308 | 2507.0 | 586970.0 |
| c1_log | pair_309 | 2254.0 | 586970.0 |
| c1_log | pair_310 | 2507.0 | 586970.0 |
| c1_log | pair_311 | 2487.0 | 586970.0 |
| c1_log | pair_312 | 2254.0 | 586970.0 |
| c1_log | pair_313 | 2254.0 | 586970.0 |
| c1_log | pair_314 | 1425.0 | 586970.0 |
| c1_log | pair_315 | 2457.0 | 586970.0 |
| c1_log | pair_316 | 2517.0 | 586970.0 |
| c1_log | pair_317 | 1425.0 | 586970.0 |
| c1_log | pair_318 | 1425.0 | 586970.0 |
| c1_log | pair_319 | 2497.0 | 586970.0 |
| c1_log | pair_320 | 2527.0 | 586970.0 |
| c1_log | pair_321 | 2497.0 | 586970.0 |
| c1_log | pair_322 | 2507.0 | 586970.0 |
| c1_log | pair_323 | 2507.0 | 586970.0 |
| c1_log | pair_324 | 2254.0 | 586970.0 |
| c1_log | pair_325 | 2507.0 | 586970.0 |
| c1_log | pair_326 | 2487.0 | 586970.0 |
| c1_log | pair_327 | 2254.0 | 586970.0 |
| c1_log | pair_328 | 2254.0 | 586970.0 |
| c1_log | pair_329 | 1425.0 | 586970.0 |
| c1_log | pair_330 | 2457.0 | 586970.0 |
| c1_log | pair_331 | 1431.0 | 586970.0 |
| c1_log | pair_332 | 1431.0 | 586970.0 |
| c1_log | pair_333 | 2497.0 | 586970.0 |
| c1_log | pair_334 | 2527.0 | 586970.0 |
| c1_log | pair_335 | 2497.0 | 586970.0 |
| c1_log | pair_336 | 2507.0 | 586970.0 |
| c1_log | pair_337 | 2507.0 | 586970.0 |
| c1_log | pair_338 | 2263.0 | 586970.0 |
| c1_log | pair_339 | 2507.0 | 586970.0 |
| c1_log | pair_340 | 2487.0 | 586970.0 |
| c1_log | pair_341 | 2263.0 | 586970.0 |
| c1_log | pair_342 | 2263.0 | 586970.0 |
| c1_log | pair_343 | 1431.0 | 586970.0 |
| c1_log | pair_344 | 2457.0 | 586970.0 |
| c1_log | pair_345 | 1419.0 | 586970.0 |
| c1_log | pair_346 | 1419.0 | 586970.0 |
| c1_log | pair_347 | 1437.0 | 586970.0 |
| c1_log | pair_348 | 1419.0 | 586970.0 |
| c1_log | pair_349 | 1425.0 | 586970.0 |
| c1_log | pair_350 | 1425.0 | 586970.0 |
| c1_log | pair_351 | 1413.0 | 586970.0 |
| c1_log | pair_352 | 1425.0 | 586970.0 |
| c1_log | pair_353 | 1413.0 | 586970.0 |
| c1_log | pair_354 | 1413.0 | 586970.0 |
| c1_log | pair_355 | 1419.0 | 586970.0 |
| c1_log | pair_356 | 1455.0 | 586970.0 |
| c1_log | pair_357 | 1395.0 | 586970.0 |
| c1_log | pair_358 | 1419.0 | 586970.0 |
| c1_log | pair_359 | 1437.0 | 586970.0 |
| c1_log | pair_360 | 1419.0 | 586970.0 |
| c1_log | pair_361 | 1425.0 | 586970.0 |
| c1_log | pair_362 | 1425.0 | 586970.0 |
| c1_log | pair_363 | 1413.0 | 586970.0 |
| c1_log | pair_364 | 1425.0 | 586970.0 |
| c1_log | pair_365 | 1413.0 | 586970.0 |
| c1_log | pair_366 | 1413.0 | 586970.0 |
| c1_log | pair_367 | 1419.0 | 586970.0 |
| c1_log | pair_368 | 1455.0 | 586970.0 |
| c1_log | pair_369 | 1395.0 | 586970.0 |
| c1_log | pair_370 | 2749.0 | 586970.0 |
| c1_log | pair_371 | 4079.0 | 586970.0 |
| c1_log | pair_372 | 2749.0 | 586970.0 |
| c1_log | pair_373 | 4095.0 | 586970.0 |
| c1_log | pair_374 | 2245.0 | 586970.0 |
| c1_log | pair_375 | 4095.0 | 586970.0 |
| c1_log | pair_376 | 4063.0 | 586970.0 |
| c1_log | pair_377 | 2245.0 | 586970.0 |
| c1_log | pair_378 | 2245.0 | 586970.0 |
| c1_log | pair_379 | 1419.0 | 586970.0 |
| c1_log | pair_380 | 2497.0 | 586970.0 |
| c1_log | pair_381 | 2749.0 | 586970.0 |
| c1_log | pair_382 | 2760.0 | 586970.0 |
| c1_log | pair_383 | 2760.0 | 586970.0 |
| c1_log | pair_384 | 2272.0 | 586970.0 |
| c1_log | pair_385 | 2760.0 | 586970.0 |
| c1_log | pair_386 | 2738.0 | 586970.0 |
| c1_log | pair_387 | 2272.0 | 586970.0 |
| c1_log | pair_388 | 2272.0 | 586970.0 |
| c1_log | pair_389 | 1437.0 | 586970.0 |
| c1_log | pair_390 | 2527.0 | 586970.0 |
| c1_log | pair_391 | 2749.0 | 586970.0 |
| c1_log | pair_392 | 6669.0 | 586970.0 |
| c1_log | pair_393 | 2245.0 | 586970.0 |
| c1_log | pair_394 | 4331.0 | 586970.0 |
| c1_log | pair_395 | 6921.0 | 586970.0 |
| c1_log | pair_396 | 2245.0 | 586970.0 |
| c1_log | pair_397 | 2245.0 | 586970.0 |
| c1_log | pair_398 | 1419.0 | 586970.0 |
| c1_log | pair_399 | 2497.0 | 586970.0 |
| c1_log | pair_400 | 2760.0 | 586970.0 |
| c1_log | pair_401 | 2254.0 | 586970.0 |
| c1_log | pair_402 | 2760.0 | 586970.0 |
| c1_log | pair_403 | 2738.0 | 586970.0 |
| c1_log | pair_404 | 2254.0 | 586970.0 |
| c1_log | pair_405 | 2254.0 | 586970.0 |
| c1_log | pair_406 | 1425.0 | 586970.0 |
| c1_log | pair_407 | 2507.0 | 586970.0 |
| c1_log | pair_408 | 2254.0 | 586970.0 |
| c1_log | pair_409 | 4348.0 | 586970.0 |
| c1_log | pair_410 | 6643.0 | 586970.0 |
| c1_log | pair_411 | 2254.0 | 586970.0 |
| c1_log | pair_412 | 2254.0 | 586970.0 |
| c1_log | pair_413 | 1425.0 | 586970.0 |
| c1_log | pair_414 | 2507.0 | 586970.0 |
| c1_log | pair_415 | 2254.0 | 586970.0 |
| c1_log | pair_416 | 2236.0 | 586970.0 |
| c1_log | pair_417 | 2236.0 | 586970.0 |
| c1_log | pair_418 | 2245.0 | 586970.0 |
| c1_log | pair_419 | 1413.0 | 586970.0 |
| c1_log | pair_420 | 2209.0 | 586970.0 |
| c1_log | pair_421 | 4314.0 | 586970.0 |
| c1_log | pair_422 | 2254.0 | 586970.0 |
| c1_log | pair_423 | 2254.0 | 586970.0 |
| c1_log | pair_424 | 1425.0 | 586970.0 |
| c1_log | pair_425 | 2507.0 | 586970.0 |
| c1_log | pair_426 | 2236.0 | 586970.0 |
| c1_log | pair_427 | 2236.0 | 586970.0 |
| c1_log | pair_428 | 1413.0 | 586970.0 |
| c1_log | pair_429 | 2487.0 | 586970.0 |
| c1_log | pair_430 | 2245.0 | 586970.0 |
| c1_log | pair_431 | 1413.0 | 586970.0 |
| c1_log | pair_432 | 2209.0 | 586970.0 |
| c1_log | pair_433 | 1419.0 | 586970.0 |
| c1_log | pair_434 | 2209.0 | 586970.0 |
| c1_log | pair_435 | 1395.0 | 586970.0 |
| c1_power2 | pair_001 | 4056.0 | 586970.0 |
| c1_power2 | pair_002 | 2586.0 | 586970.0 |
| c1_power2 | pair_003 | 2586.0 | 586970.0 |
| c1_power2 | pair_004 | 2586.0 | 586970.0 |
| c1_power2 | pair_005 | 1302.0 | 586970.0 |
| c1_power2 | pair_006 | 2586.0 | 586970.0 |
| c1_power2 | pair_007 | 5154.0 | 586970.0 |
| c1_power2 | pair_008 | 2586.0 | 586970.0 |
| c1_power2 | pair_009 | 10224.0 | 586970.0 |
| c1_power2 | pair_010 | 2586.0 | 586970.0 |
| c1_power2 | pair_011 | 2586.0 | 586970.0 |
| c1_power2 | pair_012 | 2586.0 | 586970.0 |
| c1_power2 | pair_013 | 2586.0 | 586970.0 |
| c1_power2 | pair_014 | 2586.0 | 586970.0 |
| c1_power2 | pair_015 | 2586.0 | 586970.0 |
| c1_power2 | pair_016 | 1302.0 | 586970.0 |
| c1_power2 | pair_017 | 1302.0 | 586970.0 |
| c1_power2 | pair_018 | 5154.0 | 586970.0 |
| c1_power2 | pair_019 | 2586.0 | 586970.0 |
| c1_power2 | pair_020 | 10533.0 | 586970.0 |
| c1_power2 | pair_021 | 2586.0 | 586970.0 |
| c1_power2 | pair_022 | 10290.0 | 586970.0 |
| c1_power2 | pair_023 | 2586.0 | 586970.0 |
| c1_power2 | pair_024 | 5154.0 | 586970.0 |
| c1_power2 | pair_025 | 10290.0 | 586970.0 |
| c1_power2 | pair_026 | 2586.0 | 586970.0 |
| c1_power2 | pair_027 | 2586.0 | 586970.0 |
| c1_power2 | pair_028 | 1302.0 | 586970.0 |
| c1_power2 | pair_029 | 2586.0 | 586970.0 |
| c1_power2 | pair_030 | 2139.0 | 586970.0 |
| c1_power2 | pair_031 | 2069.0 | 586970.0 |
| c1_power2 | pair_032 | 2098.0 | 586970.0 |
| c1_power2 | pair_033 | 1077.0 | 586970.0 |
| c1_power2 | pair_034 | 2139.0 | 586970.0 |
| c1_power2 | pair_035 | 2044.0 | 586970.0 |
| c1_power2 | pair_036 | 2069.0 | 586970.0 |
| c1_power2 | pair_037 | 2008.0 | 586970.0 |
| c1_power2 | pair_038 | 2069.0 | 586970.0 |
| c1_power2 | pair_039 | 2209.0 | 586970.0 |
| c1_power2 | pair_040 | 2209.0 | 586970.0 |
| c1_power2 | pair_041 | 2139.0 | 586970.0 |
| c1_power2 | pair_042 | 2139.0 | 586970.0 |
| c1_power2 | pair_043 | 2139.0 | 586970.0 |
| c1_power2 | pair_044 | 1077.0 | 586970.0 |
| c1_power2 | pair_045 | 1077.0 | 586970.0 |
| c1_power2 | pair_046 | 2035.0 | 586970.0 |
| c1_power2 | pair_047 | 2069.0 | 586970.0 |
| c1_power2 | pair_048 | 2035.0 | 586970.0 |
| c1_power2 | pair_049 | 2069.0 | 586970.0 |
| c1_power2 | pair_050 | 2044.0 | 586970.0 |
| c1_power2 | pair_051 | 2209.0 | 586970.0 |
| c1_power2 | pair_052 | 2044.0 | 586970.0 |
| c1_power2 | pair_053 | 2026.0 | 586970.0 |
| c1_power2 | pair_054 | 2209.0 | 586970.0 |
| c1_power2 | pair_055 | 2209.0 | 586970.0 |
| c1_power2 | pair_056 | 1077.0 | 586970.0 |
| c1_power2 | pair_057 | 2139.0 | 586970.0 |
| c1_power2 | pair_058 | 2166.0 | 586970.0 |
| c1_power2 | pair_059 | 2238.0 | 586970.0 |
| c1_power2 | pair_060 | 1117.0 | 586970.0 |
| c1_power2 | pair_061 | 2157.0 | 586970.0 |
| c1_power2 | pair_062 | 2184.0 | 586970.0 |
| c1_power2 | pair_063 | 2202.0 | 586970.0 |
| c1_power2 | pair_064 | 2148.0 | 586970.0 |
| c1_power2 | pair_065 | 2184.0 | 586970.0 |
| c1_power2 | pair_066 | 2281.0 | 586970.0 |
| c1_power2 | pair_067 | 2281.0 | 586970.0 |
| c1_power2 | pair_068 | 2184.0 | 586970.0 |
| c1_power2 | pair_069 | 2184.0 | 586970.0 |
| c1_power2 | pair_070 | 2193.0 | 586970.0 |
| c1_power2 | pair_071 | 1117.0 | 586970.0 |
| c1_power2 | pair_072 | 1117.0 | 586970.0 |
| c1_power2 | pair_073 | 2175.0 | 586970.0 |
| c1_power2 | pair_074 | 2202.0 | 586970.0 |
| c1_power2 | pair_075 | 2175.0 | 586970.0 |
| c1_power2 | pair_076 | 2184.0 | 586970.0 |
| c1_power2 | pair_077 | 2184.0 | 586970.0 |
| c1_power2 | pair_078 | 2281.0 | 586970.0 |
| c1_power2 | pair_079 | 2184.0 | 586970.0 |
| c1_power2 | pair_080 | 2166.0 | 586970.0 |
| c1_power2 | pair_081 | 2281.0 | 586970.0 |
| c1_power2 | pair_082 | 2281.0 | 586970.0 |
| c1_power2 | pair_083 | 1117.0 | 586970.0 |
| c1_power2 | pair_084 | 2139.0 | 586970.0 |
| c1_power2 | pair_085 | 2168.0 | 586970.0 |
| c1_power2 | pair_086 | 1092.0 | 586970.0 |
| c1_power2 | pair_087 | 2166.0 | 586970.0 |
| c1_power2 | pair_088 | 2114.0 | 586970.0 |
| c1_power2 | pair_089 | 2132.0 | 586970.0 |
| c1_power2 | pair_090 | 2078.0 | 586970.0 |
| c1_power2 | pair_091 | 2114.0 | 586970.0 |
| c1_power2 | pair_092 | 2236.0 | 586970.0 |
| c1_power2 | pair_093 | 2236.0 | 586970.0 |
| c1_power2 | pair_094 | 2166.0 | 586970.0 |
| c1_power2 | pair_095 | 2166.0 | 586970.0 |
| c1_power2 | pair_096 | 2166.0 | 586970.0 |
| c1_power2 | pair_097 | 1092.0 | 586970.0 |
| c1_power2 | pair_098 | 1092.0 | 586970.0 |
| c1_power2 | pair_099 | 2105.0 | 586970.0 |
| c1_power2 | pair_100 | 2132.0 | 586970.0 |
| c1_power2 | pair_101 | 2105.0 | 586970.0 |
| c1_power2 | pair_102 | 2114.0 | 586970.0 |
| c1_power2 | pair_103 | 2114.0 | 586970.0 |
| c1_power2 | pair_104 | 2236.0 | 586970.0 |
| c1_power2 | pair_105 | 2114.0 | 586970.0 |
| c1_power2 | pair_106 | 2096.0 | 586970.0 |
| c1_power2 | pair_107 | 2236.0 | 586970.0 |
| c1_power2 | pair_108 | 2236.0 | 586970.0 |
| c1_power2 | pair_109 | 1092.0 | 586970.0 |
| c1_power2 | pair_110 | 2166.0 | 586970.0 |
| c1_power2 | pair_111 | 1132.0 | 586970.0 |
| c1_power2 | pair_112 | 2238.0 | 586970.0 |
| c1_power2 | pair_113 | 2044.0 | 586970.0 |
| c1_power2 | pair_114 | 2168.0 | 586970.0 |
| c1_power2 | pair_115 | 2008.0 | 586970.0 |
| c1_power2 | pair_116 | 2168.0 | 586970.0 |
| c1_power2 | pair_117 | 2308.0 | 586970.0 |
| c1_power2 | pair_118 | 2308.0 | 586970.0 |
| c1_power2 | pair_119 | 2238.0 | 586970.0 |
| c1_power2 | pair_120 | 2238.0 | 586970.0 |
| c1_power2 | pair_121 | 2238.0 | 586970.0 |
| c1_power2 | pair_122 | 1132.0 | 586970.0 |
| c1_power2 | pair_123 | 1132.0 | 586970.0 |
| c1_power2 | pair_124 | 2035.0 | 586970.0 |
| c1_power2 | pair_125 | 2168.0 | 586970.0 |
| c1_power2 | pair_126 | 2035.0 | 586970.0 |
| c1_power2 | pair_127 | 2168.0 | 586970.0 |
| c1_power2 | pair_128 | 2044.0 | 586970.0 |
| c1_power2 | pair_129 | 2308.0 | 586970.0 |
| c1_power2 | pair_130 | 2044.0 | 586970.0 |
| c1_power2 | pair_131 | 2026.0 | 586970.0 |
| c1_power2 | pair_132 | 2308.0 | 586970.0 |
| c1_power2 | pair_133 | 2308.0 | 586970.0 |
| c1_power2 | pair_134 | 1132.0 | 586970.0 |
| c1_power2 | pair_135 | 2238.0 | 586970.0 |
| c1_power2 | pair_136 | 1087.0 | 586970.0 |
| c1_power2 | pair_137 | 1102.0 | 586970.0 |
| c1_power2 | pair_138 | 1112.0 | 586970.0 |
| c1_power2 | pair_139 | 1082.0 | 586970.0 |
| c1_power2 | pair_140 | 1102.0 | 586970.0 |
| c1_power2 | pair_141 | 1087.0 | 586970.0 |
| c1_power2 | pair_142 | 1112.0 | 586970.0 |
| c1_power2 | pair_143 | 1102.0 | 586970.0 |
| c1_power2 | pair_144 | 1102.0 | 586970.0 |
| c1_power2 | pair_145 | 1107.0 | 586970.0 |
| c1_power2 | pair_146 | 1102.0 | 586970.0 |
| c1_power2 | pair_147 | 1097.0 | 586970.0 |
| c1_power2 | pair_148 | 1097.0 | 586970.0 |
| c1_power2 | pair_149 | 1112.0 | 586970.0 |
| c1_power2 | pair_150 | 1097.0 | 586970.0 |
| c1_power2 | pair_151 | 1102.0 | 586970.0 |
| c1_power2 | pair_152 | 1102.0 | 586970.0 |
| c1_power2 | pair_153 | 1092.0 | 586970.0 |
| c1_power2 | pair_154 | 1102.0 | 586970.0 |
| c1_power2 | pair_155 | 1092.0 | 586970.0 |
| c1_power2 | pair_156 | 1092.0 | 586970.0 |
| c1_power2 | pair_157 | 1097.0 | 586970.0 |
| c1_power2 | pair_158 | 1127.0 | 586970.0 |
| c1_power2 | pair_159 | 1077.0 | 586970.0 |
| c1_power2 | pair_160 | 2184.0 | 586970.0 |
| c1_power2 | pair_161 | 2202.0 | 586970.0 |
| c1_power2 | pair_162 | 2148.0 | 586970.0 |
| c1_power2 | pair_163 | 2184.0 | 586970.0 |
| c1_power2 | pair_164 | 2227.0 | 586970.0 |
| c1_power2 | pair_165 | 2227.0 | 586970.0 |
| c1_power2 | pair_166 | 2184.0 | 586970.0 |
| c1_power2 | pair_167 | 2184.0 | 586970.0 |
| c1_power2 | pair_168 | 2193.0 | 586970.0 |
| c1_power2 | pair_169 | 1087.0 | 586970.0 |
| c1_power2 | pair_170 | 1087.0 | 586970.0 |
| c1_power2 | pair_171 | 2175.0 | 586970.0 |
| c1_power2 | pair_172 | 2202.0 | 586970.0 |
| c1_power2 | pair_173 | 2175.0 | 586970.0 |
| c1_power2 | pair_174 | 2184.0 | 586970.0 |
| c1_power2 | pair_175 | 2184.0 | 586970.0 |
| c1_power2 | pair_176 | 2227.0 | 586970.0 |
| c1_power2 | pair_177 | 2184.0 | 586970.0 |
| c1_power2 | pair_178 | 2166.0 | 586970.0 |
| c1_power2 | pair_179 | 2227.0 | 586970.0 |
| c1_power2 | pair_180 | 2227.0 | 586970.0 |
| c1_power2 | pair_181 | 1087.0 | 586970.0 |
| c1_power2 | pair_182 | 2139.0 | 586970.0 |
| c1_power2 | pair_183 | 2114.0 | 586970.0 |
| c1_power2 | pair_184 | 4000.0 | 586970.0 |
| c1_power2 | pair_185 | 2114.0 | 586970.0 |
| c1_power2 | pair_186 | 2254.0 | 586970.0 |
| c1_power2 | pair_187 | 2254.0 | 586970.0 |
| c1_power2 | pair_188 | 2184.0 | 586970.0 |
| c1_power2 | pair_189 | 2184.0 | 586970.0 |
| c1_power2 | pair_190 | 2184.0 | 586970.0 |
| c1_power2 | pair_191 | 1102.0 | 586970.0 |
| c1_power2 | pair_192 | 1102.0 | 586970.0 |
| c1_power2 | pair_193 | 4418.0 | 586970.0 |
| c1_power2 | pair_194 | 2114.0 | 586970.0 |
| c1_power2 | pair_195 | 4051.0 | 586970.0 |
| c1_power2 | pair_196 | 2114.0 | 586970.0 |
| c1_power2 | pair_197 | 4068.0 | 586970.0 |
| c1_power2 | pair_198 | 2254.0 | 586970.0 |
| c1_power2 | pair_199 | 4348.0 | 586970.0 |
| c1_power2 | pair_200 | 4034.0 | 586970.0 |
| c1_power2 | pair_201 | 2254.0 | 586970.0 |
| c1_power2 | pair_202 | 2254.0 | 586970.0 |
| c1_power2 | pair_203 | 1102.0 | 586970.0 |
| c1_power2 | pair_204 | 2184.0 | 586970.0 |
| c1_power2 | pair_205 | 2078.0 | 586970.0 |
| c1_power2 | pair_206 | 2114.0 | 586970.0 |
| c1_power2 | pair_207 | 2272.0 | 586970.0 |
| c1_power2 | pair_208 | 2272.0 | 586970.0 |
| c1_power2 | pair_209 | 2202.0 | 586970.0 |
| c1_power2 | pair_210 | 2202.0 | 586970.0 |
| c1_power2 | pair_211 | 2202.0 | 586970.0 |
| c1_power2 | pair_212 | 1112.0 | 586970.0 |
| c1_power2 | pair_213 | 1112.0 | 586970.0 |
| c1_power2 | pair_214 | 2105.0 | 586970.0 |
| c1_power2 | pair_215 | 2132.0 | 586970.0 |
| c1_power2 | pair_216 | 2105.0 | 586970.0 |
| c1_power2 | pair_217 | 2114.0 | 586970.0 |
| c1_power2 | pair_218 | 2114.0 | 586970.0 |
| c1_power2 | pair_219 | 2272.0 | 586970.0 |
| c1_power2 | pair_220 | 2114.0 | 586970.0 |
| c1_power2 | pair_221 | 2096.0 | 586970.0 |
| c1_power2 | pair_222 | 2272.0 | 586970.0 |
| c1_power2 | pair_223 | 2272.0 | 586970.0 |
| c1_power2 | pair_224 | 1112.0 | 586970.0 |
| c1_power2 | pair_225 | 2202.0 | 586970.0 |
| c1_power2 | pair_226 | 2078.0 | 586970.0 |
| c1_power2 | pair_227 | 2218.0 | 586970.0 |
| c1_power2 | pair_228 | 2218.0 | 586970.0 |
| c1_power2 | pair_229 | 2148.0 | 586970.0 |
| c1_power2 | pair_230 | 2148.0 | 586970.0 |
| c1_power2 | pair_231 | 2148.0 | 586970.0 |
| c1_power2 | pair_232 | 1082.0 | 586970.0 |
| c1_power2 | pair_233 | 1082.0 | 586970.0 |
| c1_power2 | pair_234 | 4350.0 | 586970.0 |
| c1_power2 | pair_235 | 2078.0 | 586970.0 |
| c1_power2 | pair_236 | 8853.0 | 586970.0 |
| c1_power2 | pair_237 | 2078.0 | 586970.0 |
| c1_power2 | pair_238 | 8824.0 | 586970.0 |
| c1_power2 | pair_239 | 2218.0 | 586970.0 |
| c1_power2 | pair_240 | 4280.0 | 586970.0 |
| c1_power2 | pair_241 | 8820.0 | 586970.0 |
| c1_power2 | pair_242 | 2218.0 | 586970.0 |
| c1_power2 | pair_243 | 2218.0 | 586970.0 |
| c1_power2 | pair_244 | 1082.0 | 586970.0 |
| c1_power2 | pair_245 | 2148.0 | 586970.0 |
| c1_power2 | pair_246 | 2254.0 | 586970.0 |
| c1_power2 | pair_247 | 2254.0 | 586970.0 |
| c1_power2 | pair_248 | 2184.0 | 586970.0 |
| c1_power2 | pair_249 | 2184.0 | 586970.0 |
| c1_power2 | pair_250 | 2184.0 | 586970.0 |
| c1_power2 | pair_251 | 1102.0 | 586970.0 |
| c1_power2 | pair_252 | 1102.0 | 586970.0 |
| c1_power2 | pair_253 | 2105.0 | 586970.0 |
| c1_power2 | pair_254 | 2132.0 | 586970.0 |
| c1_power2 | pair_255 | 2105.0 | 586970.0 |
| c1_power2 | pair_256 | 2114.0 | 586970.0 |
| c1_power2 | pair_257 | 2114.0 | 586970.0 |
| c1_power2 | pair_258 | 2254.0 | 586970.0 |
| c1_power2 | pair_259 | 2114.0 | 586970.0 |
| c1_power2 | pair_260 | 2096.0 | 586970.0 |
| c1_power2 | pair_261 | 2254.0 | 586970.0 |
| c1_power2 | pair_262 | 2254.0 | 586970.0 |
| c1_power2 | pair_263 | 1102.0 | 586970.0 |
| c1_power2 | pair_264 | 2184.0 | 586970.0 |
| c1_power2 | pair_265 | 2272.0 | 586970.0 |
| c1_power2 | pair_266 | 2254.0 | 586970.0 |
| c1_power2 | pair_267 | 2254.0 | 586970.0 |
| c1_power2 | pair_268 | 2263.0 | 586970.0 |
| c1_power2 | pair_269 | 1087.0 | 586970.0 |
| c1_power2 | pair_270 | 1087.0 | 586970.0 |
| c1_power2 | pair_271 | 2245.0 | 586970.0 |
| c1_power2 | pair_272 | 2272.0 | 586970.0 |
| c1_power2 | pair_273 | 2245.0 | 586970.0 |
| c1_power2 | pair_274 | 2254.0 | 586970.0 |
| c1_power2 | pair_275 | 2254.0 | 586970.0 |
| c1_power2 | pair_276 | 2236.0 | 586970.0 |
| c1_power2 | pair_277 | 2254.0 | 586970.0 |
| c1_power2 | pair_278 | 2236.0 | 586970.0 |
| c1_power2 | pair_279 | 2236.0 | 586970.0 |
| c1_power2 | pair_280 | 2245.0 | 586970.0 |
| c1_power2 | pair_281 | 1087.0 | 586970.0 |
| c1_power2 | pair_282 | 2209.0 | 586970.0 |
| c1_power2 | pair_283 | 2254.0 | 586970.0 |
| c1_power2 | pair_284 | 2254.0 | 586970.0 |
| c1_power2 | pair_285 | 2263.0 | 586970.0 |
| c1_power2 | pair_286 | 1112.0 | 586970.0 |
| c1_power2 | pair_287 | 1112.0 | 586970.0 |
| c1_power2 | pair_288 | 2245.0 | 586970.0 |
| c1_power2 | pair_289 | 2272.0 | 586970.0 |
| c1_power2 | pair_290 | 2245.0 | 586970.0 |
| c1_power2 | pair_291 | 2254.0 | 586970.0 |
| c1_power2 | pair_292 | 2254.0 | 586970.0 |
| c1_power2 | pair_293 | 2236.0 | 586970.0 |
| c1_power2 | pair_294 | 2254.0 | 586970.0 |
| c1_power2 | pair_295 | 2236.0 | 586970.0 |
| c1_power2 | pair_296 | 2236.0 | 586970.0 |
| c1_power2 | pair_297 | 2245.0 | 586970.0 |
| c1_power2 | pair_298 | 1112.0 | 586970.0 |
| c1_power2 | pair_299 | 2209.0 | 586970.0 |
| c1_power2 | pair_300 | 2184.0 | 586970.0 |
| c1_power2 | pair_301 | 2193.0 | 586970.0 |
| c1_power2 | pair_302 | 1102.0 | 586970.0 |
| c1_power2 | pair_303 | 1102.0 | 586970.0 |
| c1_power2 | pair_304 | 2175.0 | 586970.0 |
| c1_power2 | pair_305 | 2202.0 | 586970.0 |
| c1_power2 | pair_306 | 2175.0 | 586970.0 |
| c1_power2 | pair_307 | 2184.0 | 586970.0 |
| c1_power2 | pair_308 | 2184.0 | 586970.0 |
| c1_power2 | pair_309 | 2254.0 | 586970.0 |
| c1_power2 | pair_310 | 2184.0 | 586970.0 |
| c1_power2 | pair_311 | 2166.0 | 586970.0 |
| c1_power2 | pair_312 | 2254.0 | 586970.0 |
| c1_power2 | pair_313 | 2254.0 | 586970.0 |
| c1_power2 | pair_314 | 1102.0 | 586970.0 |
| c1_power2 | pair_315 | 2139.0 | 586970.0 |
| c1_power2 | pair_316 | 2193.0 | 586970.0 |
| c1_power2 | pair_317 | 1102.0 | 586970.0 |
| c1_power2 | pair_318 | 1102.0 | 586970.0 |
| c1_power2 | pair_319 | 2175.0 | 586970.0 |
| c1_power2 | pair_320 | 2202.0 | 586970.0 |
| c1_power2 | pair_321 | 2175.0 | 586970.0 |
| c1_power2 | pair_322 | 2184.0 | 586970.0 |
| c1_power2 | pair_323 | 2184.0 | 586970.0 |
| c1_power2 | pair_324 | 2254.0 | 586970.0 |
| c1_power2 | pair_325 | 2184.0 | 586970.0 |
| c1_power2 | pair_326 | 2166.0 | 586970.0 |
| c1_power2 | pair_327 | 2254.0 | 586970.0 |
| c1_power2 | pair_328 | 2254.0 | 586970.0 |
| c1_power2 | pair_329 | 1102.0 | 586970.0 |
| c1_power2 | pair_330 | 2139.0 | 586970.0 |
| c1_power2 | pair_331 | 1107.0 | 586970.0 |
| c1_power2 | pair_332 | 1107.0 | 586970.0 |
| c1_power2 | pair_333 | 2175.0 | 586970.0 |
| c1_power2 | pair_334 | 2202.0 | 586970.0 |
| c1_power2 | pair_335 | 2175.0 | 586970.0 |
| c1_power2 | pair_336 | 2184.0 | 586970.0 |
| c1_power2 | pair_337 | 2184.0 | 586970.0 |
| c1_power2 | pair_338 | 2263.0 | 586970.0 |
| c1_power2 | pair_339 | 2184.0 | 586970.0 |
| c1_power2 | pair_340 | 2166.0 | 586970.0 |
| c1_power2 | pair_341 | 2263.0 | 586970.0 |
| c1_power2 | pair_342 | 2263.0 | 586970.0 |
| c1_power2 | pair_343 | 1107.0 | 586970.0 |
| c1_power2 | pair_344 | 2139.0 | 586970.0 |
| c1_power2 | pair_345 | 1097.0 | 586970.0 |
| c1_power2 | pair_346 | 1097.0 | 586970.0 |
| c1_power2 | pair_347 | 1112.0 | 586970.0 |
| c1_power2 | pair_348 | 1097.0 | 586970.0 |
| c1_power2 | pair_349 | 1102.0 | 586970.0 |
| c1_power2 | pair_350 | 1102.0 | 586970.0 |
| c1_power2 | pair_351 | 1092.0 | 586970.0 |
| c1_power2 | pair_352 | 1102.0 | 586970.0 |
| c1_power2 | pair_353 | 1092.0 | 586970.0 |
| c1_power2 | pair_354 | 1092.0 | 586970.0 |
| c1_power2 | pair_355 | 1097.0 | 586970.0 |
| c1_power2 | pair_356 | 1127.0 | 586970.0 |
| c1_power2 | pair_357 | 1077.0 | 586970.0 |
| c1_power2 | pair_358 | 1097.0 | 586970.0 |
| c1_power2 | pair_359 | 1112.0 | 586970.0 |
| c1_power2 | pair_360 | 1097.0 | 586970.0 |
| c1_power2 | pair_361 | 1102.0 | 586970.0 |
| c1_power2 | pair_362 | 1102.0 | 586970.0 |
| c1_power2 | pair_363 | 1092.0 | 586970.0 |
| c1_power2 | pair_364 | 1102.0 | 586970.0 |
| c1_power2 | pair_365 | 1092.0 | 586970.0 |
| c1_power2 | pair_366 | 1092.0 | 586970.0 |
| c1_power2 | pair_367 | 1097.0 | 586970.0 |
| c1_power2 | pair_368 | 1127.0 | 586970.0 |
| c1_power2 | pair_369 | 1077.0 | 586970.0 |
| c1_power2 | pair_370 | 2105.0 | 586970.0 |
| c1_power2 | pair_371 | 4401.0 | 586970.0 |
| c1_power2 | pair_372 | 2105.0 | 586970.0 |
| c1_power2 | pair_373 | 4418.0 | 586970.0 |
| c1_power2 | pair_374 | 2245.0 | 586970.0 |
| c1_power2 | pair_375 | 4418.0 | 586970.0 |
| c1_power2 | pair_376 | 4384.0 | 586970.0 |
| c1_power2 | pair_377 | 2245.0 | 586970.0 |
| c1_power2 | pair_378 | 2245.0 | 586970.0 |
| c1_power2 | pair_379 | 1097.0 | 586970.0 |
| c1_power2 | pair_380 | 2175.0 | 586970.0 |
| c1_power2 | pair_381 | 2105.0 | 586970.0 |
| c1_power2 | pair_382 | 2114.0 | 586970.0 |
| c1_power2 | pair_383 | 2114.0 | 586970.0 |
| c1_power2 | pair_384 | 2272.0 | 586970.0 |
| c1_power2 | pair_385 | 2114.0 | 586970.0 |
| c1_power2 | pair_386 | 2096.0 | 586970.0 |
| c1_power2 | pair_387 | 2272.0 | 586970.0 |
| c1_power2 | pair_388 | 2272.0 | 586970.0 |
| c1_power2 | pair_389 | 1112.0 | 586970.0 |
| c1_power2 | pair_390 | 2202.0 | 586970.0 |
| c1_power2 | pair_391 | 2105.0 | 586970.0 |
| c1_power2 | pair_392 | 8923.0 | 586970.0 |
| c1_power2 | pair_393 | 2245.0 | 586970.0 |
| c1_power2 | pair_394 | 4331.0 | 586970.0 |
| c1_power2 | pair_395 | 8853.0 | 586970.0 |
| c1_power2 | pair_396 | 2245.0 | 586970.0 |
| c1_power2 | pair_397 | 2245.0 | 586970.0 |
| c1_power2 | pair_398 | 1097.0 | 586970.0 |
| c1_power2 | pair_399 | 2175.0 | 586970.0 |
| c1_power2 | pair_400 | 2114.0 | 586970.0 |
| c1_power2 | pair_401 | 2254.0 | 586970.0 |
| c1_power2 | pair_402 | 2114.0 | 586970.0 |
| c1_power2 | pair_403 | 2096.0 | 586970.0 |
| c1_power2 | pair_404 | 2254.0 | 586970.0 |
| c1_power2 | pair_405 | 2254.0 | 586970.0 |
| c1_power2 | pair_406 | 1102.0 | 586970.0 |
| c1_power2 | pair_407 | 2184.0 | 586970.0 |
| c1_power2 | pair_408 | 2254.0 | 586970.0 |
| c1_power2 | pair_409 | 4348.0 | 586970.0 |
| c1_power2 | pair_410 | 8890.0 | 586970.0 |
| c1_power2 | pair_411 | 2254.0 | 586970.0 |
| c1_power2 | pair_412 | 2254.0 | 586970.0 |
| c1_power2 | pair_413 | 1102.0 | 586970.0 |
| c1_power2 | pair_414 | 2184.0 | 586970.0 |
| c1_power2 | pair_415 | 2254.0 | 586970.0 |
| c1_power2 | pair_416 | 2236.0 | 586970.0 |
| c1_power2 | pair_417 | 2236.0 | 586970.0 |
| c1_power2 | pair_418 | 2245.0 | 586970.0 |
| c1_power2 | pair_419 | 1092.0 | 586970.0 |
| c1_power2 | pair_420 | 2209.0 | 586970.0 |
| c1_power2 | pair_421 | 4314.0 | 586970.0 |
| c1_power2 | pair_422 | 2254.0 | 586970.0 |
| c1_power2 | pair_423 | 2254.0 | 586970.0 |
| c1_power2 | pair_424 | 1102.0 | 586970.0 |
| c1_power2 | pair_425 | 2184.0 | 586970.0 |
| c1_power2 | pair_426 | 2236.0 | 586970.0 |
| c1_power2 | pair_427 | 2236.0 | 586970.0 |
| c1_power2 | pair_428 | 1092.0 | 586970.0 |
| c1_power2 | pair_429 | 2166.0 | 586970.0 |
| c1_power2 | pair_430 | 2245.0 | 586970.0 |
| c1_power2 | pair_431 | 1092.0 | 586970.0 |
| c1_power2 | pair_432 | 2209.0 | 586970.0 |
| c1_power2 | pair_433 | 1097.0 | 586970.0 |
| c1_power2 | pair_434 | 2209.0 | 586970.0 |
| c1_power2 | pair_435 | 1077.0 | 586970.0 |
| c1_sqrt | pair_001 | 5019.0 | 586970.0 |
| c1_sqrt | pair_002 | 2907.0 | 586970.0 |
| c1_sqrt | pair_003 | 3228.0 | 586970.0 |
| c1_sqrt | pair_004 | 3549.0 | 586970.0 |
| c1_sqrt | pair_005 | 1623.0 | 586970.0 |
| c1_sqrt | pair_006 | 2907.0 | 586970.0 |
| c1_sqrt | pair_007 | 6438.0 | 586970.0 |
| c1_sqrt | pair_008 | 3228.0 | 586970.0 |
| c1_sqrt | pair_009 | 8310.0 | 586970.0 |
| c1_sqrt | pair_010 | 3228.0 | 586970.0 |
| c1_sqrt | pair_011 | 2586.0 | 586970.0 |
| c1_sqrt | pair_012 | 2586.0 | 586970.0 |
| c1_sqrt | pair_013 | 2907.0 | 586970.0 |
| c1_sqrt | pair_014 | 2907.0 | 586970.0 |
| c1_sqrt | pair_015 | 2907.0 | 586970.0 |
| c1_sqrt | pair_016 | 1623.0 | 586970.0 |
| c1_sqrt | pair_017 | 1623.0 | 586970.0 |
| c1_sqrt | pair_018 | 4833.0 | 586970.0 |
| c1_sqrt | pair_019 | 3228.0 | 586970.0 |
| c1_sqrt | pair_020 | 8601.0 | 586970.0 |
| c1_sqrt | pair_021 | 3228.0 | 586970.0 |
| c1_sqrt | pair_022 | 8043.0 | 586970.0 |
| c1_sqrt | pair_023 | 2586.0 | 586970.0 |
| c1_sqrt | pair_024 | 5154.0 | 586970.0 |
| c1_sqrt | pair_025 | 8364.0 | 586970.0 |
| c1_sqrt | pair_026 | 2586.0 | 586970.0 |
| c1_sqrt | pair_027 | 2586.0 | 586970.0 |
| c1_sqrt | pair_028 | 1623.0 | 586970.0 |
| c1_sqrt | pair_029 | 2907.0 | 586970.0 |
| c1_sqrt | pair_030 | 2457.0 | 586970.0 |
| c1_sqrt | pair_031 | 2705.0 | 586970.0 |
| c1_sqrt | pair_032 | 3085.0 | 586970.0 |
| c1_sqrt | pair_033 | 1395.0 | 586970.0 |
| c1_sqrt | pair_034 | 2457.0 | 586970.0 |
| c1_sqrt | pair_035 | 3013.0 | 586970.0 |
| c1_sqrt | pair_036 | 2705.0 | 586970.0 |
| c1_sqrt | pair_037 | 2965.0 | 586970.0 |
| c1_sqrt | pair_038 | 2705.0 | 586970.0 |
| c1_sqrt | pair_039 | 2209.0 | 586970.0 |
| c1_sqrt | pair_040 | 2209.0 | 586970.0 |
| c1_sqrt | pair_041 | 2457.0 | 586970.0 |
| c1_sqrt | pair_042 | 2457.0 | 586970.0 |
| c1_sqrt | pair_043 | 2457.0 | 586970.0 |
| c1_sqrt | pair_044 | 1395.0 | 586970.0 |
| c1_sqrt | pair_045 | 1395.0 | 586970.0 |
| c1_sqrt | pair_046 | 3001.0 | 586970.0 |
| c1_sqrt | pair_047 | 2705.0 | 586970.0 |
| c1_sqrt | pair_048 | 3001.0 | 586970.0 |
| c1_sqrt | pair_049 | 2705.0 | 586970.0 |
| c1_sqrt | pair_050 | 3013.0 | 586970.0 |
| c1_sqrt | pair_051 | 2209.0 | 586970.0 |
| c1_sqrt | pair_052 | 3013.0 | 586970.0 |
| c1_sqrt | pair_053 | 2989.0 | 586970.0 |
| c1_sqrt | pair_054 | 2209.0 | 586970.0 |
| c1_sqrt | pair_055 | 2209.0 | 586970.0 |
| c1_sqrt | pair_056 | 1395.0 | 586970.0 |
| c1_sqrt | pair_057 | 2457.0 | 586970.0 |
| c1_sqrt | pair_058 | 2487.0 | 586970.0 |
| c1_sqrt | pair_059 | 2567.0 | 586970.0 |
| c1_sqrt | pair_060 | 1443.0 | 586970.0 |
| c1_sqrt | pair_061 | 2477.0 | 586970.0 |
| c1_sqrt | pair_062 | 2507.0 | 586970.0 |
| c1_sqrt | pair_063 | 2527.0 | 586970.0 |
| c1_sqrt | pair_064 | 2467.0 | 586970.0 |
| c1_sqrt | pair_065 | 2507.0 | 586970.0 |
| c1_sqrt | pair_066 | 2281.0 | 586970.0 |
| c1_sqrt | pair_067 | 2281.0 | 586970.0 |
| c1_sqrt | pair_068 | 2507.0 | 586970.0 |
| c1_sqrt | pair_069 | 2507.0 | 586970.0 |
| c1_sqrt | pair_070 | 2517.0 | 586970.0 |
| c1_sqrt | pair_071 | 1443.0 | 586970.0 |
| c1_sqrt | pair_072 | 1443.0 | 586970.0 |
| c1_sqrt | pair_073 | 2497.0 | 586970.0 |
| c1_sqrt | pair_074 | 2527.0 | 586970.0 |
| c1_sqrt | pair_075 | 2497.0 | 586970.0 |
| c1_sqrt | pair_076 | 2507.0 | 586970.0 |
| c1_sqrt | pair_077 | 2507.0 | 586970.0 |
| c1_sqrt | pair_078 | 2281.0 | 586970.0 |
| c1_sqrt | pair_079 | 2507.0 | 586970.0 |
| c1_sqrt | pair_080 | 2487.0 | 586970.0 |
| c1_sqrt | pair_081 | 2281.0 | 586970.0 |
| c1_sqrt | pair_082 | 2281.0 | 586970.0 |
| c1_sqrt | pair_083 | 1443.0 | 586970.0 |
| c1_sqrt | pair_084 | 2457.0 | 586970.0 |
| c1_sqrt | pair_085 | 2826.0 | 586970.0 |
| c1_sqrt | pair_086 | 1413.0 | 586970.0 |
| c1_sqrt | pair_087 | 2487.0 | 586970.0 |
| c1_sqrt | pair_088 | 2760.0 | 586970.0 |
| c1_sqrt | pair_089 | 2782.0 | 586970.0 |
| c1_sqrt | pair_090 | 2716.0 | 586970.0 |
| c1_sqrt | pair_091 | 2760.0 | 586970.0 |
| c1_sqrt | pair_092 | 2236.0 | 586970.0 |
| c1_sqrt | pair_093 | 2236.0 | 586970.0 |
| c1_sqrt | pair_094 | 2487.0 | 586970.0 |
| c1_sqrt | pair_095 | 2487.0 | 586970.0 |
| c1_sqrt | pair_096 | 2487.0 | 586970.0 |
| c1_sqrt | pair_097 | 1413.0 | 586970.0 |
| c1_sqrt | pair_098 | 1413.0 | 586970.0 |
| c1_sqrt | pair_099 | 2749.0 | 586970.0 |
| c1_sqrt | pair_100 | 2782.0 | 586970.0 |
| c1_sqrt | pair_101 | 2749.0 | 586970.0 |
| c1_sqrt | pair_102 | 2760.0 | 586970.0 |
| c1_sqrt | pair_103 | 2760.0 | 586970.0 |
| c1_sqrt | pair_104 | 2236.0 | 586970.0 |
| c1_sqrt | pair_105 | 2760.0 | 586970.0 |
| c1_sqrt | pair_106 | 2738.0 | 586970.0 |
| c1_sqrt | pair_107 | 2236.0 | 586970.0 |
| c1_sqrt | pair_108 | 2236.0 | 586970.0 |
| c1_sqrt | pair_109 | 1413.0 | 586970.0 |
| c1_sqrt | pair_110 | 2487.0 | 586970.0 |
| c1_sqrt | pair_111 | 1461.0 | 586970.0 |
| c1_sqrt | pair_112 | 2567.0 | 586970.0 |
| c1_sqrt | pair_113 | 3013.0 | 586970.0 |
| c1_sqrt | pair_114 | 2826.0 | 586970.0 |
| c1_sqrt | pair_115 | 2965.0 | 586970.0 |
| c1_sqrt | pair_116 | 2826.0 | 586970.0 |
| c1_sqrt | pair_117 | 2308.0 | 586970.0 |
| c1_sqrt | pair_118 | 2308.0 | 586970.0 |
| c1_sqrt | pair_119 | 2567.0 | 586970.0 |
| c1_sqrt | pair_120 | 2567.0 | 586970.0 |
| c1_sqrt | pair_121 | 2567.0 | 586970.0 |
| c1_sqrt | pair_122 | 1461.0 | 586970.0 |
| c1_sqrt | pair_123 | 1461.0 | 586970.0 |
| c1_sqrt | pair_124 | 3001.0 | 586970.0 |
| c1_sqrt | pair_125 | 2826.0 | 586970.0 |
| c1_sqrt | pair_126 | 3001.0 | 586970.0 |
| c1_sqrt | pair_127 | 2826.0 | 586970.0 |
| c1_sqrt | pair_128 | 3013.0 | 586970.0 |
| c1_sqrt | pair_129 | 2308.0 | 586970.0 |
| c1_sqrt | pair_130 | 3013.0 | 586970.0 |
| c1_sqrt | pair_131 | 2989.0 | 586970.0 |
| c1_sqrt | pair_132 | 2308.0 | 586970.0 |
| c1_sqrt | pair_133 | 2308.0 | 586970.0 |
| c1_sqrt | pair_134 | 1461.0 | 586970.0 |
| c1_sqrt | pair_135 | 2567.0 | 586970.0 |
| c1_sqrt | pair_136 | 1407.0 | 586970.0 |
| c1_sqrt | pair_137 | 1425.0 | 586970.0 |
| c1_sqrt | pair_138 | 1437.0 | 586970.0 |
| c1_sqrt | pair_139 | 1401.0 | 586970.0 |
| c1_sqrt | pair_140 | 1425.0 | 586970.0 |
| c1_sqrt | pair_141 | 1407.0 | 586970.0 |
| c1_sqrt | pair_142 | 1437.0 | 586970.0 |
| c1_sqrt | pair_143 | 1425.0 | 586970.0 |
| c1_sqrt | pair_144 | 1425.0 | 586970.0 |
| c1_sqrt | pair_145 | 1431.0 | 586970.0 |
| c1_sqrt | pair_146 | 1425.0 | 586970.0 |
| c1_sqrt | pair_147 | 1419.0 | 586970.0 |
| c1_sqrt | pair_148 | 1419.0 | 586970.0 |
| c1_sqrt | pair_149 | 1437.0 | 586970.0 |
| c1_sqrt | pair_150 | 1419.0 | 586970.0 |
| c1_sqrt | pair_151 | 1425.0 | 586970.0 |
| c1_sqrt | pair_152 | 1425.0 | 586970.0 |
| c1_sqrt | pair_153 | 1413.0 | 586970.0 |
| c1_sqrt | pair_154 | 1425.0 | 586970.0 |
| c1_sqrt | pair_155 | 1413.0 | 586970.0 |
| c1_sqrt | pair_156 | 1413.0 | 586970.0 |
| c1_sqrt | pair_157 | 1419.0 | 586970.0 |
| c1_sqrt | pair_158 | 1455.0 | 586970.0 |
| c1_sqrt | pair_159 | 1395.0 | 586970.0 |
| c1_sqrt | pair_160 | 2507.0 | 586970.0 |
| c1_sqrt | pair_161 | 2527.0 | 586970.0 |
| c1_sqrt | pair_162 | 2467.0 | 586970.0 |
| c1_sqrt | pair_163 | 2507.0 | 586970.0 |
| c1_sqrt | pair_164 | 2227.0 | 586970.0 |
| c1_sqrt | pair_165 | 2227.0 | 586970.0 |
| c1_sqrt | pair_166 | 2507.0 | 586970.0 |
| c1_sqrt | pair_167 | 2507.0 | 586970.0 |
| c1_sqrt | pair_168 | 2517.0 | 586970.0 |
| c1_sqrt | pair_169 | 1407.0 | 586970.0 |
| c1_sqrt | pair_170 | 1407.0 | 586970.0 |
| c1_sqrt | pair_171 | 2497.0 | 586970.0 |
| c1_sqrt | pair_172 | 2527.0 | 586970.0 |
| c1_sqrt | pair_173 | 2497.0 | 586970.0 |
| c1_sqrt | pair_174 | 2507.0 | 586970.0 |
| c1_sqrt | pair_175 | 2507.0 | 586970.0 |
| c1_sqrt | pair_176 | 2227.0 | 586970.0 |
| c1_sqrt | pair_177 | 2507.0 | 586970.0 |
| c1_sqrt | pair_178 | 2487.0 | 586970.0 |
| c1_sqrt | pair_179 | 2227.0 | 586970.0 |
| c1_sqrt | pair_180 | 2227.0 | 586970.0 |
| c1_sqrt | pair_181 | 1407.0 | 586970.0 |
| c1_sqrt | pair_182 | 2457.0 | 586970.0 |
| c1_sqrt | pair_183 | 2760.0 | 586970.0 |
| c1_sqrt | pair_184 | 5276.0 | 586970.0 |
| c1_sqrt | pair_185 | 2760.0 | 586970.0 |
| c1_sqrt | pair_186 | 2254.0 | 586970.0 |
| c1_sqrt | pair_187 | 2254.0 | 586970.0 |
| c1_sqrt | pair_188 | 2507.0 | 586970.0 |
| c1_sqrt | pair_189 | 2507.0 | 586970.0 |
| c1_sqrt | pair_190 | 2507.0 | 586970.0 |
| c1_sqrt | pair_191 | 1425.0 | 586970.0 |
| c1_sqrt | pair_192 | 1425.0 | 586970.0 |
| c1_sqrt | pair_193 | 4095.0 | 586970.0 |
| c1_sqrt | pair_194 | 2760.0 | 586970.0 |
| c1_sqrt | pair_195 | 5339.0 | 586970.0 |
| c1_sqrt | pair_196 | 2760.0 | 586970.0 |
| c1_sqrt | pair_197 | 5360.0 | 586970.0 |
| c1_sqrt | pair_198 | 2254.0 | 586970.0 |
| c1_sqrt | pair_199 | 4348.0 | 586970.0 |
| c1_sqrt | pair_200 | 5318.0 | 586970.0 |
| c1_sqrt | pair_201 | 2254.0 | 586970.0 |
| c1_sqrt | pair_202 | 2254.0 | 586970.0 |
| c1_sqrt | pair_203 | 1425.0 | 586970.0 |
| c1_sqrt | pair_204 | 2507.0 | 586970.0 |
| c1_sqrt | pair_205 | 2716.0 | 586970.0 |
| c1_sqrt | pair_206 | 2760.0 | 586970.0 |
| c1_sqrt | pair_207 | 2272.0 | 586970.0 |
| c1_sqrt | pair_208 | 2272.0 | 586970.0 |
| c1_sqrt | pair_209 | 2527.0 | 586970.0 |
| c1_sqrt | pair_210 | 2527.0 | 586970.0 |
| c1_sqrt | pair_211 | 2527.0 | 586970.0 |
| c1_sqrt | pair_212 | 1437.0 | 586970.0 |
| c1_sqrt | pair_213 | 1437.0 | 586970.0 |
| c1_sqrt | pair_214 | 2749.0 | 586970.0 |
| c1_sqrt | pair_215 | 2782.0 | 586970.0 |
| c1_sqrt | pair_216 | 2749.0 | 586970.0 |
| c1_sqrt | pair_217 | 2760.0 | 586970.0 |
| c1_sqrt | pair_218 | 2760.0 | 586970.0 |
| c1_sqrt | pair_219 | 2272.0 | 586970.0 |
| c1_sqrt | pair_220 | 2760.0 | 586970.0 |
| c1_sqrt | pair_221 | 2738.0 | 586970.0 |
| c1_sqrt | pair_222 | 2272.0 | 586970.0 |
| c1_sqrt | pair_223 | 2272.0 | 586970.0 |
| c1_sqrt | pair_224 | 1437.0 | 586970.0 |
| c1_sqrt | pair_225 | 2527.0 | 586970.0 |
| c1_sqrt | pair_226 | 2716.0 | 586970.0 |
| c1_sqrt | pair_227 | 2218.0 | 586970.0 |
| c1_sqrt | pair_228 | 2218.0 | 586970.0 |
| c1_sqrt | pair_229 | 2467.0 | 586970.0 |
| c1_sqrt | pair_230 | 2467.0 | 586970.0 |
| c1_sqrt | pair_231 | 2467.0 | 586970.0 |
| c1_sqrt | pair_232 | 1401.0 | 586970.0 |
| c1_sqrt | pair_233 | 1401.0 | 586970.0 |
| c1_sqrt | pair_234 | 4031.0 | 586970.0 |
| c1_sqrt | pair_235 | 2716.0 | 586970.0 |
| c1_sqrt | pair_236 | 6921.0 | 586970.0 |
| c1_sqrt | pair_237 | 2716.0 | 586970.0 |
| c1_sqrt | pair_238 | 6591.0 | 586970.0 |
| c1_sqrt | pair_239 | 2218.0 | 586970.0 |
| c1_sqrt | pair_240 | 4280.0 | 586970.0 |
| c1_sqrt | pair_241 | 6894.0 | 586970.0 |
| c1_sqrt | pair_242 | 2218.0 | 586970.0 |
| c1_sqrt | pair_243 | 2218.0 | 586970.0 |
| c1_sqrt | pair_244 | 1401.0 | 586970.0 |
| c1_sqrt | pair_245 | 2467.0 | 586970.0 |
| c1_sqrt | pair_246 | 2254.0 | 586970.0 |
| c1_sqrt | pair_247 | 2254.0 | 586970.0 |
| c1_sqrt | pair_248 | 2507.0 | 586970.0 |
| c1_sqrt | pair_249 | 2507.0 | 586970.0 |
| c1_sqrt | pair_250 | 2507.0 | 586970.0 |
| c1_sqrt | pair_251 | 1425.0 | 586970.0 |
| c1_sqrt | pair_252 | 1425.0 | 586970.0 |
| c1_sqrt | pair_253 | 2749.0 | 586970.0 |
| c1_sqrt | pair_254 | 2782.0 | 586970.0 |
| c1_sqrt | pair_255 | 2749.0 | 586970.0 |
| c1_sqrt | pair_256 | 2760.0 | 586970.0 |
| c1_sqrt | pair_257 | 2760.0 | 586970.0 |
| c1_sqrt | pair_258 | 2254.0 | 586970.0 |
| c1_sqrt | pair_259 | 2760.0 | 586970.0 |
| c1_sqrt | pair_260 | 2738.0 | 586970.0 |
| c1_sqrt | pair_261 | 2254.0 | 586970.0 |
| c1_sqrt | pair_262 | 2254.0 | 586970.0 |
| c1_sqrt | pair_263 | 1425.0 | 586970.0 |
| c1_sqrt | pair_264 | 2507.0 | 586970.0 |
| c1_sqrt | pair_265 | 2272.0 | 586970.0 |
| c1_sqrt | pair_266 | 2254.0 | 586970.0 |
| c1_sqrt | pair_267 | 2254.0 | 586970.0 |
| c1_sqrt | pair_268 | 2263.0 | 586970.0 |
| c1_sqrt | pair_269 | 1407.0 | 586970.0 |
| c1_sqrt | pair_270 | 1407.0 | 586970.0 |
| c1_sqrt | pair_271 | 2245.0 | 586970.0 |
| c1_sqrt | pair_272 | 2272.0 | 586970.0 |
| c1_sqrt | pair_273 | 2245.0 | 586970.0 |
| c1_sqrt | pair_274 | 2254.0 | 586970.0 |
| c1_sqrt | pair_275 | 2254.0 | 586970.0 |
| c1_sqrt | pair_276 | 2236.0 | 586970.0 |
| c1_sqrt | pair_277 | 2254.0 | 586970.0 |
| c1_sqrt | pair_278 | 2236.0 | 586970.0 |
| c1_sqrt | pair_279 | 2236.0 | 586970.0 |
| c1_sqrt | pair_280 | 2245.0 | 586970.0 |
| c1_sqrt | pair_281 | 1407.0 | 586970.0 |
| c1_sqrt | pair_282 | 2209.0 | 586970.0 |
| c1_sqrt | pair_283 | 2254.0 | 586970.0 |
| c1_sqrt | pair_284 | 2254.0 | 586970.0 |
| c1_sqrt | pair_285 | 2263.0 | 586970.0 |
| c1_sqrt | pair_286 | 1437.0 | 586970.0 |
| c1_sqrt | pair_287 | 1437.0 | 586970.0 |
| c1_sqrt | pair_288 | 2245.0 | 586970.0 |
| c1_sqrt | pair_289 | 2272.0 | 586970.0 |
| c1_sqrt | pair_290 | 2245.0 | 586970.0 |
| c1_sqrt | pair_291 | 2254.0 | 586970.0 |
| c1_sqrt | pair_292 | 2254.0 | 586970.0 |
| c1_sqrt | pair_293 | 2236.0 | 586970.0 |
| c1_sqrt | pair_294 | 2254.0 | 586970.0 |
| c1_sqrt | pair_295 | 2236.0 | 586970.0 |
| c1_sqrt | pair_296 | 2236.0 | 586970.0 |
| c1_sqrt | pair_297 | 2245.0 | 586970.0 |
| c1_sqrt | pair_298 | 1437.0 | 586970.0 |
| c1_sqrt | pair_299 | 2209.0 | 586970.0 |
| c1_sqrt | pair_300 | 2507.0 | 586970.0 |
| c1_sqrt | pair_301 | 2517.0 | 586970.0 |
| c1_sqrt | pair_302 | 1425.0 | 586970.0 |
| c1_sqrt | pair_303 | 1425.0 | 586970.0 |
| c1_sqrt | pair_304 | 2497.0 | 586970.0 |
| c1_sqrt | pair_305 | 2527.0 | 586970.0 |
| c1_sqrt | pair_306 | 2497.0 | 586970.0 |
| c1_sqrt | pair_307 | 2507.0 | 586970.0 |
| c1_sqrt | pair_308 | 2507.0 | 586970.0 |
| c1_sqrt | pair_309 | 2254.0 | 586970.0 |
| c1_sqrt | pair_310 | 2507.0 | 586970.0 |
| c1_sqrt | pair_311 | 2487.0 | 586970.0 |
| c1_sqrt | pair_312 | 2254.0 | 586970.0 |
| c1_sqrt | pair_313 | 2254.0 | 586970.0 |
| c1_sqrt | pair_314 | 1425.0 | 586970.0 |
| c1_sqrt | pair_315 | 2457.0 | 586970.0 |
| c1_sqrt | pair_316 | 2517.0 | 586970.0 |
| c1_sqrt | pair_317 | 1425.0 | 586970.0 |
| c1_sqrt | pair_318 | 1425.0 | 586970.0 |
| c1_sqrt | pair_319 | 2497.0 | 586970.0 |
| c1_sqrt | pair_320 | 2527.0 | 586970.0 |
| c1_sqrt | pair_321 | 2497.0 | 586970.0 |
| c1_sqrt | pair_322 | 2507.0 | 586970.0 |
| c1_sqrt | pair_323 | 2507.0 | 586970.0 |
| c1_sqrt | pair_324 | 2254.0 | 586970.0 |
| c1_sqrt | pair_325 | 2507.0 | 586970.0 |
| c1_sqrt | pair_326 | 2487.0 | 586970.0 |
| c1_sqrt | pair_327 | 2254.0 | 586970.0 |
| c1_sqrt | pair_328 | 2254.0 | 586970.0 |
| c1_sqrt | pair_329 | 1425.0 | 586970.0 |
| c1_sqrt | pair_330 | 2457.0 | 586970.0 |
| c1_sqrt | pair_331 | 1431.0 | 586970.0 |
| c1_sqrt | pair_332 | 1431.0 | 586970.0 |
| c1_sqrt | pair_333 | 2497.0 | 586970.0 |
| c1_sqrt | pair_334 | 2527.0 | 586970.0 |
| c1_sqrt | pair_335 | 2497.0 | 586970.0 |
| c1_sqrt | pair_336 | 2507.0 | 586970.0 |
| c1_sqrt | pair_337 | 2507.0 | 586970.0 |
| c1_sqrt | pair_338 | 2263.0 | 586970.0 |
| c1_sqrt | pair_339 | 2507.0 | 586970.0 |
| c1_sqrt | pair_340 | 2487.0 | 586970.0 |
| c1_sqrt | pair_341 | 2263.0 | 586970.0 |
| c1_sqrt | pair_342 | 2263.0 | 586970.0 |
| c1_sqrt | pair_343 | 1431.0 | 586970.0 |
| c1_sqrt | pair_344 | 2457.0 | 586970.0 |
| c1_sqrt | pair_345 | 1419.0 | 586970.0 |
| c1_sqrt | pair_346 | 1419.0 | 586970.0 |
| c1_sqrt | pair_347 | 1437.0 | 586970.0 |
| c1_sqrt | pair_348 | 1419.0 | 586970.0 |
| c1_sqrt | pair_349 | 1425.0 | 586970.0 |
| c1_sqrt | pair_350 | 1425.0 | 586970.0 |
| c1_sqrt | pair_351 | 1413.0 | 586970.0 |
| c1_sqrt | pair_352 | 1425.0 | 586970.0 |
| c1_sqrt | pair_353 | 1413.0 | 586970.0 |
| c1_sqrt | pair_354 | 1413.0 | 586970.0 |
| c1_sqrt | pair_355 | 1419.0 | 586970.0 |
| c1_sqrt | pair_356 | 1455.0 | 586970.0 |
| c1_sqrt | pair_357 | 1395.0 | 586970.0 |
| c1_sqrt | pair_358 | 1419.0 | 586970.0 |
| c1_sqrt | pair_359 | 1437.0 | 586970.0 |
| c1_sqrt | pair_360 | 1419.0 | 586970.0 |
| c1_sqrt | pair_361 | 1425.0 | 586970.0 |
| c1_sqrt | pair_362 | 1425.0 | 586970.0 |
| c1_sqrt | pair_363 | 1413.0 | 586970.0 |
| c1_sqrt | pair_364 | 1425.0 | 586970.0 |
| c1_sqrt | pair_365 | 1413.0 | 586970.0 |
| c1_sqrt | pair_366 | 1413.0 | 586970.0 |
| c1_sqrt | pair_367 | 1419.0 | 586970.0 |
| c1_sqrt | pair_368 | 1455.0 | 586970.0 |
| c1_sqrt | pair_369 | 1395.0 | 586970.0 |
| c1_sqrt | pair_370 | 2749.0 | 586970.0 |
| c1_sqrt | pair_371 | 4079.0 | 586970.0 |
| c1_sqrt | pair_372 | 2749.0 | 586970.0 |
| c1_sqrt | pair_373 | 4095.0 | 586970.0 |
| c1_sqrt | pair_374 | 2245.0 | 586970.0 |
| c1_sqrt | pair_375 | 4095.0 | 586970.0 |
| c1_sqrt | pair_376 | 4063.0 | 586970.0 |
| c1_sqrt | pair_377 | 2245.0 | 586970.0 |
| c1_sqrt | pair_378 | 2245.0 | 586970.0 |
| c1_sqrt | pair_379 | 1419.0 | 586970.0 |
| c1_sqrt | pair_380 | 2497.0 | 586970.0 |
| c1_sqrt | pair_381 | 2749.0 | 586970.0 |
| c1_sqrt | pair_382 | 2760.0 | 586970.0 |
| c1_sqrt | pair_383 | 2760.0 | 586970.0 |
| c1_sqrt | pair_384 | 2272.0 | 586970.0 |
| c1_sqrt | pair_385 | 2760.0 | 586970.0 |
| c1_sqrt | pair_386 | 2738.0 | 586970.0 |
| c1_sqrt | pair_387 | 2272.0 | 586970.0 |
| c1_sqrt | pair_388 | 2272.0 | 586970.0 |
| c1_sqrt | pair_389 | 1437.0 | 586970.0 |
| c1_sqrt | pair_390 | 2527.0 | 586970.0 |
| c1_sqrt | pair_391 | 2749.0 | 586970.0 |
| c1_sqrt | pair_392 | 6669.0 | 586970.0 |
| c1_sqrt | pair_393 | 2245.0 | 586970.0 |
| c1_sqrt | pair_394 | 4331.0 | 586970.0 |
| c1_sqrt | pair_395 | 6921.0 | 586970.0 |
| c1_sqrt | pair_396 | 2245.0 | 586970.0 |
| c1_sqrt | pair_397 | 2245.0 | 586970.0 |
| c1_sqrt | pair_398 | 1419.0 | 586970.0 |
| c1_sqrt | pair_399 | 2497.0 | 586970.0 |
| c1_sqrt | pair_400 | 2760.0 | 586970.0 |
| c1_sqrt | pair_401 | 2254.0 | 586970.0 |
| c1_sqrt | pair_402 | 2760.0 | 586970.0 |
| c1_sqrt | pair_403 | 2738.0 | 586970.0 |
| c1_sqrt | pair_404 | 2254.0 | 586970.0 |
| c1_sqrt | pair_405 | 2254.0 | 586970.0 |
| c1_sqrt | pair_406 | 1425.0 | 586970.0 |
| c1_sqrt | pair_407 | 2507.0 | 586970.0 |
| c1_sqrt | pair_408 | 2254.0 | 586970.0 |
| c1_sqrt | pair_409 | 4348.0 | 586970.0 |
| c1_sqrt | pair_410 | 6643.0 | 586970.0 |
| c1_sqrt | pair_411 | 2254.0 | 586970.0 |
| c1_sqrt | pair_412 | 2254.0 | 586970.0 |
| c1_sqrt | pair_413 | 1425.0 | 586970.0 |
| c1_sqrt | pair_414 | 2507.0 | 586970.0 |
| c1_sqrt | pair_415 | 2254.0 | 586970.0 |
| c1_sqrt | pair_416 | 2236.0 | 586970.0 |
| c1_sqrt | pair_417 | 2236.0 | 586970.0 |
| c1_sqrt | pair_418 | 2245.0 | 586970.0 |
| c1_sqrt | pair_419 | 1413.0 | 586970.0 |
| c1_sqrt | pair_420 | 2209.0 | 586970.0 |
| c1_sqrt | pair_421 | 4314.0 | 586970.0 |
| c1_sqrt | pair_422 | 2254.0 | 586970.0 |
| c1_sqrt | pair_423 | 2254.0 | 586970.0 |
| c1_sqrt | pair_424 | 1425.0 | 586970.0 |
| c1_sqrt | pair_425 | 2507.0 | 586970.0 |
| c1_sqrt | pair_426 | 2236.0 | 586970.0 |
| c1_sqrt | pair_427 | 2236.0 | 586970.0 |
| c1_sqrt | pair_428 | 1413.0 | 586970.0 |
| c1_sqrt | pair_429 | 2487.0 | 586970.0 |
| c1_sqrt | pair_430 | 2245.0 | 586970.0 |
| c1_sqrt | pair_431 | 1413.0 | 586970.0 |
| c1_sqrt | pair_432 | 2209.0 | 586970.0 |
| c1_sqrt | pair_433 | 1419.0 | 586970.0 |
| c1_sqrt | pair_434 | 2209.0 | 586970.0 |
| c1_sqrt | pair_435 | 1395.0 | 586970.0 |
| corrected_fdxt | pair_001 | 4056.0 | 586970.0 |
| corrected_fdxt | pair_002 | 1944.0 | 586970.0 |
| corrected_fdxt | pair_003 | 2265.0 | 586970.0 |
| corrected_fdxt | pair_004 | 2586.0 | 586970.0 |
| corrected_fdxt | pair_005 | 981.0 | 586970.0 |
| corrected_fdxt | pair_006 | 1944.0 | 586970.0 |
| corrected_fdxt | pair_007 | 5154.0 | 586970.0 |
| corrected_fdxt | pair_008 | 2265.0 | 586970.0 |
| corrected_fdxt | pair_009 | 6715.0 | 586970.0 |
| corrected_fdxt | pair_010 | 2265.0 | 586970.0 |
| corrected_fdxt | pair_011 | 1623.0 | 586970.0 |
| corrected_fdxt | pair_012 | 1623.0 | 586970.0 |
| corrected_fdxt | pair_013 | 1944.0 | 586970.0 |
| corrected_fdxt | pair_014 | 1944.0 | 586970.0 |
| corrected_fdxt | pair_015 | 1944.0 | 586970.0 |
| corrected_fdxt | pair_016 | 981.0 | 586970.0 |
| corrected_fdxt | pair_017 | 981.0 | 586970.0 |
| corrected_fdxt | pair_018 | 3549.0 | 586970.0 |
| corrected_fdxt | pair_019 | 2265.0 | 586970.0 |
| corrected_fdxt | pair_020 | 6991.0 | 586970.0 |
| corrected_fdxt | pair_021 | 2265.0 | 586970.0 |
| corrected_fdxt | pair_022 | 6438.0 | 586970.0 |
| corrected_fdxt | pair_023 | 1623.0 | 586970.0 |
| corrected_fdxt | pair_024 | 3870.0 | 586970.0 |
| corrected_fdxt | pair_025 | 6759.0 | 586970.0 |
| corrected_fdxt | pair_026 | 1623.0 | 586970.0 |
| corrected_fdxt | pair_027 | 1623.0 | 586970.0 |
| corrected_fdxt | pair_028 | 981.0 | 586970.0 |
| corrected_fdxt | pair_029 | 1944.0 | 586970.0 |
| corrected_fdxt | pair_030 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_031 | 1751.0 | 586970.0 |
| corrected_fdxt | pair_032 | 2098.0 | 586970.0 |
| corrected_fdxt | pair_033 | 759.0 | 586970.0 |
| corrected_fdxt | pair_034 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_035 | 2044.0 | 586970.0 |
| corrected_fdxt | pair_036 | 1751.0 | 586970.0 |
| corrected_fdxt | pair_037 | 2008.0 | 586970.0 |
| corrected_fdxt | pair_038 | 1751.0 | 586970.0 |
| corrected_fdxt | pair_039 | 1255.0 | 586970.0 |
| corrected_fdxt | pair_040 | 1255.0 | 586970.0 |
| corrected_fdxt | pair_041 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_042 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_043 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_044 | 759.0 | 586970.0 |
| corrected_fdxt | pair_045 | 759.0 | 586970.0 |
| corrected_fdxt | pair_046 | 2035.0 | 586970.0 |
| corrected_fdxt | pair_047 | 1751.0 | 586970.0 |
| corrected_fdxt | pair_048 | 2035.0 | 586970.0 |
| corrected_fdxt | pair_049 | 1751.0 | 586970.0 |
| corrected_fdxt | pair_050 | 2044.0 | 586970.0 |
| corrected_fdxt | pair_051 | 1255.0 | 586970.0 |
| corrected_fdxt | pair_052 | 2044.0 | 586970.0 |
| corrected_fdxt | pair_053 | 2026.0 | 586970.0 |
| corrected_fdxt | pair_054 | 1255.0 | 586970.0 |
| corrected_fdxt | pair_055 | 1255.0 | 586970.0 |
| corrected_fdxt | pair_056 | 759.0 | 586970.0 |
| corrected_fdxt | pair_057 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_058 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_059 | 1580.0 | 586970.0 |
| corrected_fdxt | pair_060 | 791.0 | 586970.0 |
| corrected_fdxt | pair_061 | 1517.0 | 586970.0 |
| corrected_fdxt | pair_062 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_063 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_064 | 1510.0 | 586970.0 |
| corrected_fdxt | pair_065 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_066 | 1303.0 | 586970.0 |
| corrected_fdxt | pair_067 | 1303.0 | 586970.0 |
| corrected_fdxt | pair_068 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_069 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_070 | 1545.0 | 586970.0 |
| corrected_fdxt | pair_071 | 791.0 | 586970.0 |
| corrected_fdxt | pair_072 | 791.0 | 586970.0 |
| corrected_fdxt | pair_073 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_074 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_075 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_076 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_077 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_078 | 1303.0 | 586970.0 |
| corrected_fdxt | pair_079 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_080 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_081 | 1303.0 | 586970.0 |
| corrected_fdxt | pair_082 | 1303.0 | 586970.0 |
| corrected_fdxt | pair_083 | 791.0 | 586970.0 |
| corrected_fdxt | pair_084 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_085 | 1839.0 | 586970.0 |
| corrected_fdxt | pair_086 | 771.0 | 586970.0 |
| corrected_fdxt | pair_087 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_088 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_089 | 1807.0 | 586970.0 |
| corrected_fdxt | pair_090 | 1759.0 | 586970.0 |
| corrected_fdxt | pair_091 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_092 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_093 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_094 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_095 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_096 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_097 | 771.0 | 586970.0 |
| corrected_fdxt | pair_098 | 771.0 | 586970.0 |
| corrected_fdxt | pair_099 | 1783.0 | 586970.0 |
| corrected_fdxt | pair_100 | 1807.0 | 586970.0 |
| corrected_fdxt | pair_101 | 1783.0 | 586970.0 |
| corrected_fdxt | pair_102 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_103 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_104 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_105 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_106 | 1775.0 | 586970.0 |
| corrected_fdxt | pair_107 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_108 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_109 | 771.0 | 586970.0 |
| corrected_fdxt | pair_110 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_111 | 803.0 | 586970.0 |
| corrected_fdxt | pair_112 | 1580.0 | 586970.0 |
| corrected_fdxt | pair_113 | 2044.0 | 586970.0 |
| corrected_fdxt | pair_114 | 1839.0 | 586970.0 |
| corrected_fdxt | pair_115 | 2008.0 | 586970.0 |
| corrected_fdxt | pair_116 | 1839.0 | 586970.0 |
| corrected_fdxt | pair_117 | 1321.0 | 586970.0 |
| corrected_fdxt | pair_118 | 1321.0 | 586970.0 |
| corrected_fdxt | pair_119 | 1580.0 | 586970.0 |
| corrected_fdxt | pair_120 | 1580.0 | 586970.0 |
| corrected_fdxt | pair_121 | 1580.0 | 586970.0 |
| corrected_fdxt | pair_122 | 803.0 | 586970.0 |
| corrected_fdxt | pair_123 | 803.0 | 586970.0 |
| corrected_fdxt | pair_124 | 2035.0 | 586970.0 |
| corrected_fdxt | pair_125 | 1839.0 | 586970.0 |
| corrected_fdxt | pair_126 | 2035.0 | 586970.0 |
| corrected_fdxt | pair_127 | 1839.0 | 586970.0 |
| corrected_fdxt | pair_128 | 2044.0 | 586970.0 |
| corrected_fdxt | pair_129 | 1321.0 | 586970.0 |
| corrected_fdxt | pair_130 | 2044.0 | 586970.0 |
| corrected_fdxt | pair_131 | 2026.0 | 586970.0 |
| corrected_fdxt | pair_132 | 1321.0 | 586970.0 |
| corrected_fdxt | pair_133 | 1321.0 | 586970.0 |
| corrected_fdxt | pair_134 | 803.0 | 586970.0 |
| corrected_fdxt | pair_135 | 1580.0 | 586970.0 |
| corrected_fdxt | pair_136 | 767.0 | 586970.0 |
| corrected_fdxt | pair_137 | 779.0 | 586970.0 |
| corrected_fdxt | pair_138 | 787.0 | 586970.0 |
| corrected_fdxt | pair_139 | 763.0 | 586970.0 |
| corrected_fdxt | pair_140 | 779.0 | 586970.0 |
| corrected_fdxt | pair_141 | 767.0 | 586970.0 |
| corrected_fdxt | pair_142 | 787.0 | 586970.0 |
| corrected_fdxt | pair_143 | 779.0 | 586970.0 |
| corrected_fdxt | pair_144 | 779.0 | 586970.0 |
| corrected_fdxt | pair_145 | 783.0 | 586970.0 |
| corrected_fdxt | pair_146 | 779.0 | 586970.0 |
| corrected_fdxt | pair_147 | 775.0 | 586970.0 |
| corrected_fdxt | pair_148 | 775.0 | 586970.0 |
| corrected_fdxt | pair_149 | 787.0 | 586970.0 |
| corrected_fdxt | pair_150 | 775.0 | 586970.0 |
| corrected_fdxt | pair_151 | 779.0 | 586970.0 |
| corrected_fdxt | pair_152 | 779.0 | 586970.0 |
| corrected_fdxt | pair_153 | 771.0 | 586970.0 |
| corrected_fdxt | pair_154 | 779.0 | 586970.0 |
| corrected_fdxt | pair_155 | 771.0 | 586970.0 |
| corrected_fdxt | pair_156 | 771.0 | 586970.0 |
| corrected_fdxt | pair_157 | 775.0 | 586970.0 |
| corrected_fdxt | pair_158 | 799.0 | 586970.0 |
| corrected_fdxt | pair_159 | 759.0 | 586970.0 |
| corrected_fdxt | pair_160 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_161 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_162 | 1510.0 | 586970.0 |
| corrected_fdxt | pair_163 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_164 | 1267.0 | 586970.0 |
| corrected_fdxt | pair_165 | 1267.0 | 586970.0 |
| corrected_fdxt | pair_166 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_167 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_168 | 1545.0 | 586970.0 |
| corrected_fdxt | pair_169 | 767.0 | 586970.0 |
| corrected_fdxt | pair_170 | 767.0 | 586970.0 |
| corrected_fdxt | pair_171 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_172 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_173 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_174 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_175 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_176 | 1267.0 | 586970.0 |
| corrected_fdxt | pair_177 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_178 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_179 | 1267.0 | 586970.0 |
| corrected_fdxt | pair_180 | 1267.0 | 586970.0 |
| corrected_fdxt | pair_181 | 767.0 | 586970.0 |
| corrected_fdxt | pair_182 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_183 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_184 | 4000.0 | 586970.0 |
| corrected_fdxt | pair_185 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_186 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_187 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_188 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_189 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_190 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_191 | 779.0 | 586970.0 |
| corrected_fdxt | pair_192 | 779.0 | 586970.0 |
| corrected_fdxt | pair_193 | 2803.0 | 586970.0 |
| corrected_fdxt | pair_194 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_195 | 4051.0 | 586970.0 |
| corrected_fdxt | pair_196 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_197 | 4068.0 | 586970.0 |
| corrected_fdxt | pair_198 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_199 | 3056.0 | 586970.0 |
| corrected_fdxt | pair_200 | 4034.0 | 586970.0 |
| corrected_fdxt | pair_201 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_202 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_203 | 779.0 | 586970.0 |
| corrected_fdxt | pair_204 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_205 | 1759.0 | 586970.0 |
| corrected_fdxt | pair_206 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_207 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_208 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_209 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_210 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_211 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_212 | 787.0 | 586970.0 |
| corrected_fdxt | pair_213 | 787.0 | 586970.0 |
| corrected_fdxt | pair_214 | 1783.0 | 586970.0 |
| corrected_fdxt | pair_215 | 1807.0 | 586970.0 |
| corrected_fdxt | pair_216 | 1783.0 | 586970.0 |
| corrected_fdxt | pair_217 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_218 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_219 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_220 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_221 | 1775.0 | 586970.0 |
| corrected_fdxt | pair_222 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_223 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_224 | 787.0 | 586970.0 |
| corrected_fdxt | pair_225 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_226 | 1759.0 | 586970.0 |
| corrected_fdxt | pair_227 | 1261.0 | 586970.0 |
| corrected_fdxt | pair_228 | 1261.0 | 586970.0 |
| corrected_fdxt | pair_229 | 1510.0 | 586970.0 |
| corrected_fdxt | pair_230 | 1510.0 | 586970.0 |
| corrected_fdxt | pair_231 | 1510.0 | 586970.0 |
| corrected_fdxt | pair_232 | 763.0 | 586970.0 |
| corrected_fdxt | pair_233 | 763.0 | 586970.0 |
| corrected_fdxt | pair_234 | 2755.0 | 586970.0 |
| corrected_fdxt | pair_235 | 1759.0 | 586970.0 |
| corrected_fdxt | pair_236 | 5311.0 | 586970.0 |
| corrected_fdxt | pair_237 | 1759.0 | 586970.0 |
| corrected_fdxt | pair_238 | 4996.0 | 586970.0 |
| corrected_fdxt | pair_239 | 1261.0 | 586970.0 |
| corrected_fdxt | pair_240 | 3004.0 | 586970.0 |
| corrected_fdxt | pair_241 | 5289.0 | 586970.0 |
| corrected_fdxt | pair_242 | 1261.0 | 586970.0 |
| corrected_fdxt | pair_243 | 1261.0 | 586970.0 |
| corrected_fdxt | pair_244 | 763.0 | 586970.0 |
| corrected_fdxt | pair_245 | 1510.0 | 586970.0 |
| corrected_fdxt | pair_246 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_247 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_248 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_249 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_250 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_251 | 779.0 | 586970.0 |
| corrected_fdxt | pair_252 | 779.0 | 586970.0 |
| corrected_fdxt | pair_253 | 1783.0 | 586970.0 |
| corrected_fdxt | pair_254 | 1807.0 | 586970.0 |
| corrected_fdxt | pair_255 | 1783.0 | 586970.0 |
| corrected_fdxt | pair_256 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_257 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_258 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_259 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_260 | 1775.0 | 586970.0 |
| corrected_fdxt | pair_261 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_262 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_263 | 779.0 | 586970.0 |
| corrected_fdxt | pair_264 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_265 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_266 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_267 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_268 | 1291.0 | 586970.0 |
| corrected_fdxt | pair_269 | 767.0 | 586970.0 |
| corrected_fdxt | pair_270 | 767.0 | 586970.0 |
| corrected_fdxt | pair_271 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_272 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_273 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_274 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_275 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_276 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_277 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_278 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_279 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_280 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_281 | 767.0 | 586970.0 |
| corrected_fdxt | pair_282 | 1255.0 | 586970.0 |
| corrected_fdxt | pair_283 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_284 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_285 | 1291.0 | 586970.0 |
| corrected_fdxt | pair_286 | 787.0 | 586970.0 |
| corrected_fdxt | pair_287 | 787.0 | 586970.0 |
| corrected_fdxt | pair_288 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_289 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_290 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_291 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_292 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_293 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_294 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_295 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_296 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_297 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_298 | 787.0 | 586970.0 |
| corrected_fdxt | pair_299 | 1255.0 | 586970.0 |
| corrected_fdxt | pair_300 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_301 | 1545.0 | 586970.0 |
| corrected_fdxt | pair_302 | 779.0 | 586970.0 |
| corrected_fdxt | pair_303 | 779.0 | 586970.0 |
| corrected_fdxt | pair_304 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_305 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_306 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_307 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_308 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_309 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_310 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_311 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_312 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_313 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_314 | 779.0 | 586970.0 |
| corrected_fdxt | pair_315 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_316 | 1545.0 | 586970.0 |
| corrected_fdxt | pair_317 | 779.0 | 586970.0 |
| corrected_fdxt | pair_318 | 779.0 | 586970.0 |
| corrected_fdxt | pair_319 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_320 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_321 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_322 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_323 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_324 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_325 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_326 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_327 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_328 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_329 | 779.0 | 586970.0 |
| corrected_fdxt | pair_330 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_331 | 783.0 | 586970.0 |
| corrected_fdxt | pair_332 | 783.0 | 586970.0 |
| corrected_fdxt | pair_333 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_334 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_335 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_336 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_337 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_338 | 1291.0 | 586970.0 |
| corrected_fdxt | pair_339 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_340 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_341 | 1291.0 | 586970.0 |
| corrected_fdxt | pair_342 | 1291.0 | 586970.0 |
| corrected_fdxt | pair_343 | 783.0 | 586970.0 |
| corrected_fdxt | pair_344 | 1503.0 | 586970.0 |
| corrected_fdxt | pair_345 | 775.0 | 586970.0 |
| corrected_fdxt | pair_346 | 775.0 | 586970.0 |
| corrected_fdxt | pair_347 | 787.0 | 586970.0 |
| corrected_fdxt | pair_348 | 775.0 | 586970.0 |
| corrected_fdxt | pair_349 | 779.0 | 586970.0 |
| corrected_fdxt | pair_350 | 779.0 | 586970.0 |
| corrected_fdxt | pair_351 | 771.0 | 586970.0 |
| corrected_fdxt | pair_352 | 779.0 | 586970.0 |
| corrected_fdxt | pair_353 | 771.0 | 586970.0 |
| corrected_fdxt | pair_354 | 771.0 | 586970.0 |
| corrected_fdxt | pair_355 | 775.0 | 586970.0 |
| corrected_fdxt | pair_356 | 799.0 | 586970.0 |
| corrected_fdxt | pair_357 | 759.0 | 586970.0 |
| corrected_fdxt | pair_358 | 775.0 | 586970.0 |
| corrected_fdxt | pair_359 | 787.0 | 586970.0 |
| corrected_fdxt | pair_360 | 775.0 | 586970.0 |
| corrected_fdxt | pair_361 | 779.0 | 586970.0 |
| corrected_fdxt | pair_362 | 779.0 | 586970.0 |
| corrected_fdxt | pair_363 | 771.0 | 586970.0 |
| corrected_fdxt | pair_364 | 779.0 | 586970.0 |
| corrected_fdxt | pair_365 | 771.0 | 586970.0 |
| corrected_fdxt | pair_366 | 771.0 | 586970.0 |
| corrected_fdxt | pair_367 | 775.0 | 586970.0 |
| corrected_fdxt | pair_368 | 799.0 | 586970.0 |
| corrected_fdxt | pair_369 | 759.0 | 586970.0 |
| corrected_fdxt | pair_370 | 1783.0 | 586970.0 |
| corrected_fdxt | pair_371 | 2791.0 | 586970.0 |
| corrected_fdxt | pair_372 | 1783.0 | 586970.0 |
| corrected_fdxt | pair_373 | 2803.0 | 586970.0 |
| corrected_fdxt | pair_374 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_375 | 2803.0 | 586970.0 |
| corrected_fdxt | pair_376 | 2779.0 | 586970.0 |
| corrected_fdxt | pair_377 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_378 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_379 | 775.0 | 586970.0 |
| corrected_fdxt | pair_380 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_381 | 1783.0 | 586970.0 |
| corrected_fdxt | pair_382 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_383 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_384 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_385 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_386 | 1775.0 | 586970.0 |
| corrected_fdxt | pair_387 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_388 | 1297.0 | 586970.0 |
| corrected_fdxt | pair_389 | 787.0 | 586970.0 |
| corrected_fdxt | pair_390 | 1552.0 | 586970.0 |
| corrected_fdxt | pair_391 | 1783.0 | 586970.0 |
| corrected_fdxt | pair_392 | 5059.0 | 586970.0 |
| corrected_fdxt | pair_393 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_394 | 3043.0 | 586970.0 |
| corrected_fdxt | pair_395 | 5311.0 | 586970.0 |
| corrected_fdxt | pair_396 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_397 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_398 | 775.0 | 586970.0 |
| corrected_fdxt | pair_399 | 1531.0 | 586970.0 |
| corrected_fdxt | pair_400 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_401 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_402 | 1791.0 | 586970.0 |
| corrected_fdxt | pair_403 | 1775.0 | 586970.0 |
| corrected_fdxt | pair_404 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_405 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_406 | 779.0 | 586970.0 |
| corrected_fdxt | pair_407 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_408 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_409 | 3056.0 | 586970.0 |
| corrected_fdxt | pair_410 | 5038.0 | 586970.0 |
| corrected_fdxt | pair_411 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_412 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_413 | 779.0 | 586970.0 |
| corrected_fdxt | pair_414 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_415 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_416 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_417 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_418 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_419 | 771.0 | 586970.0 |
| corrected_fdxt | pair_420 | 1255.0 | 586970.0 |
| corrected_fdxt | pair_421 | 3030.0 | 586970.0 |
| corrected_fdxt | pair_422 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_423 | 1285.0 | 586970.0 |
| corrected_fdxt | pair_424 | 779.0 | 586970.0 |
| corrected_fdxt | pair_425 | 1538.0 | 586970.0 |
| corrected_fdxt | pair_426 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_427 | 1273.0 | 586970.0 |
| corrected_fdxt | pair_428 | 771.0 | 586970.0 |
| corrected_fdxt | pair_429 | 1524.0 | 586970.0 |
| corrected_fdxt | pair_430 | 1279.0 | 586970.0 |
| corrected_fdxt | pair_431 | 771.0 | 586970.0 |
| corrected_fdxt | pair_432 | 1255.0 | 586970.0 |
| corrected_fdxt | pair_433 | 775.0 | 586970.0 |
| corrected_fdxt | pair_434 | 1255.0 | 586970.0 |
| corrected_fdxt | pair_435 | 759.0 | 586970.0 |

## Table 4 — Volume / leakage-proxy measurements

| Variant | Workload | Mean observed volume | Mean cover volume | Exposure proxy |
|---|---|---:|---:|---:|
| c1_hybrid | pair_001 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_002 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_003 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_004 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_005 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_006 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_007 | 25.00 | 9.00 | 0.6400 |
| c1_hybrid | pair_008 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_009 | 31.00 | 10.00 | 0.6774 |
| c1_hybrid | pair_010 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_011 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_012 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_013 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_014 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_015 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_016 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_017 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_018 | 19.00 | 8.00 | 0.5789 |
| c1_hybrid | pair_019 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_020 | 31.00 | 10.00 | 0.6774 |
| c1_hybrid | pair_021 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_022 | 30.00 | 10.00 | 0.6667 |
| c1_hybrid | pair_023 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_024 | 20.00 | 8.00 | 0.6000 |
| c1_hybrid | pair_025 | 31.00 | 10.00 | 0.6774 |
| c1_hybrid | pair_026 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_027 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_028 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_029 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_030 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_031 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_032 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_033 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_034 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_035 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_036 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_037 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_038 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_039 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_040 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_041 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_042 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_043 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_044 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_045 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_046 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_047 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_048 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_049 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_050 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_051 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_052 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_053 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_054 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_055 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_056 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_057 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_058 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_059 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_060 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_061 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_062 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_063 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_064 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_065 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_066 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_067 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_068 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_069 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_070 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_071 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_072 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_073 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_074 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_075 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_076 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_077 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_078 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_079 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_080 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_081 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_082 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_083 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_084 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_085 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_086 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_087 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_088 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_089 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_090 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_091 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_092 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_093 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_094 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_095 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_096 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_097 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_098 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_099 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_100 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_101 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_102 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_103 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_104 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_105 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_106 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_107 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_108 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_109 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_110 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_111 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_112 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_113 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_114 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_115 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_116 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_117 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_118 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_119 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_120 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_121 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_122 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_123 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_124 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_125 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_126 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_127 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_128 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_129 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_130 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_131 | 15.00 | 7.00 | 0.5333 |
| c1_hybrid | pair_132 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_133 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_134 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_135 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_136 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_137 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_138 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_139 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_140 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_141 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_142 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_143 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_144 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_145 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_146 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_147 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_148 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_149 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_150 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_151 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_152 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_153 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_154 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_155 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_156 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_157 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_158 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_159 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_160 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_161 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_162 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_163 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_164 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_165 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_166 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_167 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_168 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_169 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_170 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_171 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_172 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_173 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_174 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_175 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_176 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_177 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_178 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_179 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_180 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_181 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_182 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_183 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_184 | 25.00 | 9.00 | 0.6400 |
| c1_hybrid | pair_185 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_186 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_187 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_188 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_189 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_190 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_191 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_192 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_193 | 19.00 | 8.00 | 0.5789 |
| c1_hybrid | pair_194 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_195 | 25.00 | 9.00 | 0.6400 |
| c1_hybrid | pair_196 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_197 | 25.00 | 9.00 | 0.6400 |
| c1_hybrid | pair_198 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_199 | 20.00 | 8.00 | 0.6000 |
| c1_hybrid | pair_200 | 25.00 | 9.00 | 0.6400 |
| c1_hybrid | pair_201 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_202 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_203 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_204 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_205 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_206 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_207 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_208 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_209 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_210 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_211 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_212 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_213 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_214 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_215 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_216 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_217 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_218 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_219 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_220 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_221 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_222 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_223 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_224 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_225 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_226 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_227 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_228 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_229 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_230 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_231 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_232 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_233 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_234 | 19.00 | 8.00 | 0.5789 |
| c1_hybrid | pair_235 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_236 | 31.00 | 10.00 | 0.6774 |
| c1_hybrid | pair_237 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_238 | 30.00 | 10.00 | 0.6667 |
| c1_hybrid | pair_239 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_240 | 20.00 | 8.00 | 0.6000 |
| c1_hybrid | pair_241 | 31.00 | 10.00 | 0.6774 |
| c1_hybrid | pair_242 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_243 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_244 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_245 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_246 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_247 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_248 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_249 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_250 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_251 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_252 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_253 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_254 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_255 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_256 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_257 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_258 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_259 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_260 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_261 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_262 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_263 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_264 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_265 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_266 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_267 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_268 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_269 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_270 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_271 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_272 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_273 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_274 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_275 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_276 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_277 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_278 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_279 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_280 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_281 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_282 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_283 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_284 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_285 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_286 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_287 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_288 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_289 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_290 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_291 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_292 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_293 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_294 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_295 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_296 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_297 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_298 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_299 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_300 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_301 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_302 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_303 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_304 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_305 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_306 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_307 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_308 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_309 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_310 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_311 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_312 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_313 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_314 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_315 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_316 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_317 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_318 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_319 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_320 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_321 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_322 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_323 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_324 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_325 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_326 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_327 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_328 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_329 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_330 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_331 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_332 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_333 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_334 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_335 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_336 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_337 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_338 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_339 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_340 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_341 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_342 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_343 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_344 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_345 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_346 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_347 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_348 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_349 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_350 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_351 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_352 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_353 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_354 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_355 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_356 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_357 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_358 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_359 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_360 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_361 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_362 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_363 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_364 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_365 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_366 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_367 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_368 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_369 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_370 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_371 | 19.00 | 8.00 | 0.5789 |
| c1_hybrid | pair_372 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_373 | 19.00 | 8.00 | 0.5789 |
| c1_hybrid | pair_374 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_375 | 19.00 | 8.00 | 0.5789 |
| c1_hybrid | pair_376 | 19.00 | 8.00 | 0.5789 |
| c1_hybrid | pair_377 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_378 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_379 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_380 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_381 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_382 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_383 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_384 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_385 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_386 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_387 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_388 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_389 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_390 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_391 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_392 | 30.00 | 10.00 | 0.6667 |
| c1_hybrid | pair_393 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_394 | 20.00 | 8.00 | 0.6000 |
| c1_hybrid | pair_395 | 31.00 | 10.00 | 0.6774 |
| c1_hybrid | pair_396 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_397 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_398 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_399 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_400 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_401 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_402 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_403 | 13.00 | 6.00 | 0.5385 |
| c1_hybrid | pair_404 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_405 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_406 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_407 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_408 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_409 | 20.00 | 8.00 | 0.6000 |
| c1_hybrid | pair_410 | 30.00 | 10.00 | 0.6667 |
| c1_hybrid | pair_411 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_412 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_413 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_414 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_415 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_416 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_417 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_418 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_419 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_420 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_421 | 20.00 | 8.00 | 0.6000 |
| c1_hybrid | pair_422 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_423 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_424 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_425 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_426 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_427 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_428 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_429 | 12.00 | 6.00 | 0.5000 |
| c1_hybrid | pair_430 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_431 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_432 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_433 | 7.00 | 4.00 | 0.4286 |
| c1_hybrid | pair_434 | 11.00 | 6.00 | 0.4545 |
| c1_hybrid | pair_435 | 7.00 | 4.00 | 0.4286 |
| c1_log | pair_001 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_002 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_003 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_004 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_005 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_006 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_007 | 21.00 | 5.00 | 0.7619 |
| c1_log | pair_008 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_009 | 26.00 | 5.00 | 0.8077 |
| c1_log | pair_010 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_011 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_012 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_013 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_014 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_015 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_016 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_017 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_018 | 15.00 | 4.00 | 0.7333 |
| c1_log | pair_019 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_020 | 26.00 | 5.00 | 0.8077 |
| c1_log | pair_021 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_022 | 25.00 | 5.00 | 0.8000 |
| c1_log | pair_023 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_024 | 16.00 | 4.00 | 0.7500 |
| c1_log | pair_025 | 26.00 | 5.00 | 0.8077 |
| c1_log | pair_026 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_027 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_028 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_029 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_030 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_031 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_032 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_033 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_034 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_035 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_036 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_037 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_038 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_039 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_040 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_041 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_042 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_043 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_044 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_045 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_046 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_047 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_048 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_049 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_050 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_051 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_052 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_053 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_054 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_055 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_056 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_057 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_058 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_059 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_060 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_061 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_062 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_063 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_064 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_065 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_066 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_067 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_068 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_069 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_070 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_071 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_072 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_073 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_074 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_075 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_076 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_077 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_078 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_079 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_080 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_081 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_082 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_083 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_084 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_085 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_086 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_087 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_088 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_089 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_090 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_091 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_092 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_093 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_094 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_095 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_096 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_097 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_098 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_099 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_100 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_101 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_102 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_103 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_104 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_105 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_106 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_107 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_108 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_109 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_110 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_111 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_112 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_113 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_114 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_115 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_116 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_117 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_118 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_119 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_120 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_121 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_122 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_123 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_124 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_125 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_126 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_127 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_128 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_129 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_130 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_131 | 12.00 | 4.00 | 0.6667 |
| c1_log | pair_132 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_133 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_134 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_135 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_136 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_137 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_138 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_139 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_140 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_141 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_142 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_143 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_144 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_145 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_146 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_147 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_148 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_149 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_150 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_151 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_152 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_153 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_154 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_155 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_156 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_157 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_158 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_159 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_160 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_161 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_162 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_163 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_164 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_165 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_166 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_167 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_168 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_169 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_170 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_171 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_172 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_173 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_174 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_175 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_176 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_177 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_178 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_179 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_180 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_181 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_182 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_183 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_184 | 21.00 | 5.00 | 0.7619 |
| c1_log | pair_185 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_186 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_187 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_188 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_189 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_190 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_191 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_192 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_193 | 15.00 | 4.00 | 0.7333 |
| c1_log | pair_194 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_195 | 21.00 | 5.00 | 0.7619 |
| c1_log | pair_196 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_197 | 21.00 | 5.00 | 0.7619 |
| c1_log | pair_198 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_199 | 16.00 | 4.00 | 0.7500 |
| c1_log | pair_200 | 21.00 | 5.00 | 0.7619 |
| c1_log | pair_201 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_202 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_203 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_204 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_205 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_206 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_207 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_208 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_209 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_210 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_211 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_212 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_213 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_214 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_215 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_216 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_217 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_218 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_219 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_220 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_221 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_222 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_223 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_224 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_225 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_226 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_227 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_228 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_229 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_230 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_231 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_232 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_233 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_234 | 15.00 | 4.00 | 0.7333 |
| c1_log | pair_235 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_236 | 26.00 | 5.00 | 0.8077 |
| c1_log | pair_237 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_238 | 25.00 | 5.00 | 0.8000 |
| c1_log | pair_239 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_240 | 16.00 | 4.00 | 0.7500 |
| c1_log | pair_241 | 26.00 | 5.00 | 0.8077 |
| c1_log | pair_242 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_243 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_244 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_245 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_246 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_247 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_248 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_249 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_250 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_251 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_252 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_253 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_254 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_255 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_256 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_257 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_258 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_259 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_260 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_261 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_262 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_263 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_264 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_265 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_266 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_267 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_268 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_269 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_270 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_271 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_272 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_273 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_274 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_275 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_276 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_277 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_278 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_279 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_280 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_281 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_282 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_283 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_284 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_285 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_286 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_287 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_288 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_289 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_290 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_291 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_292 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_293 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_294 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_295 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_296 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_297 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_298 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_299 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_300 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_301 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_302 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_303 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_304 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_305 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_306 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_307 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_308 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_309 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_310 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_311 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_312 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_313 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_314 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_315 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_316 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_317 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_318 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_319 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_320 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_321 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_322 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_323 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_324 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_325 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_326 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_327 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_328 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_329 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_330 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_331 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_332 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_333 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_334 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_335 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_336 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_337 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_338 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_339 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_340 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_341 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_342 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_343 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_344 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_345 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_346 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_347 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_348 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_349 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_350 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_351 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_352 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_353 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_354 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_355 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_356 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_357 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_358 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_359 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_360 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_361 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_362 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_363 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_364 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_365 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_366 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_367 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_368 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_369 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_370 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_371 | 15.00 | 4.00 | 0.7333 |
| c1_log | pair_372 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_373 | 15.00 | 4.00 | 0.7333 |
| c1_log | pair_374 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_375 | 15.00 | 4.00 | 0.7333 |
| c1_log | pair_376 | 15.00 | 4.00 | 0.7333 |
| c1_log | pair_377 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_378 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_379 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_380 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_381 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_382 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_383 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_384 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_385 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_386 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_387 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_388 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_389 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_390 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_391 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_392 | 25.00 | 5.00 | 0.8000 |
| c1_log | pair_393 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_394 | 16.00 | 4.00 | 0.7500 |
| c1_log | pair_395 | 26.00 | 5.00 | 0.8077 |
| c1_log | pair_396 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_397 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_398 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_399 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_400 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_401 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_402 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_403 | 10.00 | 3.00 | 0.7000 |
| c1_log | pair_404 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_405 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_406 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_407 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_408 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_409 | 16.00 | 4.00 | 0.7500 |
| c1_log | pair_410 | 25.00 | 5.00 | 0.8000 |
| c1_log | pair_411 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_412 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_413 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_414 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_415 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_416 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_417 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_418 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_419 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_420 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_421 | 16.00 | 4.00 | 0.7500 |
| c1_log | pair_422 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_423 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_424 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_425 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_426 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_427 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_428 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_429 | 9.00 | 3.00 | 0.6667 |
| c1_log | pair_430 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_431 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_432 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_433 | 5.00 | 2.00 | 0.6000 |
| c1_log | pair_434 | 8.00 | 3.00 | 0.6250 |
| c1_log | pair_435 | 5.00 | 2.00 | 0.6000 |
| c1_power2 | pair_001 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_002 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_003 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_004 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_005 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_006 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_007 | 16.00 | 0.00 | 1.0000 |
| c1_power2 | pair_008 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_009 | 32.00 | 11.00 | 0.6562 |
| c1_power2 | pair_010 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_011 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_012 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_013 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_014 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_015 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_016 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_017 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_018 | 16.00 | 5.00 | 0.6875 |
| c1_power2 | pair_019 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_020 | 32.00 | 11.00 | 0.6562 |
| c1_power2 | pair_021 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_022 | 32.00 | 12.00 | 0.6250 |
| c1_power2 | pair_023 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_024 | 16.00 | 4.00 | 0.7500 |
| c1_power2 | pair_025 | 32.00 | 11.00 | 0.6562 |
| c1_power2 | pair_026 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_027 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_028 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_029 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_030 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_031 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_032 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_033 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_034 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_035 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_036 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_037 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_038 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_039 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_040 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_041 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_042 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_043 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_044 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_045 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_046 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_047 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_048 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_049 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_050 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_051 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_052 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_053 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_054 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_055 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_056 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_057 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_058 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_059 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_060 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_061 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_062 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_063 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_064 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_065 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_066 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_067 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_068 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_069 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_070 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_071 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_072 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_073 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_074 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_075 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_076 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_077 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_078 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_079 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_080 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_081 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_082 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_083 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_084 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_085 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_086 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_087 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_088 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_089 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_090 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_091 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_092 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_093 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_094 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_095 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_096 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_097 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_098 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_099 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_100 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_101 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_102 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_103 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_104 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_105 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_106 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_107 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_108 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_109 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_110 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_111 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_112 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_113 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_114 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_115 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_116 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_117 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_118 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_119 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_120 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_121 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_122 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_123 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_124 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_125 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_126 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_127 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_128 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_129 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_130 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_131 | 8.00 | 0.00 | 1.0000 |
| c1_power2 | pair_132 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_133 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_134 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_135 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_136 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_137 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_138 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_139 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_140 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_141 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_142 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_143 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_144 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_145 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_146 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_147 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_148 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_149 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_150 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_151 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_152 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_153 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_154 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_155 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_156 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_157 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_158 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_159 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_160 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_161 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_162 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_163 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_164 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_165 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_166 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_167 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_168 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_169 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_170 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_171 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_172 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_173 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_174 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_175 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_176 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_177 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_178 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_179 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_180 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_181 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_182 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_183 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_184 | 16.00 | 0.00 | 1.0000 |
| c1_power2 | pair_185 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_186 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_187 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_188 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_189 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_190 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_191 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_192 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_193 | 16.00 | 5.00 | 0.6875 |
| c1_power2 | pair_194 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_195 | 16.00 | 0.00 | 1.0000 |
| c1_power2 | pair_196 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_197 | 16.00 | 0.00 | 1.0000 |
| c1_power2 | pair_198 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_199 | 16.00 | 4.00 | 0.7500 |
| c1_power2 | pair_200 | 16.00 | 0.00 | 1.0000 |
| c1_power2 | pair_201 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_202 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_203 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_204 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_205 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_206 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_207 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_208 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_209 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_210 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_211 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_212 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_213 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_214 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_215 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_216 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_217 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_218 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_219 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_220 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_221 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_222 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_223 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_224 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_225 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_226 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_227 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_228 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_229 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_230 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_231 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_232 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_233 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_234 | 16.00 | 5.00 | 0.6875 |
| c1_power2 | pair_235 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_236 | 32.00 | 11.00 | 0.6562 |
| c1_power2 | pair_237 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_238 | 32.00 | 12.00 | 0.6250 |
| c1_power2 | pair_239 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_240 | 16.00 | 4.00 | 0.7500 |
| c1_power2 | pair_241 | 32.00 | 11.00 | 0.6562 |
| c1_power2 | pair_242 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_243 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_244 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_245 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_246 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_247 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_248 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_249 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_250 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_251 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_252 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_253 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_254 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_255 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_256 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_257 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_258 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_259 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_260 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_261 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_262 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_263 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_264 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_265 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_266 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_267 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_268 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_269 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_270 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_271 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_272 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_273 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_274 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_275 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_276 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_277 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_278 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_279 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_280 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_281 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_282 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_283 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_284 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_285 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_286 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_287 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_288 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_289 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_290 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_291 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_292 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_293 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_294 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_295 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_296 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_297 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_298 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_299 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_300 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_301 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_302 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_303 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_304 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_305 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_306 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_307 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_308 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_309 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_310 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_311 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_312 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_313 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_314 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_315 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_316 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_317 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_318 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_319 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_320 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_321 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_322 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_323 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_324 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_325 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_326 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_327 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_328 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_329 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_330 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_331 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_332 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_333 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_334 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_335 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_336 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_337 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_338 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_339 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_340 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_341 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_342 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_343 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_344 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_345 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_346 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_347 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_348 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_349 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_350 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_351 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_352 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_353 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_354 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_355 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_356 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_357 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_358 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_359 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_360 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_361 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_362 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_363 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_364 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_365 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_366 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_367 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_368 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_369 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_370 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_371 | 16.00 | 5.00 | 0.6875 |
| c1_power2 | pair_372 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_373 | 16.00 | 5.00 | 0.6875 |
| c1_power2 | pair_374 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_375 | 16.00 | 5.00 | 0.6875 |
| c1_power2 | pair_376 | 16.00 | 5.00 | 0.6875 |
| c1_power2 | pair_377 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_378 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_379 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_380 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_381 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_382 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_383 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_384 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_385 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_386 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_387 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_388 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_389 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_390 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_391 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_392 | 32.00 | 12.00 | 0.6250 |
| c1_power2 | pair_393 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_394 | 16.00 | 4.00 | 0.7500 |
| c1_power2 | pair_395 | 32.00 | 11.00 | 0.6562 |
| c1_power2 | pair_396 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_397 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_398 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_399 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_400 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_401 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_402 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_403 | 8.00 | 1.00 | 0.8750 |
| c1_power2 | pair_404 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_405 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_406 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_407 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_408 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_409 | 16.00 | 4.00 | 0.7500 |
| c1_power2 | pair_410 | 32.00 | 12.00 | 0.6250 |
| c1_power2 | pair_411 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_412 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_413 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_414 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_415 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_416 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_417 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_418 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_419 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_420 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_421 | 16.00 | 4.00 | 0.7500 |
| c1_power2 | pair_422 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_423 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_424 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_425 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_426 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_427 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_428 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_429 | 8.00 | 2.00 | 0.7500 |
| c1_power2 | pair_430 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_431 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_432 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_433 | 4.00 | 1.00 | 0.7500 |
| c1_power2 | pair_434 | 8.00 | 3.00 | 0.6250 |
| c1_power2 | pair_435 | 4.00 | 1.00 | 0.7500 |
| c1_sqrt | pair_001 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_002 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_003 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_004 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_005 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_006 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_007 | 20.00 | 4.00 | 0.8000 |
| c1_sqrt | pair_008 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_009 | 26.00 | 5.00 | 0.8077 |
| c1_sqrt | pair_010 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_011 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_012 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_013 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_014 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_015 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_016 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_017 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_018 | 15.00 | 4.00 | 0.7333 |
| c1_sqrt | pair_019 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_020 | 26.00 | 5.00 | 0.8077 |
| c1_sqrt | pair_021 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_022 | 25.00 | 5.00 | 0.8000 |
| c1_sqrt | pair_023 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_024 | 16.00 | 4.00 | 0.7500 |
| c1_sqrt | pair_025 | 26.00 | 5.00 | 0.8077 |
| c1_sqrt | pair_026 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_027 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_028 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_029 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_030 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_031 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_032 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_033 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_034 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_035 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_036 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_037 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_038 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_039 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_040 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_041 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_042 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_043 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_044 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_045 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_046 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_047 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_048 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_049 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_050 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_051 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_052 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_053 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_054 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_055 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_056 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_057 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_058 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_059 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_060 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_061 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_062 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_063 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_064 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_065 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_066 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_067 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_068 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_069 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_070 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_071 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_072 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_073 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_074 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_075 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_076 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_077 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_078 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_079 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_080 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_081 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_082 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_083 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_084 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_085 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_086 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_087 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_088 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_089 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_090 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_091 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_092 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_093 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_094 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_095 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_096 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_097 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_098 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_099 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_100 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_101 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_102 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_103 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_104 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_105 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_106 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_107 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_108 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_109 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_110 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_111 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_112 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_113 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_114 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_115 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_116 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_117 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_118 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_119 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_120 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_121 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_122 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_123 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_124 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_125 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_126 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_127 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_128 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_129 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_130 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_131 | 11.00 | 3.00 | 0.7273 |
| c1_sqrt | pair_132 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_133 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_134 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_135 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_136 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_137 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_138 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_139 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_140 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_141 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_142 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_143 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_144 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_145 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_146 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_147 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_148 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_149 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_150 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_151 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_152 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_153 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_154 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_155 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_156 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_157 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_158 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_159 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_160 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_161 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_162 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_163 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_164 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_165 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_166 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_167 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_168 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_169 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_170 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_171 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_172 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_173 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_174 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_175 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_176 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_177 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_178 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_179 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_180 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_181 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_182 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_183 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_184 | 20.00 | 4.00 | 0.8000 |
| c1_sqrt | pair_185 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_186 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_187 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_188 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_189 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_190 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_191 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_192 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_193 | 15.00 | 4.00 | 0.7333 |
| c1_sqrt | pair_194 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_195 | 20.00 | 4.00 | 0.8000 |
| c1_sqrt | pair_196 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_197 | 20.00 | 4.00 | 0.8000 |
| c1_sqrt | pair_198 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_199 | 16.00 | 4.00 | 0.7500 |
| c1_sqrt | pair_200 | 20.00 | 4.00 | 0.8000 |
| c1_sqrt | pair_201 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_202 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_203 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_204 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_205 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_206 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_207 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_208 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_209 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_210 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_211 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_212 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_213 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_214 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_215 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_216 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_217 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_218 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_219 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_220 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_221 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_222 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_223 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_224 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_225 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_226 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_227 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_228 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_229 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_230 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_231 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_232 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_233 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_234 | 15.00 | 4.00 | 0.7333 |
| c1_sqrt | pair_235 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_236 | 26.00 | 5.00 | 0.8077 |
| c1_sqrt | pair_237 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_238 | 25.00 | 5.00 | 0.8000 |
| c1_sqrt | pair_239 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_240 | 16.00 | 4.00 | 0.7500 |
| c1_sqrt | pair_241 | 26.00 | 5.00 | 0.8077 |
| c1_sqrt | pair_242 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_243 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_244 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_245 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_246 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_247 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_248 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_249 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_250 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_251 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_252 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_253 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_254 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_255 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_256 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_257 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_258 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_259 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_260 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_261 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_262 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_263 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_264 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_265 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_266 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_267 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_268 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_269 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_270 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_271 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_272 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_273 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_274 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_275 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_276 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_277 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_278 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_279 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_280 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_281 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_282 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_283 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_284 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_285 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_286 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_287 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_288 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_289 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_290 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_291 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_292 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_293 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_294 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_295 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_296 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_297 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_298 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_299 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_300 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_301 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_302 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_303 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_304 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_305 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_306 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_307 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_308 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_309 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_310 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_311 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_312 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_313 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_314 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_315 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_316 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_317 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_318 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_319 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_320 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_321 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_322 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_323 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_324 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_325 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_326 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_327 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_328 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_329 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_330 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_331 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_332 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_333 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_334 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_335 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_336 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_337 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_338 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_339 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_340 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_341 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_342 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_343 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_344 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_345 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_346 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_347 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_348 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_349 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_350 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_351 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_352 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_353 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_354 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_355 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_356 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_357 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_358 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_359 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_360 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_361 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_362 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_363 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_364 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_365 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_366 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_367 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_368 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_369 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_370 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_371 | 15.00 | 4.00 | 0.7333 |
| c1_sqrt | pair_372 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_373 | 15.00 | 4.00 | 0.7333 |
| c1_sqrt | pair_374 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_375 | 15.00 | 4.00 | 0.7333 |
| c1_sqrt | pair_376 | 15.00 | 4.00 | 0.7333 |
| c1_sqrt | pair_377 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_378 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_379 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_380 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_381 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_382 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_383 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_384 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_385 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_386 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_387 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_388 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_389 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_390 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_391 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_392 | 25.00 | 5.00 | 0.8000 |
| c1_sqrt | pair_393 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_394 | 16.00 | 4.00 | 0.7500 |
| c1_sqrt | pair_395 | 26.00 | 5.00 | 0.8077 |
| c1_sqrt | pair_396 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_397 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_398 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_399 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_400 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_401 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_402 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_403 | 10.00 | 3.00 | 0.7000 |
| c1_sqrt | pair_404 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_405 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_406 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_407 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_408 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_409 | 16.00 | 4.00 | 0.7500 |
| c1_sqrt | pair_410 | 25.00 | 5.00 | 0.8000 |
| c1_sqrt | pair_411 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_412 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_413 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_414 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_415 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_416 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_417 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_418 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_419 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_420 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_421 | 16.00 | 4.00 | 0.7500 |
| c1_sqrt | pair_422 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_423 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_424 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_425 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_426 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_427 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_428 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_429 | 9.00 | 3.00 | 0.6667 |
| c1_sqrt | pair_430 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_431 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_432 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_433 | 5.00 | 2.00 | 0.6000 |
| c1_sqrt | pair_434 | 8.00 | 3.00 | 0.6250 |
| c1_sqrt | pair_435 | 5.00 | 2.00 | 0.6000 |
| corrected_fdxt | pair_001 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_002 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_003 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_004 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_005 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_006 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_007 | 16.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_008 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_009 | 21.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_010 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_011 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_012 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_013 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_014 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_015 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_016 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_017 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_018 | 11.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_019 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_020 | 21.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_021 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_022 | 20.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_023 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_024 | 12.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_025 | 21.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_026 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_027 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_028 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_029 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_030 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_031 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_032 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_033 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_034 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_035 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_036 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_037 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_038 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_039 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_040 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_041 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_042 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_043 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_044 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_045 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_046 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_047 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_048 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_049 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_050 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_051 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_052 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_053 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_054 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_055 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_056 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_057 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_058 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_059 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_060 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_061 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_062 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_063 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_064 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_065 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_066 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_067 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_068 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_069 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_070 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_071 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_072 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_073 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_074 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_075 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_076 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_077 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_078 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_079 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_080 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_081 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_082 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_083 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_084 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_085 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_086 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_087 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_088 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_089 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_090 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_091 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_092 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_093 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_094 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_095 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_096 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_097 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_098 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_099 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_100 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_101 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_102 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_103 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_104 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_105 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_106 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_107 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_108 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_109 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_110 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_111 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_112 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_113 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_114 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_115 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_116 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_117 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_118 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_119 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_120 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_121 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_122 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_123 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_124 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_125 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_126 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_127 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_128 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_129 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_130 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_131 | 8.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_132 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_133 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_134 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_135 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_136 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_137 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_138 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_139 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_140 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_141 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_142 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_143 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_144 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_145 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_146 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_147 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_148 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_149 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_150 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_151 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_152 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_153 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_154 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_155 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_156 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_157 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_158 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_159 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_160 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_161 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_162 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_163 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_164 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_165 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_166 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_167 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_168 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_169 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_170 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_171 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_172 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_173 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_174 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_175 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_176 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_177 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_178 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_179 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_180 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_181 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_182 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_183 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_184 | 16.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_185 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_186 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_187 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_188 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_189 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_190 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_191 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_192 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_193 | 11.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_194 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_195 | 16.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_196 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_197 | 16.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_198 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_199 | 12.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_200 | 16.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_201 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_202 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_203 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_204 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_205 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_206 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_207 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_208 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_209 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_210 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_211 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_212 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_213 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_214 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_215 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_216 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_217 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_218 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_219 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_220 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_221 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_222 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_223 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_224 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_225 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_226 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_227 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_228 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_229 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_230 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_231 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_232 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_233 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_234 | 11.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_235 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_236 | 21.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_237 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_238 | 20.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_239 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_240 | 12.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_241 | 21.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_242 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_243 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_244 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_245 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_246 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_247 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_248 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_249 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_250 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_251 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_252 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_253 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_254 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_255 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_256 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_257 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_258 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_259 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_260 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_261 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_262 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_263 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_264 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_265 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_266 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_267 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_268 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_269 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_270 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_271 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_272 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_273 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_274 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_275 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_276 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_277 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_278 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_279 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_280 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_281 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_282 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_283 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_284 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_285 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_286 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_287 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_288 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_289 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_290 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_291 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_292 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_293 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_294 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_295 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_296 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_297 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_298 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_299 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_300 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_301 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_302 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_303 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_304 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_305 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_306 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_307 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_308 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_309 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_310 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_311 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_312 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_313 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_314 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_315 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_316 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_317 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_318 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_319 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_320 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_321 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_322 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_323 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_324 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_325 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_326 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_327 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_328 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_329 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_330 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_331 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_332 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_333 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_334 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_335 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_336 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_337 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_338 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_339 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_340 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_341 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_342 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_343 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_344 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_345 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_346 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_347 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_348 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_349 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_350 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_351 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_352 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_353 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_354 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_355 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_356 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_357 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_358 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_359 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_360 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_361 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_362 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_363 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_364 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_365 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_366 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_367 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_368 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_369 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_370 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_371 | 11.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_372 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_373 | 11.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_374 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_375 | 11.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_376 | 11.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_377 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_378 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_379 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_380 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_381 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_382 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_383 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_384 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_385 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_386 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_387 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_388 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_389 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_390 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_391 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_392 | 20.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_393 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_394 | 12.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_395 | 21.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_396 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_397 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_398 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_399 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_400 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_401 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_402 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_403 | 7.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_404 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_405 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_406 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_407 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_408 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_409 | 12.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_410 | 20.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_411 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_412 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_413 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_414 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_415 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_416 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_417 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_418 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_419 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_420 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_421 | 12.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_422 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_423 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_424 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_425 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_426 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_427 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_428 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_429 | 6.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_430 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_431 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_432 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_433 | 3.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_434 | 5.00 | 0.00 | 1.0000 |
| corrected_fdxt | pair_435 | 3.00 | 0.00 | 1.0000 |

## Table 5 — C1 overhead

| Variant | Workload | Mean C1 overhead |
|---|---|---:|
| c1_hybrid | pair_001 | 0.8750 |
| c1_hybrid | pair_002 | 1.0000 |
| c1_hybrid | pair_003 | 0.8571 |
| c1_hybrid | pair_004 | 0.8750 |
| c1_hybrid | pair_005 | 1.3333 |
| c1_hybrid | pair_006 | 1.0000 |
| c1_hybrid | pair_007 | 0.5625 |
| c1_hybrid | pair_008 | 0.8571 |
| c1_hybrid | pair_009 | 0.4762 |
| c1_hybrid | pair_010 | 0.8571 |
| c1_hybrid | pair_011 | 1.2000 |
| c1_hybrid | pair_012 | 1.2000 |
| c1_hybrid | pair_013 | 1.0000 |
| c1_hybrid | pair_014 | 1.0000 |
| c1_hybrid | pair_015 | 1.0000 |
| c1_hybrid | pair_016 | 1.3333 |
| c1_hybrid | pair_017 | 1.3333 |
| c1_hybrid | pair_018 | 0.7273 |
| c1_hybrid | pair_019 | 0.8571 |
| c1_hybrid | pair_020 | 0.4762 |
| c1_hybrid | pair_021 | 0.8571 |
| c1_hybrid | pair_022 | 0.5000 |
| c1_hybrid | pair_023 | 1.2000 |
| c1_hybrid | pair_024 | 0.6667 |
| c1_hybrid | pair_025 | 0.4762 |
| c1_hybrid | pair_026 | 1.2000 |
| c1_hybrid | pair_027 | 1.2000 |
| c1_hybrid | pair_028 | 1.3333 |
| c1_hybrid | pair_029 | 1.0000 |
| c1_hybrid | pair_030 | 1.0000 |
| c1_hybrid | pair_031 | 0.8571 |
| c1_hybrid | pair_032 | 0.8750 |
| c1_hybrid | pair_033 | 1.3333 |
| c1_hybrid | pair_034 | 1.0000 |
| c1_hybrid | pair_035 | 0.8750 |
| c1_hybrid | pair_036 | 0.8571 |
| c1_hybrid | pair_037 | 0.8750 |
| c1_hybrid | pair_038 | 0.8571 |
| c1_hybrid | pair_039 | 1.2000 |
| c1_hybrid | pair_040 | 1.2000 |
| c1_hybrid | pair_041 | 1.0000 |
| c1_hybrid | pair_042 | 1.0000 |
| c1_hybrid | pair_043 | 1.0000 |
| c1_hybrid | pair_044 | 1.3333 |
| c1_hybrid | pair_045 | 1.3333 |
| c1_hybrid | pair_046 | 0.8750 |
| c1_hybrid | pair_047 | 0.8571 |
| c1_hybrid | pair_048 | 0.8750 |
| c1_hybrid | pair_049 | 0.8571 |
| c1_hybrid | pair_050 | 0.8750 |
| c1_hybrid | pair_051 | 1.2000 |
| c1_hybrid | pair_052 | 0.8750 |
| c1_hybrid | pair_053 | 0.8750 |
| c1_hybrid | pair_054 | 1.2000 |
| c1_hybrid | pair_055 | 1.2000 |
| c1_hybrid | pair_056 | 1.3333 |
| c1_hybrid | pair_057 | 1.0000 |
| c1_hybrid | pair_058 | 1.0000 |
| c1_hybrid | pair_059 | 1.0000 |
| c1_hybrid | pair_060 | 1.3333 |
| c1_hybrid | pair_061 | 1.0000 |
| c1_hybrid | pair_062 | 1.0000 |
| c1_hybrid | pair_063 | 1.0000 |
| c1_hybrid | pair_064 | 1.0000 |
| c1_hybrid | pair_065 | 1.0000 |
| c1_hybrid | pair_066 | 1.2000 |
| c1_hybrid | pair_067 | 1.2000 |
| c1_hybrid | pair_068 | 1.0000 |
| c1_hybrid | pair_069 | 1.0000 |
| c1_hybrid | pair_070 | 1.0000 |
| c1_hybrid | pair_071 | 1.3333 |
| c1_hybrid | pair_072 | 1.3333 |
| c1_hybrid | pair_073 | 1.0000 |
| c1_hybrid | pair_074 | 1.0000 |
| c1_hybrid | pair_075 | 1.0000 |
| c1_hybrid | pair_076 | 1.0000 |
| c1_hybrid | pair_077 | 1.0000 |
| c1_hybrid | pair_078 | 1.2000 |
| c1_hybrid | pair_079 | 1.0000 |
| c1_hybrid | pair_080 | 1.0000 |
| c1_hybrid | pair_081 | 1.2000 |
| c1_hybrid | pair_082 | 1.2000 |
| c1_hybrid | pair_083 | 1.3333 |
| c1_hybrid | pair_084 | 1.0000 |
| c1_hybrid | pair_085 | 0.8571 |
| c1_hybrid | pair_086 | 1.3333 |
| c1_hybrid | pair_087 | 1.0000 |
| c1_hybrid | pair_088 | 0.8571 |
| c1_hybrid | pair_089 | 0.8571 |
| c1_hybrid | pair_090 | 0.8571 |
| c1_hybrid | pair_091 | 0.8571 |
| c1_hybrid | pair_092 | 1.2000 |
| c1_hybrid | pair_093 | 1.2000 |
| c1_hybrid | pair_094 | 1.0000 |
| c1_hybrid | pair_095 | 1.0000 |
| c1_hybrid | pair_096 | 1.0000 |
| c1_hybrid | pair_097 | 1.3333 |
| c1_hybrid | pair_098 | 1.3333 |
| c1_hybrid | pair_099 | 0.8571 |
| c1_hybrid | pair_100 | 0.8571 |
| c1_hybrid | pair_101 | 0.8571 |
| c1_hybrid | pair_102 | 0.8571 |
| c1_hybrid | pair_103 | 0.8571 |
| c1_hybrid | pair_104 | 1.2000 |
| c1_hybrid | pair_105 | 0.8571 |
| c1_hybrid | pair_106 | 0.8571 |
| c1_hybrid | pair_107 | 1.2000 |
| c1_hybrid | pair_108 | 1.2000 |
| c1_hybrid | pair_109 | 1.3333 |
| c1_hybrid | pair_110 | 1.0000 |
| c1_hybrid | pair_111 | 1.3333 |
| c1_hybrid | pair_112 | 1.0000 |
| c1_hybrid | pair_113 | 0.8750 |
| c1_hybrid | pair_114 | 0.8571 |
| c1_hybrid | pair_115 | 0.8750 |
| c1_hybrid | pair_116 | 0.8571 |
| c1_hybrid | pair_117 | 1.2000 |
| c1_hybrid | pair_118 | 1.2000 |
| c1_hybrid | pair_119 | 1.0000 |
| c1_hybrid | pair_120 | 1.0000 |
| c1_hybrid | pair_121 | 1.0000 |
| c1_hybrid | pair_122 | 1.3333 |
| c1_hybrid | pair_123 | 1.3333 |
| c1_hybrid | pair_124 | 0.8750 |
| c1_hybrid | pair_125 | 0.8571 |
| c1_hybrid | pair_126 | 0.8750 |
| c1_hybrid | pair_127 | 0.8571 |
| c1_hybrid | pair_128 | 0.8750 |
| c1_hybrid | pair_129 | 1.2000 |
| c1_hybrid | pair_130 | 0.8750 |
| c1_hybrid | pair_131 | 0.8750 |
| c1_hybrid | pair_132 | 1.2000 |
| c1_hybrid | pair_133 | 1.2000 |
| c1_hybrid | pair_134 | 1.3333 |
| c1_hybrid | pair_135 | 1.0000 |
| c1_hybrid | pair_136 | 1.3333 |
| c1_hybrid | pair_137 | 1.3333 |
| c1_hybrid | pair_138 | 1.3333 |
| c1_hybrid | pair_139 | 1.3333 |
| c1_hybrid | pair_140 | 1.3333 |
| c1_hybrid | pair_141 | 1.3333 |
| c1_hybrid | pair_142 | 1.3333 |
| c1_hybrid | pair_143 | 1.3333 |
| c1_hybrid | pair_144 | 1.3333 |
| c1_hybrid | pair_145 | 1.3333 |
| c1_hybrid | pair_146 | 1.3333 |
| c1_hybrid | pair_147 | 1.3333 |
| c1_hybrid | pair_148 | 1.3333 |
| c1_hybrid | pair_149 | 1.3333 |
| c1_hybrid | pair_150 | 1.3333 |
| c1_hybrid | pair_151 | 1.3333 |
| c1_hybrid | pair_152 | 1.3333 |
| c1_hybrid | pair_153 | 1.3333 |
| c1_hybrid | pair_154 | 1.3333 |
| c1_hybrid | pair_155 | 1.3333 |
| c1_hybrid | pair_156 | 1.3333 |
| c1_hybrid | pair_157 | 1.3333 |
| c1_hybrid | pair_158 | 1.3333 |
| c1_hybrid | pair_159 | 1.3333 |
| c1_hybrid | pair_160 | 1.0000 |
| c1_hybrid | pair_161 | 1.0000 |
| c1_hybrid | pair_162 | 1.0000 |
| c1_hybrid | pair_163 | 1.0000 |
| c1_hybrid | pair_164 | 1.2000 |
| c1_hybrid | pair_165 | 1.2000 |
| c1_hybrid | pair_166 | 1.0000 |
| c1_hybrid | pair_167 | 1.0000 |
| c1_hybrid | pair_168 | 1.0000 |
| c1_hybrid | pair_169 | 1.3333 |
| c1_hybrid | pair_170 | 1.3333 |
| c1_hybrid | pair_171 | 1.0000 |
| c1_hybrid | pair_172 | 1.0000 |
| c1_hybrid | pair_173 | 1.0000 |
| c1_hybrid | pair_174 | 1.0000 |
| c1_hybrid | pair_175 | 1.0000 |
| c1_hybrid | pair_176 | 1.2000 |
| c1_hybrid | pair_177 | 1.0000 |
| c1_hybrid | pair_178 | 1.0000 |
| c1_hybrid | pair_179 | 1.2000 |
| c1_hybrid | pair_180 | 1.2000 |
| c1_hybrid | pair_181 | 1.3333 |
| c1_hybrid | pair_182 | 1.0000 |
| c1_hybrid | pair_183 | 0.8571 |
| c1_hybrid | pair_184 | 0.5625 |
| c1_hybrid | pair_185 | 0.8571 |
| c1_hybrid | pair_186 | 1.2000 |
| c1_hybrid | pair_187 | 1.2000 |
| c1_hybrid | pair_188 | 1.0000 |
| c1_hybrid | pair_189 | 1.0000 |
| c1_hybrid | pair_190 | 1.0000 |
| c1_hybrid | pair_191 | 1.3333 |
| c1_hybrid | pair_192 | 1.3333 |
| c1_hybrid | pair_193 | 0.7273 |
| c1_hybrid | pair_194 | 0.8571 |
| c1_hybrid | pair_195 | 0.5625 |
| c1_hybrid | pair_196 | 0.8571 |
| c1_hybrid | pair_197 | 0.5625 |
| c1_hybrid | pair_198 | 1.2000 |
| c1_hybrid | pair_199 | 0.6667 |
| c1_hybrid | pair_200 | 0.5625 |
| c1_hybrid | pair_201 | 1.2000 |
| c1_hybrid | pair_202 | 1.2000 |
| c1_hybrid | pair_203 | 1.3333 |
| c1_hybrid | pair_204 | 1.0000 |
| c1_hybrid | pair_205 | 0.8571 |
| c1_hybrid | pair_206 | 0.8571 |
| c1_hybrid | pair_207 | 1.2000 |
| c1_hybrid | pair_208 | 1.2000 |
| c1_hybrid | pair_209 | 1.0000 |
| c1_hybrid | pair_210 | 1.0000 |
| c1_hybrid | pair_211 | 1.0000 |
| c1_hybrid | pair_212 | 1.3333 |
| c1_hybrid | pair_213 | 1.3333 |
| c1_hybrid | pair_214 | 0.8571 |
| c1_hybrid | pair_215 | 0.8571 |
| c1_hybrid | pair_216 | 0.8571 |
| c1_hybrid | pair_217 | 0.8571 |
| c1_hybrid | pair_218 | 0.8571 |
| c1_hybrid | pair_219 | 1.2000 |
| c1_hybrid | pair_220 | 0.8571 |
| c1_hybrid | pair_221 | 0.8571 |
| c1_hybrid | pair_222 | 1.2000 |
| c1_hybrid | pair_223 | 1.2000 |
| c1_hybrid | pair_224 | 1.3333 |
| c1_hybrid | pair_225 | 1.0000 |
| c1_hybrid | pair_226 | 0.8571 |
| c1_hybrid | pair_227 | 1.2000 |
| c1_hybrid | pair_228 | 1.2000 |
| c1_hybrid | pair_229 | 1.0000 |
| c1_hybrid | pair_230 | 1.0000 |
| c1_hybrid | pair_231 | 1.0000 |
| c1_hybrid | pair_232 | 1.3333 |
| c1_hybrid | pair_233 | 1.3333 |
| c1_hybrid | pair_234 | 0.7273 |
| c1_hybrid | pair_235 | 0.8571 |
| c1_hybrid | pair_236 | 0.4762 |
| c1_hybrid | pair_237 | 0.8571 |
| c1_hybrid | pair_238 | 0.5000 |
| c1_hybrid | pair_239 | 1.2000 |
| c1_hybrid | pair_240 | 0.6667 |
| c1_hybrid | pair_241 | 0.4762 |
| c1_hybrid | pair_242 | 1.2000 |
| c1_hybrid | pair_243 | 1.2000 |
| c1_hybrid | pair_244 | 1.3333 |
| c1_hybrid | pair_245 | 1.0000 |
| c1_hybrid | pair_246 | 1.2000 |
| c1_hybrid | pair_247 | 1.2000 |
| c1_hybrid | pair_248 | 1.0000 |
| c1_hybrid | pair_249 | 1.0000 |
| c1_hybrid | pair_250 | 1.0000 |
| c1_hybrid | pair_251 | 1.3333 |
| c1_hybrid | pair_252 | 1.3333 |
| c1_hybrid | pair_253 | 0.8571 |
| c1_hybrid | pair_254 | 0.8571 |
| c1_hybrid | pair_255 | 0.8571 |
| c1_hybrid | pair_256 | 0.8571 |
| c1_hybrid | pair_257 | 0.8571 |
| c1_hybrid | pair_258 | 1.2000 |
| c1_hybrid | pair_259 | 0.8571 |
| c1_hybrid | pair_260 | 0.8571 |
| c1_hybrid | pair_261 | 1.2000 |
| c1_hybrid | pair_262 | 1.2000 |
| c1_hybrid | pair_263 | 1.3333 |
| c1_hybrid | pair_264 | 1.0000 |
| c1_hybrid | pair_265 | 1.2000 |
| c1_hybrid | pair_266 | 1.2000 |
| c1_hybrid | pair_267 | 1.2000 |
| c1_hybrid | pair_268 | 1.2000 |
| c1_hybrid | pair_269 | 1.3333 |
| c1_hybrid | pair_270 | 1.3333 |
| c1_hybrid | pair_271 | 1.2000 |
| c1_hybrid | pair_272 | 1.2000 |
| c1_hybrid | pair_273 | 1.2000 |
| c1_hybrid | pair_274 | 1.2000 |
| c1_hybrid | pair_275 | 1.2000 |
| c1_hybrid | pair_276 | 1.2000 |
| c1_hybrid | pair_277 | 1.2000 |
| c1_hybrid | pair_278 | 1.2000 |
| c1_hybrid | pair_279 | 1.2000 |
| c1_hybrid | pair_280 | 1.2000 |
| c1_hybrid | pair_281 | 1.3333 |
| c1_hybrid | pair_282 | 1.2000 |
| c1_hybrid | pair_283 | 1.2000 |
| c1_hybrid | pair_284 | 1.2000 |
| c1_hybrid | pair_285 | 1.2000 |
| c1_hybrid | pair_286 | 1.3333 |
| c1_hybrid | pair_287 | 1.3333 |
| c1_hybrid | pair_288 | 1.2000 |
| c1_hybrid | pair_289 | 1.2000 |
| c1_hybrid | pair_290 | 1.2000 |
| c1_hybrid | pair_291 | 1.2000 |
| c1_hybrid | pair_292 | 1.2000 |
| c1_hybrid | pair_293 | 1.2000 |
| c1_hybrid | pair_294 | 1.2000 |
| c1_hybrid | pair_295 | 1.2000 |
| c1_hybrid | pair_296 | 1.2000 |
| c1_hybrid | pair_297 | 1.2000 |
| c1_hybrid | pair_298 | 1.3333 |
| c1_hybrid | pair_299 | 1.2000 |
| c1_hybrid | pair_300 | 1.0000 |
| c1_hybrid | pair_301 | 1.0000 |
| c1_hybrid | pair_302 | 1.3333 |
| c1_hybrid | pair_303 | 1.3333 |
| c1_hybrid | pair_304 | 1.0000 |
| c1_hybrid | pair_305 | 1.0000 |
| c1_hybrid | pair_306 | 1.0000 |
| c1_hybrid | pair_307 | 1.0000 |
| c1_hybrid | pair_308 | 1.0000 |
| c1_hybrid | pair_309 | 1.2000 |
| c1_hybrid | pair_310 | 1.0000 |
| c1_hybrid | pair_311 | 1.0000 |
| c1_hybrid | pair_312 | 1.2000 |
| c1_hybrid | pair_313 | 1.2000 |
| c1_hybrid | pair_314 | 1.3333 |
| c1_hybrid | pair_315 | 1.0000 |
| c1_hybrid | pair_316 | 1.0000 |
| c1_hybrid | pair_317 | 1.3333 |
| c1_hybrid | pair_318 | 1.3333 |
| c1_hybrid | pair_319 | 1.0000 |
| c1_hybrid | pair_320 | 1.0000 |
| c1_hybrid | pair_321 | 1.0000 |
| c1_hybrid | pair_322 | 1.0000 |
| c1_hybrid | pair_323 | 1.0000 |
| c1_hybrid | pair_324 | 1.2000 |
| c1_hybrid | pair_325 | 1.0000 |
| c1_hybrid | pair_326 | 1.0000 |
| c1_hybrid | pair_327 | 1.2000 |
| c1_hybrid | pair_328 | 1.2000 |
| c1_hybrid | pair_329 | 1.3333 |
| c1_hybrid | pair_330 | 1.0000 |
| c1_hybrid | pair_331 | 1.3333 |
| c1_hybrid | pair_332 | 1.3333 |
| c1_hybrid | pair_333 | 1.0000 |
| c1_hybrid | pair_334 | 1.0000 |
| c1_hybrid | pair_335 | 1.0000 |
| c1_hybrid | pair_336 | 1.0000 |
| c1_hybrid | pair_337 | 1.0000 |
| c1_hybrid | pair_338 | 1.2000 |
| c1_hybrid | pair_339 | 1.0000 |
| c1_hybrid | pair_340 | 1.0000 |
| c1_hybrid | pair_341 | 1.2000 |
| c1_hybrid | pair_342 | 1.2000 |
| c1_hybrid | pair_343 | 1.3333 |
| c1_hybrid | pair_344 | 1.0000 |
| c1_hybrid | pair_345 | 1.3333 |
| c1_hybrid | pair_346 | 1.3333 |
| c1_hybrid | pair_347 | 1.3333 |
| c1_hybrid | pair_348 | 1.3333 |
| c1_hybrid | pair_349 | 1.3333 |
| c1_hybrid | pair_350 | 1.3333 |
| c1_hybrid | pair_351 | 1.3333 |
| c1_hybrid | pair_352 | 1.3333 |
| c1_hybrid | pair_353 | 1.3333 |
| c1_hybrid | pair_354 | 1.3333 |
| c1_hybrid | pair_355 | 1.3333 |
| c1_hybrid | pair_356 | 1.3333 |
| c1_hybrid | pair_357 | 1.3333 |
| c1_hybrid | pair_358 | 1.3333 |
| c1_hybrid | pair_359 | 1.3333 |
| c1_hybrid | pair_360 | 1.3333 |
| c1_hybrid | pair_361 | 1.3333 |
| c1_hybrid | pair_362 | 1.3333 |
| c1_hybrid | pair_363 | 1.3333 |
| c1_hybrid | pair_364 | 1.3333 |
| c1_hybrid | pair_365 | 1.3333 |
| c1_hybrid | pair_366 | 1.3333 |
| c1_hybrid | pair_367 | 1.3333 |
| c1_hybrid | pair_368 | 1.3333 |
| c1_hybrid | pair_369 | 1.3333 |
| c1_hybrid | pair_370 | 0.8571 |
| c1_hybrid | pair_371 | 0.7273 |
| c1_hybrid | pair_372 | 0.8571 |
| c1_hybrid | pair_373 | 0.7273 |
| c1_hybrid | pair_374 | 1.2000 |
| c1_hybrid | pair_375 | 0.7273 |
| c1_hybrid | pair_376 | 0.7273 |
| c1_hybrid | pair_377 | 1.2000 |
| c1_hybrid | pair_378 | 1.2000 |
| c1_hybrid | pair_379 | 1.3333 |
| c1_hybrid | pair_380 | 1.0000 |
| c1_hybrid | pair_381 | 0.8571 |
| c1_hybrid | pair_382 | 0.8571 |
| c1_hybrid | pair_383 | 0.8571 |
| c1_hybrid | pair_384 | 1.2000 |
| c1_hybrid | pair_385 | 0.8571 |
| c1_hybrid | pair_386 | 0.8571 |
| c1_hybrid | pair_387 | 1.2000 |
| c1_hybrid | pair_388 | 1.2000 |
| c1_hybrid | pair_389 | 1.3333 |
| c1_hybrid | pair_390 | 1.0000 |
| c1_hybrid | pair_391 | 0.8571 |
| c1_hybrid | pair_392 | 0.5000 |
| c1_hybrid | pair_393 | 1.2000 |
| c1_hybrid | pair_394 | 0.6667 |
| c1_hybrid | pair_395 | 0.4762 |
| c1_hybrid | pair_396 | 1.2000 |
| c1_hybrid | pair_397 | 1.2000 |
| c1_hybrid | pair_398 | 1.3333 |
| c1_hybrid | pair_399 | 1.0000 |
| c1_hybrid | pair_400 | 0.8571 |
| c1_hybrid | pair_401 | 1.2000 |
| c1_hybrid | pair_402 | 0.8571 |
| c1_hybrid | pair_403 | 0.8571 |
| c1_hybrid | pair_404 | 1.2000 |
| c1_hybrid | pair_405 | 1.2000 |
| c1_hybrid | pair_406 | 1.3333 |
| c1_hybrid | pair_407 | 1.0000 |
| c1_hybrid | pair_408 | 1.2000 |
| c1_hybrid | pair_409 | 0.6667 |
| c1_hybrid | pair_410 | 0.5000 |
| c1_hybrid | pair_411 | 1.2000 |
| c1_hybrid | pair_412 | 1.2000 |
| c1_hybrid | pair_413 | 1.3333 |
| c1_hybrid | pair_414 | 1.0000 |
| c1_hybrid | pair_415 | 1.2000 |
| c1_hybrid | pair_416 | 1.2000 |
| c1_hybrid | pair_417 | 1.2000 |
| c1_hybrid | pair_418 | 1.2000 |
| c1_hybrid | pair_419 | 1.3333 |
| c1_hybrid | pair_420 | 1.2000 |
| c1_hybrid | pair_421 | 0.6667 |
| c1_hybrid | pair_422 | 1.2000 |
| c1_hybrid | pair_423 | 1.2000 |
| c1_hybrid | pair_424 | 1.3333 |
| c1_hybrid | pair_425 | 1.0000 |
| c1_hybrid | pair_426 | 1.2000 |
| c1_hybrid | pair_427 | 1.2000 |
| c1_hybrid | pair_428 | 1.3333 |
| c1_hybrid | pair_429 | 1.0000 |
| c1_hybrid | pair_430 | 1.2000 |
| c1_hybrid | pair_431 | 1.3333 |
| c1_hybrid | pair_432 | 1.2000 |
| c1_hybrid | pair_433 | 1.3333 |
| c1_hybrid | pair_434 | 1.2000 |
| c1_hybrid | pair_435 | 1.3333 |
| c1_log | pair_001 | 0.5000 |
| c1_log | pair_002 | 0.5000 |
| c1_log | pair_003 | 0.4286 |
| c1_log | pair_004 | 0.5000 |
| c1_log | pair_005 | 0.6667 |
| c1_log | pair_006 | 0.5000 |
| c1_log | pair_007 | 0.3125 |
| c1_log | pair_008 | 0.4286 |
| c1_log | pair_009 | 0.2381 |
| c1_log | pair_010 | 0.4286 |
| c1_log | pair_011 | 0.6000 |
| c1_log | pair_012 | 0.6000 |
| c1_log | pair_013 | 0.5000 |
| c1_log | pair_014 | 0.5000 |
| c1_log | pair_015 | 0.5000 |
| c1_log | pair_016 | 0.6667 |
| c1_log | pair_017 | 0.6667 |
| c1_log | pair_018 | 0.3636 |
| c1_log | pair_019 | 0.4286 |
| c1_log | pair_020 | 0.2381 |
| c1_log | pair_021 | 0.4286 |
| c1_log | pair_022 | 0.2500 |
| c1_log | pair_023 | 0.6000 |
| c1_log | pair_024 | 0.3333 |
| c1_log | pair_025 | 0.2381 |
| c1_log | pair_026 | 0.6000 |
| c1_log | pair_027 | 0.6000 |
| c1_log | pair_028 | 0.6667 |
| c1_log | pair_029 | 0.5000 |
| c1_log | pair_030 | 0.5000 |
| c1_log | pair_031 | 0.4286 |
| c1_log | pair_032 | 0.5000 |
| c1_log | pair_033 | 0.6667 |
| c1_log | pair_034 | 0.5000 |
| c1_log | pair_035 | 0.5000 |
| c1_log | pair_036 | 0.4286 |
| c1_log | pair_037 | 0.5000 |
| c1_log | pair_038 | 0.4286 |
| c1_log | pair_039 | 0.6000 |
| c1_log | pair_040 | 0.6000 |
| c1_log | pair_041 | 0.5000 |
| c1_log | pair_042 | 0.5000 |
| c1_log | pair_043 | 0.5000 |
| c1_log | pair_044 | 0.6667 |
| c1_log | pair_045 | 0.6667 |
| c1_log | pair_046 | 0.5000 |
| c1_log | pair_047 | 0.4286 |
| c1_log | pair_048 | 0.5000 |
| c1_log | pair_049 | 0.4286 |
| c1_log | pair_050 | 0.5000 |
| c1_log | pair_051 | 0.6000 |
| c1_log | pair_052 | 0.5000 |
| c1_log | pair_053 | 0.5000 |
| c1_log | pair_054 | 0.6000 |
| c1_log | pair_055 | 0.6000 |
| c1_log | pair_056 | 0.6667 |
| c1_log | pair_057 | 0.5000 |
| c1_log | pair_058 | 0.5000 |
| c1_log | pair_059 | 0.5000 |
| c1_log | pair_060 | 0.6667 |
| c1_log | pair_061 | 0.5000 |
| c1_log | pair_062 | 0.5000 |
| c1_log | pair_063 | 0.5000 |
| c1_log | pair_064 | 0.5000 |
| c1_log | pair_065 | 0.5000 |
| c1_log | pair_066 | 0.6000 |
| c1_log | pair_067 | 0.6000 |
| c1_log | pair_068 | 0.5000 |
| c1_log | pair_069 | 0.5000 |
| c1_log | pair_070 | 0.5000 |
| c1_log | pair_071 | 0.6667 |
| c1_log | pair_072 | 0.6667 |
| c1_log | pair_073 | 0.5000 |
| c1_log | pair_074 | 0.5000 |
| c1_log | pair_075 | 0.5000 |
| c1_log | pair_076 | 0.5000 |
| c1_log | pair_077 | 0.5000 |
| c1_log | pair_078 | 0.6000 |
| c1_log | pair_079 | 0.5000 |
| c1_log | pair_080 | 0.5000 |
| c1_log | pair_081 | 0.6000 |
| c1_log | pair_082 | 0.6000 |
| c1_log | pair_083 | 0.6667 |
| c1_log | pair_084 | 0.5000 |
| c1_log | pair_085 | 0.4286 |
| c1_log | pair_086 | 0.6667 |
| c1_log | pair_087 | 0.5000 |
| c1_log | pair_088 | 0.4286 |
| c1_log | pair_089 | 0.4286 |
| c1_log | pair_090 | 0.4286 |
| c1_log | pair_091 | 0.4286 |
| c1_log | pair_092 | 0.6000 |
| c1_log | pair_093 | 0.6000 |
| c1_log | pair_094 | 0.5000 |
| c1_log | pair_095 | 0.5000 |
| c1_log | pair_096 | 0.5000 |
| c1_log | pair_097 | 0.6667 |
| c1_log | pair_098 | 0.6667 |
| c1_log | pair_099 | 0.4286 |
| c1_log | pair_100 | 0.4286 |
| c1_log | pair_101 | 0.4286 |
| c1_log | pair_102 | 0.4286 |
| c1_log | pair_103 | 0.4286 |
| c1_log | pair_104 | 0.6000 |
| c1_log | pair_105 | 0.4286 |
| c1_log | pair_106 | 0.4286 |
| c1_log | pair_107 | 0.6000 |
| c1_log | pair_108 | 0.6000 |
| c1_log | pair_109 | 0.6667 |
| c1_log | pair_110 | 0.5000 |
| c1_log | pair_111 | 0.6667 |
| c1_log | pair_112 | 0.5000 |
| c1_log | pair_113 | 0.5000 |
| c1_log | pair_114 | 0.4286 |
| c1_log | pair_115 | 0.5000 |
| c1_log | pair_116 | 0.4286 |
| c1_log | pair_117 | 0.6000 |
| c1_log | pair_118 | 0.6000 |
| c1_log | pair_119 | 0.5000 |
| c1_log | pair_120 | 0.5000 |
| c1_log | pair_121 | 0.5000 |
| c1_log | pair_122 | 0.6667 |
| c1_log | pair_123 | 0.6667 |
| c1_log | pair_124 | 0.5000 |
| c1_log | pair_125 | 0.4286 |
| c1_log | pair_126 | 0.5000 |
| c1_log | pair_127 | 0.4286 |
| c1_log | pair_128 | 0.5000 |
| c1_log | pair_129 | 0.6000 |
| c1_log | pair_130 | 0.5000 |
| c1_log | pair_131 | 0.5000 |
| c1_log | pair_132 | 0.6000 |
| c1_log | pair_133 | 0.6000 |
| c1_log | pair_134 | 0.6667 |
| c1_log | pair_135 | 0.5000 |
| c1_log | pair_136 | 0.6667 |
| c1_log | pair_137 | 0.6667 |
| c1_log | pair_138 | 0.6667 |
| c1_log | pair_139 | 0.6667 |
| c1_log | pair_140 | 0.6667 |
| c1_log | pair_141 | 0.6667 |
| c1_log | pair_142 | 0.6667 |
| c1_log | pair_143 | 0.6667 |
| c1_log | pair_144 | 0.6667 |
| c1_log | pair_145 | 0.6667 |
| c1_log | pair_146 | 0.6667 |
| c1_log | pair_147 | 0.6667 |
| c1_log | pair_148 | 0.6667 |
| c1_log | pair_149 | 0.6667 |
| c1_log | pair_150 | 0.6667 |
| c1_log | pair_151 | 0.6667 |
| c1_log | pair_152 | 0.6667 |
| c1_log | pair_153 | 0.6667 |
| c1_log | pair_154 | 0.6667 |
| c1_log | pair_155 | 0.6667 |
| c1_log | pair_156 | 0.6667 |
| c1_log | pair_157 | 0.6667 |
| c1_log | pair_158 | 0.6667 |
| c1_log | pair_159 | 0.6667 |
| c1_log | pair_160 | 0.5000 |
| c1_log | pair_161 | 0.5000 |
| c1_log | pair_162 | 0.5000 |
| c1_log | pair_163 | 0.5000 |
| c1_log | pair_164 | 0.6000 |
| c1_log | pair_165 | 0.6000 |
| c1_log | pair_166 | 0.5000 |
| c1_log | pair_167 | 0.5000 |
| c1_log | pair_168 | 0.5000 |
| c1_log | pair_169 | 0.6667 |
| c1_log | pair_170 | 0.6667 |
| c1_log | pair_171 | 0.5000 |
| c1_log | pair_172 | 0.5000 |
| c1_log | pair_173 | 0.5000 |
| c1_log | pair_174 | 0.5000 |
| c1_log | pair_175 | 0.5000 |
| c1_log | pair_176 | 0.6000 |
| c1_log | pair_177 | 0.5000 |
| c1_log | pair_178 | 0.5000 |
| c1_log | pair_179 | 0.6000 |
| c1_log | pair_180 | 0.6000 |
| c1_log | pair_181 | 0.6667 |
| c1_log | pair_182 | 0.5000 |
| c1_log | pair_183 | 0.4286 |
| c1_log | pair_184 | 0.3125 |
| c1_log | pair_185 | 0.4286 |
| c1_log | pair_186 | 0.6000 |
| c1_log | pair_187 | 0.6000 |
| c1_log | pair_188 | 0.5000 |
| c1_log | pair_189 | 0.5000 |
| c1_log | pair_190 | 0.5000 |
| c1_log | pair_191 | 0.6667 |
| c1_log | pair_192 | 0.6667 |
| c1_log | pair_193 | 0.3636 |
| c1_log | pair_194 | 0.4286 |
| c1_log | pair_195 | 0.3125 |
| c1_log | pair_196 | 0.4286 |
| c1_log | pair_197 | 0.3125 |
| c1_log | pair_198 | 0.6000 |
| c1_log | pair_199 | 0.3333 |
| c1_log | pair_200 | 0.3125 |
| c1_log | pair_201 | 0.6000 |
| c1_log | pair_202 | 0.6000 |
| c1_log | pair_203 | 0.6667 |
| c1_log | pair_204 | 0.5000 |
| c1_log | pair_205 | 0.4286 |
| c1_log | pair_206 | 0.4286 |
| c1_log | pair_207 | 0.6000 |
| c1_log | pair_208 | 0.6000 |
| c1_log | pair_209 | 0.5000 |
| c1_log | pair_210 | 0.5000 |
| c1_log | pair_211 | 0.5000 |
| c1_log | pair_212 | 0.6667 |
| c1_log | pair_213 | 0.6667 |
| c1_log | pair_214 | 0.4286 |
| c1_log | pair_215 | 0.4286 |
| c1_log | pair_216 | 0.4286 |
| c1_log | pair_217 | 0.4286 |
| c1_log | pair_218 | 0.4286 |
| c1_log | pair_219 | 0.6000 |
| c1_log | pair_220 | 0.4286 |
| c1_log | pair_221 | 0.4286 |
| c1_log | pair_222 | 0.6000 |
| c1_log | pair_223 | 0.6000 |
| c1_log | pair_224 | 0.6667 |
| c1_log | pair_225 | 0.5000 |
| c1_log | pair_226 | 0.4286 |
| c1_log | pair_227 | 0.6000 |
| c1_log | pair_228 | 0.6000 |
| c1_log | pair_229 | 0.5000 |
| c1_log | pair_230 | 0.5000 |
| c1_log | pair_231 | 0.5000 |
| c1_log | pair_232 | 0.6667 |
| c1_log | pair_233 | 0.6667 |
| c1_log | pair_234 | 0.3636 |
| c1_log | pair_235 | 0.4286 |
| c1_log | pair_236 | 0.2381 |
| c1_log | pair_237 | 0.4286 |
| c1_log | pair_238 | 0.2500 |
| c1_log | pair_239 | 0.6000 |
| c1_log | pair_240 | 0.3333 |
| c1_log | pair_241 | 0.2381 |
| c1_log | pair_242 | 0.6000 |
| c1_log | pair_243 | 0.6000 |
| c1_log | pair_244 | 0.6667 |
| c1_log | pair_245 | 0.5000 |
| c1_log | pair_246 | 0.6000 |
| c1_log | pair_247 | 0.6000 |
| c1_log | pair_248 | 0.5000 |
| c1_log | pair_249 | 0.5000 |
| c1_log | pair_250 | 0.5000 |
| c1_log | pair_251 | 0.6667 |
| c1_log | pair_252 | 0.6667 |
| c1_log | pair_253 | 0.4286 |
| c1_log | pair_254 | 0.4286 |
| c1_log | pair_255 | 0.4286 |
| c1_log | pair_256 | 0.4286 |
| c1_log | pair_257 | 0.4286 |
| c1_log | pair_258 | 0.6000 |
| c1_log | pair_259 | 0.4286 |
| c1_log | pair_260 | 0.4286 |
| c1_log | pair_261 | 0.6000 |
| c1_log | pair_262 | 0.6000 |
| c1_log | pair_263 | 0.6667 |
| c1_log | pair_264 | 0.5000 |
| c1_log | pair_265 | 0.6000 |
| c1_log | pair_266 | 0.6000 |
| c1_log | pair_267 | 0.6000 |
| c1_log | pair_268 | 0.6000 |
| c1_log | pair_269 | 0.6667 |
| c1_log | pair_270 | 0.6667 |
| c1_log | pair_271 | 0.6000 |
| c1_log | pair_272 | 0.6000 |
| c1_log | pair_273 | 0.6000 |
| c1_log | pair_274 | 0.6000 |
| c1_log | pair_275 | 0.6000 |
| c1_log | pair_276 | 0.6000 |
| c1_log | pair_277 | 0.6000 |
| c1_log | pair_278 | 0.6000 |
| c1_log | pair_279 | 0.6000 |
| c1_log | pair_280 | 0.6000 |
| c1_log | pair_281 | 0.6667 |
| c1_log | pair_282 | 0.6000 |
| c1_log | pair_283 | 0.6000 |
| c1_log | pair_284 | 0.6000 |
| c1_log | pair_285 | 0.6000 |
| c1_log | pair_286 | 0.6667 |
| c1_log | pair_287 | 0.6667 |
| c1_log | pair_288 | 0.6000 |
| c1_log | pair_289 | 0.6000 |
| c1_log | pair_290 | 0.6000 |
| c1_log | pair_291 | 0.6000 |
| c1_log | pair_292 | 0.6000 |
| c1_log | pair_293 | 0.6000 |
| c1_log | pair_294 | 0.6000 |
| c1_log | pair_295 | 0.6000 |
| c1_log | pair_296 | 0.6000 |
| c1_log | pair_297 | 0.6000 |
| c1_log | pair_298 | 0.6667 |
| c1_log | pair_299 | 0.6000 |
| c1_log | pair_300 | 0.5000 |
| c1_log | pair_301 | 0.5000 |
| c1_log | pair_302 | 0.6667 |
| c1_log | pair_303 | 0.6667 |
| c1_log | pair_304 | 0.5000 |
| c1_log | pair_305 | 0.5000 |
| c1_log | pair_306 | 0.5000 |
| c1_log | pair_307 | 0.5000 |
| c1_log | pair_308 | 0.5000 |
| c1_log | pair_309 | 0.6000 |
| c1_log | pair_310 | 0.5000 |
| c1_log | pair_311 | 0.5000 |
| c1_log | pair_312 | 0.6000 |
| c1_log | pair_313 | 0.6000 |
| c1_log | pair_314 | 0.6667 |
| c1_log | pair_315 | 0.5000 |
| c1_log | pair_316 | 0.5000 |
| c1_log | pair_317 | 0.6667 |
| c1_log | pair_318 | 0.6667 |
| c1_log | pair_319 | 0.5000 |
| c1_log | pair_320 | 0.5000 |
| c1_log | pair_321 | 0.5000 |
| c1_log | pair_322 | 0.5000 |
| c1_log | pair_323 | 0.5000 |
| c1_log | pair_324 | 0.6000 |
| c1_log | pair_325 | 0.5000 |
| c1_log | pair_326 | 0.5000 |
| c1_log | pair_327 | 0.6000 |
| c1_log | pair_328 | 0.6000 |
| c1_log | pair_329 | 0.6667 |
| c1_log | pair_330 | 0.5000 |
| c1_log | pair_331 | 0.6667 |
| c1_log | pair_332 | 0.6667 |
| c1_log | pair_333 | 0.5000 |
| c1_log | pair_334 | 0.5000 |
| c1_log | pair_335 | 0.5000 |
| c1_log | pair_336 | 0.5000 |
| c1_log | pair_337 | 0.5000 |
| c1_log | pair_338 | 0.6000 |
| c1_log | pair_339 | 0.5000 |
| c1_log | pair_340 | 0.5000 |
| c1_log | pair_341 | 0.6000 |
| c1_log | pair_342 | 0.6000 |
| c1_log | pair_343 | 0.6667 |
| c1_log | pair_344 | 0.5000 |
| c1_log | pair_345 | 0.6667 |
| c1_log | pair_346 | 0.6667 |
| c1_log | pair_347 | 0.6667 |
| c1_log | pair_348 | 0.6667 |
| c1_log | pair_349 | 0.6667 |
| c1_log | pair_350 | 0.6667 |
| c1_log | pair_351 | 0.6667 |
| c1_log | pair_352 | 0.6667 |
| c1_log | pair_353 | 0.6667 |
| c1_log | pair_354 | 0.6667 |
| c1_log | pair_355 | 0.6667 |
| c1_log | pair_356 | 0.6667 |
| c1_log | pair_357 | 0.6667 |
| c1_log | pair_358 | 0.6667 |
| c1_log | pair_359 | 0.6667 |
| c1_log | pair_360 | 0.6667 |
| c1_log | pair_361 | 0.6667 |
| c1_log | pair_362 | 0.6667 |
| c1_log | pair_363 | 0.6667 |
| c1_log | pair_364 | 0.6667 |
| c1_log | pair_365 | 0.6667 |
| c1_log | pair_366 | 0.6667 |
| c1_log | pair_367 | 0.6667 |
| c1_log | pair_368 | 0.6667 |
| c1_log | pair_369 | 0.6667 |
| c1_log | pair_370 | 0.4286 |
| c1_log | pair_371 | 0.3636 |
| c1_log | pair_372 | 0.4286 |
| c1_log | pair_373 | 0.3636 |
| c1_log | pair_374 | 0.6000 |
| c1_log | pair_375 | 0.3636 |
| c1_log | pair_376 | 0.3636 |
| c1_log | pair_377 | 0.6000 |
| c1_log | pair_378 | 0.6000 |
| c1_log | pair_379 | 0.6667 |
| c1_log | pair_380 | 0.5000 |
| c1_log | pair_381 | 0.4286 |
| c1_log | pair_382 | 0.4286 |
| c1_log | pair_383 | 0.4286 |
| c1_log | pair_384 | 0.6000 |
| c1_log | pair_385 | 0.4286 |
| c1_log | pair_386 | 0.4286 |
| c1_log | pair_387 | 0.6000 |
| c1_log | pair_388 | 0.6000 |
| c1_log | pair_389 | 0.6667 |
| c1_log | pair_390 | 0.5000 |
| c1_log | pair_391 | 0.4286 |
| c1_log | pair_392 | 0.2500 |
| c1_log | pair_393 | 0.6000 |
| c1_log | pair_394 | 0.3333 |
| c1_log | pair_395 | 0.2381 |
| c1_log | pair_396 | 0.6000 |
| c1_log | pair_397 | 0.6000 |
| c1_log | pair_398 | 0.6667 |
| c1_log | pair_399 | 0.5000 |
| c1_log | pair_400 | 0.4286 |
| c1_log | pair_401 | 0.6000 |
| c1_log | pair_402 | 0.4286 |
| c1_log | pair_403 | 0.4286 |
| c1_log | pair_404 | 0.6000 |
| c1_log | pair_405 | 0.6000 |
| c1_log | pair_406 | 0.6667 |
| c1_log | pair_407 | 0.5000 |
| c1_log | pair_408 | 0.6000 |
| c1_log | pair_409 | 0.3333 |
| c1_log | pair_410 | 0.2500 |
| c1_log | pair_411 | 0.6000 |
| c1_log | pair_412 | 0.6000 |
| c1_log | pair_413 | 0.6667 |
| c1_log | pair_414 | 0.5000 |
| c1_log | pair_415 | 0.6000 |
| c1_log | pair_416 | 0.6000 |
| c1_log | pair_417 | 0.6000 |
| c1_log | pair_418 | 0.6000 |
| c1_log | pair_419 | 0.6667 |
| c1_log | pair_420 | 0.6000 |
| c1_log | pair_421 | 0.3333 |
| c1_log | pair_422 | 0.6000 |
| c1_log | pair_423 | 0.6000 |
| c1_log | pair_424 | 0.6667 |
| c1_log | pair_425 | 0.5000 |
| c1_log | pair_426 | 0.6000 |
| c1_log | pair_427 | 0.6000 |
| c1_log | pair_428 | 0.6667 |
| c1_log | pair_429 | 0.5000 |
| c1_log | pair_430 | 0.6000 |
| c1_log | pair_431 | 0.6667 |
| c1_log | pair_432 | 0.6000 |
| c1_log | pair_433 | 0.6667 |
| c1_log | pair_434 | 0.6000 |
| c1_log | pair_435 | 0.6667 |
| c1_power2 | pair_001 | 0.0000 |
| c1_power2 | pair_002 | 0.3333 |
| c1_power2 | pair_003 | 0.1429 |
| c1_power2 | pair_004 | 0.0000 |
| c1_power2 | pair_005 | 0.3333 |
| c1_power2 | pair_006 | 0.3333 |
| c1_power2 | pair_007 | 0.0000 |
| c1_power2 | pair_008 | 0.1429 |
| c1_power2 | pair_009 | 0.5238 |
| c1_power2 | pair_010 | 0.1429 |
| c1_power2 | pair_011 | 0.6000 |
| c1_power2 | pair_012 | 0.6000 |
| c1_power2 | pair_013 | 0.3333 |
| c1_power2 | pair_014 | 0.3333 |
| c1_power2 | pair_015 | 0.3333 |
| c1_power2 | pair_016 | 0.3333 |
| c1_power2 | pair_017 | 0.3333 |
| c1_power2 | pair_018 | 0.4545 |
| c1_power2 | pair_019 | 0.1429 |
| c1_power2 | pair_020 | 0.5238 |
| c1_power2 | pair_021 | 0.1429 |
| c1_power2 | pair_022 | 0.6000 |
| c1_power2 | pair_023 | 0.6000 |
| c1_power2 | pair_024 | 0.3333 |
| c1_power2 | pair_025 | 0.5238 |
| c1_power2 | pair_026 | 0.6000 |
| c1_power2 | pair_027 | 0.6000 |
| c1_power2 | pair_028 | 0.3333 |
| c1_power2 | pair_029 | 0.3333 |
| c1_power2 | pair_030 | 0.3333 |
| c1_power2 | pair_031 | 0.1429 |
| c1_power2 | pair_032 | 0.0000 |
| c1_power2 | pair_033 | 0.3333 |
| c1_power2 | pair_034 | 0.3333 |
| c1_power2 | pair_035 | 0.0000 |
| c1_power2 | pair_036 | 0.1429 |
| c1_power2 | pair_037 | 0.0000 |
| c1_power2 | pair_038 | 0.1429 |
| c1_power2 | pair_039 | 0.6000 |
| c1_power2 | pair_040 | 0.6000 |
| c1_power2 | pair_041 | 0.3333 |
| c1_power2 | pair_042 | 0.3333 |
| c1_power2 | pair_043 | 0.3333 |
| c1_power2 | pair_044 | 0.3333 |
| c1_power2 | pair_045 | 0.3333 |
| c1_power2 | pair_046 | 0.0000 |
| c1_power2 | pair_047 | 0.1429 |
| c1_power2 | pair_048 | 0.0000 |
| c1_power2 | pair_049 | 0.1429 |
| c1_power2 | pair_050 | 0.0000 |
| c1_power2 | pair_051 | 0.6000 |
| c1_power2 | pair_052 | 0.0000 |
| c1_power2 | pair_053 | 0.0000 |
| c1_power2 | pair_054 | 0.6000 |
| c1_power2 | pair_055 | 0.6000 |
| c1_power2 | pair_056 | 0.3333 |
| c1_power2 | pair_057 | 0.3333 |
| c1_power2 | pair_058 | 0.3333 |
| c1_power2 | pair_059 | 0.3333 |
| c1_power2 | pair_060 | 0.3333 |
| c1_power2 | pair_061 | 0.3333 |
| c1_power2 | pair_062 | 0.3333 |
| c1_power2 | pair_063 | 0.3333 |
| c1_power2 | pair_064 | 0.3333 |
| c1_power2 | pair_065 | 0.3333 |
| c1_power2 | pair_066 | 0.6000 |
| c1_power2 | pair_067 | 0.6000 |
| c1_power2 | pair_068 | 0.3333 |
| c1_power2 | pair_069 | 0.3333 |
| c1_power2 | pair_070 | 0.3333 |
| c1_power2 | pair_071 | 0.3333 |
| c1_power2 | pair_072 | 0.3333 |
| c1_power2 | pair_073 | 0.3333 |
| c1_power2 | pair_074 | 0.3333 |
| c1_power2 | pair_075 | 0.3333 |
| c1_power2 | pair_076 | 0.3333 |
| c1_power2 | pair_077 | 0.3333 |
| c1_power2 | pair_078 | 0.6000 |
| c1_power2 | pair_079 | 0.3333 |
| c1_power2 | pair_080 | 0.3333 |
| c1_power2 | pair_081 | 0.6000 |
| c1_power2 | pair_082 | 0.6000 |
| c1_power2 | pair_083 | 0.3333 |
| c1_power2 | pair_084 | 0.3333 |
| c1_power2 | pair_085 | 0.1429 |
| c1_power2 | pair_086 | 0.3333 |
| c1_power2 | pair_087 | 0.3333 |
| c1_power2 | pair_088 | 0.1429 |
| c1_power2 | pair_089 | 0.1429 |
| c1_power2 | pair_090 | 0.1429 |
| c1_power2 | pair_091 | 0.1429 |
| c1_power2 | pair_092 | 0.6000 |
| c1_power2 | pair_093 | 0.6000 |
| c1_power2 | pair_094 | 0.3333 |
| c1_power2 | pair_095 | 0.3333 |
| c1_power2 | pair_096 | 0.3333 |
| c1_power2 | pair_097 | 0.3333 |
| c1_power2 | pair_098 | 0.3333 |
| c1_power2 | pair_099 | 0.1429 |
| c1_power2 | pair_100 | 0.1429 |
| c1_power2 | pair_101 | 0.1429 |
| c1_power2 | pair_102 | 0.1429 |
| c1_power2 | pair_103 | 0.1429 |
| c1_power2 | pair_104 | 0.6000 |
| c1_power2 | pair_105 | 0.1429 |
| c1_power2 | pair_106 | 0.1429 |
| c1_power2 | pair_107 | 0.6000 |
| c1_power2 | pair_108 | 0.6000 |
| c1_power2 | pair_109 | 0.3333 |
| c1_power2 | pair_110 | 0.3333 |
| c1_power2 | pair_111 | 0.3333 |
| c1_power2 | pair_112 | 0.3333 |
| c1_power2 | pair_113 | 0.0000 |
| c1_power2 | pair_114 | 0.1429 |
| c1_power2 | pair_115 | 0.0000 |
| c1_power2 | pair_116 | 0.1429 |
| c1_power2 | pair_117 | 0.6000 |
| c1_power2 | pair_118 | 0.6000 |
| c1_power2 | pair_119 | 0.3333 |
| c1_power2 | pair_120 | 0.3333 |
| c1_power2 | pair_121 | 0.3333 |
| c1_power2 | pair_122 | 0.3333 |
| c1_power2 | pair_123 | 0.3333 |
| c1_power2 | pair_124 | 0.0000 |
| c1_power2 | pair_125 | 0.1429 |
| c1_power2 | pair_126 | 0.0000 |
| c1_power2 | pair_127 | 0.1429 |
| c1_power2 | pair_128 | 0.0000 |
| c1_power2 | pair_129 | 0.6000 |
| c1_power2 | pair_130 | 0.0000 |
| c1_power2 | pair_131 | 0.0000 |
| c1_power2 | pair_132 | 0.6000 |
| c1_power2 | pair_133 | 0.6000 |
| c1_power2 | pair_134 | 0.3333 |
| c1_power2 | pair_135 | 0.3333 |
| c1_power2 | pair_136 | 0.3333 |
| c1_power2 | pair_137 | 0.3333 |
| c1_power2 | pair_138 | 0.3333 |
| c1_power2 | pair_139 | 0.3333 |
| c1_power2 | pair_140 | 0.3333 |
| c1_power2 | pair_141 | 0.3333 |
| c1_power2 | pair_142 | 0.3333 |
| c1_power2 | pair_143 | 0.3333 |
| c1_power2 | pair_144 | 0.3333 |
| c1_power2 | pair_145 | 0.3333 |
| c1_power2 | pair_146 | 0.3333 |
| c1_power2 | pair_147 | 0.3333 |
| c1_power2 | pair_148 | 0.3333 |
| c1_power2 | pair_149 | 0.3333 |
| c1_power2 | pair_150 | 0.3333 |
| c1_power2 | pair_151 | 0.3333 |
| c1_power2 | pair_152 | 0.3333 |
| c1_power2 | pair_153 | 0.3333 |
| c1_power2 | pair_154 | 0.3333 |
| c1_power2 | pair_155 | 0.3333 |
| c1_power2 | pair_156 | 0.3333 |
| c1_power2 | pair_157 | 0.3333 |
| c1_power2 | pair_158 | 0.3333 |
| c1_power2 | pair_159 | 0.3333 |
| c1_power2 | pair_160 | 0.3333 |
| c1_power2 | pair_161 | 0.3333 |
| c1_power2 | pair_162 | 0.3333 |
| c1_power2 | pair_163 | 0.3333 |
| c1_power2 | pair_164 | 0.6000 |
| c1_power2 | pair_165 | 0.6000 |
| c1_power2 | pair_166 | 0.3333 |
| c1_power2 | pair_167 | 0.3333 |
| c1_power2 | pair_168 | 0.3333 |
| c1_power2 | pair_169 | 0.3333 |
| c1_power2 | pair_170 | 0.3333 |
| c1_power2 | pair_171 | 0.3333 |
| c1_power2 | pair_172 | 0.3333 |
| c1_power2 | pair_173 | 0.3333 |
| c1_power2 | pair_174 | 0.3333 |
| c1_power2 | pair_175 | 0.3333 |
| c1_power2 | pair_176 | 0.6000 |
| c1_power2 | pair_177 | 0.3333 |
| c1_power2 | pair_178 | 0.3333 |
| c1_power2 | pair_179 | 0.6000 |
| c1_power2 | pair_180 | 0.6000 |
| c1_power2 | pair_181 | 0.3333 |
| c1_power2 | pair_182 | 0.3333 |
| c1_power2 | pair_183 | 0.1429 |
| c1_power2 | pair_184 | 0.0000 |
| c1_power2 | pair_185 | 0.1429 |
| c1_power2 | pair_186 | 0.6000 |
| c1_power2 | pair_187 | 0.6000 |
| c1_power2 | pair_188 | 0.3333 |
| c1_power2 | pair_189 | 0.3333 |
| c1_power2 | pair_190 | 0.3333 |
| c1_power2 | pair_191 | 0.3333 |
| c1_power2 | pair_192 | 0.3333 |
| c1_power2 | pair_193 | 0.4545 |
| c1_power2 | pair_194 | 0.1429 |
| c1_power2 | pair_195 | 0.0000 |
| c1_power2 | pair_196 | 0.1429 |
| c1_power2 | pair_197 | 0.0000 |
| c1_power2 | pair_198 | 0.6000 |
| c1_power2 | pair_199 | 0.3333 |
| c1_power2 | pair_200 | 0.0000 |
| c1_power2 | pair_201 | 0.6000 |
| c1_power2 | pair_202 | 0.6000 |
| c1_power2 | pair_203 | 0.3333 |
| c1_power2 | pair_204 | 0.3333 |
| c1_power2 | pair_205 | 0.1429 |
| c1_power2 | pair_206 | 0.1429 |
| c1_power2 | pair_207 | 0.6000 |
| c1_power2 | pair_208 | 0.6000 |
| c1_power2 | pair_209 | 0.3333 |
| c1_power2 | pair_210 | 0.3333 |
| c1_power2 | pair_211 | 0.3333 |
| c1_power2 | pair_212 | 0.3333 |
| c1_power2 | pair_213 | 0.3333 |
| c1_power2 | pair_214 | 0.1429 |
| c1_power2 | pair_215 | 0.1429 |
| c1_power2 | pair_216 | 0.1429 |
| c1_power2 | pair_217 | 0.1429 |
| c1_power2 | pair_218 | 0.1429 |
| c1_power2 | pair_219 | 0.6000 |
| c1_power2 | pair_220 | 0.1429 |
| c1_power2 | pair_221 | 0.1429 |
| c1_power2 | pair_222 | 0.6000 |
| c1_power2 | pair_223 | 0.6000 |
| c1_power2 | pair_224 | 0.3333 |
| c1_power2 | pair_225 | 0.3333 |
| c1_power2 | pair_226 | 0.1429 |
| c1_power2 | pair_227 | 0.6000 |
| c1_power2 | pair_228 | 0.6000 |
| c1_power2 | pair_229 | 0.3333 |
| c1_power2 | pair_230 | 0.3333 |
| c1_power2 | pair_231 | 0.3333 |
| c1_power2 | pair_232 | 0.3333 |
| c1_power2 | pair_233 | 0.3333 |
| c1_power2 | pair_234 | 0.4545 |
| c1_power2 | pair_235 | 0.1429 |
| c1_power2 | pair_236 | 0.5238 |
| c1_power2 | pair_237 | 0.1429 |
| c1_power2 | pair_238 | 0.6000 |
| c1_power2 | pair_239 | 0.6000 |
| c1_power2 | pair_240 | 0.3333 |
| c1_power2 | pair_241 | 0.5238 |
| c1_power2 | pair_242 | 0.6000 |
| c1_power2 | pair_243 | 0.6000 |
| c1_power2 | pair_244 | 0.3333 |
| c1_power2 | pair_245 | 0.3333 |
| c1_power2 | pair_246 | 0.6000 |
| c1_power2 | pair_247 | 0.6000 |
| c1_power2 | pair_248 | 0.3333 |
| c1_power2 | pair_249 | 0.3333 |
| c1_power2 | pair_250 | 0.3333 |
| c1_power2 | pair_251 | 0.3333 |
| c1_power2 | pair_252 | 0.3333 |
| c1_power2 | pair_253 | 0.1429 |
| c1_power2 | pair_254 | 0.1429 |
| c1_power2 | pair_255 | 0.1429 |
| c1_power2 | pair_256 | 0.1429 |
| c1_power2 | pair_257 | 0.1429 |
| c1_power2 | pair_258 | 0.6000 |
| c1_power2 | pair_259 | 0.1429 |
| c1_power2 | pair_260 | 0.1429 |
| c1_power2 | pair_261 | 0.6000 |
| c1_power2 | pair_262 | 0.6000 |
| c1_power2 | pair_263 | 0.3333 |
| c1_power2 | pair_264 | 0.3333 |
| c1_power2 | pair_265 | 0.6000 |
| c1_power2 | pair_266 | 0.6000 |
| c1_power2 | pair_267 | 0.6000 |
| c1_power2 | pair_268 | 0.6000 |
| c1_power2 | pair_269 | 0.3333 |
| c1_power2 | pair_270 | 0.3333 |
| c1_power2 | pair_271 | 0.6000 |
| c1_power2 | pair_272 | 0.6000 |
| c1_power2 | pair_273 | 0.6000 |
| c1_power2 | pair_274 | 0.6000 |
| c1_power2 | pair_275 | 0.6000 |
| c1_power2 | pair_276 | 0.6000 |
| c1_power2 | pair_277 | 0.6000 |
| c1_power2 | pair_278 | 0.6000 |
| c1_power2 | pair_279 | 0.6000 |
| c1_power2 | pair_280 | 0.6000 |
| c1_power2 | pair_281 | 0.3333 |
| c1_power2 | pair_282 | 0.6000 |
| c1_power2 | pair_283 | 0.6000 |
| c1_power2 | pair_284 | 0.6000 |
| c1_power2 | pair_285 | 0.6000 |
| c1_power2 | pair_286 | 0.3333 |
| c1_power2 | pair_287 | 0.3333 |
| c1_power2 | pair_288 | 0.6000 |
| c1_power2 | pair_289 | 0.6000 |
| c1_power2 | pair_290 | 0.6000 |
| c1_power2 | pair_291 | 0.6000 |
| c1_power2 | pair_292 | 0.6000 |
| c1_power2 | pair_293 | 0.6000 |
| c1_power2 | pair_294 | 0.6000 |
| c1_power2 | pair_295 | 0.6000 |
| c1_power2 | pair_296 | 0.6000 |
| c1_power2 | pair_297 | 0.6000 |
| c1_power2 | pair_298 | 0.3333 |
| c1_power2 | pair_299 | 0.6000 |
| c1_power2 | pair_300 | 0.3333 |
| c1_power2 | pair_301 | 0.3333 |
| c1_power2 | pair_302 | 0.3333 |
| c1_power2 | pair_303 | 0.3333 |
| c1_power2 | pair_304 | 0.3333 |
| c1_power2 | pair_305 | 0.3333 |
| c1_power2 | pair_306 | 0.3333 |
| c1_power2 | pair_307 | 0.3333 |
| c1_power2 | pair_308 | 0.3333 |
| c1_power2 | pair_309 | 0.6000 |
| c1_power2 | pair_310 | 0.3333 |
| c1_power2 | pair_311 | 0.3333 |
| c1_power2 | pair_312 | 0.6000 |
| c1_power2 | pair_313 | 0.6000 |
| c1_power2 | pair_314 | 0.3333 |
| c1_power2 | pair_315 | 0.3333 |
| c1_power2 | pair_316 | 0.3333 |
| c1_power2 | pair_317 | 0.3333 |
| c1_power2 | pair_318 | 0.3333 |
| c1_power2 | pair_319 | 0.3333 |
| c1_power2 | pair_320 | 0.3333 |
| c1_power2 | pair_321 | 0.3333 |
| c1_power2 | pair_322 | 0.3333 |
| c1_power2 | pair_323 | 0.3333 |
| c1_power2 | pair_324 | 0.6000 |
| c1_power2 | pair_325 | 0.3333 |
| c1_power2 | pair_326 | 0.3333 |
| c1_power2 | pair_327 | 0.6000 |
| c1_power2 | pair_328 | 0.6000 |
| c1_power2 | pair_329 | 0.3333 |
| c1_power2 | pair_330 | 0.3333 |
| c1_power2 | pair_331 | 0.3333 |
| c1_power2 | pair_332 | 0.3333 |
| c1_power2 | pair_333 | 0.3333 |
| c1_power2 | pair_334 | 0.3333 |
| c1_power2 | pair_335 | 0.3333 |
| c1_power2 | pair_336 | 0.3333 |
| c1_power2 | pair_337 | 0.3333 |
| c1_power2 | pair_338 | 0.6000 |
| c1_power2 | pair_339 | 0.3333 |
| c1_power2 | pair_340 | 0.3333 |
| c1_power2 | pair_341 | 0.6000 |
| c1_power2 | pair_342 | 0.6000 |
| c1_power2 | pair_343 | 0.3333 |
| c1_power2 | pair_344 | 0.3333 |
| c1_power2 | pair_345 | 0.3333 |
| c1_power2 | pair_346 | 0.3333 |
| c1_power2 | pair_347 | 0.3333 |
| c1_power2 | pair_348 | 0.3333 |
| c1_power2 | pair_349 | 0.3333 |
| c1_power2 | pair_350 | 0.3333 |
| c1_power2 | pair_351 | 0.3333 |
| c1_power2 | pair_352 | 0.3333 |
| c1_power2 | pair_353 | 0.3333 |
| c1_power2 | pair_354 | 0.3333 |
| c1_power2 | pair_355 | 0.3333 |
| c1_power2 | pair_356 | 0.3333 |
| c1_power2 | pair_357 | 0.3333 |
| c1_power2 | pair_358 | 0.3333 |
| c1_power2 | pair_359 | 0.3333 |
| c1_power2 | pair_360 | 0.3333 |
| c1_power2 | pair_361 | 0.3333 |
| c1_power2 | pair_362 | 0.3333 |
| c1_power2 | pair_363 | 0.3333 |
| c1_power2 | pair_364 | 0.3333 |
| c1_power2 | pair_365 | 0.3333 |
| c1_power2 | pair_366 | 0.3333 |
| c1_power2 | pair_367 | 0.3333 |
| c1_power2 | pair_368 | 0.3333 |
| c1_power2 | pair_369 | 0.3333 |
| c1_power2 | pair_370 | 0.1429 |
| c1_power2 | pair_371 | 0.4545 |
| c1_power2 | pair_372 | 0.1429 |
| c1_power2 | pair_373 | 0.4545 |
| c1_power2 | pair_374 | 0.6000 |
| c1_power2 | pair_375 | 0.4545 |
| c1_power2 | pair_376 | 0.4545 |
| c1_power2 | pair_377 | 0.6000 |
| c1_power2 | pair_378 | 0.6000 |
| c1_power2 | pair_379 | 0.3333 |
| c1_power2 | pair_380 | 0.3333 |
| c1_power2 | pair_381 | 0.1429 |
| c1_power2 | pair_382 | 0.1429 |
| c1_power2 | pair_383 | 0.1429 |
| c1_power2 | pair_384 | 0.6000 |
| c1_power2 | pair_385 | 0.1429 |
| c1_power2 | pair_386 | 0.1429 |
| c1_power2 | pair_387 | 0.6000 |
| c1_power2 | pair_388 | 0.6000 |
| c1_power2 | pair_389 | 0.3333 |
| c1_power2 | pair_390 | 0.3333 |
| c1_power2 | pair_391 | 0.1429 |
| c1_power2 | pair_392 | 0.6000 |
| c1_power2 | pair_393 | 0.6000 |
| c1_power2 | pair_394 | 0.3333 |
| c1_power2 | pair_395 | 0.5238 |
| c1_power2 | pair_396 | 0.6000 |
| c1_power2 | pair_397 | 0.6000 |
| c1_power2 | pair_398 | 0.3333 |
| c1_power2 | pair_399 | 0.3333 |
| c1_power2 | pair_400 | 0.1429 |
| c1_power2 | pair_401 | 0.6000 |
| c1_power2 | pair_402 | 0.1429 |
| c1_power2 | pair_403 | 0.1429 |
| c1_power2 | pair_404 | 0.6000 |
| c1_power2 | pair_405 | 0.6000 |
| c1_power2 | pair_406 | 0.3333 |
| c1_power2 | pair_407 | 0.3333 |
| c1_power2 | pair_408 | 0.6000 |
| c1_power2 | pair_409 | 0.3333 |
| c1_power2 | pair_410 | 0.6000 |
| c1_power2 | pair_411 | 0.6000 |
| c1_power2 | pair_412 | 0.6000 |
| c1_power2 | pair_413 | 0.3333 |
| c1_power2 | pair_414 | 0.3333 |
| c1_power2 | pair_415 | 0.6000 |
| c1_power2 | pair_416 | 0.6000 |
| c1_power2 | pair_417 | 0.6000 |
| c1_power2 | pair_418 | 0.6000 |
| c1_power2 | pair_419 | 0.3333 |
| c1_power2 | pair_420 | 0.6000 |
| c1_power2 | pair_421 | 0.3333 |
| c1_power2 | pair_422 | 0.6000 |
| c1_power2 | pair_423 | 0.6000 |
| c1_power2 | pair_424 | 0.3333 |
| c1_power2 | pair_425 | 0.3333 |
| c1_power2 | pair_426 | 0.6000 |
| c1_power2 | pair_427 | 0.6000 |
| c1_power2 | pair_428 | 0.3333 |
| c1_power2 | pair_429 | 0.3333 |
| c1_power2 | pair_430 | 0.6000 |
| c1_power2 | pair_431 | 0.3333 |
| c1_power2 | pair_432 | 0.6000 |
| c1_power2 | pair_433 | 0.3333 |
| c1_power2 | pair_434 | 0.6000 |
| c1_power2 | pair_435 | 0.3333 |
| c1_sqrt | pair_001 | 0.3750 |
| c1_sqrt | pair_002 | 0.5000 |
| c1_sqrt | pair_003 | 0.4286 |
| c1_sqrt | pair_004 | 0.3750 |
| c1_sqrt | pair_005 | 0.6667 |
| c1_sqrt | pair_006 | 0.5000 |
| c1_sqrt | pair_007 | 0.2500 |
| c1_sqrt | pair_008 | 0.4286 |
| c1_sqrt | pair_009 | 0.2381 |
| c1_sqrt | pair_010 | 0.4286 |
| c1_sqrt | pair_011 | 0.6000 |
| c1_sqrt | pair_012 | 0.6000 |
| c1_sqrt | pair_013 | 0.5000 |
| c1_sqrt | pair_014 | 0.5000 |
| c1_sqrt | pair_015 | 0.5000 |
| c1_sqrt | pair_016 | 0.6667 |
| c1_sqrt | pair_017 | 0.6667 |
| c1_sqrt | pair_018 | 0.3636 |
| c1_sqrt | pair_019 | 0.4286 |
| c1_sqrt | pair_020 | 0.2381 |
| c1_sqrt | pair_021 | 0.4286 |
| c1_sqrt | pair_022 | 0.2500 |
| c1_sqrt | pair_023 | 0.6000 |
| c1_sqrt | pair_024 | 0.3333 |
| c1_sqrt | pair_025 | 0.2381 |
| c1_sqrt | pair_026 | 0.6000 |
| c1_sqrt | pair_027 | 0.6000 |
| c1_sqrt | pair_028 | 0.6667 |
| c1_sqrt | pair_029 | 0.5000 |
| c1_sqrt | pair_030 | 0.5000 |
| c1_sqrt | pair_031 | 0.4286 |
| c1_sqrt | pair_032 | 0.3750 |
| c1_sqrt | pair_033 | 0.6667 |
| c1_sqrt | pair_034 | 0.5000 |
| c1_sqrt | pair_035 | 0.3750 |
| c1_sqrt | pair_036 | 0.4286 |
| c1_sqrt | pair_037 | 0.3750 |
| c1_sqrt | pair_038 | 0.4286 |
| c1_sqrt | pair_039 | 0.6000 |
| c1_sqrt | pair_040 | 0.6000 |
| c1_sqrt | pair_041 | 0.5000 |
| c1_sqrt | pair_042 | 0.5000 |
| c1_sqrt | pair_043 | 0.5000 |
| c1_sqrt | pair_044 | 0.6667 |
| c1_sqrt | pair_045 | 0.6667 |
| c1_sqrt | pair_046 | 0.3750 |
| c1_sqrt | pair_047 | 0.4286 |
| c1_sqrt | pair_048 | 0.3750 |
| c1_sqrt | pair_049 | 0.4286 |
| c1_sqrt | pair_050 | 0.3750 |
| c1_sqrt | pair_051 | 0.6000 |
| c1_sqrt | pair_052 | 0.3750 |
| c1_sqrt | pair_053 | 0.3750 |
| c1_sqrt | pair_054 | 0.6000 |
| c1_sqrt | pair_055 | 0.6000 |
| c1_sqrt | pair_056 | 0.6667 |
| c1_sqrt | pair_057 | 0.5000 |
| c1_sqrt | pair_058 | 0.5000 |
| c1_sqrt | pair_059 | 0.5000 |
| c1_sqrt | pair_060 | 0.6667 |
| c1_sqrt | pair_061 | 0.5000 |
| c1_sqrt | pair_062 | 0.5000 |
| c1_sqrt | pair_063 | 0.5000 |
| c1_sqrt | pair_064 | 0.5000 |
| c1_sqrt | pair_065 | 0.5000 |
| c1_sqrt | pair_066 | 0.6000 |
| c1_sqrt | pair_067 | 0.6000 |
| c1_sqrt | pair_068 | 0.5000 |
| c1_sqrt | pair_069 | 0.5000 |
| c1_sqrt | pair_070 | 0.5000 |
| c1_sqrt | pair_071 | 0.6667 |
| c1_sqrt | pair_072 | 0.6667 |
| c1_sqrt | pair_073 | 0.5000 |
| c1_sqrt | pair_074 | 0.5000 |
| c1_sqrt | pair_075 | 0.5000 |
| c1_sqrt | pair_076 | 0.5000 |
| c1_sqrt | pair_077 | 0.5000 |
| c1_sqrt | pair_078 | 0.6000 |
| c1_sqrt | pair_079 | 0.5000 |
| c1_sqrt | pair_080 | 0.5000 |
| c1_sqrt | pair_081 | 0.6000 |
| c1_sqrt | pair_082 | 0.6000 |
| c1_sqrt | pair_083 | 0.6667 |
| c1_sqrt | pair_084 | 0.5000 |
| c1_sqrt | pair_085 | 0.4286 |
| c1_sqrt | pair_086 | 0.6667 |
| c1_sqrt | pair_087 | 0.5000 |
| c1_sqrt | pair_088 | 0.4286 |
| c1_sqrt | pair_089 | 0.4286 |
| c1_sqrt | pair_090 | 0.4286 |
| c1_sqrt | pair_091 | 0.4286 |
| c1_sqrt | pair_092 | 0.6000 |
| c1_sqrt | pair_093 | 0.6000 |
| c1_sqrt | pair_094 | 0.5000 |
| c1_sqrt | pair_095 | 0.5000 |
| c1_sqrt | pair_096 | 0.5000 |
| c1_sqrt | pair_097 | 0.6667 |
| c1_sqrt | pair_098 | 0.6667 |
| c1_sqrt | pair_099 | 0.4286 |
| c1_sqrt | pair_100 | 0.4286 |
| c1_sqrt | pair_101 | 0.4286 |
| c1_sqrt | pair_102 | 0.4286 |
| c1_sqrt | pair_103 | 0.4286 |
| c1_sqrt | pair_104 | 0.6000 |
| c1_sqrt | pair_105 | 0.4286 |
| c1_sqrt | pair_106 | 0.4286 |
| c1_sqrt | pair_107 | 0.6000 |
| c1_sqrt | pair_108 | 0.6000 |
| c1_sqrt | pair_109 | 0.6667 |
| c1_sqrt | pair_110 | 0.5000 |
| c1_sqrt | pair_111 | 0.6667 |
| c1_sqrt | pair_112 | 0.5000 |
| c1_sqrt | pair_113 | 0.3750 |
| c1_sqrt | pair_114 | 0.4286 |
| c1_sqrt | pair_115 | 0.3750 |
| c1_sqrt | pair_116 | 0.4286 |
| c1_sqrt | pair_117 | 0.6000 |
| c1_sqrt | pair_118 | 0.6000 |
| c1_sqrt | pair_119 | 0.5000 |
| c1_sqrt | pair_120 | 0.5000 |
| c1_sqrt | pair_121 | 0.5000 |
| c1_sqrt | pair_122 | 0.6667 |
| c1_sqrt | pair_123 | 0.6667 |
| c1_sqrt | pair_124 | 0.3750 |
| c1_sqrt | pair_125 | 0.4286 |
| c1_sqrt | pair_126 | 0.3750 |
| c1_sqrt | pair_127 | 0.4286 |
| c1_sqrt | pair_128 | 0.3750 |
| c1_sqrt | pair_129 | 0.6000 |
| c1_sqrt | pair_130 | 0.3750 |
| c1_sqrt | pair_131 | 0.3750 |
| c1_sqrt | pair_132 | 0.6000 |
| c1_sqrt | pair_133 | 0.6000 |
| c1_sqrt | pair_134 | 0.6667 |
| c1_sqrt | pair_135 | 0.5000 |
| c1_sqrt | pair_136 | 0.6667 |
| c1_sqrt | pair_137 | 0.6667 |
| c1_sqrt | pair_138 | 0.6667 |
| c1_sqrt | pair_139 | 0.6667 |
| c1_sqrt | pair_140 | 0.6667 |
| c1_sqrt | pair_141 | 0.6667 |
| c1_sqrt | pair_142 | 0.6667 |
| c1_sqrt | pair_143 | 0.6667 |
| c1_sqrt | pair_144 | 0.6667 |
| c1_sqrt | pair_145 | 0.6667 |
| c1_sqrt | pair_146 | 0.6667 |
| c1_sqrt | pair_147 | 0.6667 |
| c1_sqrt | pair_148 | 0.6667 |
| c1_sqrt | pair_149 | 0.6667 |
| c1_sqrt | pair_150 | 0.6667 |
| c1_sqrt | pair_151 | 0.6667 |
| c1_sqrt | pair_152 | 0.6667 |
| c1_sqrt | pair_153 | 0.6667 |
| c1_sqrt | pair_154 | 0.6667 |
| c1_sqrt | pair_155 | 0.6667 |
| c1_sqrt | pair_156 | 0.6667 |
| c1_sqrt | pair_157 | 0.6667 |
| c1_sqrt | pair_158 | 0.6667 |
| c1_sqrt | pair_159 | 0.6667 |
| c1_sqrt | pair_160 | 0.5000 |
| c1_sqrt | pair_161 | 0.5000 |
| c1_sqrt | pair_162 | 0.5000 |
| c1_sqrt | pair_163 | 0.5000 |
| c1_sqrt | pair_164 | 0.6000 |
| c1_sqrt | pair_165 | 0.6000 |
| c1_sqrt | pair_166 | 0.5000 |
| c1_sqrt | pair_167 | 0.5000 |
| c1_sqrt | pair_168 | 0.5000 |
| c1_sqrt | pair_169 | 0.6667 |
| c1_sqrt | pair_170 | 0.6667 |
| c1_sqrt | pair_171 | 0.5000 |
| c1_sqrt | pair_172 | 0.5000 |
| c1_sqrt | pair_173 | 0.5000 |
| c1_sqrt | pair_174 | 0.5000 |
| c1_sqrt | pair_175 | 0.5000 |
| c1_sqrt | pair_176 | 0.6000 |
| c1_sqrt | pair_177 | 0.5000 |
| c1_sqrt | pair_178 | 0.5000 |
| c1_sqrt | pair_179 | 0.6000 |
| c1_sqrt | pair_180 | 0.6000 |
| c1_sqrt | pair_181 | 0.6667 |
| c1_sqrt | pair_182 | 0.5000 |
| c1_sqrt | pair_183 | 0.4286 |
| c1_sqrt | pair_184 | 0.2500 |
| c1_sqrt | pair_185 | 0.4286 |
| c1_sqrt | pair_186 | 0.6000 |
| c1_sqrt | pair_187 | 0.6000 |
| c1_sqrt | pair_188 | 0.5000 |
| c1_sqrt | pair_189 | 0.5000 |
| c1_sqrt | pair_190 | 0.5000 |
| c1_sqrt | pair_191 | 0.6667 |
| c1_sqrt | pair_192 | 0.6667 |
| c1_sqrt | pair_193 | 0.3636 |
| c1_sqrt | pair_194 | 0.4286 |
| c1_sqrt | pair_195 | 0.2500 |
| c1_sqrt | pair_196 | 0.4286 |
| c1_sqrt | pair_197 | 0.2500 |
| c1_sqrt | pair_198 | 0.6000 |
| c1_sqrt | pair_199 | 0.3333 |
| c1_sqrt | pair_200 | 0.2500 |
| c1_sqrt | pair_201 | 0.6000 |
| c1_sqrt | pair_202 | 0.6000 |
| c1_sqrt | pair_203 | 0.6667 |
| c1_sqrt | pair_204 | 0.5000 |
| c1_sqrt | pair_205 | 0.4286 |
| c1_sqrt | pair_206 | 0.4286 |
| c1_sqrt | pair_207 | 0.6000 |
| c1_sqrt | pair_208 | 0.6000 |
| c1_sqrt | pair_209 | 0.5000 |
| c1_sqrt | pair_210 | 0.5000 |
| c1_sqrt | pair_211 | 0.5000 |
| c1_sqrt | pair_212 | 0.6667 |
| c1_sqrt | pair_213 | 0.6667 |
| c1_sqrt | pair_214 | 0.4286 |
| c1_sqrt | pair_215 | 0.4286 |
| c1_sqrt | pair_216 | 0.4286 |
| c1_sqrt | pair_217 | 0.4286 |
| c1_sqrt | pair_218 | 0.4286 |
| c1_sqrt | pair_219 | 0.6000 |
| c1_sqrt | pair_220 | 0.4286 |
| c1_sqrt | pair_221 | 0.4286 |
| c1_sqrt | pair_222 | 0.6000 |
| c1_sqrt | pair_223 | 0.6000 |
| c1_sqrt | pair_224 | 0.6667 |
| c1_sqrt | pair_225 | 0.5000 |
| c1_sqrt | pair_226 | 0.4286 |
| c1_sqrt | pair_227 | 0.6000 |
| c1_sqrt | pair_228 | 0.6000 |
| c1_sqrt | pair_229 | 0.5000 |
| c1_sqrt | pair_230 | 0.5000 |
| c1_sqrt | pair_231 | 0.5000 |
| c1_sqrt | pair_232 | 0.6667 |
| c1_sqrt | pair_233 | 0.6667 |
| c1_sqrt | pair_234 | 0.3636 |
| c1_sqrt | pair_235 | 0.4286 |
| c1_sqrt | pair_236 | 0.2381 |
| c1_sqrt | pair_237 | 0.4286 |
| c1_sqrt | pair_238 | 0.2500 |
| c1_sqrt | pair_239 | 0.6000 |
| c1_sqrt | pair_240 | 0.3333 |
| c1_sqrt | pair_241 | 0.2381 |
| c1_sqrt | pair_242 | 0.6000 |
| c1_sqrt | pair_243 | 0.6000 |
| c1_sqrt | pair_244 | 0.6667 |
| c1_sqrt | pair_245 | 0.5000 |
| c1_sqrt | pair_246 | 0.6000 |
| c1_sqrt | pair_247 | 0.6000 |
| c1_sqrt | pair_248 | 0.5000 |
| c1_sqrt | pair_249 | 0.5000 |
| c1_sqrt | pair_250 | 0.5000 |
| c1_sqrt | pair_251 | 0.6667 |
| c1_sqrt | pair_252 | 0.6667 |
| c1_sqrt | pair_253 | 0.4286 |
| c1_sqrt | pair_254 | 0.4286 |
| c1_sqrt | pair_255 | 0.4286 |
| c1_sqrt | pair_256 | 0.4286 |
| c1_sqrt | pair_257 | 0.4286 |
| c1_sqrt | pair_258 | 0.6000 |
| c1_sqrt | pair_259 | 0.4286 |
| c1_sqrt | pair_260 | 0.4286 |
| c1_sqrt | pair_261 | 0.6000 |
| c1_sqrt | pair_262 | 0.6000 |
| c1_sqrt | pair_263 | 0.6667 |
| c1_sqrt | pair_264 | 0.5000 |
| c1_sqrt | pair_265 | 0.6000 |
| c1_sqrt | pair_266 | 0.6000 |
| c1_sqrt | pair_267 | 0.6000 |
| c1_sqrt | pair_268 | 0.6000 |
| c1_sqrt | pair_269 | 0.6667 |
| c1_sqrt | pair_270 | 0.6667 |
| c1_sqrt | pair_271 | 0.6000 |
| c1_sqrt | pair_272 | 0.6000 |
| c1_sqrt | pair_273 | 0.6000 |
| c1_sqrt | pair_274 | 0.6000 |
| c1_sqrt | pair_275 | 0.6000 |
| c1_sqrt | pair_276 | 0.6000 |
| c1_sqrt | pair_277 | 0.6000 |
| c1_sqrt | pair_278 | 0.6000 |
| c1_sqrt | pair_279 | 0.6000 |
| c1_sqrt | pair_280 | 0.6000 |
| c1_sqrt | pair_281 | 0.6667 |
| c1_sqrt | pair_282 | 0.6000 |
| c1_sqrt | pair_283 | 0.6000 |
| c1_sqrt | pair_284 | 0.6000 |
| c1_sqrt | pair_285 | 0.6000 |
| c1_sqrt | pair_286 | 0.6667 |
| c1_sqrt | pair_287 | 0.6667 |
| c1_sqrt | pair_288 | 0.6000 |
| c1_sqrt | pair_289 | 0.6000 |
| c1_sqrt | pair_290 | 0.6000 |
| c1_sqrt | pair_291 | 0.6000 |
| c1_sqrt | pair_292 | 0.6000 |
| c1_sqrt | pair_293 | 0.6000 |
| c1_sqrt | pair_294 | 0.6000 |
| c1_sqrt | pair_295 | 0.6000 |
| c1_sqrt | pair_296 | 0.6000 |
| c1_sqrt | pair_297 | 0.6000 |
| c1_sqrt | pair_298 | 0.6667 |
| c1_sqrt | pair_299 | 0.6000 |
| c1_sqrt | pair_300 | 0.5000 |
| c1_sqrt | pair_301 | 0.5000 |
| c1_sqrt | pair_302 | 0.6667 |
| c1_sqrt | pair_303 | 0.6667 |
| c1_sqrt | pair_304 | 0.5000 |
| c1_sqrt | pair_305 | 0.5000 |
| c1_sqrt | pair_306 | 0.5000 |
| c1_sqrt | pair_307 | 0.5000 |
| c1_sqrt | pair_308 | 0.5000 |
| c1_sqrt | pair_309 | 0.6000 |
| c1_sqrt | pair_310 | 0.5000 |
| c1_sqrt | pair_311 | 0.5000 |
| c1_sqrt | pair_312 | 0.6000 |
| c1_sqrt | pair_313 | 0.6000 |
| c1_sqrt | pair_314 | 0.6667 |
| c1_sqrt | pair_315 | 0.5000 |
| c1_sqrt | pair_316 | 0.5000 |
| c1_sqrt | pair_317 | 0.6667 |
| c1_sqrt | pair_318 | 0.6667 |
| c1_sqrt | pair_319 | 0.5000 |
| c1_sqrt | pair_320 | 0.5000 |
| c1_sqrt | pair_321 | 0.5000 |
| c1_sqrt | pair_322 | 0.5000 |
| c1_sqrt | pair_323 | 0.5000 |
| c1_sqrt | pair_324 | 0.6000 |
| c1_sqrt | pair_325 | 0.5000 |
| c1_sqrt | pair_326 | 0.5000 |
| c1_sqrt | pair_327 | 0.6000 |
| c1_sqrt | pair_328 | 0.6000 |
| c1_sqrt | pair_329 | 0.6667 |
| c1_sqrt | pair_330 | 0.5000 |
| c1_sqrt | pair_331 | 0.6667 |
| c1_sqrt | pair_332 | 0.6667 |
| c1_sqrt | pair_333 | 0.5000 |
| c1_sqrt | pair_334 | 0.5000 |
| c1_sqrt | pair_335 | 0.5000 |
| c1_sqrt | pair_336 | 0.5000 |
| c1_sqrt | pair_337 | 0.5000 |
| c1_sqrt | pair_338 | 0.6000 |
| c1_sqrt | pair_339 | 0.5000 |
| c1_sqrt | pair_340 | 0.5000 |
| c1_sqrt | pair_341 | 0.6000 |
| c1_sqrt | pair_342 | 0.6000 |
| c1_sqrt | pair_343 | 0.6667 |
| c1_sqrt | pair_344 | 0.5000 |
| c1_sqrt | pair_345 | 0.6667 |
| c1_sqrt | pair_346 | 0.6667 |
| c1_sqrt | pair_347 | 0.6667 |
| c1_sqrt | pair_348 | 0.6667 |
| c1_sqrt | pair_349 | 0.6667 |
| c1_sqrt | pair_350 | 0.6667 |
| c1_sqrt | pair_351 | 0.6667 |
| c1_sqrt | pair_352 | 0.6667 |
| c1_sqrt | pair_353 | 0.6667 |
| c1_sqrt | pair_354 | 0.6667 |
| c1_sqrt | pair_355 | 0.6667 |
| c1_sqrt | pair_356 | 0.6667 |
| c1_sqrt | pair_357 | 0.6667 |
| c1_sqrt | pair_358 | 0.6667 |
| c1_sqrt | pair_359 | 0.6667 |
| c1_sqrt | pair_360 | 0.6667 |
| c1_sqrt | pair_361 | 0.6667 |
| c1_sqrt | pair_362 | 0.6667 |
| c1_sqrt | pair_363 | 0.6667 |
| c1_sqrt | pair_364 | 0.6667 |
| c1_sqrt | pair_365 | 0.6667 |
| c1_sqrt | pair_366 | 0.6667 |
| c1_sqrt | pair_367 | 0.6667 |
| c1_sqrt | pair_368 | 0.6667 |
| c1_sqrt | pair_369 | 0.6667 |
| c1_sqrt | pair_370 | 0.4286 |
| c1_sqrt | pair_371 | 0.3636 |
| c1_sqrt | pair_372 | 0.4286 |
| c1_sqrt | pair_373 | 0.3636 |
| c1_sqrt | pair_374 | 0.6000 |
| c1_sqrt | pair_375 | 0.3636 |
| c1_sqrt | pair_376 | 0.3636 |
| c1_sqrt | pair_377 | 0.6000 |
| c1_sqrt | pair_378 | 0.6000 |
| c1_sqrt | pair_379 | 0.6667 |
| c1_sqrt | pair_380 | 0.5000 |
| c1_sqrt | pair_381 | 0.4286 |
| c1_sqrt | pair_382 | 0.4286 |
| c1_sqrt | pair_383 | 0.4286 |
| c1_sqrt | pair_384 | 0.6000 |
| c1_sqrt | pair_385 | 0.4286 |
| c1_sqrt | pair_386 | 0.4286 |
| c1_sqrt | pair_387 | 0.6000 |
| c1_sqrt | pair_388 | 0.6000 |
| c1_sqrt | pair_389 | 0.6667 |
| c1_sqrt | pair_390 | 0.5000 |
| c1_sqrt | pair_391 | 0.4286 |
| c1_sqrt | pair_392 | 0.2500 |
| c1_sqrt | pair_393 | 0.6000 |
| c1_sqrt | pair_394 | 0.3333 |
| c1_sqrt | pair_395 | 0.2381 |
| c1_sqrt | pair_396 | 0.6000 |
| c1_sqrt | pair_397 | 0.6000 |
| c1_sqrt | pair_398 | 0.6667 |
| c1_sqrt | pair_399 | 0.5000 |
| c1_sqrt | pair_400 | 0.4286 |
| c1_sqrt | pair_401 | 0.6000 |
| c1_sqrt | pair_402 | 0.4286 |
| c1_sqrt | pair_403 | 0.4286 |
| c1_sqrt | pair_404 | 0.6000 |
| c1_sqrt | pair_405 | 0.6000 |
| c1_sqrt | pair_406 | 0.6667 |
| c1_sqrt | pair_407 | 0.5000 |
| c1_sqrt | pair_408 | 0.6000 |
| c1_sqrt | pair_409 | 0.3333 |
| c1_sqrt | pair_410 | 0.2500 |
| c1_sqrt | pair_411 | 0.6000 |
| c1_sqrt | pair_412 | 0.6000 |
| c1_sqrt | pair_413 | 0.6667 |
| c1_sqrt | pair_414 | 0.5000 |
| c1_sqrt | pair_415 | 0.6000 |
| c1_sqrt | pair_416 | 0.6000 |
| c1_sqrt | pair_417 | 0.6000 |
| c1_sqrt | pair_418 | 0.6000 |
| c1_sqrt | pair_419 | 0.6667 |
| c1_sqrt | pair_420 | 0.6000 |
| c1_sqrt | pair_421 | 0.3333 |
| c1_sqrt | pair_422 | 0.6000 |
| c1_sqrt | pair_423 | 0.6000 |
| c1_sqrt | pair_424 | 0.6667 |
| c1_sqrt | pair_425 | 0.5000 |
| c1_sqrt | pair_426 | 0.6000 |
| c1_sqrt | pair_427 | 0.6000 |
| c1_sqrt | pair_428 | 0.6667 |
| c1_sqrt | pair_429 | 0.5000 |
| c1_sqrt | pair_430 | 0.6000 |
| c1_sqrt | pair_431 | 0.6667 |
| c1_sqrt | pair_432 | 0.6000 |
| c1_sqrt | pair_433 | 0.6667 |
| c1_sqrt | pair_434 | 0.6000 |
| c1_sqrt | pair_435 | 0.6667 |
| corrected_fdxt | pair_001 | 0.0000 |
| corrected_fdxt | pair_002 | 0.0000 |
| corrected_fdxt | pair_003 | 0.0000 |
| corrected_fdxt | pair_004 | 0.0000 |
| corrected_fdxt | pair_005 | 0.0000 |
| corrected_fdxt | pair_006 | 0.0000 |
| corrected_fdxt | pair_007 | 0.0000 |
| corrected_fdxt | pair_008 | 0.0000 |
| corrected_fdxt | pair_009 | 0.0000 |
| corrected_fdxt | pair_010 | 0.0000 |
| corrected_fdxt | pair_011 | 0.0000 |
| corrected_fdxt | pair_012 | 0.0000 |
| corrected_fdxt | pair_013 | 0.0000 |
| corrected_fdxt | pair_014 | 0.0000 |
| corrected_fdxt | pair_015 | 0.0000 |
| corrected_fdxt | pair_016 | 0.0000 |
| corrected_fdxt | pair_017 | 0.0000 |
| corrected_fdxt | pair_018 | 0.0000 |
| corrected_fdxt | pair_019 | 0.0000 |
| corrected_fdxt | pair_020 | 0.0000 |
| corrected_fdxt | pair_021 | 0.0000 |
| corrected_fdxt | pair_022 | 0.0000 |
| corrected_fdxt | pair_023 | 0.0000 |
| corrected_fdxt | pair_024 | 0.0000 |
| corrected_fdxt | pair_025 | 0.0000 |
| corrected_fdxt | pair_026 | 0.0000 |
| corrected_fdxt | pair_027 | 0.0000 |
| corrected_fdxt | pair_028 | 0.0000 |
| corrected_fdxt | pair_029 | 0.0000 |
| corrected_fdxt | pair_030 | 0.0000 |
| corrected_fdxt | pair_031 | 0.0000 |
| corrected_fdxt | pair_032 | 0.0000 |
| corrected_fdxt | pair_033 | 0.0000 |
| corrected_fdxt | pair_034 | 0.0000 |
| corrected_fdxt | pair_035 | 0.0000 |
| corrected_fdxt | pair_036 | 0.0000 |
| corrected_fdxt | pair_037 | 0.0000 |
| corrected_fdxt | pair_038 | 0.0000 |
| corrected_fdxt | pair_039 | 0.0000 |
| corrected_fdxt | pair_040 | 0.0000 |
| corrected_fdxt | pair_041 | 0.0000 |
| corrected_fdxt | pair_042 | 0.0000 |
| corrected_fdxt | pair_043 | 0.0000 |
| corrected_fdxt | pair_044 | 0.0000 |
| corrected_fdxt | pair_045 | 0.0000 |
| corrected_fdxt | pair_046 | 0.0000 |
| corrected_fdxt | pair_047 | 0.0000 |
| corrected_fdxt | pair_048 | 0.0000 |
| corrected_fdxt | pair_049 | 0.0000 |
| corrected_fdxt | pair_050 | 0.0000 |
| corrected_fdxt | pair_051 | 0.0000 |
| corrected_fdxt | pair_052 | 0.0000 |
| corrected_fdxt | pair_053 | 0.0000 |
| corrected_fdxt | pair_054 | 0.0000 |
| corrected_fdxt | pair_055 | 0.0000 |
| corrected_fdxt | pair_056 | 0.0000 |
| corrected_fdxt | pair_057 | 0.0000 |
| corrected_fdxt | pair_058 | 0.0000 |
| corrected_fdxt | pair_059 | 0.0000 |
| corrected_fdxt | pair_060 | 0.0000 |
| corrected_fdxt | pair_061 | 0.0000 |
| corrected_fdxt | pair_062 | 0.0000 |
| corrected_fdxt | pair_063 | 0.0000 |
| corrected_fdxt | pair_064 | 0.0000 |
| corrected_fdxt | pair_065 | 0.0000 |
| corrected_fdxt | pair_066 | 0.0000 |
| corrected_fdxt | pair_067 | 0.0000 |
| corrected_fdxt | pair_068 | 0.0000 |
| corrected_fdxt | pair_069 | 0.0000 |
| corrected_fdxt | pair_070 | 0.0000 |
| corrected_fdxt | pair_071 | 0.0000 |
| corrected_fdxt | pair_072 | 0.0000 |
| corrected_fdxt | pair_073 | 0.0000 |
| corrected_fdxt | pair_074 | 0.0000 |
| corrected_fdxt | pair_075 | 0.0000 |
| corrected_fdxt | pair_076 | 0.0000 |
| corrected_fdxt | pair_077 | 0.0000 |
| corrected_fdxt | pair_078 | 0.0000 |
| corrected_fdxt | pair_079 | 0.0000 |
| corrected_fdxt | pair_080 | 0.0000 |
| corrected_fdxt | pair_081 | 0.0000 |
| corrected_fdxt | pair_082 | 0.0000 |
| corrected_fdxt | pair_083 | 0.0000 |
| corrected_fdxt | pair_084 | 0.0000 |
| corrected_fdxt | pair_085 | 0.0000 |
| corrected_fdxt | pair_086 | 0.0000 |
| corrected_fdxt | pair_087 | 0.0000 |
| corrected_fdxt | pair_088 | 0.0000 |
| corrected_fdxt | pair_089 | 0.0000 |
| corrected_fdxt | pair_090 | 0.0000 |
| corrected_fdxt | pair_091 | 0.0000 |
| corrected_fdxt | pair_092 | 0.0000 |
| corrected_fdxt | pair_093 | 0.0000 |
| corrected_fdxt | pair_094 | 0.0000 |
| corrected_fdxt | pair_095 | 0.0000 |
| corrected_fdxt | pair_096 | 0.0000 |
| corrected_fdxt | pair_097 | 0.0000 |
| corrected_fdxt | pair_098 | 0.0000 |
| corrected_fdxt | pair_099 | 0.0000 |
| corrected_fdxt | pair_100 | 0.0000 |
| corrected_fdxt | pair_101 | 0.0000 |
| corrected_fdxt | pair_102 | 0.0000 |
| corrected_fdxt | pair_103 | 0.0000 |
| corrected_fdxt | pair_104 | 0.0000 |
| corrected_fdxt | pair_105 | 0.0000 |
| corrected_fdxt | pair_106 | 0.0000 |
| corrected_fdxt | pair_107 | 0.0000 |
| corrected_fdxt | pair_108 | 0.0000 |
| corrected_fdxt | pair_109 | 0.0000 |
| corrected_fdxt | pair_110 | 0.0000 |
| corrected_fdxt | pair_111 | 0.0000 |
| corrected_fdxt | pair_112 | 0.0000 |
| corrected_fdxt | pair_113 | 0.0000 |
| corrected_fdxt | pair_114 | 0.0000 |
| corrected_fdxt | pair_115 | 0.0000 |
| corrected_fdxt | pair_116 | 0.0000 |
| corrected_fdxt | pair_117 | 0.0000 |
| corrected_fdxt | pair_118 | 0.0000 |
| corrected_fdxt | pair_119 | 0.0000 |
| corrected_fdxt | pair_120 | 0.0000 |
| corrected_fdxt | pair_121 | 0.0000 |
| corrected_fdxt | pair_122 | 0.0000 |
| corrected_fdxt | pair_123 | 0.0000 |
| corrected_fdxt | pair_124 | 0.0000 |
| corrected_fdxt | pair_125 | 0.0000 |
| corrected_fdxt | pair_126 | 0.0000 |
| corrected_fdxt | pair_127 | 0.0000 |
| corrected_fdxt | pair_128 | 0.0000 |
| corrected_fdxt | pair_129 | 0.0000 |
| corrected_fdxt | pair_130 | 0.0000 |
| corrected_fdxt | pair_131 | 0.0000 |
| corrected_fdxt | pair_132 | 0.0000 |
| corrected_fdxt | pair_133 | 0.0000 |
| corrected_fdxt | pair_134 | 0.0000 |
| corrected_fdxt | pair_135 | 0.0000 |
| corrected_fdxt | pair_136 | 0.0000 |
| corrected_fdxt | pair_137 | 0.0000 |
| corrected_fdxt | pair_138 | 0.0000 |
| corrected_fdxt | pair_139 | 0.0000 |
| corrected_fdxt | pair_140 | 0.0000 |
| corrected_fdxt | pair_141 | 0.0000 |
| corrected_fdxt | pair_142 | 0.0000 |
| corrected_fdxt | pair_143 | 0.0000 |
| corrected_fdxt | pair_144 | 0.0000 |
| corrected_fdxt | pair_145 | 0.0000 |
| corrected_fdxt | pair_146 | 0.0000 |
| corrected_fdxt | pair_147 | 0.0000 |
| corrected_fdxt | pair_148 | 0.0000 |
| corrected_fdxt | pair_149 | 0.0000 |
| corrected_fdxt | pair_150 | 0.0000 |
| corrected_fdxt | pair_151 | 0.0000 |
| corrected_fdxt | pair_152 | 0.0000 |
| corrected_fdxt | pair_153 | 0.0000 |
| corrected_fdxt | pair_154 | 0.0000 |
| corrected_fdxt | pair_155 | 0.0000 |
| corrected_fdxt | pair_156 | 0.0000 |
| corrected_fdxt | pair_157 | 0.0000 |
| corrected_fdxt | pair_158 | 0.0000 |
| corrected_fdxt | pair_159 | 0.0000 |
| corrected_fdxt | pair_160 | 0.0000 |
| corrected_fdxt | pair_161 | 0.0000 |
| corrected_fdxt | pair_162 | 0.0000 |
| corrected_fdxt | pair_163 | 0.0000 |
| corrected_fdxt | pair_164 | 0.0000 |
| corrected_fdxt | pair_165 | 0.0000 |
| corrected_fdxt | pair_166 | 0.0000 |
| corrected_fdxt | pair_167 | 0.0000 |
| corrected_fdxt | pair_168 | 0.0000 |
| corrected_fdxt | pair_169 | 0.0000 |
| corrected_fdxt | pair_170 | 0.0000 |
| corrected_fdxt | pair_171 | 0.0000 |
| corrected_fdxt | pair_172 | 0.0000 |
| corrected_fdxt | pair_173 | 0.0000 |
| corrected_fdxt | pair_174 | 0.0000 |
| corrected_fdxt | pair_175 | 0.0000 |
| corrected_fdxt | pair_176 | 0.0000 |
| corrected_fdxt | pair_177 | 0.0000 |
| corrected_fdxt | pair_178 | 0.0000 |
| corrected_fdxt | pair_179 | 0.0000 |
| corrected_fdxt | pair_180 | 0.0000 |
| corrected_fdxt | pair_181 | 0.0000 |
| corrected_fdxt | pair_182 | 0.0000 |
| corrected_fdxt | pair_183 | 0.0000 |
| corrected_fdxt | pair_184 | 0.0000 |
| corrected_fdxt | pair_185 | 0.0000 |
| corrected_fdxt | pair_186 | 0.0000 |
| corrected_fdxt | pair_187 | 0.0000 |
| corrected_fdxt | pair_188 | 0.0000 |
| corrected_fdxt | pair_189 | 0.0000 |
| corrected_fdxt | pair_190 | 0.0000 |
| corrected_fdxt | pair_191 | 0.0000 |
| corrected_fdxt | pair_192 | 0.0000 |
| corrected_fdxt | pair_193 | 0.0000 |
| corrected_fdxt | pair_194 | 0.0000 |
| corrected_fdxt | pair_195 | 0.0000 |
| corrected_fdxt | pair_196 | 0.0000 |
| corrected_fdxt | pair_197 | 0.0000 |
| corrected_fdxt | pair_198 | 0.0000 |
| corrected_fdxt | pair_199 | 0.0000 |
| corrected_fdxt | pair_200 | 0.0000 |
| corrected_fdxt | pair_201 | 0.0000 |
| corrected_fdxt | pair_202 | 0.0000 |
| corrected_fdxt | pair_203 | 0.0000 |
| corrected_fdxt | pair_204 | 0.0000 |
| corrected_fdxt | pair_205 | 0.0000 |
| corrected_fdxt | pair_206 | 0.0000 |
| corrected_fdxt | pair_207 | 0.0000 |
| corrected_fdxt | pair_208 | 0.0000 |
| corrected_fdxt | pair_209 | 0.0000 |
| corrected_fdxt | pair_210 | 0.0000 |
| corrected_fdxt | pair_211 | 0.0000 |
| corrected_fdxt | pair_212 | 0.0000 |
| corrected_fdxt | pair_213 | 0.0000 |
| corrected_fdxt | pair_214 | 0.0000 |
| corrected_fdxt | pair_215 | 0.0000 |
| corrected_fdxt | pair_216 | 0.0000 |
| corrected_fdxt | pair_217 | 0.0000 |
| corrected_fdxt | pair_218 | 0.0000 |
| corrected_fdxt | pair_219 | 0.0000 |
| corrected_fdxt | pair_220 | 0.0000 |
| corrected_fdxt | pair_221 | 0.0000 |
| corrected_fdxt | pair_222 | 0.0000 |
| corrected_fdxt | pair_223 | 0.0000 |
| corrected_fdxt | pair_224 | 0.0000 |
| corrected_fdxt | pair_225 | 0.0000 |
| corrected_fdxt | pair_226 | 0.0000 |
| corrected_fdxt | pair_227 | 0.0000 |
| corrected_fdxt | pair_228 | 0.0000 |
| corrected_fdxt | pair_229 | 0.0000 |
| corrected_fdxt | pair_230 | 0.0000 |
| corrected_fdxt | pair_231 | 0.0000 |
| corrected_fdxt | pair_232 | 0.0000 |
| corrected_fdxt | pair_233 | 0.0000 |
| corrected_fdxt | pair_234 | 0.0000 |
| corrected_fdxt | pair_235 | 0.0000 |
| corrected_fdxt | pair_236 | 0.0000 |
| corrected_fdxt | pair_237 | 0.0000 |
| corrected_fdxt | pair_238 | 0.0000 |
| corrected_fdxt | pair_239 | 0.0000 |
| corrected_fdxt | pair_240 | 0.0000 |
| corrected_fdxt | pair_241 | 0.0000 |
| corrected_fdxt | pair_242 | 0.0000 |
| corrected_fdxt | pair_243 | 0.0000 |
| corrected_fdxt | pair_244 | 0.0000 |
| corrected_fdxt | pair_245 | 0.0000 |
| corrected_fdxt | pair_246 | 0.0000 |
| corrected_fdxt | pair_247 | 0.0000 |
| corrected_fdxt | pair_248 | 0.0000 |
| corrected_fdxt | pair_249 | 0.0000 |
| corrected_fdxt | pair_250 | 0.0000 |
| corrected_fdxt | pair_251 | 0.0000 |
| corrected_fdxt | pair_252 | 0.0000 |
| corrected_fdxt | pair_253 | 0.0000 |
| corrected_fdxt | pair_254 | 0.0000 |
| corrected_fdxt | pair_255 | 0.0000 |
| corrected_fdxt | pair_256 | 0.0000 |
| corrected_fdxt | pair_257 | 0.0000 |
| corrected_fdxt | pair_258 | 0.0000 |
| corrected_fdxt | pair_259 | 0.0000 |
| corrected_fdxt | pair_260 | 0.0000 |
| corrected_fdxt | pair_261 | 0.0000 |
| corrected_fdxt | pair_262 | 0.0000 |
| corrected_fdxt | pair_263 | 0.0000 |
| corrected_fdxt | pair_264 | 0.0000 |
| corrected_fdxt | pair_265 | 0.0000 |
| corrected_fdxt | pair_266 | 0.0000 |
| corrected_fdxt | pair_267 | 0.0000 |
| corrected_fdxt | pair_268 | 0.0000 |
| corrected_fdxt | pair_269 | 0.0000 |
| corrected_fdxt | pair_270 | 0.0000 |
| corrected_fdxt | pair_271 | 0.0000 |
| corrected_fdxt | pair_272 | 0.0000 |
| corrected_fdxt | pair_273 | 0.0000 |
| corrected_fdxt | pair_274 | 0.0000 |
| corrected_fdxt | pair_275 | 0.0000 |
| corrected_fdxt | pair_276 | 0.0000 |
| corrected_fdxt | pair_277 | 0.0000 |
| corrected_fdxt | pair_278 | 0.0000 |
| corrected_fdxt | pair_279 | 0.0000 |
| corrected_fdxt | pair_280 | 0.0000 |
| corrected_fdxt | pair_281 | 0.0000 |
| corrected_fdxt | pair_282 | 0.0000 |
| corrected_fdxt | pair_283 | 0.0000 |
| corrected_fdxt | pair_284 | 0.0000 |
| corrected_fdxt | pair_285 | 0.0000 |
| corrected_fdxt | pair_286 | 0.0000 |
| corrected_fdxt | pair_287 | 0.0000 |
| corrected_fdxt | pair_288 | 0.0000 |
| corrected_fdxt | pair_289 | 0.0000 |
| corrected_fdxt | pair_290 | 0.0000 |
| corrected_fdxt | pair_291 | 0.0000 |
| corrected_fdxt | pair_292 | 0.0000 |
| corrected_fdxt | pair_293 | 0.0000 |
| corrected_fdxt | pair_294 | 0.0000 |
| corrected_fdxt | pair_295 | 0.0000 |
| corrected_fdxt | pair_296 | 0.0000 |
| corrected_fdxt | pair_297 | 0.0000 |
| corrected_fdxt | pair_298 | 0.0000 |
| corrected_fdxt | pair_299 | 0.0000 |
| corrected_fdxt | pair_300 | 0.0000 |
| corrected_fdxt | pair_301 | 0.0000 |
| corrected_fdxt | pair_302 | 0.0000 |
| corrected_fdxt | pair_303 | 0.0000 |
| corrected_fdxt | pair_304 | 0.0000 |
| corrected_fdxt | pair_305 | 0.0000 |
| corrected_fdxt | pair_306 | 0.0000 |
| corrected_fdxt | pair_307 | 0.0000 |
| corrected_fdxt | pair_308 | 0.0000 |
| corrected_fdxt | pair_309 | 0.0000 |
| corrected_fdxt | pair_310 | 0.0000 |
| corrected_fdxt | pair_311 | 0.0000 |
| corrected_fdxt | pair_312 | 0.0000 |
| corrected_fdxt | pair_313 | 0.0000 |
| corrected_fdxt | pair_314 | 0.0000 |
| corrected_fdxt | pair_315 | 0.0000 |
| corrected_fdxt | pair_316 | 0.0000 |
| corrected_fdxt | pair_317 | 0.0000 |
| corrected_fdxt | pair_318 | 0.0000 |
| corrected_fdxt | pair_319 | 0.0000 |
| corrected_fdxt | pair_320 | 0.0000 |
| corrected_fdxt | pair_321 | 0.0000 |
| corrected_fdxt | pair_322 | 0.0000 |
| corrected_fdxt | pair_323 | 0.0000 |
| corrected_fdxt | pair_324 | 0.0000 |
| corrected_fdxt | pair_325 | 0.0000 |
| corrected_fdxt | pair_326 | 0.0000 |
| corrected_fdxt | pair_327 | 0.0000 |
| corrected_fdxt | pair_328 | 0.0000 |
| corrected_fdxt | pair_329 | 0.0000 |
| corrected_fdxt | pair_330 | 0.0000 |
| corrected_fdxt | pair_331 | 0.0000 |
| corrected_fdxt | pair_332 | 0.0000 |
| corrected_fdxt | pair_333 | 0.0000 |
| corrected_fdxt | pair_334 | 0.0000 |
| corrected_fdxt | pair_335 | 0.0000 |
| corrected_fdxt | pair_336 | 0.0000 |
| corrected_fdxt | pair_337 | 0.0000 |
| corrected_fdxt | pair_338 | 0.0000 |
| corrected_fdxt | pair_339 | 0.0000 |
| corrected_fdxt | pair_340 | 0.0000 |
| corrected_fdxt | pair_341 | 0.0000 |
| corrected_fdxt | pair_342 | 0.0000 |
| corrected_fdxt | pair_343 | 0.0000 |
| corrected_fdxt | pair_344 | 0.0000 |
| corrected_fdxt | pair_345 | 0.0000 |
| corrected_fdxt | pair_346 | 0.0000 |
| corrected_fdxt | pair_347 | 0.0000 |
| corrected_fdxt | pair_348 | 0.0000 |
| corrected_fdxt | pair_349 | 0.0000 |
| corrected_fdxt | pair_350 | 0.0000 |
| corrected_fdxt | pair_351 | 0.0000 |
| corrected_fdxt | pair_352 | 0.0000 |
| corrected_fdxt | pair_353 | 0.0000 |
| corrected_fdxt | pair_354 | 0.0000 |
| corrected_fdxt | pair_355 | 0.0000 |
| corrected_fdxt | pair_356 | 0.0000 |
| corrected_fdxt | pair_357 | 0.0000 |
| corrected_fdxt | pair_358 | 0.0000 |
| corrected_fdxt | pair_359 | 0.0000 |
| corrected_fdxt | pair_360 | 0.0000 |
| corrected_fdxt | pair_361 | 0.0000 |
| corrected_fdxt | pair_362 | 0.0000 |
| corrected_fdxt | pair_363 | 0.0000 |
| corrected_fdxt | pair_364 | 0.0000 |
| corrected_fdxt | pair_365 | 0.0000 |
| corrected_fdxt | pair_366 | 0.0000 |
| corrected_fdxt | pair_367 | 0.0000 |
| corrected_fdxt | pair_368 | 0.0000 |
| corrected_fdxt | pair_369 | 0.0000 |
| corrected_fdxt | pair_370 | 0.0000 |
| corrected_fdxt | pair_371 | 0.0000 |
| corrected_fdxt | pair_372 | 0.0000 |
| corrected_fdxt | pair_373 | 0.0000 |
| corrected_fdxt | pair_374 | 0.0000 |
| corrected_fdxt | pair_375 | 0.0000 |
| corrected_fdxt | pair_376 | 0.0000 |
| corrected_fdxt | pair_377 | 0.0000 |
| corrected_fdxt | pair_378 | 0.0000 |
| corrected_fdxt | pair_379 | 0.0000 |
| corrected_fdxt | pair_380 | 0.0000 |
| corrected_fdxt | pair_381 | 0.0000 |
| corrected_fdxt | pair_382 | 0.0000 |
| corrected_fdxt | pair_383 | 0.0000 |
| corrected_fdxt | pair_384 | 0.0000 |
| corrected_fdxt | pair_385 | 0.0000 |
| corrected_fdxt | pair_386 | 0.0000 |
| corrected_fdxt | pair_387 | 0.0000 |
| corrected_fdxt | pair_388 | 0.0000 |
| corrected_fdxt | pair_389 | 0.0000 |
| corrected_fdxt | pair_390 | 0.0000 |
| corrected_fdxt | pair_391 | 0.0000 |
| corrected_fdxt | pair_392 | 0.0000 |
| corrected_fdxt | pair_393 | 0.0000 |
| corrected_fdxt | pair_394 | 0.0000 |
| corrected_fdxt | pair_395 | 0.0000 |
| corrected_fdxt | pair_396 | 0.0000 |
| corrected_fdxt | pair_397 | 0.0000 |
| corrected_fdxt | pair_398 | 0.0000 |
| corrected_fdxt | pair_399 | 0.0000 |
| corrected_fdxt | pair_400 | 0.0000 |
| corrected_fdxt | pair_401 | 0.0000 |
| corrected_fdxt | pair_402 | 0.0000 |
| corrected_fdxt | pair_403 | 0.0000 |
| corrected_fdxt | pair_404 | 0.0000 |
| corrected_fdxt | pair_405 | 0.0000 |
| corrected_fdxt | pair_406 | 0.0000 |
| corrected_fdxt | pair_407 | 0.0000 |
| corrected_fdxt | pair_408 | 0.0000 |
| corrected_fdxt | pair_409 | 0.0000 |
| corrected_fdxt | pair_410 | 0.0000 |
| corrected_fdxt | pair_411 | 0.0000 |
| corrected_fdxt | pair_412 | 0.0000 |
| corrected_fdxt | pair_413 | 0.0000 |
| corrected_fdxt | pair_414 | 0.0000 |
| corrected_fdxt | pair_415 | 0.0000 |
| corrected_fdxt | pair_416 | 0.0000 |
| corrected_fdxt | pair_417 | 0.0000 |
| corrected_fdxt | pair_418 | 0.0000 |
| corrected_fdxt | pair_419 | 0.0000 |
| corrected_fdxt | pair_420 | 0.0000 |
| corrected_fdxt | pair_421 | 0.0000 |
| corrected_fdxt | pair_422 | 0.0000 |
| corrected_fdxt | pair_423 | 0.0000 |
| corrected_fdxt | pair_424 | 0.0000 |
| corrected_fdxt | pair_425 | 0.0000 |
| corrected_fdxt | pair_426 | 0.0000 |
| corrected_fdxt | pair_427 | 0.0000 |
| corrected_fdxt | pair_428 | 0.0000 |
| corrected_fdxt | pair_429 | 0.0000 |
| corrected_fdxt | pair_430 | 0.0000 |
| corrected_fdxt | pair_431 | 0.0000 |
| corrected_fdxt | pair_432 | 0.0000 |
| corrected_fdxt | pair_433 | 0.0000 |
| corrected_fdxt | pair_434 | 0.0000 |
| corrected_fdxt | pair_435 | 0.0000 |

## Table 6 — Response-volume leakage analysis

The leakage model gives the observer only the encrypted response volume and asks it to infer the genuine candidate volume. Results use deterministic 5-fold cross-validation; repeated timing runs are collapsed to one point per query. Lower inference accuracy, correlation, mutual information, and normalized mutual information, and higher MAE/RMSE/collision rate indicate less direct volume information under this model. This is an experimental leakage measurement, not a formal cryptographic proof.

| Variant | Queries | Inference accuracy (95% CI) | MAE (95% CI) | RMSE (95% CI) | Pearson r | MI (bits, 95% CI) | NMI | Collision rate (95% CI) | Unique observed volumes |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| c1_hybrid | 435 | 1.0000 (0.9931–1.0000) | 0.0000 (0.0000–0.0759) | 0.0000 (0.0000–1.0515) | 0.9927 | 2.4836 (2.3543–2.5763) | 1.0000 | 0.0000 (0.0000–0.0000) | 10 |
| c1_log | 435 | 1.0000 (0.9931–1.0000) | 0.0000 (0.0000–0.0759) | 0.0000 (0.0000–1.0515) | 0.9970 | 2.4836 (2.3543–2.5763) | 1.0000 | 0.0000 (0.0000–0.0000) | 10 |
| c1_power2 | 435 | 0.5218 (0.4736–0.5701) | 0.5517 (0.4920–0.6322) | 0.8763 (0.7820–1.0483) | 0.9536 | 1.1927 (1.0775–1.2929) | 0.6930 | 0.4054 (0.3557–0.4575) | 4 |
| c1_sqrt | 435 | 1.0000 (0.9931–1.0000) | 0.0000 (0.0000–0.0759) | 0.0000 (0.0000–1.0515) | 0.9974 | 2.4836 (2.3543–2.5763) | 1.0000 | 0.0000 (0.0000–0.0000) | 10 |
| corrected_fdxt | 435 | 1.0000 (0.9931–1.0000) | 0.0000 (0.0000–0.0759) | 0.0000 (0.0000–1.0515) | 1.0000 | 2.4836 (2.3543–2.5763) | 1.0000 | 0.0000 (0.0000–0.0000) | 10 |
