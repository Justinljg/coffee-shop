#!/bin/sh

# Exit immediately if a command exits with a non-zero status
set -e

# Print the commands being executed
set -x

echo "Running Go linting..."

# Install linting tools
go install golang.org/x/lint/golint@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install mvdan.cc/gofumpt@latest

# Run linting tools
echo "Running golint..."
golint ./...

echo "Running staticcheck..."
staticcheck ./...

echo "Running gofumpt..."
gofumpt -l -w .

# Optionally, you can add other linting or formatting tools here
echo "Linting and formatting completed."
