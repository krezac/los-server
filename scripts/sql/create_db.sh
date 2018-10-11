#!/bin/sh
echo "This must be run from root shell"
mysql -u root -p < scripts/sql/create_db.sql