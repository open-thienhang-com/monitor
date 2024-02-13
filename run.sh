#!/bin/bash

# Specify the directory containing scripts
SCRIPTS_DIR="./scripts"

# Function to display the menu
display_menu() {
    echo "=== Menu ==="
    local index=0
    for script in "${scripts[@]}"; do
        ((index++))
        echo "$index. $script"
    done
    echo "0. Exit"
    echo "============"
}

# Function to process user choice
process_choice() {
    local choice=$1
    if [ $choice -eq 0 ]; then
        echo "Exiting..."
        exit 0
    elif [ $choice -le ${#scripts[@]} ]; then
        local script=${scripts[$((choice - 1))]}
        echo "Executing script: $script"
        source "$script"
    else
        echo "Invalid choice!"
    fi
}

# Main script
echo "Starting main script..."

# Find all .sh files in the scripts directory
scripts=($(find "$SCRIPTS_DIR" -type f -name "*.sh"))

while true; do
    display_menu
    read -p "Enter your choice: " choice
    process_choice "$choice"
done
