FROM golang:alpine AS builder
COPY . /app
WORKDIR /app
RUN go build -o whoami

FROM alpine
WORKDIR /app
COPY --from=builder /app/whoami /app/
ENTRYPOINT ./whoami