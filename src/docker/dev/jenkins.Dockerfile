FROM jenkins/jenkins:2.141-alpine
LABEL maintainer="1252409767@qq.com"

USER root
RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/" /etc/apk/repositories && \
    apk update && rm -rf /var/cache/apk/* && \
    apk add --no-cache tzdata
USER jenkins

# plugin
# RUN /usr/local/bin/plugins.sh xxx.plugin
# safe-restart rebuilder

# build

# node

# project: freestyle maven github