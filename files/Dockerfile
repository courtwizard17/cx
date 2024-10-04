FROM golang:latest

WORKDIR /app

RUN go build -o server

RUN mkdir docs
COPY ./docs/*.pdf ./docs/

EXPOSE 8080

CMD ["./server"]
