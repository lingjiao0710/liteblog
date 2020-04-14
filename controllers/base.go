package controllers

import (
	"log"

	"github.com/astaxie/beego"
)

type NestPreparer interface {
	NestPrepare()
}

type BaseController struct {
	beego.Controller
}

func (ctx *BaseController) Prepare() {
	log.Println("BaseController")

	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	//判断子类是否实现了NestPreparer接口，如果实现了就调用接口方法
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
