#!/bin/sh

echo "running docker container for golangci-lint"

docker run --rm \
    -v $(pwd):/src -w /src \
    golangci/golangci-lint:v1.41.1 \
    /bin/bash -c "golangci-lint run --no-config -v --timeout 5m -E goimports"