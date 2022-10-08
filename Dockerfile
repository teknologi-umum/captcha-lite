FROM golang:1.19-bullseye AS builder

WORKDIR /usr/app

COPY . .

RUN go build -o captcha-lite .

FROM debian:bullseye

WORKDIR /app

COPY --from=builder /usr/app/captcha-lite /app/captcha-lite

CMD [ "./captcha-lite" ]
