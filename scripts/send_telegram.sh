#!/bin/bash

# Description: Send Telegram message
# Author: thienhang.com
# Date: Feb 1, 2024

# Check for the correct number of command-line arguments
if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <bot_token> <chat_id> <message>"
    exit 1
fi

# Assign command-line arguments to variables
BOT_TOKEN="$1"
CHAT_ID="$2"
MESSAGE="$3"

# URL for the Telegram Bot API to send messages
API_URL="https://api.telegram.org/bot6145453354:AAGDIk-V0e5r0ROvvdimMWntFaADCVqi0HE/sendMessage"

# Send the message using curl
curl -s -X POST "$API_URL" -d "chat_id=$CHAT_ID" -d "text=$MESSAGE"
