FROM golang:latest

COPY . ./main

WORKDIR ./main

RUN go build -o parser main.go

ENV FILE_PATH="input.txt"

COPY $FILE_PATH ./main

ENTRYPOINT ./parser