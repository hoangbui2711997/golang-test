package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"log"
	"my-first-beego-project/controllers"
	"time"
)

func init() {
	// use to register a router address
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/test-tpl", &controllers.BlogsController{})
	beego.Router("/employees", &controllers.FirstController{}, "get:GetEmployees")
	beego.Router("/home", &controllers.SessionController{}, "get:Home")
	beego.Router("/login", &controllers.SessionController{}, "get:Login")
	beego.Router("/logout", &controllers.SessionController{}, "get:Logout")
	beego.Router("/cache", &controllers.CacheController{}, "get:GetFromCache")
	beego.Router("/test-params/:key", &controllers.CacheController{}, "get:GetResult")
	//beego.InsertFilter("/*", beego.BeforeRouter, filters.LogManager)

	beego.Get("/test", func(ctx *context.Context){
		// print hello world
		ctx.Output.Body([]byte("hello world"))
	})

	beego.Post("/alice",func(ctx *context.Context){
		ctx.Output.Body([]byte("bob"))
	})

	beego.Any("/foo",func(ctx *context.Context){
		ctx.Output.Body([]byte("bar"))
	})

	// life cycle of request
	// first
	beego.InsertFilter("/cache", beego.BeforeStatic, func(ctx *context.Context) {
		log.Println("BeforeStatic:::::IP :: " + ctx.Request.RemoteAddr + ", Time:: " + time.Now().Format(time.RFC850))
	})
	// second
	beego.InsertFilter("/cache", beego.BeforeRouter, func(ctx *context.Context) {
		log.Println("BeforeRouter:::::IP :: " + ctx.Request.RemoteAddr + ", Time:: " + time.Now().Format(time.RFC850))
	})
	// third
	beego.InsertFilter("/cache", beego.BeforeExec, func(ctx *context.Context) {
		log.Println("BeforeExec:::::IP :: " + ctx.Request.RemoteAddr + ", Time:: " + time.Now().Format(time.RFC850))
	})
	// fourth
	beego.InsertFilter("/cache", beego.AfterExec, func(ctx *context.Context) {
		log.Println("AfterExec:::::IP :: " + ctx.Request.RemoteAddr + ", Time:: " + time.Now().Format(time.RFC850))
	})
	// five
	beego.InsertFilter("/cache", beego.FinishRouter, func(ctx *context.Context) {
		log.Println("FinishRouter:::::IP :: " + ctx.Request.RemoteAddr + ", Time:: " + time.Now().Format(time.RFC850))
	})
}
