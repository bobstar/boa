package models

import (
	//~ "os"
	//~ "crypto/md5"
	"fmt"

	//~ "github.com/astaxie/beego"
	//~ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Topic struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
	Name          string        `json:"author"`
	ArticleNumber int
}

//查找主题是否存在
func FindTopic(name string) bool {
	var t Topic
	err := C_topic.Find(bson.M{"name": name}).One(&t)
	if err == nil {
		fmt.Println(err)
		return true
	} else {
		fmt.Println(err)
		return false
	}

}

//查找主题
func FindTopicByName(name string) (t *Topic) {
	err := C_topic.Find(bson.M{"name": name}).One(&t)
	if err != nil {
		panic(err)
	}
	return
}

//新增主题
func AddTopic(topic string) {
	NewId := bson.NewObjectId()
	err := C_topic.Insert(bson.M{"_id": NewId, "name": topic})
	if err != nil {
		panic(err)
	}
}

//更新主题
func UpdateTopic(topic *Topic) bool {
	err := C_topic.UpdateId(topic.Id, topic)
	if err != nil {
		panic(err)
		return false
	}

	return true
}

//所有主题
func AllTopics() (topics []Topic) {
	err := C_topic.Find(nil).Sort("$natural").All(&topics)
	if err != nil {
		panic(err)
	}

	return topics
}
