package models

import (
	//~ "os"
	//~ "crypto/md5"
	"fmt"

	"time"
	//~ "github.com/astaxie/beego"
	//~ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Article struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	Title        string        `json:"title"`
	Author       string        `json:"author"`
	Content      string        `json:"content"`
	Topic        string        `json:"topic"`
	Status       string        `json:"status"`
	CreatedTime  time.Time
	ModifiedTime time.Time
	kickup       int8 //点赞数目
	Kickdown     int8 //踩
	Score        int8 //得分
	HaveRead     bool //已读
	Favorite     bool //收藏

}

//所有文章
func AllArticles() (articles []Article) {
	err := C_article.Find(nil).All(&articles)
	if err != nil {
		panic(err)
	}

	return articles
}

//主题--文章列表
func ArticlesByTopic(topic string) (articles []Article) {
	err := C_article.Find(bson.M{"topic": topic}).All(&articles)
	if err != nil {
		panic(err)
	}

	return articles
}

//作者--文章列表
func ArticlesByAuthor(author string) (articles []Article) {
	err := C_article.Find(bson.M{"author": author}).All(&articles)
	if err != nil {
		panic(err)
	}

	return articles
}

//状态
func ArticlesByStatus(status string) (articles []Article) {

	err := C_article.Find(bson.M{"status": status}).All(&articles)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hex: ", articles[0].Id.Hex())

	return
}

//新增文章
func AddArticle(article Article) {
	NewId := bson.NewObjectId()
	err := C_article.Insert(bson.M{"_id": NewId, "title": article.Title, "author": article.Author, "topic": article.Topic, "content": article.Content})
	if err != nil {
		panic(err)
	}
}

//更新文章
func UpdateArticle(article Article) bool {
	err := C_article.UpdateId(article.Id, article)
	if err != nil {
		panic(err)
		return false
	}

	return true
}

//删除文章
func DeleteArticle(idStr string) bool {
	id := bson.ObjectIdHex(idStr)
	err := C_article.RemoveId(id)
	if err != nil {
		panic(err)
		return false
	}

	return true
}

//关键字查找
func ArticleByKeyword(key string) (articles []Article) {
	err := C_article.Find(bson.M{"keywords": key}).All(&articles)
	if err != nil {
		panic(err)
	}

	return articles
}

//ID
func ArticleById(id string) (article Article) {

	err := C_article.FindId(bson.ObjectIdHex(id)).One(&article)
	//~ err :=C_article.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&article)   //为什么必须是&article
	if err != nil {
		panic(err)
	}

	return
}

//标题
func ArticleByTitle(title string) (article *Article) {
	err := C_article.Find(bson.M{"title": title}).One(&article)
	if err != nil {
		panic(err)
	}

	return
}
