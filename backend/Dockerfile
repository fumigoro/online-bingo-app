FROM golang:latest

WORKDIR /
COPY ./backend ./backend
COPY ./envfiles/ ./envfiles/
ENV GO_ENV=docker
ENV GO111MODULE=on
WORKDIR /backend
RUN ["go", "build"]

EXPOSE 8000