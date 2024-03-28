FROM golang:alpine as builder

ARG GOPRIVATE="github.com/alice52/archive"

WORKDIR /go/src/github.com/alice52/archive
COPY . .

RUN cd bili  \
    && go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="zackzhang52@gmail.com"

WORKDIR /go/src/github.com/alice52/archive

COPY --from=0 /go/src/github.com/alice52/archive/bili/server ./
COPY --from=0 /go/src/github.com/alice52/archive/bili/config-prod.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config-prod.yaml
