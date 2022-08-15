package handler

import (
	"path"
	"time"
	"yayar/internal/conf"
	"yayar/internal/data"
	"yayar/pkg/sts"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Data *data.Data
	Conf *conf.AppConf
}
type UserHandlerOptions func(handler *UserHandler)

func (handler *UserHandler) InitUserHandler(opts ...UserHandlerOptions) {
	for _, opt := range opts {
		opt(handler)
	}
}
func InitUserHandlerWithData(data *data.Data) UserHandlerOptions {
	return func(handler *UserHandler) {
		handler.Data = data
	}
}
func InitUserHandlerWithConf(conf *conf.AppConf) UserHandlerOptions {
	return func(handler *UserHandler) {
		handler.Conf = conf
	}
}
func (handler *UserHandler) SingUp(ctx *gin.Context) {

}

func (handler *UserHandler) SignIn(ctx *gin.Context) {

}

func (handler *UserHandler) User(ctx *gin.Context) {

}

func (handler *UserHandler) CosTempToken(ctx *gin.Context) {
	appid := handler.Conf.COS_APP_ID
	bucket := handler.Conf.COS_BUCKET_NAME + "-" + appid
	resp := Response{}
	if appid == "" {
		resp.Serve(ctx, WithError("NOT AVAILABLE"))
		return
	}
	c := sts.NewClient(
		handler.Conf.COS_SECRET_ID,
		handler.Conf.COS_SECRET_KEY,
		nil, // sts.Host("默认域名sts.tencentcloudapi.com")
	)
	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          handler.Conf.COS_REGION,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					// 权限列表请看 https://cloud.tencent.com/document/product/436/31923
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
					},
					Effect: "allow",
					Resource: []string{
						// path.Join(handler.Conf.COS_RES_DIR, "/*") ==> /distribute/app/*
						"qcs::cos:ap-shanghai:uid/" + appid + ":" + bucket + path.Join(handler.Conf.COS_RES_DIR, "/*"),
					},
				},
			},
		},
	}
	res, err := c.GetCredential(opt)
	if err != nil {
		resp.Serve(ctx, WithError(err.Error()))
		return
	}
	d := struct {
		Data            *sts.CredentialResult
		COS_BUCKET_NAME string
		COS_APP_ID      string
		COS_REGION      string
		COS_CDN_URL     string
		COS_RES_DIR     string
	}{
		Data:            res,
		COS_BUCKET_NAME: handler.Conf.COS_BUCKET_NAME,
		COS_APP_ID:      handler.Conf.COS_APP_ID,
		COS_REGION:      handler.Conf.COS_REGION,
		COS_CDN_URL:     handler.Conf.COS_CDN_URL,
		COS_RES_DIR:     handler.Conf.COS_RES_DIR,
	}
	resp.Serve(ctx, WithData(d, nil))
}
