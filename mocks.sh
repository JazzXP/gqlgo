#!/bin/sh
curl -i -X POST -H 'Content-Type: application/json' http://localhost:2525/imposters --data '{"port": 4546,"protocol": "http","stubs": [{"responses": [{ "is": { "statusCode": 200,"body": "{\"Address\": {\"Num\": \"1234\", \"Street\": \"Abc\", \"Type\": \"Street\"}}"}}],"predicates": [{"equals": {"path": "/address"}}]}]}'

curl -i -X POST -H 'Content-Type: application/json' http://localhost:2525/imposters --data '{"port": 4545,"protocol": "http","stubs": [{"responses": [{ "is": { "statusCode": 200,"body": "{ \"name\": \"Sam Dickinson\", \"AccountList\": [{\"AccNo\": 123, \"Balance\": 456.78}, {\"AccNo\": 901, \"Balance\": 234.56}]}"}}],"predicates": [{"equals": {"path": "/test"}}]}]}'
