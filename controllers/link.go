package controllers

import (
	// "boa/models"
	// "fmt"

	//~ "github.com/astaxie/beego"
	//~
	//~ "html/template"
	//~ "time"
)

type LinksController struct {
	BaseController
}

func (this *LinksController) Get() {
	this.TplName = "链接/链接.html"
}
