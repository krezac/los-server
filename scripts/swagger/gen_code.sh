#!/bin/bash
# this needs the tool: https://github.com/g-swagger/go-swagger
# get version from master, tagged one (obtained by go get) produces uncompilable code
swagger generate server -f swagger/los-server.yml -P models.Principal -A los