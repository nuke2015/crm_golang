package controllers

import(
	"gopkg.in/mgo.v2/bson"
	."crm/models"
	"time"
)

type MenuController struct {
	BaseController
}

type menu struct{
	Id      bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Title   string
	Url     string
	Pid     string
	Ctime   time.Time
}

//分类列表
func (this *MenuController) Index() {
	var m Mongobase
	total,_:=m.Use("menu").Count(nil)
	p,err:=this.GetInt("page");
	if err!=nil || p<1{
		p=1
	}
	limit:=5
	skip:=(p-1)*limit
	result:=m.Use("menu").Query(nil,skip,limit)
	pager:=NewPager(p,total,limit,"/menu/index",true).ToString()
	this.Data["Result"]=result
	this.Data["Pager"]=pager
	this.display("menu_index")
}

//添加分类
func (this *MenuController) Add(){
	this.display("menu_add")
}

//添加入库
func (this *MenuController) Insert(){
	var m Mongobase
	a:=&menu{}
	a.Title=this.GetString("title")
	a.Url=this.GetString("url")
	a.Pid = this.GetString("pid")
	a.Ctime=time.Now()
	m.Use("menu").Insert(a)
	this.Redirect("/menu/index",302)
}

//编辑入库
func (this *MenuController) Edit(){
	var m Mongobase
	//http://127.0.0.1:8080/menu/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where:=bson.M{"_id":mid}
    result:=m.Use("menu").One(where)
	// this.dump(where)
	this.Data["result"]=result
	this.display("menu_edit")
}

//添加入库
func (this *MenuController) Update(){
	var m Mongobase
	//http://127.0.0.1:8080/menu/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where:=bson.M{"_id":mid}
	a:=bson.M{"title":this.GetString("title"),"content":this.GetString("content"),"url":this.GetString("url"),"pid":this.GetString("pid")}
	update:=bson.M{"$set":a}
	m.Use("menu").Update(where,update)
	this.Redirect("/menu/index",302)
}

//删除
func (this *MenuController) Delete(){
	var m Mongobase
	//http://127.0.0.1:8080/menu/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where:=bson.M{"_id":mid}
	m.Use("menu").Remove(where)
	this.Redirect("/menu/index",302)
}
