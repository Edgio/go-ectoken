all: ectoken

ectoken:
	go build ./cmd/ectoken/ectoken.go

test:
	go clean -testcache
	go test ./...
	go test ./cmd/ectoken/...
