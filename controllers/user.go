package controllers

import (
	"github.com/astaxie/beego"
	"ScaleAffWeb/repositories"
	//"ScaleAffWeb/models"
	"github.com/astaxie/beego/session"
)
// UserController operations for User
type UserController struct {
	beego.Controller
	Session session.Store
}

func (c *UserController) Prepare() {
	c.Session = c.StartSession()
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Form", c.Form)
}

func (c *UserController) Index() {
	c.Data["Roles"],_ = repositories.GetRoleRepository().GetAll()
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user/index.tpl"
	//c.Layout = "templates/v1/admin.tpl"
}

func (c *UserController) Form() {
	//id := c.GetString("id")
	//log.Info(id)
	//c.Data["DataRoles"] = []models.Role{}
	//if id != "" {
	//	data, _ := repositories.GetUserRepository().GetById(id)
	//	c.Data["DataRoles"] = data.Roles
	//	log.Info(data)
	//	log.Info(data.Roles)
	//}
	//log.Info(c.Data["DataRoles"])
	c.Data["Roles"],_ = repositories.GetRoleRepository().GetAll()
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "user/form.tpl"
	//c.Layout = "templates/v1/admin.tpl"
}
