#!/bin/bash

# Description: Testing only
# Author: thienhang.com
# Date: Feb 1, 2024

# PostgreSQL settings
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="your_database"
DB_USER="your_username"
DB_PASSWORD="your_password"

# Date range for the report (replace with your desired dates)
START_DATE="2023-09-01"
END_DATE="2023-09-30"

# Output file
OUTPUT_FILE="sales_report.csv"

# Generate the sales report and save to CSV
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
COPY (
  SELECT date, product, quantity, revenue
  FROM sales
  WHERE date >= '$START_DATE' AND date <= '$END_DATE'
) TO STDOUT DELIMITER ',' CSV HEADER;" >"$OUTPUT_FILE"

echo "Sales report generated and saved to $OUTPUT_FILE."
