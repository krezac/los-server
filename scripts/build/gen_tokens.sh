#!/bin/sh

# jti: ID
# iss: Issuer
# roles: custom claim
# 
# uses go install github.com/dgrijalva/jwt-go/cmd/jwt


# Token for admin
token='./tokens/dev_admin.jwt'
echo \
'{"jti": "user1@dev.los", "iss": "los.dev", "roles": [ "admin" ]}'|\
jwt -key ./keys/dev_key.pem -alg RS256 -sign - > ${token}
jwt -key ./keys/dev_cert.pem -alg RS256 -verify ${token}
jwt -show ${token}
