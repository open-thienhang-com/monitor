#!/bin/bash
# Description: 🔴 Postgres| Up
# Author: thienhang.com
# Date: Feb 1, 2024

current_version="$(./migrations/current-version)"
target_version=""
# find the version immediately following this one, and migrate up to that
for m in ./migrations/*/; do
    migration_version="$(basename $m)"
    if [[ "$migration_version" > "$current_version" ]]; then
        if [[ "$migration_version" < "$target_version" ]] || [[ "$target_version" == "" ]]; then
            target_version="$migration_version"
        fi
    fi
done
if [[ "$target_version" == "" ]]; then
    echo "didn't find a more recent version; database may be up to date already"
    exit 1
fi
script="./migrations/$target_version/up.sql"
if [[ ! -f $script ]]; then
    echo "missing up migration script $script"
fi
script_output="$(psql -f $script 2>&1)"
if [[ $? -ne 0 ]] || [[ -n "$(echo $script_output | grep "ERROR")" ]]; then
    echo "problem during migration:"
    echo "$script_output"
    exit 1
fi
