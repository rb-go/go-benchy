# benchy

Golang benchmarks

[ToC]

## jwt

github.com/rb-pkg/benchy/jwt

```
go test -benchmem -cpu 8 -benchtime=10s -bench=. github.com/rb-pkg/benchy/jwt
```

```
goos: windows
goarch: amd64
pkg: github.com/rb-pkg/benchy/jwt
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkJWT_Create_String_Empty_Parallel/golang-jwt_jwt-8               3278571              1876 ns/op            4447 B/op         47 allocs/op
BenchmarkJWT_Create_String_Empty_Parallel/lestrrat-go_jwx-8              1607312              3870 ns/op            7559 B/op        117 allocs/op
BenchmarkJWT_Create_String_Empty_Parallel/gbrlsnchs_jwt-8                5460390               918.6 ns/op          1892 B/op         20 allocs/op
BenchmarkJWT_Create_String_Empty_Parallel/pascaldekloe_jwt-8             2478529              2563 ns/op            6353 B/op         67 allocs/op
BenchmarkJWT_Create_String_Empty_Parallel/brianvoe_sjwt-8                1783620              3739 ns/op            8591 B/op        106 allocs/op
BenchmarkJWT_Create_String_Empty_Parallel/cristalhq_jwt-8                5966274              1020 ns/op            1908 B/op         18 allocs/op
======================================================================================================================================================
BenchmarkJWT_Create_Bytes_Empty_Parallel/golang-jwt_jwt-8                2887369              1984 ns/op            4896 B/op         48 allocs/op
BenchmarkJWT_Create_Bytes_Empty_Parallel/lestrrat-go_jwx-8               1644874              3713 ns/op            6982 B/op        116 allocs/op
BenchmarkJWT_Create_Bytes_Empty_Parallel/gbrlsnchs_jwt-8                 5982262               988.6 ns/op          1443 B/op         19 allocs/op
BenchmarkJWT_Create_Bytes_Empty_Parallel/pascaldekloe_jwt-8              2252778              2694 ns/op            5776 B/op         66 allocs/op
BenchmarkJWT_Create_Bytes_Empty_Parallel/brianvoe_sjwt-8                 1450950              4422 ns/op            9040 B/op        107 allocs/op
BenchmarkJWT_Create_Bytes_Empty_Parallel/cristalhq_jwt-8                 7715965               959.3 ns/op          1459 B/op         17 allocs/op
======================================================================================================================================================
BenchmarkJWT_Create_String_Filled_Parallel/golang-jwt_jwt-8                18264            299187 ns/op         2001539 B/op         51 allocs/op
BenchmarkJWT_Create_String_Filled_Parallel/lestrrat-go_jwx-8               13263            499470 ns/op         2409402 B/op        130 allocs/op
BenchmarkJWT_Create_String_Filled_Parallel/gbrlsnchs_jwt-8                 25752            264798 ns/op         1000723 B/op         24 allocs/op
BenchmarkJWT_Create_String_Filled_Parallel/pascaldekloe_jwt-8              23678            271481 ns/op         1019753 B/op         76 allocs/op
BenchmarkJWT_Create_String_Filled_Parallel/brianvoe_sjwt-8                  9652            603288 ns/op         3254096 B/op        373 allocs/op
BenchmarkJWT_Create_String_Filled_Parallel/cristalhq_jwt-8                 25813            241166 ns/op          984360 B/op         22 allocs/op
======================================================================================================================================================
BenchmarkJWT_Create_Bytes_Filled_Parallel/golang-jwt_jwt-8                 19153            314752 ns/op         2353052 B/op         53 allocs/op
BenchmarkJWT_Create_Bytes_Filled_Parallel/lestrrat-go_jwx-8                13466            442402 ns/op         2082176 B/op        129 allocs/op
BenchmarkJWT_Create_Bytes_Filled_Parallel/gbrlsnchs_jwt-8                  25282            217014 ns/op          673543 B/op         22 allocs/op
BenchmarkJWT_Create_Bytes_Filled_Parallel/pascaldekloe_jwt-8               23976            252375 ns/op          676713 B/op         74 allocs/op
BenchmarkJWT_Create_Bytes_Filled_Parallel/brianvoe_sjwt-8                   9177            659257 ns/op         3581393 B/op        374 allocs/op
BenchmarkJWT_Create_Bytes_Filled_Parallel/cristalhq_jwt-8                  26420            228148 ns/op          650399 B/op         20 allocs/op
======================================================================================================================================================
BenchmarkJWT_Create_String_Empty/golang-jwt_jwt-8                         945513              6551 ns/op            4444 B/op         47 allocs/op
BenchmarkJWT_Create_String_Empty/lestrrat-go_jwx-8                        353697             17987 ns/op            7552 B/op        117 allocs/op
BenchmarkJWT_Create_String_Empty/gbrlsnchs_jwt-8                         1340521              4564 ns/op            1891 B/op         20 allocs/op
BenchmarkJWT_Create_String_Empty/pascaldekloe_jwt-8                       577951             10485 ns/op            6347 B/op         67 allocs/op
BenchmarkJWT_Create_String_Empty/brianvoe_sjwt-8                          476108             14252 ns/op            8582 B/op        106 allocs/op
BenchmarkJWT_Create_String_Empty/cristalhq_jwt-8                         1518392              4061 ns/op            1907 B/op         18 allocs/op
======================================================================================================================================================
BenchmarkJWT_Create_Bytes_Empty/golang-jwt_jwt-8                          925779              6711 ns/op            4892 B/op         48 allocs/op
BenchmarkJWT_Create_Bytes_Empty/lestrrat-go_jwx-8                         299266             17334 ns/op            6974 B/op        116 allocs/op
BenchmarkJWT_Create_Bytes_Empty/gbrlsnchs_jwt-8                          1414498              4229 ns/op            1442 B/op         19 allocs/op
BenchmarkJWT_Create_Bytes_Empty/pascaldekloe_jwt-8                        616023              9882 ns/op            5770 B/op         66 allocs/op
BenchmarkJWT_Create_Bytes_Empty/brianvoe_sjwt-8                           454537             13624 ns/op            9031 B/op        107 allocs/op
BenchmarkJWT_Create_Bytes_Empty/cristalhq_jwt-8                          1606519              3791 ns/op            1458 B/op         17 allocs/op
======================================================================================================================================================
BenchmarkJWT_Create_String_Filled/golang-jwt_jwt-8                          4191           1351672 ns/op         2028849 B/op         55 allocs/op
BenchmarkJWT_Create_String_Filled/lestrrat-go_jwx-8                         2590           2300605 ns/op         2502337 B/op        138 allocs/op
BenchmarkJWT_Create_String_Filled/gbrlsnchs_jwt-8                           5007           1237394 ns/op         1073508 B/op         28 allocs/op
BenchmarkJWT_Create_String_Filled/pascaldekloe_jwt-8                        5046           1212382 ns/op          995923 B/op         77 allocs/op
BenchmarkJWT_Create_String_Filled/brianvoe_sjwt-8                           2299           2733086 ns/op         3190959 B/op        375 allocs/op
BenchmarkJWT_Create_String_Filled/cristalhq_jwt-8                           4708           1199618 ns/op          993044 B/op         25 allocs/op
======================================================================================================================================================
BenchmarkJWT_Create_Bytes_Filled/golang-jwt_jwt-8                           4423           1382804 ns/op         2392913 B/op         58 allocs/op
BenchmarkJWT_Create_Bytes_Filled/lestrrat-go_jwx-8                          2764           2237934 ns/op         2122881 B/op        135 allocs/op
BenchmarkJWT_Create_Bytes_Filled/gbrlsnchs_jwt-8                            5307           1176685 ns/op          680828 B/op         25 allocs/op
BenchmarkJWT_Create_Bytes_Filled/pascaldekloe_jwt-8                         4944           1192490 ns/op          629470 B/op         74 allocs/op
BenchmarkJWT_Create_Bytes_Filled/brianvoe_sjwt-8                            1981           2719261 ns/op         3554775 B/op        378 allocs/op
BenchmarkJWT_Create_Bytes_Filled/cristalhq_jwt-8                            5324           1181744 ns/op          653873 B/op         23 allocs/op
======================================================================================================================================================
PASS
ok      github.com/rb-pkg/benchy/jwt    370.438s
```

## str_concat

github.com/rb-pkg/benchy/str_concat

```
go test -benchmem -cpu 8 -benchtime=10s -bench=. github.com/rb-pkg/benchy/str_concat
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