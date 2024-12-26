FROM golang:1.23.2

WORKDIR /app

COPY src/notification-svc/go.mod src/notification-svc/go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/src/notification-svc

RUN go mod tidy
RUN go build -o notification-svc main.go

CMD ["./notification-svc"]
