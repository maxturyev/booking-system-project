FROM golang:1.23.2

WORKDIR /app

COPY src/booking-svc/go.mod src/booking-svc/go.sum ./
COPY src/grpc ../grpc
RUN go mod download

COPY . .

WORKDIR /app/src/booking-svc

RUN go mod tidy
RUN go build -o booking-svc main.go

EXPOSE 9091

CMD ["./booking-svc"]
