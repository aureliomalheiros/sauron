FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build -o open-notify cmd/main.go


FROM alpine:3.16

COPY --from=builder /app/open-notify /app/open-notify
EXPOSE 7777
WORKDIR /app

CMD [ "./open-notify" ]
