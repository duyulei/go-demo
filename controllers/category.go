package controllers

import (
	"encoding/json"
	"fmt"
	"rompapi/models"

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
		if nameLen := len(v.Name); nameLen < 1 {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "名称不能为空", nil}

			c.ServeJSON()
			return
		}
		if exist := models.CheckCategoryName(v.Name); exist {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "分类名称存在", nil}

			c.ServeJSON()
			return
		}
		if category, err := models.AddCategory(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = Response{200, 200, "添加分类成功", category}
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

// @Title Getall
// @router /getall [post]
func (c *CategoryController) GetAllCategory() {
	var v models.Category
	var p models.Pagination
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &p); err == nil {
			if err, categories := models.GetAllCategory(&v, &p); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = Response{200, 200, "查询分类成功", categories}
				c.ServeJSON()
				return
			} else {
				c.Ctx.Output.SetStatus(403)
				c.Data["json"] = Response{200, 200, "查询分类失败", err}
				c.ServeJSON()
				return
			}
		}
	}
	return
}

// @Title Delete
// @router /delete [post]
func (c *CategoryController) Delete() {
	var v models.Category
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Id = 13 // 写死的ID
		ers := models.DeleteCategory(&v)
		if ers == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = Response{200, 200, "删除成功", nil}
			c.ServeJSON()
			return
		} else {
			c.Ctx.Output.SetStatus(403)
			c.Data["json"] = Response{200, 200, "删除失败", fmt.Sprintf("%s", ers)}
			c.ServeJSON()
			return
		}
	}
	c.Data["json"] = Response{403, 403, "参数错误", nil}
	c.ServeJSON()
	return
}

// @Title Updata
// @router /updata [post]
func (c *CategoryController) Updata() {
	var v models.Category
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if nameLen := len(v.Name); nameLen < 1 {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "名称不能为空", nil}

			c.ServeJSON()
			return
		}
		if exist := models.CheckCategoryName(v.Name); exist {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "分类名称存在", nil}

			c.ServeJSON()
			return
		}
		v.Id = 12 // 写死的ID
		if err := models.UpdateCategoryById(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = Response{200, 200, "修改分类成功", nil}
		} else {
			c.Data["json"] = &Response{1, 1, "修改分类失败", err.Error()}
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = Response{403, 403, "参数错误", nil}
	c.ServeJSON()
	return
}
