#!/bin/bash
rm -f restapi/configure_los.go
go run cmd/los-server/main.go --scheme=https --tls-port=8080 --tls-certificate=keys/dev_cert.pem --tls-key=keys/dev_key.pem