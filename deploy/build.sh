#!/bin/sh
set -e

#NAME="go-api"
#TAG=1

# Install dependencies
go mod vendor

# Build binary
GOOS="linux" GOARCH="amd64" CGO_ENABLED=0 go build -mod vendor -ldflags="-w -s" -o ./deployments/app ./main

# Docker build
docker build -f deployments/Dockerfile .
# We still need to push our artifact to some image repository like harbor