package controllers

import(
	"gopkg.in/mgo.v2/bson"
	."crm/models"
	"time"
)

type CategoryController struct {
	AuthController
}

type category struct{
	Id      bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Title   string
	Ctime   time.Time
}

//分类列表
func (this *CategoryController) Index() {
	var m Mongobase
	total,_:=m.Use("category").Count(nil)
	p,err:=this.GetInt("page");
	if err!=nil || p<1{
		p=1
	}
	limit:=5
	skip:=(p-1)*limit
	result:=m.Use("category").Query(nil,skip,limit)
	pager:=NewPager(p,total,limit,"/category/index",true).ToString()
	this.Data["Result"]=result
	this.Data["Pager"]=pager
	this.display("category_index")
}

//添加分类
func (this *CategoryController) Add(){
	this.display("category_add")
}

//添加入库
func (this *CategoryController) Insert(){
	a:=&category{}
	a.Title=this.GetString("title")
	a.Ctime=time.Now()
	var m Mongobase
	m.Use("category").Insert(a)
	this.Redirect("/category/index",302)
}

//编辑入库
func (this *CategoryController) Edit(){
	var m Mongobase
	//http://127.0.0.1:8080/category/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where:=bson.M{"_id":mid}
    result:=m.Use("category").One(where)
	// this.dump(where)
	this.Data["result"]=result
	this.display("category_edit")
}

//添加入库
func (this *CategoryController) Update(){
	var m Mongobase
	//http://127.0.0.1:8080/category/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where:=bson.M{"_id":mid}

	a:=bson.M{"title":this.GetString("title"),"content":this.GetString("content")}
	update:=bson.M{"$set":a}
	m.Use("category").Update(where,update)
	this.Redirect("/category/index",302)
}

//删除
func (this *CategoryController) Delete(){
	var m Mongobase
	//http://127.0.0.1:8080/category/edit/?id=54c3be428c40e46373e69efe
	mid := m.MongoId(this.GetString("id"))
	where:=bson.M{"_id":mid}
	m.Use("category").Remove(where)
	this.Redirect("/category/index",302)
}
