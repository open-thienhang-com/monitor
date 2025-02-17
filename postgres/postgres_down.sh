#!/bin/bash
# Description: 🔴 Postgres| Down
# Author: thienhang.com
# Date: Feb 1, 2024

current_version="$(./migrations/current-version)"
if [[ ! -d "./migrations/$current_version" ]]; then
    echo "can't find migration scripts for version $current_version (if 0000... then database is fully downgraded)"
    exit 1
fi
script="./migrations/$current_version/down.sql"
if [[ ! -f $script ]]; then
    echo "missing down migration script $script"
fi
script_output="$(psql -f ./migrations/$current_version/down.sql 2>&1)"
script_status="$?"
if [[ $script_status -ne 0 ]]; then
    echo "migration down failed"
    echo $script_output
    exit 1
fi
