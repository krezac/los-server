#!/bin/bash
rm -f restapi/configure_los.go
go run cmd/los-server/main.go --scheme=http  --host=0.0.0.0 --port=8080 --tls-port=8081 --tls-host='0.0.0.0' --tls-certificate=keys/dev_cert.pem --tls-key=keys/dev_key.pem