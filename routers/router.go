// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"rompapi/controllers"
	"rompapi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"encoding/json"
	"rompapi/utils"
)
//过滤器函数
func filFunc(ctx *context.Context){
	var f models.Token
	if path := ctx.Request.URL.Path; path != "/romp/user/login" && path != "/romp/user/regist" {
		if err := json.Unmarshal(ctx.Input.RequestBody, &f); err == nil {
			if (f.Token==""){
				ctx.ResponseWriter.WriteHeader(403)
				ctx.Output.JSON(controllers.Response{403, 403,"TOKEN不能为空", nil}, false, true)
				return
			}else{
				if ok := utils.CheckToken(f.Token); !ok{
					ctx.ResponseWriter.WriteHeader(403)
					ctx.Output.JSON(controllers.Response{403, 403,"TOKEN无效", nil}, false, true)
					return
				}
			}
		}
	}
	
}

func init() {
	beego.InsertFilter("/romp/*",beego.BeforeExec,filFunc)
	
	ns := beego.NewNamespace("/romp",
		beego.NSNamespace("/article",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/category",
			beego.NSInclude(
				&controllers.CategoryController{},
			),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
