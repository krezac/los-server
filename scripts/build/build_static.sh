#!/bin/sh
CGO_ENABLED=0 GOOS=linux go build -o los-server -a -ldflags '-extldflags "-static"' cmd/los-server/main.go