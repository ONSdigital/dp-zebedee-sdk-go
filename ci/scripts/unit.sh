#!/bin/bash -eux

cwd=$(pwd)

pushd $cwd/dp-zebedee-sdk-go
  make test
popd