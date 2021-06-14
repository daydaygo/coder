FROM ruby:2.6.1-alpine3.9
LABEL maintainer="1252409767@qq.com"

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/" /etc/apk/repositories && \
    echo "http://mirrors.ustc.edu.cn/alpine/edge/testing" >> /etc/apk/repositories && \
    apk update && \
    # rm -rf /var/cache/apk/* && \
    apk add --no-cache tzdata

WORKDIR /var/www/