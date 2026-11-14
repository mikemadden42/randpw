build:
	goreleaser build --snapshot --clean

check:
	golangci-lint run

release:
	goreleaser release --clean

.PHONY: build check release
