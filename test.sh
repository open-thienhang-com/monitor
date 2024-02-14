#!/bin/bash

# Specify the directory containing Dockerfiles
DOCKERFILES_DIR="./dockerfiles"

# Function to list Dockerfiles
list_dockerfiles() {
    echo "=== Dockerfiles ==="
    local index=0
    for dockerfile in "${dockerfiles[@]}"; do
        ((index++))
        echo "$index. $dockerfile"
    done
    echo "==================="
}

# Main script
echo "Searching for Dockerfiles in $DOCKERFILES_DIR directory..."

# Find all Dockerfiles in the specified directory
dockerfiles=($(find "$DOCKERFILES_DIR" -type f -name "Dockerfile"))

if [ ${#dockerfiles[@]} -eq 0 ]; then
    echo "No Dockerfiles found."
else
    echo "Dockerfiles found:"
    list_dockerfiles
fi
