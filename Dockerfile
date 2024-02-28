# Build stage
FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .

COPY migrations ./migrations

EXPOSE 8080
CMD [ "/app/main" ]
