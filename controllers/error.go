package controllers

import (
	"github.com/astaxie/beego"
	"ScaleAffWeb/util"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) NotAuth() {
	url := c.Ctx.Input.URL()
	if(util.IsApi(url)){
		util.NotAuthJson(c.Ctx)
		return
	}
	c.Data["content"] = "not auth"
	c.TplName = "error/404.tpl"
}

func (c *ErrorController) Error404() {
	url := c.Ctx.Input.URL()
	if(util.IsApi(url)){
		util.NotAuthJson(c.Ctx)
		return
	}
	c.Data["content"] = "page not found"
	c.TplName = "error/404.tpl"
}

func (c *ErrorController) Error500() {
	url := c.Ctx.Input.URL()
	if(util.IsApi(url)){
		util.NotAuthJson(c.Ctx)
		return
	}
	c.Data["content"] = "internal server error"
	c.TplName = "error/500.tpl"
}

func (c *ErrorController) ErrorDb() {
	url := c.Ctx.Input.URL()
	if(util.IsApi(url)){
		util.NotAuthJson(c.Ctx)
		return
	}
	c.Data["content"] = "database is now down"
	c.TplName = "error/dberror.tpl"
}