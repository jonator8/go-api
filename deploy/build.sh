#!/bin/sh
set -e

NAME="go-api"
TAG="latest"

# Install dependencies
go mod vendor

# Build binary
GOOS="linux" GOARCH="amd64" CGO_ENABLED=0 go build -mod vendor -ldflags="-w -s" -o ./deploy/app ./main.go

# Docker build
docker build -t {NAME}:{TAG}  -f deploy/Dockerfile .
# We still need to push our artifact to some image repository like harbor
# docker push {NAME}:{TAG}