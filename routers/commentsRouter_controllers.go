package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["rompapi/controllers:ArticleController"] = append(beego.GlobalControllerRouter["rompapi/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "AddArticle",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["rompapi/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "AddCategory",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["rompapi/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["rompapi/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "GetAllCategory",
            Router: `/getall`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:CategoryController"] = append(beego.GlobalControllerRouter["rompapi/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Updata",
            Router: `/updata`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:CommentController"] = append(beego.GlobalControllerRouter["rompapi/controllers:CommentController"],
        beego.ControllerComments{
            Method: "AddComment",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:CommentController"] = append(beego.GlobalControllerRouter["rompapi/controllers:CommentController"],
        beego.ControllerComments{
            Method: "GetArticleComment",
            Router: `/artcomment`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:CommentController"] = append(beego.GlobalControllerRouter["rompapi/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/delete`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:CommentController"] = append(beego.GlobalControllerRouter["rompapi/controllers:CommentController"],
        beego.ControllerComments{
            Method: "GetUSerComment",
            Router: `/usercomment`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:UserController"] = append(beego.GlobalControllerRouter["rompapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:UserController"] = append(beego.GlobalControllerRouter["rompapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Regist",
            Router: `/regist`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["rompapi/controllers:UserController"] = append(beego.GlobalControllerRouter["rompapi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Test",
            Router: `/test`,
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
