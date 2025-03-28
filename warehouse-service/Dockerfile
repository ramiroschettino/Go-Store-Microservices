FROM golang:1.23-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod y go.sum primero para aprovechar el cache
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod tidy

# Copiar todo el contenido de tu proyecto local al contenedor
COPY . .

# Compilar el binario de tu aplicación
RUN go build -o warehouse-service ./cmd/warehouse-service

# Exponer el puerto necesario
EXPOSE 50051

# Comando para ejecutar el contenedor
CMD ["/app/warehouse-service"]
