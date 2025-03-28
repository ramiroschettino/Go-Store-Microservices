FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o warehouse-service .
EXPOSE 50051
CMD ["/app/warehouse-service"]
