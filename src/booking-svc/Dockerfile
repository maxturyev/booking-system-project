FROM golang:1.23.2 AS builder

RUN apt-get update
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/app
COPY . .
RUN go install

FROM scratch
COPY --from=builder /go/bin/app .
ENTRYPOINT ["./app"]

#RUN go mod download
#
#RUN go mod tidy
#RUN go build -o booking-svc

#EXPOSE 9091

#ENV KAFKA-BROKER=localhost:9092
#ENV HOTEL-SERVER-ADDR=localhost:50051

#CMD ["./booking-svc"]