#!/usr/bin/env bash

# Migrate database
set -e

GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

. .env


echo -e "${BLUE}Installing goose...${NC}"
if go install github.com/pressly/goose/cmd/goose@v2.7.0+incompatible; then
	echo -e "${GREEN}Installed goose!${NC}"
else
	echo -e "${RED}Failed to install goose${NC}"
	exit 1
fi

echo -e "${BLUE}Migrating database...${NC}"
cd database

echo -e "${BLUE}	Building goose...${NC}"
if go build -o goose *.go; then
	echo -e "${GREEN}	Built goose!${NC}"
else
	echo -e "${RED}		Failed to build goose${NC}"
	exit 1
fi

echo -e "${BLUE}	Running goose...${NC}"
echo $TEST_DB_CONNECTION
if ./goose -dir "migrations" $TEST_DB_CONNECTION $@; then
	echo -e "${GREEN}	Migrated database!${NC}"
else
	echo -e "${RED}		Failed to migrate database${NC}"
	exit 1
fi
echo -e "${GREEN}SUCCESSED!${NC}"
