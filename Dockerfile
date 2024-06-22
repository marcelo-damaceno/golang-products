FROM golang:1.21.6

WORKDIR /app

COPY . .

RUN GOOS=windows GOARCH=386 go build -o server

