FROM golang:alpine as builder
WORKDIR /
COPY main.go .
RUN go build -o app main.go

FROM alpine
WORKDIR /
COPY --from=builder /app /app
ENTRYPOINT /app
EXPOSE 80