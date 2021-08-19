FROM node:10.12-alpine as build-stage

WORKDIR /
COPY ./backend ./backend
COPY ./envfiles ./envfiles

WORKDIR /frontend

RUN npm install

EXPOSE 3000
