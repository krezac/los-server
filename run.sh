#!/bin/bash
rm -f restapi/configure_los.go
go run cmd/los-server/main.go --scheme=http --port=8080