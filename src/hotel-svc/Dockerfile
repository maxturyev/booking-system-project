FROM golang:1.23.2


WORKDIR /app

COPY go.mod go.sum ./
COPY ../grpc/go.mod ../grpc/go.sum ../

RUN go mod download

#COPY hotel-svc/ .

RUN go mod tidy
RUN go build -o hotel-svc


EXPOSE 9090
EXPOSE 50051

ENV KAFKA-BROKER=kafka:9092

CMD [ "./hotel-svc" ]