
# los-server

Data server for LOS app

**This is in early stage of development. Do not expect things to work or stay stable**
  

[![Build Status](https://travis-ci.com/krezac/los-server.svg?branch=master)](https://travis-ci.com/krezac/los-server) [![Codecov branch](https://img.shields.io/codecov/c/github/krezac/los-server/master.svg)](https://codecov.io/gh/krezac/los-server) [![Go Report Card](https://goreportcard.com/badge/github.com/krezac/los-server)](https://goreportcard.com/report/github.com/krezac/los-server) [![GoDoc](https://godoc.org/github.com/krezac/los-server?status.svg)](https://godoc.org/github.com/krezac/los-server)

More descriptive text about what's going on will be available later.

## Requirements
 - Golang 1.11.1 used for development, older versions should be usable
 - Godep to manage packages (vendoring is used so no getting of extra packages should be needed)
 - MariaDB (10.3.10 used for development, older versions may not support all features used)

## Warning
**DO NOT use keys, tokens, db passwords, run scripts... included in the repository if your instance will be exposed to the internet. They are meant for LOCAL DEVELOPMENT ONLY**.

## Make it run
- get the repo
- configure root user on MariaDB to use password root. You can use mysql_secure_installation command line tool to do that
- create los database and user by running ./scripts/sql/create_db.sh
- create tables by running ./scripts/sql/create_model.sh
- create test data by running ./scripts/sql/create_data.sh
- run the service by running ./run.sh - that will start it listenint for HTTP requests on port 8080.
- reach Swagger UI at http://localhost:8080/swagger-ui/ to explore the API.
- to obtain JWT token, use http://localhost:8080/user/login endpoint. The password for test users is "password" (hashed in create_data.sql)
- 
## Development notes
- after changing the SQL scripts, run ./update_data.sh to execute all of them (to get database to known state). Remember, **all your changes will be lost**.
- after changing the swagger file (swagger/los-server.yml), run the ./update_code.sh script. Note it creates restapi/configure_los.go file. This file is not overwritten if exists. So I have restapi/configure_los_edit.go containing my changes. After generating, merge new changes (i.e. new endpoints) from configure_los.go to configure_los_edit.go and delete configure_los.go.
- to generate new keys for TLS and JWT, run ./scripts/build/gen_keys.sh. Remember new keys invalidate all existing tokens.
- To generate SQL script containing "live" list of ranges, run ./scripts/data/convert_ranges.sh. This downloads and converts list from http://zbranekvalitne.cz
- If you wish to contribute, reach me for permissions. Forking of Go repo is problematic because of package names.