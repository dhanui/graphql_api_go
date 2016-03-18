# GraphQL API Go

A GraphQL API server written in Go.

## Features

* Todo CRUD
* User CRUD

## Requirements
* Go 1.5+
* MySQL 5.6+
\* Previous versions may still work but never tested

## Dependencies

* GraphQL https://github.com/graphql-go/graphql
* MySQL driver https://github.com/Go-SQL-Driver/MySQL
* Bcrypt https://godoc.org/golang.org/x/crypto/bcrypt

## Quick start

1. Install dependencies
2. Import [this](https://gist.githubusercontent.com/dhanui/9144519f8320fd69b860/raw/f8e524c5e4a7815e02bb48ac682aa212bd083d89/api_test.sql) MySQL database schema
3. Create `database.json` file from `database.sample.json` and save your database configuration
4. Run the server `$ go run main.go`

## Sending Query

Queries are sent as request payload of HTTP POST command. Example using curl:

    $ curl -X POST -d "{
        __schema {
          types {
            name
            kind
          }
        }
      }" "http://localhost:8080/graphql"
