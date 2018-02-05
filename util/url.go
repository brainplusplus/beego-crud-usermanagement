package util

import (
	"ScaleAffWeb/responses"
	"encoding/json"
	"github.com/astaxie/beego/context"
	"strings"
	"github.com/astaxie/beego"
)

func IsApi(url string) bool {
	return strings.HasPrefix(url, "/api/")
}

func IsApiResponseList(url string) bool {
	return IsApi(url) && strings.Contains(url, "/list")
}

func IsUrlContains(url string,routes *[]string) bool {
	for _, route := range *routes {
		if strings.Contains(url,"/"+route) {
			return true
		}
	}
	return false
}

func NotAuthJson(ctx *context.Context){
	url := ctx.Input.URL()
	if(IsApi(url)){
		if(IsApiResponseList(url)){
			NotAuthResponseListJson(ctx)
		}else{
			NotAuthResponseJson(ctx)
		}
	}
}

func NotAuthView(ctx *context.Context){
	ctx.Redirect(302, "/error/not_auth")
}

func NotLoginView(ctx *context.Context){
	ctx.Redirect(302, "/auth/login")
}

func RedirectSuccessRoute(ctx *context.Context){
	ctx.Redirect(302, beego.AppConfig.String("success_route"))
}

func NotAuthResponseJson(ctx *context.Context){
	ctx.Output.Header("Content-Type", "application/json")
	resp := responses.Response{}
	resp.Success = false
	resp.Message = "Failed Authenticated! check your session"
	json, _ := json.Marshal(resp)
	ctx.WriteString(string(json))
}

func NotAuthResponseListJson(ctx *context.Context){
	ctx.Output.Header("Content-Type", "application/json")
	resp := responses.ResponseList{}
	resp.Rows = []int{}
	resp.RecordsFiltered = 0
	resp.RecordsTotal = 0
	resp.Success = false
	resp.Message = "Failed Authenticated! check your session"
	json, _ := json.Marshal(resp)
	ctx.WriteString(string(json))
}