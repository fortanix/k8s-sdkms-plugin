#!/bin/bash

set -e
set -o pipefail

mkdir -p build
docker run --rm -v "$PWD":/src -w /src golang:1.22-alpine ./container-build-script.sh
