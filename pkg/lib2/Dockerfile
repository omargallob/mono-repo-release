# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o lib2 ./...

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/lib2 ./lib2
ENTRYPOINT ["./lib2"]
