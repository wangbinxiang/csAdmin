package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin"
	"github.com/wangbinxiang/csAdmin/models"
	_ "github.com/wangbinxiang/csAdmin/routers"
)

func main() {
	admin.Run()
	models.MysqlRegisterDB()
	beego.Run()
}
