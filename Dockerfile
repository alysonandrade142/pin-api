# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copiar os arquivos go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Fazer o download das dependências
RUN go mod download

# Copiar todo o código-fonte para dentro da imagem
COPY . .

# Construir a aplicação Go com o main localizado em cmd/api/main.go
RUN go build -o main ./cmd/api

# Final stage: criar uma imagem final baseada em Alpine
FROM alpine:latest

WORKDIR /root/

# Copiar o binário gerado da fase anterior
COPY --from=builder /app/main .

# Expor a porta que a aplicação vai usar
EXPOSE 8080

# Comando que será executado ao rodar o container
CMD ["./main"]
