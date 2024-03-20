FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux go build -o main .

ENTRYPOINT ["./main"]
