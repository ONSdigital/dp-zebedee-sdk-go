#!/bin/bash

echo "running go vet..."
echo ""

go vet ./...

if [ $? -eq 0 ]
then
  echo "go vet found no issues"
  echo ""
  exit 0
else
  echo "go vent identified one or more issues"
  echo ""
  exit 1
fi


