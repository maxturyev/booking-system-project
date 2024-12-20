FROM golang:1.23.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

#COPY notifatication-svc/ .

RUN go mod tidy
RUN go build -o notifatication-svc

EXPOSE 9092

ENV KAFKA-BROKER=localhost:9092

CMD ["./notification-svc"]