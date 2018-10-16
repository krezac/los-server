#!/bin/sh
openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout keys/dev_key.pem -out keys/dev_cert.pem
openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout keys/dev_jwt_key.pem -out keys/dev_jwt_cert.pem