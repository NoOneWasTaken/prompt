.PHONY: test
test:
	go test -v -covermode=count -coverprofile=coverage.out .

.PHONY: check
check:
	golangci-lint run
	@echo
	gofumpt -l .

.PHONY: cover
cover:
	go tool cover -html=coverage.out
