# Usar una imagen base de Go
FROM golang:1.20-alpine

# Crear un directorio de trabajo
WORKDIR /app

# Copiar el archivo del servidor a /app
COPY server.go .

# Compilar el servidor
RUN go build -o server .

# Exponer el puerto 8080
EXPOSE 8080

# Ejecutar el servidor
CMD ["./server"]
