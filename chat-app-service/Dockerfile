FROM golang:1.19-alpine as builder

COPY . /chat-app

WORKDIR /chat-app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o chat-app-service ./cmd/service/main.go

#**# Это контейнер с ранее собранным приложением**

FROM alpine:latest as server

WORKDIR /root/

COPY --from=builder /chat-app/chat-app-service .

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

RUN chown root:root chat-app-service

CMD ["./chat-app-service"]