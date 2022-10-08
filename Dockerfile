FROM golang:1.19-bullseye AS builder

WORKDIR /usr/app

COPY . .

RUN go build .

FROM debian:bullseye

WORKDIR /app

COPY --from=builder /usr/app/teknologi-umum-bot /app/teknologi-umum-bot

CMD [ "./teknologi-umum-bot" ]
