#FROM golang:1.23.2
#
#WORKDIR /app
#
#COPY go.mod go.sum ./
#
#RUN go mod download
#
#RUN go mod tidy
#RUN go build -o booking-svc
#
#EXPOSE 9091
#
#ENV KAFKA-BROKER=localhost:9092
#ENV HOTEL-SERVER-ADDR=localhost:50051
#
#CMD ["./booking-svc"]

FROM golang:1.23.2 AS builder

WORKDIR /app

COPY src/booking-svc/go.mod src/booking-svc/go.sum ./
COPY src/grpc ../grpc
RUN go mod download

COPY . /app

WORKDIR /app/src/booking-svc

RUN go mod tidy
RUN go build -o booking-svc main.go

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/src/booking-svc/booking-svc /app/

ENV KAFKA-BROKER=localhost:9092
ENV HOTEL-SERVER-ADDR=localhost:50051
EXPOSE 9091
CMD ["/app/booking-svc"]
