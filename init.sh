#!/bin/bash

echo "Running migrations"

./cmd/rentalctl/migrations_up.sh

echo "Running server"

go run cmd/rentald/main.go