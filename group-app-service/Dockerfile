FROM golang:1.19-alpine as builder

COPY . /group-app

WORKDIR /group-app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -o group-app-service ./cmd/service/main.go

#**# Это контейнер с ранее собранным приложением**

FROM alpine:latest as server

WORKDIR /root/

COPY --from=builder /group-app/group-app-service .

COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

RUN chown root:root group-app-service

CMD ["./group-app-service"]