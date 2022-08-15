package utils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	cos "github.com/tencentyun/cos-go-sdk-v5"
)

type TTCos struct {
	Client *cos.Client
}

//COS错误定义部分
const (
	CosUninitialClientError = "COS客户端未初始化"
)

func IsCosUninitialClientError(err error) bool {
	return err.Error() == CosUninitialClientError
}

// COS初始化部分
// 提供两个初始化可选方法 InitWithPersistKeys 和 InitWithTempKeys
// 前者为使用永久密钥方式、后者为使用临时密钥
// 初始化完成后应通过判断 Client 是否为nil 来确定初始化是否成功
type TTCosInitOption func(ttcos *TTCos)

func (ttcos *TTCos) Init(opts ...TTCosInitOption) {
	for _, opt := range opts {
		opt(ttcos)
	}
}

type TTCOSConfig struct {
	BucketName string
	AppID      string
	Region     string
	SecretID   string
	SecretKey  string
}
type TTCOSTmpConfig struct {
	Base         TTCOSConfig
	SessionToken string
}

func InitWithPersistKeys(config TTCOSConfig) TTCosInitOption {
	return func(ttcos *TTCos) {
		bucketURL, _ := url.Parse(fmt.Sprintf("https://%s-%s.cos.%s.myqcloud.com",
			config.BucketName,
			config.AppID,
			config.Region))
		serviceURL, _ := url.Parse(fmt.Sprintf("https://cos.%s.myqcloud.com", config.Region))
		b := &cos.BaseURL{
			BucketURL:  bucketURL,
			ServiceURL: serviceURL,
		}
		ttcos.Client = cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  config.SecretID,
				SecretKey: config.SecretKey,
			},
		})
	}
}

func InitWithTempKeys(config TTCOSTmpConfig) TTCosInitOption {
	return func(ttcos *TTCos) {
		bucketURL, _ := url.Parse(fmt.Sprintf("https://%s-%s.cos.%s.myqcloud.com",
			config.Base.BucketName,
			config.Base.AppID,
			config.Base.Region))
		b := &cos.BaseURL{BucketURL: bucketURL}
		ttcos.Client = cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:     config.Base.SecretID,
				SecretKey:    config.Base.SecretKey,
				SessionToken: config.SessionToken,
			},
		})
	}
}

// COS操作部分
// 包括上传对象、查询对象列表、下载对象、删除对象等等
// objectPath 即 cos存储文件的key、也是其在cos云端显示的存储路径

//上传对象
func (ttcos *TTCos) UploadObject(ctx context.Context, objectPath, localPath string, options *cos.MultiUploadOptions) (*cos.CompleteMultipartUploadResult, *cos.Response, error) {
	if ttcos.Client == nil {
		return nil, nil, errors.New(CosUninitialClientError)
	}
	return ttcos.Client.Object.Upload(ctx, objectPath, localPath, options)
}

//删除对象
func (ttcos *TTCos) DeleteObject(ctx context.Context, objectPath string) (*cos.Response, error) {
	if ttcos.Client == nil {
		return nil, errors.New(CosUninitialClientError)
	}
	return ttcos.Client.Object.Delete(ctx, objectPath)
}

//移动对象
//fromURL为对象的完整访问路径，fromPath为对象的COS存储路径,destPath为对象的目标COS存储路径
func (ttcos *TTCos) MoveObject(ctx context.Context, sourceURL, fromPath, destPath string) (*cos.Response, error) {
	amendedSourceURL := strings.TrimPrefix(sourceURL, "https://")
	if ttcos.Client == nil {
		return nil, errors.New(CosUninitialClientError)
	}
	_, _, err := ttcos.Client.Object.Copy(ctx, destPath, amendedSourceURL, nil)
	if err == nil {
		res, err := ttcos.Client.Object.Delete(ctx, fromPath, nil)
		if err != nil {
			// Error
			return nil, err
		}
		return res, nil
	}
	return nil, err
}

//查询对象是否存在
func (ttcos *TTCos) IsExist(ctx context.Context, objectPath string) (bool, error) {
	if ttcos.Client == nil {
		return false, errors.New(CosUninitialClientError)
	}
	ok, err := ttcos.Client.Object.IsExist(ctx, objectPath)
	if err == nil && ok {
		//exist
		return true, nil
	} else if err != nil {
		//failed somehow
		return false, err
	} else {
		//unexist
		return false, nil
	}
}
