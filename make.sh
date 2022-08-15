#!/bin/sh
echo =============== READY TO BUILD PROJECT =================

function isMac() {
    if [[ "$(uname)" == *Darwin* ]]; then
        return 1
    else
        return 0
    fi
}
isMac
isMacOS=$?
wd=$(pwd)
echo "💡 开始编译 view/downloader "
cd $wd/view/downloader && yarn build

echo "💡 开始编译 view/index "
cd $wd/view/index && yarn build
cd $wd

loadWebSource() {
    echo "💡 开始迁移 web项目源文件 "
    if [ -d ./docker/web ]; then
        if [ -d ./docker/web/view ]; then
            rm -rf ./docker/web/view
        fi
    fi
    mkdir -p ./docker/web/view
    cp -R view/status ./docker/web/view/
    if [ -d view/index/build ]; then
        cp -R view/index/build ./docker/web/view/index/
    else 
        echo "index网页项目未编译、请先检查并编译"
        exit -1
    fi
    if [ -d view/downloader/dist ]; then
        cp -R view/downloader/dist ./docker/web/view/downloader/
    else 
        echo "downloader网页项目未编译、请先检查并编译"
        exit -2
    fi
}

loadWebSource

echo "=============== 🍉 MAKE DONE ================="
