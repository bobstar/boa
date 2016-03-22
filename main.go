package main

import (

	_ "boa/routers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"gopkg.in/mgo.v2/bson"
	"strings"
)


//模板函数
func MongoIdToString(id bson.ObjectId)(string){
    idstr:=id.Hex()
    return idstr
}





//语言国际化
var langTypes []string // Languages that are supported.

func init() {

	langTypes = strings.Split(beego.AppConfig.String("langs"), "|")
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}




func main() {


	beego.AddFuncMap("idtostring",MongoIdToString)    //模板函数
	beego.AddFuncMap("i18n", i18n.Tr)

    beego.Run("0.0.0.0:8080")
}





