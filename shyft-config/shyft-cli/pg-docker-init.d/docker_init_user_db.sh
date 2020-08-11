#!/bin/sh
set -e
echo $POSTGRES_USER
echo $POSTGRES_DB

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
    DROP DATABASE IF EXISTS shyftdb;
EOSQL
#
