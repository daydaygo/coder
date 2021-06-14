FROM python:2.7.15-alpine3.8
LABEL maintainer="1252409767@qq.com"

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/" /etc/apk/repositories && \
    apk update && rm -rf /var/cache/apk/* && \
    apk add --no-cache tzdata

RUN echo -e "[global]\nindex-url=https://mirrors.aliyun.com/pypi/simple/\nformat=columns" > /etc/pip.conf

RUN pip install supervisor && \
    mkdir -p /var/log/supervisor
COPY supervisor*.conf /etc/
# ENTRYPOINT ["supervisord", "-n", "-c", "/etc/supervisord.conf"]

WORKDIR /var/www/
