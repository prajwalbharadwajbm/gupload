FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gupload ./cmd/server

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/gupload .
RUN apk --no-cache add ca-certificates && \
    addgroup -S app && adduser -S app -G app
RUN mkdir -p /app/uploads && \
    chown -R app:app /app
USER app
EXPOSE 8080
CMD ["./gupload"]