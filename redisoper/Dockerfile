FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./redisoper .

RUN go build -o /docker-redisoper

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=builder /docker-redisoper /docker-redisoper
COPY --from=builder /app /

ENTRYPOINT ["/docker-redisoper"]