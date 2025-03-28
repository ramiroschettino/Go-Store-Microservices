FROM golang:1.23-alpine AS builder

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# 1. Copiar archivos de módulo primero
COPY go.mod go.sum ./

# 2. Descargar dependencias
RUN go mod download

# 3. Copiar el código fuente
COPY . .

# 4. Solución clave: Configurar el replace para checkout-service
# Esto asume que checkout-service está en el mismo nivel que warehouse-service
RUN go mod edit -replace github.com/ramiroschettino/Go-Store-Microservices/checkout-service=/go/pkg/mod/github.com/ramiroschettino/Go-Store-Microservices/checkout-service@v0.0.0

# 5. Compilar la aplicación
RUN go build -o warehouse-service ./cmd/warehouse-service

# --- Etapa de ejecución ---
FROM alpine:latest

WORKDIR /app

# Copiar el binario compilado
COPY --from=builder /app/warehouse-service .

# Exponer el puerto necesario
EXPOSE 50051

# Comando para ejecutar el contenedor
CMD ["/app/warehouse-service"]