# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

ENV GO111MODULE=on
ENV DEV=production

RUN go mod download

COPY *.go ./