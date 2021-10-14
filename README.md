# benchy

Golang benchmarks

[ToC]

## jwt

github.com/riftbit/go-benchy/jwt

```
go test -benchmem -cpu 8 -benchtime=1s -timeout 30m -bench=. github.com/riftbit/go-benchy/jwt
```

```
goos: darwin
goarch: amd64
pkg: github.com/riftbit/go-benchy/jwt
cpu: Intel(R) Core(TM) i7-4960HQ CPU @ 2.60GHz
BenchmarkJWT_Create_HS256_String_Empty_Parallel/cristalhq_jwt-8                   677205              1597 ns/op            1908 B/op         18 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/kataras_jwt-8                     661843              1952 ns/op            2879 B/op         17 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/gbrlsnchs_jwt-8                   642675              2224 ns/op            1892 B/op         20 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/golang-jwt_jwt-8                  391484              3458 ns/op            4447 B/op         47 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/pascaldekloe_jwt-8                252896              5480 ns/op            6352 B/op         67 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/lestrrat-go_jwx-8                 181424              7227 ns/op            7247 B/op        112 allocs/op
BenchmarkJWT_Create_HS256_String_Empty_Parallel/brianvoe_sjwt-8                   165558              8031 ns/op            8589 B/op        106 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/cristalhq_jwt-8                    612157              1765 ns/op            1459 B/op         17 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/kataras_jwt-8                      653535              2199 ns/op            2431 B/op         16 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/gbrlsnchs_jwt-8                    695529              1972 ns/op            1443 B/op         19 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/golang-jwt_jwt-8                   357440              3822 ns/op            4896 B/op         48 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/pascaldekloe_jwt-8                 251233              5426 ns/op            5775 B/op         66 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/lestrrat-go_jwx-8                  185168              7642 ns/op            6669 B/op        111 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty_Parallel/brianvoe_sjwt-8                    144860              7367 ns/op            9039 B/op        107 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_String_Filled_Parallel/cristalhq_jwt-8                    2307            578580 ns/op         1076315 B/op         28 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/kataras_jwt-8                      2235            627947 ns/op         1846147 B/op         27 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/gbrlsnchs_jwt-8                    2352            536527 ns/op         1102050 B/op         29 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/golang-jwt_jwt-8                   1952            616412 ns/op         2134593 B/op         56 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/pascaldekloe_jwt-8                 2431            583891 ns/op         1184220 B/op         80 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/lestrrat-go_jwx-8                  1352           1023171 ns/op         2677793 B/op        132 allocs/op
BenchmarkJWT_Create_HS256_String_Filled_Parallel/brianvoe_sjwt-8                     907           1331361 ns/op         3403252 B/op        379 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/cristalhq_jwt-8                     2576            528062 ns/op          703430 B/op         23 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/kataras_jwt-8                       2167            615565 ns/op         1488835 B/op         25 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/gbrlsnchs_jwt-8                     2524            502466 ns/op          793151 B/op         27 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/golang-jwt_jwt-8                    1790            643380 ns/op         2449036 B/op         57 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/pascaldekloe_jwt-8                  2594            540766 ns/op          779342 B/op         77 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/lestrrat-go_jwx-8                   1315           1095949 ns/op         2331736 B/op        130 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled_Parallel/brianvoe_sjwt-8                      991           1447677 ns/op         3736817 B/op        380 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_String_Empty/cristalhq_jwt-8                            230805              5082 ns/op            1906 B/op         18 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/kataras_jwt-8                              240042              5021 ns/op            2878 B/op         17 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/gbrlsnchs_jwt-8                            204442              5810 ns/op            1890 B/op         20 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/golang-jwt_jwt-8                           144945              8290 ns/op            4443 B/op         47 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/pascaldekloe_jwt-8                          89580             12979 ns/op            6346 B/op         67 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/lestrrat-go_jwx-8                           68577             17424 ns/op            7234 B/op        112 allocs/op
BenchmarkJWT_Create_HS256_String_Empty/brianvoe_sjwt-8                             68838             17962 ns/op            8579 B/op        106 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_Bytes_Empty/cristalhq_jwt-8                             232214              5499 ns/op            1458 B/op         17 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/kataras_jwt-8                               183043              5566 ns/op            2429 B/op         16 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/gbrlsnchs_jwt-8                             208639              5997 ns/op            1442 B/op         19 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/golang-jwt_jwt-8                            130406              8904 ns/op            4892 B/op         48 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/pascaldekloe_jwt-8                           89032             12877 ns/op            5769 B/op         66 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/lestrrat-go_jwx-8                            69428             17343 ns/op            6658 B/op        111 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Empty/brianvoe_sjwt-8                              69170             16967 ns/op            9028 B/op        107 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_String_Filled/cristalhq_jwt-8                              819           1435493 ns/op          976619 B/op         25 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/kataras_jwt-8                                756           1575589 ns/op         1595363 B/op         21 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/gbrlsnchs_jwt-8                              820           1432677 ns/op          962177 B/op         25 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/golang-jwt_jwt-8                             699           1699802 ns/op         2009891 B/op         55 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/pascaldekloe_jwt-8                           835           1416294 ns/op          927285 B/op         76 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/lestrrat-go_jwx-8                            410           2897394 ns/op         2379544 B/op        130 allocs/op
BenchmarkJWT_Create_HS256_String_Filled/brianvoe_sjwt-8                              362           3267921 ns/op         3033101 B/op        371 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_HS256_Bytes_Filled/cristalhq_jwt-8                               846           1365822 ns/op          625826 B/op         21 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/kataras_jwt-8                                 787           1485125 ns/op         1233633 B/op         19 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/gbrlsnchs_jwt-8                               852           1372203 ns/op          628462 B/op         22 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/golang-jwt_jwt-8                              668           1750591 ns/op         2358505 B/op         58 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/pascaldekloe_jwt-8                            870           1400016 ns/op          613979 B/op         74 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/lestrrat-go_jwx-8                             417           2809138 ns/op         2013133 B/op        126 allocs/op
BenchmarkJWT_Create_HS256_Bytes_Filled/brianvoe_sjwt-8                               358           3344122 ns/op         3396427 B/op        373 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_String_Empty_Parallel/cristalhq_jwt-8                     2415            461881 ns/op           33163 B/op        128 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/kataras_jwt-8                       2444            517255 ns/op           33909 B/op        120 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/gbrlsnchs_jwt-8                     2204            483819 ns/op           33021 B/op        129 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/golang-jwt_jwt-8                    2208            554211 ns/op           35504 B/op        151 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/pascaldekloe_jwt-8                  2266            496044 ns/op           37122 B/op        172 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/lestrrat-go_jwx-8                   2241            506258 ns/op           38634 B/op        215 allocs/op
BenchmarkJWT_Create_RS256_String_Empty_Parallel/brianvoe_sjwt-8                        1        2004514899 ns/op           16784 B/op        145 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/cristalhq_jwt-8                      2422            458063 ns/op           32390 B/op        127 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/kataras_jwt-8                        2230            492953 ns/op           33138 B/op        119 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/gbrlsnchs_jwt-8                      2288            517039 ns/op           32251 B/op        128 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/golang-jwt_jwt-8                     2152            542483 ns/op           36273 B/op        152 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/pascaldekloe_jwt-8                   2428            494983 ns/op           36227 B/op        171 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/lestrrat-go_jwx-8                    2259            510958 ns/op           37723 B/op        214 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty_Parallel/brianvoe_sjwt-8                         1        2001021696 ns/op           14960 B/op        143 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_String_Filled_Parallel/cristalhq_jwt-8                    1214           1013365 ns/op         1207312 B/op        146 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/kataras_jwt-8                      1012           1306944 ns/op         2020628 B/op        145 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/gbrlsnchs_jwt-8                    1177           1211326 ns/op         1322027 B/op        150 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/golang-jwt_jwt-8                    924           1446613 ns/op         2354210 B/op        176 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/pascaldekloe_jwt-8                 1119           1228906 ns/op         1362016 B/op        195 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/lestrrat-go_jwx-8                   794           1715866 ns/op         3147734 B/op        247 allocs/op
BenchmarkJWT_Create_RS256_String_Filled_Parallel/brianvoe_sjwt-8                       1        2006606369 ns/op         3493080 B/op        406 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/cristalhq_jwt-8                     1309            931511 ns/op          797680 B/op        138 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/kataras_jwt-8                       1114           1295243 ns/op         1669396 B/op        142 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/gbrlsnchs_jwt-8                     1276           1052957 ns/op          867403 B/op        141 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/golang-jwt_jwt-8                    1045           1377791 ns/op         2639734 B/op        176 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/pascaldekloe_jwt-8                  1198           1110751 ns/op         1016451 B/op        191 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/lestrrat-go_jwx-8                    895           1543570 ns/op         2566073 B/op        240 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled_Parallel/brianvoe_sjwt-8                        1        2006455002 ns/op         3291856 B/op        394 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_String_Empty/cristalhq_jwt-8                               625           1871093 ns/op           33098 B/op        127 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/kataras_jwt-8                                 630           1857430 ns/op           33868 B/op        120 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/gbrlsnchs_jwt-8                               634           1853934 ns/op           32962 B/op        128 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/golang-jwt_jwt-8                              631           1919964 ns/op           35428 B/op        150 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/pascaldekloe_jwt-8                            588           1964041 ns/op           37081 B/op        172 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/lestrrat-go_jwx-8                             600           2005797 ns/op           38536 B/op        214 allocs/op
BenchmarkJWT_Create_RS256_String_Empty/brianvoe_sjwt-8                                 1        2000645689 ns/op           13008 B/op        115 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_Bytes_Empty/cristalhq_jwt-8                                608           1954014 ns/op           32343 B/op        126 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/kataras_jwt-8                                  620           1905134 ns/op           33109 B/op        119 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/gbrlsnchs_jwt-8                                607           1951283 ns/op           32198 B/op        127 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/golang-jwt_jwt-8                               573           1975158 ns/op           36200 B/op        151 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/pascaldekloe_jwt-8                             585           2010715 ns/op           36163 B/op        170 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/lestrrat-go_jwx-8                              601           2061420 ns/op           37640 B/op        213 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Empty/brianvoe_sjwt-8                                  1        2005324013 ns/op           14240 B/op        121 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_String_Filled/cristalhq_jwt-8                              349           3447641 ns/op          958669 B/op        132 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/kataras_jwt-8                                330           3620872 ns/op         1656462 B/op        128 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/gbrlsnchs_jwt-8                              352           3473367 ns/op          992530 B/op        134 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/golang-jwt_jwt-8                             316           3778908 ns/op         2075521 B/op        164 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/pascaldekloe_jwt-8                           358           3378262 ns/op          975096 B/op        182 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/lestrrat-go_jwx-8                            242           5067636 ns/op         2457851 B/op        238 allocs/op
BenchmarkJWT_Create_RS256_String_Filled/brianvoe_sjwt-8                                1        2006222069 ns/op         2971648 B/op        373 allocs/op
===========================================================================================================================================================
BenchmarkJWT_Create_RS256_Bytes_Filled/cristalhq_jwt-8                               351           3323940 ns/op          627311 B/op        130 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/kataras_jwt-8                                 330           3562200 ns/op         1299061 B/op        125 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/gbrlsnchs_jwt-8                               332           3372268 ns/op          672592 B/op        132 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/golang-jwt_jwt-8                              314           3834987 ns/op         2399321 B/op        167 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/pascaldekloe_jwt-8                            354           3354720 ns/op          649263 B/op        180 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/lestrrat-go_jwx-8                             235           5080096 ns/op         2192718 B/op        236 allocs/op
BenchmarkJWT_Create_RS256_Bytes_Filled/brianvoe_sjwt-8                                 1        2006394635 ns/op         3907208 B/op        384 allocs/op
===========================================================================================================================================================
PASS
ok      github.com/riftbit/go-benchy/jwt        163.485s
```

