# Docker 部署指南

## 部署一共分为两步

### 1. 编译 web 项目

> docker/view 下面已默认有编译好的文件；无修改 Web 需要、可跳过该步骤；

编译 web 项目可在当前项目根路径执行

> /bin/sh make.sh

### 2. 执行 compose up 编译镜像并运行服务

进入到 ./docker 目录下执行

> docker compose up -d

## 配置文件说明

默认服务端运行端口为: 8600

默认 Web 应用运行端口为: 8601

配置在 **./docker/docker-compose.yml** 文件中
需要修改的配置主要为

### 一. COS 相关配置,主要有

```
COS_SECRET_ID => example-id
COS_SECRET_KEY => example-key
COS_BUCKET_NAME => example
COS_APP_ID => 100000
COS_REGION => ap-shanghai
COS_CDN_URL => https://...
```

### 二. Server 访问地址

```
SERVER_ROOT
```

这里默认为 localhost:8600 、必须修改为外部网络访问 Server 时使用的地址、否则 web 端会无法 fetch 服务端的接口

> 例如使用 nginx 配置了域名 server.example.com:8600 映射本地 server 服务的地址 localhost:8600

此时 SERVER_ROOT 就应该改为 http://server.example.com:8600

### 三. COS 的 STS 临时密钥生成接口权限说明

web 项目在访问 server 端接口请求临时密钥时、STS 接口给予的权限如下

```
Action: []string{
	"name/cos:PostObject",
	"name/cos:PutObject",
	"name/cos:InitiateMultipartUpload",
	"name/cos:ListMultipartUploads",
	"name/cos:ListParts",
	"name/cos:UploadPart",
	"name/cos:CompleteMultipartUpload",
	"name/cos:GetBucket",
	"name/cos:GetObject",
	"name/cos:DeleteObject",
}
```

STS 授予允许操作的路径默认为 **/distribute/app**

通过修改 docker-compose.yml 中的 **COS_RES_DIR** 参数来修改操作路径
App 的安装文件会在此路径下存储、上传时会在次路径下生成 /tmp 文件夹作为临时文件上传路径
