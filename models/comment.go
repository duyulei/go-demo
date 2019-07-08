package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/astaxie/beego"
	"time"
)


func init() {
	orm.RegisterModel(new(Comment))
}

type Comment struct {
	Id       int `json:"id" orm:"column(id);pk;unique;auto;int(11)"`
    Content    string `json:"content" orm:"column(content);size(256)"`
	Create_time  time.Time `orm:"auto_now_add;type(datetime)"`
	Article     *Article   `json:"article_id" orm:"rel(fk)"`
}

