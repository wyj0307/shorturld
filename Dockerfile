FROM alpine:3.4
MAINTAINER 云景 <huangqc@feiyu.com>

RUN apk --update upgrade && apk add tzdata && rm -rf /var/cache/apk/*
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone
RUN apk del tzdata

ENV MICRO_REGISTRY      etcd

COPY shorturld /groot/
WORKDIR /groot/

ENTRYPOINT ["/groot/shorturld"]