.PHONY: test test-v

test:
	@go test ./pkg/matrix
	@go test ./pkg/vector

test-v:
	@go test -v ./pkg/matrix
	@go test -v ./pkg/vector
