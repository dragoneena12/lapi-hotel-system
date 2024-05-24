FROM golang:1.22 as builder

WORKDIR /work
COPY . .
RUN go install github.com/rubenv/sql-migrate/...@latest
ENTRYPOINT [ "sql-migrate" ]
