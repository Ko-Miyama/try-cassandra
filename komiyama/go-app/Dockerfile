FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod .
COPY main.go .
RUN go mod tidy
RUN go build -o app main.go
EXPOSE 8080
CMD ["/app/app"]
