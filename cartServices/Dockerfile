FROM golang:1.22-alpine3.19

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go get github.com/rabbitmq/amqp091-go

EXPOSE 5500

CMD ["go", "run", "main.go"]