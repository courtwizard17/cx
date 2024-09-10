FROM golang:latest

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN go build -o server

RUN mkdir docs
COPY ./docs/*.pdf ./docs/

EXPOSE 8080

CMD ["./server"]
