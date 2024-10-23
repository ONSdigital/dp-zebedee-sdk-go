#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-zebedee-sdk-go
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.1
  make lint
popd
