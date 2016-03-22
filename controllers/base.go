package controllers

import (
	"github.com/astaxie/beego"
	//~ "time"
	"boa/models"
	"fmt"
	"github.com/beego/i18n"

	//~ "gopkg.in/mgo.v2/bson"
)

//基础结构体，供其他结构体匿名嵌入，在beego基础上增加了其他“全局”变量和方法
type BaseController struct {
	beego.Controller
	i18n.Locale
	User    models.User
	IsLogin bool
	IsAdmin bool
}

func (this *BaseController) Prepare() {
	this.Data["Title"] = beego.AppConfig.String("title")
	this.Data["LoginName"] = this.GetSession("loginname")

	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["notice"]; ok {
		fmt.Println(n)
	} else if n, ok = flash.Data["warning"]; ok {
		fmt.Println(n)
	} else if n, ok = flash.Data["error"]; ok {
		fmt.Println(n)
	} else {
		fmt.Println("没有Flash信息")
	}

	//语言处理
	this.Lang = ""
	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}
	if len(this.Lang) == 0 {
		this.Lang = "zh-CN"
	}
	this.Data["Lang"] = this.Lang

	//登陆状态检测
	if this.Ctx.Request.Method == "POST" && this.Ctx.Request.URL.Path != "/注册" && this.Ctx.Request.URL.Path != "/" {
		if this.GetSession("loginname") == "" {
			this.Redirect("/", 302)
		} else {
			this.IsLogin = true
		}

	}
}
