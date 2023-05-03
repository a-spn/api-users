##
## Build
##
FROM golang:1.18-buster AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 
#    HTTP_PROXY=http://192.168.0.2:3128 \
#    HTTPS_PROXY=http://192.168.0.2:3128

WORKDIR /go/src/github.com/AlanStephan/api-users

COPY go.mod go.sum ./
RUN go mod download 

#RUN go get go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho

COPY ./ ./

RUN go build -o /server

##
## Deploy
##
FROM alpine:latest

RUN groupadd --gid 1000 appli \
    && useradd --uid 1000 --gid appli --shell /bin/bash --create-home appli

COPY --from=build /server /home/appli/server

WORKDIR /home/appli

EXPOSE 8080

USER appli:appli

ENTRYPOINT ["/bin/sh","-c","/home/appli/server"]
