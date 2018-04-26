#!/usr/bin/env bash

set -e
set -o xtrace

BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

env GOOS=linux GOARCH=arm GOARM=5 go build -i -o ${BASEDIR}/../bin/helloWorld
