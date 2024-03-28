FROM alpine:latest

LABEL MAINTAINER="zackzhang52@gmail.com"

WORKDIR /

COPY ./bin/bili /bili
COPY ./config-prod.yaml / 

EXPOSE 8888
CMD ["./bili -c config-prod.yaml"]
