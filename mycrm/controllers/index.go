package controllers

import(
    "gopkg.in/mgo.v2/bson"
    ."mycrm/models"
    "time"
)

type IndexController struct {
    BaseController
}

type myuser struct{
	Id       bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Username string
	Password string
	Group    string
	Ctime    time.Time
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

// 注册
func (this *IndexController) Reg(){
    this.display("index_reg");
}

// 注册接收
func (this *IndexController) Doreg(){
	username:=this.GetString("username")
    password:=this.GetString("password")
    var m Mongobase
    where:=bson.M{"username":username}
    result:=m.Use("user").One(where)
    if result != nil {
        this.dump("用户已存在!");
    }else{
    	u:= &myuser{}
    	u.Username=username
    	u.Password=password
        m.Use("user").Insert(u);
        this.Redirect("/index/index",302)
    }    
}
