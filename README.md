# benchy

Golang benchmarks

[ToC]

## jwt

github.com/rb-pkg/benchy/jwt

```
go test -benchmem -cpu 8 -benchtime=5s -bench=. github.com/rb-pkg/benchy/jwt
```

```
goos: windows
goarch: amd64
pkg: github.com/rb-pkg/benchy/jwt
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkJWT_Create_HS256_String_Empty_Parallel/cristalhq_jwt-8                  6729601               865.5 ns/op          1908 B/op         18 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/kataras_jwt-8                    6470664               976.7 ns/op          2879 B/op         17 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/gbrlsnchs_jwt-8                  6585448               997.2 ns/op          1892 B/op         20 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/golang-jwt_jwt-8                 3410836              1796 ns/op            4447 B/op         47 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/pascaldekloe_jwt-8               2131717              2744 ns/op            6353 B/op         67 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/lestrrat-go_jwx-8                1346372              4624 ns/op            7560 B/op        117 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/brianvoe_sjwt-8                  1267494              4440 ns/op            8591 B/op        106 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/cristalhq_jwt-8                   6336261              1032 ns/op            1459 B/op         17 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/kataras_jwt-8                     5880807              1070 ns/op            2430 B/op         16 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/gbrlsnchs_jwt-8                   5093265               995.3 ns/op          1443 B/op         19 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/golang-jwt_jwt-8                  2950632              1998 ns/op            4896 B/op         48 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/pascaldekloe_jwt-8                2235133              2864 ns/op            5776 B/op         66 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/lestrrat-go_jwx-8                 1369180              4071 ns/op            6982 B/op        116 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/brianvoe_sjwt-8                   1375792              4815 ns/op            9041 B/op        107 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_String_Filled_Parallel/cristalhq_jwt-8                   24813            236007 ns/op          977493 B/op         22 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/kataras_jwt-8                     21600            285153 ns/op         1654266 B/op         21 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/gbrlsnchs_jwt-8                   25327            256683 ns/op         1007053 B/op         24 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/golang-jwt_jwt-8                  19360            319405 ns/op         1994945 B/op         51 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/pascaldekloe_jwt-8                21698            280484 ns/op         1015266 B/op         76 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/lestrrat-go_jwx-8                 10000            522498 ns/op         2410444 B/op        130 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/brianvoe_sjwt-8                    9488            676205 ns/op         3240924 B/op        373 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/cristalhq_jwt-8                    24081            230737 ns/op          643649 B/op         20 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/kataras_jwt-8                      21697            275052 ns/op         1304034 B/op         19 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/gbrlsnchs_jwt-8                    25261            239621 ns/op          666367 B/op         22 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/golang-jwt_jwt-8                   18718            335724 ns/op         2330094 B/op         53 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/pascaldekloe_jwt-8                 22484            262326 ns/op          669935 B/op         74 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/lestrrat-go_jwx-8                  10000            526618 ns/op         2063442 B/op        129 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/brianvoe_sjwt-8                     9442            712270 ns/op         3556852 B/op        373 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_String_Empty/cristalhq_jwt-8                           1441412              4188 ns/op            1907 B/op         18 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/kataras_jwt-8                             1577556              3803 ns/op            2878 B/op         17 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/gbrlsnchs_jwt-8                           1363226              4508 ns/op            1891 B/op         20 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/golang-jwt_jwt-8                           986400              6559 ns/op            4444 B/op         47 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/pascaldekloe_jwt-8                         571303             10701 ns/op            6347 B/op         67 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/lestrrat-go_jwx-8                          303213             18292 ns/op            7551 B/op        117 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/brianvoe_sjwt-8                            361449             13943 ns/op            8583 B/op        106 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_Bytes_Empty/cristalhq_jwt-8                            1493443              3949 ns/op            1458 B/op         17 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/kataras_jwt-8                              1608156              3817 ns/op            2430 B/op         16 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/gbrlsnchs_jwt-8                            1357873              4503 ns/op            1442 B/op         19 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/golang-jwt_jwt-8                            973845              6874 ns/op            4892 B/op         48 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/pascaldekloe_jwt-8                          490125             10412 ns/op            5770 B/op         66 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/lestrrat-go_jwx-8                           345258             18560 ns/op            6975 B/op        116 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/brianvoe_sjwt-8                             367647             14401 ns/op            9031 B/op        107 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_String_Filled/cristalhq_jwt-8                             4881           1261939 ns/op         1004775 B/op         26 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/kataras_jwt-8                               4064           1336669 ns/op         1668509 B/op         24 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/gbrlsnchs_jwt-8                             4528           1258238 ns/op         1073980 B/op         28 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/golang-jwt_jwt-8                            4293           1362127 ns/op         2038962 B/op         56 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/pascaldekloe_jwt-8                          5118           1193201 ns/op          997252 B/op         77 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/lestrrat-go_jwx-8                           2554           2313033 ns/op         2521889 B/op        139 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/brianvoe_sjwt-8                             2060           2729500 ns/op         3148824 B/op        374 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_Bytes_Filled/cristalhq_jwt-8                              5350           1158534 ns/op          649513 B/op         22 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/kataras_jwt-8                                4981           1276562 ns/op         1329520 B/op         21 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/gbrlsnchs_jwt-8                              4660           1165868 ns/op          719283 B/op         26 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/golang-jwt_jwt-8                             4072           1399263 ns/op         2386729 B/op         58 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/pascaldekloe_jwt-8                           4712           1154428 ns/op          627386 B/op         74 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/lestrrat-go_jwx-8                            2727           2326986 ns/op         2151910 B/op        136 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/brianvoe_sjwt-8                              2223           2820559 ns/op         3600349 B/op        380 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_String_Empty_Parallel/cristalhq_jwt-8                    29467            204022 ns/op           33163 B/op        128 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/kataras_jwt-8                      29823            201180 ns/op           33916 B/op        120 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/gbrlsnchs_jwt-8                    29258            201040 ns/op           33027 B/op        129 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/golang-jwt_jwt-8                   28426            211522 ns/op           35503 B/op        151 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/pascaldekloe_jwt-8                 28802            233035 ns/op           37142 B/op        172 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/lestrrat-go_jwx-8                  26888            222327 ns/op           38960 B/op        220 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/brianvoe_sjwt-8                       20         301496260 ns/op            8930 B/op        109 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/cristalhq_jwt-8                     32690            195981 ns/op           32392 B/op        127 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/kataras_jwt-8                       29587            200569 ns/op           33146 B/op        119 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/gbrlsnchs_jwt-8                     28492            192163 ns/op           32256 B/op        128 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/golang-jwt_jwt-8                    29641            203710 ns/op           36278 B/op        152 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/pascaldekloe_jwt-8                  29700            207809 ns/op           36241 B/op        171 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/lestrrat-go_jwx-8                   29258            206133 ns/op           38059 B/op        219 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/brianvoe_sjwt-8                        20         301554505 ns/op            9478 B/op        110 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_String_Filled_Parallel/cristalhq_jwt-8                   12082            525217 ns/op         1107564 B/op        139 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/kataras_jwt-8                     10000            600749 ns/op         1858422 B/op        134 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/gbrlsnchs_jwt-8                    9024            577765 ns/op         1173396 B/op        141 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/golang-jwt_jwt-8                  10000            603977 ns/op         2172226 B/op        164 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/pascaldekloe_jwt-8                10000            571476 ns/op         1244908 B/op        189 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/lestrrat-go_jwx-8                  8170            712526 ns/op         2825931 B/op        244 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/brianvoe_sjwt-8                      20         301997285 ns/op         3441898 B/op        382 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/cristalhq_jwt-8                    12314            483560 ns/op          770873 B/op        135 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/kataras_jwt-8                      10000            583137 ns/op         1524677 B/op        132 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/gbrlsnchs_jwt-8                    10000            519083 ns/op          815306 B/op        137 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/golang-jwt_jwt-8                    9609            617743 ns/op         2509426 B/op        166 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/pascaldekloe_jwt-8                 10000            528734 ns/op          881399 B/op        186 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/lestrrat-go_jwx-8                   8856            724964 ns/op         2433817 B/op        241 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/brianvoe_sjwt-8                       20         301739380 ns/op         3844042 B/op        385 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_String_Empty/cristalhq_jwt-8                              4864           1249585 ns/op           33104 B/op        127 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/kataras_jwt-8                                5239           1250540 ns/op           33871 B/op        120 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/gbrlsnchs_jwt-8                              5103           1279207 ns/op           32963 B/op        128 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/golang-jwt_jwt-8                             5101           1238380 ns/op           35426 B/op        150 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/pascaldekloe_jwt-8                           4848           1327163 ns/op           37067 B/op        171 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/lestrrat-go_jwx-8                            5025           1230861 ns/op           38841 B/op        219 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/brianvoe_sjwt-8                                 3        2004140267 ns/op           11280 B/op        115 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_Bytes_Empty/cristalhq_jwt-8                               5094           1227387 ns/op           32336 B/op        126 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/kataras_jwt-8                                 5125           1224969 ns/op           33098 B/op        119 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/gbrlsnchs_jwt-8                               5083           1238822 ns/op           32194 B/op        127 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/golang-jwt_jwt-8                              5042           1232718 ns/op           36198 B/op        151 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/pascaldekloe_jwt-8                            4914           1276245 ns/op           36161 B/op        170 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/lestrrat-go_jwx-8                             5031           1285528 ns/op           37939 B/op        218 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/brianvoe_sjwt-8                                  3        2008724400 ns/op           10170 B/op        109 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_String_Filled/cristalhq_jwt-8                             2380           2581383 ns/op         1076561 B/op        136 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/kataras_jwt-8                               2278           2727765 ns/op         1719228 B/op        132 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/gbrlsnchs_jwt-8                             2222           2543487 ns/op         1203935 B/op        141 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/golang-jwt_jwt-8                            2272           2750685 ns/op         2127165 B/op        166 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/pascaldekloe_jwt-8                          2413           2599030 ns/op         1011657 B/op        183 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/lestrrat-go_jwx-8                           1622           3922579 ns/op         2756870 B/op        249 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/brianvoe_sjwt-8                                3        2008628833 ns/op         3492704 B/op        385 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_Bytes_Filled/cristalhq_jwt-8                              2336           2389406 ns/op          686019 B/op        132 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/kataras_jwt-8                                2175           2552903 ns/op         1392358 B/op        129 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/gbrlsnchs_jwt-8                              2572           2402060 ns/op          713277 B/op        134 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/golang-jwt_jwt-8                             2226           2802241 ns/op         2449529 B/op        169 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/pascaldekloe_jwt-8                           2551           2416418 ns/op          647124 B/op        180 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/lestrrat-go_jwx-8                            1585           3767040 ns/op         2430886 B/op        245 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/brianvoe_sjwt-8                                 3        2009364800 ns/op         3812690 B/op        387 allocs/op
===========================================================================================================================================================
PASS
ok      github.com/rb-pkg/benchy/jwt    892.727s
```

