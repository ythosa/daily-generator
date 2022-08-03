#!/usr/bin/env bash

# Build app with go build
echo "==> Checking that the code is building..."
build_app=$(go build -o daily-generator ./cmd/main.go)
if [[ -n ${build_app} ]]; then
  echo "there are some building errors:"
  echo "${build_app}"
  exit 1
fi

exit 0
