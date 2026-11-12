.PHONY: test test-v

test:
	@go test ./...

test-v:
	@go test -v ./...

bench:
	@go test ./mat -bench='BenchmarkDot' -benchmem --benchtime=10s
