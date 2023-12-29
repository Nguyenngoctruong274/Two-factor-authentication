FROM golang:1.20.7-alpine3.17
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main authentication/cmd/main/main.go
EXPOSE 8080
CMD ["/app/main"]