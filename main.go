package main

import (
	_ "rompapi/routers"
	_ "github.com/astaxie/beego/session/mysql"
	"rompapi/models"
	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.Session.SessionOn = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
