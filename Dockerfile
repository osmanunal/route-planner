FROM golang:1.23-alpine as builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main api/cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

RUN apk update && \
    apk add --no-cache tzdata ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

CMD ["./main"] 