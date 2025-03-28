# Usar la imagen base de Golang
FROM golang:1.23-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar go.mod y go.sum primero para manejar dependencias
COPY go.mod go.sum ./

# Ejecutar go mod tidy para gestionar las dependencias
RUN go mod tidy

# Copiar todo el código fuente al contenedor
COPY . .

# Establecer el directorio de trabajo a la carpeta donde está cmd
WORKDIR /app/cmd

# Compilar el binario usando el archivo main.go en cmd
RUN go build -o /app/user-service

# Exponer el puerto en el que el servicio estará corriendo
EXPOSE 50054

# Comando para ejecutar el servicio
CMD ["/app/user-service"]
