package filters

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"ScaleAffWeb/logging"
	"ScaleAffWeb/services"
	"ScaleAffWeb/models"
	"ScaleAffWeb/util"
	"ScaleAffWeb/objects"
)

var (
	log = logging.MustGetLogger()
 	administratorRoutes *[]string
	vendorRoutes *[]string
	affiliateRoutes *[]string
	anonymousRoutes *[]string
)

func InitAuthFilter(){
	var FilterUser = func(ctx *context.Context){
		url := ctx.Input.URL()
		usernameSession := ctx.Input.Session("Username")
		var username string = ""
		if usernameSession != nil {
			username = usernameSession.(string)
		}
		//log.Info(username)
		respRoles := services.GetUserRoleService().GetAllRoleByUsername(username)
		roles := respRoles.Data.([]models.Role)
		//log.Info(roles)
		if url == "/" {
			if username == "" {
				util.NotLoginView(ctx)
			}else{
				util.RedirectSuccessRoute(ctx)
			}
			return
		}else {
			if util.IsUrlContains(url,getAnonymousRoutes()) {
				return
			}else{
				if username == "" {
					if util.IsApi(url) {
						util.NotAuthJson(ctx)
					}else{
						util.NotLoginView(ctx)
					}
					return
				}

				if isAdministrator(&roles) && isAllowedUrlAdministrator(url) {
					return
				} else if isAffiliate(&roles) && isAllowedUrlAffiliate(url) {
					return
				} else if isVendor(&roles) && isAllowedUrlVendor(url) {
					return
				}else{
					if util.IsApi(url) {
						util.NotAuthJson(ctx)
					}else{
						util.NotAuthView(ctx)
					}
					return
				}
			}
		}
	//
	//	_, ok := ctx.Input.Session("uid").(int)
	//	if !ok {
	//		//util.NotAuthJson(ctx)
	////		ctx.Redirect(302, "/auth/login")
	//	}

	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
}

func isAdministrator(roles *[]models.Role) bool {
	return hasRole(roles,objects.ADMINISTRATOR)
}

func isVendor(roles *[]models.Role) bool {
	return hasRole(roles,objects.VENDOR)
}

func isAffiliate(roles *[]models.Role) bool {
	return hasRole(roles,objects.AFFILIATE)
}

func isAllowedUrl(url string,routes *[]string) bool {
	for _, route := range *routes {
		if strings.Contains(url,"/"+route) {
			return true
		}
	}
	return false
}

func isAllowedUrlAdministrator(url string) bool {
	return isAllowedUrl(url,administratorRoutes)
}

func isAllowedUrlVendor(url string) bool {
	return isAllowedUrl(url,vendorRoutes)
}

func isAllowedUrlAffiliate(url string) bool {
	return isAllowedUrl(url,affiliateRoutes)
}

func hasRole(roles *[]models.Role,roleUser string) bool {
	for _, role := range *roles {
		if roleUser == role.Id {
			return true
		}
	}
	return false
}

func getAdministratorRoutes()(*[]string){
	return administratorRoutes
}

func getVendorRoutes()(*[]string){
	return vendorRoutes
}

func getAffiliateRoutes()(*[]string){
	return affiliateRoutes
}

func getAnonymousRoutes()(*[]string){
	return anonymousRoutes
}

func InitPathRuleRoutes(){
	setAdministratorRoutes()
	setVendorRoutes()
	setAffiliateRoutes()
	setAnonymousRoutes()
}

func setAdministratorRoutes(){
	routes := strings.Split(beego.AppConfig.String("administrator_routes"),",")
	administratorRoutes = &routes
}

func setVendorRoutes(){
	routes := strings.Split(beego.AppConfig.String("vendor_routes"),",")
	vendorRoutes = &routes
}

func setAffiliateRoutes(){
	routes := strings.Split(beego.AppConfig.String("affiliate_routes"),",")
	affiliateRoutes = &routes
}

func setAnonymousRoutes(){
	routes := strings.Split(beego.AppConfig.String("anonymous_routes"),",")
	anonymousRoutes = &routes
}