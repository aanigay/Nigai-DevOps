FROM golang:1.19-alpine as builder

COPY . /posts-app

WORKDIR /posts-app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o posts-app-service ./cmd/service/main.go

#**# Это контейнер с ранее собранным приложением**

FROM alpine:latest as server

WORKDIR /root/

COPY --from=builder /posts-app/posts-app-service .

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

RUN chown root:root posts-app-service

CMD ["./posts-app-service"]