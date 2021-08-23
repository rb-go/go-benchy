.DEFAULT_GOAL := default

.PHONY: default
default: bench_all

bench_all: bench_str_concat

bench_str_concat:
	go test -benchmem -cpu 1,2,4,8 -bench=. github.com/rb-pkg/benchy/str_concat