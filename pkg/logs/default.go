package logs

import (
	"fmt"
	"os"
	"path"

	"github.com/golang-module/carbon/v2"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func Info(args ...interface{}) {
	logger.Info(args...)
}
func Warn(args ...interface{}) {
	logger.Warn(args...)
}
func Error(args ...interface{}) {
	logger.Error(args...)
}
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

type LoggerConig struct {
	OutputDir *string
	Level     log.Level
}

func InitLogger(conf LoggerConig) {
	logger = log.New()
	logger.SetLevel(conf.Level)
	if conf.OutputDir != nil {
		err := os.MkdirAll(*conf.OutputDir, 0775)
		if err != nil {
			Fatal("Log Output Dir Error : ", err)
		}
		today := carbon.Now().ToDateString()
		fpath := path.Join(*conf.OutputDir, fmt.Sprintf("%s.log", today))
		if _, err := os.Stat(fpath); os.IsNotExist(err) {
			//file not exit
			nf, err := os.Create(fpath)
			if err != nil {
				Fatal(err)
			}
			logger.Out = nf
		} else {
			//file exit || other error
			if err != nil {
				Fatal(err)
			} else {
				f, err := os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
				if err != nil {
					Fatal(err)
				}
				logger.Out = f
			}
		}
	}
}
