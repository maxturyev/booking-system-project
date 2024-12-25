FROM golang:1.23.2

WORKDIR /app

COPY src/hotel-svc/go.mod src/hotel-svc/go.sum ./
COPY src/grpc ../grpc
RUN go mod download

COPY . .

WORKDIR /app/src/hotel-svc

RUN go mod tidy
RUN go build -o hotel-svc main.go

EXPOSE 9090
EXPOSE 50051

CMD ["./hotel-svc"]
