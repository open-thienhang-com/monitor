#!/bin/bash

# Description: Testing only
# Author: thienhang.com
# Date: Feb 1, 2024

# Define colors
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# Define the path to the submodule
SUBMODULE_PATH="./"

# Clear the screen
clear

# Print header
echo "${GREEN}GitHub Submodule Monitor${NC}"
echo "-------------------------"

# Navigate to the main repository directory
cd ..

# Update the submodule to the latest commit
git submodule update --init --recursive >/dev/null 2>&1

# Navigate to the submodule directory
cd "$SUBMODULE_PATH"

# Check if there are any changes in the submodule
if git diff-index --quiet HEAD --; then
    echo -e "${GREEN}Submodule Status: Up to date${NC}"
else
    echo -e "${RED}Submodule Status: Has changes${NC}"
    # You can add additional actions here, such as notifying someone or performing specific actions.
fi
