package routers

import (
	"boa/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:HomeLogin")
    beego.Router("/加入", &controllers.MainController{}, "*:Join")
	beego.Router("/教程", &controllers.MainController{}, "*:Tech")
    beego.Router("/视频/:path/:name", &controllers.MainController{}, "*:PlayVideo")
    beego.Router("/链接", &controllers.LinksController{})

	beego.Router("/注册", &controllers.RegisterController{})
	beego.Router("/登陆", &controllers.LoginController{})
	beego.Router("/找回密码", &controllers.LoginController{})
	beego.Router("/注销", &controllers.LogoutController{})



	beego.Router("/首页", &controllers.MainController{}, "*:Home")


	beego.Router("/收件箱", &controllers.MainController{}, "*:Inbox")
	beego.Router("/发件箱", &controllers.MainController{}, "*:Outbox")
	beego.Router("/草稿箱", &controllers.MainController{}, "*:Drafts")
	beego.Router("/设置", &controllers.MainController{}, "*:EmailSet")
	beego.Router("/写邮件", &controllers.MainController{}, "*:NewEmail")
	beego.Router("/邮件", &controllers.MainController{}, "*:Email")

	beego.Router("/个人任务", &controllers.MainController{}, "*:Task")
	beego.Router("/团队任务", &controllers.MainController{}, "*:TeamTask")
	beego.Router("/历史任务", &controllers.MainController{}, "*:HistoryTask")

	beego.Router("/展示", &controllers.MainController{}, "*:Show")

	beego.Router("/公共文件", &controllers.MainController{}, "*:PublicFiles")

	beego.Router("/聊天室", &controllers.MainController{}, "*:IM")
	beego.Router("/聊天室/ws", &controllers.WSController{})

	beego.Router("/报表", &controllers.MainController{}, "*:Report")

	beego.Router("/通讯录", &controllers.MainController{}, "*:Contact")


	beego.Router("/办公用品", &controllers.MainController{}, "*:OfficeSupplies")
	beego.Router("/固定资产", &controllers.MainController{}, "*:FixedAssets")
	beego.Router("/会议预约", &controllers.MainController{}, "*:Meeting")

	
}
