FROM golang:latest

RUN mkdir -p /go/src/demo
WORKDIR /go/src/demo

RUN go get -v github.com/prometheus/client_golang/prometheus

COPY main.go .
RUN go build -o main

EXPOSE 8080

CMD [ "./main" ]