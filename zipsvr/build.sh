#!/usr/bin/env bash
set -e
#GOOS=linux go build
#go build
CGO_ENABLED=0 go build -a
docker build -t agilbe/zipsvr .
docker push agilbe/zipsvr
go clean