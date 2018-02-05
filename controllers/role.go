package controllers

import (
	"github.com/astaxie/beego"
)

// RoleController operations for Role
type RoleController struct {
	beego.Controller
}

// URLMapping ...
func (c *RoleController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Form", c.Form)
}

func (c *RoleController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "role/index.tpl"
	//c.Layout = "templates/v1/admin.tpl"
}

func (c *RoleController) Form() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "role/form.tpl"
	//c.Layout = "templates/v1/admin.tpl"
}
