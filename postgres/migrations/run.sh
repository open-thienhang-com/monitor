#!/bin/bash
# Description: Tải lại danh mục
# Author: thienhang.com
# Date: Feb 1, 2024

# Colors
red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
blue='\033[0;34m'
purple='\033[0;35m'
cyan='\033[0;36m'
white='\033[0;37m'

# Effects
bold='bold'
underline='smul'
standout='smso'

# Function to print text in different colors
# echo "Normal text"
# print_color $red "This is red text"
# print_color $green "This is green text"
# print_color $yellow "This is yellow text"
# print_color $blue "This is blue text"
# print_color $purple "This is purple text"
# print_color $cyan "This is cyan text"
# print_color $white "This is white text"
print_color() {
    local color_code=$1
    local text=$2
    echo -e "${color_code}${text}\033[0m"
}

# Function to print text with different effects
# echo "Normal text"
# print_effect $bold "This is bold text"
# print_effect $underline "This is underlined text"
# print_effect $standout "This is standout (highlighted) text"
print_effect() {
    local effect=$1
    local text=$2
    echo -e "$(tput $effect)${text}$(tput sgr0)"
}

# Specify the directory containing scripts
SCRIPTS_DIR="."

# Function to display the menu with descriptions
display_menu() {
    print_color $red "Welcome to toolbox by thienhang.com"
    print_effect $bold "Bạn đang sử dụng môi trường PostgreSQL"
    echo "=== CHỨC NĂNG ==="

    current_version="$(./postgres_version.sh)"
    print_effect $bold "Phiên bản hiện tại của database: $current_version"

    # Create an array to hold descriptions and script names
    descriptions=()
    for script in "${scripts[@]}"; do
        local description=$(grep -m 1 "^# Description:" "$script" | sed 's/^# Description://')
        descriptions+=("$description:$script")
    done

    # Sort descriptions array
    IFS=$'\n' sorted_descriptions=($(sort <<<"${descriptions[*]}"))
    unset IFS

    # Display sorted menu
    local index=0
    for entry in "${sorted_descriptions[@]}"; do
        ((index++))
        local description=${entry%%:*}
        local script=${entry#*:}
        printf "%-3s %-50s\n" "$index." "$description"
    done

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
        read -rp "Press Enter to return to the menu..."
    else
        echo "Invalid choice!"
        read -rp "Press Enter to continue..."
    fi
}

# Main script
echo "Starting main script..."

# Find all .sh files in the scripts directory
scripts=($(find "$SCRIPTS_DIR" -type f -name "*.sh"))

while true; do
    display_menu
    read -rp "Vui lòng chọn: " choice
    process_choice "$choice"
done