## str_concat

github.com/riftbit/go-benchy/str_concat

```
go test -benchmem -cpu 8 -benchtime=5s -timeout 30m -bench=. github.com/riftbit/go-benchy/str_concat
```

```
goos: darwin
goarch: amd64
pkg: github.com/riftbit/go-benchy/str_concat
cpu: Intel(R) Core(TM) i7-4960HQ CPU @ 2.60GHz
Benchmark_Pluses-8                      63049274                95.07 ns/op          288 B/op          1 allocs/op
Benchmark_Pluses_Parallel-8             99938961                59.20 ns/op          288 B/op          1 allocs/op
Benchmark_PlusesFromEnv-8               55123735               108.6 ns/op           320 B/op          1 allocs/op
Benchmark_PlusesFromEnv_Parallel-8      82028496                71.09 ns/op          320 B/op          1 allocs/op
Benchmark_Fmt-8                         20146476               250.4 ns/op           304 B/op          2 allocs/op
Benchmark_Fmt_Parallel-8                50952112               146.3 ns/op           304 B/op          2 allocs/op
Benchmark_StringBuilder-8               40346128               180.6 ns/op           368 B/op          4 allocs/op
Benchmark_StringBuilder_Parallel-8      80789652                93.16 ns/op          328 B/op          3 allocs/op
Benchmark_ByteBufferPool-8              52562928               111.6 ns/op           288 B/op          1 allocs/op
Benchmark_ByteBufferPool_Parallel-8     76494380                97.75 ns/op          288 B/op          1 allocs/op
PASS
ok      github.com/riftbit/go-benchy/str_concat 65.870s
```


