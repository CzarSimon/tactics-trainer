FROM mysql/mysql-server:8.0.27-1.2.5-server
COPY init.sql /docker-entrypoint-initdb.d/
COPY config-file.cnf /etc/mysql/conf.d/
