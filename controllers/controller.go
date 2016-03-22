package controllers

import (
//~ "github.com/astaxie/beego"
//~ "github.com/beego/i18n"
)

//#################################################################  导航

type MainController struct {
	BaseController
}

func (this *MainController) HomeLogin() {

	this.TplName = "导航/home.html"
}

func (this *MainController) Join() {

	this.TplName = "加入/加入.html"
}


func (this *MainController) Tech() {

	this.TplName = "教程/教程列表.html"
}

func (this *MainController) PlayVideo() {
    path:= this.Ctx.Input.Param(":path")
    name := this.Ctx.Input.Param(":name")
    
    this.Data["Video"] = path+"/"+name
	this.TplName = "教程/video.html"
}


func (this *MainController) Contact() {

	this.TplName = "通讯录/通讯录.html"
}

func (this *MainController) Show() {
	this.TplName = "3D/3D.html"
}
func (this *MainController) Home() {

	this.TplName = "首页/首页.html"
}

func (this *MainController) Register() {

	this.TplName = "注册/注册.html"
}

func (this *MainController) Login() {

	this.TplName = "登陆/登陆.html"
}
//##################################################################    首页顶部导航
func (this *MainController) Task() {

	this.TplName = "待办任务/个人任务.html"
}

func (this *MainController) TeamTask() {

	this.TplName = "待办任务/团队任务.html"
}

func (this *MainController) HistoryTask() {

	this.TplName = "待办任务/历史任务.html"
}

func (this *MainController) PublicFiles() {

	this.TplName = "公共文件/公共文件.html"
}

//##################################################################    菜单项
func (this *MainController) Meeting() {

	this.TplName = "会议预约/会议预约.html"
}

func (this *MainController) OfficeSupplies() {

	this.TplName = "办公用品/办公用品.html"
}

func (this *MainController) FixedAssets() {

	this.TplName = "固定资产/固定资产.html"
}

func (this *MainController) Report() {

	this.TplName = "报表/报表.html"
}

//###################################################################### 测试

func (this *MainController) Test() {

	this.TplName = "视频背景.html"
}

func (this *MainController) Together() {

	this.TplName = "测试/together.html"
}

func (this *MainController) VideoList() {

	this.TplName = "视频/视频列表.html"
}

func (this *MainController) Video() {

	this.Data["Videoname"] = this.Ctx.Input.Param(":name")
	this.TplName = "视频/视频.html"
}
