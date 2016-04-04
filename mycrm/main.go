package main

import (
	_ "mycrm/routers"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	beego.AddFuncMap("hex",hex)
	beego.Run()
}

//mongoid
func hex(mid bson.ObjectId) string {
	return mid.Hex();
}

