FROM golang:latest

WORKDIR /app

COPY server.go .

RUN go build -o server server.go

EXPOSE 8081

CMD ["./server"]
