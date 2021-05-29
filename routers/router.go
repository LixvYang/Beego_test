package routers

import (
	"beego/controllers/fronted"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &frontend.IndexController{})
}


