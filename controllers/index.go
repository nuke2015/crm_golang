package controllers

import(
	"gopkg.in/mgo.v2/bson"
	."crm/models"
)

type IndexController struct {
	BaseController
}

//登陆入口
func (this *IndexController) Index() {
	this.display("index_index")
}


//登陆执行
func (this *IndexController) Login(){
	username:=this.GetString("username")
	password:=this.GetString("password")
	var m Mongobase
	where:=bson.M{"username":username,"password":password}
    result:=m.Use("user").One(where)
	if result != nil {
		user:=result.(bson.M)
		uid:=user["_id"].(bson.ObjectId)
		this.SetSession("username",username)
		this.SetSession("userid",uid.Hex())
		this.Redirect("/main/index",302)
	}else{
		this.dump("用户名或密码错误!");
	}
}

//登出执行
func (this *IndexController) Logout(){
	this.SetSession("username","");
	this.SetSession("userid","");
	this.Redirect("/index/index",302)
}

