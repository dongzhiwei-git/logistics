FROM golang
MAINTAINER King
RUN mkdir -p /logistics
WORKDIR /logistics
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN git clone https://gitee.com/zmr_0210/logistics.git
WORKDIR /logistics/logistics
RUN go build -o logistics
EXPOSE 8000