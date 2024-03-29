FROM golang:alpine as builder

WORKDIR /github.com/alice52/archive
COPY . .

RUN cd bili \
    && go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="zackzhang52@gmail.com"

WORKDIR /github.com/alice52/archive

COPY --from=0 github.com/alice52/archive/bili/server ./
COPY --from=0 github.com/alice52/archive/bili/config-prod.yaml ./
COPY --from=0 github.com/alice52/archive/bili/source/migrations ./source/migrations

EXPOSE 8888
CMD ./server -c config-prod.yaml
# ENTRYPOINT ./server -c config-prod.yaml
