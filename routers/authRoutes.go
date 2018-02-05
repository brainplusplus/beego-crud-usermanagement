package routers

import (
	"github.com/astaxie/beego"
	"ScaleAffWeb/controllers"
)

func InitAuthRoutes(){
	beego.Router("/auth/login", &controllers.AuthController{}, "get:Login")
	beego.Router("/auth/savesession", &controllers.AuthApiController{}, "get:SaveSession")
	beego.Router("/api/auth/login", &controllers.AuthApiController{}, "post:Login")
	beego.Router("/api/auth/validate", &controllers.AuthApiController{}, "get:Validate")
}