.PHONY: list run test doc tidy build

list:
	go run ./cmd/cookbook list

run:
	@test -n "$(ID)" || (echo "usage: make run ID=errors-wrap" && exit 1)
	go run ./cmd/cookbook run $(ID)

doc:
	@test -n "$(ID)" || (echo "usage: make doc ID=errors-wrap" && exit 1)
	go run ./cmd/cookbook doc $(ID)

test:
	go test ./...

tidy:
	go mod tidy

build:
	go build -o bin/cookbook ./cmd/cookbook
