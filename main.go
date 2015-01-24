package main

import (
	"crm/controllers"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	beego.Router("/", &controllers.MainController{},"*:Index")
	beego.AutoRouter(&controllers.MainController{})
	beego.AddFuncMap("hex",hex)
	beego.Run()
}

//mongoid
func hex(mid bson.ObjectId) string {
	return mid.Hex();
}

