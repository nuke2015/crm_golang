package controllers

import(
	"gopkg.in/mgo.v2/bson"
	."crm/models"
	"time"
)

type MainController struct {
	BaseController
}

type article struct{
	Id      bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Title   string
	Content string
	Ctime   time.Time
}

//文章列表
func (this *MainController) Index() {
	var m Mongobase
	total,_:=m.Use("article").Count(nil)
	p,err:=this.GetInt("page");
	if err!=nil || p<1{
		p=1
	}
	limit:=5
	skip:=(p-1)*limit
	result:=m.Use("article").Query(nil,skip,limit)
	pager:=NewPager(p,total,limit,"/main/index",true).ToString()
	this.Data["Result"]=result
	this.Data["Pager"]=pager
	this.display("index")
}

//添加文章
func (this *MainController) Add(){
	this.display("add")
}

//添加入库
func (this *MainController) Insert(){
	a:=&article{}
	a.Title=this.GetString("title")
	a.Content=this.GetString("content")
	a.Ctime=time.Now()
	var m Mongobase
	m.Use("article").Insert(a)
	this.Redirect("/",302)
}

//编辑入库
func (this *MainController) Edit(){
	var m Mongobase
	//http://127.0.0.1:8080/main/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where:=bson.M{"_id":mid}
    result:=m.Use("article").One(where)
	// this.dump(where)
	this.Data["result"]=result
	this.display("edit")
}

//添加入库
func (this *MainController) Update(){
	var m Mongobase
	//http://127.0.0.1:8080/main/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where:=bson.M{"_id":mid}

	a:=bson.M{"title":this.GetString("title"),"content":this.GetString("content")}
	update:=bson.M{"$set":a}
	m.Use("article").Update(where,update)
	this.Redirect("/",302)
}

//删除
func (this *MainController) Delete(){
	var m Mongobase
	//http://127.0.0.1:8080/main/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where:=bson.M{"_id":mid}
	m.Use("article").Remove(where)
	this.Redirect("/",302)
}
