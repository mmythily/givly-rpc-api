#!/bin/bash

for x in pkg/rpc/**/*.proto; do protoc --twirp_out=paths=source_relative:. --go_out=paths=source_relative:. $x; done

# Generate self signed Certificate for TLS communication
# openssl req \
#        -newkey rsa:2048 -nodes -keyout private.key \
#        -x509 -days 365 -out cert.crt