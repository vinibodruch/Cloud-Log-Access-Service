FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o cloud-log-api -ldflags "-s -w" .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/cloud-log-api .

EXPOSE 8080

ENTRYPOINT ["./cloud-log-api"]
