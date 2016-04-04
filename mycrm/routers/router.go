package routers

import (
	"mycrm/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.CategoryController{})
	beego.AutoRouter(&controllers.MenuController{})
	beego.AutoRouter(&controllers.UserController{})
}

