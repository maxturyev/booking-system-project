FROM golang:1.23.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

#COPY booking-svc/ .

RUN go mod tidy
RUN go build -o booking-svc

EXPOSE 9091

ENV KAFKA-BROKER=localhost:9092
ENV HOTEL-SERVER-ADDR=localhost:50051

CMD ["./booking-svc"]