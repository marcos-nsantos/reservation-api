FROM golang:1.22.3-alpine3.19 as builder
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/server/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
