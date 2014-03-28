package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wangbinxiang/codesave/models"
)

func MysqlRegisterDB() {

	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlHost := beego.AppConfig.String("mysqlhost")
	mysqlPort := beego.AppConfig.String("mysqlport")
	mysqlDB := beego.AppConfig.String("mysqldb")

	orm.RegisterDataBase("codesave", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDB))
	models.Orm = orm.NewOrm()
	models.Orm.Using("codesave")
}
