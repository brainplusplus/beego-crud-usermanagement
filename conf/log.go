package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/astaxie/beego"
	"os"
)

func SetupLogging() {
	logrus.SetLevel(logrus.DebugLevel)
	if beego.AppConfig.String("runmode") == "prod" {
		log.Info("SETUP LOGGING")
		logPath := beego.AppConfig.String("log_path")

		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
		if err != nil {
			log.Fatal("Cannot log to file", err.Error())
		}

		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(file)
	}
}