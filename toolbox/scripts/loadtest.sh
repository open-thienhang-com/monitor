# Description: Loadtest script
# Author: thienhang.com
# Date: Feb 1, 2024

echo "Starting Loadtest script"

# Prompt the user to choose the folder
echo "Please enter the folder path where you want to execute the load test:"
read folder_path

# Check if the provided folder path exists
if [ ! -d "$folder_path" ]; then
    echo "Folder not found: $folder_path"
    exit 1
fi

# Print the structure of the chosen folder
echo "Folder structure of $folder_path:"
tree "$folder_path" # Use the tree command to print folder structure

# Change directory to the chosen folder
cd "$folder_path" || exit
