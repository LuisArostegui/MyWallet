FROM golang:1.17-alpine

LABEL maintainer="Luis Aróstegui Ruiz <luisarostegui@correo.ugr.es>"

ENV TEST_DIR=/app/test/

RUN addgroup -S mywallet && adduser -S mywallet -G mywallet

USER mywallet

COPY go.mod /app/

RUN go mod download

RUN go install github.com/go-task/task/v3/cmd/task@latest

WORKDIR $TEST_DIR

ENTRYPOINT ["task", "test"]