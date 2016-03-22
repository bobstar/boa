package models

import (
	//~ "os"
	//~ "crypto/md5"
	"fmt"

	//~ "github.com/astaxie/beego"
	//~ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Link struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Url         string        `json:"url" bson:"url"`
	Description string        `json:"description" bson:"description"`
}

//所有链接
func AllLinks() (Links []Link) {
	err := C_link.Find(nil).Sort("name").All(&Links)
	if err != nil {
		panic(err)
	}

	return Links
}

//新增链接
func AddLink(link Link) {
	NewId := bson.NewObjectId()
	fmt.Println(link.Name, link.Url)
	err := C_link.Insert(bson.M{"_id": NewId, "name": link.Name, "url": link.Url, "description": link.Description})
	if err != nil {
		panic(err)
	}
}

//删除链接
func DeleteLink(id string) {
	err := C_link.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		panic(err)
	}

	fmt.Println(bson.IsObjectIdHex(id))
	fmt.Println(bson.ObjectIdHex(id))
}
