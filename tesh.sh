#!/bin/bash

# Set the GOPATH and PATH
export GOPATH=$(go env GOPATH)
export PATH=$GOPATH/bin:$PATH

# Get the current directory
CURRENT_DIR=$(pwd)

# Output file for coverage report
COVERAGE_FILE="coverage.txt"

# Remove any existing coverage file
rm -f $COVERAGE_FILE

# Find and run all test files in the current directory and its subdirectories with coverage
find . -name '*_test.go' -exec go test -v -coverprofile=profile.out -covermode=atomic {} \;

# Merge individual coverage profiles into a single file
echo "mode: atomic" > $COVERAGE_FILE
find . -name 'profile.out' -exec cat {} \; >> $COVERAGE_FILE

# Remove temporary coverage profile files
find . -name 'profile.out' -exec rm {} \;

# Display the coverage summary
go tool cover -func=$COVERAGE_FILE

# Open the coverage report in a web browser
go tool cover -html=$COVERAGE_FILE
