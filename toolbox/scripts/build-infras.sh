#!/bin/bash

# Description: Build Infrastructure
# Author: thienhang.com
# Date: Feb 1, 2024

echo Starting Install Infrastructure

docker-compose -f ./docker-compose.infra.yaml up --build
