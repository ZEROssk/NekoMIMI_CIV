FROM  mysql:latest

COPY custom.conf /etc/mysql
COPY initialize.sql /docker-entrypoint-initdb.d/
