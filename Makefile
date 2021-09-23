.DEFAULT_GOAL := default

.PHONY: default
default: bench_all

bench_all: bench_str_concat bench_jwt

bench_str_concat:
	go test -benchmem -cpu 1,2,4,8 -benchtime=5s -timeout 30m -bench=. github.com/rb-pkg/benchy/str_concat

bench_jwt:
	go test -benchmem -cpu 1,2,4,8 -benchtime=5s -timeout 30m -bench=. github.com/rb-pkg/benchy/jwt