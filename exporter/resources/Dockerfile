FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . /app

RUN go build  -o resources cmd/main.go


FROM alpine:3.16

COPY --from=builder /app/resources /app/resources
EXPOSE 7788
WORKDIR /app

CMD ["./resources"]
