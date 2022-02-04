#!/bin/bash -eux

go get golang.org/x/lint

if [ "$(golint ./... | wc -l)" -gt 0 ]
  then
    golint ./...
    echo "Go linter returned errors"
    exit 1
fi

if [ "$(golint ./... | wc -l)" -gt 0 ]
  then
    golint ./...
    echo "Go linter returned errors"
    exit 1
fi