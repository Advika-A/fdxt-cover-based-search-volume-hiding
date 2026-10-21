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
| c1_hybrid | pair_001 | 0.010 | 1883.309 | 14.444 | 12.796 | 27.910 | 28.417 | 2.723 |
| c1_hybrid | pair_002 | 0.010 | 1883.309 | 9.795 | 10.021 | 20.303 | 20.210 | 0.374 |
| c1_hybrid | pair_003 | 0.010 | 1883.309 | 10.415 | 10.322 | 21.199 | 21.039 | 0.384 |
| c1_hybrid | pair_004 | 0.010 | 1883.309 | 11.825 | 11.738 | 24.138 | 24.084 | 0.173 |
| c1_hybrid | pair_005 | 0.010 | 1883.309 | 6.647 | 5.487 | 12.437 | 11.356 | 1.964 |
| c1_hybrid | pair_006 | 0.010 | 1883.309 | 9.607 | 9.510 | 19.634 | 19.441 | 0.493 |
| c1_hybrid | pair_007 | 0.010 | 1883.309 | 20.010 | 19.682 | 40.824 | 40.612 | 0.522 |
| c1_hybrid | pair_008 | 0.010 | 1883.309 | 10.906 | 10.846 | 22.298 | 20.926 | 2.512 |
| c1_hybrid | pair_009 | 0.010 | 1883.309 | 27.420 | 26.552 | 55.267 | 52.417 | 6.879 |
| c1_hybrid | pair_010 | 0.010 | 1883.309 | 11.141 | 11.215 | 22.987 | 21.077 | 3.327 |
| c1_hybrid | pair_011 | 0.010 | 1883.309 | 9.326 | 9.417 | 19.380 | 18.573 | 2.060 |
| c1_hybrid | pair_012 | 0.010 | 1883.309 | 9.548 | 9.413 | 19.514 | 18.371 | 2.513 |
| c1_hybrid | pair_013 | 0.010 | 1883.309 | 10.200 | 9.715 | 20.507 | 20.231 | 0.950 |
| c1_hybrid | pair_014 | 0.010 | 1883.309 | 10.146 | 9.453 | 20.082 | 19.482 | 1.178 |
| c1_hybrid | pair_015 | 0.010 | 1883.309 | 9.555 | 9.408 | 19.407 | 19.463 | 0.139 |
| c1_hybrid | pair_016 | 0.010 | 1883.309 | 5.589 | 5.579 | 11.390 | 11.418 | 0.131 |
| c1_hybrid | pair_017 | 0.010 | 1883.309 | 5.516 | 5.438 | 11.224 | 11.258 | 0.093 |
| c1_hybrid | pair_018 | 0.010 | 1883.309 | 14.959 | 14.846 | 30.721 | 30.693 | 0.372 |
| c1_hybrid | pair_019 | 0.010 | 1883.309 | 10.243 | 10.148 | 20.789 | 20.794 | 0.183 |
| c1_hybrid | pair_020 | 0.010 | 1883.309 | 24.493 | 24.191 | 50.271 | 50.071 | 0.620 |
| c1_hybrid | pair_021 | 0.010 | 1883.309 | 10.219 | 10.123 | 21.705 | 20.822 | 1.753 |
| c1_hybrid | pair_022 | 0.010 | 1883.309 | 24.258 | 23.755 | 51.120 | 51.054 | 2.872 |
| c1_hybrid | pair_023 | 0.010 | 1883.309 | 9.814 | 9.624 | 20.011 | 19.252 | 2.365 |
| c1_hybrid | pair_024 | 0.010 | 1883.309 | 16.226 | 16.190 | 33.206 | 33.091 | 1.193 |
| c1_hybrid | pair_025 | 0.010 | 1883.309 | 25.918 | 24.329 | 51.395 | 51.377 | 1.835 |
| c1_hybrid | pair_026 | 0.010 | 1883.309 | 8.726 | 8.616 | 17.743 | 17.709 | 0.135 |
| c1_hybrid | pair_027 | 0.010 | 1883.309 | 8.620 | 8.545 | 17.535 | 17.558 | 0.080 |
| c1_hybrid | pair_028 | 0.010 | 1883.309 | 5.525 | 5.480 | 11.262 | 11.256 | 0.027 |
| c1_hybrid | pair_029 | 0.010 | 1883.309 | 9.416 | 9.379 | 19.177 | 19.179 | 0.119 |
| c1_hybrid | pair_030 | 0.010 | 1883.309 | 9.404 | 9.321 | 19.140 | 19.135 | 0.049 |
| c1_hybrid | pair_031 | 0.010 | 1883.309 | 10.248 | 10.126 | 20.788 | 20.782 | 0.020 |
| c1_hybrid | pair_032 | 0.010 | 1883.309 | 11.768 | 11.810 | 37.982 | 25.050 | 23.494 |
| c1_hybrid | pair_033 | 0.010 | 1883.309 | 6.066 | 7.880 | 14.316 | 13.015 | 3.878 |
| c1_hybrid | pair_034 | 0.010 | 1883.309 | 14.117 | 10.854 | 25.371 | 19.373 | 10.611 |
| c1_hybrid | pair_035 | 0.010 | 1883.309 | 12.961 | 11.866 | 26.578 | 27.745 | 2.033 |
| c1_hybrid | pair_036 | 0.010 | 1883.309 | 11.900 | 10.809 | 23.452 | 21.017 | 4.372 |
| c1_hybrid | pair_037 | 0.010 | 1883.309 | 12.137 | 12.462 | 25.271 | 24.347 | 1.844 |
| c1_hybrid | pair_038 | 0.010 | 1883.309 | 10.292 | 10.304 | 20.992 | 21.019 | 0.202 |
| c1_hybrid | pair_039 | 0.010 | 1883.309 | 8.684 | 8.646 | 17.745 | 17.671 | 0.138 |
| c1_hybrid | pair_040 | 0.010 | 1883.309 | 8.634 | 8.612 | 17.700 | 17.603 | 0.242 |
| c1_hybrid | pair_041 | 0.010 | 1883.309 | 9.460 | 9.856 | 19.742 | 19.310 | 0.929 |
| c1_hybrid | pair_042 | 0.010 | 1883.309 | 9.553 | 9.387 | 19.344 | 19.283 | 0.303 |
| c1_hybrid | pair_043 | 0.010 | 1883.309 | 9.403 | 9.377 | 19.208 | 19.112 | 0.257 |
| c1_hybrid | pair_044 | 0.010 | 1883.309 | 5.520 | 5.454 | 12.006 | 11.409 | 1.235 |
| c1_hybrid | pair_045 | 0.010 | 1883.309 | 6.047 | 6.065 | 12.375 | 11.639 | 1.487 |
| c1_hybrid | pair_046 | 0.010 | 1883.309 | 12.666 | 12.513 | 27.734 | 29.173 | 3.229 |
| c1_hybrid | pair_047 | 0.010 | 1883.309 | 11.191 | 10.978 | 22.870 | 22.118 | 2.701 |
| c1_hybrid | pair_048 | 0.010 | 1883.309 | 12.437 | 12.162 | 25.249 | 25.345 | 0.733 |
| c1_hybrid | pair_049 | 0.010 | 1883.309 | 10.498 | 10.312 | 21.247 | 21.135 | 0.627 |
| c1_hybrid | pair_050 | 0.010 | 1883.309 | 12.009 | 11.848 | 24.430 | 24.219 | 0.585 |
| c1_hybrid | pair_051 | 0.010 | 1883.309 | 11.103 | 8.891 | 20.386 | 18.782 | 3.771 |
| c1_hybrid | pair_052 | 0.010 | 1883.309 | 12.685 | 12.483 | 25.772 | 24.491 | 2.805 |
| c1_hybrid | pair_053 | 0.010 | 1883.309 | 11.844 | 11.826 | 24.147 | 23.981 | 0.481 |
| c1_hybrid | pair_054 | 0.010 | 1883.309 | 8.585 | 8.538 | 17.550 | 17.566 | 0.043 |
| c1_hybrid | pair_055 | 0.010 | 1883.309 | 8.637 | 8.562 | 17.739 | 17.647 | 0.252 |
| c1_hybrid | pair_056 | 0.010 | 1883.309 | 6.175 | 5.579 | 11.975 | 11.254 | 1.283 |
| c1_hybrid | pair_057 | 0.010 | 1883.309 | 10.071 | 9.742 | 20.305 | 19.186 | 2.022 |
| c1_hybrid | pair_058 | 0.010 | 1883.309 | 9.465 | 9.431 | 19.354 | 19.364 | 0.113 |
| c1_hybrid | pair_059 | 0.010 | 1883.309 | 9.503 | 9.386 | 19.282 | 19.306 | 0.134 |
| c1_hybrid | pair_060 | 0.010 | 1883.309 | 5.483 | 5.504 | 11.261 | 11.279 | 0.132 |
| c1_hybrid | pair_061 | 0.010 | 1883.309 | 9.451 | 9.391 | 19.204 | 19.188 | 0.231 |
| c1_hybrid | pair_062 | 0.010 | 1883.309 | 9.452 | 9.532 | 19.407 | 19.132 | 0.516 |
| c1_hybrid | pair_063 | 0.010 | 1883.309 | 10.083 | 9.518 | 20.018 | 19.161 | 1.716 |
| c1_hybrid | pair_064 | 0.010 | 1883.309 | 9.690 | 9.684 | 20.320 | 19.323 | 1.944 |
| c1_hybrid | pair_065 | 0.010 | 1883.309 | 11.266 | 10.314 | 22.081 | 19.479 | 4.597 |
| c1_hybrid | pair_066 | 0.010 | 1883.309 | 9.232 | 9.219 | 19.008 | 17.752 | 2.320 |
| c1_hybrid | pair_067 | 0.010 | 1883.309 | 8.740 | 8.643 | 17.691 | 17.614 | 0.321 |
| c1_hybrid | pair_068 | 0.010 | 1883.309 | 9.511 | 9.370 | 19.864 | 19.582 | 0.862 |
| c1_hybrid | pair_069 | 0.010 | 1883.309 | 9.767 | 9.382 | 19.703 | 19.709 | 0.438 |
| c1_hybrid | pair_070 | 0.010 | 1883.309 | 9.472 | 9.386 | 19.297 | 19.293 | 0.007 |
| c1_hybrid | pair_071 | 0.010 | 1883.309 | 5.655 | 5.527 | 11.444 | 11.267 | 0.391 |
| c1_hybrid | pair_072 | 0.010 | 1883.309 | 5.490 | 5.463 | 11.200 | 11.225 | 0.088 |
| c1_hybrid | pair_073 | 0.010 | 1883.309 | 9.416 | 9.663 | 19.403 | 19.343 | 0.526 |
| c1_hybrid | pair_074 | 0.010 | 1883.309 | 9.730 | 9.651 | 19.827 | 19.297 | 1.037 |
| c1_hybrid | pair_075 | 0.010 | 1883.309 | 9.425 | 9.331 | 19.236 | 19.258 | 0.059 |
| c1_hybrid | pair_076 | 0.010 | 1883.309 | 9.391 | 9.510 | 19.294 | 19.193 | 0.264 |
| c1_hybrid | pair_077 | 0.010 | 1883.309 | 9.368 | 10.463 | 20.242 | 19.378 | 1.863 |
| c1_hybrid | pair_078 | 0.010 | 1883.309 | 8.776 | 8.619 | 17.753 | 17.634 | 0.213 |
| c1_hybrid | pair_079 | 0.010 | 1883.309 | 9.504 | 9.326 | 19.216 | 19.171 | 0.121 |
| c1_hybrid | pair_080 | 0.010 | 1883.309 | 9.412 | 9.353 | 19.161 | 19.154 | 0.122 |
| c1_hybrid | pair_081 | 0.010 | 1883.309 | 8.636 | 8.610 | 17.748 | 17.792 | 0.115 |
| c1_hybrid | pair_082 | 0.010 | 1883.309 | 8.660 | 8.951 | 18.462 | 18.900 | 0.971 |
| c1_hybrid | pair_083 | 0.010 | 1883.309 | 5.537 | 6.107 | 11.888 | 12.108 | 0.658 |
| c1_hybrid | pair_084 | 0.010 | 1883.309 | 9.828 | 9.408 | 19.711 | 19.321 | 0.918 |
| c1_hybrid | pair_085 | 0.010 | 1883.309 | 10.258 | 10.228 | 20.924 | 20.997 | 0.300 |
| c1_hybrid | pair_086 | 0.010 | 1883.309 | 5.468 | 5.455 | 11.223 | 11.270 | 0.106 |
| c1_hybrid | pair_087 | 0.010 | 1883.309 | 9.467 | 9.328 | 19.180 | 19.091 | 0.181 |
| c1_hybrid | pair_088 | 0.010 | 1883.309 | 10.226 | 10.375 | 21.034 | 21.070 | 0.079 |
| c1_hybrid | pair_089 | 0.010 | 1883.309 | 10.306 | 11.149 | 25.779 | 23.533 | 6.172 |
| c1_hybrid | pair_090 | 0.010 | 1883.309 | 14.156 | 10.582 | 25.183 | 21.321 | 6.730 |
| c1_hybrid | pair_091 | 0.010 | 1883.309 | 11.023 | 11.050 | 22.717 | 21.405 | 2.588 |
| c1_hybrid | pair_092 | 0.010 | 1883.309 | 8.761 | 9.015 | 18.248 | 17.871 | 0.713 |
| c1_hybrid | pair_093 | 0.010 | 1883.309 | 9.171 | 8.901 | 18.454 | 17.917 | 0.976 |
| c1_hybrid | pair_094 | 0.010 | 1883.309 | 9.646 | 9.711 | 19.972 | 20.082 | 0.464 |
| c1_hybrid | pair_095 | 0.010 | 1883.309 | 11.477 | 10.985 | 22.869 | 19.528 | 6.042 |
| c1_hybrid | pair_096 | 0.010 | 1883.309 | 10.359 | 10.080 | 20.828 | 19.502 | 2.561 |
| c1_hybrid | pair_097 | 0.010 | 1883.309 | 6.908 | 5.852 | 13.010 | 12.707 | 1.896 |
| c1_hybrid | pair_098 | 0.010 | 1883.309 | 5.815 | 5.937 | 12.006 | 11.334 | 1.173 |
| c1_hybrid | pair_099 | 0.010 | 1883.309 | 11.367 | 11.619 | 23.488 | 20.913 | 4.486 |
| c1_hybrid | pair_100 | 0.010 | 1883.309 | 11.721 | 11.332 | 23.517 | 21.004 | 4.572 |
| c1_hybrid | pair_101 | 0.010 | 1883.309 | 10.543 | 10.735 | 21.624 | 20.929 | 1.254 |
| c1_hybrid | pair_102 | 0.010 | 1883.309 | 10.523 | 10.271 | 21.188 | 21.236 | 0.112 |
| c1_hybrid | pair_103 | 0.010 | 1883.309 | 10.275 | 10.348 | 21.080 | 21.095 | 0.056 |
| c1_hybrid | pair_104 | 0.010 | 1883.309 | 8.736 | 8.960 | 18.116 | 17.842 | 0.570 |
| c1_hybrid | pair_105 | 0.010 | 1883.309 | 10.313 | 10.241 | 21.050 | 21.041 | 0.042 |
| c1_hybrid | pair_106 | 0.010 | 1883.309 | 10.310 | 10.239 | 20.989 | 20.954 | 0.069 |
| c1_hybrid | pair_107 | 0.010 | 1883.309 | 8.741 | 8.693 | 18.802 | 18.075 | 1.591 |
| c1_hybrid | pair_108 | 0.010 | 1883.309 | 9.051 | 8.602 | 18.993 | 18.491 | 1.406 |
| c1_hybrid | pair_109 | 0.010 | 1883.309 | 5.570 | 5.599 | 11.428 | 11.467 | 0.322 |
| c1_hybrid | pair_110 | 0.010 | 1883.309 | 9.639 | 10.106 | 20.239 | 20.717 | 0.936 |
| c1_hybrid | pair_111 | 0.010 | 1883.309 | 5.691 | 5.494 | 11.459 | 11.527 | 0.220 |
| c1_hybrid | pair_112 | 0.010 | 1883.309 | 9.755 | 10.673 | 21.093 | 19.899 | 2.736 |
| c1_hybrid | pair_113 | 0.010 | 1883.309 | 11.823 | 11.815 | 24.175 | 24.187 | 0.068 |
| c1_hybrid | pair_114 | 0.010 | 1883.309 | 10.233 | 10.200 | 20.864 | 20.972 | 0.288 |
| c1_hybrid | pair_115 | 0.010 | 1883.309 | 11.865 | 11.673 | 24.028 | 23.981 | 0.205 |
| c1_hybrid | pair_116 | 0.010 | 1883.309 | 10.992 | 10.436 | 21.886 | 22.065 | 0.740 |
| c1_hybrid | pair_117 | 0.010 | 1883.309 | 8.791 | 8.603 | 17.774 | 17.704 | 0.201 |
| c1_hybrid | pair_118 | 0.010 | 1883.309 | 8.629 | 8.605 | 17.605 | 17.574 | 0.218 |
| c1_hybrid | pair_119 | 0.010 | 1883.309 | 9.620 | 9.467 | 19.485 | 19.171 | 0.590 |
| c1_hybrid | pair_120 | 0.010 | 1883.309 | 10.833 | 9.426 | 20.616 | 19.121 | 2.608 |
| c1_hybrid | pair_121 | 0.010 | 1883.309 | 9.379 | 9.540 | 19.298 | 19.343 | 0.094 |
| c1_hybrid | pair_122 | 0.010 | 1883.309 | 5.577 | 5.516 | 11.353 | 11.407 | 0.098 |
| c1_hybrid | pair_123 | 0.010 | 1883.309 | 5.556 | 5.441 | 11.181 | 11.145 | 0.145 |
| c1_hybrid | pair_124 | 0.010 | 1883.309 | 13.573 | 11.767 | 26.125 | 24.448 | 3.517 |
| c1_hybrid | pair_125 | 0.010 | 1883.309 | 10.207 | 10.207 | 20.864 | 20.991 | 0.256 |
| c1_hybrid | pair_126 | 0.010 | 1883.309 | 11.753 | 11.672 | 24.931 | 25.348 | 0.849 |
| c1_hybrid | pair_127 | 0.010 | 1883.309 | 10.687 | 10.804 | 25.484 | 20.813 | 8.096 |
| c1_hybrid | pair_128 | 0.010 | 1883.309 | 12.112 | 11.868 | 24.826 | 24.595 | 0.854 |
| c1_hybrid | pair_129 | 0.010 | 1883.309 | 8.883 | 8.605 | 17.819 | 17.694 | 0.485 |
| c1_hybrid | pair_130 | 0.010 | 1883.309 | 11.866 | 11.708 | 24.073 | 24.034 | 0.177 |
| c1_hybrid | pair_131 | 0.010 | 1883.309 | 11.785 | 11.913 | 24.234 | 24.073 | 0.384 |
| c1_hybrid | pair_132 | 0.010 | 1883.309 | 8.657 | 8.623 | 17.723 | 17.712 | 0.160 |
| c1_hybrid | pair_133 | 0.010 | 1883.309 | 8.636 | 8.565 | 17.550 | 17.573 | 0.100 |
| c1_hybrid | pair_134 | 0.010 | 1883.309 | 5.643 | 5.509 | 11.450 | 11.479 | 0.160 |
| c1_hybrid | pair_135 | 0.010 | 1883.309 | 9.634 | 9.381 | 19.497 | 19.494 | 0.140 |
| c1_hybrid | pair_136 | 0.010 | 1883.309 | 5.473 | 5.449 | 11.159 | 11.127 | 0.110 |
| c1_hybrid | pair_137 | 0.010 | 1883.309 | 5.495 | 5.435 | 11.138 | 11.133 | 0.066 |
| c1_hybrid | pair_138 | 0.010 | 1883.309 | 6.518 | 5.576 | 12.293 | 11.309 | 1.901 |
| c1_hybrid | pair_139 | 0.010 | 1883.309 | 5.502 | 5.609 | 11.380 | 11.225 | 0.312 |
| c1_hybrid | pair_140 | 0.010 | 1883.309 | 5.580 | 5.571 | 11.410 | 11.353 | 0.123 |
| c1_hybrid | pair_141 | 0.010 | 1883.309 | 5.551 | 5.446 | 11.231 | 11.198 | 0.068 |
| c1_hybrid | pair_142 | 0.010 | 1883.309 | 5.515 | 5.573 | 11.270 | 11.275 | 0.185 |
| c1_hybrid | pair_143 | 0.010 | 1883.309 | 5.548 | 5.423 | 11.196 | 11.235 | 0.081 |
| c1_hybrid | pair_144 | 0.010 | 1883.309 | 6.070 | 5.491 | 11.757 | 11.434 | 0.628 |
| c1_hybrid | pair_145 | 0.010 | 1883.309 | 5.438 | 5.527 | 11.203 | 11.216 | 0.038 |
| c1_hybrid | pair_146 | 0.010 | 1883.309 | 5.532 | 5.403 | 11.161 | 11.162 | 0.013 |
| c1_hybrid | pair_147 | 0.010 | 1883.309 | 5.500 | 5.408 | 11.127 | 11.142 | 0.085 |
| c1_hybrid | pair_148 | 0.010 | 1883.309 | 5.508 | 5.563 | 11.271 | 11.244 | 0.100 |
| c1_hybrid | pair_149 | 0.010 | 1883.309 | 5.536 | 5.443 | 12.641 | 11.241 | 2.464 |
| c1_hybrid | pair_150 | 0.010 | 1883.309 | 5.497 | 5.495 | 11.209 | 11.191 | 0.035 |
| c1_hybrid | pair_151 | 0.010 | 1883.309 | 5.501 | 5.523 | 11.262 | 11.245 | 0.060 |
| c1_hybrid | pair_152 | 0.010 | 1883.309 | 5.528 | 5.431 | 11.199 | 11.210 | 0.062 |
| c1_hybrid | pair_153 | 0.010 | 1883.309 | 5.478 | 5.539 | 11.227 | 11.190 | 0.084 |
| c1_hybrid | pair_154 | 0.010 | 1883.309 | 5.594 | 5.606 | 11.416 | 11.203 | 0.543 |
| c1_hybrid | pair_155 | 0.010 | 1883.309 | 6.146 | 5.552 | 11.967 | 11.191 | 1.378 |
| c1_hybrid | pair_156 | 0.010 | 1883.309 | 5.660 | 5.601 | 11.475 | 11.242 | 0.470 |
| c1_hybrid | pair_157 | 0.010 | 1883.309 | 5.693 | 5.838 | 11.761 | 11.349 | 0.793 |
| c1_hybrid | pair_158 | 0.010 | 1883.309 | 5.792 | 5.548 | 11.630 | 11.286 | 0.672 |
| c1_hybrid | pair_159 | 0.010 | 1883.309 | 5.733 | 5.655 | 11.597 | 11.138 | 0.803 |
| c1_hybrid | pair_160 | 0.010 | 1883.309 | 9.439 | 9.356 | 19.227 | 19.243 | 0.113 |
| c1_hybrid | pair_161 | 0.010 | 1883.309 | 10.596 | 10.313 | 21.401 | 19.611 | 3.535 |
| c1_hybrid | pair_162 | 0.010 | 1883.309 | 10.230 | 10.199 | 20.809 | 19.362 | 2.857 |
| c1_hybrid | pair_163 | 0.010 | 1883.309 | 10.386 | 10.227 | 20.976 | 19.321 | 3.073 |
| c1_hybrid | pair_164 | 0.010 | 1883.309 | 9.518 | 9.442 | 19.365 | 17.622 | 3.081 |
| c1_hybrid | pair_165 | 0.010 | 1883.309 | 8.768 | 9.383 | 18.522 | 17.803 | 1.513 |
| c1_hybrid | pair_166 | 0.010 | 1883.309 | 9.586 | 9.399 | 21.397 | 19.901 | 3.274 |
| c1_hybrid | pair_167 | 0.010 | 1883.309 | 9.482 | 9.563 | 19.448 | 19.163 | 0.679 |
| c1_hybrid | pair_168 | 0.010 | 1883.309 | 9.424 | 9.376 | 19.196 | 19.125 | 0.217 |
| c1_hybrid | pair_169 | 0.010 | 1883.309 | 5.496 | 5.461 | 11.202 | 11.184 | 0.069 |
| c1_hybrid | pair_170 | 0.010 | 1883.309 | 5.521 | 5.508 | 11.278 | 11.247 | 0.065 |
| c1_hybrid | pair_171 | 0.010 | 1883.309 | 9.650 | 10.868 | 20.935 | 21.082 | 1.721 |
| c1_hybrid | pair_172 | 0.010 | 1883.309 | 10.472 | 9.933 | 20.992 | 21.823 | 1.586 |
| c1_hybrid | pair_173 | 0.010 | 1883.309 | 10.396 | 17.680 | 28.822 | 24.726 | 12.284 |
| c1_hybrid | pair_174 | 0.010 | 1883.309 | 10.645 | 11.979 | 23.300 | 22.953 | 1.188 |
| c1_hybrid | pair_175 | 0.010 | 1883.309 | 11.205 | 10.951 | 22.717 | 24.046 | 2.362 |
| c1_hybrid | pair_176 | 0.010 | 1883.309 | 9.272 | 9.012 | 18.783 | 18.484 | 1.181 |
| c1_hybrid | pair_177 | 0.010 | 1883.309 | 9.941 | 9.734 | 20.118 | 19.396 | 1.509 |
| c1_hybrid | pair_178 | 0.010 | 1883.309 | 9.517 | 9.448 | 19.359 | 19.456 | 0.360 |
| c1_hybrid | pair_179 | 0.010 | 1883.309 | 8.746 | 8.524 | 17.637 | 17.594 | 0.118 |
| c1_hybrid | pair_180 | 0.010 | 1883.309 | 8.654 | 8.712 | 17.733 | 17.617 | 0.219 |
| c1_hybrid | pair_181 | 0.010 | 1883.309 | 5.521 | 5.485 | 11.216 | 11.223 | 0.061 |
| c1_hybrid | pair_182 | 0.010 | 1883.309 | 9.431 | 9.379 | 19.490 | 19.233 | 0.647 |
| c1_hybrid | pair_183 | 0.010 | 1883.309 | 11.115 | 11.104 | 22.830 | 20.700 | 3.703 |
| c1_hybrid | pair_184 | 0.010 | 1883.309 | 20.799 | 20.561 | 42.810 | 41.025 | 4.207 |
| c1_hybrid | pair_185 | 0.010 | 1883.309 | 10.200 | 10.305 | 20.913 | 20.794 | 0.313 |
| c1_hybrid | pair_186 | 0.010 | 1883.309 | 8.583 | 8.569 | 17.566 | 17.594 | 0.142 |
| c1_hybrid | pair_187 | 0.010 | 1883.309 | 8.639 | 8.550 | 17.615 | 17.652 | 0.088 |
| c1_hybrid | pair_188 | 0.010 | 1883.309 | 9.460 | 9.367 | 19.558 | 19.498 | 0.340 |
| c1_hybrid | pair_189 | 0.010 | 1883.309 | 10.557 | 10.589 | 21.638 | 20.633 | 2.211 |
| c1_hybrid | pair_190 | 0.010 | 1883.309 | 12.052 | 11.990 | 24.540 | 24.343 | 5.140 |
| c1_hybrid | pair_191 | 0.010 | 1883.309 | 7.762 | 6.144 | 14.249 | 14.638 | 2.762 |
| c1_hybrid | pair_192 | 0.010 | 1883.309 | 6.218 | 6.910 | 13.419 | 14.085 | 1.916 |
| c1_hybrid | pair_193 | 0.010 | 1883.309 | 19.811 | 18.024 | 38.575 | 37.265 | 8.524 |
| c1_hybrid | pair_194 | 0.010 | 1883.309 | 10.717 | 10.586 | 21.782 | 21.793 | 0.984 |
| c1_hybrid | pair_195 | 0.010 | 1883.309 | 19.959 | 19.781 | 41.125 | 41.166 | 1.164 |
| c1_hybrid | pair_196 | 0.010 | 1883.309 | 10.553 | 10.167 | 21.284 | 21.118 | 0.578 |
| c1_hybrid | pair_197 | 0.010 | 1883.309 | 20.437 | 20.352 | 41.772 | 42.402 | 1.249 |
| c1_hybrid | pair_198 | 0.010 | 1883.309 | 9.038 | 8.996 | 18.451 | 17.658 | 1.376 |
| c1_hybrid | pair_199 | 0.010 | 1883.309 | 16.060 | 15.635 | 32.444 | 32.049 | 0.754 |
| c1_hybrid | pair_200 | 0.010 | 1883.309 | 19.732 | 19.518 | 40.227 | 40.161 | 0.215 |
| c1_hybrid | pair_201 | 0.010 | 1883.309 | 8.725 | 8.532 | 17.871 | 17.774 | 0.307 |
| c1_hybrid | pair_202 | 0.010 | 1883.309 | 8.580 | 8.577 | 17.559 | 17.574 | 0.028 |
| c1_hybrid | pair_203 | 0.010 | 1883.309 | 5.559 | 5.455 | 11.275 | 11.285 | 0.084 |
| c1_hybrid | pair_204 | 0.010 | 1883.309 | 9.395 | 9.414 | 19.198 | 19.204 | 0.030 |
| c1_hybrid | pair_205 | 0.010 | 1883.309 | 10.276 | 10.120 | 20.851 | 20.897 | 0.220 |
| c1_hybrid | pair_206 | 0.010 | 1883.309 | 10.807 | 10.128 | 21.362 | 20.876 | 1.001 |
| c1_hybrid | pair_207 | 0.010 | 1883.309 | 8.670 | 8.575 | 17.626 | 17.602 | 0.053 |
| c1_hybrid | pair_208 | 0.010 | 1883.309 | 8.617 | 8.665 | 17.745 | 17.789 | 0.083 |
| c1_hybrid | pair_209 | 0.010 | 1883.309 | 10.260 | 9.577 | 20.300 | 19.462 | 1.619 |
| c1_hybrid | pair_210 | 0.010 | 1883.309 | 9.624 | 9.676 | 19.709 | 19.327 | 0.703 |
| c1_hybrid | pair_211 | 0.010 | 1883.309 | 9.397 | 9.437 | 19.308 | 19.279 | 0.254 |
| c1_hybrid | pair_212 | 0.010 | 1883.309 | 5.488 | 5.620 | 11.380 | 11.367 | 0.233 |
| c1_hybrid | pair_213 | 0.010 | 1883.309 | 5.548 | 5.646 | 11.437 | 11.419 | 0.256 |
| c1_hybrid | pair_214 | 0.010 | 1883.309 | 10.225 | 10.268 | 20.959 | 21.037 | 0.160 |
| c1_hybrid | pair_215 | 0.010 | 1883.309 | 10.195 | 10.142 | 20.795 | 20.782 | 0.056 |
| c1_hybrid | pair_216 | 0.010 | 1883.309 | 10.428 | 10.239 | 21.096 | 21.218 | 0.274 |
| c1_hybrid | pair_217 | 0.010 | 1883.309 | 10.533 | 10.355 | 21.750 | 22.146 | 0.835 |
| c1_hybrid | pair_218 | 0.010 | 1883.309 | 10.328 | 10.169 | 20.951 | 20.966 | 0.077 |
| c1_hybrid | pair_219 | 0.010 | 1883.309 | 8.691 | 8.611 | 18.213 | 17.856 | 0.857 |
| c1_hybrid | pair_220 | 0.010 | 1883.309 | 10.388 | 11.117 | 21.929 | 21.178 | 1.607 |
| c1_hybrid | pair_221 | 0.010 | 1883.309 | 10.200 | 10.173 | 20.769 | 20.744 | 0.088 |
| c1_hybrid | pair_222 | 0.010 | 1883.309 | 8.723 | 8.977 | 18.050 | 17.555 | 0.885 |
| c1_hybrid | pair_223 | 0.010 | 1883.309 | 8.722 | 8.522 | 17.823 | 17.759 | 0.332 |
| c1_hybrid | pair_224 | 0.010 | 1883.309 | 7.860 | 5.977 | 14.242 | 11.258 | 5.211 |
| c1_hybrid | pair_225 | 0.010 | 1883.309 | 10.170 | 9.983 | 20.567 | 19.288 | 2.220 |
| c1_hybrid | pair_226 | 0.010 | 1883.309 | 10.834 | 10.537 | 21.897 | 20.906 | 1.809 |
| c1_hybrid | pair_227 | 0.010 | 1883.309 | 8.745 | 8.746 | 17.833 | 17.745 | 0.417 |
| c1_hybrid | pair_228 | 0.010 | 1883.309 | 8.564 | 8.626 | 17.632 | 17.663 | 0.073 |
| c1_hybrid | pair_229 | 0.010 | 1883.309 | 9.363 | 9.325 | 19.075 | 19.086 | 0.036 |
| c1_hybrid | pair_230 | 0.010 | 1883.309 | 9.578 | 9.585 | 19.629 | 19.247 | 0.792 |
| c1_hybrid | pair_231 | 0.010 | 1883.309 | 9.901 | 10.036 | 20.451 | 19.513 | 1.849 |
| c1_hybrid | pair_232 | 0.010 | 1883.309 | 5.820 | 5.765 | 11.838 | 11.162 | 1.261 |
| c1_hybrid | pair_233 | 0.010 | 1883.309 | 5.972 | 5.591 | 11.820 | 11.181 | 1.119 |
| c1_hybrid | pair_234 | 0.010 | 1883.309 | 15.360 | 15.162 | 84.223 | 30.536 | 92.991 |
| c1_hybrid | pair_235 | 0.010 | 1883.309 | 12.177 | 10.950 | 23.604 | 20.806 | 4.952 |
| c1_hybrid | pair_236 | 0.010 | 1883.309 | 24.537 | 24.139 | 49.873 | 50.010 | 0.336 |
| c1_hybrid | pair_237 | 0.010 | 1883.309 | 10.756 | 11.141 | 23.903 | 25.490 | 2.812 |
| c1_hybrid | pair_238 | 0.010 | 1883.309 | 24.428 | 24.692 | 50.399 | 48.643 | 3.466 |
| c1_hybrid | pair_239 | 0.010 | 1883.309 | 8.672 | 8.664 | 17.752 | 17.735 | 0.067 |
| c1_hybrid | pair_240 | 0.010 | 1883.309 | 15.912 | 15.732 | 32.395 | 32.375 | 0.156 |
| c1_hybrid | pair_241 | 0.010 | 1883.309 | 24.778 | 24.488 | 50.398 | 50.664 | 0.480 |
| c1_hybrid | pair_242 | 0.010 | 1883.309 | 9.405 | 8.798 | 18.582 | 18.000 | 1.480 |
| c1_hybrid | pair_243 | 0.010 | 1883.309 | 9.117 | 8.746 | 19.055 | 17.767 | 2.444 |
| c1_hybrid | pair_244 | 0.010 | 1883.309 | 6.515 | 6.226 | 13.007 | 12.834 | 1.552 |
| c1_hybrid | pair_245 | 0.010 | 1883.309 | 10.326 | 10.531 | 21.387 | 19.287 | 3.835 |
| c1_hybrid | pair_246 | 0.010 | 1883.309 | 9.277 | 9.020 | 18.937 | 18.163 | 1.766 |
| c1_hybrid | pair_247 | 0.010 | 1883.309 | 8.648 | 9.053 | 18.312 | 17.943 | 0.949 |
| c1_hybrid | pair_248 | 0.010 | 1883.309 | 10.608 | 10.543 | 21.931 | 22.928 | 2.248 |
| c1_hybrid | pair_249 | 0.010 | 1883.309 | 9.851 | 9.750 | 19.962 | 19.999 | 0.720 |
| c1_hybrid | pair_250 | 0.010 | 1883.309 | 9.830 | 9.843 | 20.240 | 20.294 | 0.829 |
| c1_hybrid | pair_251 | 0.010 | 1883.309 | 5.900 | 5.674 | 11.864 | 11.808 | 0.713 |
| c1_hybrid | pair_252 | 0.010 | 1883.309 | 5.701 | 5.613 | 11.582 | 11.655 | 0.371 |
| c1_hybrid | pair_253 | 0.010 | 1883.309 | 10.615 | 10.640 | 21.723 | 22.143 | 0.752 |
| c1_hybrid | pair_254 | 0.010 | 1883.309 | 10.741 | 10.418 | 21.543 | 21.162 | 0.794 |
| c1_hybrid | pair_255 | 0.010 | 1883.309 | 10.485 | 10.561 | 21.505 | 20.981 | 1.027 |
| c1_hybrid | pair_256 | 0.010 | 1883.309 | 10.216 | 10.477 | 21.043 | 20.990 | 0.458 |
| c1_hybrid | pair_257 | 0.010 | 1883.309 | 10.310 | 10.374 | 21.189 | 21.229 | 0.371 |
| c1_hybrid | pair_258 | 0.010 | 1883.309 | 8.751 | 8.666 | 17.860 | 17.830 | 0.117 |
| c1_hybrid | pair_259 | 0.010 | 1883.309 | 10.327 | 10.246 | 21.046 | 21.027 | 0.140 |
| c1_hybrid | pair_260 | 0.010 | 1883.309 | 10.480 | 10.325 | 21.146 | 21.109 | 0.113 |
| c1_hybrid | pair_261 | 0.010 | 1883.309 | 9.770 | 8.664 | 18.821 | 17.911 | 1.787 |
| c1_hybrid | pair_262 | 0.010 | 1883.309 | 8.715 | 8.707 | 17.822 | 17.771 | 0.112 |
| c1_hybrid | pair_263 | 0.010 | 1883.309 | 6.143 | 5.593 | 12.021 | 11.525 | 0.980 |
| c1_hybrid | pair_264 | 0.010 | 1883.309 | 9.601 | 9.434 | 19.400 | 19.474 | 0.189 |
| c1_hybrid | pair_265 | 0.010 | 1883.309 | 8.766 | 8.629 | 17.850 | 17.836 | 0.053 |
| c1_hybrid | pair_266 | 0.010 | 1883.309 | 8.744 | 8.664 | 17.844 | 17.906 | 0.327 |
| c1_hybrid | pair_267 | 0.010 | 1883.309 | 8.702 | 8.612 | 17.656 | 17.653 | 0.188 |
| c1_hybrid | pair_268 | 0.010 | 1883.309 | 8.688 | 8.645 | 17.725 | 17.703 | 0.161 |
| c1_hybrid | pair_269 | 0.010 | 1883.309 | 6.211 | 5.587 | 12.076 | 11.399 | 1.368 |
| c1_hybrid | pair_270 | 0.010 | 1883.309 | 5.563 | 6.608 | 12.395 | 12.702 | 1.275 |
| c1_hybrid | pair_271 | 0.010 | 1883.309 | 8.861 | 8.669 | 17.924 | 17.750 | 0.459 |
| c1_hybrid | pair_272 | 0.010 | 1883.309 | 8.658 | 8.678 | 18.019 | 18.040 | 0.462 |
| c1_hybrid | pair_273 | 0.010 | 1883.309 | 12.137 | 9.068 | 21.599 | 18.329 | 6.343 |
| c1_hybrid | pair_274 | 0.010 | 1883.309 | 9.591 | 9.117 | 19.152 | 17.705 | 2.653 |
| c1_hybrid | pair_275 | 0.010 | 1883.309 | 8.805 | 8.655 | 17.833 | 17.704 | 0.336 |
| c1_hybrid | pair_276 | 0.010 | 1883.309 | 8.674 | 8.593 | 17.676 | 17.682 | 0.015 |
| c1_hybrid | pair_277 | 0.010 | 1883.309 | 8.757 | 8.598 | 17.768 | 17.720 | 0.163 |
| c1_hybrid | pair_278 | 0.010 | 1883.309 | 8.633 | 8.581 | 17.669 | 17.666 | 0.109 |
| c1_hybrid | pair_279 | 0.010 | 1883.309 | 8.681 | 8.695 | 20.788 | 17.574 | 5.607 |
| c1_hybrid | pair_280 | 0.010 | 1883.309 | 9.151 | 8.684 | 18.290 | 17.702 | 1.083 |
| c1_hybrid | pair_281 | 0.010 | 1883.309 | 5.518 | 5.579 | 11.300 | 11.223 | 0.197 |
| c1_hybrid | pair_282 | 0.010 | 1883.309 | 8.723 | 8.729 | 17.863 | 17.735 | 0.299 |
| c1_hybrid | pair_283 | 0.010 | 1883.309 | 8.689 | 8.572 | 18.641 | 17.691 | 1.799 |
| c1_hybrid | pair_284 | 0.010 | 1883.309 | 9.316 | 9.102 | 18.874 | 17.633 | 2.289 |
| c1_hybrid | pair_285 | 0.010 | 1883.309 | 8.959 | 8.985 | 18.431 | 17.655 | 1.424 |
| c1_hybrid | pair_286 | 0.010 | 1883.309 | 5.748 | 5.582 | 11.634 | 11.303 | 0.713 |
| c1_hybrid | pair_287 | 0.010 | 1883.309 | 5.608 | 5.502 | 11.486 | 11.248 | 0.445 |
| c1_hybrid | pair_288 | 0.010 | 1883.309 | 9.827 | 8.964 | 19.211 | 18.829 | 1.332 |
| c1_hybrid | pair_289 | 0.010 | 1883.309 | 8.893 | 8.860 | 18.142 | 17.764 | 0.758 |
| c1_hybrid | pair_290 | 0.010 | 1883.309 | 8.878 | 8.815 | 18.071 | 17.745 | 0.616 |
| c1_hybrid | pair_291 | 0.010 | 1883.309 | 8.690 | 8.628 | 17.966 | 17.801 | 0.444 |
| c1_hybrid | pair_292 | 0.010 | 1883.309 | 9.281 | 9.058 | 18.775 | 17.966 | 1.742 |
| c1_hybrid | pair_293 | 0.010 | 1883.309 | 8.900 | 8.678 | 17.917 | 18.042 | 0.219 |
| c1_hybrid | pair_294 | 0.010 | 1883.309 | 8.648 | 8.596 | 17.723 | 17.718 | 0.033 |
| c1_hybrid | pair_295 | 0.010 | 1883.309 | 8.717 | 8.565 | 17.708 | 17.623 | 0.300 |
| c1_hybrid | pair_296 | 0.010 | 1883.309 | 8.742 | 8.605 | 17.758 | 17.772 | 0.306 |
| c1_hybrid | pair_297 | 0.010 | 1883.309 | 8.656 | 8.660 | 17.698 | 17.537 | 0.317 |
| c1_hybrid | pair_298 | 0.010 | 1883.309 | 5.584 | 5.514 | 11.825 | 11.670 | 0.738 |
| c1_hybrid | pair_299 | 0.010 | 1883.309 | 8.760 | 8.610 | 17.758 | 17.809 | 0.354 |
| c1_hybrid | pair_300 | 0.010 | 1883.309 | 9.607 | 9.481 | 19.594 | 19.505 | 0.513 |
| c1_hybrid | pair_301 | 0.010 | 1883.309 | 9.838 | 9.756 | 20.121 | 20.019 | 0.996 |
| c1_hybrid | pair_302 | 0.010 | 1883.309 | 5.699 | 5.832 | 11.823 | 11.673 | 0.649 |
| c1_hybrid | pair_303 | 0.010 | 1883.309 | 5.908 | 5.834 | 11.932 | 11.993 | 0.734 |
| c1_hybrid | pair_304 | 0.010 | 1883.309 | 10.188 | 9.809 | 20.461 | 20.271 | 1.422 |
| c1_hybrid | pair_305 | 0.010 | 1883.309 | 10.155 | 10.279 | 20.885 | 21.486 | 1.599 |
| c1_hybrid | pair_306 | 0.010 | 1883.309 | 10.254 | 10.104 | 20.793 | 20.699 | 1.269 |
| c1_hybrid | pair_307 | 0.010 | 1883.309 | 10.003 | 9.857 | 20.314 | 20.443 | 0.922 |
| c1_hybrid | pair_308 | 0.010 | 1883.309 | 9.734 | 9.722 | 19.885 | 19.525 | 0.915 |
| c1_hybrid | pair_309 | 0.010 | 1883.309 | 8.673 | 8.549 | 17.685 | 17.771 | 0.198 |
| c1_hybrid | pair_310 | 0.010 | 1883.309 | 9.470 | 9.376 | 19.289 | 19.322 | 0.096 |
| c1_hybrid | pair_311 | 0.010 | 1883.309 | 9.427 | 9.363 | 19.125 | 19.138 | 0.046 |
| c1_hybrid | pair_312 | 0.010 | 1883.309 | 9.131 | 9.082 | 18.696 | 17.808 | 1.705 |
| c1_hybrid | pair_313 | 0.010 | 1883.309 | 9.068 | 9.173 | 18.813 | 17.772 | 1.875 |
| c1_hybrid | pair_314 | 0.010 | 1883.309 | 5.515 | 5.521 | 11.335 | 11.370 | 0.080 |
| c1_hybrid | pair_315 | 0.010 | 1883.309 | 9.518 | 9.491 | 19.398 | 19.328 | 0.171 |
| c1_hybrid | pair_316 | 0.010 | 1883.309 | 10.522 | 10.019 | 20.895 | 19.826 | 2.314 |
| c1_hybrid | pair_317 | 0.010 | 1883.309 | 5.595 | 5.559 | 11.402 | 11.500 | 0.241 |
| c1_hybrid | pair_318 | 0.010 | 1883.309 | 5.572 | 5.532 | 11.370 | 11.346 | 0.091 |
| c1_hybrid | pair_319 | 0.010 | 1883.309 | 9.574 | 9.442 | 19.427 | 19.480 | 0.153 |
| c1_hybrid | pair_320 | 0.010 | 1883.309 | 9.406 | 9.445 | 19.254 | 19.238 | 0.078 |
| c1_hybrid | pair_321 | 0.010 | 1883.309 | 9.532 | 9.396 | 19.273 | 19.346 | 0.253 |
| c1_hybrid | pair_322 | 0.010 | 1883.309 | 9.493 | 9.362 | 111.036 | 19.272 | 158.993 |
| c1_hybrid | pair_323 | 0.010 | 1883.309 | 12.948 | 10.653 | 24.065 | 19.466 | 8.193 |
| c1_hybrid | pair_324 | 0.010 | 1883.309 | 9.581 | 9.458 | 19.423 | 17.617 | 3.273 |
| c1_hybrid | pair_325 | 0.010 | 1883.309 | 10.468 | 10.201 | 21.163 | 19.593 | 3.083 |
| c1_hybrid | pair_326 | 0.010 | 1883.309 | 9.968 | 9.945 | 20.439 | 19.495 | 1.826 |
| c1_hybrid | pair_327 | 0.010 | 1883.309 | 8.847 | 8.706 | 17.927 | 17.835 | 0.292 |
| c1_hybrid | pair_328 | 0.010 | 1883.309 | 8.673 | 8.603 | 17.594 | 17.658 | 0.132 |
| c1_hybrid | pair_329 | 0.010 | 1883.309 | 5.482 | 5.539 | 11.299 | 11.345 | 0.084 |
| c1_hybrid | pair_330 | 0.010 | 1883.309 | 9.482 | 9.347 | 20.287 | 19.175 | 2.045 |
| c1_hybrid | pair_331 | 0.010 | 1883.309 | 8.582 | 5.570 | 14.347 | 11.464 | 5.231 |
| c1_hybrid | pair_332 | 0.010 | 1883.309 | 5.537 | 5.437 | 11.242 | 11.279 | 0.114 |
| c1_hybrid | pair_333 | 0.010 | 1883.309 | 9.462 | 10.068 | 19.958 | 19.287 | 1.273 |
| c1_hybrid | pair_334 | 0.010 | 1883.309 | 9.453 | 9.655 | 19.513 | 19.245 | 0.490 |
| c1_hybrid | pair_335 | 0.010 | 1883.309 | 9.488 | 9.446 | 19.361 | 19.331 | 0.119 |
| c1_hybrid | pair_336 | 0.010 | 1883.309 | 9.416 | 9.383 | 19.176 | 19.054 | 0.223 |
| c1_hybrid | pair_337 | 0.010 | 1883.309 | 9.866 | 9.747 | 22.332 | 19.881 | 4.837 |
| c1_hybrid | pair_338 | 0.010 | 1883.309 | 8.692 | 8.552 | 18.076 | 17.711 | 0.761 |
| c1_hybrid | pair_339 | 0.010 | 1883.309 | 9.642 | 9.698 | 19.802 | 19.301 | 0.875 |
| c1_hybrid | pair_340 | 0.010 | 1883.309 | 9.423 | 9.419 | 19.332 | 19.366 | 0.107 |
| c1_hybrid | pair_341 | 0.010 | 1883.309 | 8.572 | 8.597 | 17.536 | 17.550 | 0.134 |
| c1_hybrid | pair_342 | 0.010 | 1883.309 | 8.654 | 8.507 | 17.566 | 17.569 | 0.053 |
| c1_hybrid | pair_343 | 0.010 | 1883.309 | 5.511 | 5.467 | 11.178 | 11.222 | 0.093 |
| c1_hybrid | pair_344 | 0.010 | 1883.309 | 9.396 | 9.375 | 19.163 | 19.186 | 0.084 |
| c1_hybrid | pair_345 | 0.010 | 1883.309 | 5.534 | 5.451 | 11.216 | 11.219 | 0.085 |
| c1_hybrid | pair_346 | 0.010 | 1883.309 | 5.495 | 5.441 | 11.194 | 11.173 | 0.038 |
| c1_hybrid | pair_347 | 0.010 | 1883.309 | 5.466 | 5.499 | 11.210 | 11.172 | 0.085 |
| c1_hybrid | pair_348 | 0.010 | 1883.309 | 5.511 | 5.433 | 11.184 | 11.175 | 0.024 |
| c1_hybrid | pair_349 | 0.010 | 1883.309 | 5.465 | 5.463 | 11.178 | 11.150 | 0.094 |
| c1_hybrid | pair_350 | 0.010 | 1883.309 | 5.497 | 5.473 | 11.203 | 11.188 | 0.031 |
| c1_hybrid | pair_351 | 0.010 | 1883.309 | 5.487 | 5.559 | 11.281 | 11.228 | 0.264 |
| c1_hybrid | pair_352 | 0.010 | 1883.309 | 5.472 | 5.473 | 11.251 | 11.245 | 0.068 |
| c1_hybrid | pair_353 | 0.010 | 1883.309 | 5.477 | 5.447 | 11.166 | 11.174 | 0.019 |
| c1_hybrid | pair_354 | 0.010 | 1883.309 | 5.544 | 5.522 | 11.304 | 11.272 | 0.153 |
| c1_hybrid | pair_355 | 0.010 | 1883.309 | 6.186 | 5.535 | 12.005 | 11.237 | 1.367 |
| c1_hybrid | pair_356 | 0.010 | 1883.309 | 6.027 | 5.721 | 12.025 | 11.520 | 1.246 |
| c1_hybrid | pair_357 | 0.010 | 1883.309 | 5.733 | 5.506 | 11.564 | 11.280 | 0.724 |
| c1_hybrid | pair_358 | 0.010 | 1883.309 | 5.699 | 5.740 | 12.290 | 12.760 | 0.986 |
| c1_hybrid | pair_359 | 0.010 | 1883.309 | 6.511 | 6.080 | 12.876 | 11.651 | 2.511 |
| c1_hybrid | pair_360 | 0.010 | 1883.309 | 6.587 | 6.059 | 13.003 | 11.302 | 3.141 |
| c1_hybrid | pair_361 | 0.010 | 1883.309 | 6.184 | 5.529 | 11.968 | 11.326 | 1.424 |
| c1_hybrid | pair_362 | 0.010 | 1883.309 | 5.652 | 5.560 | 11.425 | 11.112 | 0.612 |
| c1_hybrid | pair_363 | 0.010 | 1883.309 | 5.524 | 5.474 | 11.244 | 11.214 | 0.087 |
| c1_hybrid | pair_364 | 0.010 | 1883.309 | 5.648 | 5.448 | 11.389 | 11.435 | 0.145 |
| c1_hybrid | pair_365 | 0.010 | 1883.309 | 6.025 | 5.963 | 12.198 | 11.175 | 1.876 |
| c1_hybrid | pair_366 | 0.010 | 1883.309 | 6.116 | 5.899 | 12.289 | 11.495 | 1.536 |
| c1_hybrid | pair_367 | 0.010 | 1883.309 | 5.546 | 5.680 | 11.463 | 11.358 | 0.426 |
| c1_hybrid | pair_368 | 0.010 | 1883.309 | 5.561 | 5.461 | 11.255 | 11.198 | 0.103 |
| c1_hybrid | pair_369 | 0.010 | 1883.309 | 5.539 | 5.434 | 11.215 | 11.257 | 0.097 |
| c1_hybrid | pair_370 | 0.010 | 1883.309 | 11.370 | 10.369 | 22.155 | 21.237 | 1.802 |
| c1_hybrid | pair_371 | 0.010 | 1883.309 | 15.033 | 14.816 | 30.533 | 30.501 | 0.087 |
| c1_hybrid | pair_372 | 0.010 | 1883.309 | 10.216 | 10.245 | 20.883 | 20.786 | 0.183 |
| c1_hybrid | pair_373 | 0.010 | 1883.309 | 14.913 | 15.734 | 31.405 | 31.417 | 1.103 |
| c1_hybrid | pair_374 | 0.010 | 1883.309 | 9.204 | 8.958 | 18.584 | 18.678 | 1.097 |
| c1_hybrid | pair_375 | 0.010 | 1883.309 | 15.363 | 15.420 | 32.603 | 32.074 | 2.569 |
| c1_hybrid | pair_376 | 0.010 | 1883.309 | 15.589 | 14.991 | 31.240 | 31.094 | 0.771 |
| c1_hybrid | pair_377 | 0.010 | 1883.309 | 8.648 | 8.597 | 17.657 | 17.630 | 0.065 |
| c1_hybrid | pair_378 | 0.010 | 1883.309 | 8.597 | 8.567 | 17.585 | 17.582 | 0.024 |
| c1_hybrid | pair_379 | 0.010 | 1883.309 | 5.489 | 5.443 | 11.121 | 11.140 | 0.069 |
| c1_hybrid | pair_380 | 0.010 | 1883.309 | 9.460 | 9.476 | 19.306 | 19.133 | 0.339 |
| c1_hybrid | pair_381 | 0.010 | 1883.309 | 10.240 | 10.097 | 20.764 | 20.822 | 0.136 |
| c1_hybrid | pair_382 | 0.010 | 1883.309 | 10.161 | 10.239 | 21.611 | 21.058 | 1.140 |
| c1_hybrid | pair_383 | 0.010 | 1883.309 | 10.791 | 10.796 | 22.003 | 22.169 | 1.196 |
| c1_hybrid | pair_384 | 0.010 | 1883.309 | 8.833 | 8.987 | 18.281 | 17.989 | 0.683 |
| c1_hybrid | pair_385 | 0.010 | 1883.309 | 10.456 | 10.317 | 21.192 | 20.900 | 0.654 |
| c1_hybrid | pair_386 | 0.010 | 1883.309 | 11.469 | 10.337 | 22.213 | 20.956 | 2.333 |
| c1_hybrid | pair_387 | 0.010 | 1883.309 | 8.662 | 8.593 | 17.657 | 17.581 | 0.148 |
| c1_hybrid | pair_388 | 0.010 | 1883.309 | 8.756 | 8.578 | 17.709 | 17.765 | 0.141 |
| c1_hybrid | pair_389 | 0.010 | 1883.309 | 5.587 | 5.474 | 11.283 | 11.049 | 0.433 |
| c1_hybrid | pair_390 | 0.010 | 1883.309 | 9.418 | 9.435 | 211.848 | 19.238 | 333.747 |
| c1_hybrid | pair_391 | 0.010 | 1883.309 | 12.688 | 10.764 | 24.002 | 20.673 | 5.776 |
| c1_hybrid | pair_392 | 0.010 | 1883.309 | 25.130 | 24.820 | 51.697 | 49.398 | 4.960 |
| c1_hybrid | pair_393 | 0.010 | 1883.309 | 8.646 | 8.745 | 17.695 | 17.602 | 0.324 |
| c1_hybrid | pair_394 | 0.010 | 1883.309 | 15.918 | 15.946 | 32.610 | 32.308 | 0.996 |
| c1_hybrid | pair_395 | 0.010 | 1883.309 | 24.915 | 25.242 | 51.375 | 51.509 | 1.819 |
| c1_hybrid | pair_396 | 0.010 | 1883.309 | 8.961 | 9.812 | 19.216 | 19.383 | 1.523 |
| c1_hybrid | pair_397 | 0.010 | 1883.309 | 8.800 | 8.862 | 18.051 | 17.649 | 0.743 |
| c1_hybrid | pair_398 | 0.010 | 1883.309 | 6.250 | 5.694 | 12.988 | 13.684 | 1.537 |
| c1_hybrid | pair_399 | 0.010 | 1883.309 | 10.285 | 10.608 | 21.266 | 19.829 | 3.066 |
| c1_hybrid | pair_400 | 0.010 | 1883.309 | 10.762 | 10.709 | 21.963 | 20.909 | 1.980 |
| c1_hybrid | pair_401 | 0.010 | 1883.309 | 9.031 | 8.926 | 18.431 | 17.674 | 1.499 |
| c1_hybrid | pair_402 | 0.010 | 1883.309 | 10.536 | 10.611 | 21.684 | 20.862 | 1.588 |
| c1_hybrid | pair_403 | 0.010 | 1883.309 | 10.368 | 10.161 | 20.970 | 20.880 | 0.309 |
| c1_hybrid | pair_404 | 0.010 | 1883.309 | 8.679 | 8.616 | 17.675 | 17.703 | 0.149 |
| c1_hybrid | pair_405 | 0.010 | 1883.309 | 8.715 | 9.450 | 18.488 | 17.701 | 1.464 |
| c1_hybrid | pair_406 | 0.010 | 1883.309 | 5.481 | 5.524 | 11.214 | 11.124 | 0.177 |
| c1_hybrid | pair_407 | 0.010 | 1883.309 | 9.550 | 9.401 | 19.285 | 19.304 | 0.085 |
| c1_hybrid | pair_408 | 0.010 | 1883.309 | 8.720 | 8.668 | 17.766 | 17.793 | 0.068 |
| c1_hybrid | pair_409 | 0.010 | 1883.309 | 15.781 | 15.766 | 32.624 | 32.428 | 0.425 |
| c1_hybrid | pair_410 | 0.010 | 1883.309 | 23.729 | 23.440 | 48.327 | 48.540 | 0.433 |
| c1_hybrid | pair_411 | 0.010 | 1883.309 | 8.730 | 8.614 | 17.680 | 17.588 | 0.184 |
| c1_hybrid | pair_412 | 0.010 | 1883.309 | 8.662 | 8.621 | 17.691 | 17.681 | 0.110 |
| c1_hybrid | pair_413 | 0.010 | 1883.309 | 5.581 | 5.514 | 11.319 | 11.234 | 0.268 |
| c1_hybrid | pair_414 | 0.010 | 1883.309 | 10.089 | 9.544 | 20.054 | 19.297 | 1.390 |
| c1_hybrid | pair_415 | 0.010 | 1883.309 | 8.627 | 8.628 | 17.888 | 17.903 | 0.312 |
| c1_hybrid | pair_416 | 0.010 | 1883.309 | 9.194 | 8.945 | 18.668 | 17.837 | 1.664 |
| c1_hybrid | pair_417 | 0.010 | 1883.309 | 8.699 | 8.624 | 17.791 | 17.809 | 0.242 |
| c1_hybrid | pair_418 | 0.010 | 1883.309 | 8.701 | 8.640 | 17.638 | 17.624 | 0.079 |
| c1_hybrid | pair_419 | 0.010 | 1883.309 | 5.575 | 5.438 | 11.225 | 11.280 | 0.097 |
| c1_hybrid | pair_420 | 0.010 | 1883.309 | 8.683 | 8.620 | 17.677 | 17.689 | 0.061 |
| c1_hybrid | pair_421 | 0.010 | 1883.309 | 16.487 | 16.278 | 33.632 | 32.640 | 1.807 |
| c1_hybrid | pair_422 | 0.010 | 1883.309 | 8.729 | 8.794 | 17.867 | 17.738 | 0.224 |
| c1_hybrid | pair_423 | 0.010 | 1883.309 | 8.711 | 8.557 | 17.677 | 17.634 | 0.132 |
| c1_hybrid | pair_424 | 0.010 | 1883.309 | 5.573 | 5.502 | 11.303 | 11.326 | 0.108 |
| c1_hybrid | pair_425 | 0.010 | 1883.309 | 10.169 | 9.854 | 20.458 | 19.260 | 2.188 |
| c1_hybrid | pair_426 | 0.010 | 1883.309 | 8.725 | 8.735 | 17.823 | 17.607 | 0.520 |
| c1_hybrid | pair_427 | 0.010 | 1883.309 | 8.649 | 8.545 | 17.896 | 17.597 | 0.537 |
| c1_hybrid | pair_428 | 0.010 | 1883.309 | 5.918 | 5.884 | 12.058 | 11.323 | 1.349 |
| c1_hybrid | pair_429 | 0.010 | 1883.309 | 9.966 | 9.598 | 20.001 | 19.251 | 1.367 |
| c1_hybrid | pair_430 | 0.010 | 1883.309 | 9.149 | 9.072 | 18.673 | 17.755 | 1.795 |
| c1_hybrid | pair_431 | 0.010 | 1883.309 | 5.541 | 5.567 | 11.409 | 11.367 | 0.256 |
| c1_hybrid | pair_432 | 0.010 | 1883.309 | 8.765 | 8.808 | 18.001 | 18.094 | 0.244 |
| c1_hybrid | pair_433 | 0.010 | 1883.309 | 5.516 | 5.469 | 11.243 | 11.294 | 0.119 |
| c1_hybrid | pair_434 | 0.010 | 1883.309 | 8.638 | 8.592 | 18.663 | 18.585 | 1.024 |
| c1_hybrid | pair_435 | 0.010 | 1883.309 | 5.854 | 5.815 | 11.936 | 11.209 | 1.263 |
| c1_log | pair_001 | 0.003 | 1878.336 | 11.100 | 10.959 | 22.384 | 21.823 | 3.465 |
| c1_log | pair_002 | 0.003 | 1878.336 | 8.018 | 7.886 | 16.316 | 15.908 | 1.918 |
| c1_log | pair_003 | 0.003 | 1878.336 | 9.597 | 9.276 | 19.138 | 19.394 | 2.211 |
| c1_log | pair_004 | 0.003 | 1878.336 | 10.939 | 10.906 | 22.349 | 22.451 | 2.910 |
| c1_log | pair_005 | 0.003 | 1878.336 | 4.324 | 4.468 | 9.019 | 9.291 | 0.867 |
| c1_log | pair_006 | 0.003 | 1878.336 | 7.242 | 7.026 | 14.506 | 14.495 | 0.100 |
| c1_log | pair_007 | 0.003 | 1878.336 | 16.753 | 16.669 | 34.069 | 34.122 | 0.178 |
| c1_log | pair_008 | 0.003 | 1878.336 | 9.449 | 9.487 | 19.215 | 18.985 | 2.609 |
| c1_log | pair_009 | 0.003 | 1878.336 | 24.986 | 24.380 | 50.461 | 46.444 | 10.567 |
| c1_log | pair_010 | 0.003 | 1878.336 | 9.546 | 10.003 | 20.019 | 18.023 | 3.907 |
| c1_log | pair_011 | 0.003 | 1878.336 | 6.404 | 6.408 | 13.040 | 12.999 | 0.267 |
| c1_log | pair_012 | 0.003 | 1878.336 | 8.612 | 7.737 | 16.899 | 14.465 | 5.324 |
| c1_log | pair_013 | 0.003 | 1878.336 | 7.878 | 7.545 | 16.641 | 16.447 | 2.186 |
| c1_log | pair_014 | 0.003 | 1878.336 | 7.692 | 7.190 | 15.769 | 15.398 | 1.576 |
| c1_log | pair_015 | 0.003 | 1878.336 | 7.265 | 7.168 | 14.797 | 14.624 | 0.469 |
| c1_log | pair_016 | 0.003 | 1878.336 | 4.059 | 4.006 | 8.292 | 8.224 | 0.278 |
| c1_log | pair_017 | 0.003 | 1878.336 | 4.844 | 3.959 | 9.287 | 9.184 | 1.082 |
| c1_log | pair_018 | 0.003 | 1878.336 | 12.087 | 11.869 | 25.326 | 24.314 | 2.015 |
| c1_log | pair_019 | 0.003 | 1878.336 | 8.677 | 10.597 | 19.939 | 19.267 | 4.169 |
| c1_log | pair_020 | 0.003 | 1878.336 | 23.667 | 24.040 | 51.920 | 55.720 | 8.616 |
| c1_log | pair_021 | 0.003 | 1878.336 | 7.953 | 7.846 | 16.080 | 16.143 | 0.122 |
| c1_log | pair_022 | 0.003 | 1878.336 | 22.400 | 20.977 | 44.155 | 40.180 | 7.223 |
| c1_log | pair_023 | 0.003 | 1878.336 | 6.345 | 6.241 | 12.818 | 12.803 | 0.032 |
| c1_log | pair_024 | 0.003 | 1878.336 | 12.785 | 12.724 | 25.992 | 26.229 | 0.452 |
| c1_log | pair_025 | 0.003 | 1878.336 | 20.933 | 20.563 | 42.984 | 42.472 | 1.454 |
| c1_log | pair_026 | 0.003 | 1878.336 | 6.377 | 8.323 | 14.943 | 12.809 | 3.882 |
| c1_log | pair_027 | 0.003 | 1878.336 | 7.329 | 6.826 | 14.443 | 12.896 | 2.776 |
| c1_log | pair_028 | 0.003 | 1878.336 | 4.010 | 3.901 | 8.177 | 8.043 | 0.275 |
| c1_log | pair_029 | 0.003 | 1878.336 | 7.150 | 7.101 | 14.498 | 14.502 | 0.189 |
| c1_log | pair_030 | 0.003 | 1878.336 | 7.084 | 7.095 | 14.400 | 14.357 | 0.129 |
| c1_log | pair_031 | 0.003 | 1878.336 | 7.914 | 7.832 | 16.003 | 15.991 | 0.163 |
| c1_log | pair_032 | 0.003 | 1878.336 | 9.452 | 9.449 | 19.248 | 19.262 | 0.044 |
| c1_log | pair_033 | 0.003 | 1878.336 | 3.942 | 3.928 | 8.023 | 7.984 | 0.092 |
| c1_log | pair_034 | 0.003 | 1878.336 | 7.069 | 7.021 | 14.349 | 14.348 | 0.081 |
| c1_log | pair_035 | 0.003 | 1878.336 | 9.539 | 11.057 | 21.127 | 19.429 | 3.232 |
| c1_log | pair_036 | 0.003 | 1878.336 | 7.955 | 7.883 | 16.358 | 16.292 | 0.419 |
| c1_log | pair_037 | 0.003 | 1878.336 | 9.854 | 9.838 | 20.100 | 19.292 | 1.573 |
| c1_log | pair_038 | 0.003 | 1878.336 | 8.540 | 8.536 | 17.556 | 16.135 | 2.809 |
| c1_log | pair_039 | 0.003 | 1878.336 | 6.315 | 6.271 | 12.872 | 12.871 | 0.142 |
| c1_log | pair_040 | 0.003 | 1878.336 | 6.400 | 6.233 | 12.858 | 12.736 | 0.226 |
| c1_log | pair_041 | 0.003 | 1878.336 | 7.315 | 7.011 | 14.579 | 14.428 | 0.390 |
| c1_log | pair_042 | 0.003 | 1878.336 | 7.069 | 7.033 | 14.358 | 14.319 | 0.071 |
| c1_log | pair_043 | 0.003 | 1878.336 | 7.352 | 6.997 | 14.676 | 14.327 | 0.626 |
| c1_log | pair_044 | 0.003 | 1878.336 | 3.975 | 3.919 | 8.226 | 8.327 | 0.189 |
| c1_log | pair_045 | 0.003 | 1878.336 | 6.606 | 5.487 | 12.350 | 8.036 | 7.540 |
| c1_log | pair_046 | 0.003 | 1878.336 | 10.401 | 10.419 | 21.235 | 19.680 | 2.986 |
| c1_log | pair_047 | 0.003 | 1878.336 | 8.195 | 8.050 | 16.481 | 16.453 | 0.661 |
| c1_log | pair_048 | 0.003 | 1878.336 | 10.327 | 10.710 | 21.642 | 19.251 | 4.141 |
| c1_log | pair_049 | 0.003 | 1878.336 | 8.731 | 8.619 | 17.520 | 16.196 | 2.444 |
| c1_log | pair_050 | 0.003 | 1878.336 | 9.829 | 9.566 | 19.715 | 19.323 | 0.737 |
| c1_log | pair_051 | 0.003 | 1878.336 | 6.859 | 6.745 | 13.908 | 13.011 | 1.755 |
| c1_log | pair_052 | 0.003 | 1878.336 | 9.604 | 9.674 | 19.651 | 19.390 | 0.458 |
| c1_log | pair_053 | 0.003 | 1878.336 | 9.930 | 10.098 | 20.468 | 20.580 | 0.885 |
| c1_log | pair_054 | 0.003 | 1878.336 | 6.546 | 6.491 | 13.206 | 13.263 | 0.338 |
| c1_log | pair_055 | 0.003 | 1878.336 | 6.489 | 6.850 | 13.518 | 12.914 | 1.195 |
| c1_log | pair_056 | 0.003 | 1878.336 | 4.249 | 4.048 | 8.458 | 8.205 | 0.489 |
| c1_log | pair_057 | 0.003 | 1878.336 | 7.487 | 7.208 | 14.946 | 14.553 | 0.787 |
| c1_log | pair_058 | 0.003 | 1878.336 | 7.200 | 10.192 | 17.708 | 14.426 | 5.690 |
| c1_log | pair_059 | 0.003 | 1878.336 | 7.294 | 7.792 | 15.299 | 14.400 | 1.616 |
| c1_log | pair_060 | 0.003 | 1878.336 | 4.465 | 4.480 | 9.114 | 8.306 | 1.578 |
| c1_log | pair_061 | 0.003 | 1878.336 | 7.935 | 7.800 | 16.065 | 14.301 | 3.107 |
| c1_log | pair_062 | 0.003 | 1878.336 | 8.093 | 7.668 | 16.006 | 14.433 | 2.779 |
| c1_log | pair_063 | 0.003 | 1878.336 | 7.754 | 7.620 | 15.673 | 14.452 | 2.211 |
| c1_log | pair_064 | 0.003 | 1878.336 | 7.552 | 7.258 | 15.075 | 15.289 | 0.635 |
| c1_log | pair_065 | 0.003 | 1878.336 | 7.398 | 7.232 | 14.855 | 14.728 | 0.614 |
| c1_log | pair_066 | 0.003 | 1878.336 | 6.397 | 6.289 | 12.984 | 12.986 | 0.146 |
| c1_log | pair_067 | 0.003 | 1878.336 | 6.281 | 6.479 | 12.943 | 12.959 | 0.312 |
| c1_log | pair_068 | 0.003 | 1878.336 | 7.167 | 7.140 | 14.847 | 14.894 | 0.463 |
| c1_log | pair_069 | 0.003 | 1878.336 | 7.645 | 7.481 | 15.366 | 15.454 | 0.976 |
| c1_log | pair_070 | 0.003 | 1878.336 | 10.621 | 8.704 | 19.617 | 17.596 | 6.550 |
| c1_log | pair_071 | 0.003 | 1878.336 | 4.393 | 4.271 | 8.847 | 8.180 | 1.275 |
| c1_log | pair_072 | 0.003 | 1878.336 | 4.280 | 4.211 | 8.624 | 8.228 | 1.020 |
| c1_log | pair_073 | 0.003 | 1878.336 | 7.578 | 7.264 | 15.082 | 14.401 | 1.360 |
| c1_log | pair_074 | 0.003 | 1878.336 | 7.293 | 7.372 | 15.028 | 15.077 | 0.778 |
| c1_log | pair_075 | 0.003 | 1878.336 | 7.560 | 9.141 | 16.928 | 14.906 | 3.805 |
| c1_log | pair_076 | 0.003 | 1878.336 | 7.183 | 7.162 | 14.560 | 14.575 | 0.260 |
| c1_log | pair_077 | 0.003 | 1878.336 | 7.154 | 7.118 | 14.484 | 14.614 | 0.331 |
| c1_log | pair_078 | 0.003 | 1878.336 | 6.359 | 6.379 | 12.939 | 12.796 | 0.352 |
| c1_log | pair_079 | 0.003 | 1878.336 | 7.147 | 7.171 | 17.416 | 14.781 | 4.929 |
| c1_log | pair_080 | 0.003 | 1878.336 | 7.378 | 6.989 | 14.652 | 14.600 | 0.264 |
| c1_log | pair_081 | 0.003 | 1878.336 | 6.639 | 6.291 | 13.527 | 13.666 | 0.534 |
| c1_log | pair_082 | 0.003 | 1878.336 | 6.306 | 6.408 | 12.885 | 12.913 | 0.147 |
| c1_log | pair_083 | 0.003 | 1878.336 | 4.120 | 3.921 | 8.184 | 8.269 | 0.260 |
| c1_log | pair_084 | 0.003 | 1878.336 | 7.253 | 7.598 | 15.168 | 14.771 | 0.703 |
| c1_log | pair_085 | 0.003 | 1878.336 | 7.895 | 7.867 | 16.019 | 15.979 | 0.086 |
| c1_log | pair_086 | 0.003 | 1878.336 | 4.609 | 3.999 | 8.750 | 8.379 | 0.684 |
| c1_log | pair_087 | 0.003 | 1878.336 | 7.132 | 6.994 | 14.279 | 14.299 | 0.037 |
| c1_log | pair_088 | 0.003 | 1878.336 | 7.870 | 7.908 | 15.977 | 16.015 | 0.244 |
| c1_log | pair_089 | 0.003 | 1878.336 | 7.988 | 7.856 | 16.105 | 16.053 | 0.258 |
| c1_log | pair_090 | 0.003 | 1878.336 | 7.887 | 7.906 | 16.041 | 15.969 | 0.170 |
| c1_log | pair_091 | 0.003 | 1878.336 | 7.954 | 7.938 | 16.152 | 16.020 | 0.234 |
| c1_log | pair_092 | 0.003 | 1878.336 | 6.357 | 6.345 | 12.919 | 12.868 | 0.175 |
| c1_log | pair_093 | 0.003 | 1878.336 | 6.346 | 6.198 | 12.747 | 12.746 | 0.110 |
| c1_log | pair_094 | 0.003 | 1878.336 | 7.106 | 7.086 | 14.412 | 14.295 | 0.211 |
| c1_log | pair_095 | 0.003 | 1878.336 | 7.077 | 7.032 | 14.732 | 14.274 | 0.834 |
| c1_log | pair_096 | 0.003 | 1878.336 | 8.729 | 8.842 | 17.830 | 14.523 | 6.014 |
| c1_log | pair_097 | 0.003 | 1878.336 | 4.301 | 4.033 | 8.466 | 8.038 | 0.812 |
| c1_log | pair_098 | 0.003 | 1878.336 | 3.999 | 4.271 | 8.388 | 7.907 | 0.878 |
| c1_log | pair_099 | 0.003 | 1878.336 | 8.736 | 8.769 | 17.772 | 15.801 | 3.427 |
| c1_log | pair_100 | 0.003 | 1878.336 | 8.586 | 8.719 | 17.580 | 16.178 | 2.771 |
| c1_log | pair_101 | 0.003 | 1878.336 | 8.238 | 8.241 | 16.845 | 16.285 | 1.444 |
| c1_log | pair_102 | 0.003 | 1878.336 | 7.899 | 7.923 | 16.097 | 16.104 | 0.177 |
| c1_log | pair_103 | 0.003 | 1878.336 | 7.966 | 7.829 | 16.079 | 16.279 | 0.371 |
| c1_log | pair_104 | 0.003 | 1878.336 | 6.334 | 6.280 | 12.893 | 12.925 | 0.123 |
| c1_log | pair_105 | 0.003 | 1878.336 | 7.903 | 7.817 | 15.954 | 15.962 | 0.163 |
| c1_log | pair_106 | 0.003 | 1878.336 | 7.942 | 7.872 | 16.100 | 16.114 | 0.321 |
| c1_log | pair_107 | 0.003 | 1878.336 | 6.477 | 6.422 | 13.166 | 12.814 | 0.789 |
| c1_log | pair_108 | 0.003 | 1878.336 | 9.422 | 6.861 | 16.521 | 12.855 | 6.432 |
| c1_log | pair_109 | 0.003 | 1878.336 | 4.419 | 4.318 | 8.915 | 8.110 | 1.426 |
| c1_log | pair_110 | 0.003 | 1878.336 | 7.575 | 7.593 | 15.643 | 14.776 | 1.890 |
| c1_log | pair_111 | 0.003 | 1878.336 | 4.150 | 4.039 | 8.516 | 8.564 | 0.097 |
| c1_log | pair_112 | 0.003 | 1878.336 | 7.266 | 7.645 | 15.367 | 15.292 | 0.280 |
| c1_log | pair_113 | 0.003 | 1878.336 | 9.787 | 9.637 | 19.791 | 19.649 | 0.418 |
| c1_log | pair_114 | 0.003 | 1878.336 | 7.845 | 7.907 | 16.026 | 16.038 | 0.211 |
| c1_log | pair_115 | 0.003 | 1878.336 | 9.497 | 9.360 | 19.132 | 19.087 | 0.129 |
| c1_log | pair_116 | 0.003 | 1878.336 | 8.002 | 7.839 | 16.114 | 16.110 | 0.079 |
| c1_log | pair_117 | 0.003 | 1878.336 | 6.278 | 6.285 | 12.793 | 12.827 | 0.073 |
| c1_log | pair_118 | 0.003 | 1878.336 | 6.403 | 6.260 | 12.825 | 12.793 | 0.155 |
| c1_log | pair_119 | 0.003 | 1878.336 | 7.138 | 7.099 | 14.509 | 14.419 | 0.337 |
| c1_log | pair_120 | 0.003 | 1878.336 | 7.108 | 7.023 | 14.365 | 14.397 | 0.119 |
| c1_log | pair_121 | 0.003 | 1878.336 | 7.179 | 7.127 | 15.035 | 14.638 | 0.697 |
| c1_log | pair_122 | 0.003 | 1878.336 | 4.088 | 4.001 | 8.207 | 8.267 | 0.224 |
| c1_log | pair_123 | 0.003 | 1878.336 | 3.996 | 3.948 | 8.099 | 8.137 | 0.097 |
| c1_log | pair_124 | 0.003 | 1878.336 | 9.470 | 9.516 | 19.298 | 19.404 | 0.193 |
| c1_log | pair_125 | 0.003 | 1878.336 | 7.922 | 7.893 | 16.080 | 16.065 | 0.223 |
| c1_log | pair_126 | 0.003 | 1878.336 | 9.508 | 9.527 | 19.476 | 19.533 | 0.204 |
| c1_log | pair_127 | 0.003 | 1878.336 | 8.133 | 9.132 | 17.497 | 16.021 | 2.704 |
| c1_log | pair_128 | 0.003 | 1878.336 | 9.620 | 9.508 | 19.430 | 19.196 | 0.427 |
| c1_log | pair_129 | 0.003 | 1878.336 | 6.404 | 6.376 | 12.997 | 13.039 | 0.182 |
| c1_log | pair_130 | 0.003 | 1878.336 | 9.456 | 9.537 | 19.333 | 19.260 | 0.378 |
| c1_log | pair_131 | 0.003 | 1878.336 | 9.915 | 9.568 | 19.796 | 19.579 | 0.390 |
| c1_log | pair_132 | 0.003 | 1878.336 | 6.373 | 6.245 | 12.836 | 12.748 | 0.257 |
| c1_log | pair_133 | 0.003 | 1878.336 | 6.333 | 6.284 | 12.841 | 12.702 | 0.242 |
| c1_log | pair_134 | 0.003 | 1878.336 | 3.996 | 3.910 | 8.036 | 8.019 | 0.186 |
| c1_log | pair_135 | 0.003 | 1878.336 | 7.098 | 7.033 | 14.389 | 14.348 | 0.178 |
| c1_log | pair_136 | 0.003 | 1878.336 | 3.953 | 3.917 | 7.993 | 7.958 | 0.102 |
| c1_log | pair_137 | 0.003 | 1878.336 | 4.001 | 3.911 | 8.023 | 8.043 | 0.110 |
| c1_log | pair_138 | 0.003 | 1878.336 | 4.218 | 4.140 | 8.498 | 8.187 | 0.783 |
| c1_log | pair_139 | 0.003 | 1878.336 | 4.323 | 4.177 | 8.641 | 8.017 | 1.123 |
| c1_log | pair_140 | 0.003 | 1878.336 | 4.316 | 4.008 | 8.459 | 7.973 | 0.912 |
| c1_log | pair_141 | 0.003 | 1878.336 | 4.344 | 4.211 | 8.764 | 7.910 | 1.491 |
| c1_log | pair_142 | 0.003 | 1878.336 | 3.982 | 3.896 | 8.034 | 7.931 | 0.220 |
| c1_log | pair_143 | 0.003 | 1878.336 | 3.995 | 3.904 | 7.996 | 7.898 | 0.196 |
| c1_log | pair_144 | 0.003 | 1878.336 | 4.297 | 3.959 | 8.375 | 8.477 | 0.496 |
| c1_log | pair_145 | 0.003 | 1878.336 | 3.935 | 3.929 | 7.981 | 7.982 | 0.040 |
| c1_log | pair_146 | 0.003 | 1878.336 | 3.972 | 3.935 | 8.009 | 8.047 | 0.110 |
| c1_log | pair_147 | 0.003 | 1878.336 | 3.980 | 3.910 | 8.036 | 7.983 | 0.109 |
| c1_log | pair_148 | 0.003 | 1878.336 | 3.977 | 3.916 | 8.021 | 8.008 | 0.067 |
| c1_log | pair_149 | 0.003 | 1878.336 | 3.987 | 3.933 | 8.052 | 8.034 | 0.095 |
| c1_log | pair_150 | 0.003 | 1878.336 | 4.186 | 3.931 | 8.275 | 8.036 | 0.458 |
| c1_log | pair_151 | 0.003 | 1878.336 | 4.018 | 3.956 | 8.133 | 8.155 | 0.061 |
| c1_log | pair_152 | 0.003 | 1878.336 | 4.014 | 4.077 | 8.226 | 8.168 | 0.213 |
| c1_log | pair_153 | 0.003 | 1878.336 | 4.204 | 4.083 | 8.431 | 8.089 | 0.687 |
| c1_log | pair_154 | 0.003 | 1878.336 | 4.067 | 4.024 | 8.284 | 8.060 | 0.453 |
| c1_log | pair_155 | 0.003 | 1878.336 | 4.079 | 3.969 | 8.210 | 8.100 | 0.241 |
| c1_log | pair_156 | 0.003 | 1878.336 | 4.237 | 3.971 | 8.294 | 8.319 | 0.271 |
| c1_log | pair_157 | 0.003 | 1878.336 | 4.064 | 4.051 | 8.283 | 8.053 | 0.414 |
| c1_log | pair_158 | 0.003 | 1878.336 | 5.647 | 4.220 | 10.023 | 8.850 | 2.883 |
| c1_log | pair_159 | 0.003 | 1878.336 | 4.461 | 5.983 | 10.678 | 9.238 | 3.647 |
| c1_log | pair_160 | 0.003 | 1878.336 | 7.446 | 7.494 | 15.219 | 15.105 | 0.886 |
| c1_log | pair_161 | 0.003 | 1878.336 | 7.399 | 7.069 | 14.677 | 14.490 | 0.371 |
| c1_log | pair_162 | 0.003 | 1878.336 | 7.032 | 7.000 | 14.251 | 14.223 | 0.085 |
| c1_log | pair_163 | 0.003 | 1878.336 | 7.064 | 6.986 | 14.249 | 14.263 | 0.027 |
| c1_log | pair_164 | 0.003 | 1878.336 | 6.326 | 6.222 | 12.779 | 12.810 | 0.056 |
| c1_log | pair_165 | 0.003 | 1878.336 | 6.274 | 6.201 | 12.679 | 12.726 | 0.084 |
| c1_log | pair_166 | 0.003 | 1878.336 | 7.072 | 7.055 | 14.353 | 14.349 | 0.009 |
| c1_log | pair_167 | 0.003 | 1878.336 | 7.106 | 7.005 | 14.317 | 14.291 | 0.059 |
| c1_log | pair_168 | 0.003 | 1878.336 | 7.066 | 7.052 | 17.476 | 14.396 | 5.495 |
| c1_log | pair_169 | 0.003 | 1878.336 | 5.099 | 4.775 | 10.034 | 8.032 | 3.523 |
| c1_log | pair_170 | 0.003 | 1878.336 | 4.407 | 4.233 | 8.760 | 8.085 | 1.185 |
| c1_log | pair_171 | 0.003 | 1878.336 | 8.154 | 7.461 | 15.907 | 14.441 | 2.763 |
| c1_log | pair_172 | 0.003 | 1878.336 | 7.116 | 7.185 | 14.542 | 14.369 | 0.478 |
| c1_log | pair_173 | 0.003 | 1878.336 | 7.090 | 7.199 | 14.546 | 14.504 | 0.198 |
| c1_log | pair_174 | 0.003 | 1878.336 | 7.046 | 7.741 | 15.157 | 14.347 | 1.482 |
| c1_log | pair_175 | 0.003 | 1878.336 | 7.063 | 7.045 | 14.317 | 14.268 | 0.088 |
| c1_log | pair_176 | 0.003 | 1878.336 | 6.262 | 6.300 | 12.830 | 12.847 | 0.086 |
| c1_log | pair_177 | 0.003 | 1878.336 | 7.571 | 7.249 | 15.044 | 14.228 | 1.424 |
| c1_log | pair_178 | 0.003 | 1878.336 | 7.378 | 7.280 | 15.100 | 14.316 | 1.421 |
| c1_log | pair_179 | 0.003 | 1878.336 | 6.583 | 6.293 | 13.176 | 12.787 | 0.712 |
| c1_log | pair_180 | 0.003 | 1878.336 | 6.413 | 6.410 | 13.020 | 12.939 | 0.236 |
| c1_log | pair_181 | 0.003 | 1878.336 | 3.997 | 4.029 | 8.264 | 7.895 | 0.643 |
| c1_log | pair_182 | 0.003 | 1878.336 | 7.044 | 7.099 | 14.437 | 14.336 | 0.215 |
| c1_log | pair_183 | 0.003 | 1878.336 | 7.857 | 7.795 | 15.868 | 15.807 | 0.169 |
| c1_log | pair_184 | 0.003 | 1878.336 | 16.528 | 16.473 | 33.627 | 33.516 | 0.234 |
| c1_log | pair_185 | 0.003 | 1878.336 | 7.791 | 7.753 | 15.884 | 15.814 | 0.217 |
| c1_log | pair_186 | 0.003 | 1878.336 | 6.948 | 7.089 | 14.290 | 12.801 | 2.633 |
| c1_log | pair_187 | 0.003 | 1878.336 | 6.711 | 6.571 | 13.665 | 12.666 | 1.763 |
| c1_log | pair_188 | 0.003 | 1878.336 | 7.162 | 7.338 | 14.709 | 14.805 | 0.397 |
| c1_log | pair_189 | 0.003 | 1878.336 | 7.785 | 7.209 | 15.221 | 14.854 | 1.262 |
| c1_log | pair_190 | 0.003 | 1878.336 | 7.116 | 7.045 | 14.419 | 14.377 | 0.112 |
| c1_log | pair_191 | 0.003 | 1878.336 | 3.946 | 3.925 | 8.025 | 8.022 | 0.120 |
| c1_log | pair_192 | 0.003 | 1878.336 | 3.955 | 3.956 | 8.038 | 7.953 | 0.238 |
| c1_log | pair_193 | 0.003 | 1878.336 | 11.670 | 11.626 | 23.938 | 23.830 | 0.392 |
| c1_log | pair_194 | 0.003 | 1878.336 | 7.968 | 7.783 | 16.020 | 15.939 | 0.249 |
| c1_log | pair_195 | 0.003 | 1878.336 | 16.487 | 16.389 | 33.488 | 33.488 | 0.109 |
| c1_log | pair_196 | 0.003 | 1878.336 | 7.788 | 7.796 | 15.834 | 15.804 | 0.070 |
| c1_log | pair_197 | 0.003 | 1878.336 | 16.453 | 16.375 | 33.399 | 33.435 | 0.067 |
| c1_log | pair_198 | 0.003 | 1878.336 | 6.300 | 6.264 | 12.794 | 12.724 | 0.151 |
| c1_log | pair_199 | 0.003 | 1878.336 | 12.493 | 13.018 | 26.577 | 27.257 | 1.242 |
| c1_log | pair_200 | 0.003 | 1878.336 | 16.794 | 16.430 | 33.886 | 34.258 | 0.721 |
| c1_log | pair_201 | 0.003 | 1878.336 | 7.196 | 6.370 | 13.742 | 12.882 | 1.659 |
| c1_log | pair_202 | 0.003 | 1878.336 | 6.302 | 6.286 | 12.786 | 12.786 | 0.075 |
| c1_log | pair_203 | 0.003 | 1878.336 | 3.978 | 3.876 | 7.968 | 7.958 | 0.060 |
| c1_log | pair_204 | 0.003 | 1878.336 | 7.100 | 7.001 | 14.325 | 14.330 | 0.054 |
| c1_log | pair_205 | 0.003 | 1878.336 | 7.946 | 7.949 | 16.143 | 16.267 | 0.243 |
| c1_log | pair_206 | 0.003 | 1878.336 | 7.908 | 8.318 | 16.461 | 15.999 | 0.972 |
| c1_log | pair_207 | 0.003 | 1878.336 | 6.375 | 6.232 | 12.833 | 12.704 | 0.284 |
| c1_log | pair_208 | 0.003 | 1878.336 | 6.264 | 6.165 | 13.245 | 12.547 | 1.210 |
| c1_log | pair_209 | 0.003 | 1878.336 | 8.377 | 9.078 | 17.848 | 14.289 | 6.207 |
| c1_log | pair_210 | 0.003 | 1878.336 | 7.350 | 7.203 | 14.750 | 14.337 | 0.820 |
| c1_log | pair_211 | 0.003 | 1878.336 | 7.473 | 7.383 | 15.156 | 14.344 | 1.589 |
| c1_log | pair_212 | 0.003 | 1878.336 | 3.944 | 3.881 | 7.986 | 7.948 | 0.091 |
| c1_log | pair_213 | 0.003 | 1878.336 | 3.927 | 3.857 | 7.921 | 7.928 | 0.090 |
| c1_log | pair_214 | 0.003 | 1878.336 | 7.858 | 7.775 | 15.863 | 15.883 | 0.064 |
| c1_log | pair_215 | 0.003 | 1878.336 | 7.931 | 7.806 | 15.982 | 15.932 | 0.294 |
| c1_log | pair_216 | 0.003 | 1878.336 | 8.936 | 7.811 | 16.996 | 16.015 | 1.867 |
| c1_log | pair_217 | 0.003 | 1878.336 | 7.855 | 7.773 | 15.827 | 15.828 | 0.040 |
| c1_log | pair_218 | 0.003 | 1878.336 | 8.097 | 7.810 | 16.163 | 15.913 | 0.496 |
| c1_log | pair_219 | 0.003 | 1878.336 | 6.485 | 8.377 | 16.606 | 12.788 | 6.765 |
| c1_log | pair_220 | 0.003 | 1878.336 | 8.036 | 7.870 | 16.163 | 15.882 | 0.540 |
| c1_log | pair_221 | 0.003 | 1878.336 | 7.957 | 7.792 | 16.041 | 16.017 | 0.170 |
| c1_log | pair_222 | 0.003 | 1878.336 | 6.332 | 6.256 | 12.754 | 12.739 | 0.044 |
| c1_log | pair_223 | 0.003 | 1878.336 | 6.301 | 6.266 | 12.781 | 12.770 | 0.090 |
| c1_log | pair_224 | 0.003 | 1878.336 | 3.929 | 3.889 | 7.921 | 7.953 | 0.126 |
| c1_log | pair_225 | 0.003 | 1878.336 | 7.110 | 6.992 | 14.334 | 14.332 | 0.142 |
| c1_log | pair_226 | 0.003 | 1878.336 | 7.881 | 7.788 | 15.930 | 15.889 | 0.076 |
| c1_log | pair_227 | 0.003 | 1878.336 | 6.298 | 6.282 | 12.825 | 12.855 | 0.147 |
| c1_log | pair_228 | 0.003 | 1878.336 | 6.283 | 6.199 | 13.090 | 12.651 | 0.817 |
| c1_log | pair_229 | 0.003 | 1878.336 | 7.114 | 7.114 | 14.483 | 14.446 | 0.106 |
| c1_log | pair_230 | 0.003 | 1878.336 | 7.046 | 7.086 | 14.393 | 14.383 | 0.280 |
| c1_log | pair_231 | 0.003 | 1878.336 | 7.128 | 7.185 | 14.570 | 14.535 | 0.118 |
| c1_log | pair_232 | 0.003 | 1878.336 | 3.942 | 3.894 | 7.981 | 8.010 | 0.094 |
| c1_log | pair_233 | 0.003 | 1878.336 | 3.927 | 3.886 | 7.941 | 7.999 | 0.155 |
| c1_log | pair_234 | 0.003 | 1878.336 | 11.835 | 11.659 | 23.879 | 23.999 | 0.210 |
| c1_log | pair_235 | 0.003 | 1878.336 | 7.825 | 7.763 | 15.789 | 15.807 | 0.054 |
| c1_log | pair_236 | 0.003 | 1878.336 | 20.592 | 21.102 | 43.130 | 43.802 | 1.233 |
| c1_log | pair_237 | 0.003 | 1878.336 | 8.386 | 8.026 | 16.653 | 15.927 | 1.340 |
| c1_log | pair_238 | 0.003 | 1878.336 | 19.738 | 19.508 | 39.969 | 39.671 | 0.545 |
| c1_log | pair_239 | 0.003 | 1878.336 | 6.285 | 6.228 | 12.763 | 12.837 | 0.134 |
| c1_log | pair_240 | 0.003 | 1878.336 | 12.549 | 12.466 | 25.487 | 25.443 | 0.138 |
| c1_log | pair_241 | 0.003 | 1878.336 | 20.591 | 20.357 | 42.068 | 42.357 | 0.566 |
| c1_log | pair_242 | 0.003 | 1878.336 | 6.352 | 6.218 | 12.796 | 12.803 | 0.019 |
| c1_log | pair_243 | 0.003 | 1878.336 | 6.416 | 6.326 | 12.986 | 12.880 | 0.326 |
| c1_log | pair_244 | 0.003 | 1878.336 | 3.943 | 3.916 | 7.992 | 7.987 | 0.060 |
| c1_log | pair_245 | 0.003 | 1878.336 | 7.026 | 6.992 | 14.238 | 14.281 | 0.114 |
| c1_log | pair_246 | 0.003 | 1878.336 | 6.298 | 6.234 | 12.767 | 12.796 | 0.080 |
| c1_log | pair_247 | 0.003 | 1878.336 | 6.990 | 6.407 | 13.619 | 12.987 | 1.353 |
| c1_log | pair_248 | 0.003 | 1878.336 | 7.152 | 7.174 | 14.588 | 14.431 | 0.387 |
| c1_log | pair_249 | 0.003 | 1878.336 | 7.252 | 7.080 | 15.354 | 15.242 | 1.021 |
| c1_log | pair_250 | 0.003 | 1878.336 | 8.667 | 8.459 | 17.378 | 15.828 | 3.909 |
| c1_log | pair_251 | 0.003 | 1878.336 | 4.484 | 4.358 | 8.989 | 8.781 | 1.165 |
| c1_log | pair_252 | 0.003 | 1878.336 | 4.300 | 4.425 | 8.953 | 8.669 | 1.209 |
| c1_log | pair_253 | 0.003 | 1878.336 | 8.493 | 8.335 | 17.067 | 17.454 | 0.701 |
| c1_log | pair_254 | 0.003 | 1878.336 | 8.355 | 8.228 | 16.968 | 16.650 | 1.291 |
| c1_log | pair_255 | 0.003 | 1878.336 | 7.939 | 7.889 | 16.103 | 15.965 | 0.288 |
| c1_log | pair_256 | 0.003 | 1878.336 | 7.972 | 7.766 | 15.967 | 15.821 | 0.301 |
| c1_log | pair_257 | 0.003 | 1878.336 | 7.824 | 7.792 | 15.859 | 15.778 | 0.160 |
| c1_log | pair_258 | 0.003 | 1878.336 | 6.277 | 6.192 | 12.713 | 12.699 | 0.055 |
| c1_log | pair_259 | 0.003 | 1878.336 | 7.796 | 7.762 | 15.793 | 15.765 | 0.064 |
| c1_log | pair_260 | 0.003 | 1878.336 | 7.838 | 7.788 | 15.819 | 15.772 | 0.101 |
| c1_log | pair_261 | 0.003 | 1878.336 | 6.274 | 8.207 | 14.678 | 12.666 | 3.527 |
| c1_log | pair_262 | 0.003 | 1878.336 | 6.642 | 6.303 | 13.205 | 12.733 | 0.841 |
| c1_log | pair_263 | 0.003 | 1878.336 | 4.277 | 3.967 | 8.356 | 8.559 | 0.412 |
| c1_log | pair_264 | 0.003 | 1878.336 | 8.574 | 7.075 | 15.865 | 14.850 | 2.251 |
| c1_log | pair_265 | 0.003 | 1878.336 | 6.596 | 6.704 | 13.492 | 13.623 | 0.940 |
| c1_log | pair_266 | 0.003 | 1878.336 | 6.408 | 6.991 | 13.638 | 13.276 | 1.172 |
| c1_log | pair_267 | 0.003 | 1878.336 | 7.769 | 7.231 | 15.220 | 16.187 | 1.981 |
| c1_log | pair_268 | 0.003 | 1878.336 | 6.455 | 6.231 | 12.869 | 12.805 | 0.177 |
| c1_log | pair_269 | 0.003 | 1878.336 | 3.912 | 3.935 | 8.002 | 7.989 | 0.050 |
| c1_log | pair_270 | 0.003 | 1878.336 | 3.956 | 3.891 | 7.974 | 8.002 | 0.050 |
| c1_log | pair_271 | 0.003 | 1878.336 | 6.348 | 6.247 | 12.757 | 12.793 | 0.078 |
| c1_log | pair_272 | 0.003 | 1878.336 | 6.362 | 6.326 | 12.901 | 12.915 | 0.158 |
| c1_log | pair_273 | 0.003 | 1878.336 | 6.333 | 6.276 | 12.828 | 12.938 | 0.215 |
| c1_log | pair_274 | 0.003 | 1878.336 | 6.342 | 6.256 | 12.806 | 12.829 | 0.094 |
| c1_log | pair_275 | 0.003 | 1878.336 | 6.371 | 6.322 | 12.893 | 12.957 | 0.197 |
| c1_log | pair_276 | 0.003 | 1878.336 | 6.326 | 6.220 | 12.804 | 12.780 | 0.091 |
| c1_log | pair_277 | 0.003 | 1878.336 | 6.292 | 6.177 | 12.642 | 12.602 | 0.105 |
| c1_log | pair_278 | 0.003 | 1878.336 | 6.330 | 6.212 | 12.733 | 12.818 | 0.172 |
| c1_log | pair_279 | 0.003 | 1878.336 | 6.503 | 6.379 | 13.128 | 12.987 | 0.435 |
| c1_log | pair_280 | 0.003 | 1878.336 | 6.408 | 6.210 | 12.841 | 12.746 | 0.228 |
| c1_log | pair_281 | 0.003 | 1878.336 | 3.949 | 4.036 | 8.109 | 8.065 | 0.272 |
| c1_log | pair_282 | 0.003 | 1878.336 | 6.360 | 6.570 | 13.117 | 12.994 | 0.371 |
| c1_log | pair_283 | 0.003 | 1878.336 | 6.337 | 6.319 | 12.815 | 12.700 | 0.310 |
| c1_log | pair_284 | 0.003 | 1878.336 | 6.668 | 6.202 | 13.050 | 12.811 | 0.571 |
| c1_log | pair_285 | 0.003 | 1878.336 | 6.276 | 6.300 | 12.798 | 12.789 | 0.046 |
| c1_log | pair_286 | 0.003 | 1878.336 | 3.968 | 3.893 | 7.971 | 7.964 | 0.015 |
| c1_log | pair_287 | 0.003 | 1878.336 | 3.939 | 3.887 | 7.931 | 7.941 | 0.078 |
| c1_log | pair_288 | 0.003 | 1878.336 | 6.964 | 6.372 | 13.535 | 12.890 | 1.428 |
| c1_log | pair_289 | 0.003 | 1878.336 | 6.370 | 6.156 | 12.811 | 12.721 | 0.193 |
| c1_log | pair_290 | 0.003 | 1878.336 | 6.397 | 6.171 | 12.809 | 12.766 | 0.233 |
| c1_log | pair_291 | 0.003 | 1878.336 | 6.226 | 6.400 | 12.884 | 12.937 | 0.239 |
| c1_log | pair_292 | 0.003 | 1878.336 | 6.291 | 6.207 | 12.668 | 12.672 | 0.026 |
| c1_log | pair_293 | 0.003 | 1878.336 | 6.322 | 6.355 | 12.894 | 12.915 | 0.221 |
| c1_log | pair_294 | 0.003 | 1878.336 | 6.359 | 6.238 | 12.765 | 12.706 | 0.287 |
| c1_log | pair_295 | 0.003 | 1878.336 | 6.277 | 6.177 | 12.650 | 12.631 | 0.106 |
| c1_log | pair_296 | 0.003 | 1878.336 | 6.330 | 6.314 | 12.866 | 12.703 | 0.290 |
| c1_log | pair_297 | 0.003 | 1878.336 | 6.260 | 6.210 | 13.037 | 12.873 | 0.512 |
| c1_log | pair_298 | 0.003 | 1878.336 | 4.102 | 16.874 | 21.215 | 8.141 | 22.857 |
| c1_log | pair_299 | 0.003 | 1878.336 | 7.600 | 6.971 | 14.826 | 12.828 | 3.612 |
| c1_log | pair_300 | 0.003 | 1878.336 | 7.478 | 7.221 | 15.001 | 14.560 | 0.866 |
| c1_log | pair_301 | 0.003 | 1878.336 | 7.574 | 7.656 | 15.481 | 14.320 | 2.100 |
| c1_log | pair_302 | 0.003 | 1878.336 | 4.213 | 4.028 | 8.330 | 7.973 | 0.674 |
| c1_log | pair_303 | 0.003 | 1878.336 | 4.029 | 3.896 | 8.059 | 7.940 | 0.211 |
| c1_log | pair_304 | 0.003 | 1878.336 | 7.088 | 7.076 | 14.384 | 14.383 | 0.070 |
| c1_log | pair_305 | 0.003 | 1878.336 | 7.088 | 7.052 | 14.337 | 14.310 | 0.179 |
| c1_log | pair_306 | 0.003 | 1878.336 | 7.105 | 7.257 | 14.598 | 14.536 | 0.318 |
| c1_log | pair_307 | 0.003 | 1878.336 | 7.096 | 7.014 | 14.343 | 14.294 | 0.112 |
| c1_log | pair_308 | 0.003 | 1878.336 | 7.212 | 7.206 | 14.888 | 14.753 | 0.536 |
| c1_log | pair_309 | 0.003 | 1878.336 | 9.495 | 7.088 | 16.852 | 14.892 | 5.547 |
| c1_log | pair_310 | 0.003 | 1878.336 | 8.532 | 7.969 | 16.818 | 15.853 | 1.762 |
| c1_log | pair_311 | 0.003 | 1878.336 | 7.457 | 7.208 | 14.997 | 14.581 | 0.932 |
| c1_log | pair_312 | 0.003 | 1878.336 | 6.475 | 6.877 | 13.634 | 12.936 | 1.278 |
| c1_log | pair_313 | 0.003 | 1878.336 | 6.686 | 6.530 | 13.544 | 12.986 | 1.307 |
| c1_log | pair_314 | 0.003 | 1878.336 | 3.934 | 4.030 | 8.079 | 8.039 | 0.258 |
| c1_log | pair_315 | 0.003 | 1878.336 | 7.089 | 7.173 | 14.552 | 14.668 | 0.259 |
| c1_log | pair_316 | 0.003 | 1878.336 | 7.084 | 6.989 | 14.305 | 14.281 | 0.066 |
| c1_log | pair_317 | 0.003 | 1878.336 | 3.924 | 3.903 | 7.935 | 7.961 | 0.099 |
| c1_log | pair_318 | 0.003 | 1878.336 | 3.935 | 3.858 | 7.903 | 7.913 | 0.149 |
| c1_log | pair_319 | 0.003 | 1878.336 | 7.018 | 7.030 | 14.298 | 14.310 | 0.121 |
| c1_log | pair_320 | 0.003 | 1878.336 | 7.052 | 6.965 | 14.288 | 14.298 | 0.097 |
| c1_log | pair_321 | 0.003 | 1878.336 | 7.046 | 6.979 | 14.201 | 14.193 | 0.019 |
| c1_log | pair_322 | 0.003 | 1878.336 | 7.094 | 7.741 | 15.118 | 14.318 | 1.582 |
| c1_log | pair_323 | 0.003 | 1878.336 | 7.299 | 6.969 | 14.560 | 14.377 | 0.577 |
| c1_log | pair_324 | 0.003 | 1878.336 | 6.749 | 6.230 | 13.171 | 12.979 | 0.706 |
| c1_log | pair_325 | 0.003 | 1878.336 | 7.108 | 7.041 | 14.351 | 14.301 | 0.151 |
| c1_log | pair_326 | 0.003 | 1878.336 | 11.746 | 7.388 | 19.377 | 14.512 | 8.602 |
| c1_log | pair_327 | 0.003 | 1878.336 | 6.301 | 6.233 | 12.720 | 12.680 | 0.076 |
| c1_log | pair_328 | 0.003 | 1878.336 | 6.304 | 6.225 | 12.729 | 12.707 | 0.095 |
| c1_log | pair_329 | 0.003 | 1878.336 | 3.943 | 3.872 | 7.971 | 7.987 | 0.037 |
| c1_log | pair_330 | 0.003 | 1878.336 | 7.072 | 7.032 | 14.342 | 14.377 | 0.217 |
| c1_log | pair_331 | 0.003 | 1878.336 | 3.936 | 3.856 | 7.936 | 7.946 | 0.024 |
| c1_log | pair_332 | 0.003 | 1878.336 | 3.922 | 3.867 | 7.902 | 7.923 | 0.046 |
| c1_log | pair_333 | 0.003 | 1878.336 | 7.009 | 7.003 | 14.258 | 14.279 | 0.041 |
| c1_log | pair_334 | 0.003 | 1878.336 | 7.045 | 7.103 | 14.327 | 14.270 | 0.131 |
| c1_log | pair_335 | 0.003 | 1878.336 | 7.032 | 7.014 | 14.238 | 14.231 | 0.078 |
| c1_log | pair_336 | 0.003 | 1878.336 | 7.434 | 7.097 | 15.066 | 15.138 | 0.775 |
| c1_log | pair_337 | 0.003 | 1878.336 | 7.118 | 7.068 | 14.460 | 14.407 | 0.114 |
| c1_log | pair_338 | 0.003 | 1878.336 | 6.319 | 6.356 | 12.894 | 12.891 | 0.137 |
| c1_log | pair_339 | 0.003 | 1878.336 | 7.048 | 7.007 | 14.302 | 14.290 | 0.023 |
| c1_log | pair_340 | 0.003 | 1878.336 | 7.046 | 7.002 | 14.283 | 14.264 | 0.080 |
| c1_log | pair_341 | 0.003 | 1878.336 | 6.275 | 6.233 | 12.731 | 12.746 | 0.033 |
| c1_log | pair_342 | 0.003 | 1878.336 | 6.309 | 6.261 | 12.773 | 12.801 | 0.169 |
| c1_log | pair_343 | 0.003 | 1878.336 | 3.959 | 3.869 | 7.993 | 7.949 | 0.114 |
| c1_log | pair_344 | 0.003 | 1878.336 | 7.208 | 7.065 | 14.453 | 14.407 | 0.263 |
| c1_log | pair_345 | 0.003 | 1878.336 | 3.912 | 4.042 | 8.140 | 7.964 | 0.363 |
| c1_log | pair_346 | 0.003 | 1878.336 | 4.225 | 3.996 | 8.317 | 8.037 | 0.592 |
| c1_log | pair_347 | 0.003 | 1878.336 | 3.958 | 4.005 | 8.071 | 8.114 | 0.101 |
| c1_log | pair_348 | 0.003 | 1878.336 | 3.966 | 3.894 | 7.995 | 7.993 | 0.042 |
| c1_log | pair_349 | 0.003 | 1878.336 | 3.919 | 3.858 | 7.950 | 7.954 | 0.014 |
| c1_log | pair_350 | 0.003 | 1878.336 | 3.959 | 3.858 | 7.941 | 7.923 | 0.038 |
| c1_log | pair_351 | 0.003 | 1878.336 | 3.940 | 3.862 | 7.915 | 7.903 | 0.030 |
| c1_log | pair_352 | 0.003 | 1878.336 | 3.956 | 3.856 | 7.946 | 7.959 | 0.038 |
| c1_log | pair_353 | 0.003 | 1878.336 | 3.954 | 3.848 | 7.971 | 7.938 | 0.115 |
| c1_log | pair_354 | 0.003 | 1878.336 | 3.935 | 3.855 | 7.879 | 7.877 | 0.035 |
| c1_log | pair_355 | 0.003 | 1878.336 | 3.937 | 3.859 | 7.940 | 7.876 | 0.124 |
| c1_log | pair_356 | 0.003 | 1878.336 | 3.965 | 3.865 | 7.965 | 7.922 | 0.105 |
| c1_log | pair_357 | 0.003 | 1878.336 | 3.908 | 3.863 | 7.870 | 7.845 | 0.069 |
| c1_log | pair_358 | 0.003 | 1878.336 | 3.921 | 3.884 | 7.957 | 8.020 | 0.146 |
| c1_log | pair_359 | 0.003 | 1878.336 | 3.930 | 3.876 | 7.935 | 7.907 | 0.086 |
| c1_log | pair_360 | 0.003 | 1878.336 | 3.892 | 3.896 | 7.935 | 7.931 | 0.039 |
| c1_log | pair_361 | 0.003 | 1878.336 | 3.912 | 3.881 | 7.954 | 7.977 | 0.052 |
| c1_log | pair_362 | 0.003 | 1878.336 | 3.896 | 3.854 | 7.883 | 7.845 | 0.071 |
| c1_log | pair_363 | 0.003 | 1878.336 | 3.944 | 3.892 | 7.949 | 8.011 | 0.116 |
| c1_log | pair_364 | 0.003 | 1878.336 | 3.898 | 3.908 | 7.888 | 7.922 | 0.072 |
| c1_log | pair_365 | 0.003 | 1878.336 | 3.936 | 4.103 | 8.245 | 7.916 | 0.586 |
| c1_log | pair_366 | 0.003 | 1878.336 | 3.888 | 3.937 | 7.945 | 7.946 | 0.041 |
| c1_log | pair_367 | 0.003 | 1878.336 | 3.934 | 3.884 | 8.404 | 8.005 | 0.841 |
| c1_log | pair_368 | 0.003 | 1878.336 | 4.046 | 3.871 | 8.050 | 7.942 | 0.222 |
| c1_log | pair_369 | 0.003 | 1878.336 | 4.031 | 3.900 | 8.061 | 8.063 | 0.163 |
| c1_log | pair_370 | 0.003 | 1878.336 | 7.962 | 7.966 | 16.117 | 15.952 | 0.344 |
| c1_log | pair_371 | 0.003 | 1878.336 | 12.124 | 11.915 | 24.467 | 24.010 | 1.041 |
| c1_log | pair_372 | 0.003 | 1878.336 | 8.038 | 8.045 | 16.386 | 16.030 | 0.715 |
| c1_log | pair_373 | 0.003 | 1878.336 | 12.177 | 12.582 | 25.112 | 24.009 | 2.138 |
| c1_log | pair_374 | 0.003 | 1878.336 | 6.592 | 6.490 | 13.321 | 12.971 | 0.933 |
| c1_log | pair_375 | 0.003 | 1878.336 | 12.164 | 11.916 | 27.956 | 25.641 | 5.455 |
| c1_log | pair_376 | 0.003 | 1878.336 | 13.517 | 12.767 | 26.926 | 23.966 | 5.257 |
| c1_log | pair_377 | 0.003 | 1878.336 | 6.826 | 6.713 | 13.855 | 12.763 | 1.933 |
| c1_log | pair_378 | 0.003 | 1878.336 | 6.955 | 6.715 | 13.989 | 12.900 | 2.025 |
| c1_log | pair_379 | 0.003 | 1878.336 | 4.207 | 4.242 | 8.686 | 8.043 | 1.214 |
| c1_log | pair_380 | 0.003 | 1878.336 | 7.687 | 7.699 | 15.722 | 14.396 | 2.353 |
| c1_log | pair_381 | 0.003 | 1878.336 | 8.461 | 8.484 | 17.274 | 16.208 | 1.961 |
| c1_log | pair_382 | 0.003 | 1878.336 | 8.724 | 8.672 | 17.786 | 17.466 | 2.116 |
| c1_log | pair_383 | 0.003 | 1878.336 | 8.547 | 9.686 | 18.684 | 19.030 | 2.472 |
| c1_log | pair_384 | 0.003 | 1878.336 | 6.493 | 6.405 | 13.138 | 13.312 | 0.321 |
| c1_log | pair_385 | 0.003 | 1878.336 | 8.328 | 8.153 | 16.804 | 15.975 | 1.475 |
| c1_log | pair_386 | 0.003 | 1878.336 | 7.965 | 7.897 | 16.164 | 16.078 | 0.188 |
| c1_log | pair_387 | 0.003 | 1878.336 | 6.364 | 6.321 | 12.891 | 12.851 | 0.311 |
| c1_log | pair_388 | 0.003 | 1878.336 | 6.548 | 6.453 | 13.269 | 13.396 | 0.500 |
| c1_log | pair_389 | 0.003 | 1878.336 | 3.956 | 3.966 | 8.029 | 8.028 | 0.061 |
| c1_log | pair_390 | 0.003 | 1878.336 | 7.102 | 7.091 | 14.425 | 14.461 | 0.105 |
| c1_log | pair_391 | 0.003 | 1878.336 | 7.886 | 8.734 | 16.913 | 16.110 | 1.445 |
| c1_log | pair_392 | 0.003 | 1878.336 | 23.053 | 23.096 | 47.099 | 40.427 | 11.803 |
| c1_log | pair_393 | 0.003 | 1878.336 | 7.447 | 6.620 | 14.681 | 14.145 | 1.904 |
| c1_log | pair_394 | 0.003 | 1878.336 | 15.041 | 13.725 | 29.298 | 30.522 | 3.057 |
| c1_log | pair_395 | 0.003 | 1878.336 | 22.625 | 23.449 | 46.939 | 45.170 | 5.754 |
| c1_log | pair_396 | 0.003 | 1878.336 | 6.344 | 6.281 | 12.805 | 12.662 | 0.264 |
| c1_log | pair_397 | 0.003 | 1878.336 | 7.273 | 6.597 | 14.439 | 13.964 | 1.929 |
| c1_log | pair_398 | 0.003 | 1878.336 | 4.184 | 4.016 | 8.354 | 8.015 | 0.648 |
| c1_log | pair_399 | 0.003 | 1878.336 | 7.315 | 7.338 | 14.870 | 14.340 | 1.015 |
| c1_log | pair_400 | 0.003 | 1878.336 | 7.989 | 7.924 | 16.155 | 16.323 | 0.313 |
| c1_log | pair_401 | 0.003 | 1878.336 | 6.479 | 6.386 | 13.084 | 13.145 | 0.366 |
| c1_log | pair_402 | 0.003 | 1878.336 | 7.967 | 7.829 | 16.026 | 15.932 | 0.251 |
| c1_log | pair_403 | 0.003 | 1878.336 | 7.899 | 7.849 | 15.947 | 15.820 | 0.366 |
| c1_log | pair_404 | 0.003 | 1878.336 | 6.406 | 6.325 | 13.294 | 13.095 | 0.466 |
| c1_log | pair_405 | 0.003 | 1878.336 | 9.364 | 6.515 | 16.131 | 12.814 | 5.880 |
| c1_log | pair_406 | 0.003 | 1878.336 | 4.282 | 4.221 | 8.632 | 8.075 | 1.132 |
| c1_log | pair_407 | 0.003 | 1878.336 | 7.862 | 7.772 | 15.939 | 14.525 | 2.652 |
| c1_log | pair_408 | 0.003 | 1878.336 | 6.998 | 6.854 | 14.320 | 13.072 | 2.372 |
| c1_log | pair_409 | 0.003 | 1878.336 | 13.352 | 13.234 | 27.110 | 25.914 | 2.438 |
| c1_log | pair_410 | 0.003 | 1878.336 | 19.750 | 19.941 | 40.455 | 40.198 | 0.522 |
| c1_log | pair_411 | 0.003 | 1878.336 | 6.550 | 7.471 | 14.282 | 12.826 | 2.616 |
| c1_log | pair_412 | 0.003 | 1878.336 | 7.589 | 7.168 | 14.969 | 12.801 | 3.910 |
| c1_log | pair_413 | 0.003 | 1878.336 | 4.146 | 4.105 | 8.436 | 8.184 | 0.721 |
| c1_log | pair_414 | 0.003 | 1878.336 | 8.562 | 8.687 | 17.499 | 15.071 | 4.964 |
| c1_log | pair_415 | 0.003 | 1878.336 | 6.421 | 6.429 | 13.053 | 13.037 | 0.270 |
| c1_log | pair_416 | 0.003 | 1878.336 | 6.290 | 6.263 | 12.763 | 12.724 | 0.104 |
| c1_log | pair_417 | 0.003 | 1878.336 | 6.356 | 6.271 | 13.463 | 12.967 | 1.007 |
| c1_log | pair_418 | 0.003 | 1878.336 | 6.285 | 6.264 | 12.777 | 12.733 | 0.097 |
| c1_log | pair_419 | 0.003 | 1878.336 | 3.929 | 3.887 | 7.930 | 7.947 | 0.070 |
| c1_log | pair_420 | 0.003 | 1878.336 | 6.256 | 6.173 | 12.627 | 12.636 | 0.022 |
| c1_log | pair_421 | 0.003 | 1878.336 | 12.570 | 12.455 | 25.378 | 25.379 | 0.003 |
| c1_log | pair_422 | 0.003 | 1878.336 | 6.287 | 6.223 | 12.783 | 12.765 | 0.116 |
| c1_log | pair_423 | 0.003 | 1878.336 | 6.276 | 6.240 | 12.756 | 12.732 | 0.090 |
| c1_log | pair_424 | 0.003 | 1878.336 | 3.895 | 3.876 | 7.910 | 7.896 | 0.082 |
| c1_log | pair_425 | 0.003 | 1878.336 | 7.000 | 7.004 | 14.224 | 14.196 | 0.125 |
| c1_log | pair_426 | 0.003 | 1878.336 | 6.388 | 6.206 | 12.775 | 12.663 | 0.290 |
| c1_log | pair_427 | 0.003 | 1878.336 | 6.227 | 6.223 | 14.882 | 12.727 | 3.737 |
| c1_log | pair_428 | 0.003 | 1878.336 | 4.885 | 4.727 | 9.772 | 7.969 | 3.153 |
| c1_log | pair_429 | 0.003 | 1878.336 | 7.045 | 7.022 | 14.296 | 14.295 | 0.051 |
| c1_log | pair_430 | 0.003 | 1878.336 | 6.359 | 6.281 | 12.869 | 12.687 | 0.374 |
| c1_log | pair_431 | 0.003 | 1878.336 | 3.988 | 3.859 | 7.968 | 7.917 | 0.129 |
| c1_log | pair_432 | 0.003 | 1878.336 | 7.577 | 6.265 | 14.051 | 12.757 | 2.369 |
| c1_log | pair_433 | 0.003 | 1878.336 | 3.975 | 3.934 | 8.036 | 8.020 | 0.046 |
| c1_log | pair_434 | 0.003 | 1878.336 | 6.344 | 6.259 | 12.850 | 12.900 | 0.192 |
| c1_log | pair_435 | 0.003 | 1878.336 | 3.987 | 3.883 | 8.074 | 8.025 | 0.126 |
| c1_power2 | pair_001 | 0.006 | 1856.420 | 7.010 | 6.504 | 13.586 | 13.268 | 1.137 |
| c1_power2 | pair_002 | 0.006 | 1856.420 | 6.459 | 6.727 | 13.387 | 13.399 | 0.640 |
| c1_power2 | pair_003 | 0.006 | 1856.420 | 6.485 | 6.885 | 13.514 | 12.975 | 1.228 |
| c1_power2 | pair_004 | 0.006 | 1856.420 | 6.423 | 6.280 | 12.779 | 12.774 | 0.117 |
| c1_power2 | pair_005 | 0.006 | 1856.420 | 3.219 | 3.194 | 6.519 | 6.560 | 0.134 |
| c1_power2 | pair_006 | 0.006 | 1856.420 | 6.343 | 6.356 | 12.863 | 12.856 | 0.138 |
| c1_power2 | pair_007 | 0.006 | 1856.420 | 12.598 | 12.497 | 25.352 | 25.362 | 0.092 |
| c1_power2 | pair_008 | 0.006 | 1856.420 | 6.326 | 6.271 | 12.712 | 12.750 | 0.085 |
| c1_power2 | pair_009 | 0.006 | 1856.420 | 25.044 | 25.735 | 53.487 | 53.908 | 2.062 |
| c1_power2 | pair_010 | 0.006 | 1856.420 | 6.399 | 6.308 | 12.820 | 12.847 | 0.252 |
| c1_power2 | pair_011 | 0.006 | 1856.420 | 6.436 | 6.283 | 12.967 | 13.019 | 0.222 |
| c1_power2 | pair_012 | 0.006 | 1856.420 | 6.378 | 6.326 | 12.995 | 13.064 | 0.186 |
| c1_power2 | pair_013 | 0.006 | 1856.420 | 6.463 | 6.254 | 12.931 | 12.939 | 0.021 |
| c1_power2 | pair_014 | 0.006 | 1856.420 | 6.427 | 6.328 | 12.929 | 12.976 | 0.145 |
| c1_power2 | pair_015 | 0.006 | 1856.420 | 10.820 | 9.040 | 20.037 | 21.771 | 6.683 |
| c1_power2 | pair_016 | 0.006 | 1856.420 | 3.517 | 3.627 | 7.235 | 7.113 | 0.930 |
| c1_power2 | pair_017 | 0.006 | 1856.420 | 4.464 | 3.576 | 8.130 | 8.291 | 0.626 |
| c1_power2 | pair_018 | 0.006 | 1856.420 | 14.214 | 13.786 | 28.652 | 29.925 | 2.243 |
| c1_power2 | pair_019 | 0.006 | 1856.420 | 7.165 | 6.961 | 14.410 | 15.119 | 1.279 |
| c1_power2 | pair_020 | 0.006 | 1856.420 | 26.386 | 26.503 | 55.049 | 51.686 | 6.000 |
| c1_power2 | pair_021 | 0.006 | 1856.420 | 6.501 | 6.744 | 13.501 | 12.824 | 1.193 |
| c1_power2 | pair_022 | 0.006 | 1856.420 | 26.272 | 26.776 | 54.478 | 51.614 | 4.990 |
| c1_power2 | pair_023 | 0.006 | 1856.420 | 6.351 | 6.269 | 12.879 | 12.855 | 0.162 |
| c1_power2 | pair_024 | 0.006 | 1856.420 | 12.814 | 12.515 | 25.716 | 25.802 | 0.267 |
| c1_power2 | pair_025 | 0.006 | 1856.420 | 25.295 | 25.023 | 52.200 | 51.719 | 0.984 |
| c1_power2 | pair_026 | 0.006 | 1856.420 | 6.368 | 6.252 | 12.904 | 12.848 | 0.219 |
| c1_power2 | pair_027 | 0.006 | 1856.420 | 6.276 | 6.268 | 12.763 | 12.781 | 0.101 |
| c1_power2 | pair_028 | 0.006 | 1856.420 | 3.185 | 3.091 | 6.345 | 6.345 | 0.043 |
| c1_power2 | pair_029 | 0.006 | 1856.420 | 6.294 | 6.200 | 12.636 | 12.616 | 0.126 |
| c1_power2 | pair_030 | 0.006 | 1856.420 | 6.379 | 7.382 | 13.965 | 12.845 | 1.995 |
| c1_power2 | pair_031 | 0.006 | 1856.420 | 8.205 | 6.307 | 14.664 | 12.705 | 3.419 |
| c1_power2 | pair_032 | 0.006 | 1856.420 | 6.685 | 6.601 | 13.378 | 12.548 | 1.463 |
| c1_power2 | pair_033 | 0.006 | 1856.420 | 3.181 | 3.151 | 6.393 | 6.312 | 0.178 |
| c1_power2 | pair_034 | 0.006 | 1856.420 | 6.516 | 6.466 | 13.140 | 12.911 | 0.649 |
| c1_power2 | pair_035 | 0.006 | 1856.420 | 6.249 | 6.297 | 12.615 | 12.538 | 0.240 |
| c1_power2 | pair_036 | 0.006 | 1856.420 | 6.434 | 6.647 | 13.216 | 12.976 | 0.714 |
| c1_power2 | pair_037 | 0.006 | 1856.420 | 6.984 | 6.833 | 13.892 | 12.941 | 2.041 |
| c1_power2 | pair_038 | 0.006 | 1856.420 | 6.908 | 6.726 | 13.759 | 12.638 | 2.003 |
| c1_power2 | pair_039 | 0.006 | 1856.420 | 6.963 | 6.729 | 13.838 | 12.795 | 1.947 |
| c1_power2 | pair_040 | 0.006 | 1856.420 | 6.920 | 6.894 | 14.064 | 12.864 | 2.134 |
| c1_power2 | pair_041 | 0.006 | 1856.420 | 6.571 | 6.701 | 13.531 | 12.848 | 1.216 |
| c1_power2 | pair_042 | 0.006 | 1856.420 | 6.386 | 6.337 | 12.896 | 12.783 | 0.270 |
| c1_power2 | pair_043 | 0.006 | 1856.420 | 6.318 | 6.424 | 12.915 | 12.787 | 0.368 |
| c1_power2 | pair_044 | 0.006 | 1856.420 | 3.353 | 3.067 | 6.472 | 6.333 | 0.278 |
| c1_power2 | pair_045 | 0.006 | 1856.420 | 3.304 | 3.127 | 6.514 | 6.414 | 0.205 |
| c1_power2 | pair_046 | 0.006 | 1856.420 | 6.350 | 6.193 | 12.611 | 12.567 | 0.188 |
| c1_power2 | pair_047 | 0.006 | 1856.420 | 6.307 | 6.204 | 12.623 | 12.635 | 0.027 |
| c1_power2 | pair_048 | 0.006 | 1856.420 | 6.233 | 6.215 | 12.517 | 12.521 | 0.027 |
| c1_power2 | pair_049 | 0.006 | 1856.420 | 6.272 | 6.226 | 12.611 | 12.570 | 0.125 |
| c1_power2 | pair_050 | 0.006 | 1856.420 | 6.442 | 6.309 | 12.822 | 12.757 | 0.214 |
| c1_power2 | pair_051 | 0.006 | 1856.420 | 6.314 | 6.255 | 12.804 | 12.750 | 0.096 |
| c1_power2 | pair_052 | 0.006 | 1856.420 | 6.244 | 6.261 | 12.576 | 12.587 | 0.066 |
| c1_power2 | pair_053 | 0.006 | 1856.420 | 6.254 | 6.233 | 12.556 | 12.598 | 0.086 |
| c1_power2 | pair_054 | 0.006 | 1856.420 | 6.230 | 6.261 | 12.749 | 12.739 | 0.048 |
| c1_power2 | pair_055 | 0.006 | 1856.420 | 6.363 | 6.273 | 12.889 | 12.894 | 0.112 |
| c1_power2 | pair_056 | 0.006 | 1856.420 | 3.187 | 3.108 | 6.387 | 6.398 | 0.075 |
| c1_power2 | pair_057 | 0.006 | 1856.420 | 6.253 | 6.188 | 12.611 | 12.649 | 0.075 |
| c1_power2 | pair_058 | 0.006 | 1856.420 | 6.225 | 6.849 | 13.240 | 12.545 | 1.205 |
| c1_power2 | pair_059 | 0.006 | 1856.420 | 9.473 | 7.087 | 16.952 | 16.680 | 4.503 |
| c1_power2 | pair_060 | 0.006 | 1856.420 | 3.219 | 3.132 | 6.451 | 6.380 | 0.182 |
| c1_power2 | pair_061 | 0.006 | 1856.420 | 6.379 | 6.308 | 12.830 | 12.824 | 0.047 |
| c1_power2 | pair_062 | 0.006 | 1856.420 | 6.676 | 6.243 | 13.084 | 12.677 | 0.722 |
| c1_power2 | pair_063 | 0.006 | 1856.420 | 6.291 | 6.351 | 12.805 | 12.666 | 0.414 |
| c1_power2 | pair_064 | 0.006 | 1856.420 | 6.442 | 6.427 | 13.005 | 12.821 | 0.426 |
| c1_power2 | pair_065 | 0.006 | 1856.420 | 6.387 | 6.517 | 13.101 | 13.275 | 0.397 |
| c1_power2 | pair_066 | 0.006 | 1856.420 | 6.729 | 7.541 | 14.490 | 13.090 | 2.545 |
| c1_power2 | pair_067 | 0.006 | 1856.420 | 6.919 | 7.195 | 14.313 | 12.670 | 2.921 |
| c1_power2 | pair_068 | 0.006 | 1856.420 | 6.943 | 6.848 | 13.978 | 12.694 | 2.292 |
| c1_power2 | pair_069 | 0.006 | 1856.420 | 6.774 | 6.753 | 13.728 | 12.804 | 1.738 |
| c1_power2 | pair_070 | 0.006 | 1856.420 | 6.696 | 6.769 | 13.636 | 12.622 | 1.806 |
| c1_power2 | pair_071 | 0.006 | 1856.420 | 3.511 | 3.247 | 6.869 | 6.364 | 0.908 |
| c1_power2 | pair_072 | 0.006 | 1856.420 | 3.468 | 3.211 | 6.894 | 6.350 | 1.001 |
| c1_power2 | pair_073 | 0.006 | 1856.420 | 6.742 | 6.607 | 13.688 | 12.594 | 1.916 |
| c1_power2 | pair_074 | 0.006 | 1856.420 | 6.403 | 6.363 | 12.952 | 12.686 | 0.510 |
| c1_power2 | pair_075 | 0.006 | 1856.420 | 7.169 | 6.476 | 13.935 | 12.701 | 2.211 |
| c1_power2 | pair_076 | 0.006 | 1856.420 | 6.244 | 6.367 | 12.767 | 12.636 | 0.351 |
| c1_power2 | pair_077 | 0.006 | 1856.420 | 6.338 | 6.489 | 13.027 | 12.755 | 0.479 |
| c1_power2 | pair_078 | 0.006 | 1856.420 | 6.296 | 6.784 | 13.355 | 12.899 | 1.110 |
| c1_power2 | pair_079 | 0.006 | 1856.420 | 6.329 | 6.248 | 12.730 | 12.656 | 0.130 |
| c1_power2 | pair_080 | 0.006 | 1856.420 | 6.328 | 6.292 | 12.786 | 12.875 | 0.189 |
| c1_power2 | pair_081 | 0.006 | 1856.420 | 6.351 | 6.277 | 12.854 | 12.688 | 0.348 |
| c1_power2 | pair_082 | 0.006 | 1856.420 | 6.290 | 6.428 | 12.984 | 12.790 | 0.387 |
| c1_power2 | pair_083 | 0.006 | 1856.420 | 3.566 | 3.488 | 7.155 | 6.421 | 1.483 |
| c1_power2 | pair_084 | 0.006 | 1856.420 | 7.665 | 7.684 | 15.619 | 12.831 | 5.033 |
| c1_power2 | pair_085 | 0.006 | 1856.420 | 7.617 | 7.452 | 15.244 | 12.803 | 4.413 |
| c1_power2 | pair_086 | 0.006 | 1856.420 | 3.278 | 3.169 | 10.079 | 6.757 | 6.069 |
| c1_power2 | pair_087 | 0.006 | 1856.420 | 6.399 | 6.518 | 13.083 | 12.943 | 0.291 |
| c1_power2 | pair_088 | 0.006 | 1856.420 | 6.544 | 6.442 | 13.098 | 12.736 | 0.703 |
| c1_power2 | pair_089 | 0.006 | 1856.420 | 6.506 | 6.480 | 13.104 | 12.913 | 0.726 |
| c1_power2 | pair_090 | 0.006 | 1856.420 | 6.538 | 6.351 | 13.017 | 12.746 | 0.544 |
| c1_power2 | pair_091 | 0.006 | 1856.420 | 6.299 | 6.254 | 12.670 | 12.680 | 0.053 |
| c1_power2 | pair_092 | 0.006 | 1856.420 | 6.334 | 6.332 | 12.866 | 12.867 | 0.054 |
| c1_power2 | pair_093 | 0.006 | 1856.420 | 6.392 | 6.220 | 12.795 | 12.725 | 0.135 |
| c1_power2 | pair_094 | 0.006 | 1856.420 | 6.367 | 6.263 | 12.761 | 12.723 | 0.177 |
| c1_power2 | pair_095 | 0.006 | 1856.420 | 6.365 | 6.309 | 12.864 | 12.830 | 0.142 |
| c1_power2 | pair_096 | 0.006 | 1856.420 | 6.336 | 6.278 | 12.786 | 12.855 | 0.137 |
| c1_power2 | pair_097 | 0.006 | 1856.420 | 3.179 | 3.139 | 6.373 | 6.313 | 0.136 |
| c1_power2 | pair_098 | 0.006 | 1856.420 | 3.290 | 4.026 | 7.789 | 7.809 | 1.559 |
| c1_power2 | pair_099 | 0.006 | 1856.420 | 8.287 | 6.881 | 15.334 | 14.344 | 3.407 |
| c1_power2 | pair_100 | 0.006 | 1856.420 | 7.034 | 7.313 | 14.563 | 14.141 | 2.137 |
| c1_power2 | pair_101 | 0.006 | 1856.420 | 7.247 | 6.709 | 14.096 | 13.254 | 2.038 |
| c1_power2 | pair_102 | 0.006 | 1856.420 | 6.847 | 6.493 | 13.478 | 12.811 | 1.282 |
| c1_power2 | pair_103 | 0.006 | 1856.420 | 6.573 | 6.637 | 13.430 | 12.573 | 1.598 |
| c1_power2 | pair_104 | 0.006 | 1856.420 | 6.399 | 6.353 | 12.981 | 12.832 | 0.354 |
| c1_power2 | pair_105 | 0.006 | 1856.420 | 6.271 | 6.257 | 12.636 | 12.574 | 0.153 |
| c1_power2 | pair_106 | 0.006 | 1856.420 | 6.369 | 6.244 | 12.706 | 12.653 | 0.173 |
| c1_power2 | pair_107 | 0.006 | 1856.420 | 6.314 | 6.251 | 12.809 | 12.780 | 0.050 |
| c1_power2 | pair_108 | 0.006 | 1856.420 | 6.467 | 6.249 | 13.001 | 13.139 | 0.280 |
| c1_power2 | pair_109 | 0.006 | 1856.420 | 3.384 | 3.579 | 7.040 | 6.376 | 1.227 |
| c1_power2 | pair_110 | 0.006 | 1856.420 | 6.925 | 6.408 | 13.528 | 13.420 | 0.597 |
| c1_power2 | pair_111 | 0.006 | 1856.420 | 3.249 | 3.224 | 6.536 | 6.487 | 0.190 |
| c1_power2 | pair_112 | 0.006 | 1856.420 | 6.809 | 6.418 | 13.372 | 13.707 | 0.683 |
| c1_power2 | pair_113 | 0.006 | 1856.420 | 6.542 | 6.528 | 13.143 | 12.792 | 0.826 |
| c1_power2 | pair_114 | 0.006 | 1856.420 | 6.401 | 6.932 | 13.477 | 12.766 | 1.370 |
| c1_power2 | pair_115 | 0.006 | 1856.420 | 6.347 | 6.241 | 12.656 | 12.732 | 0.149 |
| c1_power2 | pair_116 | 0.006 | 1856.420 | 6.360 | 6.898 | 13.390 | 12.869 | 1.040 |
| c1_power2 | pair_117 | 0.006 | 1856.420 | 6.333 | 6.318 | 58.943 | 12.927 | 79.752 |
| c1_power2 | pair_118 | 0.006 | 1856.420 | 8.902 | 7.165 | 16.333 | 13.211 | 5.541 |
| c1_power2 | pair_119 | 0.006 | 1856.420 | 6.891 | 6.563 | 13.695 | 12.774 | 1.639 |
| c1_power2 | pair_120 | 0.006 | 1856.420 | 6.482 | 6.391 | 13.043 | 12.897 | 0.378 |
| c1_power2 | pair_121 | 0.006 | 1856.420 | 6.701 | 6.844 | 13.870 | 13.380 | 1.260 |
| c1_power2 | pair_122 | 0.006 | 1856.420 | 3.319 | 3.248 | 6.673 | 6.420 | 0.528 |
| c1_power2 | pair_123 | 0.006 | 1856.420 | 3.212 | 3.159 | 6.477 | 6.475 | 0.066 |
| c1_power2 | pair_124 | 0.006 | 1856.420 | 6.288 | 6.407 | 12.764 | 12.665 | 0.260 |
| c1_power2 | pair_125 | 0.006 | 1856.420 | 6.456 | 6.839 | 13.427 | 13.723 | 0.579 |
| c1_power2 | pair_126 | 0.006 | 1856.420 | 6.394 | 6.364 | 12.831 | 12.591 | 0.556 |
| c1_power2 | pair_127 | 0.006 | 1856.420 | 6.383 | 6.413 | 12.915 | 12.894 | 0.099 |
| c1_power2 | pair_128 | 0.006 | 1856.420 | 6.337 | 6.273 | 12.681 | 12.569 | 0.204 |
| c1_power2 | pair_129 | 0.006 | 1856.420 | 6.315 | 6.217 | 12.694 | 12.716 | 0.039 |
| c1_power2 | pair_130 | 0.006 | 1856.420 | 6.233 | 6.347 | 12.650 | 12.494 | 0.290 |
| c1_power2 | pair_131 | 0.006 | 1856.420 | 6.216 | 6.231 | 12.519 | 12.532 | 0.029 |
| c1_power2 | pair_132 | 0.006 | 1856.420 | 6.306 | 6.274 | 12.782 | 12.770 | 0.029 |
| c1_power2 | pair_133 | 0.006 | 1856.420 | 8.032 | 6.307 | 14.516 | 12.680 | 3.202 |
| c1_power2 | pair_134 | 0.006 | 1856.420 | 3.275 | 3.236 | 6.561 | 6.370 | 0.359 |
| c1_power2 | pair_135 | 0.006 | 1856.420 | 6.316 | 7.065 | 13.567 | 12.791 | 1.582 |
| c1_power2 | pair_136 | 0.006 | 1856.420 | 3.394 | 3.250 | 6.733 | 6.432 | 0.593 |
| c1_power2 | pair_137 | 0.006 | 1856.420 | 3.146 | 3.265 | 6.479 | 6.325 | 0.298 |
| c1_power2 | pair_138 | 0.006 | 1856.420 | 3.341 | 3.318 | 6.740 | 6.371 | 0.680 |
| c1_power2 | pair_139 | 0.006 | 1856.420 | 3.309 | 3.242 | 6.608 | 6.491 | 0.341 |
| c1_power2 | pair_140 | 0.006 | 1856.420 | 3.305 | 3.315 | 6.699 | 6.439 | 0.576 |
| c1_power2 | pair_141 | 0.006 | 1856.420 | 3.299 | 3.410 | 6.918 | 6.776 | 0.605 |
| c1_power2 | pair_142 | 0.006 | 1856.420 | 3.540 | 3.364 | 6.983 | 6.738 | 0.825 |
| c1_power2 | pair_143 | 0.006 | 1856.420 | 3.258 | 3.234 | 6.560 | 6.421 | 0.269 |
| c1_power2 | pair_144 | 0.006 | 1856.420 | 3.268 | 3.134 | 6.510 | 6.452 | 0.197 |
| c1_power2 | pair_145 | 0.006 | 1856.420 | 3.159 | 3.100 | 6.312 | 6.305 | 0.039 |
| c1_power2 | pair_146 | 0.006 | 1856.420 | 3.150 | 3.110 | 6.353 | 6.359 | 0.017 |
| c1_power2 | pair_147 | 0.006 | 1856.420 | 3.145 | 3.098 | 6.320 | 6.310 | 0.031 |
| c1_power2 | pair_148 | 0.006 | 1856.420 | 3.287 | 3.150 | 6.526 | 6.411 | 0.258 |
| c1_power2 | pair_149 | 0.006 | 1856.420 | 3.178 | 3.158 | 6.402 | 6.413 | 0.028 |
| c1_power2 | pair_150 | 0.006 | 1856.420 | 3.266 | 3.232 | 6.575 | 6.389 | 0.342 |
| c1_power2 | pair_151 | 0.006 | 1856.420 | 3.124 | 3.177 | 6.347 | 6.309 | 0.097 |
| c1_power2 | pair_152 | 0.006 | 1856.420 | 3.127 | 3.151 | 6.362 | 6.379 | 0.068 |
| c1_power2 | pair_153 | 0.006 | 1856.420 | 3.136 | 3.099 | 6.329 | 6.331 | 0.007 |
| c1_power2 | pair_154 | 0.006 | 1856.420 | 3.397 | 3.308 | 6.752 | 6.378 | 0.720 |
| c1_power2 | pair_155 | 0.006 | 1856.420 | 3.244 | 3.421 | 6.764 | 6.432 | 0.693 |
| c1_power2 | pair_156 | 0.006 | 1856.420 | 3.462 | 3.362 | 6.881 | 6.389 | 0.890 |
| c1_power2 | pair_157 | 0.006 | 1856.420 | 3.505 | 3.342 | 6.930 | 6.586 | 0.689 |
| c1_power2 | pair_158 | 0.006 | 1856.420 | 3.450 | 3.312 | 6.958 | 6.560 | 0.864 |
| c1_power2 | pair_159 | 0.006 | 1856.420 | 3.258 | 3.189 | 6.526 | 6.491 | 0.173 |
| c1_power2 | pair_160 | 0.006 | 1856.420 | 6.443 | 6.231 | 12.830 | 12.758 | 0.307 |
| c1_power2 | pair_161 | 0.006 | 1856.420 | 6.469 | 6.362 | 13.045 | 12.833 | 0.569 |
| c1_power2 | pair_162 | 0.006 | 1856.420 | 6.441 | 6.660 | 13.267 | 13.466 | 0.504 |
| c1_power2 | pair_163 | 0.006 | 1856.420 | 6.826 | 6.744 | 13.815 | 13.425 | 1.118 |
| c1_power2 | pair_164 | 0.006 | 1856.420 | 6.335 | 6.403 | 12.927 | 12.818 | 0.265 |
| c1_power2 | pair_165 | 0.006 | 1856.420 | 6.444 | 6.430 | 13.047 | 12.939 | 0.446 |
| c1_power2 | pair_166 | 0.006 | 1856.420 | 6.601 | 6.338 | 13.146 | 13.139 | 0.073 |
| c1_power2 | pair_167 | 0.006 | 1856.420 | 6.332 | 6.382 | 12.885 | 12.788 | 0.226 |
| c1_power2 | pair_168 | 0.006 | 1856.420 | 6.537 | 6.452 | 13.150 | 12.777 | 0.749 |
| c1_power2 | pair_169 | 0.006 | 1856.420 | 3.355 | 3.164 | 6.603 | 6.358 | 0.439 |
| c1_power2 | pair_170 | 0.006 | 1856.420 | 3.367 | 3.233 | 6.708 | 6.395 | 0.571 |
| c1_power2 | pair_171 | 0.006 | 1856.420 | 6.647 | 6.480 | 13.347 | 12.938 | 0.781 |
| c1_power2 | pair_172 | 0.006 | 1856.420 | 6.544 | 6.325 | 13.023 | 13.191 | 0.453 |
| c1_power2 | pair_173 | 0.006 | 1856.420 | 6.299 | 6.297 | 12.770 | 12.819 | 0.135 |
| c1_power2 | pair_174 | 0.006 | 1856.420 | 6.270 | 6.294 | 12.707 | 12.705 | 0.115 |
| c1_power2 | pair_175 | 0.006 | 1856.420 | 6.317 | 6.307 | 12.822 | 12.752 | 0.190 |
| c1_power2 | pair_176 | 0.006 | 1856.420 | 6.315 | 6.201 | 12.726 | 12.735 | 0.062 |
| c1_power2 | pair_177 | 0.006 | 1856.420 | 6.264 | 6.225 | 12.688 | 12.701 | 0.027 |
| c1_power2 | pair_178 | 0.006 | 1856.420 | 6.290 | 6.254 | 12.717 | 12.725 | 0.018 |
| c1_power2 | pair_179 | 0.006 | 1856.420 | 8.526 | 6.574 | 15.256 | 12.650 | 4.589 |
| c1_power2 | pair_180 | 0.006 | 1856.420 | 6.468 | 6.357 | 13.811 | 13.435 | 1.299 |
| c1_power2 | pair_181 | 0.006 | 1856.420 | 3.297 | 5.338 | 8.718 | 6.366 | 4.102 |
| c1_power2 | pair_182 | 0.006 | 1856.420 | 6.645 | 6.531 | 13.334 | 13.476 | 0.529 |
| c1_power2 | pair_183 | 0.006 | 1856.420 | 7.162 | 7.696 | 15.030 | 15.941 | 2.049 |
| c1_power2 | pair_184 | 0.006 | 1856.420 | 18.096 | 14.586 | 32.945 | 32.263 | 7.996 |
| c1_power2 | pair_185 | 0.006 | 1856.420 | 7.924 | 7.694 | 15.837 | 16.218 | 1.625 |
| c1_power2 | pair_186 | 0.006 | 1856.420 | 6.589 | 6.713 | 13.566 | 13.251 | 0.881 |
| c1_power2 | pair_187 | 0.006 | 1856.420 | 6.561 | 6.555 | 13.319 | 13.042 | 0.822 |
| c1_power2 | pair_188 | 0.006 | 1856.420 | 7.016 | 6.842 | 14.096 | 12.791 | 2.361 |
| c1_power2 | pair_189 | 0.006 | 1856.420 | 6.620 | 6.700 | 14.107 | 14.730 | 1.099 |
| c1_power2 | pair_190 | 0.006 | 1856.420 | 7.933 | 7.664 | 15.882 | 13.583 | 4.537 |
| c1_power2 | pair_191 | 0.006 | 1856.420 | 3.411 | 3.507 | 7.109 | 6.812 | 0.930 |
| c1_power2 | pair_192 | 0.006 | 1856.420 | 3.403 | 3.259 | 6.735 | 6.908 | 0.306 |
| c1_power2 | pair_193 | 0.006 | 1856.420 | 13.301 | 13.083 | 26.782 | 25.877 | 1.910 |
| c1_power2 | pair_194 | 0.006 | 1856.420 | 6.389 | 6.265 | 12.805 | 12.683 | 0.268 |
| c1_power2 | pair_195 | 0.006 | 1856.420 | 12.593 | 12.684 | 25.525 | 25.599 | 0.159 |
| c1_power2 | pair_196 | 0.006 | 1856.420 | 6.292 | 6.236 | 12.637 | 12.637 | 0.140 |
| c1_power2 | pair_197 | 0.006 | 1856.420 | 12.644 | 12.531 | 25.408 | 25.438 | 0.302 |
| c1_power2 | pair_198 | 0.006 | 1856.420 | 6.319 | 6.197 | 12.684 | 12.666 | 0.061 |
| c1_power2 | pair_199 | 0.006 | 1856.420 | 12.937 | 12.639 | 25.942 | 25.902 | 0.617 |
| c1_power2 | pair_200 | 0.006 | 1856.420 | 12.591 | 12.556 | 25.398 | 25.413 | 0.254 |
| c1_power2 | pair_201 | 0.006 | 1856.420 | 6.332 | 6.234 | 12.810 | 12.829 | 0.151 |
| c1_power2 | pair_202 | 0.006 | 1856.420 | 6.292 | 6.316 | 12.806 | 12.831 | 0.043 |
| c1_power2 | pair_203 | 0.006 | 1856.420 | 3.221 | 3.156 | 6.443 | 6.422 | 0.210 |
| c1_power2 | pair_204 | 0.006 | 1856.420 | 6.244 | 6.219 | 12.604 | 12.560 | 0.144 |
| c1_power2 | pair_205 | 0.006 | 1856.420 | 6.348 | 6.265 | 12.724 | 12.660 | 0.128 |
| c1_power2 | pair_206 | 0.006 | 1856.420 | 6.333 | 6.432 | 12.893 | 12.680 | 0.444 |
| c1_power2 | pair_207 | 0.006 | 1856.420 | 6.329 | 6.296 | 12.810 | 12.829 | 0.203 |
| c1_power2 | pair_208 | 0.006 | 1856.420 | 10.797 | 7.766 | 18.763 | 13.025 | 10.130 |
| c1_power2 | pair_209 | 0.006 | 1856.420 | 6.501 | 6.307 | 12.974 | 12.966 | 0.084 |
| c1_power2 | pair_210 | 0.006 | 1856.420 | 6.447 | 6.969 | 13.562 | 12.761 | 1.388 |
| c1_power2 | pair_211 | 0.006 | 1856.420 | 6.486 | 6.321 | 12.999 | 12.957 | 0.186 |
| c1_power2 | pair_212 | 0.006 | 1856.420 | 3.169 | 3.117 | 6.391 | 6.374 | 0.043 |
| c1_power2 | pair_213 | 0.006 | 1856.420 | 3.189 | 3.102 | 6.394 | 6.411 | 0.032 |
| c1_power2 | pair_214 | 0.006 | 1856.420 | 6.383 | 6.343 | 12.843 | 12.771 | 0.182 |
| c1_power2 | pair_215 | 0.006 | 1856.420 | 6.483 | 6.361 | 12.956 | 13.006 | 0.159 |
| c1_power2 | pair_216 | 0.006 | 1856.420 | 6.476 | 6.392 | 13.034 | 13.132 | 0.197 |
| c1_power2 | pair_217 | 0.006 | 1856.420 | 6.370 | 6.307 | 12.808 | 12.781 | 0.073 |
| c1_power2 | pair_218 | 0.006 | 1856.420 | 6.710 | 7.157 | 14.029 | 12.787 | 2.157 |
| c1_power2 | pair_219 | 0.006 | 1856.420 | 6.353 | 6.300 | 12.896 | 12.896 | 0.022 |
| c1_power2 | pair_220 | 0.006 | 1856.420 | 6.405 | 6.334 | 12.845 | 12.864 | 0.061 |
| c1_power2 | pair_221 | 0.006 | 1856.420 | 6.278 | 6.307 | 12.695 | 12.680 | 0.050 |
| c1_power2 | pair_222 | 0.006 | 1856.420 | 6.348 | 6.394 | 12.971 | 12.850 | 0.264 |
| c1_power2 | pair_223 | 0.006 | 1856.420 | 6.409 | 6.317 | 12.912 | 12.938 | 0.070 |
| c1_power2 | pair_224 | 0.006 | 1856.420 | 3.372 | 3.181 | 6.651 | 6.424 | 0.420 |
| c1_power2 | pair_225 | 0.006 | 1856.420 | 6.645 | 6.891 | 13.803 | 13.343 | 1.332 |
| c1_power2 | pair_226 | 0.006 | 1856.420 | 6.463 | 6.365 | 12.951 | 12.989 | 0.136 |
| c1_power2 | pair_227 | 0.006 | 1856.420 | 6.475 | 6.329 | 13.072 | 13.276 | 0.425 |
| c1_power2 | pair_228 | 0.006 | 1856.420 | 6.330 | 6.418 | 13.002 | 12.924 | 0.332 |
| c1_power2 | pair_229 | 0.006 | 1856.420 | 6.428 | 6.248 | 12.842 | 12.758 | 0.255 |
| c1_power2 | pair_230 | 0.006 | 1856.420 | 6.385 | 6.235 | 12.780 | 12.818 | 0.206 |
| c1_power2 | pair_231 | 0.006 | 1856.420 | 6.271 | 6.199 | 12.625 | 12.615 | 0.115 |
| c1_power2 | pair_232 | 0.006 | 1856.420 | 3.151 | 3.231 | 6.447 | 6.333 | 0.222 |
| c1_power2 | pair_233 | 0.006 | 1856.420 | 3.138 | 3.105 | 6.330 | 6.321 | 0.024 |
| c1_power2 | pair_234 | 0.006 | 1856.420 | 12.637 | 14.344 | 27.511 | 25.904 | 3.233 |
| c1_power2 | pair_235 | 0.006 | 1856.420 | 6.295 | 6.294 | 12.725 | 12.786 | 0.164 |
| c1_power2 | pair_236 | 0.006 | 1856.420 | 25.340 | 26.251 | 52.899 | 53.111 | 1.088 |
| c1_power2 | pair_237 | 0.006 | 1856.420 | 6.342 | 6.331 | 12.771 | 12.782 | 0.161 |
| c1_power2 | pair_238 | 0.006 | 1856.420 | 25.224 | 25.064 | 51.571 | 51.439 | 0.237 |
| c1_power2 | pair_239 | 0.006 | 1856.420 | 6.391 | 6.376 | 13.036 | 12.905 | 0.408 |
| c1_power2 | pair_240 | 0.006 | 1856.420 | 13.172 | 12.931 | 26.546 | 25.660 | 1.900 |
| c1_power2 | pair_241 | 0.006 | 1856.420 | 25.101 | 25.057 | 52.201 | 51.925 | 1.114 |
| c1_power2 | pair_242 | 0.006 | 1856.420 | 7.169 | 6.675 | 14.160 | 12.778 | 2.476 |
| c1_power2 | pair_243 | 0.006 | 1856.420 | 6.442 | 6.291 | 12.991 | 12.823 | 0.467 |
| c1_power2 | pair_244 | 0.006 | 1856.420 | 3.240 | 3.116 | 6.430 | 6.333 | 0.172 |
| c1_power2 | pair_245 | 0.006 | 1856.420 | 6.421 | 6.300 | 12.904 | 12.806 | 0.407 |
| c1_power2 | pair_246 | 0.006 | 1856.420 | 6.490 | 6.207 | 12.893 | 12.702 | 0.380 |
| c1_power2 | pair_247 | 0.006 | 1856.420 | 6.290 | 6.272 | 12.777 | 12.804 | 0.064 |
| c1_power2 | pair_248 | 0.006 | 1856.420 | 6.371 | 6.280 | 12.828 | 12.710 | 0.234 |
| c1_power2 | pair_249 | 0.006 | 1856.420 | 6.453 | 6.290 | 12.921 | 12.750 | 0.301 |
| c1_power2 | pair_250 | 0.006 | 1856.420 | 6.274 | 6.345 | 12.798 | 12.841 | 0.090 |
| c1_power2 | pair_251 | 0.006 | 1856.420 | 3.179 | 3.107 | 6.362 | 6.394 | 0.072 |
| c1_power2 | pair_252 | 0.006 | 1856.420 | 3.139 | 3.101 | 6.308 | 6.345 | 0.087 |
| c1_power2 | pair_253 | 0.006 | 1856.420 | 6.269 | 6.242 | 12.623 | 12.523 | 0.178 |
| c1_power2 | pair_254 | 0.006 | 1856.420 | 6.363 | 6.205 | 12.665 | 12.588 | 0.177 |
| c1_power2 | pair_255 | 0.006 | 1856.420 | 6.304 | 6.190 | 12.628 | 12.662 | 0.060 |
| c1_power2 | pair_256 | 0.006 | 1856.420 | 6.227 | 6.214 | 12.560 | 12.579 | 0.047 |
| c1_power2 | pair_257 | 0.006 | 1856.420 | 6.272 | 6.201 | 12.615 | 12.581 | 0.114 |
| c1_power2 | pair_258 | 0.006 | 1856.420 | 6.256 | 6.246 | 12.714 | 12.725 | 0.064 |
| c1_power2 | pair_259 | 0.006 | 1856.420 | 6.278 | 6.180 | 12.609 | 12.530 | 0.138 |
| c1_power2 | pair_260 | 0.006 | 1856.420 | 6.339 | 6.205 | 12.674 | 12.719 | 0.153 |
| c1_power2 | pair_261 | 0.006 | 1856.420 | 6.279 | 6.201 | 12.677 | 12.676 | 0.178 |
| c1_power2 | pair_262 | 0.006 | 1856.420 | 6.307 | 7.430 | 14.265 | 13.758 | 1.933 |
| c1_power2 | pair_263 | 0.006 | 1856.420 | 3.575 | 5.730 | 9.385 | 8.962 | 3.310 |
| c1_power2 | pair_264 | 0.006 | 1856.420 | 7.772 | 7.507 | 15.431 | 15.757 | 2.663 |
| c1_power2 | pair_265 | 0.006 | 1856.420 | 8.038 | 7.542 | 15.917 | 16.966 | 2.803 |
| c1_power2 | pair_266 | 0.006 | 1856.420 | 7.454 | 7.445 | 15.193 | 14.225 | 3.114 |
| c1_power2 | pair_267 | 0.006 | 1856.420 | 7.623 | 7.867 | 15.864 | 15.122 | 3.703 |
| c1_power2 | pair_268 | 0.006 | 1856.420 | 6.967 | 6.898 | 14.253 | 13.106 | 2.515 |
| c1_power2 | pair_269 | 0.006 | 1856.420 | 3.582 | 3.193 | 6.846 | 6.343 | 0.933 |
| c1_power2 | pair_270 | 0.006 | 1856.420 | 3.334 | 3.419 | 6.800 | 6.310 | 0.856 |
| c1_power2 | pair_271 | 0.006 | 1856.420 | 6.742 | 6.516 | 13.575 | 12.808 | 1.501 |
| c1_power2 | pair_272 | 0.006 | 1856.420 | 6.444 | 6.311 | 13.005 | 13.018 | 0.239 |
| c1_power2 | pair_273 | 0.006 | 1856.420 | 6.470 | 6.531 | 13.186 | 13.377 | 0.437 |
| c1_power2 | pair_274 | 0.006 | 1856.420 | 6.459 | 6.432 | 13.130 | 13.132 | 0.394 |
| c1_power2 | pair_275 | 0.006 | 1856.420 | 6.481 | 6.398 | 13.098 | 12.729 | 0.642 |
| c1_power2 | pair_276 | 0.006 | 1856.420 | 6.426 | 6.542 | 13.203 | 12.888 | 0.655 |
| c1_power2 | pair_277 | 0.006 | 1856.420 | 6.340 | 6.293 | 12.803 | 12.822 | 0.078 |
| c1_power2 | pair_278 | 0.006 | 1856.420 | 6.411 | 6.262 | 12.897 | 12.803 | 0.251 |
| c1_power2 | pair_279 | 0.006 | 1856.420 | 6.340 | 6.291 | 12.851 | 12.904 | 0.262 |
| c1_power2 | pair_280 | 0.006 | 1856.420 | 6.283 | 7.025 | 13.833 | 13.527 | 1.421 |
| c1_power2 | pair_281 | 0.006 | 1856.420 | 3.286 | 3.115 | 6.444 | 6.340 | 0.327 |
| c1_power2 | pair_282 | 0.006 | 1856.420 | 6.307 | 6.226 | 12.706 | 12.701 | 0.088 |
| c1_power2 | pair_283 | 0.006 | 1856.420 | 6.326 | 6.244 | 12.799 | 12.799 | 0.029 |
| c1_power2 | pair_284 | 0.006 | 1856.420 | 6.379 | 6.331 | 12.944 | 12.857 | 0.332 |
| c1_power2 | pair_285 | 0.006 | 1856.420 | 6.297 | 6.277 | 12.772 | 12.749 | 0.134 |
| c1_power2 | pair_286 | 0.006 | 1856.420 | 3.233 | 3.143 | 6.411 | 6.297 | 0.274 |
| c1_power2 | pair_287 | 0.006 | 1856.420 | 3.157 | 3.137 | 6.371 | 6.383 | 0.070 |
| c1_power2 | pair_288 | 0.006 | 1856.420 | 6.221 | 6.294 | 12.712 | 12.729 | 0.092 |
| c1_power2 | pair_289 | 0.006 | 1856.420 | 6.298 | 7.072 | 13.593 | 12.899 | 1.470 |
| c1_power2 | pair_290 | 0.006 | 1856.420 | 6.275 | 6.366 | 12.896 | 12.854 | 0.128 |
| c1_power2 | pair_291 | 0.006 | 1856.420 | 6.300 | 6.277 | 12.793 | 12.853 | 0.121 |
| c1_power2 | pair_292 | 0.006 | 1856.420 | 6.314 | 6.318 | 12.868 | 12.741 | 0.240 |
| c1_power2 | pair_293 | 0.006 | 1856.420 | 6.294 | 6.432 | 12.902 | 12.922 | 0.201 |
| c1_power2 | pair_294 | 0.006 | 1856.420 | 6.405 | 6.255 | 13.151 | 13.225 | 0.460 |
| c1_power2 | pair_295 | 0.006 | 1856.420 | 10.700 | 7.265 | 18.198 | 12.902 | 9.289 |
| c1_power2 | pair_296 | 0.006 | 1856.420 | 6.306 | 6.342 | 12.864 | 12.896 | 0.166 |
| c1_power2 | pair_297 | 0.006 | 1856.420 | 7.157 | 6.389 | 14.104 | 12.806 | 2.364 |
| c1_power2 | pair_298 | 0.006 | 1856.420 | 3.148 | 3.252 | 6.462 | 6.299 | 0.330 |
| c1_power2 | pair_299 | 0.006 | 1856.420 | 6.335 | 6.241 | 12.767 | 12.828 | 0.194 |
| c1_power2 | pair_300 | 0.006 | 1856.420 | 6.300 | 6.242 | 12.706 | 12.701 | 0.027 |
| c1_power2 | pair_301 | 0.006 | 1856.420 | 6.273 | 6.166 | 12.595 | 12.576 | 0.091 |
| c1_power2 | pair_302 | 0.006 | 1856.420 | 3.175 | 3.108 | 6.338 | 6.346 | 0.076 |
| c1_power2 | pair_303 | 0.006 | 1856.420 | 3.150 | 3.100 | 6.289 | 6.301 | 0.042 |
| c1_power2 | pair_304 | 0.006 | 1856.420 | 6.332 | 6.228 | 12.723 | 12.660 | 0.191 |
| c1_power2 | pair_305 | 0.006 | 1856.420 | 6.271 | 6.259 | 12.677 | 12.689 | 0.104 |
| c1_power2 | pair_306 | 0.006 | 1856.420 | 6.357 | 6.171 | 12.633 | 12.620 | 0.048 |
| c1_power2 | pair_307 | 0.006 | 1856.420 | 6.221 | 6.192 | 12.620 | 12.631 | 0.033 |
| c1_power2 | pair_308 | 0.006 | 1856.420 | 6.252 | 6.222 | 12.610 | 12.594 | 0.102 |
| c1_power2 | pair_309 | 0.006 | 1856.420 | 6.282 | 6.232 | 12.748 | 12.759 | 0.162 |
| c1_power2 | pair_310 | 0.006 | 1856.420 | 6.247 | 6.216 | 12.619 | 12.639 | 0.062 |
| c1_power2 | pair_311 | 0.006 | 1856.420 | 6.397 | 6.243 | 12.827 | 12.712 | 0.317 |
| c1_power2 | pair_312 | 0.006 | 1856.420 | 6.328 | 6.237 | 12.786 | 12.742 | 0.090 |
| c1_power2 | pair_313 | 0.006 | 1856.420 | 6.326 | 6.187 | 12.716 | 12.787 | 0.166 |
| c1_power2 | pair_314 | 0.006 | 1856.420 | 3.178 | 3.133 | 6.347 | 6.284 | 0.152 |
| c1_power2 | pair_315 | 0.006 | 1856.420 | 6.246 | 6.241 | 12.670 | 12.674 | 0.077 |
| c1_power2 | pair_316 | 0.006 | 1856.420 | 6.317 | 6.280 | 12.775 | 12.778 | 0.065 |
| c1_power2 | pair_317 | 0.006 | 1856.420 | 3.145 | 3.133 | 6.347 | 6.348 | 0.017 |
| c1_power2 | pair_318 | 0.006 | 1856.420 | 3.141 | 3.141 | 6.352 | 6.365 | 0.030 |
| c1_power2 | pair_319 | 0.006 | 1856.420 | 6.317 | 6.302 | 12.814 | 12.785 | 0.109 |
| c1_power2 | pair_320 | 0.006 | 1856.420 | 6.263 | 6.284 | 12.683 | 12.703 | 0.129 |
| c1_power2 | pair_321 | 0.006 | 1856.420 | 6.265 | 6.328 | 12.764 | 12.772 | 0.101 |
| c1_power2 | pair_322 | 0.006 | 1856.420 | 6.316 | 6.287 | 13.436 | 13.004 | 1.100 |
| c1_power2 | pair_323 | 0.006 | 1856.420 | 6.887 | 6.491 | 13.526 | 13.021 | 1.284 |
| c1_power2 | pair_324 | 0.006 | 1856.420 | 6.431 | 6.252 | 12.916 | 12.906 | 0.266 |
| c1_power2 | pair_325 | 0.006 | 1856.420 | 6.315 | 6.263 | 12.736 | 12.715 | 0.056 |
| c1_power2 | pair_326 | 0.006 | 1856.420 | 6.692 | 7.191 | 14.128 | 12.979 | 2.106 |
| c1_power2 | pair_327 | 0.006 | 1856.420 | 6.310 | 6.221 | 12.756 | 12.795 | 0.134 |
| c1_power2 | pair_328 | 0.006 | 1856.420 | 6.253 | 6.298 | 12.787 | 12.749 | 0.156 |
| c1_power2 | pair_329 | 0.006 | 1856.420 | 3.179 | 3.106 | 6.378 | 6.417 | 0.103 |
| c1_power2 | pair_330 | 0.006 | 1856.420 | 6.242 | 6.293 | 12.718 | 12.726 | 0.079 |
| c1_power2 | pair_331 | 0.006 | 1856.420 | 3.183 | 3.130 | 6.409 | 6.335 | 0.151 |
| c1_power2 | pair_332 | 0.006 | 1856.420 | 3.133 | 3.115 | 6.347 | 6.338 | 0.023 |
| c1_power2 | pair_333 | 0.006 | 1856.420 | 6.276 | 6.342 | 12.786 | 12.725 | 0.107 |
| c1_power2 | pair_334 | 0.006 | 1856.420 | 6.257 | 6.313 | 13.771 | 12.808 | 1.821 |
| c1_power2 | pair_335 | 0.006 | 1856.420 | 6.433 | 6.322 | 12.929 | 12.799 | 0.310 |
| c1_power2 | pair_336 | 0.006 | 1856.420 | 6.574 | 6.339 | 13.066 | 13.082 | 0.277 |
| c1_power2 | pair_337 | 0.006 | 1856.420 | 7.173 | 7.685 | 15.004 | 13.080 | 3.703 |
| c1_power2 | pair_338 | 0.006 | 1856.420 | 6.330 | 6.320 | 12.843 | 12.912 | 0.152 |
| c1_power2 | pair_339 | 0.006 | 1856.420 | 6.480 | 6.331 | 13.012 | 12.835 | 0.331 |
| c1_power2 | pair_340 | 0.006 | 1856.420 | 6.380 | 6.253 | 12.827 | 12.784 | 0.169 |
| c1_power2 | pair_341 | 0.006 | 1856.420 | 6.447 | 6.414 | 13.088 | 13.250 | 0.427 |
| c1_power2 | pair_342 | 0.006 | 1856.420 | 6.432 | 6.942 | 13.635 | 13.607 | 1.011 |
| c1_power2 | pair_343 | 0.006 | 1856.420 | 3.164 | 3.152 | 6.375 | 6.414 | 0.084 |
| c1_power2 | pair_344 | 0.006 | 1856.420 | 6.331 | 6.228 | 12.747 | 12.755 | 0.189 |
| c1_power2 | pair_345 | 0.006 | 1856.420 | 3.145 | 3.098 | 6.318 | 6.296 | 0.043 |
| c1_power2 | pair_346 | 0.006 | 1856.420 | 3.135 | 3.089 | 6.297 | 6.288 | 0.023 |
| c1_power2 | pair_347 | 0.006 | 1856.420 | 3.130 | 3.139 | 6.334 | 6.300 | 0.078 |
| c1_power2 | pair_348 | 0.006 | 1856.420 | 3.195 | 3.075 | 6.312 | 6.352 | 0.081 |
| c1_power2 | pair_349 | 0.006 | 1856.420 | 3.164 | 3.095 | 6.321 | 6.319 | 0.039 |
| c1_power2 | pair_350 | 0.006 | 1856.420 | 3.155 | 3.192 | 6.419 | 6.333 | 0.165 |
| c1_power2 | pair_351 | 0.006 | 1856.420 | 3.146 | 3.168 | 6.377 | 6.376 | 0.112 |
| c1_power2 | pair_352 | 0.006 | 1856.420 | 3.175 | 3.792 | 7.060 | 6.299 | 1.322 |
| c1_power2 | pair_353 | 0.006 | 1856.420 | 3.232 | 3.117 | 6.439 | 6.356 | 0.217 |
| c1_power2 | pair_354 | 0.006 | 1856.420 | 3.245 | 3.130 | 6.461 | 6.288 | 0.312 |
| c1_power2 | pair_355 | 0.006 | 1856.420 | 3.122 | 3.190 | 6.375 | 6.329 | 0.109 |
| c1_power2 | pair_356 | 0.006 | 1856.420 | 3.150 | 3.177 | 6.387 | 6.336 | 0.129 |
| c1_power2 | pair_357 | 0.006 | 1856.420 | 3.148 | 3.097 | 6.313 | 6.299 | 0.052 |
| c1_power2 | pair_358 | 0.006 | 1856.420 | 3.215 | 3.135 | 6.398 | 6.340 | 0.136 |
| c1_power2 | pair_359 | 0.006 | 1856.420 | 3.202 | 3.713 | 6.997 | 6.553 | 0.909 |
| c1_power2 | pair_360 | 0.006 | 1856.420 | 3.413 | 3.098 | 6.567 | 6.366 | 0.402 |
| c1_power2 | pair_361 | 0.006 | 1856.420 | 3.140 | 3.105 | 6.361 | 6.363 | 0.012 |
| c1_power2 | pair_362 | 0.006 | 1856.420 | 3.141 | 3.192 | 6.390 | 6.372 | 0.039 |
| c1_power2 | pair_363 | 0.006 | 1856.420 | 3.231 | 3.114 | 6.417 | 6.346 | 0.135 |
| c1_power2 | pair_364 | 0.006 | 1856.420 | 3.130 | 3.084 | 6.287 | 6.308 | 0.074 |
| c1_power2 | pair_365 | 0.006 | 1856.420 | 3.163 | 3.107 | 6.317 | 6.319 | 0.013 |
| c1_power2 | pair_366 | 0.006 | 1856.420 | 3.148 | 3.187 | 6.408 | 6.335 | 0.143 |
| c1_power2 | pair_367 | 0.006 | 1856.420 | 3.163 | 3.159 | 6.415 | 6.418 | 0.036 |
| c1_power2 | pair_368 | 0.006 | 1856.420 | 3.159 | 3.105 | 6.329 | 6.376 | 0.098 |
| c1_power2 | pair_369 | 0.006 | 1856.420 | 3.125 | 3.116 | 6.310 | 6.289 | 0.037 |
| c1_power2 | pair_370 | 0.006 | 1856.420 | 6.224 | 6.217 | 12.538 | 12.549 | 0.036 |
| c1_power2 | pair_371 | 0.006 | 1856.420 | 12.558 | 12.498 | 25.513 | 25.577 | 0.214 |
| c1_power2 | pair_372 | 0.006 | 1856.420 | 6.306 | 6.224 | 12.633 | 12.637 | 0.093 |
| c1_power2 | pair_373 | 0.006 | 1856.420 | 12.515 | 12.443 | 25.831 | 25.469 | 0.761 |
| c1_power2 | pair_374 | 0.006 | 1856.420 | 6.264 | 6.204 | 12.698 | 12.687 | 0.035 |
| c1_power2 | pair_375 | 0.006 | 1856.420 | 12.465 | 12.540 | 25.469 | 25.486 | 0.109 |
| c1_power2 | pair_376 | 0.006 | 1856.420 | 12.686 | 12.437 | 25.838 | 25.726 | 0.490 |
| c1_power2 | pair_377 | 0.006 | 1856.420 | 6.380 | 6.226 | 12.853 | 12.825 | 0.107 |
| c1_power2 | pair_378 | 0.006 | 1856.420 | 6.316 | 7.210 | 16.519 | 15.755 | 4.323 |
| c1_power2 | pair_379 | 0.006 | 1856.420 | 3.413 | 3.288 | 6.794 | 6.566 | 0.648 |
| c1_power2 | pair_380 | 0.006 | 1856.420 | 6.470 | 6.371 | 13.025 | 12.720 | 0.640 |
| c1_power2 | pair_381 | 0.006 | 1856.420 | 6.306 | 6.251 | 12.701 | 12.701 | 0.078 |
| c1_power2 | pair_382 | 0.006 | 1856.420 | 6.309 | 6.172 | 12.590 | 12.546 | 0.076 |
| c1_power2 | pair_383 | 0.006 | 1856.420 | 6.264 | 6.198 | 12.605 | 12.596 | 0.078 |
| c1_power2 | pair_384 | 0.006 | 1856.420 | 6.307 | 6.237 | 12.765 | 12.770 | 0.047 |
| c1_power2 | pair_385 | 0.006 | 1856.420 | 6.245 | 6.217 | 12.585 | 12.577 | 0.101 |
| c1_power2 | pair_386 | 0.006 | 1856.420 | 6.354 | 6.234 | 12.709 | 12.688 | 0.130 |
| c1_power2 | pair_387 | 0.006 | 1856.420 | 6.226 | 6.235 | 12.713 | 12.736 | 0.141 |
| c1_power2 | pair_388 | 0.006 | 1856.420 | 6.330 | 6.297 | 12.852 | 12.807 | 0.260 |
| c1_power2 | pair_389 | 0.006 | 1856.420 | 3.153 | 3.140 | 6.357 | 6.357 | 0.039 |
| c1_power2 | pair_390 | 0.006 | 1856.420 | 6.272 | 6.192 | 12.631 | 12.665 | 0.063 |
| c1_power2 | pair_391 | 0.006 | 1856.420 | 6.271 | 6.186 | 12.568 | 12.603 | 0.079 |
| c1_power2 | pair_392 | 0.006 | 1856.420 | 25.153 | 25.034 | 51.478 | 51.512 | 0.133 |
| c1_power2 | pair_393 | 0.006 | 1856.420 | 6.369 | 6.336 | 12.905 | 12.648 | 0.478 |
| c1_power2 | pair_394 | 0.006 | 1856.420 | 12.634 | 12.460 | 25.491 | 25.456 | 0.269 |
| c1_power2 | pair_395 | 0.006 | 1856.420 | 25.195 | 25.209 | 51.658 | 51.220 | 0.989 |
| c1_power2 | pair_396 | 0.006 | 1856.420 | 6.281 | 6.314 | 16.695 | 12.786 | 6.775 |
| c1_power2 | pair_397 | 0.006 | 1856.420 | 6.305 | 6.215 | 12.711 | 12.678 | 0.107 |
| c1_power2 | pair_398 | 0.006 | 1856.420 | 3.195 | 3.131 | 6.391 | 6.405 | 0.030 |
| c1_power2 | pair_399 | 0.006 | 1856.420 | 6.245 | 6.248 | 15.265 | 14.658 | 2.934 |
| c1_power2 | pair_400 | 0.006 | 1856.420 | 10.108 | 6.692 | 17.009 | 13.115 | 7.144 |
| c1_power2 | pair_401 | 0.006 | 1856.420 | 6.848 | 6.883 | 14.075 | 12.861 | 2.249 |
| c1_power2 | pair_402 | 0.006 | 1856.420 | 6.985 | 6.644 | 13.808 | 13.152 | 1.675 |
| c1_power2 | pair_403 | 0.006 | 1856.420 | 6.892 | 6.685 | 13.806 | 13.612 | 1.350 |
| c1_power2 | pair_404 | 0.006 | 1856.420 | 6.555 | 6.362 | 13.156 | 13.337 | 0.367 |
| c1_power2 | pair_405 | 0.006 | 1856.420 | 6.608 | 6.426 | 13.245 | 13.089 | 0.531 |
| c1_power2 | pair_406 | 0.006 | 1856.420 | 3.333 | 3.261 | 6.745 | 6.702 | 0.482 |
| c1_power2 | pair_407 | 0.006 | 1856.420 | 6.782 | 6.553 | 13.521 | 13.705 | 0.735 |
| c1_power2 | pair_408 | 0.006 | 1856.420 | 6.387 | 6.650 | 13.294 | 12.899 | 0.711 |
| c1_power2 | pair_409 | 0.006 | 1856.420 | 12.791 | 13.065 | 26.414 | 25.948 | 1.097 |
| c1_power2 | pair_410 | 0.006 | 1856.420 | 26.016 | 26.771 | 54.039 | 54.557 | 1.406 |
| c1_power2 | pair_411 | 0.006 | 1856.420 | 6.383 | 6.255 | 12.849 | 12.876 | 0.091 |
| c1_power2 | pair_412 | 0.006 | 1856.420 | 6.315 | 6.309 | 12.835 | 12.832 | 0.101 |
| c1_power2 | pair_413 | 0.006 | 1856.420 | 3.182 | 3.127 | 6.395 | 6.374 | 0.039 |
| c1_power2 | pair_414 | 0.006 | 1856.420 | 6.292 | 6.284 | 12.824 | 12.882 | 0.172 |
| c1_power2 | pair_415 | 0.006 | 1856.420 | 6.348 | 6.236 | 12.785 | 12.745 | 0.079 |
| c1_power2 | pair_416 | 0.006 | 1856.420 | 6.394 | 6.307 | 12.889 | 12.899 | 0.244 |
| c1_power2 | pair_417 | 0.006 | 1856.420 | 6.347 | 6.503 | 13.042 | 12.842 | 0.524 |
| c1_power2 | pair_418 | 0.006 | 1856.420 | 6.582 | 6.604 | 13.430 | 12.910 | 1.005 |
| c1_power2 | pair_419 | 0.006 | 1856.420 | 3.346 | 3.392 | 6.786 | 6.347 | 0.808 |
| c1_power2 | pair_420 | 0.006 | 1856.420 | 6.896 | 7.000 | 14.141 | 12.923 | 2.227 |
| c1_power2 | pair_421 | 0.006 | 1856.420 | 12.739 | 12.625 | 25.747 | 25.649 | 0.522 |
| c1_power2 | pair_422 | 0.006 | 1856.420 | 7.022 | 6.811 | 14.321 | 13.808 | 1.770 |
| c1_power2 | pair_423 | 0.006 | 1856.420 | 7.507 | 7.126 | 14.873 | 14.108 | 2.446 |
| c1_power2 | pair_424 | 0.006 | 1856.420 | 4.073 | 4.872 | 9.079 | 9.022 | 1.204 |
| c1_power2 | pair_425 | 0.006 | 1856.420 | 7.800 | 8.719 | 16.768 | 14.227 | 4.845 |
| c1_power2 | pair_426 | 0.006 | 1856.420 | 7.619 | 7.511 | 15.463 | 16.306 | 1.909 |
| c1_power2 | pair_427 | 0.006 | 1856.420 | 7.162 | 6.840 | 14.314 | 13.648 | 1.479 |
| c1_power2 | pair_428 | 0.006 | 1856.420 | 3.375 | 3.229 | 6.662 | 6.398 | 0.531 |
| c1_power2 | pair_429 | 0.006 | 1856.420 | 6.335 | 6.811 | 13.290 | 13.432 | 0.475 |
| c1_power2 | pair_430 | 0.006 | 1856.420 | 6.931 | 6.773 | 13.935 | 13.764 | 1.302 |
| c1_power2 | pair_431 | 0.006 | 1856.420 | 3.227 | 3.291 | 6.587 | 6.570 | 0.145 |
| c1_power2 | pair_432 | 0.006 | 1856.420 | 6.494 | 6.538 | 13.260 | 12.811 | 0.822 |
| c1_power2 | pair_433 | 0.006 | 1856.420 | 3.338 | 3.149 | 6.560 | 6.354 | 0.427 |
| c1_power2 | pair_434 | 0.006 | 1856.420 | 6.352 | 6.558 | 13.198 | 12.857 | 0.662 |
| c1_power2 | pair_435 | 0.006 | 1856.420 | 3.151 | 3.220 | 6.431 | 6.405 | 0.228 |
| c1_sqrt | pair_001 | 0.004 | 1923.541 | 10.405 | 8.782 | 19.539 | 18.684 | 2.196 |
| c1_sqrt | pair_002 | 0.004 | 1923.541 | 7.495 | 7.451 | 15.291 | 14.949 | 1.331 |
| c1_sqrt | pair_003 | 0.004 | 1923.541 | 8.165 | 8.081 | 16.599 | 16.278 | 0.647 |
| c1_sqrt | pair_004 | 0.004 | 1923.541 | 8.830 | 8.775 | 17.806 | 17.803 | 0.546 |
| c1_sqrt | pair_005 | 0.004 | 1923.541 | 4.019 | 4.010 | 8.142 | 8.149 | 0.137 |
| c1_sqrt | pair_006 | 0.004 | 1923.541 | 7.118 | 7.198 | 14.556 | 14.418 | 0.496 |
| c1_sqrt | pair_007 | 0.004 | 1923.541 | 17.670 | 16.379 | 34.642 | 34.090 | 2.130 |
| c1_sqrt | pair_008 | 0.004 | 1923.541 | 7.874 | 7.830 | 15.936 | 15.888 | 0.141 |
| c1_sqrt | pair_009 | 0.004 | 1923.541 | 22.710 | 22.176 | 45.975 | 41.930 | 7.221 |
| c1_sqrt | pair_010 | 0.004 | 1923.541 | 8.177 | 7.944 | 16.921 | 17.129 | 0.459 |
| c1_sqrt | pair_011 | 0.004 | 1923.541 | 6.680 | 6.524 | 13.548 | 13.566 | 0.652 |
| c1_sqrt | pair_012 | 0.004 | 1923.541 | 7.829 | 8.414 | 16.864 | 15.732 | 3.200 |
| c1_sqrt | pair_013 | 0.004 | 1923.541 | 9.717 | 8.822 | 18.983 | 17.340 | 3.984 |
| c1_sqrt | pair_014 | 0.004 | 1923.541 | 8.153 | 8.036 | 16.485 | 16.560 | 1.504 |
| c1_sqrt | pair_015 | 0.004 | 1923.541 | 9.336 | 8.013 | 17.804 | 16.858 | 3.766 |
| c1_sqrt | pair_016 | 0.004 | 1923.541 | 4.455 | 4.438 | 9.044 | 8.958 | 0.999 |
| c1_sqrt | pair_017 | 0.004 | 1923.541 | 4.100 | 4.115 | 8.319 | 8.014 | 0.560 |
| c1_sqrt | pair_018 | 0.004 | 1923.541 | 12.297 | 12.221 | 24.969 | 24.345 | 1.339 |
| c1_sqrt | pair_019 | 0.004 | 1923.541 | 8.138 | 8.270 | 16.678 | 15.920 | 1.314 |
| c1_sqrt | pair_020 | 0.004 | 1923.541 | 21.181 | 21.171 | 43.205 | 41.682 | 2.990 |
| c1_sqrt | pair_021 | 0.004 | 1923.541 | 8.378 | 8.059 | 16.685 | 16.475 | 0.749 |
| c1_sqrt | pair_022 | 0.004 | 1923.541 | 21.176 | 21.339 | 43.609 | 42.963 | 4.137 |
| c1_sqrt | pair_023 | 0.004 | 1923.541 | 6.651 | 6.903 | 13.868 | 14.323 | 1.082 |
| c1_sqrt | pair_024 | 0.004 | 1923.541 | 14.002 | 13.802 | 28.523 | 29.350 | 2.727 |
| c1_sqrt | pair_025 | 0.004 | 1923.541 | 22.864 | 22.009 | 45.943 | 45.580 | 2.270 |
| c1_sqrt | pair_026 | 0.004 | 1923.541 | 6.762 | 6.488 | 13.489 | 13.702 | 0.579 |
| c1_sqrt | pair_027 | 0.004 | 1923.541 | 6.784 | 6.508 | 13.489 | 13.260 | 0.709 |
| c1_sqrt | pair_028 | 0.004 | 1923.541 | 4.010 | 4.101 | 8.298 | 8.138 | 0.297 |
| c1_sqrt | pair_029 | 0.004 | 1923.541 | 7.274 | 7.038 | 14.582 | 14.607 | 0.134 |
| c1_sqrt | pair_030 | 0.004 | 1923.541 | 7.148 | 7.491 | 14.906 | 14.437 | 0.822 |
| c1_sqrt | pair_031 | 0.004 | 1923.541 | 8.555 | 8.690 | 17.633 | 18.133 | 1.205 |
| c1_sqrt | pair_032 | 0.004 | 1923.541 | 9.098 | 8.774 | 18.189 | 17.952 | 0.484 |
| c1_sqrt | pair_033 | 0.004 | 1923.541 | 4.229 | 4.147 | 8.489 | 8.119 | 0.721 |
| c1_sqrt | pair_034 | 0.004 | 1923.541 | 7.324 | 7.190 | 14.734 | 14.549 | 0.400 |
| c1_sqrt | pair_035 | 0.004 | 1923.541 | 8.791 | 9.014 | 18.068 | 18.186 | 0.447 |
| c1_sqrt | pair_036 | 0.004 | 1923.541 | 7.989 | 7.847 | 16.034 | 16.065 | 0.126 |
| c1_sqrt | pair_037 | 0.004 | 1923.541 | 8.899 | 8.773 | 17.982 | 17.768 | 0.518 |
| c1_sqrt | pair_038 | 0.004 | 1923.541 | 7.997 | 7.983 | 16.157 | 16.182 | 0.219 |
| c1_sqrt | pair_039 | 0.004 | 1923.541 | 8.137 | 6.292 | 14.707 | 13.347 | 2.887 |
| c1_sqrt | pair_040 | 0.004 | 1923.541 | 6.343 | 6.272 | 12.813 | 12.835 | 0.077 |
| c1_sqrt | pair_041 | 0.004 | 1923.541 | 7.199 | 7.255 | 14.729 | 14.810 | 0.189 |
| c1_sqrt | pair_042 | 0.004 | 1923.541 | 7.542 | 7.405 | 15.203 | 15.378 | 0.797 |
| c1_sqrt | pair_043 | 0.004 | 1923.541 | 7.291 | 7.211 | 14.877 | 14.745 | 0.494 |
| c1_sqrt | pair_044 | 0.004 | 1923.541 | 4.145 | 4.321 | 8.622 | 8.080 | 1.062 |
| c1_sqrt | pair_045 | 0.004 | 1923.541 | 4.868 | 4.764 | 9.791 | 8.072 | 3.068 |
| c1_sqrt | pair_046 | 0.004 | 1923.541 | 8.749 | 8.586 | 17.601 | 17.461 | 0.281 |
| c1_sqrt | pair_047 | 0.004 | 1923.541 | 7.886 | 7.840 | 15.984 | 15.906 | 0.179 |
| c1_sqrt | pair_048 | 0.004 | 1923.541 | 8.986 | 8.830 | 18.074 | 17.573 | 0.942 |
| c1_sqrt | pair_049 | 0.004 | 1923.541 | 8.417 | 8.060 | 16.689 | 15.967 | 1.313 |
| c1_sqrt | pair_050 | 0.004 | 1923.541 | 8.819 | 8.693 | 17.769 | 17.748 | 0.291 |
| c1_sqrt | pair_051 | 0.004 | 1923.541 | 6.279 | 6.330 | 12.812 | 12.810 | 0.080 |
| c1_sqrt | pair_052 | 0.004 | 1923.541 | 9.548 | 10.036 | 19.919 | 17.632 | 4.041 |
| c1_sqrt | pair_053 | 0.004 | 1923.541 | 9.439 | 8.850 | 18.546 | 17.497 | 1.916 |
| c1_sqrt | pair_054 | 0.004 | 1923.541 | 6.580 | 6.698 | 13.465 | 12.848 | 1.193 |
| c1_sqrt | pair_055 | 0.004 | 1923.541 | 7.579 | 7.811 | 15.650 | 12.872 | 4.912 |
| c1_sqrt | pair_056 | 0.004 | 1923.541 | 6.771 | 5.302 | 12.686 | 9.110 | 7.107 |
| c1_sqrt | pair_057 | 0.004 | 1923.541 | 8.696 | 7.450 | 16.366 | 15.004 | 2.913 |
| c1_sqrt | pair_058 | 0.004 | 1923.541 | 7.525 | 7.616 | 15.464 | 14.453 | 1.887 |
| c1_sqrt | pair_059 | 0.004 | 1923.541 | 9.511 | 9.711 | 19.580 | 14.477 | 8.900 |
| c1_sqrt | pair_060 | 0.004 | 1923.541 | 4.315 | 4.224 | 8.705 | 7.970 | 1.323 |
| c1_sqrt | pair_061 | 0.004 | 1923.541 | 9.202 | 8.452 | 17.925 | 14.917 | 5.565 |
| c1_sqrt | pair_062 | 0.004 | 1923.541 | 7.392 | 7.289 | 14.883 | 14.451 | 0.892 |
| c1_sqrt | pair_063 | 0.004 | 1923.541 | 8.720 | 7.901 | 16.930 | 14.414 | 4.413 |
| c1_sqrt | pair_064 | 0.004 | 1923.541 | 7.992 | 7.706 | 15.945 | 14.407 | 2.853 |
| c1_sqrt | pair_065 | 0.004 | 1923.541 | 8.286 | 7.508 | 16.046 | 14.819 | 2.678 |
| c1_sqrt | pair_066 | 0.004 | 1923.541 | 6.745 | 6.654 | 13.640 | 13.202 | 1.153 |
| c1_sqrt | pair_067 | 0.004 | 1923.541 | 6.694 | 6.621 | 13.545 | 13.450 | 0.960 |
| c1_sqrt | pair_068 | 0.004 | 1923.541 | 7.676 | 8.146 | 16.042 | 14.528 | 2.786 |
| c1_sqrt | pair_069 | 0.004 | 1923.541 | 8.340 | 7.301 | 15.825 | 16.345 | 1.405 |
| c1_sqrt | pair_070 | 0.004 | 1923.541 | 7.402 | 7.650 | 15.300 | 14.411 | 1.640 |
| c1_sqrt | pair_071 | 0.004 | 1923.541 | 4.464 | 4.092 | 8.684 | 8.004 | 1.286 |
| c1_sqrt | pair_072 | 0.004 | 1923.541 | 4.150 | 4.515 | 8.811 | 8.031 | 1.557 |
| c1_sqrt | pair_073 | 0.004 | 1923.541 | 8.610 | 8.012 | 16.872 | 14.487 | 4.284 |
| c1_sqrt | pair_074 | 0.004 | 1923.541 | 7.406 | 8.116 | 188.059 | 19.034 | 296.857 |
| c1_sqrt | pair_075 | 0.004 | 1923.541 | 9.000 | 8.511 | 17.800 | 16.739 | 4.011 |
| c1_sqrt | pair_076 | 0.004 | 1923.541 | 8.442 | 9.324 | 18.179 | 18.423 | 0.642 |
| c1_sqrt | pair_077 | 0.004 | 1923.541 | 7.725 | 7.504 | 15.476 | 15.106 | 1.606 |
| c1_sqrt | pair_078 | 0.004 | 1923.541 | 7.102 | 7.279 | 14.690 | 12.978 | 3.148 |
| c1_sqrt | pair_079 | 0.004 | 1923.541 | 7.894 | 8.299 | 16.517 | 14.397 | 3.764 |
| c1_sqrt | pair_080 | 0.004 | 1923.541 | 7.546 | 7.295 | 15.085 | 14.561 | 1.073 |
| c1_sqrt | pair_081 | 0.004 | 1923.541 | 6.881 | 6.843 | 13.995 | 12.868 | 2.036 |
| c1_sqrt | pair_082 | 0.004 | 1923.541 | 7.386 | 7.059 | 14.712 | 12.909 | 3.393 |
| c1_sqrt | pair_083 | 0.004 | 1923.541 | 4.492 | 4.474 | 9.294 | 8.514 | 1.752 |
| c1_sqrt | pair_084 | 0.004 | 1923.541 | 7.520 | 7.712 | 15.445 | 15.645 | 0.730 |
| c1_sqrt | pair_085 | 0.004 | 1923.541 | 8.469 | 8.966 | 17.785 | 18.649 | 1.613 |
| c1_sqrt | pair_086 | 0.004 | 1923.541 | 3.991 | 3.932 | 8.092 | 8.127 | 0.080 |
| c1_sqrt | pair_087 | 0.004 | 1923.541 | 7.310 | 7.273 | 15.062 | 14.551 | 0.938 |
| c1_sqrt | pair_088 | 0.004 | 1923.541 | 10.540 | 9.003 | 19.786 | 16.322 | 6.321 |
| c1_sqrt | pair_089 | 0.004 | 1923.541 | 8.493 | 8.538 | 17.375 | 16.088 | 2.395 |
| c1_sqrt | pair_090 | 0.004 | 1923.541 | 8.413 | 8.535 | 17.316 | 16.087 | 2.223 |
| c1_sqrt | pair_091 | 0.004 | 1923.541 | 8.090 | 8.197 | 16.587 | 16.463 | 0.522 |
| c1_sqrt | pair_092 | 0.004 | 1923.541 | 6.478 | 6.298 | 13.028 | 13.097 | 0.183 |
| c1_sqrt | pair_093 | 0.004 | 1923.541 | 6.312 | 6.496 | 13.029 | 12.934 | 0.496 |
| c1_sqrt | pair_094 | 0.004 | 1923.541 | 7.142 | 7.276 | 14.726 | 14.566 | 0.316 |
| c1_sqrt | pair_095 | 0.004 | 1923.541 | 8.334 | 7.288 | 15.872 | 14.557 | 2.368 |
| c1_sqrt | pair_096 | 0.004 | 1923.541 | 8.266 | 8.391 | 16.990 | 14.564 | 4.370 |
| c1_sqrt | pair_097 | 0.004 | 1923.541 | 4.675 | 4.577 | 9.423 | 8.231 | 2.262 |
| c1_sqrt | pair_098 | 0.004 | 1923.541 | 4.528 | 5.453 | 10.430 | 8.013 | 4.229 |
| c1_sqrt | pair_099 | 0.004 | 1923.541 | 10.845 | 9.474 | 20.598 | 16.407 | 7.551 |
| c1_sqrt | pair_100 | 0.004 | 1923.541 | 8.738 | 8.170 | 17.126 | 16.132 | 1.912 |
| c1_sqrt | pair_101 | 0.004 | 1923.541 | 8.171 | 8.104 | 16.520 | 16.070 | 0.921 |
| c1_sqrt | pair_102 | 0.004 | 1923.541 | 8.045 | 7.921 | 16.252 | 16.109 | 0.278 |
| c1_sqrt | pair_103 | 0.004 | 1923.541 | 8.137 | 8.143 | 16.507 | 16.105 | 0.851 |
| c1_sqrt | pair_104 | 0.004 | 1923.541 | 6.584 | 6.521 | 13.319 | 12.886 | 0.790 |
| c1_sqrt | pair_105 | 0.004 | 1923.541 | 8.247 | 8.059 | 16.512 | 16.103 | 0.743 |
| c1_sqrt | pair_106 | 0.004 | 1923.541 | 8.056 | 8.062 | 16.370 | 16.082 | 0.622 |
| c1_sqrt | pair_107 | 0.004 | 1923.541 | 6.420 | 6.284 | 12.909 | 12.834 | 0.174 |
| c1_sqrt | pair_108 | 0.004 | 1923.541 | 6.437 | 6.410 | 13.042 | 13.053 | 0.035 |
| c1_sqrt | pair_109 | 0.004 | 1923.541 | 3.989 | 3.942 | 8.041 | 8.072 | 0.098 |
| c1_sqrt | pair_110 | 0.004 | 1923.541 | 7.692 | 7.231 | 15.266 | 14.573 | 1.420 |
| c1_sqrt | pair_111 | 0.004 | 1923.541 | 3.939 | 4.404 | 8.441 | 8.057 | 0.798 |
| c1_sqrt | pair_112 | 0.004 | 1923.541 | 7.326 | 7.228 | 15.193 | 15.090 | 0.452 |
| c1_sqrt | pair_113 | 0.004 | 1923.541 | 8.688 | 8.772 | 17.701 | 17.615 | 0.187 |
| c1_sqrt | pair_114 | 0.004 | 1923.541 | 7.959 | 7.855 | 16.109 | 16.101 | 0.080 |
| c1_sqrt | pair_115 | 0.004 | 1923.541 | 8.946 | 8.820 | 18.020 | 18.058 | 0.208 |
| c1_sqrt | pair_116 | 0.004 | 1923.541 | 8.121 | 8.195 | 16.554 | 16.383 | 0.559 |
| c1_sqrt | pair_117 | 0.004 | 1923.541 | 7.081 | 6.673 | 13.981 | 14.226 | 1.041 |
| c1_sqrt | pair_118 | 0.004 | 1923.541 | 7.247 | 6.601 | 14.076 | 14.243 | 1.212 |
| c1_sqrt | pair_119 | 0.004 | 1923.541 | 7.219 | 7.128 | 14.560 | 14.436 | 0.260 |
| c1_sqrt | pair_120 | 0.004 | 1923.541 | 7.214 | 7.131 | 14.566 | 14.461 | 0.185 |
| c1_sqrt | pair_121 | 0.004 | 1923.541 | 7.102 | 7.100 | 14.429 | 14.468 | 0.095 |
| c1_sqrt | pair_122 | 0.004 | 1923.541 | 3.969 | 3.953 | 8.041 | 8.066 | 0.059 |
| c1_sqrt | pair_123 | 0.004 | 1923.541 | 3.989 | 3.902 | 8.030 | 8.018 | 0.027 |
| c1_sqrt | pair_124 | 0.004 | 1923.541 | 8.731 | 8.688 | 17.726 | 17.884 | 0.285 |
| c1_sqrt | pair_125 | 0.004 | 1923.541 | 7.921 | 7.832 | 16.031 | 16.064 | 0.143 |
| c1_sqrt | pair_126 | 0.004 | 1923.541 | 8.961 | 8.800 | 17.983 | 17.675 | 0.590 |
| c1_sqrt | pair_127 | 0.004 | 1923.541 | 7.999 | 7.870 | 16.103 | 16.121 | 0.201 |
| c1_sqrt | pair_128 | 0.004 | 1923.541 | 8.723 | 9.973 | 18.954 | 17.740 | 2.199 |
| c1_sqrt | pair_129 | 0.004 | 1923.541 | 6.524 | 6.290 | 13.075 | 13.121 | 0.144 |
| c1_sqrt | pair_130 | 0.004 | 1923.541 | 8.792 | 8.697 | 17.754 | 17.629 | 0.229 |
| c1_sqrt | pair_131 | 0.004 | 1923.541 | 9.064 | 8.795 | 18.130 | 17.525 | 1.102 |
| c1_sqrt | pair_132 | 0.004 | 1923.541 | 6.346 | 6.485 | 13.003 | 12.897 | 0.357 |
| c1_sqrt | pair_133 | 0.004 | 1923.541 | 6.375 | 6.438 | 13.360 | 13.018 | 0.858 |
| c1_sqrt | pair_134 | 0.004 | 1923.541 | 4.056 | 3.997 | 8.215 | 8.128 | 0.220 |
| c1_sqrt | pair_135 | 0.004 | 1923.541 | 7.145 | 7.089 | 14.470 | 14.431 | 0.213 |
| c1_sqrt | pair_136 | 0.004 | 1923.541 | 3.952 | 3.961 | 8.004 | 7.979 | 0.138 |
| c1_sqrt | pair_137 | 0.004 | 1923.541 | 4.031 | 4.003 | 8.150 | 8.108 | 0.244 |
| c1_sqrt | pair_138 | 0.004 | 1923.541 | 4.007 | 3.897 | 8.062 | 8.083 | 0.152 |
| c1_sqrt | pair_139 | 0.004 | 1923.541 | 4.143 | 3.951 | 8.238 | 8.120 | 0.246 |
| c1_sqrt | pair_140 | 0.004 | 1923.541 | 4.746 | 4.344 | 9.241 | 8.195 | 1.865 |
| c1_sqrt | pair_141 | 0.004 | 1923.541 | 4.674 | 4.348 | 9.149 | 8.264 | 1.884 |
| c1_sqrt | pair_142 | 0.004 | 1923.541 | 4.914 | 4.363 | 9.445 | 8.477 | 1.989 |
| c1_sqrt | pair_143 | 0.004 | 1923.541 | 4.417 | 4.622 | 9.170 | 8.592 | 1.514 |
| c1_sqrt | pair_144 | 0.004 | 1923.541 | 4.427 | 4.694 | 9.201 | 8.259 | 1.897 |
| c1_sqrt | pair_145 | 0.004 | 1923.541 | 4.373 | 4.814 | 9.275 | 8.044 | 2.239 |
| c1_sqrt | pair_146 | 0.004 | 1923.541 | 4.253 | 4.130 | 8.529 | 8.085 | 0.875 |
| c1_sqrt | pair_147 | 0.004 | 1923.541 | 4.403 | 4.527 | 9.061 | 8.373 | 1.569 |
| c1_sqrt | pair_148 | 0.004 | 1923.541 | 4.565 | 4.504 | 9.262 | 8.649 | 1.599 |
| c1_sqrt | pair_149 | 0.004 | 1923.541 | 4.747 | 4.330 | 9.378 | 8.912 | 1.143 |
| c1_sqrt | pair_150 | 0.004 | 1923.541 | 4.938 | 4.855 | 10.231 | 10.177 | 1.092 |
| c1_sqrt | pair_151 | 0.004 | 1923.541 | 4.621 | 4.142 | 8.954 | 8.980 | 0.787 |
| c1_sqrt | pair_152 | 0.004 | 1923.541 | 4.027 | 4.149 | 8.400 | 8.520 | 0.295 |
| c1_sqrt | pair_153 | 0.004 | 1923.541 | 4.527 | 4.249 | 8.906 | 8.351 | 1.165 |
| c1_sqrt | pair_154 | 0.004 | 1923.541 | 4.105 | 4.044 | 8.277 | 8.381 | 0.391 |
| c1_sqrt | pair_155 | 0.004 | 1923.541 | 4.185 | 4.048 | 8.361 | 8.422 | 0.151 |
| c1_sqrt | pair_156 | 0.004 | 1923.541 | 4.114 | 4.007 | 8.235 | 8.015 | 0.407 |
| c1_sqrt | pair_157 | 0.004 | 1923.541 | 4.265 | 4.036 | 8.387 | 8.054 | 0.771 |
| c1_sqrt | pair_158 | 0.004 | 1923.541 | 4.578 | 4.038 | 8.733 | 7.982 | 1.419 |
| c1_sqrt | pair_159 | 0.004 | 1923.541 | 4.052 | 4.150 | 8.381 | 8.269 | 0.490 |
| c1_sqrt | pair_160 | 0.004 | 1923.541 | 7.553 | 7.443 | 15.251 | 14.907 | 1.020 |
| c1_sqrt | pair_161 | 0.004 | 1923.541 | 7.852 | 7.592 | 15.708 | 14.694 | 2.124 |
| c1_sqrt | pair_162 | 0.004 | 1923.541 | 8.044 | 7.836 | 16.206 | 14.642 | 3.039 |
| c1_sqrt | pair_163 | 0.004 | 1923.541 | 7.651 | 7.965 | 15.985 | 14.741 | 2.440 |
| c1_sqrt | pair_164 | 0.004 | 1923.541 | 6.822 | 6.791 | 13.827 | 13.044 | 1.520 |
| c1_sqrt | pair_165 | 0.004 | 1923.541 | 6.929 | 6.367 | 13.503 | 13.043 | 1.035 |
| c1_sqrt | pair_166 | 0.004 | 1923.541 | 7.615 | 7.537 | 15.482 | 14.668 | 1.595 |
| c1_sqrt | pair_167 | 0.004 | 1923.541 | 7.319 | 7.254 | 15.482 | 15.264 | 1.285 |
| c1_sqrt | pair_168 | 0.004 | 1923.541 | 8.116 | 7.887 | 16.350 | 16.432 | 1.803 |
| c1_sqrt | pair_169 | 0.004 | 1923.541 | 4.287 | 4.341 | 8.732 | 8.311 | 1.094 |
| c1_sqrt | pair_170 | 0.004 | 1923.541 | 4.489 | 4.316 | 9.071 | 9.312 | 0.967 |
| c1_sqrt | pair_171 | 0.004 | 1923.541 | 7.395 | 7.559 | 15.247 | 14.857 | 1.170 |
| c1_sqrt | pair_172 | 0.004 | 1923.541 | 9.159 | 8.543 | 17.945 | 14.326 | 6.354 |
| c1_sqrt | pair_173 | 0.004 | 1923.541 | 7.832 | 7.680 | 15.823 | 15.655 | 0.762 |
| c1_sqrt | pair_174 | 0.004 | 1923.541 | 7.497 | 7.189 | 14.876 | 14.825 | 0.428 |
| c1_sqrt | pair_175 | 0.004 | 1923.541 | 7.377 | 7.257 | 14.884 | 14.712 | 0.536 |
| c1_sqrt | pair_176 | 0.004 | 1923.541 | 6.699 | 6.534 | 13.453 | 13.636 | 0.618 |
| c1_sqrt | pair_177 | 0.004 | 1923.541 | 7.389 | 7.416 | 15.201 | 14.984 | 1.010 |
| c1_sqrt | pair_178 | 0.004 | 1923.541 | 7.908 | 7.308 | 15.479 | 15.319 | 0.947 |
| c1_sqrt | pair_179 | 0.004 | 1923.541 | 6.820 | 6.861 | 13.904 | 14.431 | 1.026 |
| c1_sqrt | pair_180 | 0.004 | 1923.541 | 6.819 | 7.372 | 14.543 | 15.317 | 1.470 |
| c1_sqrt | pair_181 | 0.004 | 1923.541 | 4.060 | 4.028 | 8.245 | 8.124 | 0.343 |
| c1_sqrt | pair_182 | 0.004 | 1923.541 | 7.073 | 7.111 | 14.475 | 14.455 | 0.169 |
| c1_sqrt | pair_183 | 0.004 | 1923.541 | 7.921 | 7.918 | 16.081 | 16.124 | 0.172 |
| c1_sqrt | pair_184 | 0.004 | 1923.541 | 15.677 | 15.814 | 32.028 | 31.735 | 0.533 |
| c1_sqrt | pair_185 | 0.004 | 1923.541 | 8.260 | 7.928 | 16.490 | 16.112 | 0.792 |
| c1_sqrt | pair_186 | 0.004 | 1923.541 | 6.366 | 6.283 | 12.877 | 12.807 | 0.171 |
| c1_sqrt | pair_187 | 0.004 | 1923.541 | 6.339 | 6.242 | 12.865 | 12.905 | 0.097 |
| c1_sqrt | pair_188 | 0.004 | 1923.541 | 7.026 | 6.986 | 14.260 | 14.315 | 0.100 |
| c1_sqrt | pair_189 | 0.004 | 1923.541 | 7.043 | 7.459 | 14.729 | 14.278 | 0.833 |
| c1_sqrt | pair_190 | 0.004 | 1923.541 | 7.151 | 7.037 | 14.438 | 14.389 | 0.135 |
| c1_sqrt | pair_191 | 0.004 | 1923.541 | 3.920 | 3.899 | 7.953 | 7.964 | 0.089 |
| c1_sqrt | pair_192 | 0.004 | 1923.541 | 3.907 | 3.872 | 7.887 | 7.903 | 0.041 |
| c1_sqrt | pair_193 | 0.004 | 1923.541 | 11.685 | 11.694 | 23.796 | 23.812 | 0.077 |
| c1_sqrt | pair_194 | 0.004 | 1923.541 | 8.058 | 7.990 | 16.302 | 15.837 | 0.813 |
| c1_sqrt | pair_195 | 0.004 | 1923.541 | 16.024 | 16.323 | 32.890 | 31.854 | 1.993 |
| c1_sqrt | pair_196 | 0.004 | 1923.541 | 8.077 | 8.025 | 16.375 | 15.778 | 1.073 |
| c1_sqrt | pair_197 | 0.004 | 1923.541 | 16.243 | 16.170 | 32.991 | 32.055 | 1.842 |
| c1_sqrt | pair_198 | 0.004 | 1923.541 | 6.695 | 6.466 | 13.343 | 12.853 | 0.890 |
| c1_sqrt | pair_199 | 0.004 | 1923.541 | 13.229 | 13.028 | 26.716 | 25.975 | 1.624 |
| c1_sqrt | pair_200 | 0.004 | 1923.541 | 16.901 | 16.262 | 33.747 | 33.826 | 1.067 |
| c1_sqrt | pair_201 | 0.004 | 1923.541 | 6.615 | 6.425 | 13.342 | 13.456 | 0.641 |
| c1_sqrt | pair_202 | 0.004 | 1923.541 | 6.448 | 7.073 | 13.775 | 12.871 | 1.700 |
| c1_sqrt | pair_203 | 0.004 | 1923.541 | 4.203 | 4.246 | 8.622 | 8.045 | 1.065 |
| c1_sqrt | pair_204 | 0.004 | 1923.541 | 7.330 | 7.231 | 14.769 | 14.307 | 0.883 |
| c1_sqrt | pair_205 | 0.004 | 1923.541 | 8.258 | 8.315 | 16.916 | 17.071 | 0.809 |
| c1_sqrt | pair_206 | 0.004 | 1923.541 | 8.263 | 8.025 | 16.657 | 16.254 | 1.078 |
| c1_sqrt | pair_207 | 0.004 | 1923.541 | 6.508 | 7.067 | 13.787 | 13.847 | 1.158 |
| c1_sqrt | pair_208 | 0.004 | 1923.541 | 6.906 | 6.548 | 13.688 | 14.113 | 0.779 |
| c1_sqrt | pair_209 | 0.004 | 1923.541 | 7.404 | 7.502 | 15.137 | 14.561 | 1.333 |
| c1_sqrt | pair_210 | 0.004 | 1923.541 | 7.229 | 7.404 | 14.957 | 14.328 | 1.096 |
| c1_sqrt | pair_211 | 0.004 | 1923.541 | 7.697 | 7.341 | 15.252 | 14.572 | 1.439 |
| c1_sqrt | pair_212 | 0.004 | 1923.541 | 4.014 | 3.927 | 8.049 | 7.966 | 0.196 |
| c1_sqrt | pair_213 | 0.004 | 1923.541 | 4.126 | 3.929 | 8.189 | 7.911 | 0.508 |
| c1_sqrt | pair_214 | 0.004 | 1923.541 | 7.894 | 7.812 | 15.955 | 15.823 | 0.263 |
| c1_sqrt | pair_215 | 0.004 | 1923.541 | 8.042 | 8.962 | 17.251 | 15.931 | 2.353 |
| c1_sqrt | pair_216 | 0.004 | 1923.541 | 7.842 | 7.815 | 15.934 | 15.877 | 0.106 |
| c1_sqrt | pair_217 | 0.004 | 1923.541 | 8.115 | 8.021 | 16.368 | 16.023 | 0.819 |
| c1_sqrt | pair_218 | 0.004 | 1923.541 | 8.105 | 8.046 | 16.424 | 16.013 | 0.882 |
| c1_sqrt | pair_219 | 0.004 | 1923.541 | 6.288 | 6.300 | 12.797 | 12.681 | 0.229 |
| c1_sqrt | pair_220 | 0.004 | 1923.541 | 7.842 | 7.752 | 15.799 | 15.797 | 0.113 |
| c1_sqrt | pair_221 | 0.004 | 1923.541 | 7.797 | 7.796 | 15.846 | 15.802 | 0.118 |
| c1_sqrt | pair_222 | 0.004 | 1923.541 | 9.686 | 6.538 | 16.469 | 13.380 | 6.081 |
| c1_sqrt | pair_223 | 0.004 | 1923.541 | 6.366 | 6.270 | 12.858 | 12.855 | 0.061 |
| c1_sqrt | pair_224 | 0.004 | 1923.541 | 3.971 | 3.890 | 8.002 | 8.019 | 0.043 |
| c1_sqrt | pair_225 | 0.004 | 1923.541 | 7.075 | 7.017 | 14.311 | 14.303 | 0.117 |
| c1_sqrt | pair_226 | 0.004 | 1923.541 | 7.835 | 7.809 | 15.856 | 15.805 | 0.112 |
| c1_sqrt | pair_227 | 0.004 | 1923.541 | 6.268 | 6.261 | 12.708 | 12.728 | 0.061 |
| c1_sqrt | pair_228 | 0.004 | 1923.541 | 6.313 | 6.244 | 12.791 | 12.827 | 0.153 |
| c1_sqrt | pair_229 | 0.004 | 1923.541 | 7.048 | 7.010 | 14.253 | 14.229 | 0.076 |
| c1_sqrt | pair_230 | 0.004 | 1923.541 | 7.057 | 6.995 | 14.288 | 14.331 | 0.115 |
| c1_sqrt | pair_231 | 0.004 | 1923.541 | 7.059 | 7.012 | 14.558 | 14.395 | 0.332 |
| c1_sqrt | pair_232 | 0.004 | 1923.541 | 4.151 | 4.220 | 8.489 | 7.962 | 0.958 |
| c1_sqrt | pair_233 | 0.004 | 1923.541 | 4.214 | 4.398 | 8.746 | 7.935 | 1.422 |
| c1_sqrt | pair_234 | 0.004 | 1923.541 | 12.323 | 12.756 | 25.494 | 25.352 | 1.713 |
| c1_sqrt | pair_235 | 0.004 | 1923.541 | 8.335 | 8.244 | 16.898 | 16.082 | 1.687 |
| c1_sqrt | pair_236 | 0.004 | 1923.541 | 23.519 | 23.749 | 48.259 | 48.036 | 4.385 |
| c1_sqrt | pair_237 | 0.004 | 1923.541 | 8.216 | 8.220 | 16.773 | 16.630 | 0.894 |
| c1_sqrt | pair_238 | 0.004 | 1923.541 | 19.819 | 20.041 | 40.558 | 40.339 | 1.046 |
| c1_sqrt | pair_239 | 0.004 | 1923.541 | 6.357 | 6.553 | 13.067 | 12.953 | 0.490 |
| c1_sqrt | pair_240 | 0.004 | 1923.541 | 12.990 | 12.597 | 26.029 | 25.675 | 0.711 |
| c1_sqrt | pair_241 | 0.004 | 1923.541 | 20.435 | 20.359 | 43.230 | 41.648 | 2.906 |
| c1_sqrt | pair_242 | 0.004 | 1923.541 | 8.415 | 6.869 | 15.823 | 12.987 | 5.184 |
| c1_sqrt | pair_243 | 0.004 | 1923.541 | 6.337 | 6.255 | 12.757 | 12.738 | 0.071 |
| c1_sqrt | pair_244 | 0.004 | 1923.541 | 3.967 | 3.923 | 8.067 | 8.049 | 0.098 |
| c1_sqrt | pair_245 | 0.004 | 1923.541 | 7.074 | 7.029 | 14.360 | 14.375 | 0.038 |
| c1_sqrt | pair_246 | 0.004 | 1923.541 | 6.954 | 6.617 | 13.823 | 13.315 | 1.170 |
| c1_sqrt | pair_247 | 0.004 | 1923.541 | 6.511 | 6.413 | 13.122 | 12.932 | 0.394 |
| c1_sqrt | pair_248 | 0.004 | 1923.541 | 7.202 | 7.116 | 14.569 | 14.583 | 0.257 |
| c1_sqrt | pair_249 | 0.004 | 1923.541 | 7.140 | 7.028 | 14.779 | 14.731 | 0.487 |
| c1_sqrt | pair_250 | 0.004 | 1923.541 | 7.865 | 8.091 | 16.145 | 14.574 | 3.012 |
| c1_sqrt | pair_251 | 0.004 | 1923.541 | 4.844 | 4.298 | 9.291 | 9.852 | 1.120 |
| c1_sqrt | pair_252 | 0.004 | 1923.541 | 4.422 | 4.386 | 9.000 | 8.496 | 1.258 |
| c1_sqrt | pair_253 | 0.004 | 1923.541 | 8.463 | 8.457 | 17.309 | 16.102 | 2.142 |
| c1_sqrt | pair_254 | 0.004 | 1923.541 | 8.621 | 8.740 | 17.682 | 17.093 | 2.008 |
| c1_sqrt | pair_255 | 0.004 | 1923.541 | 8.015 | 7.957 | 16.163 | 16.150 | 0.290 |
| c1_sqrt | pair_256 | 0.004 | 1923.541 | 7.894 | 7.860 | 16.000 | 15.980 | 0.087 |
| c1_sqrt | pair_257 | 0.004 | 1923.541 | 7.969 | 8.175 | 16.429 | 16.247 | 0.709 |
| c1_sqrt | pair_258 | 0.004 | 1923.541 | 6.334 | 6.250 | 12.822 | 12.807 | 0.079 |
| c1_sqrt | pair_259 | 0.004 | 1923.541 | 7.890 | 7.786 | 15.873 | 15.808 | 0.235 |
| c1_sqrt | pair_260 | 0.004 | 1923.541 | 7.921 | 7.921 | 16.064 | 16.086 | 0.144 |
| c1_sqrt | pair_261 | 0.004 | 1923.541 | 6.314 | 6.294 | 12.829 | 12.768 | 0.156 |
| c1_sqrt | pair_262 | 0.004 | 1923.541 | 6.279 | 6.281 | 12.785 | 12.794 | 0.024 |
| c1_sqrt | pair_263 | 0.004 | 1923.541 | 3.980 | 3.913 | 7.988 | 8.024 | 0.100 |
| c1_sqrt | pair_264 | 0.004 | 1923.541 | 7.045 | 7.026 | 14.300 | 14.327 | 0.049 |
| c1_sqrt | pair_265 | 0.004 | 1923.541 | 6.280 | 6.293 | 12.784 | 12.759 | 0.097 |
| c1_sqrt | pair_266 | 0.004 | 1923.541 | 6.283 | 6.212 | 12.752 | 12.681 | 0.138 |
| c1_sqrt | pair_267 | 0.004 | 1923.541 | 6.596 | 6.240 | 13.097 | 12.801 | 0.534 |
| c1_sqrt | pair_268 | 0.004 | 1923.541 | 6.334 | 6.233 | 12.785 | 12.832 | 0.134 |
| c1_sqrt | pair_269 | 0.004 | 1923.541 | 4.174 | 3.905 | 8.173 | 8.114 | 0.256 |
| c1_sqrt | pair_270 | 0.004 | 1923.541 | 3.992 | 3.886 | 8.028 | 8.023 | 0.073 |
| c1_sqrt | pair_271 | 0.004 | 1923.541 | 6.283 | 6.295 | 12.845 | 12.906 | 0.150 |
| c1_sqrt | pair_272 | 0.004 | 1923.541 | 9.507 | 6.294 | 16.067 | 12.931 | 5.646 |
| c1_sqrt | pair_273 | 0.004 | 1923.541 | 6.320 | 6.228 | 12.714 | 12.670 | 0.141 |
| c1_sqrt | pair_274 | 0.004 | 1923.541 | 6.278 | 6.273 | 12.770 | 12.774 | 0.107 |
| c1_sqrt | pair_275 | 0.004 | 1923.541 | 6.670 | 6.270 | 13.124 | 12.970 | 0.604 |
| c1_sqrt | pair_276 | 0.004 | 1923.541 | 6.289 | 6.219 | 12.731 | 12.803 | 0.152 |
| c1_sqrt | pair_277 | 0.004 | 1923.541 | 6.319 | 6.236 | 13.591 | 12.771 | 1.435 |
| c1_sqrt | pair_278 | 0.004 | 1923.541 | 6.319 | 6.369 | 12.963 | 12.953 | 0.299 |
| c1_sqrt | pair_279 | 0.004 | 1923.541 | 6.536 | 6.267 | 13.031 | 12.947 | 0.412 |
| c1_sqrt | pair_280 | 0.004 | 1923.541 | 7.221 | 6.454 | 13.921 | 12.737 | 2.075 |
| c1_sqrt | pair_281 | 0.004 | 1923.541 | 4.029 | 3.922 | 8.118 | 8.089 | 0.066 |
| c1_sqrt | pair_282 | 0.004 | 1923.541 | 6.277 | 6.337 | 12.824 | 12.750 | 0.239 |
| c1_sqrt | pair_283 | 0.004 | 1923.541 | 6.427 | 6.235 | 12.911 | 12.869 | 0.128 |
| c1_sqrt | pair_284 | 0.004 | 1923.541 | 6.309 | 6.252 | 12.765 | 12.812 | 0.150 |
| c1_sqrt | pair_285 | 0.004 | 1923.541 | 6.276 | 6.278 | 12.744 | 12.715 | 0.054 |
| c1_sqrt | pair_286 | 0.004 | 1923.541 | 4.006 | 4.005 | 8.221 | 8.002 | 0.545 |
| c1_sqrt | pair_287 | 0.004 | 1923.541 | 3.953 | 3.889 | 7.970 | 7.995 | 0.047 |
| c1_sqrt | pair_288 | 0.004 | 1923.541 | 6.325 | 6.255 | 12.793 | 12.871 | 0.157 |
| c1_sqrt | pair_289 | 0.004 | 1923.541 | 6.307 | 6.279 | 12.793 | 12.669 | 0.219 |
| c1_sqrt | pair_290 | 0.004 | 1923.541 | 6.282 | 6.628 | 13.123 | 12.967 | 0.465 |
| c1_sqrt | pair_291 | 0.004 | 1923.541 | 6.315 | 6.325 | 12.858 | 12.878 | 0.157 |
| c1_sqrt | pair_292 | 0.004 | 1923.541 | 6.395 | 6.268 | 12.816 | 12.787 | 0.286 |
| c1_sqrt | pair_293 | 0.004 | 1923.541 | 6.339 | 6.306 | 12.903 | 12.920 | 0.191 |
| c1_sqrt | pair_294 | 0.004 | 1923.541 | 6.294 | 6.303 | 12.827 | 12.831 | 0.051 |
| c1_sqrt | pair_295 | 0.004 | 1923.541 | 6.339 | 6.404 | 12.963 | 12.797 | 0.305 |
| c1_sqrt | pair_296 | 0.004 | 1923.541 | 6.250 | 6.334 | 12.769 | 12.781 | 0.053 |
| c1_sqrt | pair_297 | 0.004 | 1923.541 | 6.294 | 6.314 | 12.764 | 12.774 | 0.070 |
| c1_sqrt | pair_298 | 0.004 | 1923.541 | 4.217 | 3.942 | 8.258 | 8.012 | 0.453 |
| c1_sqrt | pair_299 | 0.004 | 1923.541 | 6.318 | 6.235 | 12.768 | 12.754 | 0.075 |
| c1_sqrt | pair_300 | 0.004 | 1923.541 | 7.069 | 7.034 | 14.330 | 14.285 | 0.103 |
| c1_sqrt | pair_301 | 0.004 | 1923.541 | 7.062 | 7.160 | 14.473 | 14.468 | 0.020 |
| c1_sqrt | pair_302 | 0.004 | 1923.541 | 3.916 | 3.913 | 7.954 | 7.938 | 0.056 |
| c1_sqrt | pair_303 | 0.004 | 1923.541 | 3.924 | 3.989 | 8.053 | 8.014 | 0.127 |
| c1_sqrt | pair_304 | 0.004 | 1923.541 | 7.069 | 7.134 | 14.437 | 14.339 | 0.228 |
| c1_sqrt | pair_305 | 0.004 | 1923.541 | 7.231 | 7.087 | 14.579 | 14.396 | 0.382 |
| c1_sqrt | pair_306 | 0.004 | 1923.541 | 7.157 | 7.026 | 14.466 | 14.469 | 0.073 |
| c1_sqrt | pair_307 | 0.004 | 1923.541 | 7.314 | 7.021 | 14.574 | 14.394 | 0.343 |
| c1_sqrt | pair_308 | 0.004 | 1923.541 | 7.043 | 7.055 | 14.345 | 14.337 | 0.088 |
| c1_sqrt | pair_309 | 0.004 | 1923.541 | 6.262 | 6.605 | 13.149 | 13.160 | 0.279 |
| c1_sqrt | pair_310 | 0.004 | 1923.541 | 7.974 | 7.043 | 16.148 | 16.720 | 1.414 |
| c1_sqrt | pair_311 | 0.004 | 1923.541 | 7.181 | 7.297 | 14.917 | 15.153 | 0.527 |
| c1_sqrt | pair_312 | 0.004 | 1923.541 | 6.381 | 6.278 | 12.959 | 12.863 | 0.294 |
| c1_sqrt | pair_313 | 0.004 | 1923.541 | 6.423 | 6.237 | 12.964 | 12.983 | 0.245 |
| c1_sqrt | pair_314 | 0.004 | 1923.541 | 4.011 | 3.976 | 294.120 | 8.141 | 495.510 |
| c1_sqrt | pair_315 | 0.004 | 1923.541 | 7.348 | 7.247 | 14.785 | 14.374 | 0.866 |
| c1_sqrt | pair_316 | 0.004 | 1923.541 | 7.085 | 7.031 | 14.322 | 14.350 | 0.067 |
| c1_sqrt | pair_317 | 0.004 | 1923.541 | 3.976 | 3.929 | 8.057 | 8.034 | 0.149 |
| c1_sqrt | pair_318 | 0.004 | 1923.541 | 3.952 | 3.953 | 8.074 | 8.072 | 0.061 |
| c1_sqrt | pair_319 | 0.004 | 1923.541 | 7.126 | 7.020 | 14.385 | 14.284 | 0.191 |
| c1_sqrt | pair_320 | 0.004 | 1923.541 | 7.067 | 6.961 | 14.200 | 14.195 | 0.084 |
| c1_sqrt | pair_321 | 0.004 | 1923.541 | 7.026 | 7.047 | 14.375 | 14.337 | 0.124 |
| c1_sqrt | pair_322 | 0.004 | 1923.541 | 7.085 | 7.062 | 14.348 | 14.291 | 0.182 |
| c1_sqrt | pair_323 | 0.004 | 1923.541 | 7.158 | 7.007 | 14.394 | 14.375 | 0.093 |
| c1_sqrt | pair_324 | 0.004 | 1923.541 | 6.255 | 6.325 | 12.790 | 12.859 | 0.194 |
| c1_sqrt | pair_325 | 0.004 | 1923.541 | 7.144 | 6.965 | 14.332 | 14.260 | 0.125 |
| c1_sqrt | pair_326 | 0.004 | 1923.541 | 7.105 | 7.059 | 14.406 | 14.280 | 0.264 |
| c1_sqrt | pair_327 | 0.004 | 1923.541 | 6.212 | 6.181 | 12.588 | 12.593 | 0.018 |
| c1_sqrt | pair_328 | 0.004 | 1923.541 | 6.273 | 6.314 | 12.813 | 12.706 | 0.214 |
| c1_sqrt | pair_329 | 0.004 | 1923.541 | 3.957 | 4.018 | 8.123 | 8.155 | 0.198 |
| c1_sqrt | pair_330 | 0.004 | 1923.541 | 7.034 | 7.093 | 14.360 | 14.326 | 0.166 |
| c1_sqrt | pair_331 | 0.004 | 1923.541 | 3.928 | 3.895 | 7.933 | 7.913 | 0.067 |
| c1_sqrt | pair_332 | 0.004 | 1923.541 | 3.893 | 3.872 | 7.868 | 7.857 | 0.082 |
| c1_sqrt | pair_333 | 0.004 | 1923.541 | 7.055 | 6.982 | 14.282 | 14.255 | 0.063 |
| c1_sqrt | pair_334 | 0.004 | 1923.541 | 7.138 | 7.042 | 14.414 | 14.394 | 0.219 |
| c1_sqrt | pair_335 | 0.004 | 1923.541 | 7.047 | 7.073 | 14.358 | 14.264 | 0.242 |
| c1_sqrt | pair_336 | 0.004 | 1923.541 | 7.102 | 7.024 | 14.355 | 14.332 | 0.235 |
| c1_sqrt | pair_337 | 0.004 | 1923.541 | 7.026 | 7.021 | 14.249 | 14.196 | 0.236 |
| c1_sqrt | pair_338 | 0.004 | 1923.541 | 6.356 | 6.217 | 12.753 | 12.675 | 0.256 |
| c1_sqrt | pair_339 | 0.004 | 1923.541 | 7.079 | 7.042 | 14.380 | 14.283 | 0.223 |
| c1_sqrt | pair_340 | 0.004 | 1923.541 | 7.015 | 7.584 | 15.180 | 14.747 | 1.226 |
| c1_sqrt | pair_341 | 0.004 | 1923.541 | 6.452 | 6.404 | 13.031 | 12.999 | 0.286 |
| c1_sqrt | pair_342 | 0.004 | 1923.541 | 6.490 | 6.681 | 13.397 | 13.483 | 0.513 |
| c1_sqrt | pair_343 | 0.004 | 1923.541 | 4.120 | 3.950 | 8.201 | 8.092 | 0.316 |
| c1_sqrt | pair_344 | 0.004 | 1923.541 | 7.284 | 7.258 | 14.870 | 14.803 | 0.720 |
| c1_sqrt | pair_345 | 0.004 | 1923.541 | 4.001 | 4.051 | 8.185 | 8.270 | 0.221 |
| c1_sqrt | pair_346 | 0.004 | 1923.541 | 3.935 | 3.979 | 8.281 | 8.185 | 0.490 |
| c1_sqrt | pair_347 | 0.004 | 1923.541 | 5.265 | 4.182 | 9.582 | 8.281 | 2.624 |
| c1_sqrt | pair_348 | 0.004 | 1923.541 | 4.390 | 3.936 | 8.443 | 8.315 | 0.399 |
| c1_sqrt | pair_349 | 0.004 | 1923.541 | 4.645 | 4.195 | 8.961 | 7.996 | 1.753 |
| c1_sqrt | pair_350 | 0.004 | 1923.541 | 4.331 | 4.227 | 8.987 | 8.220 | 1.670 |
| c1_sqrt | pair_351 | 0.004 | 1923.541 | 4.232 | 4.444 | 8.812 | 8.140 | 1.372 |
| c1_sqrt | pair_352 | 0.004 | 1923.541 | 4.515 | 4.173 | 8.834 | 8.436 | 1.149 |
| c1_sqrt | pair_353 | 0.004 | 1923.541 | 4.371 | 4.245 | 8.739 | 8.194 | 1.219 |
| c1_sqrt | pair_354 | 0.004 | 1923.541 | 4.348 | 4.142 | 8.661 | 8.285 | 0.950 |
| c1_sqrt | pair_355 | 0.004 | 1923.541 | 4.174 | 4.023 | 8.329 | 8.263 | 0.436 |
| c1_sqrt | pair_356 | 0.004 | 1923.541 | 4.321 | 4.313 | 8.767 | 8.154 | 1.231 |
| c1_sqrt | pair_357 | 0.004 | 1923.541 | 3.998 | 3.972 | 8.132 | 7.980 | 0.283 |
| c1_sqrt | pair_358 | 0.004 | 1923.541 | 3.933 | 3.959 | 8.024 | 8.065 | 0.085 |
| c1_sqrt | pair_359 | 0.004 | 1923.541 | 3.976 | 3.898 | 8.029 | 8.032 | 0.052 |
| c1_sqrt | pair_360 | 0.004 | 1923.541 | 3.992 | 3.858 | 7.943 | 7.935 | 0.038 |
| c1_sqrt | pair_361 | 0.004 | 1923.541 | 3.982 | 3.921 | 8.039 | 7.989 | 0.125 |
| c1_sqrt | pair_362 | 0.004 | 1923.541 | 4.023 | 3.986 | 8.169 | 8.134 | 0.089 |
| c1_sqrt | pair_363 | 0.004 | 1923.541 | 4.089 | 3.917 | 8.111 | 8.111 | 0.112 |
| c1_sqrt | pair_364 | 0.004 | 1923.541 | 4.038 | 3.921 | 8.128 | 8.047 | 0.141 |
| c1_sqrt | pair_365 | 0.004 | 1923.541 | 3.964 | 3.868 | 7.959 | 8.002 | 0.107 |
| c1_sqrt | pair_366 | 0.004 | 1923.541 | 3.943 | 3.871 | 7.936 | 7.897 | 0.106 |
| c1_sqrt | pair_367 | 0.004 | 1923.541 | 4.009 | 3.886 | 8.023 | 8.029 | 0.012 |
| c1_sqrt | pair_368 | 0.004 | 1923.541 | 3.987 | 3.942 | 8.035 | 7.987 | 0.093 |
| c1_sqrt | pair_369 | 0.004 | 1923.541 | 3.971 | 3.892 | 8.014 | 8.007 | 0.113 |
| c1_sqrt | pair_370 | 0.004 | 1923.541 | 7.939 | 7.848 | 16.055 | 16.073 | 0.152 |
| c1_sqrt | pair_371 | 0.004 | 1923.541 | 11.901 | 11.876 | 24.159 | 24.054 | 0.240 |
| c1_sqrt | pair_372 | 0.004 | 1923.541 | 8.011 | 7.930 | 16.225 | 16.115 | 0.227 |
| c1_sqrt | pair_373 | 0.004 | 1923.541 | 12.015 | 11.778 | 24.205 | 24.360 | 0.317 |
| c1_sqrt | pair_374 | 0.004 | 1923.541 | 6.305 | 6.333 | 12.846 | 12.873 | 0.087 |
| c1_sqrt | pair_375 | 0.004 | 1923.541 | 11.859 | 12.059 | 26.613 | 24.733 | 3.692 |
| c1_sqrt | pair_376 | 0.004 | 1923.541 | 13.778 | 12.556 | 27.075 | 24.453 | 5.127 |
| c1_sqrt | pair_377 | 0.004 | 1923.541 | 6.544 | 6.449 | 13.270 | 13.359 | 0.478 |
| c1_sqrt | pair_378 | 0.004 | 1923.541 | 6.501 | 6.442 | 13.140 | 13.158 | 0.081 |
| c1_sqrt | pair_379 | 0.004 | 1923.541 | 3.966 | 3.958 | 8.062 | 8.096 | 0.082 |
| c1_sqrt | pair_380 | 0.004 | 1923.541 | 7.150 | 7.119 | 14.502 | 14.525 | 0.202 |
| c1_sqrt | pair_381 | 0.004 | 1923.541 | 8.035 | 8.063 | 16.342 | 16.316 | 0.209 |
| c1_sqrt | pair_382 | 0.004 | 1923.541 | 7.944 | 7.877 | 16.049 | 16.114 | 0.117 |
| c1_sqrt | pair_383 | 0.004 | 1923.541 | 7.921 | 7.791 | 15.968 | 16.051 | 0.274 |
| c1_sqrt | pair_384 | 0.004 | 1923.541 | 6.465 | 6.305 | 12.950 | 13.013 | 0.193 |
| c1_sqrt | pair_385 | 0.004 | 1923.541 | 8.144 | 8.066 | 16.447 | 16.180 | 0.761 |
| c1_sqrt | pair_386 | 0.004 | 1923.541 | 8.009 | 7.894 | 16.154 | 16.189 | 0.094 |
| c1_sqrt | pair_387 | 0.004 | 1923.541 | 6.515 | 6.502 | 13.286 | 13.429 | 0.287 |
| c1_sqrt | pair_388 | 0.004 | 1923.541 | 6.536 | 6.562 | 13.324 | 12.835 | 0.938 |
| c1_sqrt | pair_389 | 0.004 | 1923.541 | 4.024 | 3.994 | 8.189 | 8.074 | 0.204 |
| c1_sqrt | pair_390 | 0.004 | 1923.541 | 7.355 | 7.289 | 14.844 | 14.533 | 0.768 |
| c1_sqrt | pair_391 | 0.004 | 1923.541 | 8.095 | 8.135 | 16.477 | 16.255 | 0.608 |
| c1_sqrt | pair_392 | 0.004 | 1923.541 | 20.106 | 20.182 | 41.024 | 40.347 | 1.729 |
| c1_sqrt | pair_393 | 0.004 | 1923.541 | 6.711 | 6.516 | 13.524 | 13.019 | 1.076 |
| c1_sqrt | pair_394 | 0.004 | 1923.541 | 13.134 | 12.634 | 26.208 | 26.170 | 0.795 |
| c1_sqrt | pair_395 | 0.004 | 1923.541 | 20.583 | 20.513 | 42.196 | 42.260 | 0.148 |
| c1_sqrt | pair_396 | 0.004 | 1923.541 | 6.463 | 6.325 | 13.027 | 13.016 | 0.061 |
| c1_sqrt | pair_397 | 0.004 | 1923.541 | 6.292 | 6.263 | 12.740 | 12.788 | 0.118 |
| c1_sqrt | pair_398 | 0.004 | 1923.541 | 3.938 | 3.947 | 8.013 | 8.035 | 0.077 |
| c1_sqrt | pair_399 | 0.004 | 1923.541 | 8.172 | 7.600 | 15.970 | 14.282 | 2.955 |
| c1_sqrt | pair_400 | 0.004 | 1923.541 | 9.602 | 8.670 | 18.701 | 19.913 | 2.445 |
| c1_sqrt | pair_401 | 0.004 | 1923.541 | 6.782 | 6.923 | 13.895 | 14.202 | 1.159 |
| c1_sqrt | pair_402 | 0.004 | 1923.541 | 8.563 | 8.050 | 16.892 | 16.453 | 1.338 |
| c1_sqrt | pair_403 | 0.004 | 1923.541 | 8.057 | 7.999 | 16.289 | 16.058 | 0.563 |
| c1_sqrt | pair_404 | 0.004 | 1923.541 | 6.361 | 6.266 | 12.844 | 12.765 | 0.195 |
| c1_sqrt | pair_405 | 0.004 | 1923.541 | 6.324 | 6.224 | 12.720 | 12.685 | 0.068 |
| c1_sqrt | pair_406 | 0.004 | 1923.541 | 3.923 | 3.910 | 7.985 | 8.016 | 0.094 |
| c1_sqrt | pair_407 | 0.004 | 1923.541 | 7.058 | 6.993 | 14.299 | 14.294 | 0.174 |
| c1_sqrt | pair_408 | 0.004 | 1923.541 | 6.289 | 6.277 | 12.767 | 12.759 | 0.029 |
| c1_sqrt | pair_409 | 0.004 | 1923.541 | 12.677 | 12.625 | 25.745 | 25.638 | 0.274 |
| c1_sqrt | pair_410 | 0.004 | 1923.541 | 20.014 | 19.668 | 40.539 | 40.420 | 0.505 |
| c1_sqrt | pair_411 | 0.004 | 1923.541 | 6.441 | 6.325 | 12.951 | 12.869 | 0.178 |
| c1_sqrt | pair_412 | 0.004 | 1923.541 | 6.429 | 6.674 | 13.314 | 12.979 | 0.667 |
| c1_sqrt | pair_413 | 0.004 | 1923.541 | 3.952 | 4.014 | 8.113 | 8.009 | 0.247 |
| c1_sqrt | pair_414 | 0.004 | 1923.541 | 7.091 | 7.009 | 14.371 | 14.297 | 0.144 |
| c1_sqrt | pair_415 | 0.004 | 1923.541 | 6.347 | 6.299 | 12.815 | 12.827 | 0.050 |
| c1_sqrt | pair_416 | 0.004 | 1923.541 | 6.361 | 6.353 | 12.927 | 12.933 | 0.089 |
| c1_sqrt | pair_417 | 0.004 | 1923.541 | 7.099 | 6.783 | 14.131 | 12.948 | 2.137 |
| c1_sqrt | pair_418 | 0.004 | 1923.541 | 6.439 | 6.671 | 13.673 | 13.834 | 0.842 |
| c1_sqrt | pair_419 | 0.004 | 1923.541 | 4.718 | 4.351 | 9.261 | 9.257 | 0.934 |
| c1_sqrt | pair_420 | 0.004 | 1923.541 | 6.892 | 6.457 | 13.550 | 13.842 | 0.667 |
| c1_sqrt | pair_421 | 0.004 | 1923.541 | 13.422 | 13.078 | 26.884 | 25.983 | 2.011 |
| c1_sqrt | pair_422 | 0.004 | 1923.541 | 8.922 | 6.595 | 15.746 | 13.516 | 4.500 |
| c1_sqrt | pair_423 | 0.004 | 1923.541 | 6.868 | 6.582 | 13.686 | 13.355 | 0.787 |
| c1_sqrt | pair_424 | 0.004 | 1923.541 | 4.005 | 4.257 | 8.393 | 8.077 | 0.567 |
| c1_sqrt | pair_425 | 0.004 | 1923.541 | 7.485 | 7.306 | 15.149 | 15.261 | 0.432 |
| c1_sqrt | pair_426 | 0.004 | 1923.541 | 6.500 | 6.404 | 13.144 | 13.020 | 0.445 |
| c1_sqrt | pair_427 | 0.004 | 1923.541 | 6.611 | 6.612 | 13.540 | 13.098 | 1.023 |
| c1_sqrt | pair_428 | 0.004 | 1923.541 | 4.059 | 4.151 | 8.381 | 8.323 | 0.169 |
| c1_sqrt | pair_429 | 0.004 | 1923.541 | 7.347 | 7.058 | 14.645 | 14.748 | 0.213 |
| c1_sqrt | pair_430 | 0.004 | 1923.541 | 6.293 | 6.271 | 12.779 | 12.812 | 0.154 |
| c1_sqrt | pair_431 | 0.004 | 1923.541 | 3.988 | 3.914 | 8.027 | 7.946 | 0.183 |
| c1_sqrt | pair_432 | 0.004 | 1923.541 | 6.319 | 6.279 | 12.848 | 12.897 | 0.110 |
| c1_sqrt | pair_433 | 0.004 | 1923.541 | 3.962 | 3.931 | 8.022 | 7.974 | 0.178 |
| c1_sqrt | pair_434 | 0.004 | 1923.541 | 6.407 | 6.321 | 12.931 | 12.722 | 0.364 |
| c1_sqrt | pair_435 | 0.004 | 1923.541 | 4.118 | 3.935 | 8.192 | 8.093 | 0.332 |
| corrected_fdxt | pair_001 | 0.004 | 2098.098 | 7.089 | 7.411 | 14.577 | 12.959 | 2.943 |
| corrected_fdxt | pair_002 | 0.004 | 2098.098 | 4.784 | 4.866 | 9.707 | 9.621 | 0.343 |
| corrected_fdxt | pair_003 | 0.004 | 2098.098 | 5.659 | 5.508 | 11.226 | 11.230 | 0.154 |
| corrected_fdxt | pair_004 | 0.004 | 2098.098 | 6.470 | 6.372 | 12.915 | 12.910 | 0.337 |
| corrected_fdxt | pair_005 | 0.004 | 2098.098 | 2.413 | 2.317 | 4.752 | 4.718 | 0.059 |
| corrected_fdxt | pair_006 | 0.004 | 2098.098 | 4.879 | 4.682 | 9.606 | 9.622 | 0.242 |
| corrected_fdxt | pair_007 | 0.004 | 2098.098 | 12.582 | 12.492 | 25.311 | 25.334 | 0.170 |
| corrected_fdxt | pair_008 | 0.004 | 2098.098 | 5.553 | 5.484 | 11.102 | 11.032 | 0.158 |
| corrected_fdxt | pair_009 | 0.004 | 2098.098 | 17.186 | 16.524 | 34.089 | 33.650 | 0.954 |
| corrected_fdxt | pair_010 | 0.004 | 2098.098 | 5.620 | 5.568 | 11.247 | 11.188 | 0.126 |
| corrected_fdxt | pair_011 | 0.004 | 2098.098 | 3.956 | 3.901 | 7.890 | 7.889 | 0.024 |
| corrected_fdxt | pair_012 | 0.004 | 2098.098 | 3.953 | 3.952 | 7.945 | 7.962 | 0.032 |
| corrected_fdxt | pair_013 | 0.004 | 2098.098 | 4.716 | 4.751 | 9.510 | 9.549 | 0.105 |
| corrected_fdxt | pair_014 | 0.004 | 2098.098 | 4.723 | 4.996 | 9.768 | 9.402 | 0.645 |
| corrected_fdxt | pair_015 | 0.004 | 2098.098 | 4.740 | 4.674 | 9.461 | 9.427 | 0.070 |
| corrected_fdxt | pair_016 | 0.004 | 2098.098 | 2.367 | 2.324 | 4.707 | 4.680 | 0.056 |
| corrected_fdxt | pair_017 | 0.004 | 2098.098 | 2.408 | 2.386 | 4.812 | 4.717 | 0.173 |
| corrected_fdxt | pair_018 | 0.004 | 2098.098 | 8.763 | 8.601 | 17.519 | 17.495 | 0.091 |
| corrected_fdxt | pair_019 | 0.004 | 2098.098 | 6.295 | 5.870 | 12.237 | 12.142 | 1.010 |
| corrected_fdxt | pair_020 | 0.004 | 2098.098 | 16.809 | 16.513 | 34.157 | 33.537 | 1.076 |
| corrected_fdxt | pair_021 | 0.004 | 2098.098 | 5.569 | 5.511 | 11.147 | 11.040 | 0.190 |
| corrected_fdxt | pair_022 | 0.004 | 2098.098 | 15.715 | 15.654 | 31.865 | 31.799 | 0.367 |
| corrected_fdxt | pair_023 | 0.004 | 2098.098 | 5.030 | 4.842 | 9.913 | 8.097 | 3.349 |
| corrected_fdxt | pair_024 | 0.004 | 2098.098 | 11.529 | 10.342 | 22.024 | 19.926 | 4.320 |
| corrected_fdxt | pair_025 | 0.004 | 2098.098 | 18.040 | 17.584 | 36.108 | 33.645 | 4.601 |
| corrected_fdxt | pair_026 | 0.004 | 2098.098 | 3.958 | 3.914 | 7.911 | 7.857 | 0.151 |
| corrected_fdxt | pair_027 | 0.004 | 2098.098 | 4.094 | 3.917 | 8.081 | 7.975 | 0.291 |
| corrected_fdxt | pair_028 | 0.004 | 2098.098 | 2.369 | 2.384 | 4.769 | 4.784 | 0.111 |
| corrected_fdxt | pair_029 | 0.004 | 2098.098 | 4.739 | 4.667 | 9.512 | 9.384 | 0.229 |
| corrected_fdxt | pair_030 | 0.004 | 2098.098 | 4.927 | 4.699 | 9.700 | 9.438 | 0.457 |
| corrected_fdxt | pair_031 | 0.004 | 2098.098 | 5.525 | 5.427 | 11.064 | 10.987 | 0.154 |
| corrected_fdxt | pair_032 | 0.004 | 2098.098 | 6.262 | 6.238 | 12.652 | 12.570 | 0.142 |
| corrected_fdxt | pair_033 | 0.004 | 2098.098 | 2.382 | 2.379 | 4.801 | 4.755 | 0.184 |
| corrected_fdxt | pair_034 | 0.004 | 2098.098 | 4.778 | 4.676 | 9.519 | 9.561 | 0.118 |
| corrected_fdxt | pair_035 | 0.004 | 2098.098 | 6.442 | 6.206 | 13.077 | 12.499 | 1.033 |
| corrected_fdxt | pair_036 | 0.004 | 2098.098 | 5.661 | 6.343 | 12.257 | 11.958 | 1.495 |
| corrected_fdxt | pair_037 | 0.004 | 2098.098 | 8.320 | 6.429 | 15.021 | 14.161 | 3.041 |
| corrected_fdxt | pair_038 | 0.004 | 2098.098 | 5.702 | 5.515 | 11.327 | 11.371 | 0.321 |
| corrected_fdxt | pair_039 | 0.004 | 2098.098 | 4.122 | 3.948 | 8.106 | 8.163 | 0.303 |
| corrected_fdxt | pair_040 | 0.004 | 2098.098 | 4.025 | 4.023 | 8.212 | 8.180 | 0.233 |
| corrected_fdxt | pair_041 | 0.004 | 2098.098 | 4.886 | 4.795 | 9.757 | 9.735 | 0.178 |
| corrected_fdxt | pair_042 | 0.004 | 2098.098 | 5.425 | 4.748 | 10.220 | 9.650 | 1.016 |
| corrected_fdxt | pair_043 | 0.004 | 2098.098 | 4.793 | 4.663 | 9.505 | 9.448 | 0.108 |
| corrected_fdxt | pair_044 | 0.004 | 2098.098 | 2.391 | 2.383 | 4.790 | 4.799 | 0.036 |
| corrected_fdxt | pair_045 | 0.004 | 2098.098 | 2.376 | 2.338 | 4.732 | 4.729 | 0.025 |
| corrected_fdxt | pair_046 | 0.004 | 2098.098 | 6.324 | 6.240 | 12.635 | 12.586 | 0.209 |
| corrected_fdxt | pair_047 | 0.004 | 2098.098 | 5.508 | 5.464 | 11.047 | 11.016 | 0.146 |
| corrected_fdxt | pair_048 | 0.004 | 2098.098 | 6.405 | 6.350 | 12.829 | 12.786 | 0.194 |
| corrected_fdxt | pair_049 | 0.004 | 2098.098 | 5.474 | 5.515 | 11.049 | 10.986 | 0.136 |
| corrected_fdxt | pair_050 | 0.004 | 2098.098 | 6.440 | 6.278 | 12.795 | 12.747 | 0.330 |
| corrected_fdxt | pair_051 | 0.004 | 2098.098 | 3.940 | 3.964 | 7.939 | 7.918 | 0.055 |
| corrected_fdxt | pair_052 | 0.004 | 2098.098 | 6.290 | 6.290 | 12.655 | 12.680 | 0.125 |
| corrected_fdxt | pair_053 | 0.004 | 2098.098 | 6.322 | 6.280 | 12.671 | 12.605 | 0.195 |
| corrected_fdxt | pair_054 | 0.004 | 2098.098 | 3.944 | 3.887 | 7.864 | 7.828 | 0.095 |
| corrected_fdxt | pair_055 | 0.004 | 2098.098 | 3.987 | 4.024 | 8.051 | 7.827 | 0.391 |
| corrected_fdxt | pair_056 | 0.004 | 2098.098 | 4.107 | 2.422 | 6.581 | 4.677 | 3.312 |
| corrected_fdxt | pair_057 | 0.004 | 2098.098 | 4.814 | 4.620 | 9.541 | 9.468 | 0.212 |
| corrected_fdxt | pair_058 | 0.004 | 2098.098 | 4.793 | 4.668 | 9.524 | 9.455 | 0.135 |
| corrected_fdxt | pair_059 | 0.004 | 2098.098 | 4.728 | 4.655 | 9.427 | 9.456 | 0.112 |
| corrected_fdxt | pair_060 | 0.004 | 2098.098 | 2.312 | 2.418 | 4.746 | 4.728 | 0.114 |
| corrected_fdxt | pair_061 | 0.004 | 2098.098 | 4.709 | 4.672 | 9.424 | 9.413 | 0.032 |
| corrected_fdxt | pair_062 | 0.004 | 2098.098 | 4.670 | 4.700 | 9.427 | 9.435 | 0.070 |
| corrected_fdxt | pair_063 | 0.004 | 2098.098 | 4.678 | 4.669 | 9.395 | 9.406 | 0.025 |
| corrected_fdxt | pair_064 | 0.004 | 2098.098 | 4.703 | 4.644 | 9.389 | 9.364 | 0.045 |
| corrected_fdxt | pair_065 | 0.004 | 2098.098 | 4.701 | 4.779 | 9.526 | 9.438 | 0.184 |
| corrected_fdxt | pair_066 | 0.004 | 2098.098 | 4.263 | 3.897 | 8.200 | 7.906 | 0.542 |
| corrected_fdxt | pair_067 | 0.004 | 2098.098 | 4.232 | 3.937 | 8.219 | 7.882 | 0.608 |
| corrected_fdxt | pair_068 | 0.004 | 2098.098 | 4.885 | 4.700 | 9.668 | 9.454 | 0.465 |
| corrected_fdxt | pair_069 | 0.004 | 2098.098 | 4.708 | 4.701 | 9.492 | 9.379 | 0.213 |
| corrected_fdxt | pair_070 | 0.004 | 2098.098 | 4.741 | 4.706 | 9.524 | 9.433 | 0.160 |
| corrected_fdxt | pair_071 | 0.004 | 2098.098 | 2.459 | 2.369 | 4.846 | 4.810 | 0.146 |
| corrected_fdxt | pair_072 | 0.004 | 2098.098 | 2.345 | 2.361 | 4.723 | 4.697 | 0.052 |
| corrected_fdxt | pair_073 | 0.004 | 2098.098 | 5.103 | 4.870 | 10.040 | 10.038 | 0.697 |
| corrected_fdxt | pair_074 | 0.004 | 2098.098 | 4.768 | 4.750 | 9.806 | 9.490 | 0.550 |
| corrected_fdxt | pair_075 | 0.004 | 2098.098 | 4.729 | 4.689 | 9.485 | 9.471 | 0.095 |
| corrected_fdxt | pair_076 | 0.004 | 2098.098 | 4.973 | 4.658 | 9.775 | 9.423 | 0.694 |
| corrected_fdxt | pair_077 | 0.004 | 2098.098 | 4.715 | 4.806 | 9.642 | 9.452 | 0.369 |
| corrected_fdxt | pair_078 | 0.004 | 2098.098 | 4.001 | 3.922 | 7.961 | 7.961 | 0.126 |
| corrected_fdxt | pair_079 | 0.004 | 2098.098 | 4.774 | 4.690 | 9.575 | 9.404 | 0.376 |
| corrected_fdxt | pair_080 | 0.004 | 2098.098 | 4.735 | 4.656 | 9.452 | 9.464 | 0.056 |
| corrected_fdxt | pair_081 | 0.004 | 2098.098 | 4.033 | 3.886 | 7.953 | 7.871 | 0.153 |
| corrected_fdxt | pair_082 | 0.004 | 2098.098 | 3.915 | 3.934 | 7.883 | 7.840 | 0.167 |
| corrected_fdxt | pair_083 | 0.004 | 2098.098 | 2.327 | 2.409 | 4.754 | 4.799 | 0.086 |
| corrected_fdxt | pair_084 | 0.004 | 2098.098 | 4.719 | 4.693 | 9.461 | 9.418 | 0.077 |
| corrected_fdxt | pair_085 | 0.004 | 2098.098 | 5.529 | 5.414 | 11.002 | 10.997 | 0.039 |
| corrected_fdxt | pair_086 | 0.004 | 2098.098 | 2.464 | 2.401 | 4.887 | 4.808 | 0.244 |
| corrected_fdxt | pair_087 | 0.004 | 2098.098 | 4.708 | 4.682 | 9.437 | 9.401 | 0.065 |
| corrected_fdxt | pair_088 | 0.004 | 2098.098 | 5.452 | 5.506 | 11.024 | 11.020 | 0.086 |
| corrected_fdxt | pair_089 | 0.004 | 2098.098 | 5.544 | 5.502 | 11.155 | 11.123 | 0.122 |
| corrected_fdxt | pair_090 | 0.004 | 2098.098 | 5.545 | 5.492 | 11.144 | 11.157 | 0.215 |
| corrected_fdxt | pair_091 | 0.004 | 2098.098 | 5.938 | 5.395 | 11.424 | 11.335 | 0.584 |
| corrected_fdxt | pair_092 | 0.004 | 2098.098 | 3.970 | 3.861 | 7.900 | 7.839 | 0.107 |
| corrected_fdxt | pair_093 | 0.004 | 2098.098 | 3.906 | 3.934 | 8.049 | 7.905 | 0.340 |
| corrected_fdxt | pair_094 | 0.004 | 2098.098 | 6.010 | 4.798 | 11.071 | 9.469 | 2.806 |
| corrected_fdxt | pair_095 | 0.004 | 2098.098 | 4.718 | 4.674 | 9.463 | 9.432 | 0.078 |
| corrected_fdxt | pair_096 | 0.004 | 2098.098 | 4.756 | 4.650 | 9.519 | 9.489 | 0.244 |
| corrected_fdxt | pair_097 | 0.004 | 2098.098 | 2.586 | 2.340 | 4.942 | 4.743 | 0.378 |
| corrected_fdxt | pair_098 | 0.004 | 2098.098 | 2.348 | 2.380 | 4.745 | 4.700 | 0.086 |
| corrected_fdxt | pair_099 | 0.004 | 2098.098 | 5.540 | 5.479 | 11.115 | 11.037 | 0.193 |
| corrected_fdxt | pair_100 | 0.004 | 2098.098 | 5.535 | 5.474 | 11.262 | 11.043 | 0.479 |
| corrected_fdxt | pair_101 | 0.004 | 2098.098 | 5.539 | 5.495 | 11.130 | 11.025 | 0.293 |
| corrected_fdxt | pair_102 | 0.004 | 2098.098 | 5.488 | 5.505 | 11.096 | 11.031 | 0.155 |
| corrected_fdxt | pair_103 | 0.004 | 2098.098 | 5.584 | 5.455 | 11.175 | 11.286 | 0.222 |
| corrected_fdxt | pair_104 | 0.004 | 2098.098 | 4.011 | 3.961 | 8.227 | 8.065 | 0.431 |
| corrected_fdxt | pair_105 | 0.004 | 2098.098 | 5.475 | 5.404 | 11.062 | 10.982 | 0.281 |
| corrected_fdxt | pair_106 | 0.004 | 2098.098 | 5.723 | 5.451 | 11.543 | 11.204 | 0.847 |
| corrected_fdxt | pair_107 | 0.004 | 2098.098 | 3.958 | 3.889 | 7.879 | 7.893 | 0.053 |
| corrected_fdxt | pair_108 | 0.004 | 2098.098 | 3.958 | 4.008 | 8.042 | 8.056 | 0.150 |
| corrected_fdxt | pair_109 | 0.004 | 2098.098 | 2.582 | 2.359 | 4.958 | 4.955 | 0.258 |
| corrected_fdxt | pair_110 | 0.004 | 2098.098 | 4.725 | 4.656 | 9.685 | 9.408 | 0.588 |
| corrected_fdxt | pair_111 | 0.004 | 2098.098 | 2.375 | 2.357 | 4.837 | 4.820 | 0.168 |
| corrected_fdxt | pair_112 | 0.004 | 2098.098 | 4.744 | 4.670 | 9.483 | 9.399 | 0.171 |
| corrected_fdxt | pair_113 | 0.004 | 2098.098 | 6.653 | 6.394 | 13.197 | 13.321 | 0.546 |
| corrected_fdxt | pair_114 | 0.004 | 2098.098 | 5.631 | 5.909 | 11.632 | 11.210 | 0.938 |
| corrected_fdxt | pair_115 | 0.004 | 2098.098 | 6.502 | 6.488 | 13.147 | 12.667 | 0.911 |
| corrected_fdxt | pair_116 | 0.004 | 2098.098 | 5.745 | 5.753 | 11.620 | 11.089 | 0.952 |
| corrected_fdxt | pair_117 | 0.004 | 2098.098 | 5.348 | 4.252 | 9.704 | 8.806 | 2.341 |
| corrected_fdxt | pair_118 | 0.004 | 2098.098 | 4.399 | 4.279 | 11.005 | 9.115 | 4.355 |
| corrected_fdxt | pair_119 | 0.004 | 2098.098 | 9.352 | 5.120 | 14.521 | 15.515 | 4.535 |
| corrected_fdxt | pair_120 | 0.004 | 2098.098 | 5.311 | 5.028 | 10.392 | 9.543 | 1.631 |
| corrected_fdxt | pair_121 | 0.004 | 2098.098 | 5.368 | 5.662 | 11.084 | 9.574 | 2.764 |
| corrected_fdxt | pair_122 | 0.004 | 2098.098 | 2.418 | 2.332 | 4.767 | 4.726 | 0.077 |
| corrected_fdxt | pair_123 | 0.004 | 2098.098 | 2.412 | 2.462 | 4.893 | 4.825 | 0.236 |
| corrected_fdxt | pair_124 | 0.004 | 2098.098 | 6.318 | 6.418 | 12.810 | 12.739 | 0.319 |
| corrected_fdxt | pair_125 | 0.004 | 2098.098 | 5.682 | 5.502 | 11.243 | 11.164 | 0.223 |
| corrected_fdxt | pair_126 | 0.004 | 2098.098 | 6.378 | 6.246 | 12.701 | 12.706 | 0.251 |
| corrected_fdxt | pair_127 | 0.004 | 2098.098 | 5.668 | 5.619 | 11.349 | 11.208 | 0.503 |
| corrected_fdxt | pair_128 | 0.004 | 2098.098 | 6.311 | 6.517 | 12.901 | 12.946 | 0.189 |
| corrected_fdxt | pair_129 | 0.004 | 2098.098 | 3.992 | 4.078 | 8.106 | 8.049 | 0.167 |
| corrected_fdxt | pair_130 | 0.004 | 2098.098 | 6.289 | 6.378 | 12.754 | 12.724 | 0.167 |
| corrected_fdxt | pair_131 | 0.004 | 2098.098 | 7.616 | 6.569 | 14.259 | 13.501 | 1.686 |
| corrected_fdxt | pair_132 | 0.004 | 2098.098 | 4.039 | 3.923 | 7.998 | 7.874 | 0.255 |
| corrected_fdxt | pair_133 | 0.004 | 2098.098 | 4.006 | 4.008 | 8.049 | 8.058 | 0.194 |
| corrected_fdxt | pair_134 | 0.004 | 2098.098 | 2.381 | 2.430 | 4.828 | 4.763 | 0.117 |
| corrected_fdxt | pair_135 | 0.004 | 2098.098 | 4.803 | 4.735 | 9.584 | 9.617 | 0.102 |
| corrected_fdxt | pair_136 | 0.004 | 2098.098 | 2.333 | 2.896 | 5.247 | 4.808 | 0.880 |
| corrected_fdxt | pair_137 | 0.004 | 2098.098 | 3.859 | 2.453 | 6.329 | 4.966 | 2.557 |
| corrected_fdxt | pair_138 | 0.004 | 2098.098 | 3.856 | 3.840 | 7.716 | 4.705 | 5.241 |
| corrected_fdxt | pair_139 | 0.004 | 2098.098 | 2.686 | 2.522 | 5.231 | 4.754 | 0.839 |
| corrected_fdxt | pair_140 | 0.004 | 2098.098 | 2.689 | 2.569 | 5.276 | 4.902 | 0.826 |
| corrected_fdxt | pair_141 | 0.004 | 2098.098 | 2.585 | 2.634 | 5.241 | 4.939 | 0.659 |
| corrected_fdxt | pair_142 | 0.004 | 2098.098 | 2.616 | 2.569 | 5.207 | 4.924 | 0.633 |
| corrected_fdxt | pair_143 | 0.004 | 2098.098 | 5.878 | 3.086 | 8.984 | 6.113 | 6.041 |
| corrected_fdxt | pair_144 | 0.004 | 2098.098 | 3.014 | 3.247 | 6.279 | 6.910 | 1.183 |
| corrected_fdxt | pair_145 | 0.004 | 2098.098 | 3.774 | 2.646 | 6.439 | 5.992 | 1.206 |
| corrected_fdxt | pair_146 | 0.004 | 2098.098 | 2.687 | 3.100 | 5.807 | 5.923 | 0.747 |
| corrected_fdxt | pair_147 | 0.004 | 2098.098 | 2.811 | 2.642 | 5.475 | 5.296 | 0.727 |
| corrected_fdxt | pair_148 | 0.004 | 2098.098 | 2.951 | 2.677 | 5.649 | 5.916 | 0.594 |
| corrected_fdxt | pair_149 | 0.004 | 2098.098 | 2.913 | 2.661 | 5.594 | 5.908 | 0.659 |
| corrected_fdxt | pair_150 | 0.004 | 2098.098 | 2.653 | 2.767 | 5.437 | 5.509 | 0.450 |
| corrected_fdxt | pair_151 | 0.004 | 2098.098 | 2.729 | 2.730 | 5.476 | 5.678 | 0.549 |
| corrected_fdxt | pair_152 | 0.004 | 2098.098 | 2.742 | 2.837 | 5.599 | 5.899 | 0.546 |
| corrected_fdxt | pair_153 | 0.004 | 2098.098 | 2.630 | 2.784 | 5.431 | 5.399 | 0.637 |
| corrected_fdxt | pair_154 | 0.004 | 2098.098 | 2.696 | 2.657 | 5.372 | 5.418 | 0.365 |
| corrected_fdxt | pair_155 | 0.004 | 2098.098 | 2.717 | 2.777 | 5.510 | 5.567 | 0.675 |
| corrected_fdxt | pair_156 | 0.004 | 2098.098 | 2.521 | 2.822 | 5.359 | 5.386 | 0.446 |
| corrected_fdxt | pair_157 | 0.004 | 2098.098 | 2.789 | 2.644 | 5.451 | 5.468 | 0.711 |
| corrected_fdxt | pair_158 | 0.004 | 2098.098 | 2.773 | 2.812 | 5.603 | 5.663 | 0.449 |
| corrected_fdxt | pair_159 | 0.004 | 2098.098 | 2.607 | 2.717 | 5.341 | 5.081 | 0.639 |
| corrected_fdxt | pair_160 | 0.004 | 2098.098 | 5.459 | 5.350 | 10.874 | 10.917 | 0.619 |
| corrected_fdxt | pair_161 | 0.004 | 2098.098 | 5.354 | 5.537 | 10.988 | 10.806 | 0.448 |
| corrected_fdxt | pair_162 | 0.004 | 2098.098 | 5.489 | 5.127 | 10.664 | 10.764 | 0.658 |
| corrected_fdxt | pair_163 | 0.004 | 2098.098 | 5.034 | 4.984 | 10.072 | 10.057 | 0.428 |
| corrected_fdxt | pair_164 | 0.004 | 2098.098 | 4.307 | 4.279 | 8.623 | 8.577 | 0.733 |
| corrected_fdxt | pair_165 | 0.004 | 2098.098 | 4.042 | 4.106 | 8.192 | 8.284 | 0.169 |
| corrected_fdxt | pair_166 | 0.004 | 2098.098 | 5.045 | 4.936 | 10.026 | 9.761 | 0.784 |
| corrected_fdxt | pair_167 | 0.004 | 2098.098 | 5.024 | 5.052 | 10.128 | 9.594 | 0.933 |
| corrected_fdxt | pair_168 | 0.004 | 2098.098 | 5.032 | 5.175 | 10.260 | 9.525 | 1.283 |
| corrected_fdxt | pair_169 | 0.004 | 2098.098 | 2.543 | 2.624 | 5.191 | 5.013 | 0.625 |
| corrected_fdxt | pair_170 | 0.004 | 2098.098 | 2.660 | 2.546 | 5.225 | 4.793 | 0.804 |
| corrected_fdxt | pair_171 | 0.004 | 2098.098 | 5.039 | 4.970 | 10.056 | 9.641 | 0.760 |
| corrected_fdxt | pair_172 | 0.004 | 2098.098 | 5.015 | 4.875 | 9.935 | 9.480 | 0.819 |
| corrected_fdxt | pair_173 | 0.004 | 2098.098 | 4.966 | 4.957 | 9.968 | 9.665 | 0.722 |
| corrected_fdxt | pair_174 | 0.004 | 2098.098 | 4.873 | 5.032 | 9.948 | 9.523 | 0.821 |
| corrected_fdxt | pair_175 | 0.004 | 2098.098 | 5.185 | 5.282 | 10.515 | 9.684 | 1.620 |
| corrected_fdxt | pair_176 | 0.004 | 2098.098 | 4.114 | 4.237 | 8.488 | 8.303 | 0.758 |
| corrected_fdxt | pair_177 | 0.004 | 2098.098 | 5.274 | 6.388 | 11.710 | 11.243 | 2.547 |
| corrected_fdxt | pair_178 | 0.004 | 2098.098 | 4.945 | 5.041 | 10.042 | 9.679 | 0.827 |
| corrected_fdxt | pair_179 | 0.004 | 2098.098 | 4.929 | 4.066 | 9.043 | 9.131 | 1.128 |
| corrected_fdxt | pair_180 | 0.004 | 2098.098 | 4.127 | 4.205 | 8.371 | 8.314 | 0.549 |
| corrected_fdxt | pair_181 | 0.004 | 2098.098 | 2.497 | 2.493 | 5.024 | 4.691 | 0.631 |
| corrected_fdxt | pair_182 | 0.004 | 2098.098 | 4.903 | 4.837 | 9.794 | 9.577 | 0.522 |
| corrected_fdxt | pair_183 | 0.004 | 2098.098 | 5.645 | 5.533 | 11.245 | 11.267 | 0.215 |
| corrected_fdxt | pair_184 | 0.004 | 2098.098 | 13.115 | 13.393 | 26.765 | 26.554 | 1.451 |
| corrected_fdxt | pair_185 | 0.004 | 2098.098 | 5.575 | 5.456 | 11.093 | 11.111 | 0.127 |
| corrected_fdxt | pair_186 | 0.004 | 2098.098 | 3.998 | 3.874 | 7.905 | 7.846 | 0.114 |
| corrected_fdxt | pair_187 | 0.004 | 2098.098 | 3.923 | 3.991 | 7.949 | 7.954 | 0.087 |
| corrected_fdxt | pair_188 | 0.004 | 2098.098 | 4.722 | 4.995 | 9.766 | 9.591 | 0.455 |
| corrected_fdxt | pair_189 | 0.004 | 2098.098 | 4.985 | 5.029 | 10.065 | 9.454 | 1.076 |
| corrected_fdxt | pair_190 | 0.004 | 2098.098 | 4.768 | 4.979 | 9.794 | 9.453 | 0.597 |
| corrected_fdxt | pair_191 | 0.004 | 2098.098 | 2.770 | 2.668 | 5.456 | 4.764 | 1.280 |
| corrected_fdxt | pair_192 | 0.004 | 2098.098 | 2.455 | 2.434 | 4.908 | 4.748 | 0.331 |
| corrected_fdxt | pair_193 | 0.004 | 2098.098 | 9.366 | 9.462 | 18.968 | 17.690 | 2.563 |
| corrected_fdxt | pair_194 | 0.004 | 2098.098 | 6.297 | 5.954 | 12.313 | 12.074 | 1.468 |
| corrected_fdxt | pair_195 | 0.004 | 2098.098 | 13.408 | 13.626 | 27.376 | 25.470 | 3.605 |
| corrected_fdxt | pair_196 | 0.004 | 2098.098 | 5.676 | 5.645 | 11.380 | 11.003 | 0.656 |
| corrected_fdxt | pair_197 | 0.004 | 2098.098 | 12.934 | 12.992 | 26.171 | 25.515 | 1.397 |
| corrected_fdxt | pair_198 | 0.004 | 2098.098 | 4.079 | 4.059 | 8.170 | 8.133 | 0.242 |
| corrected_fdxt | pair_199 | 0.004 | 2098.098 | 9.610 | 9.628 | 19.389 | 19.010 | 0.715 |
| corrected_fdxt | pair_200 | 0.004 | 2098.098 | 14.780 | 14.333 | 29.362 | 27.907 | 4.571 |
| corrected_fdxt | pair_201 | 0.004 | 2098.098 | 4.330 | 4.441 | 8.814 | 8.359 | 1.267 |
| corrected_fdxt | pair_202 | 0.004 | 2098.098 | 4.380 | 4.340 | 8.756 | 7.813 | 1.646 |
| corrected_fdxt | pair_203 | 0.004 | 2098.098 | 2.693 | 2.554 | 5.264 | 4.700 | 0.992 |
| corrected_fdxt | pair_204 | 0.004 | 2098.098 | 5.147 | 5.192 | 10.387 | 9.517 | 1.637 |
| corrected_fdxt | pair_205 | 0.004 | 2098.098 | 5.967 | 5.898 | 11.927 | 11.067 | 1.543 |
| corrected_fdxt | pair_206 | 0.004 | 2098.098 | 5.971 | 5.910 | 11.937 | 10.930 | 1.766 |
| corrected_fdxt | pair_207 | 0.004 | 2098.098 | 4.401 | 4.254 | 8.691 | 7.945 | 1.352 |
| corrected_fdxt | pair_208 | 0.004 | 2098.098 | 4.227 | 4.300 | 8.566 | 7.826 | 1.370 |
| corrected_fdxt | pair_209 | 0.004 | 2098.098 | 5.174 | 5.330 | 10.558 | 9.471 | 1.920 |
| corrected_fdxt | pair_210 | 0.004 | 2098.098 | 5.251 | 5.106 | 10.402 | 9.652 | 1.382 |
| corrected_fdxt | pair_211 | 0.004 | 2098.098 | 5.182 | 5.145 | 10.404 | 9.506 | 1.649 |
| corrected_fdxt | pair_212 | 0.004 | 2098.098 | 2.622 | 2.446 | 5.085 | 4.720 | 0.641 |
| corrected_fdxt | pair_213 | 0.004 | 2098.098 | 2.435 | 2.512 | 4.965 | 4.739 | 0.418 |
| corrected_fdxt | pair_214 | 0.004 | 2098.098 | 5.865 | 5.808 | 11.731 | 11.194 | 1.053 |
| corrected_fdxt | pair_215 | 0.004 | 2098.098 | 5.580 | 5.601 | 11.241 | 11.015 | 0.474 |
| corrected_fdxt | pair_216 | 0.004 | 2098.098 | 5.566 | 5.536 | 11.161 | 10.952 | 0.384 |
| corrected_fdxt | pair_217 | 0.004 | 2098.098 | 5.820 | 6.046 | 11.941 | 11.433 | 1.225 |
| corrected_fdxt | pair_218 | 0.004 | 2098.098 | 5.937 | 5.854 | 11.846 | 10.967 | 1.551 |
| corrected_fdxt | pair_219 | 0.004 | 2098.098 | 4.110 | 4.132 | 8.277 | 7.898 | 0.697 |
| corrected_fdxt | pair_220 | 0.004 | 2098.098 | 6.035 | 5.981 | 12.077 | 11.437 | 1.530 |
| corrected_fdxt | pair_221 | 0.004 | 2098.098 | 5.831 | 6.042 | 11.931 | 10.963 | 1.722 |
| corrected_fdxt | pair_222 | 0.004 | 2098.098 | 4.153 | 4.091 | 8.279 | 7.791 | 0.883 |
| corrected_fdxt | pair_223 | 0.004 | 2098.098 | 4.208 | 4.238 | 8.480 | 7.894 | 1.024 |
| corrected_fdxt | pair_224 | 0.004 | 2098.098 | 2.497 | 2.456 | 4.971 | 4.752 | 0.458 |
| corrected_fdxt | pair_225 | 0.004 | 2098.098 | 6.560 | 7.657 | 14.292 | 11.525 | 6.809 |
| corrected_fdxt | pair_226 | 0.004 | 2098.098 | 7.435 | 6.570 | 14.071 | 13.091 | 3.734 |
| corrected_fdxt | pair_227 | 0.004 | 2098.098 | 4.337 | 4.737 | 9.116 | 9.606 | 1.067 |
| corrected_fdxt | pair_228 | 0.004 | 2098.098 | 4.283 | 4.336 | 9.346 | 9.354 | 0.540 |
| corrected_fdxt | pair_229 | 0.004 | 2098.098 | 5.393 | 5.197 | 10.641 | 11.086 | 0.772 |
| corrected_fdxt | pair_230 | 0.004 | 2098.098 | 4.893 | 4.956 | 9.963 | 9.396 | 0.984 |
| corrected_fdxt | pair_231 | 0.004 | 2098.098 | 5.143 | 4.759 | 9.948 | 9.765 | 0.721 |
| corrected_fdxt | pair_232 | 0.004 | 2098.098 | 2.386 | 2.344 | 4.747 | 4.777 | 0.077 |
| corrected_fdxt | pair_233 | 0.004 | 2098.098 | 2.362 | 2.394 | 4.774 | 4.703 | 0.172 |
| corrected_fdxt | pair_234 | 0.004 | 2098.098 | 9.226 | 9.084 | 18.433 | 18.170 | 1.262 |
| corrected_fdxt | pair_235 | 0.004 | 2098.098 | 5.492 | 5.541 | 11.097 | 10.970 | 0.223 |
| corrected_fdxt | pair_236 | 0.004 | 2098.098 | 16.527 | 16.315 | 33.248 | 33.258 | 0.164 |
| corrected_fdxt | pair_237 | 0.004 | 2098.098 | 6.150 | 5.593 | 11.801 | 11.100 | 1.296 |
| corrected_fdxt | pair_238 | 0.004 | 2098.098 | 15.708 | 15.584 | 31.652 | 31.494 | 0.289 |
| corrected_fdxt | pair_239 | 0.004 | 2098.098 | 3.899 | 3.914 | 7.845 | 7.839 | 0.013 |
| corrected_fdxt | pair_240 | 0.004 | 2098.098 | 9.412 | 9.410 | 18.972 | 18.843 | 0.254 |
| corrected_fdxt | pair_241 | 0.004 | 2098.098 | 16.435 | 16.314 | 33.141 | 33.148 | 0.037 |
| corrected_fdxt | pair_242 | 0.004 | 2098.098 | 3.903 | 3.896 | 7.832 | 7.829 | 0.011 |
| corrected_fdxt | pair_243 | 0.004 | 2098.098 | 3.932 | 3.877 | 7.842 | 7.858 | 0.036 |
| corrected_fdxt | pair_244 | 0.004 | 2098.098 | 2.315 | 2.361 | 4.693 | 4.706 | 0.028 |
| corrected_fdxt | pair_245 | 0.004 | 2098.098 | 4.671 | 4.659 | 9.373 | 9.385 | 0.066 |
| corrected_fdxt | pair_246 | 0.004 | 2098.098 | 3.911 | 3.871 | 7.815 | 7.804 | 0.033 |
| corrected_fdxt | pair_247 | 0.004 | 2098.098 | 3.914 | 3.902 | 7.849 | 7.863 | 0.028 |
| corrected_fdxt | pair_248 | 0.004 | 2098.098 | 4.740 | 4.728 | 9.514 | 9.511 | 0.059 |
| corrected_fdxt | pair_249 | 0.004 | 2098.098 | 4.680 | 4.649 | 9.371 | 9.389 | 0.041 |
| corrected_fdxt | pair_250 | 0.004 | 2098.098 | 4.646 | 4.641 | 9.333 | 9.346 | 0.043 |
| corrected_fdxt | pair_251 | 0.004 | 2098.098 | 2.321 | 2.392 | 4.731 | 4.701 | 0.088 |
| corrected_fdxt | pair_252 | 0.004 | 2098.098 | 2.365 | 2.354 | 4.735 | 4.744 | 0.065 |
| corrected_fdxt | pair_253 | 0.004 | 2098.098 | 5.591 | 5.433 | 11.080 | 11.063 | 0.051 |
| corrected_fdxt | pair_254 | 0.004 | 2098.098 | 5.451 | 5.428 | 10.931 | 10.927 | 0.036 |
| corrected_fdxt | pair_255 | 0.004 | 2098.098 | 5.521 | 5.430 | 11.007 | 10.959 | 0.122 |
| corrected_fdxt | pair_256 | 0.004 | 2098.098 | 5.503 | 5.435 | 10.998 | 10.970 | 0.065 |
| corrected_fdxt | pair_257 | 0.004 | 2098.098 | 5.497 | 5.551 | 11.109 | 11.075 | 0.136 |
| corrected_fdxt | pair_258 | 0.004 | 2098.098 | 3.924 | 3.901 | 7.857 | 7.848 | 0.053 |
| corrected_fdxt | pair_259 | 0.004 | 2098.098 | 6.066 | 7.523 | 13.651 | 11.051 | 4.622 |
| corrected_fdxt | pair_260 | 0.004 | 2098.098 | 5.498 | 5.430 | 10.983 | 10.967 | 0.050 |
| corrected_fdxt | pair_261 | 0.004 | 2098.098 | 3.966 | 3.860 | 7.860 | 7.886 | 0.048 |
| corrected_fdxt | pair_262 | 0.004 | 2098.098 | 3.914 | 3.850 | 7.797 | 7.779 | 0.061 |
| corrected_fdxt | pair_263 | 0.004 | 2098.098 | 2.315 | 2.328 | 4.660 | 4.627 | 0.069 |
| corrected_fdxt | pair_264 | 0.004 | 2098.098 | 4.686 | 4.641 | 9.370 | 9.343 | 0.057 |
| corrected_fdxt | pair_265 | 0.004 | 2098.098 | 4.055 | 3.879 | 7.969 | 7.932 | 0.215 |
| corrected_fdxt | pair_266 | 0.004 | 2098.098 | 3.915 | 3.924 | 7.872 | 7.910 | 0.076 |
| corrected_fdxt | pair_267 | 0.004 | 2098.098 | 3.950 | 3.915 | 7.898 | 7.913 | 0.054 |
| corrected_fdxt | pair_268 | 0.004 | 2098.098 | 3.896 | 3.898 | 7.828 | 7.820 | 0.043 |
| corrected_fdxt | pair_269 | 0.004 | 2098.098 | 2.357 | 2.393 | 4.787 | 4.811 | 0.047 |
| corrected_fdxt | pair_270 | 0.004 | 2098.098 | 2.429 | 2.379 | 4.824 | 4.781 | 0.099 |
| corrected_fdxt | pair_271 | 0.004 | 2098.098 | 3.885 | 3.873 | 7.790 | 7.767 | 0.075 |
| corrected_fdxt | pair_272 | 0.004 | 2098.098 | 4.094 | 3.912 | 8.040 | 7.889 | 0.297 |
| corrected_fdxt | pair_273 | 0.004 | 2098.098 | 3.931 | 3.915 | 7.880 | 7.832 | 0.108 |
| corrected_fdxt | pair_274 | 0.004 | 2098.098 | 3.898 | 3.895 | 7.828 | 7.837 | 0.018 |
| corrected_fdxt | pair_275 | 0.004 | 2098.098 | 3.881 | 3.891 | 7.804 | 7.826 | 0.040 |
| corrected_fdxt | pair_276 | 0.004 | 2098.098 | 3.921 | 3.851 | 7.804 | 7.797 | 0.028 |
| corrected_fdxt | pair_277 | 0.004 | 2098.098 | 3.947 | 3.901 | 7.880 | 7.888 | 0.014 |
| corrected_fdxt | pair_278 | 0.004 | 2098.098 | 4.024 | 3.895 | 7.952 | 7.844 | 0.256 |
| corrected_fdxt | pair_279 | 0.004 | 2098.098 | 3.885 | 3.896 | 7.815 | 7.823 | 0.043 |
| corrected_fdxt | pair_280 | 0.004 | 2098.098 | 3.913 | 3.845 | 7.790 | 7.749 | 0.085 |
| corrected_fdxt | pair_281 | 0.004 | 2098.098 | 2.386 | 2.342 | 4.745 | 4.781 | 0.067 |
| corrected_fdxt | pair_282 | 0.004 | 2098.098 | 3.957 | 3.935 | 7.927 | 7.831 | 0.249 |
| corrected_fdxt | pair_283 | 0.004 | 2098.098 | 3.930 | 3.937 | 7.906 | 7.846 | 0.127 |
| corrected_fdxt | pair_284 | 0.004 | 2098.098 | 3.902 | 3.938 | 7.980 | 7.963 | 0.132 |
| corrected_fdxt | pair_285 | 0.004 | 2098.098 | 6.429 | 5.180 | 11.791 | 7.864 | 6.870 |
| corrected_fdxt | pair_286 | 0.004 | 2098.098 | 2.604 | 2.454 | 5.079 | 4.724 | 0.667 |
| corrected_fdxt | pair_287 | 0.004 | 2098.098 | 2.403 | 2.392 | 4.811 | 4.701 | 0.198 |
| corrected_fdxt | pair_288 | 0.004 | 2098.098 | 4.153 | 4.034 | 8.225 | 8.289 | 0.336 |
| corrected_fdxt | pair_289 | 0.004 | 2098.098 | 3.949 | 4.298 | 8.284 | 7.873 | 0.745 |
| corrected_fdxt | pair_290 | 0.004 | 2098.098 | 4.237 | 4.336 | 8.607 | 7.880 | 1.293 |
| corrected_fdxt | pair_291 | 0.004 | 2098.098 | 4.227 | 4.398 | 8.658 | 7.829 | 1.442 |
| corrected_fdxt | pair_292 | 0.004 | 2098.098 | 4.229 | 4.309 | 8.572 | 7.794 | 1.384 |
| corrected_fdxt | pair_293 | 0.004 | 2098.098 | 4.247 | 4.309 | 8.588 | 8.029 | 1.125 |
| corrected_fdxt | pair_294 | 0.004 | 2098.098 | 4.210 | 4.126 | 8.371 | 7.888 | 0.923 |
| corrected_fdxt | pair_295 | 0.004 | 2098.098 | 4.104 | 4.171 | 8.309 | 7.967 | 0.701 |
| corrected_fdxt | pair_296 | 0.004 | 2098.098 | 4.264 | 4.047 | 8.346 | 7.864 | 0.855 |
| corrected_fdxt | pair_297 | 0.004 | 2098.098 | 4.109 | 4.113 | 8.261 | 7.848 | 0.769 |
| corrected_fdxt | pair_298 | 0.004 | 2098.098 | 2.465 | 2.436 | 4.921 | 4.749 | 0.353 |
| corrected_fdxt | pair_299 | 0.004 | 2098.098 | 3.908 | 3.903 | 7.844 | 7.805 | 0.089 |
| corrected_fdxt | pair_300 | 0.004 | 2098.098 | 4.731 | 4.765 | 9.540 | 9.591 | 0.169 |
| corrected_fdxt | pair_301 | 0.004 | 2098.098 | 6.170 | 4.709 | 10.924 | 9.584 | 2.517 |
| corrected_fdxt | pair_302 | 0.004 | 2098.098 | 2.347 | 2.429 | 4.793 | 4.785 | 0.128 |
| corrected_fdxt | pair_303 | 0.004 | 2098.098 | 2.381 | 2.333 | 4.730 | 4.708 | 0.076 |
| corrected_fdxt | pair_304 | 0.004 | 2098.098 | 4.959 | 5.099 | 10.105 | 9.720 | 0.945 |
| corrected_fdxt | pair_305 | 0.004 | 2098.098 | 5.119 | 5.043 | 10.208 | 9.646 | 1.228 |
| corrected_fdxt | pair_306 | 0.004 | 2098.098 | 5.074 | 5.186 | 10.309 | 9.572 | 1.450 |
| corrected_fdxt | pair_307 | 0.004 | 2098.098 | 5.054 | 4.967 | 10.064 | 9.509 | 1.093 |
| corrected_fdxt | pair_308 | 0.004 | 2098.098 | 4.856 | 4.968 | 9.870 | 9.551 | 0.664 |
| corrected_fdxt | pair_309 | 0.004 | 2098.098 | 4.213 | 4.075 | 8.328 | 7.994 | 0.691 |
| corrected_fdxt | pair_310 | 0.004 | 2098.098 | 4.891 | 4.903 | 9.839 | 9.515 | 0.674 |
| corrected_fdxt | pair_311 | 0.004 | 2098.098 | 4.958 | 4.891 | 9.899 | 9.511 | 0.701 |
| corrected_fdxt | pair_312 | 0.004 | 2098.098 | 4.070 | 4.106 | 8.214 | 7.972 | 0.454 |
| corrected_fdxt | pair_313 | 0.004 | 2098.098 | 4.164 | 4.114 | 8.318 | 7.957 | 0.632 |
| corrected_fdxt | pair_314 | 0.004 | 2098.098 | 2.562 | 2.541 | 5.118 | 4.892 | 0.528 |
| corrected_fdxt | pair_315 | 0.004 | 2098.098 | 4.963 | 5.090 | 10.107 | 9.849 | 0.708 |
| corrected_fdxt | pair_316 | 0.004 | 2098.098 | 5.060 | 5.137 | 10.244 | 10.438 | 0.581 |
| corrected_fdxt | pair_317 | 0.004 | 2098.098 | 2.675 | 2.522 | 5.215 | 5.102 | 0.419 |
| corrected_fdxt | pair_318 | 0.004 | 2098.098 | 2.514 | 2.529 | 5.067 | 5.110 | 0.273 |
| corrected_fdxt | pair_319 | 0.004 | 2098.098 | 5.191 | 5.023 | 10.266 | 10.314 | 0.729 |
| corrected_fdxt | pair_320 | 0.004 | 2098.098 | 5.030 | 5.214 | 10.287 | 10.165 | 0.578 |
| corrected_fdxt | pair_321 | 0.004 | 2098.098 | 4.956 | 5.127 | 10.136 | 9.845 | 0.766 |
| corrected_fdxt | pair_322 | 0.004 | 2098.098 | 4.876 | 4.837 | 9.758 | 9.690 | 0.315 |
| corrected_fdxt | pair_323 | 0.004 | 2098.098 | 4.838 | 4.801 | 9.683 | 9.544 | 0.264 |
| corrected_fdxt | pair_324 | 0.004 | 2098.098 | 3.953 | 3.970 | 7.962 | 7.926 | 0.177 |
| corrected_fdxt | pair_325 | 0.004 | 2098.098 | 4.721 | 4.689 | 9.453 | 9.504 | 0.088 |
| corrected_fdxt | pair_326 | 0.004 | 2098.098 | 4.722 | 4.696 | 9.464 | 9.523 | 0.127 |
| corrected_fdxt | pair_327 | 0.004 | 2098.098 | 3.917 | 3.895 | 7.844 | 7.817 | 0.051 |
| corrected_fdxt | pair_328 | 0.004 | 2098.098 | 3.935 | 4.265 | 8.236 | 8.417 | 0.337 |
| corrected_fdxt | pair_329 | 0.004 | 2098.098 | 2.369 | 2.375 | 4.765 | 4.759 | 0.121 |
| corrected_fdxt | pair_330 | 0.004 | 2098.098 | 4.813 | 4.712 | 9.569 | 9.517 | 0.169 |
| corrected_fdxt | pair_331 | 0.004 | 2098.098 | 2.357 | 2.343 | 4.719 | 4.742 | 0.042 |
| corrected_fdxt | pair_332 | 0.004 | 2098.098 | 2.341 | 2.438 | 4.800 | 4.706 | 0.215 |
| corrected_fdxt | pair_333 | 0.004 | 2098.098 | 4.708 | 4.704 | 9.459 | 9.467 | 0.108 |
| corrected_fdxt | pair_334 | 0.004 | 2098.098 | 4.903 | 4.663 | 9.611 | 9.521 | 0.280 |
| corrected_fdxt | pair_335 | 0.004 | 2098.098 | 4.695 | 4.678 | 9.419 | 9.382 | 0.115 |
| corrected_fdxt | pair_336 | 0.004 | 2098.098 | 4.675 | 4.654 | 9.371 | 9.345 | 0.105 |
| corrected_fdxt | pair_337 | 0.004 | 2098.098 | 4.701 | 4.680 | 9.422 | 9.385 | 0.071 |
| corrected_fdxt | pair_338 | 0.004 | 2098.098 | 3.965 | 3.945 | 7.952 | 7.904 | 0.161 |
| corrected_fdxt | pair_339 | 0.004 | 2098.098 | 4.801 | 4.827 | 9.675 | 9.448 | 0.472 |
| corrected_fdxt | pair_340 | 0.004 | 2098.098 | 4.833 | 4.713 | 9.591 | 9.637 | 0.112 |
| corrected_fdxt | pair_341 | 0.004 | 2098.098 | 3.948 | 3.952 | 7.938 | 7.971 | 0.164 |
| corrected_fdxt | pair_342 | 0.004 | 2098.098 | 4.166 | 4.026 | 8.230 | 7.918 | 0.565 |
| corrected_fdxt | pair_343 | 0.004 | 2098.098 | 2.359 | 2.378 | 4.766 | 4.688 | 0.148 |
| corrected_fdxt | pair_344 | 0.004 | 2098.098 | 4.743 | 4.730 | 9.519 | 9.482 | 0.139 |
| corrected_fdxt | pair_345 | 0.004 | 2098.098 | 2.337 | 2.353 | 4.707 | 4.706 | 0.062 |
| corrected_fdxt | pair_346 | 0.004 | 2098.098 | 2.330 | 3.271 | 5.619 | 4.719 | 1.598 |
| corrected_fdxt | pair_347 | 0.004 | 2098.098 | 2.424 | 2.536 | 4.987 | 4.819 | 0.341 |
| corrected_fdxt | pair_348 | 0.004 | 2098.098 | 2.966 | 2.358 | 5.343 | 4.829 | 0.943 |
| corrected_fdxt | pair_349 | 0.004 | 2098.098 | 2.431 | 2.546 | 4.994 | 4.748 | 0.475 |
| corrected_fdxt | pair_350 | 0.004 | 2098.098 | 2.481 | 2.349 | 4.847 | 4.737 | 0.220 |
| corrected_fdxt | pair_351 | 0.004 | 2098.098 | 2.347 | 2.330 | 4.692 | 4.658 | 0.059 |
| corrected_fdxt | pair_352 | 0.004 | 2098.098 | 2.379 | 2.636 | 5.033 | 4.787 | 0.581 |
| corrected_fdxt | pair_353 | 0.004 | 2098.098 | 2.758 | 2.644 | 5.417 | 4.730 | 1.279 |
| corrected_fdxt | pair_354 | 0.004 | 2098.098 | 2.493 | 2.698 | 5.208 | 4.771 | 0.863 |
| corrected_fdxt | pair_355 | 0.004 | 2098.098 | 2.345 | 2.423 | 4.785 | 4.656 | 0.224 |
| corrected_fdxt | pair_356 | 0.004 | 2098.098 | 2.371 | 2.523 | 4.910 | 4.709 | 0.356 |
| corrected_fdxt | pair_357 | 0.004 | 2098.098 | 2.342 | 2.361 | 4.720 | 4.698 | 0.049 |
| corrected_fdxt | pair_358 | 0.004 | 2098.098 | 2.356 | 2.448 | 4.820 | 4.734 | 0.164 |
| corrected_fdxt | pair_359 | 0.004 | 2098.098 | 2.336 | 2.347 | 4.700 | 4.690 | 0.057 |
| corrected_fdxt | pair_360 | 0.004 | 2098.098 | 2.341 | 2.535 | 4.892 | 4.705 | 0.338 |
| corrected_fdxt | pair_361 | 0.004 | 2098.098 | 2.337 | 2.361 | 4.714 | 4.665 | 0.108 |
| corrected_fdxt | pair_362 | 0.004 | 2098.098 | 2.392 | 2.339 | 4.745 | 4.743 | 0.042 |
| corrected_fdxt | pair_363 | 0.004 | 2098.098 | 2.341 | 2.325 | 4.682 | 4.656 | 0.046 |
| corrected_fdxt | pair_364 | 0.004 | 2098.098 | 2.391 | 2.385 | 4.793 | 4.827 | 0.101 |
| corrected_fdxt | pair_365 | 0.004 | 2098.098 | 2.354 | 2.340 | 4.711 | 4.716 | 0.053 |
| corrected_fdxt | pair_366 | 0.004 | 2098.098 | 2.316 | 2.343 | 4.674 | 4.678 | 0.013 |
| corrected_fdxt | pair_367 | 0.004 | 2098.098 | 2.647 | 2.515 | 5.181 | 4.756 | 0.795 |
| corrected_fdxt | pair_368 | 0.004 | 2098.098 | 2.387 | 2.548 | 4.951 | 4.787 | 0.348 |
| corrected_fdxt | pair_369 | 0.004 | 2098.098 | 2.542 | 2.400 | 4.960 | 4.877 | 0.323 |
| corrected_fdxt | pair_370 | 0.004 | 2098.098 | 5.677 | 5.675 | 11.415 | 11.346 | 0.489 |
| corrected_fdxt | pair_371 | 0.004 | 2098.098 | 9.028 | 8.740 | 17.895 | 17.360 | 1.033 |
| corrected_fdxt | pair_372 | 0.004 | 2098.098 | 5.777 | 5.624 | 11.474 | 11.085 | 0.723 |
| corrected_fdxt | pair_373 | 0.004 | 2098.098 | 9.462 | 9.002 | 18.595 | 19.163 | 1.254 |
| corrected_fdxt | pair_374 | 0.004 | 2098.098 | 3.941 | 3.892 | 7.867 | 7.856 | 0.082 |
| corrected_fdxt | pair_375 | 0.004 | 2098.098 | 8.696 | 8.758 | 17.586 | 17.449 | 0.458 |
| corrected_fdxt | pair_376 | 0.004 | 2098.098 | 8.671 | 8.627 | 17.419 | 17.431 | 0.112 |
| corrected_fdxt | pair_377 | 0.004 | 2098.098 | 3.877 | 3.903 | 7.814 | 7.828 | 0.037 |
| corrected_fdxt | pair_378 | 0.004 | 2098.098 | 3.973 | 3.869 | 7.881 | 7.925 | 0.126 |
| corrected_fdxt | pair_379 | 0.004 | 2098.098 | 2.344 | 2.366 | 4.726 | 4.716 | 0.034 |
| corrected_fdxt | pair_380 | 0.004 | 2098.098 | 4.662 | 4.704 | 9.411 | 9.419 | 0.032 |
| corrected_fdxt | pair_381 | 0.004 | 2098.098 | 5.459 | 5.491 | 11.005 | 11.014 | 0.039 |
| corrected_fdxt | pair_382 | 0.004 | 2098.098 | 5.466 | 5.439 | 10.958 | 10.967 | 0.051 |
| corrected_fdxt | pair_383 | 0.004 | 2098.098 | 5.934 | 5.551 | 11.542 | 11.021 | 0.974 |
| corrected_fdxt | pair_384 | 0.004 | 2098.098 | 3.957 | 3.933 | 7.925 | 7.893 | 0.069 |
| corrected_fdxt | pair_385 | 0.004 | 2098.098 | 5.473 | 5.467 | 10.998 | 11.004 | 0.033 |
| corrected_fdxt | pair_386 | 0.004 | 2098.098 | 5.475 | 5.409 | 10.938 | 10.892 | 0.081 |
| corrected_fdxt | pair_387 | 0.004 | 2098.098 | 3.922 | 3.880 | 7.834 | 7.806 | 0.067 |
| corrected_fdxt | pair_388 | 0.004 | 2098.098 | 3.903 | 3.863 | 7.800 | 7.817 | 0.058 |
| corrected_fdxt | pair_389 | 0.004 | 2098.098 | 2.782 | 2.342 | 5.140 | 4.686 | 0.805 |
| corrected_fdxt | pair_390 | 0.004 | 2098.098 | 4.725 | 4.630 | 9.400 | 9.384 | 0.029 |
| corrected_fdxt | pair_391 | 0.004 | 2098.098 | 5.515 | 5.439 | 11.009 | 10.999 | 0.028 |
| corrected_fdxt | pair_392 | 0.004 | 2098.098 | 15.563 | 15.578 | 31.494 | 31.499 | 0.094 |
| corrected_fdxt | pair_393 | 0.004 | 2098.098 | 3.916 | 3.916 | 7.865 | 7.842 | 0.060 |
| corrected_fdxt | pair_394 | 0.004 | 2098.098 | 9.397 | 9.580 | 19.119 | 18.940 | 0.327 |
| corrected_fdxt | pair_395 | 0.004 | 2098.098 | 16.550 | 16.359 | 33.305 | 33.225 | 0.457 |
| corrected_fdxt | pair_396 | 0.004 | 2098.098 | 3.946 | 3.865 | 7.849 | 7.824 | 0.072 |
| corrected_fdxt | pair_397 | 0.004 | 2098.098 | 3.961 | 3.864 | 7.858 | 7.888 | 0.053 |
| corrected_fdxt | pair_398 | 0.004 | 2098.098 | 2.321 | 2.340 | 4.681 | 4.684 | 0.050 |
| corrected_fdxt | pair_399 | 0.004 | 2098.098 | 5.192 | 4.655 | 9.895 | 9.424 | 0.852 |
| corrected_fdxt | pair_400 | 0.004 | 2098.098 | 5.459 | 5.464 | 10.980 | 10.975 | 0.037 |
| corrected_fdxt | pair_401 | 0.004 | 2098.098 | 3.903 | 3.903 | 7.841 | 7.841 | 0.065 |
| corrected_fdxt | pair_402 | 0.004 | 2098.098 | 5.701 | 5.427 | 11.183 | 11.067 | 0.214 |
| corrected_fdxt | pair_403 | 0.004 | 2098.098 | 5.484 | 5.437 | 10.980 | 10.963 | 0.048 |
| corrected_fdxt | pair_404 | 0.004 | 2098.098 | 3.923 | 3.900 | 7.856 | 7.850 | 0.033 |
| corrected_fdxt | pair_405 | 0.004 | 2098.098 | 3.928 | 3.865 | 7.826 | 7.822 | 0.061 |
| corrected_fdxt | pair_406 | 0.004 | 2098.098 | 2.351 | 2.409 | 4.778 | 4.708 | 0.137 |
| corrected_fdxt | pair_407 | 0.004 | 2098.098 | 4.777 | 4.736 | 9.557 | 9.366 | 0.332 |
| corrected_fdxt | pair_408 | 0.004 | 2098.098 | 3.929 | 3.933 | 7.901 | 7.862 | 0.079 |
| corrected_fdxt | pair_409 | 0.004 | 2098.098 | 9.369 | 9.413 | 18.928 | 18.906 | 0.072 |
| corrected_fdxt | pair_410 | 0.004 | 2098.098 | 15.636 | 15.559 | 31.535 | 31.540 | 0.089 |
| corrected_fdxt | pair_411 | 0.004 | 2098.098 | 3.977 | 3.902 | 7.913 | 7.910 | 0.068 |
| corrected_fdxt | pair_412 | 0.004 | 2098.098 | 3.932 | 3.869 | 7.833 | 7.854 | 0.051 |
| corrected_fdxt | pair_413 | 0.004 | 2098.098 | 2.354 | 2.342 | 4.712 | 4.690 | 0.058 |
| corrected_fdxt | pair_414 | 0.004 | 2098.098 | 4.712 | 4.662 | 9.420 | 9.443 | 0.067 |
| corrected_fdxt | pair_415 | 0.004 | 2098.098 | 3.874 | 3.885 | 7.793 | 7.793 | 0.033 |
| corrected_fdxt | pair_416 | 0.004 | 2098.098 | 3.957 | 3.900 | 7.891 | 7.886 | 0.052 |
| corrected_fdxt | pair_417 | 0.004 | 2098.098 | 3.915 | 3.959 | 7.912 | 7.876 | 0.125 |
| corrected_fdxt | pair_418 | 0.004 | 2098.098 | 3.899 | 3.911 | 7.842 | 7.825 | 0.030 |
| corrected_fdxt | pair_419 | 0.004 | 2098.098 | 2.366 | 2.327 | 4.712 | 4.720 | 0.028 |
| corrected_fdxt | pair_420 | 0.004 | 2098.098 | 3.928 | 3.949 | 7.911 | 7.902 | 0.022 |
| corrected_fdxt | pair_421 | 0.004 | 2098.098 | 9.310 | 9.358 | 18.804 | 18.812 | 0.095 |
| corrected_fdxt | pair_422 | 0.004 | 2098.098 | 4.102 | 3.934 | 8.072 | 7.877 | 0.402 |
| corrected_fdxt | pair_423 | 0.004 | 2098.098 | 4.160 | 4.009 | 8.208 | 8.008 | 0.640 |
| corrected_fdxt | pair_424 | 0.004 | 2098.098 | 2.409 | 2.331 | 4.758 | 4.707 | 0.098 |
| corrected_fdxt | pair_425 | 0.004 | 2098.098 | 4.674 | 4.710 | 9.430 | 9.404 | 0.104 |
| corrected_fdxt | pair_426 | 0.004 | 2098.098 | 3.887 | 3.891 | 7.810 | 7.817 | 0.077 |
| corrected_fdxt | pair_427 | 0.004 | 2098.098 | 3.930 | 3.869 | 7.947 | 7.820 | 0.233 |
| corrected_fdxt | pair_428 | 0.004 | 2098.098 | 2.834 | 3.455 | 6.307 | 4.788 | 2.689 |
| corrected_fdxt | pair_429 | 0.004 | 2098.098 | 5.028 | 5.288 | 10.365 | 9.391 | 1.705 |
| corrected_fdxt | pair_430 | 0.004 | 2098.098 | 4.159 | 3.915 | 8.109 | 7.862 | 0.433 |
| corrected_fdxt | pair_431 | 0.004 | 2098.098 | 2.349 | 2.358 | 4.723 | 4.678 | 0.109 |
| corrected_fdxt | pair_432 | 0.004 | 2098.098 | 3.899 | 3.877 | 7.815 | 7.790 | 0.043 |
| corrected_fdxt | pair_433 | 0.004 | 2098.098 | 2.331 | 2.325 | 4.672 | 4.674 | 0.008 |
| corrected_fdxt | pair_434 | 0.004 | 2098.098 | 3.966 | 3.955 | 7.956 | 7.831 | 0.229 |
| corrected_fdxt | pair_435 | 0.004 | 2098.098 | 2.360 | 2.356 | 4.732 | 4.737 | 0.016 |

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
