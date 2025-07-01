FROM golang:1.22.2-alpine
 
WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go build -o app

CMD ["./app"]