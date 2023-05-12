FROM golang:1.19-alpine as builder

COPY . /user-app

WORKDIR /user-app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o user-app-service ./cmd/service/main.go

#**# Это контейнер с ранее собранным приложением**

FROM alpine:latest as server

WORKDIR /root/

COPY --from=builder /user-app/user-app-service .

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

RUN chown root:root user-app-service

CMD ["./user-app-service"]