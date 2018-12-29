FROM golang:1.9.4

MAINTAINER "115143589@qq.com"


WORKDIR /go/src/statutory-holidays

COPY . .

RUN go build /go/src/statutory-holidays/main.go \
    && apt-get update

CMD ["bash", "-c","go run /go/src/statutory-holidays/main"]

