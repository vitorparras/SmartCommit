# Etapa base para desenvolvimento
FROM golang:1.20 AS base

# Instalar o CompileDaemon para hot-reloading
RUN go install github.com/githubnemo/CompileDaemon@latest

# Configuração do diretório de trabalho
WORKDIR /app

# Copiar arquivos do projeto
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Expor a porta padrão
EXPOSE 8080

# Comando para hot-reloading
CMD ["CompileDaemon", "--build=go build -o backend .", "--command=./backend"]
