package repositories

import (
	"github.com/astaxie/beego"
	"ScaleAffWeb/logging"
)

var (
	datasource   *DataSource
	dbDriver   	 string
	dbUrl     	 string
	maxIdle		 int
	maxConn		 int
)

var log = logging.MustGetLogger()

func InitFactory() error {
	var err error
	maxIdle, _ 	:= beego.AppConfig.Int("db_max_idle")
	maxConn, _ 	:= beego.AppConfig.Int("db_max_conn")
	dbDriver 	:= beego.AppConfig.String("db_driver")
	dbUrl 		:= beego.AppConfig.String("db_url")

	datasource, err = NewDatabaseConnectionWithConnectionPool(dbDriver, dbUrl,maxIdle,maxConn)
	if err != nil {
		return err
	}
	log.Info("Database Connection Started")

	return nil
}
