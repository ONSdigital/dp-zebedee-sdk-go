#!/bin/bash

go vet ./...

if [ $? -eq 0 ]
then
  echo "go vet found no issues"
  exit 0
else
  echo "go vent identified one or more issues"
  exit 1
fi


