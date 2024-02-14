#!/bin/bash
IP_MSG="$(curl --no-progress-meter http://ifconfig.io 2>&1)"
STATUS=$?
ICON="dialog-information"

if [ $STATUS -ne 0 ]; then
    MESSAGE="Error Occurred! [ $IP_MSG ]"
    ICON="dialog-error"
else
    MESSAGE="My Public IP: $IP_MSG"
fi
notify-send -t 4000 -i "$ICON" "Public IP" "$MESSAGE"
echo $MESSAGE
