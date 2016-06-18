package main

import (
	"mycrm/controllers"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	
	// index
	beego.Router("/", &controllers.IndexController{},"get:Index")
	
	// control
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.CategoryController{})
	beego.AutoRouter(&controllers.MenuController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AddFuncMap("hex",hex)
	beego.Run()
}

//mongoid
func hex(mid bson.ObjectId) string {
	return mid.Hex();
}

