package models

import (
	"github.com/astaxie/beego/orm"
	_"github.com/astaxie/beego"
	"strconv"
	"runtime"
	"fmt"
)


func init() {
	orm.RegisterModel(new(Category))
}

type Category struct {
	Id       int `json:"id" orm:"column(id);pk;unique;auto;int(11)"`
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

// 根据分类ID查找分类
func GetCategoryById(id int) (err error, category *Category) {
	o := orm.NewOrm()
	category = &Category{Id: id}
	if err := o.QueryTable(new(Category)).Filter("Id", id).RelatedSel().One(category); err == nil {
		return nil, category
	}
	return err, nil

}

// 查找所有分类
func GetAllCategory(m *Category, p *Pagination) (error, []*Category) {
	fmt.Println("cpus:", runtime.NumCPU())
	page, _ := strconv.ParseInt(p.Page, 10, 0)
	pageSize, _ := strconv.ParseInt(p.PageSize, 10, 0)
	var categories []*Category
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Category)).Filter("name__contains", m.Name).Limit(pageSize, page * pageSize).All(&categories);
	count, err := o.QueryTable(new(Category)).Filter("name__contains", m.Name).Count()
	fmt.Println("count:", count)
	if err == nil {
		return nil, categories
	}
	return err, nil

}

// 添加
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

// 删除
func DeleteCategory(m *Category) (err error) {
	o := orm.NewOrm()
	category := Category{
		Id: m.Id,
	}
	err = o.Read(&category);
	if err == nil {
		var num int64
		if num, err = o.Delete(&category); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}else{
			return err
		}
	}else{
		return err
	}
	return err
}

// 更新
func UpdateCategoryById(m *Category) (err error) {
	o := orm.NewOrm()
	v := Category{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}