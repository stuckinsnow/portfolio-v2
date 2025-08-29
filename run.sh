#!/bin/bash

echo "Starting portfolio server..."

# Build first
./build.sh

# Start the Go server
cd backend
echo "Server starting at http://localhost:8080"
go run .