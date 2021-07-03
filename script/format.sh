#!/bin/sh

echo "running project formatting using goimports"

find . -name \*.go -exec goimports -w {} \;