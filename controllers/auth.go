package controllers

import (
	"github.com/astaxie/beego"
)

// UserController operations for User
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Index", c.Login)
}

func (c *AuthController) Login() {
	c.TplName = "auth/login.tpl"
}