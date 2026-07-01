.PHONY: test test-v

test:
	@go test ./mat

test-v:
	@go test -v ./mat
