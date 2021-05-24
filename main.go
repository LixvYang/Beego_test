package main

import (
	
	"encoding/gob"
	"gitee.com/shirdonl/LeastMall/common"
	"gitee.com/shirdonl/LeastMall/models"
	_ "gitee.com/shirdonl/LeastMall/routers"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/astaxie/beego/session/redis"
	_ "beego/routers"
	"github.com/astaxie/beego"

)

func main() {
	//add methods to map,for frontend HTML code.
	beego.AddFuncMap("timestampToDate",common.timestampToDate)
	models.DB.LogMode(true)
	beego.AddFuncMap("formatImage",common.FormatImage)
	beego.AddFuncMap("mul",common.Mul)
	beego.AddFuncMap("formatAttribute",common.FormatAttribute)
	beego.AddFuncMap("setting",models.GetSettingByColumn)

	//后台配置允许跨域
	beego.InsertFilter("*",beego.BeforeRouter,cors.Allow(&cors.Options{
		AllowAllOrigins: []string{"127.0.0.1"},
		AllowMethods: []string{
			"GET",
			"POST",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Authorization",
			"Access-Contorl-Allow-Origin",
			"Access-Contorl-Allow-Headers",
			"Content-Type",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Contorl-Allow-Origin",
			"Access-Contorl-Allow-Headers",
			"Content-Type",
		},
		AllowCredentials: true, //允许cookie
	}))
	//注册模型
	gob.Register(models.Administrator{})
	//关闭数据库
	defer models.DB.Close()
	//配置 Redis 用于储存session
	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	//docker-compose 请设置为redisServiceHost
	//beego.BConfig.WebConfig.Session.SessionProviderConfig="redisServiceHost:6379"

	//本地启动，设置如下
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
	beego.Run()
}

