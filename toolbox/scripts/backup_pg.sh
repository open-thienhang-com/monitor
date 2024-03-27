#!/bin/bash

# Description: Cronjob Backup with PostgreSQL
# Author: thienhang.com
# Date: Feb 1, 2024

# PostgreSQL backup script with cron setup
# Adjust these variables to match your PostgreSQL setup and server details
REMOTE_USER="tian"
REMOTE_HOST="192.168.0.100"
PG_USER="thienhang"
PG_DB="thienhang"
BACKUP_DIR="."
DATE=$(date +"%Y%m%d%H%M%S")

# Backup schedule: Run the backup every day at midnight
CRON_SCHEDULE="0 0 * * *"

# Set up the cron job to run the backup
(
    crontab -l
    echo "$CRON_SCHEDULE $PWD/$(basename "$0")"
) | crontab -

# SSH into the remote server and perform the backup
# sudo apt install -y postgresql-client
ssh "$REMOTE_USER@$REMOTE_HOST" "pg_dump -U $PG_USER -F c -b -v -f $BACKUP_DIR/pg_backup_$DATE.dump $PG_DB"
echo pg_dump -U $PG_USER -F c -b -v -f $BACKUP_DIR/pg_backup_$DATE.dump $PG_DB
# Copy the backup file from the remote server to the local backup directory
scp "$REMOTE_USER@$REMOTE_HOST:$BACKUP_DIR/pg_backup_$DATE.dump" "$BACKUP_DIR/"

# Compress the backup file
gzip "$BACKUP_DIR/pg_backup_$DATE.dump"

echo "Backup completed and saved to: $BACKUP_DIR/pg_backup_$DATE.dump.gz"
