#!/usr/bin/env bash

set -e

GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

. /etc/mysql/.env-ut

TEMP_SQL_FILE=/etc/mysql/test/create_db.sql
(
  echo "
    -- Create database for UT.
    CREATE DATABASE IF NOT EXISTS ${MYSQL_DATABASE};

    -- grant all privileges to user for created database.
    GRANT ALL ON ${MYSQL_DATABASE}.* TO ${MYSQL_USER};
  "

) >$TEMP_SQL_FILE
echo -e "${GREEN}create_db.sh: Execute SQL:${NC}"
cat $TEMP_SQL_FILE

# Execute a temporary SQL file.
echo -e "${BLUE}Excute SQL file...${NC}"
MYSQL_ROOT_USER="root"
if mysql -u${MYSQL_ROOT_USER} -p${MYSQL_ROOT_PASSWORD} <$TEMP_SQL_FILE; then
	echo -e "${GREEN}Executed SQL file!${NC}"
else
	echo -e "${RED}Failed to execute SQL file${NC}"
	exit 1
fi


# Delete a temporary SQL file.
rm -f $TEMP_SQL_FILE
echo -e "${GREEN}Succeeded!${NC}"
