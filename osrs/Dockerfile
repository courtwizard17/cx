FROM golang:latest

WORKDIR /app

COPY server.go .

RUN go build -o /usr/local/bin/server server.go && chmod +x /usr/local/bin/server

EXPOSE 8081

CMD ["/usr/local/bin/server"]
