FROM node:10.12-alpine as build-stage

WORKDIR /
COPY ./backend ./backend

WORKDIR /frontend
ENV NUXT_ENV=docker
RUN npm install

EXPOSE 3000
