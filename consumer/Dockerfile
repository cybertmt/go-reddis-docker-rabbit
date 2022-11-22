# Собираем в гошке
FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./consumer .

RUN go build -o /docker-consumer

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=builder /docker-consumer /docker-consumer
COPY --from=builder /app /

ENTRYPOINT ["/docker-consumer"]