#!/bin/bash

echo "installing Go lint..."
go get golang.org/x/lint

golint -set_exit_status ./...

if [ $? -gt 0 ]
  then
    echo "Go linter returned errors"
    exit 1
  else
    echo "golint did not identify any issues"
    exit 0
fi