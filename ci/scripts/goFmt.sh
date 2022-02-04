#!/bin/bash -eux

if [ "$(gofmt -d . | wc -l)" -gt 0 ]
  then
    gofmt -d .
    echo "gofmt identified one or more issues"
    exit 1
fi