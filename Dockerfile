FROM golang:1.20-bullseye AS builder

WORKDIR /usr/app

COPY . .

RUN go build -o captcha-lite .

FROM debian:bullseye

RUN apt-get update && apt-get install -y curl ca-certificates openssl

WORKDIR /app

COPY --from=builder /usr/app/captcha-lite /app/captcha-lite

CMD [ "./captcha-lite" ]
