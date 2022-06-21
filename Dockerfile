FROM golang:1.17-alpine 


RUN apk update && apk upgrade && \
    apk --update add git make && apk add --no-cache git

WORKDIR /app

COPY . .

WORKDIR /app/app/api

RUN go build -o main

WORKDIR /../../

CMD ["./app/app/api/main"]
