package controllers

import (
	"rompapi/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

// Operations about Users
type ArticleController struct {
	beego.Controller
}

// URLMapping ...
func (c *ArticleController) URLMapping() {
	// c.Mapping("Post", c.Post)
}

// @router /add [post]
func (c *ArticleController) AddArticle() {
	var v models.Article
	var token models.Token			
	// 用户(用token查)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &token); err == nil {
		if err, user := models.GetUserByToken(token.Token); err == nil {
			v.User = user
		}
	}else{
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = Response{403, 403, fmt.Sprintf("%s", err), nil}
		c.ServeJSON()
		return
	}
	
	// 分类
	if err, category := models.GetCategoryById(1); err == nil {
		if category !=nil {
			v.Category = category
		}else{
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "未查找到分类数据", nil}
			c.ServeJSON()
			return
		}
	}else{
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = Response{403, 403, fmt.Sprintf("%s", err), nil}
		c.ServeJSON()
		return
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if titleLen := len(v.Title); titleLen<1{
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,"标题不能为空", nil}
			c.ServeJSON()
			return
		}
		if submitLen := len(v.Submit); submitLen<1{
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,"副标题不能为空", nil}
			c.ServeJSON()
			return
		}
		if article, err := models.AddArticle(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = Response{200, 200,"添加文章成功", article}
		} else {
			c.Data["json"] = &Response{1, 1, "添加文章失败", err.Error()}
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = Response{403, 403, "参数错误", nil}
	c.ServeJSON()
	return
}