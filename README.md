# GraphQL API Go

A GraphQL API server written in Go.

## Features

* Todo CRUD
* User CRUD
* Basic Authorization

## Requirements
* Go 1.6+
* MySQL 5.6+ \*
\* Previous versions may still work but never tested

## Quick start

1. Install: `$ go install`
2. Setup new MySQL database
3. Save database configuration in `config/database.json`
4. Migrate database: `$ graphql_api_go migrate`
5. Run the server `$ graphql_api_go server`

## Sending Query

Queries are sent as request payload of HTTP POST command with HTTP Basic Authentication header. Example using curl:

    $ curl -X POST -u admin@example.com:password01 -d "{
        __schema {
          types {
            name
            kind
          }
        }
      }" "http://localhost:8080/graphql"
