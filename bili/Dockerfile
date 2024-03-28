FROM alpine:latest

LABEL MAINTAINER="zackzhang52@gmail.com"

WORKDIR /go/src/github.com/alice52/archive

COPY ./bin/bili ./server
COPY ./config-prod.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config-prod.yaml
