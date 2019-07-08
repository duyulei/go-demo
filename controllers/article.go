package controllers

import (
	"rompapi/models"
	// "rompapi/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	// "time"
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
	// user_id := "a48899f3-90e7-4c71-b9ad-bd71b77d2a2a"
	user_id := c.GetSession("user_id").(string)
	fmt.Println("ddd")
	fmt.Println(user_id)
	if err, user := models.GetUserById(user_id); err == nil {
		v.User = user
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