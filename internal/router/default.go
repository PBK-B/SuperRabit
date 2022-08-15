package router

import (
	"os"
	"path"
	"yayar/internal/conf"
	"yayar/internal/data"
	"yayar/internal/handler"
	"yayar/pkg/middleware"
	"yayar/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
)

func InitRouter(r *gin.Engine, da *data.Data, conf conf.AppConf) {
	perCosConfig := utils.TTCOSConfig{
		BucketName: conf.COS_BUCKET_NAME,
		AppID:      conf.COS_APP_ID,
		Region:     conf.COS_REGION,
		SecretID:   conf.COS_SECRET_ID,
		SecretKey:  conf.COS_SECRET_KEY,
	}
	userhdl := handler.UserHandler{}
	userhdl.InitUserHandler(
		handler.InitUserHandlerWithData(da),
		handler.InitUserHandlerWithConf(&conf),
	)
	apphdl := handler.AppHandler{}
	apphdl.InitAppHandler(
		handler.InitAppHandlerWithData(da),
		handler.InitAppHandlerWithConf(&conf),
		handler.InitAppHandlerWithTmpDir("tmp", nil),
		handler.InitAppHandlerWithCOS(&perCosConfig, nil),
	)

	r.Use(middleware.CorsMiddleware())

	pwd, _ := os.Getwd()
	vr := path.Join(pwd, "view")
	if conf.APP_MODE != "debug" {
		vr = conf.VIEW_PATH
	} //else 在debug模式下、与项目根路径下运行应用，网页应用路径在 ./view 下、因此使用pwd
	r.LoadHTMLFiles(path.Join(vr, "status/error.html"))
	r.LoadHTMLFiles(path.Join(vr, "downloader/downloader.html"))
	r.Static("/assets", path.Join(vr, "downloader/assets"))

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.Writer.WriteString("pong ~ " + carbon.Now().ToString())
	})
	user := r.Group("user")
	user.GET("cos", userhdl.CosTempToken)

	app := r.Group("app")
	app.GET("list", apphdl.List)
	app.POST("create", apphdl.Create)
	app.GET("delete", apphdl.Delete)
	app.GET("versions", apphdl.ListVersions)
	app.GET("verifyaccess", apphdl.VerifyAccessCode)
	app.GET("installdetail", apphdl.InstallDetail)
	app.GET("/install/:name/*version", apphdl.Install)

}
