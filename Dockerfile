# Stage 1
FROM golang:1.21.5-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/bin/echo-app-1 cmd/server/main.go

# Stage 2
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/bin/echo-app-1 .
EXPOSE 8080
CMD ["./echo-app-1"]
