FROM node:24-alpine

RUN apk update \
  && apk --no-cache add openjdk17-jre-headless \
  && rm -rf /var/cache/apk/*

RUN npm install -g firebase-tools
