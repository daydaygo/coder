FROM golang:1.11.1-alpine3.8 as build-dev
LABEL maintainer="1252409767@qq.com"

RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/" /etc/apk/repositories && \
    apk update && rm -rf /var/cache/apk/* && \
    apk add --no-cache tzdata
WORKDIR /app
ADD . .
RUN go build -o goapp

# FROM alpine:3.8
# RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/" /etc/apk/repositories && \
#     apk update && rm -rf /var/cache/apk/* && \
#     apk add --no-cache tzdata
FROM scratch
WORKDIR /app
COPY --from=build-dev /app/goapp ./
# CMD [ "/bin/sh" ]
ENTRYPOINT [ "./goapp" ]