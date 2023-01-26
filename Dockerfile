FROM golang:alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.12.0
WORKDIR /app
COPY --from=builder /app/main /app/
COPY env.example /app/

EXPOSE 8080
CMD ["/app/main"]