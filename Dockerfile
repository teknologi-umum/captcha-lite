FROM golang:1.18-bullseye AS builder

WORKDIR /usr/app

COPY . .

RUN go build .

FROM debian:bullseye

WORKDIR /app

COPY --from=builder /usr/app/main /app/main

CMD [ "./main" ]
