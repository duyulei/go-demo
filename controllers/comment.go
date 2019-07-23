package controllers

import (
	"encoding/json"
	"fmt"
	"rompapi/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type CommentController struct {
	beego.Controller
}

// URLMapping ...
func (c *CommentController) URLMapping() {
	// c.Mapping("Post", c.Post)
}

// @router /add [post]
func (c *CommentController) AddComment() {
	var v models.Comment
	var token models.Token
	// 用户(用token查)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &token); err == nil {
		if err, user := models.GetUserByToken(token.Token); err == nil {
			v.User = user
		}
	} else {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = Response{403, 403, fmt.Sprintf("%s", err), nil}
		c.ServeJSON()
		return
	}

	// 文章
	if err, article := models.GetArticleById(1); err == nil {
		if article != nil {
			v.Article = article
		} else {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "未查找到文章数据", nil}
			c.ServeJSON()
			return
		}
	} else {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = Response{403, 403, fmt.Sprintf("%s", err), nil}
		c.ServeJSON()
		return
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if contentLen := len(v.Content); contentLen < 1 {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "评论内容不能为空", nil}
			c.ServeJSON()
			return
		}
		if Comment, err := models.AddComment(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = Response{200, 200, "添加评论成功", Comment}
		} else {
			c.Data["json"] = &Response{1, 1, "添加评论失败", err.Error()}
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = Response{403, 403, "参数错误", nil}
	c.ServeJSON()
	return
}

// @Title Getart
// @router /artcomment [post]
func (c *CommentController) GetArticleComment() {
	var v models.Comment
	var p models.Pagination
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &p); err == nil {
			// 写死的ID
			if err, article := models.GetArticleById(1); err == nil {
				v.Article = article
			} else {
				c.Ctx.Output.SetStatus(403)
				c.Data["json"] = Response{200, 200, "获取文章失败", err}
				c.ServeJSON()
				return
			}
			if err, comments := models.GetArticleComment(&v, &p); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = Response{200, 200, "查询评论成功", comments}
				c.ServeJSON()
				return
			} else {
				c.Ctx.Output.SetStatus(403)
				c.Data["json"] = Response{200, 200, "查询评论失败", err}
				c.ServeJSON()
				return
			}
		}
	}
	return
}

// @Title Getuser
// @router /usercomment [post]
func (c *CommentController) GetUSerComment() {
	var v models.Comment
	var p models.Pagination
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, &p); err == nil {
			// 写死的ID
			user_id := "83ccdb38-f00f-4756-b66a-b1d99f2f1dc5"
			if err, user := models.GetUserById(user_id); err == nil {
				v.User = user
			} else {
				c.Ctx.Output.SetStatus(403)
				c.Data["json"] = Response{200, 200, "获取用户失败", err}
				c.ServeJSON()
				return
			}
			if err, comments := models.GetUserComment(&v, &p); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = Response{200, 200, "查询评论成功", comments}
				c.ServeJSON()
				return
			} else {
				c.Ctx.Output.SetStatus(403)
				c.Data["json"] = Response{200, 200, "查询评论失败", err}
				c.ServeJSON()
				return
			}
		}
	}
	return
}

// @Title Delete
// @router /delete [post]
func (c *CommentController) Delete() {
	var v models.Comment
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Id = 7 // 写死的ID
		ers := models.DeleteComment(&v)
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
