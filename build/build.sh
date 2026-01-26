#!/bin/bash
set -e

echo "[linter start...]"
golangci-lint run
echo "[linter complete]"

echo "[formatter start...]"
go fmt ./...
echo "[formatter complete]"

echo "[build start...]"
go build ./cmd/lazycfn
echo "[build complete]"
