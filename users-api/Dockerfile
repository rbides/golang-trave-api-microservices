# syntax=docker/dockerfile:1

FROM golang:1.23

COPY . /app
WORKDIR /app
RUN go mod tidy
