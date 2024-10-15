FROM golang:alpine AS builder

FROM golang:alpine AS builder

RUN apk add --no-cache git

COPY go.mod /app/
COPY whoami.go /app/
WORKDIR /app

RUN go mod tidy
RUN go build -o whoami

FROM alpine:latest

RUN adduser -D appuser

WORKDIR /app
COPY --from=builder /app/whoami /app/

USER appuser

CMD ["./whoami"]
