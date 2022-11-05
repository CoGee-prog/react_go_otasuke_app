FROM golang:1.19.2-alpine

RUN apk add --update &&  apk add git

COPY go.mod /go/src/app/go.mod
COPY go.sum /go/src/app/go.sum
WORKDIR /go/src/app
COPY ./Server /go/src/app

RUN go get github.com/cosmtrek/air@v1.27.3

CMD ["air", "-c", ".air.toml"]**