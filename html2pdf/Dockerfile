FROM  html2pdf:v0.2
MAINTAINER "xiongchuanhong"
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release\
    EXPOSE=8081
WORKDIR /go/src/html2pdf
ADD . /go/src/html2pdf
RUN go build
ENTRYPOINT ["./html2pdf"]