
FROM golang:latest
MAINTAINER "kl023011@gmail.com"
WORKDIR /go/src/http
COPY . /go/src/http
RUN go build .
EXPOSE 9091
ENTRYPOINT ["./http"]
