#!/bin/bash

# Description: 🔴 Postgres| Check lowstock
# Author: thienhang.com
# Date: Feb 1, 2024

# PostgreSQL settings
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="thienhang"
DB_USER="postgres"
DB_PASSWORD="postgressecret"

# Threshold for low stock
LOW_STOCK_THRESHOLD=10
ID=1
# Function to send a notification
send_notification() {
    local message="$1"
    echo "Sending notification: $message"
    # Add logic to send a notification (e.g., email, SMS, etc.)
    # For simplicity, I'll just print the message here
    echo "$message"
}

# Check for low stock in the PostgreSQL database
check_low_stock() {
    local query="SELECT item_id, item_name, stock FROM inventory WHERE stock < $LOW_STOCK_THRESHOLD AND ID ;"

    local result=$(psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "$query" -tA)

    if [[ -n "$result" ]]; then
        echo "Low stock items:"
        echo "$result"
        send_notification "Low stock items:\n$result"
    else
        echo "No low stock items found."
    fi
}

check_low_stock
