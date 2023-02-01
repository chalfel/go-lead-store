FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 8088
RUN CGO_ENABLED=0 GOOS=linux go build -o go-lead-store .

FROM alpine:3.17.0
WORKDIR /app
COPY --from=builder /app/go-lead-store .

CMD ["./go-lead-store", "server", "serve"]
