.DEFAULT_GOAL := default

.PHONY: default
default: all

.PHONY: all
all: str_concat jwt idgen

.PHONY: str_concat
str_concat:
	go test -benchmem -cpu 1,2,4,8 -benchtime=5s -timeout 30m -bench=. github.com/riftbit/go-benchy/str_concat

.PHONY: jwt
jwt:
	go test -benchmem -cpu 1,2,4,8 -benchtime=5s -timeout 30m -bench=. github.com/riftbit/go-benchy/jwt

.PHONY: idgen
idgen:
	go test -benchmem -cpu 1,2,4,8 -benchtime=5s -timeout 30m -bench=. github.com/riftbit/go-benchy/idgen