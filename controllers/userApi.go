package controllers

import (
	"github.com/astaxie/beego"
	"ScaleAffWeb/requests"
	"ScaleAffWeb/responses"
	"ScaleAffWeb/services"
)


// UserController operations for User
type UserApiController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserApiController) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Get", c.Get)
	c.Mapping("Save", c.Save)
	c.Mapping("Update", c.Update)
	c.Mapping("Delete", c.Delete)
}

func (c *UserApiController) List() {
	responseList := new(responses.ResponseList)
	responseList.Success = false

	requestList := new(requests.UserRequestList)
	requestList.Offset = 0
	requestList.Limit = 10
	requestList.Page  = 1
	requestList.Search = c.GetString("search")
	requestList.RoleId = c.GetString("role_id")
	requestList.OrderBy = "created_at DESC"

	// limit: 10 (default is 10)
	if v, err := c.GetInt64("rows"); err == nil {
		requestList.Limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("page"); err == nil {
		requestList.Page = v
	}
	requestList.Offset = (requestList.Page-1) * requestList.Limit

	// sortby: col1,col2
	if v := c.GetString("order_by"); v != "" {
		requestList.OrderBy = v
	}

	responseList = services.GetUserService().GetAllByRequest(requestList)
	c.Data["json"] = responseList
	c.ServeJSON()
}

func (c *UserApiController) Get() {
	resp := new(responses.Response)
	resp.Success = false

	id := c.GetString("id")

	resp = services.GetUserService().GetById(id)
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *UserApiController) Save() {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.UserRequest{}
	if err := c.ParseForm(&req); err == nil {
		c.Ctx.Input.Bind(&req.Roles, "Roles")
		log.Info(req)
		resp = services.GetUserService().Save(&req)
	} else {
		resp.Message = err.Error()
		resp.Success = false
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *UserApiController) Update() {
	resp := new(responses.Response)
	resp.Success = false
	req := requests.UserRequest{}
	if err := c.ParseForm(&req); err == nil {
		c.Ctx.Input.Bind(&req.Roles, "Roles")
		log.Info(req)
		resp = services.GetUserService().Update(&req)
	} else {
		resp.Message = err.Error()
		resp.Success = false
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *UserApiController) Delete() {
	resp := new(responses.Response)
	resp.Success = false

	id := c.GetString("id")

	resp = services.GetUserService().DeleteById(id)
	c.Data["json"] = resp
	c.ServeJSON()
}