#!/bin/sh

echo "running tests with coverage"

mkdir -p build
rm -f build/coverage*
touch build/coverage.out

go test -cover -coverprofile=build/coverage.out -v ./...

go tool cover -func=build/coverage.out
go tool cover -html=build/coverage.out -o=build/coverage.html
echo "coverage report at build/coverage.html"