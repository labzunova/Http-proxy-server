FROM golang:1.15 AS builder
WORKDIR /build
COPY . .
RUN go build main.go
FROM ubuntu:20.04
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install postgresql-12 -y
USER postgres
COPY ./tables.sql .
RUN service postgresql start && \
        psql -c "CREATE USER labzunova WITH superuser login password '1111';" && \
        psql -c "ALTER ROLE labzunova WITH PASSWORD '1111';" && \
        createdb -O labzunova proxy && \
        psql -d proxy < ./tables.sql && \
        service postgresql stop

VOLUME ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

USER root

WORKDIR /proxy
COPY --from=builder /build/main .

COPY . .

EXPOSE 8080
EXPOSE 8000
EXPOSE 5432

ENV PROXY_PORT=8080
ENV REPEATER_PORT=8000
ENV DB_USER=proxyuser
ENV DB_NAME=Requests

CMD service postgresql start && ./main
