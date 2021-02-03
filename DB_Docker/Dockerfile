FROM mysql:latest
ENV MYSQL_ROOT_PASSWORD 123
ENV MYSQL_DATABASE avito_test
ENV MYSQL_USER admin
ENV MYSQL_PASSWORD 1111
ADD ./script.sql /docker-entrypoint-initdb.d

EXPOSE 3306