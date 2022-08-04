FROM golang:1.17.5-alpine
RUN apk update && apk add --no-cache git

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]