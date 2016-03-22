package controllers

import (
	"boa/models"
	"fmt"
	"github.com/astaxie/beego"

	"html/template"
	"time"

)


//注册
type RegisterController struct {
	BaseController
}


func (this *RegisterController) Get() {
		this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
		this.TplName = "注册/注册.html"
	}


func (this *RegisterController) Post() {

		username := this.GetString("username")
		password1 := this.GetString("password1")
		password2 := this.GetString("password2")
		//判断用户名是否已存在
		if models.FindUserByName(username)==true {								//永远为false？？

			flash := beego.NewFlash()
			flash.Warning("用户已存在，请换一个用户名")
			flash.Store(&this.Controller)


			this.Redirect("/register",302)
			}else{
				if password1 == password2{

			    now := time.Now()
				fmt.Println(now.Format("2006-01-02 15:04:05"))

				var user models.User

				user.UserName = username
				user.Password = password1
				user.LastLoginTime = now

				models.AddUser(user)


				flash := beego.NewFlash()
				flash.Notice("恭喜您,您已成功注册！")
				flash.Store(&this.Controller)

				this.SetSession("loginname", username)
				this.Redirect("/",302)
				}else{
					flash := beego.NewFlash()
					flash.Error("密码不一致，请核对输入")
					flash.Store(&this.Controller)


					this.Redirect("/register",302)
					}

			}



	}


//登陆
type LoginController struct {
	BaseController

}


func (this *LoginController) Get() {
	this.TplName = "登陆/登陆.html"
}

func (this *LoginController) Post() {

	username := this.GetString("username")
	password := this.GetString("password")

	admin_username:=beego.AppConfig.String("admin_username")
	admin_password:=beego.AppConfig.String("admin_password")
	//是否管理员登陆
	if (username==admin_username) && (password==admin_password){
		this.SetSession("loginname", username)



		flash := beego.NewFlash()
		flash.Notice("向您致敬，伟大的管理员同志！")
		flash.Store(&this.Controller)

		fmt.Println("hello,admin")
		this.IsLogin = true
		this.IsAdmin = true
		this.Redirect("/",302)

		}else{
			//不是管理员，则判断用户是否存在
			if models.FindUserByName(username){

				//判断密码是否正确
				var user models.User
				user= models.GetUserByName(username)
				fmt.Println(user)
				if user.Password==password{

					flash := beego.NewFlash()
					flash.Notice("欢迎回来！")
					flash.Store(&this.Controller)



					this.SetSession("loginname", username)
					this.IsLogin=true
					this.Redirect("/",302)
				}else{
					flash := beego.NewFlash()
					flash.Error("密码不对")
					flash.Store(&this.Controller)
					this.Redirect("/login",302)
					}
			}else{
				flash := beego.NewFlash()
					flash.Error("用户不存在，请先注册")
					flash.Store(&this.Controller)
				this.Redirect("/register",302)
				}


		}
	}

//注销
type LogoutController struct {
	BaseController
}

//退出登录
func (this *LogoutController) Get() {
	this.SetSession("loginname", "")
	this.Ctx.SetCookie("loginname", "")

	flash := beego.NewFlash()
	flash.Warning("您已退出！")
	flash.Store(&this.Controller)


	this.Redirect("/",302)
}





