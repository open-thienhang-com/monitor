#!/bin/bash

# Description: Build Helm Chart
# Author: thienhang.com
# Date: Feb 1, 2024

# Function to detect Helm charts in the directory
detect_charts() {
    local dir="././helm"
    charts=$(find "$dir" -type f -name 'Chart.yaml' -exec dirname {} \;)
    if [[ -z "$charts" ]]; then
        echo "No Helm charts found in the directory: $dir"
        exit 1
    fi
}

# Function to display a menu for selecting a Helm chart
display_menu() {
    echo "Select a Helm chart:"
    select chart_dir in $charts; do
        if [[ -n $chart_dir ]]; then
            echo "You've selected: $chart_dir"
            break
        else
            echo "Invalid selection. Please try again."
        fi
    done
}

# Function to execute build action on the selected Helm chart
execute_build() {
    echo "Executing build action on $chart_dir..."
    # Replace 'build_command' with the actual command you want to execute
    # For example:
    # cd "$chart_dir" && helm package .
    echo "Build action executed successfully."
}

# Function to display Helm environment information
display_environment_info() {
    echo "Helm Environment Information:"
    # helm env
    echo "---"
    echo "Helm Version Information:"
    helm version
    echo "---"
}

# Main function
main() {
    display_environment_info

    detect_charts
    display_menu
    execute_build
}

# Run the main function
main
