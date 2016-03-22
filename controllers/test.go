package controllers

import (

//~ "github.com/astaxie/beego"

)

//测试
type TestController struct {
	BaseController
}

func (this *TestController) Test() {

	this.TplName = "test/test.html"

}

func (this *TestController) Test1() {

	this.TplName = "test/test1.html"

}

func (this *TestController) Test2() {

	this.TplName = "test/test2.html"

}
func (this *TestController) Test3() {

	this.TplName = "test/test3.html"

}
