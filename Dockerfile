FROM golang:1.16-alpine
MAINTAINER liujun-211
WORKDIR /opt
ENV GOPROXY=https://goproxy.cn
COPY . /opt/
RUN go mod download \
&& go build -ldflags "-s -w" -o server
ENTRYPOINT ["./server"]
