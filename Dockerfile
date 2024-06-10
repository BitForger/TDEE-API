FROM golang:1.22-alpine as builder
LABEL authors="BitForger"
WORKDIR /app

COPY ./ /app/
RUN go mod download && go mod verify && go mod tidy
RUN CGO_ENABLED=0; go build -o /app/main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
RUN chmod +x /app/main

CMD ["sh", "-c", "/app/main"]