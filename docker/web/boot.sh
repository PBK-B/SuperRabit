#!/bin/sh

# WORKDIR /volume/common/

sed -i "s~${ORIGIN}~${SERVER_ROOT}~g" `grep -rl $ORIGIN ./view`

if [ -e ./view/downloader/index.html ];then
    cd ./view/downloader && mv index.html downloader.html
    ls -al
    sed -i "s~</title>~</title><script>globalThis.id={{ .id }}</script>~g" ./downloader.html
fi

cd /go/src/app/

if [ ! -e go.mod  ];then
    go mod init
fi
go mod tidy && go run main.go