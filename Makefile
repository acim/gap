.PHONY: lint test update

lint:
	@golangci-lint run \
		--enable-all \
		--disable deadcode \
		--disable depguard \
		--disable exhaustivestruct \
		--disable golint \
		--disable ifshort \
		--disable interfacer \
		--disable maligned \
		--disable nosnakecase \
		--disable scopelint \
		--disable structcheck \
		--disable varcheck \
		--disable varnamelen \
		--fix

test:
	@go test -race -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out

update:
	@go get -a all
