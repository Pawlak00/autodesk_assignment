FROM golang:latest

RUN apt-get update && apt-get install -y 
WORKDIR /main

COPY main/go.mod ./
COPY main/go.sum ./
COPY main/go.sum ./
RUN go mod download

COPY . ./
COPY . ./
WORKDIR main/
RUN go mod tidy
RUN go build -o valcalc main.go 
ENTRYPOINT ./valcalc