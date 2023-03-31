FROM golang:1.18 as builder

WORKDIR /app
COPY . /app
