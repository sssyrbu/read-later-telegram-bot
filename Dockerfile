FROM golang:1.21.6-alpine

WORKDIR /bot

RUN apk update

COPY ./ ./

RUN go mod download && go build main.go

CMD ["./main"]
