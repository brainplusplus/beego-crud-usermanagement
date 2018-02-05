package routers

import (
	"github.com/astaxie/beego"
	"ScaleAffWeb/controllers"
)

func InitRoleRoutes(){
	beego.Router("/role/index", &controllers.RoleController{}, "get:Index")
	beego.Router("/role/form", &controllers.RoleController{}, "get:Form")
	beego.Router("/api/role/list", &controllers.RoleApiController{}, "get:List")
	beego.Router("/api/role/get", &controllers.RoleApiController{}, "get:Get")
	beego.Router("/api/role/delete", &controllers.RoleApiController{}, "get:Delete")
	beego.Router("/api/role/save", &controllers.RoleApiController{}, "post:Save")
	beego.Router("/api/role/update", &controllers.RoleApiController{}, "post:Update")
}