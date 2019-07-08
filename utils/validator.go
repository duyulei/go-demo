package utils

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

func CheckNewUserPost(username string, password string) (errorMessage string) {
	valid := validation.Validation{}
	//表单验证
	
	valid.Required(username, "Username").Message("用户名必填")
	valid.Required(password, "Password").Message("密码必填")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			//c.Ctx.ResponseWriter.WriteHeader(403)
			//c.Data["json"] = Response{403001, 403001,err.Message, ""}
			//c.ServeJSON()
			return fmt.Sprintf("%s", err.Message)
		}
	}
	return fmt.Sprintf("%s", "ok")
}

