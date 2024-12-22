#FROM golang:1.23.2
#
#WORKDIR /app
#
#COPY go.mod go.sum ./
#
#RUN go mod download
#
#COPY . .
#
#RUN go mod tidy
#RUN go build -o notifatication-svc
#
#EXPOSE 9092

FROM golang:1.23.2 AS builder

WORKDIR /app

COPY src/notification-svc/go.mod src/notification-svc/go.sum ./
RUN go mod download

COPY . /app

WORKDIR /app/src/notification-svc

RUN go mod tidy
RUN go build -o notification-svc main.go


FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/src/notification-svc/notification-svc /app/

ENV KAFKA-BROKER=localhost:9092
EXPOSE 9092
CMD ["/app/notification-svc"]
