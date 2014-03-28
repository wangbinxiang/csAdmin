package routers

import (
	"github.com/astaxie/beego"
	"github.com/wangbinxiang/csAdmin/controllers"
)

func init() {
	beego.Router("/cs/tag/index", &controllers.TagController{}, "*:Index")
	beego.Router("/cs/tag/add", &controllers.TagController{}, "*:Add")
}
