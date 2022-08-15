package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"yayar/internal/conf"
	"yayar/internal/data"
	"yayar/internal/router"
	"yayar/pkg/logs"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func main() {
	//读取配置文件
	cf := conf.NewConf()
	//配置运行模式
	isDebug := cf.APP_MODE == "debug"
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	//配置日志
	pwd, _ := os.Getwd()
	var logDirPtr *string
	if !isDebug {
		logDir := path.Join(pwd, "log")
		logDirPtr = &logDir
	}
	logs.InitLogger(logs.LoggerConig{
		OutputDir: logDirPtr,
		Level:     logrus.DebugLevel,
	})
	//配置数据库连接
	mysql_conf := data.ConfigMysql{
		User:     cf.MYSQL_USER,
		Password: cf.MYSQL_PASSWORD,
		Database: cf.MYSQL_DATABASE,
		Host:     &cf.MYSQL_HOST,
		Port:     &cf.MYSQL_PORT,
	}
	da, err := data.NewData(data.InitialConf{
		Mysql: &mysql_conf,
	})
	if err != nil {
		log.Fatalln(err)
	}
	//配置路由
	engine := gin.Default()
	router.InitRouter(engine, da, cf)

	//运行应用
	log.Printf("🚀 App is running at port: " + cf.PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", cf.PORT), engine)
}
