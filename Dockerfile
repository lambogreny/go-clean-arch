FROM golang:alpine

# Setando algumas variáveis de ambiente para a imagem
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    SERVER_PORT=:8080

#Movendo o diretório de trabalho para /build
WORKDIR /build

# Copiando os arquivos e instalando as dependências do projeto
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copiando o código para o container
COPY . .

# Realizando o build do app
RUN go build -o main .

# Movendo o arquivo binário para a pasta /dist
WORKDIR /dist

# Copiando o binário para a pasta /main
RUN cp /build/main .

# Expondo a porta do server
EXPOSE 8080

# Executando o container
CMD ["/dist/main"]