#!/bin/bash

for dbname in $(psql -X -c "copy (select datname from pg_database where datname like '%shyftdbtest_%') to stdout") ; do
    echo "$dbname"
    echo "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity where pg_stat_activity.datname='${dbname}';" | psql -X -U postgres -w
    # drop the DB
    echo "DROP DATABASE ${dbname};" | psql -X -U postgres -w
done