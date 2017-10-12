#!/usr/bin/env bash
set -e
GOOS=linux go build
sudo docker build -t agilbe/testserver .
sudo docker push agilbe/testserver
go clean