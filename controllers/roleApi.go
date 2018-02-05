package controllers

import (
	"github.com/astaxie/beego"
	"ScaleAffWeb/responses"
	"ScaleAffWeb/requests"
	"ScaleAffWeb/services"
)

// RoleController operations for Role
type RoleApiController struct {
	beego.Controller
}

// URLMapping ...
func (c *RoleApiController) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("Get", c.Get)
	c.Mapping("Save", c.Save)
	c.Mapping("Update", c.Update)
	c.Mapping("Delete", c.Delete)
}

func (c *RoleApiController) List() {
	responseList := new(responses.ResponseList)
	responseList.Success = false

	requestList := new(requests.RoleRequestList)
	requestList.Offset = 0
	requestList.Limit = 10
	requestList.Page  = 1
	requestList.Search = c.GetString("search")
	requestList.OrderBy = "id ASC"

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

	responseList = services.GetRoleService().GetAllByRequest(requestList)
	c.Data["json"] = responseList
	c.ServeJSON()
}

func (c *RoleApiController) Get() {
	resp := new(responses.Response)
	resp.Success = false

	id := c.GetString("id")

	resp = services.GetRoleService().GetById(id)
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *RoleApiController) Save() {
	resp := new(responses.Response)
	resp.Success = false

	req := requests.RoleRequest{}
	if err := c.ParseForm(&req); err == nil {
		log.Info(req)
		resp = services.GetRoleService().Save(&req)
	} else {
		resp.Message = err.Error()
		resp.Success = false
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *RoleApiController) Update() {
	resp := new(responses.Response)
	resp.Success = false
	req := requests.RoleRequest{}
	if err := c.ParseForm(&req); err == nil {
		log.Info(req)
		resp = services.GetRoleService().Update(&req)
	} else {
		resp.Message = err.Error()
		resp.Success = false
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *RoleApiController) Delete() {
	resp := new(responses.Response)
	resp.Success = false

	id := c.GetString("id")

	resp = services.GetRoleService().DeleteById(id)
	c.Data["json"] = resp
	c.ServeJSON()
}
