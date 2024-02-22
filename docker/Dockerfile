# Etapa de compilación para el backend de Go
FROM golang:1.20 AS go-builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
COPY ./.env /root/.env
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-binary ./cmd/server/main.go

# Etapa final para ejecutar la aplicación
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copia el binario compilado de Go desde la etapa de compilación de Go
COPY --from=go-builder /root/.env /root/.env
COPY --from=go-builder /go-binary .

# Expone el puerto en el que tu servidor Go escucha
EXPOSE 8080
CMD ["./go-binary"]
