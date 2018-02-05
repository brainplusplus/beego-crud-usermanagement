package routers

import (
	"github.com/astaxie/beego"
	"ScaleAffWeb/controllers"
)

func InitUserRoutes(){
	beego.Router("/user/index", &controllers.UserController{}, "get:Index")
	beego.Router("/user/form", &controllers.UserController{}, "get:Form")
	beego.Router("/api/user/list", &controllers.UserApiController{}, "get:List")
	beego.Router("/api/user/get", &controllers.UserApiController{}, "get:Get")
	beego.Router("/api/user/delete", &controllers.UserApiController{}, "get:Delete")
	beego.Router("/api/user/save", &controllers.UserApiController{}, "post:Save")
	beego.Router("/api/user/update", &controllers.UserApiController{}, "post:Update")
}
