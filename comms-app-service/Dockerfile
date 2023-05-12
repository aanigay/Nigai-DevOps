FROM golang:1.19-alpine as builder

COPY . /comms-app

WORKDIR /comms-app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o comms-app-service ./cmd/service/main.go

#**# Это контейнер с ранее собранным приложением**

FROM alpine:latest as server

WORKDIR /root/

COPY --from=builder /comms-app/comms-app-service .

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

RUN chown root:root comms-app-service

CMD ["./comms-app-service"]