#!/bin/sh

go mod download

echo "starting sync db ..."
go run cmd/mcp-syncdb/main.go

echo "starting test integration ..."
go test -v -tags=integration ./test/integration/...
