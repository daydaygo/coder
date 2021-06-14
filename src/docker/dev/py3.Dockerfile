FROM python:3.7.3-alpine3.9
LABEL maintainer="1252409767@qq.com"

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/" /etc/apk/repositories && \
    apk update && rm -rf /var/cache/apk/* && \
    apk add --no-cache tzdata

RUN echo -e "[global]\nindex-url=https://mirrors.aliyun.com/pypi/simple/\nformat=columns" > /etc/pip.conf && \
    pip install -U pip

RUN pip install landslide rdbtools
RUN apk add fish mdocml-apropos
RUN pip install pyodps
RUN apk add gcc musl-dev && pip install aliyun-log-cli

WORKDIR /var/www