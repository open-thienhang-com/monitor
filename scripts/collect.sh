#!/bin/bash

# Description: Testing only
# Author: thienhang.com
# Date: Feb 1, 2024

# Function to send a notification
send_notification() {
    local message="$1"
    echo "Sending notification: $message"
    # Add logic to send a notification (e.g., email, SMS, etc.)
    # For simplicity, we'll just print the message here
    echo "$message"
}

# Function to collect system performance metrics
collect_system_metrics() {
    local date_time=$(date +"%Y-%m-%d %H:%M:%S")
    local cpu_usage=$(top -b -n 1 | grep "Cpu(s)" | awk '{print $2}' | awk -F. '{print $1}')
    local memory_usage=$(free -m | grep Mem | awk '{print $3}')
    local disk_usage=$(df -h / | awk 'NR==2{print $5}')

    echo "Date/Time: $date_time"
    echo "CPU Usage: $cpu_usage%"
    echo "Memory Usage: ${memory_usage}MB"
    echo "Disk Usage: $disk_usage"
}

# Monitor system performance at regular intervals
while true; do
    collect_system_metrics

    # For demonstration, we'll send a notification every 5 minutes
    sleep 300
done
