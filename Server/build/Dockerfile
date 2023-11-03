FROM golang:1.19.2-alpine3.16

RUN apk add --update --no-cache tzdata && apk add git

ENV TZ=Asia/Tokyo

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@v1.41.0

CMD ["air", "-c", ".air.toml"]**