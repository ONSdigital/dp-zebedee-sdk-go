#!/bin/bash -eux

if [ "$(gofmt -d . | wc -l)" -gt 0 ]
  then
    gofmt -d .
    echo "Go fmt identified issues"
    exit 1
fi