FROM golang:1.21.0

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

RUN go install github.com/go-jet/jet/v2/cmd/jet@latest

COPY . .
RUN go mod tidy