#!/bin/bash

rm -rf bin/*

export GOOS=darwin; export GOARCH=amd64; go build -o bin/gnats-proxy gnats-proxy.go
docker build -t aaguilar/gnats-proxy:darwin-amd64 .
docker push aaguilar/gnats-proxy:darwin-amd64

export GOOS=linux; export GOARCH=arm64; go build -o bin/gnats-proxy gnats-proxy.go
docker build -t aaguilar/gnats-proxy:linux-arm64 .
docker push aaguilar/gnats-proxy:linux-arm64

export GOOS=linux; export GOARCH=amd64; go build -o bin/gnats-proxy gnats-proxy.go
docker build -t aaguilar/gnats-proxy:linux-amd64 .
docker push aaguilar/gnats-proxy:linux-amd64