## str_concat

github.com/rb-pkg/benchy/str_concat

```
go test -benchmem -cpu 8 -benchtime=5s -bench=. github.com/rb-pkg/benchy/str_concat
```

```
goos: windows
goarch: amd64
pkg: github.com/rb-pkg/benchy/str_concat
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
Benchmark_Pluses-8                      76898232                75.75 ns/op          288 B/op          1 allocs/op
Benchmark_Pluses_Parallel-8             90836283                59.25 ns/op          288 B/op          1 allocs/op
Benchmark_PlusesFromEnv-8               77947891                85.26 ns/op          320 B/op          1 allocs/op
Benchmark_PlusesFromEnv_Parallel-8      83152936                63.49 ns/op          320 B/op          1 allocs/op
Benchmark_Fmt-8                         31906712               204.0 ns/op           304 B/op          2 allocs/op
Benchmark_Fmt_Parallel-8                68297457                94.86 ns/op          304 B/op          2 allocs/op
Benchmark_StringBuilder-8               54226014               152.0 ns/op           368 B/op          4 allocs/op
Benchmark_StringBuilder_Parallel-8      69310476                72.26 ns/op          328 B/op          3 allocs/op
Benchmark_ByteBufferPool-8              63559254               104.2 ns/op           288 B/op          1 allocs/op
Benchmark_ByteBufferPool_Parallel-8     73387044                77.55 ns/op          288 B/op          1 allocs/op
PASS
ok      github.com/rb-pkg/benchy/str_concat     62.717s
```