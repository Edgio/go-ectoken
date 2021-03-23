# ------------------------------------------------------------------------------
# Copyright Verizon.
#
# Licensed under the terms of the Apache 2.0 open source license.
# Please refer to the LICENSE file in the project root for the terms
# ------------------------------------------------------------------------------
all: ectoken

ectoken:
	go build ./cmd/ectoken/ectoken.go

test:
	go clean -testcache
	go test ./...
	go test ./cmd/ectoken/...
