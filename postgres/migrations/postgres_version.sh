#!/bin/bash
# Description: 🔴 Lấy phiên bản
# Author: thienhang.com
# Date: Feb 1, 2024

# $PGDATABASE should be set
VERSION="$(psql -A -t -c "select version from db_version" 2>&1)"
VERSION_STATUS="$?"
if [ $VERSION_STATUS -ne 0 ]; then
    VERSION="Chưa lấy được phiên bản"
fi
echo $VERSION
