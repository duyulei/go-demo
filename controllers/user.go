package controllers

import (
	"rompapi/models"
	"rompapi/utils"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserController) URLMapping() {
	// c.Mapping("Post", c.Post)
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /regist [post]
func (c *UserController) Regist() {
	var v models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if usernameLen := len(v.Username); usernameLen<3{
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,"用户名不少于3位", nil}
			c.ServeJSON()
			return
		}
		if exist := models.CheckUserId(v.Id); exist{
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = "用户存在"
			c.ServeJSON()
			return
		}
		if exist := models.CheckUserName(v.Username); exist{
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,"用户名存在", nil}
			c.ServeJSON()
			return
		}
		if errorMessage := utils.CheckNewUserPost(v.Username, v.Password); errorMessage != "ok"{
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,errorMessage, ""}
			c.ServeJSON()
			return
		}
		if user, err := models.AddUser(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = Response{200, 200,"注册成功", user}
		} else {
			c.Data["json"] = &Response{1, 1, "用户注册失败", err.Error()}
		}
	}else{
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
		fmt.Println("错误是：", err)
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = fmt.Sprintf("%s", err)
		c.ServeJSON()
		return
	}
	c.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [POST]
func (c *UserController) Login() {
	var v models.User
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if errorMessage := utils.CheckNewUserPost(v.Username, v.Password); errorMessage != "ok"{
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403,errorMessage, ""}
			c.ServeJSON()
			return
		}
		if ok, user := models.Login(v); ok{
			// TODO-------------------------------------------
			if ok := utils.CheckToken(user.Token); !ok{
				token := utils.GenToken()
				models.UpdateUserToken(&user, token)
			}
			c.Ctx.Output.SetStatus(201)
			c.SetSession("user_id", user.Id)
			c.Data["json"] = Response{200, 200,"登录成功", user}
		} else {
			c.Data["json"] = Response{200, 200, "密码错误", nil}
		}
	}else{
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = Response{403, 403, fmt.Sprintf("%s", "登录失败"), nil}
		c.ServeJSON()
		return
	}
	c.ServeJSON()
}

// @router /test [POST]
func (c *UserController) Test() {
	errorMessage:=int64(time.Now().Unix() + 1000)
	c.Data["json"] = Response{200, 200, fmt.Sprintf("%v", errorMessage), nil}
	c.ServeJSON()
	return
}