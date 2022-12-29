FROM golang:1.19 as base

RUN mkdir /grpc-api
WORKDIR /grpc-api
ADD . /grpc-api

EXPOSE 8080
EXPOSE 9090

RUN go build -o main

FROM debian:buster-slim 
COPY --from=base /grpc-api .

CMD ["./main"]