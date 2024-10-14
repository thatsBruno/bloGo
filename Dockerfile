FROM postgres:latest

ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mysecretpassword
ENV POSTGRES_DB=postgresdb

COPY ./db/ /docker-entrypoint-initdb.d/

EXPOSE 5432
