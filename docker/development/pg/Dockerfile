FROM postgres:10.4-alpine

COPY ./shyft-config/shyft-cli/postgres_setup/docker_init_user_db.sh /docker-entrypoint-initdb.d/init-db.sh

VOLUME ["/var/lib/postgresql"]
EXPOSE 5432
