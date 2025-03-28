FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o user-service .
EXPOSE 50054
CMD ["/app/user-service"]
