FROM golang:1.22-bookworm

WORKDIR /bot

RUN apt update && apt install redis-server -y

COPY ./ ./
RUN redis-server --daemonize yes 
RUN go mod download && go build main.go

CMD ["./main"]
