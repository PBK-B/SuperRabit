FROM golang:1.18.3

# 数据卷: path/to/web/view => /volume/common/view

ENV GO111MODULE=on GOPROXY=https://goproxy.cn
WORKDIR /volume/common/

CMD [ "/bin/sh" ,"boot.sh" ]