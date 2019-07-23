package main

import (
	_ "rompapi/routers"
	"rompapi/models"
	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
