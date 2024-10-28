# Usa la imagen oficial de Go 1.21.6
FROM golang:1.21.6

# Establece el directorio de trabajo en /app
WORKDIR /app

# descarga  los modulos
COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env .env
EXPOSE 8080
# Compila la aplicación
RUN go build -o newsletter

# Ejecuta la aplicación al iniciar el contenedor
CMD ["./newsletter"]