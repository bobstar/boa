package controllers

import (
	"fmt"
	"net/smtp"
	"strings"

	//~ "github.com/astaxie/beego"
)

func (this *MainController) Inbox() {

	this.TplName = "邮箱/收件箱/收件箱.html"
}

func (this *MainController) Outbox() {

	this.TplName = "邮箱/发件箱/发件箱.html"
}

func (this *MainController) Drafts() {

	this.TplName = "邮箱/草稿箱/草稿箱.html"
}

func (this *MainController) NewEmail() {

	this.TplName = "邮箱/新邮件/新邮件.html"
}

func (this *MainController) Email() {

	this.TplName = "邮箱/收件箱/邮件.html"
}

func (this *MainController) EmailSet() {

	this.TplName = "邮箱/设置.html"
}




const (
	HOST        = "smtp.163.com"
	SERVER_ADDR = "smtp.163.com:25"   //邮件服务器及端口
	//pop imap
	USER        = "quanwenbo@163.com" //发送邮件的邮箱
	PASSWORD    = "quanwenbo128"      //发送邮件邮箱的密码
)

type Email struct {
	user     string "user"
	password string "password"
	smtphost string "smtphost"
	to       string "to"       //接收地址
	subject  string "subject"  //邮件主题
	msg      string "msg"      //邮件内容
	mailtype string "mailtype" //邮件类型
}

func SendMail(email Email) error {

	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	var contentType string
	if email.mailtype == "html" {
		contentType = "Content-Type: text/" + email.mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + email.to + "\r\nFrom: " + USER + "<" + USER + ">\r\nSubject: " + email.subject + "\r\n" + contentType + "\r\n\r\n" + email.msg)
	send_to := strings.Split(email.to, ";")

	err := smtp.SendMail(SERVER_ADDR, auth, USER, send_to, msg)

	return err
}

func main() {
	user := "quanwenbo@163.com"
	password := "quanwenbo128"
	host := "smtp.163.com:25"
	to := "bobstar@163.com;quanwenbo@163.com;huanghaixiang@honghe-tech.com;quanwenbo@honghe-tech.com"

	subject := "Test send email by golang"

	msg := `
    <html>
    <body>
    <h3>
    "Test send email by golang 测试。。。。。。。。。"
    </h3>
    </body>
    </html>
    `

	email := Email{user, password, host, to, subject, msg, "html"}

	fmt.Println("start send email")

	err := SendMail(email)
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
	} else {
		fmt.Println("send mail success!")
	}
}
