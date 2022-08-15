package conf

import (
	"os"
	"strconv"
)

type AppConf struct {
	APP_MODE        string
	PORT            string
	MYSQL_USER      string
	MYSQL_DATABASE  string
	MYSQL_PASSWORD  string
	MYSQL_HOST      string
	MYSQL_PORT      string
	REDIS_DATABASE  string
	REDIS_PORT      string
	MAX_RESERVE_VER int
	COS_SECRET_ID   string
	COS_SECRET_KEY  string
	COS_BUCKET_NAME string
	COS_APP_ID      string
	COS_REGION      string
	COS_CDN_URL     string
	INDEX_WEB_PORT  string
	VIEW_PATH       string
	COS_RES_DIR     string
}

func NewConf() AppConf {
	max_reserve_ver_str := os.Getenv("MAX_RESERVE_VER")
	max_reserve_ver, _ := strconv.Atoi(max_reserve_ver_str)
	var conf = AppConf{
		APP_MODE:        os.Getenv("APP_MODE"),
		PORT:            os.Getenv("PORT"),
		MYSQL_USER:      os.Getenv("MYSQL_USER"),
		MYSQL_DATABASE:  os.Getenv("MYSQL_DATABASE"),
		MYSQL_PASSWORD:  os.Getenv("MYSQL_PASSWORD"),
		MYSQL_HOST:      os.Getenv("MYSQL_HOST"),
		MYSQL_PORT:      os.Getenv("MYSQL_PORT"),
		REDIS_DATABASE:  os.Getenv("REDIS_DATABASE"),
		REDIS_PORT:      os.Getenv("REDIS_PORT"),
		MAX_RESERVE_VER: max_reserve_ver,
		COS_SECRET_ID:   os.Getenv("COS_SECRET_ID"),
		COS_SECRET_KEY:  os.Getenv("COS_SECRET_KEY"),
		COS_BUCKET_NAME: os.Getenv("COS_BUCKET_NAME"),
		COS_APP_ID:      os.Getenv("COS_APP_ID"),
		COS_REGION:      os.Getenv("COS_REGION"),
		COS_CDN_URL:     os.Getenv("COS_CDN_URL"),
		INDEX_WEB_PORT:  os.Getenv("INDEX_WEB_PORT"),
		VIEW_PATH:       os.Getenv("VIEW_PATH"),
		COS_RES_DIR:     os.Getenv("COS_RES_DIR"),
	}
	return conf
}
