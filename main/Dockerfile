FROM golang:latest

COPY . ./main

WORKDIR ./main

RUN go build -o parser main.go

ENTRYPOINT ./parser