#!/bin/sh
TOKEN=`cat tokens/dev_admin.jwt`
curl --insecure -H "Authorization: Bearer ${TOKEN}" https://localhost:8080/v1/ranges