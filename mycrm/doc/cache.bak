

c := this.NewCache()
tmp:=c.Get("hello")
this.dump(tmp)
if(tmp != nil){
	this.dump(tmp);
}else{
	c.Put("hello",time.Now(),2)
}



//封装前
package controllers

import(
	// "gopkg.in/mgo.v2/bson"
	// ."crm/models"
	"github.com/astaxie/beego/cache"
	"time"
)

type MainController struct {
	BaseController
}

type person struct{
	Title string
	Height string
	Width int64
}

func (this *MainController) Get() {
	// c,_ := cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}`)
	c,_ := cache.NewCache("memory", `{"interval":60}`)
	tmp:=c.Get("hello")
	this.dump(tmp)
	if(tmp != nil){
		this.dump(tmp);
	}else{
		c.Put("hello",time.Now(),2)
	}
}

