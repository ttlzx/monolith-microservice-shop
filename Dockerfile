FROM golang:1.9

RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/ttlzx/monolith-microservice-shop
COPY . .

RUN dep ensure
RUN go get github.com/cespare/reflex
