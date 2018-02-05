package conf

import (
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego/session"
	"github.com/astaxie/beego"
)

func SetupSession(){
	sessionconf := &session.ManagerConfig{
		CookieName: "begoosessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("mysql", sessionconf)
	go beego.GlobalSessions.GC()
}
