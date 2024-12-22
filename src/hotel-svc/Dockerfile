#FROM golang:1.23.2
#
#WORKDIR src/app
#
#COPY grpc grpc
#COPY hotel-svc/go.mod hotel-svc/go.sum ./
#
#RUN go mod download
#
#COPY hotel-svc/. src/app
#
#RUN go mod tidy
#RUN go build -o hotel-svc
#
#EXPOSE 9090
#EXPOSE 50051
#
#CMD ["./hotel-svc"]

FROM golang:1.23.2 AS builder

WORKDIR /app

COPY src/hotel-svc/go.mod src/hotel-svc/go.sum ./
RUN go mod download

COPY . /app

WORKDIR /app/src/app1

RUN go mod tidy
RUN go build -o hotel-svc main.go

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=builder /app/src/hotel-svc/hotel-svc /app/

EXPOSE 9090
EXPOSE 50051
CMD ["/app/hotel-svc"]
