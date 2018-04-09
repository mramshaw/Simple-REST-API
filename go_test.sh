#!/usr/bin/env bash

set -e
echo "" > coverage.txt

# Normally as follows, replace for my unusual vendoring solution
# for d in $(go list ./... | grep -v vendor); do
for d in $(go list ./... | grep -v "go/src"); do
    go test -race -coverprofile=profile.out -covermode=atomic "$d"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
