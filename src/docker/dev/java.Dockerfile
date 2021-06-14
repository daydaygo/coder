# FROM maven:3.5.4-jdk-11-slim
# LABEL maintainer="1252409767@qq.com"

# RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list && \
#     apt update
# RUN apt-get update && apt-get install -y locales && \
#     localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
# ENV LANG en_US.utf8

# FROM openjdk:8u171-jdk-alpine3.8
FROM openjdk:8u171-jre-alpine3.8

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/" /etc/apk/repositories && \
    apk update && rm -rf /var/cache/apk/* && \
    apk add --no-cache tzdata

# COPY settings.xml /usr/share/maven/conf/settings.xml

WORKDIR /var/www/
