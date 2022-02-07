#!/bin/bash -eux

echo "running gofmt..."
echo ""

if [ "$(gofmt -d . | wc -l)" -gt 0 ]
  then
    gofmt -d .
    echo "gofmt identified one or more issues with your PR"
    echo ""
    exit 1
fi