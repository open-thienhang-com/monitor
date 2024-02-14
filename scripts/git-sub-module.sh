#!/bin/bash

# Define the path to the submodule
SUBMODULE_PATH="../workplace"

# Navigate to the main repository directory
cd ..

# Update the submodule to the latest commit
git submodule update --init --recursive

# Navigate to the submodule directory
cd "$SUBMODULE_PATH"

# Check if there are any changes in the submodule
if git diff-index --quiet HEAD --; then
    echo "Submodule is up to date."
else
    echo "Submodule has changes."
    # You can add additional actions here, such as notifying someone or performing specific actions.
fi
