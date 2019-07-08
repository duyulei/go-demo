package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_DRIVER  = "mysql"
)
func RegisterDB() {
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	db := beego.AppConfig.String("mysqldb")
	DB_URL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pwd, url, db)
	orm.RegisterDriver(DB_DRIVER, orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", DB_URL)
	// 4. 自动创建表 参数二为是否开启创建表   参数三是否更新表
	// orm.RunSyncdb("default", true, true)
}