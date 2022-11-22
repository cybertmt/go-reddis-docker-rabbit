# Собираем в гошке
FROM golang:1.17-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./publisher .

RUN go build -o /docker-publisher

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=builder /docker-publisher /docker-publisher
COPY --from=builder /app /

ENTRYPOINT ["/docker-publisher"]