package controllers

import (
	"rompapi/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// Operations about Users
type CategoryController struct {
	beego.Controller
}

// URLMapping ...
func (c *CategoryController) URLMapping() {
}

// @router /add [post]
func (c *CategoryController) AddCategory() {
	var v models.Category
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if nameLen := len(v.Name); nameLen<1{
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,"名称不能为空", nil}

			c.ServeJSON()
			return
		}
		if exist := models.CheckCategoryName(v.Name); exist {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,"分类名称存在", nil}

			c.ServeJSON()
			return
		}
		if category, err := models.AddCategory(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = Response{200, 200,"添加分类成功", category}
		} else {
			c.Data["json"] = &Response{1, 1, "添加分类失败", err.Error()}
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = Response{403, 403, "参数错误", nil}
	c.ServeJSON()
	return
}