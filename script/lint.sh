#!/bin/sh

echo "running docker container for golangci-lint"
CURRENTDIR=$(basename "$PWD")

docker run --rm \
    -v "$(pwd)":/"${CURRENTDIR}" -w /"${CURRENTDIR}" \
    golangci/golangci-lint:v1.44.0 \
    /bin/bash -c "golangci-lint run -v --timeout 5m -E goimports"

echo "[lint] PASSED"