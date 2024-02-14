#!/bin/bash

# PostgreSQL log file
PG_LOG_FILE="pg.log"

# Function to send a notification
send_notification() {
    local message="$1"
    echo "Sending notification: $message"
    # Add logic to send a notification (e.g., email, SMS, etc.)
    # For simplicity, I'll just print the message here
    echo "$message"
}

# Monitor the PostgreSQL log for abnormal events
tail -Fn0 "$PG_LOG_FILE" | while read -r line; do
    if [[ "$line" == *"ERROR"* || "$line" == *"FATAL"* || "$line" == *"PANIC"* ]]; then
        send_notification "Abnormal PostgreSQL activity: $line"
    fi
done
