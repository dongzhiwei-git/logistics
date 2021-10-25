FROM golang:1.16-alpine
MAINTAINER King
RUN mkdir -p /logistics
WORKDIR /logistics
RUN go env -w GOPROXY=https://goproxy.cn,direct
#RUN git clone https://github.com/dongzhiwei-git/logistics.git
COPY . .
#WORKDIR /logistics/logistics
RUN go build -o logistics
EXPOSE 8000