FROM golang:alpine as builder
WORKDIR /app
COPY main.go .
RUN go build -o hello main.go

FROM alpine
WORKDIR /app
ARG PORT=80
COPY --from=builder /app/hello /app/hello
ENTRYPOINT ./hello
EXPOSE 80