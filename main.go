package main

import (
	"github.com/astaxie/beego"
	"ScaleAffWeb/routers"
	"ScaleAffWeb/middlewares"
	"ScaleAffWeb/repositories"
	"ScaleAffWeb/util"
	"ScaleAffWeb/conf"
	"ScaleAffWeb/filters"
)

func main() {
	util.InitSnowflake()
	conf.SetupLogging()
	repositories.InitFactory()
	conf.SetupSession()
	routers.InitRoutes()
	filters.InitPathRuleRoutes()
	filters.InitAuthFilter()
	middlewares.InitJWT()
	beego.Run()
}