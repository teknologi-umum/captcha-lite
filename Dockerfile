FROM golang:1.21-bookworm AS builder

WORKDIR /usr/app

COPY . .

RUN go build -o captcha-lite .

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y curl ca-certificates openssl

WORKDIR /app
COPY . .
COPY --from=builder /usr/app/captcha-lite /app/captcha-lite

CMD [ "./captcha-lite" ]
