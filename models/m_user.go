package models

import (
	//~ "os"
	"crypto/md5"
	"fmt"

	"time"
	//~ "github.com/astaxie/beego"
	//~ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id            bson.ObjectId `bson:"_id"`
	UserId        int64
	UserName      string `bson:"username"`          //用户名
	RealName      string `bson:"realname"`          //真实姓名
	NickName      string `bson:"nickname"`          //昵称
	Password      string `bson:"password"`          //密码
	PasswordHash  string `bson:"pwdhash,omitempty"` //密码哈希值
	Email         string `bson:"email"`             //邮箱
	QQ            string //QQ
	WeiXin        string
											//年龄，工号，员工编号，头像，籍贯，住址，手机号，身份证，社保，工资卡号，
	Active        bool      //是否激活
	LastLoginTime time.Time //最后一次登陆时间
	LastIp        string    //最后一次登陆IP
	LoginCount    int       //登陆总次数
	Status        string    //状态
}

//#####################################################

//新增用户
func AddUser(user User) {
	fmt.Println("添加新的注册用户", "\n", user)
	h := md5.New()
	h.Write([]byte(user.Password))
	user.PasswordHash = fmt.Sprintf("%x", h.Sum(nil))
	NewId := bson.NewObjectId()
	err := C_user.Insert(bson.M{"_id": NewId, "username": user.UserName, "password": user.Password, "passwordhash": user.PasswordHash})
	if err != nil {
		panic(err)
	}
}

//通过密码查找用户
func FindUser(user *User) *User {
	fmt.Println(user)
	h := md5.New()
	h.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", h.Sum(nil))

	u := &User{}
	err := C_user.Find(bson.M{"email": user.Email, "password": user.Password}).One(u)
	if err != nil {
		panic(err)
	}

	return u
}

//更新用户
func UpdateUser(user User) {
	//err := C_user.Update(bson.M{"username": user.username}, bson.M{"$set": bson.M{"lastlogintime": now}})

	err := C_user.Update(bson.M{"username": user.UserName},
		bson.M{"$set": bson.M{"lastlogintime": time.Now(),
			"Email":      user.Email,
			"RealName":   user.RealName,
			"NickName":   user.NickName,
			"QQ":         user.QQ,
			"Weixin":     user.WeiXin,
			"Active":     user.Active,
			"LoginCount": user.LoginCount,
			"LastIp":     user.LastIp,
			"status":     "已更新个人信息"}})
	if err != nil {
		panic(err)
	}

}

//删除用户
func RemoveUser(user User) {
	err := C_user.Remove(bson.M{"username": user.UserName})
	if err != nil {
		panic(err)
	}
}

//通过用户名查找用户是否存在
func FindUserByName(name string) bool {
	var u User
	err := C_user.Find(bson.M{"username": name}).One(&u)
	if err != nil {
		return false
	}
	return true
}

//通过用户名查找用户
func GetUserByName(name string) User {
	var u User
	err := C_user.Find(bson.M{"username": name}).One(&u)
	if err != nil {
		panic(err)
	}

	return u
}

//通过邮箱名查找用户
func GetUserByEmail(email string) User {
	var u User
	err := C_user.Find(bson.M{"email": email}).One(&u)
	if err != nil {
		panic(err)
	}

	return u
}
