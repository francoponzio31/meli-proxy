# Etapa 1: Compilación
FROM golang:1.23 AS builder

# Configurar el directorio de trabajo
WORKDIR /app

# Copiar los archivos del proyecto
COPY . .

# Descargar las dependencias y compilarlas
RUN go mod download

# Compilar el binario
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o meli_proxy main.go

# Etapa 2: Imagen final, minimizada
FROM alpine:latest

# Configurar el directorio de trabajo
WORKDIR /root/

# Copiar el binario de la etapa de construcción
COPY --from=builder /app/meli_proxy .

# Exponer el puerto de la aplicación
EXPOSE ${PORT}

# Ejecutar el binario
CMD ["./meli_proxy"]