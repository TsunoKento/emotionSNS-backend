FROM golang:1.17.5-alpine
RUN apk update && apk add --no-cache git