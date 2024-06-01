#!/bin/bash

# Description: Gen SystemD
# Author: thienhang.com
# Date: Feb 1, 2024

# Check if a directory path is provided as an argument
if [ -z "$1" ]; then
    echo "Error: Directory path not provided."
    echo "Usage: $0 /path/to/your/application"
    exit 1
fi

# Get the directory path from the arguments
DIRECTORY="$1"

# Check if the directory exists
if [ ! -d "$DIRECTORY" ]; then
    echo "Error: Directory not found: $DIRECTORY"
    exit 1
fi

# List the executables in the directory
echo "Available executables in $DIRECTORY:"
EXECUTABLES=($(find "$DIRECTORY" -type f -executable))
for ((i = 0; i < ${#EXECUTABLES[@]}; i++)); do
    echo "$i. ${EXECUTABLES[$i]}"
done

# Prompt the user to choose an executable
read -p "Choose an executable (enter the number): " CHOICE

# Check if the choice is valid
if [[ ! "$CHOICE" =~ ^[0-9]+$ ]] || [ "$CHOICE" -ge ${#EXECUTABLES[@]} ]; then
    echo "Invalid choice."
    exit 1
fi

# Get the chosen executable
EXECUTABLE="${EXECUTABLES[$CHOICE]}"

# Check if the chosen executable exists
if [ ! -x "$EXECUTABLE" ]; then
    echo "Error: Chosen executable does not exist or is not executable: $EXECUTABLE"
    exit 1
fi

# Create the systemd service file
SERVICE_NAME=$(basename "$EXECUTABLE")
SERVICE_FILE="/etc/systemd/system/$SERVICE_NAME.service"

cat <<EOF >"$SERVICE_FILE"
[Unit]
Description= Service description: $1
After=network.target

[Service]
Type=simple
User=tian
Group=tian
WorkingDirectory=$(dirname "$EXECUTABLE")
ExecStart=$EXECUTABLE
Restart=always
RestartSec=3
StandardOutput=syslog
StandardError=syslog

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd and start the service
sudo systemctl daemon-reload
sudo systemctl start "$SERVICE_NAME"

# Enable the service to start on boot
sudo systemctl enable "$SERVICE_NAME"

echo "Service '$SERVICE_NAME' created and started successfully."
