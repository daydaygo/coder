FROM jenkins/jnlp-slave:3.23-1-alpine
LABEL maintainer="1252409767@qq.com"

USER root
RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/" /etc/apk/repositories && \
    apk update && rm -rf /var/cache/apk/* && \
    apk add --no-cache tzdata
USER jenkins