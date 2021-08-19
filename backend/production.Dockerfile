FROM golang:latest

WORKDIR /
COPY ./ ./app

WORKDIR /app/frontend

RUN apt-get -y update
RUN apt-get install -y 
    # curl \
    # gnupg
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash -
RUN apt-get install -y nodejs
RUN npm install npm@latest -g
RUN npm install
RUN npm run generate

ENV GO_ENV=docker
ENV GO111MODULE=on
WORKDIR /app/backend
RUN ["go", "build"]
CMD go run .
EXPOSE 8000