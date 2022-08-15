package utils

import (
	"fmt"
	"strings"
)

type DatabaseUtil struct{}

type DSNConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
	Pairs    map[string]string
}

func (d DatabaseUtil) DefaultDSNPair() map[string]string {
	return map[string]string{
		"charset":   "utf8mb4",
		"parseTime": "True",
		"loc":       "Asia%2FShanghai",
	}
}
func (d DatabaseUtil) DSN(conf DSNConfig) string {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	if len(conf.Pairs) > 0 {
		par := ""
		for k, v := range conf.Pairs {
			par += fmt.Sprintf("%s=%s&", k, v)
		}
		par = strings.TrimRight(par, "&")
		url = fmt.Sprintf("%s?%s", url, par)
	}
	return url
}
