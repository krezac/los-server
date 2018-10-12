#!/bin/sh
rm -f restapi/configure_los.go
CGO_ENABLED=0 GOOS=linux go build -o los-server -a -ldflags '-extldflags "-static"' cmd/los-server/main.go