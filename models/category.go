package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/astaxie/beego"
)


func init() {
	orm.RegisterModel(new(Category))
}

type Category struct {
	Id       int `json:"id" orm:"column(id);pk;unique;auto=1;int(11)"`
    Name    string `json:"name" orm:"column(name);size(128)"`
    Articles []*Article `orm:"reverse(many)"`
}

func CategoryTable() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Category))
}

// 检测分类是否存在
func CheckCategoryName(name string) bool {
	exist := CategoryTable().Filter("name", name).Exist()
	return exist
}

func AddCategory(m *Category) (*Category, error) {
	o := orm.NewOrm()
	category := Category{
		Name: m.Name,
	}

	_, err := o.Insert(&category)
	if err == nil{
		return &category, err
	}

	return nil, err
}
