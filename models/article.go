package models

import (
	_"errors"
	"github.com/astaxie/beego/orm"
	_"github.com/astaxie/beego"
	"time"
)


func init() {
	orm.RegisterModel(new(Article))
}

type Article struct {
	Id       int `json:"id" orm:"column(id);pk;unique;auto;int(11)"`
    Title    string `json:"title" orm:"column(title);size(256)"`
    Submit   string `json:"submit" orm:"column(submit);size(256)"`
	Create_time  time.Time `orm:"auto_now_add;type(datetime)"`
	User     *User   `json:"user_id" orm:"rel(fk)"`
	Category *Category `json:"category_id" orm:"rel(fk)"`
    Comments []*Comment   `orm:"reverse(many)"` //反向一对多关联
}


func AddArticle(m *Article) (*Article, error) {
	o := orm.NewOrm()
	article := Article{
		Title: m.Title,
		Submit: m.Submit,
	}

	_, err := o.Insert(&article)
	if err == nil{
		return &article, err
	}

	return nil, err
}