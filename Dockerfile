FROM golang:latest

WORKDIR /project
COPY . .

RUN go mod download
RUN go build main.go

CMD ["./main"]