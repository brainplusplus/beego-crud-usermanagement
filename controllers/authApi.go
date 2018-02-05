package controllers

import (
	"github.com/astaxie/beego"
	"ScaleAffWeb/middlewares"
	"ScaleAffWeb/services"
	"ScaleAffWeb/responses"
	"ScaleAffWeb/requests"
	"ScaleAffWeb/models"
)

// UserController operations for User
type AuthApiController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthApiController) URLMapping() {

}

func (c *AuthApiController) Login() {
	req := requests.LoginRequest{}
	c.ParseForm(&req)
	log.Info(req.Username)
	log.Info(req.Password)
	resp := services.GetUserService().Login(req.Username,req.Password)
	c.Data["json"] = resp
	c.ServeJSON()
}

func(c *AuthApiController) Validate() {
	resp := new(responses.Response)
	tokenString := c.GetString("token")
	et := middlewares.EasyToken{}
	isValid, issuer, err := et.ValidateJWTToken(tokenString)
	resp.Success = isValid
	resp.Data = issuer
	resp.Message = err.Error()
	c.Data["json"] = resp
	c.ServeJSON()
}

func(c *AuthApiController) SaveSession() {
	resp := new(responses.Response)
	tokenString := c.GetString("token")
	et := middlewares.EasyToken{}
	isValid, issuer, err := et.ValidateJWTToken(tokenString)
	if isValid {
		respUser := services.GetUserService().GetByUsername(issuer)
		session := c.StartSession()
		session.Set("UserId",respUser.Data.(models.User).Id)
		session.Set("Username",issuer)
		session.Set("Token",tokenString)
		log.Info("SESSION STARTED")
		log.Info(issuer)
	}

	resp.Success = isValid
	resp.Data = issuer
	if err!= nil {
		resp.Message = err.Error()
	}
	c.Data["json"] = resp
	c.ServeJSON()
}