## idgen

github.com/riftbit/go-benchy/idgen

```
go test -benchmem -cpu 8 -benchtime=5s -timeout 30m -bench=. github.com/riftbit/go-benchy/idgen
```

```
goos: darwin
goarch: amd64
pkg: github.com/riftbit/go-benchy/idgen
cpu: Intel(R) Core(TM) i7-4960HQ CPU @ 2.60GHz
BenchmarkSnowflake_CreateID_Parallel/sonyflake-8                  154032             39081 ns/op               0 B/op          0 allocs/op
BenchmarkSnowflake_CreateID_Parallel/shaxbee-8                  13601421               434.3 ns/op             0 B/op          0 allocs/op
BenchmarkSnowflake_CreateID_Parallel/piaohua-8                  18488851               316.2 ns/op             1 B/op          0 allocs/op
BenchmarkSnowflake_CreateID_Parallel/bwmarrin-8                 24612841               244.1 ns/op             0 B/op          0 allocs/op
BenchmarkSnowflake_CreateID_Parallel/godruoyi-8                 24573907               246.5 ns/op             0 B/op          0 allocs/op
BenchmarkSnowflake_CreateID_Parallel/rs-8                       224751864               26.55 ns/op            0 B/op          0 allocs/op
BenchmarkSnowflake_CreateID_Parallel/mongo-8                    225356929               26.23 ns/op            0 B/op          0 allocs/op
===========================================================================================================================================================
BenchmarkSnowflake_CreateHex_Parallel/sonyflake-8                 155542             39079 ns/op               8 B/op          1 allocs/op
BenchmarkSnowflake_CreateHex_Parallel/shaxbee-8                 13283817               452.5 ns/op             8 B/op          1 allocs/op
BenchmarkSnowflake_CreateHex_Parallel/piaohua-8                 13795723               388.8 ns/op            50 B/op          2 allocs/op
BenchmarkSnowflake_CreateHex_Parallel/bwmarrin-8                23571288               272.0 ns/op            48 B/op          2 allocs/op
BenchmarkSnowflake_CreateHex_Parallel/godruoyi-8                24522546               244.3 ns/op             8 B/op          1 allocs/op
BenchmarkSnowflake_CreateHex_Parallel/rs-8                      159841766               37.48 ns/op           16 B/op          1 allocs/op
BenchmarkSnowflake_CreateHex_Parallel/mongo-8                   156765771               38.02 ns/op           16 B/op          1 allocs/op
===========================================================================================================================================================
BenchmarkSnowflake_CreateID/sonyflake-8                           154483             39040 ns/op               0 B/op          0 allocs/op
BenchmarkSnowflake_CreateID/shaxbee-8                           19019941               307.2 ns/op             0 B/op          0 allocs/op
BenchmarkSnowflake_CreateID/piaohua-8                           13341760               441.7 ns/op             1 B/op          0 allocs/op
BenchmarkSnowflake_CreateID/bwmarrin-8                          24635056               244.3 ns/op             0 B/op          0 allocs/op
BenchmarkSnowflake_CreateID/godruoyi-8                          24631790               244.3 ns/op             0 B/op          0 allocs/op
BenchmarkSnowflake_CreateID/rs-8                                64241464                92.02 ns/op            0 B/op          0 allocs/op
BenchmarkSnowflake_CreateID/mongo-8                             62055952                91.24 ns/op            0 B/op          0 allocs/op
===========================================================================================================================================================
BenchmarkSnowflake_CreateHex/sonyflake-8                          156964             39042 ns/op               8 B/op          1 allocs/op
BenchmarkSnowflake_CreateHex/shaxbee-8                          16448815               364.3 ns/op             8 B/op          1 allocs/op
BenchmarkSnowflake_CreateHex/piaohua-8                          17711274               341.2 ns/op            49 B/op          2 allocs/op
BenchmarkSnowflake_CreateHex/bwmarrin-8                         24630088               244.2 ns/op            48 B/op          2 allocs/op
BenchmarkSnowflake_CreateHex/godruoyi-8                         24621992               244.2 ns/op             8 B/op          1 allocs/op
BenchmarkSnowflake_CreateHex/rs-8                               53103302               112.7 ns/op            16 B/op          1 allocs/op
BenchmarkSnowflake_CreateHex/mongo-8                            53630092               109.6 ns/op            16 B/op          1 allocs/op
===========================================================================================================================================================
PASS
ok      github.com/riftbit/go-benchy/idgen      187.393s
```