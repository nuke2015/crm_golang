package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) dump(params interface{}) {
	fmt.Fprint(this.Ctx.ResponseWriter, params);
}
