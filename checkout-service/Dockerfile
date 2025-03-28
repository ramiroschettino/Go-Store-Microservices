FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o checkout-service .
EXPOSE 50052
CMD ["/app/checkout-service"]
