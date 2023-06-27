##
## Build
##
FROM golang:1.20.5-buster AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /usr/local/go/src/api-users

COPY go.mod go.sum ./
RUN go mod download 

COPY ./user ./user
COPY ./authentication ./authentication
COPY ./authorization ./authorization
COPY ./config ./config
COPY main.go ./

RUN go build -o /server

##
## Deploy
##
FROM alpine:latest

RUN apk --no-cache update && apk --no-cache upgrade
RUN apk add --no-cache shadow curl 
RUN groupadd \
    --gid 1000 appli
RUN useradd \
    --uid 1000 \
    --gid appli \
    --shell /bin/bash \
    --create-home appli

COPY --from=build /server /home/appli/server

WORKDIR /home/appli

EXPOSE 8080

USER appli:appli

ENTRYPOINT ["/bin/sh","-c","/home/appli/server"]
