FROM golang:1.18.3

#context在docker-compose.yml指定为 . 因此这里对应的工作路径是 ../docker/下
ENV GO111MODULE=on GOPROXY=https://goproxy.cn
WORKDIR /go/src/super_rabit

CMD [ "/bin/sh","boot.sh" ]