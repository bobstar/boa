package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"

	"fmt"
	"os"
)

var DB *mgo.Database
var C_user, C_article, C_topic, C_link *mgo.Collection

//连接数据库，返回数据库指针对象
func init() {
	conn := beego.AppConfig.String("db_host")
	if conn == "" {
		beego.Error("数据库地址还没有配置,请到conf内配置db主机地址.")
		os.Exit(1)
	}

	session, err := mgo.Dial(conn)
	if err != nil {
		beego.Error("MongoDB连接失败:", err.Error())
		os.Exit(1)
	}

	//~ defer session.Close()
	dbname := beego.AppConfig.String("appname")

	session.SetMode(mgo.Monotonic, true)
	fmt.Println("\n", "dbsession is ", "\n", session, "\n")
	DB = session.DB(dbname)

	fmt.Println("The db is ", "\n", DB, "\n")

	C_user = DB.C("user")
	C_article = DB.C("article")
	C_topic = DB.C("topic")
	C_link = DB.C("link")

	//此处需判断是否已存在admin
	//~ var admin User
	//~ admin.UserName = beego.AppConfig.String("admin_username")
	//~ admin.Password = beego.AppConfig.String("admin_password")
	//~ admin.Email = beego.AppConfig.String("admin_email")
	//~ AddUser(admin)

	fmt.Println("\n", "The collections are", "\n", DB.CollectionNames, C_user, C_article, C_topic, C_link, "\n")

}
