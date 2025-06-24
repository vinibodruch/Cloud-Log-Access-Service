FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && \
    CGO_ENABLED=0 go install -tags netgo -ldflags '-s -w'

COPY . .

# -ldflags "-s -w": Reduz o tamanho do binário removendo tabelas de símbolos e informações de depuração.
# CGO_ENABLED=0: Essencial para construir um binário estático que pode ser copiado para uma imagem scratch/alpine.
# -tags netgo: Garante que a biblioteca net padrão seja usada, em vez de depender de bibliotecas C para DNS, etc.
RUN CGO_ENABLED=0 go build -o /app/cloud-log-api -ldflags "-s -w" .

# --- 

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/cloud-log-api .

EXPOSE 8080

ENTRYPOINT ["/app/minhaapp"]

