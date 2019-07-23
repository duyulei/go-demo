package models

import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Comment))
}

type Comment struct {
	Id          int       `json:"id" orm:"column(id);pk;unique;auto;int(11)"`
	Content     string    `json:"content" orm:"column(content);size(256)"`
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	User        *User     `json:"user_id" orm:"rel(fk)"`
	Article     *Article  `json:"article_id" orm:"rel(fk)"`
}

// 添加评论
func AddComment(m *Comment) (*Comment, error) {
	o := orm.NewOrm()
	comment := Comment{
		Content: m.Content,
		User:    m.User,
		Article: m.Article,
	}

	_, err := o.Insert(&comment)
	if err == nil {
		return &comment, err
	}

	return nil, err
}

// 查询用户写的评论
func GetUserComment(m *Comment, p *Pagination) (error, []*Comment) {
	page, _ := strconv.ParseInt(p.Page, 10, 0)
	pageSize, _ := strconv.ParseInt(p.PageSize, 10, 0)
	var comments []*Comment
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Comment)).Filter("user_id", m.User.Id).Limit(pageSize, page*pageSize).All(&comments)
	if err == nil {
		return nil, comments
	}
	return err, nil
}

// 查询文章评论
func GetArticleComment(m *Comment, p *Pagination) (error, []*Comment) {
	page, _ := strconv.ParseInt(p.Page, 10, 0)
	pageSize, _ := strconv.ParseInt(p.PageSize, 10, 0)
	article_id := m.Article.Id
	var comments []*Comment
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Comment)).Filter("article_id", article_id).Limit(pageSize, page*pageSize).All(&comments)
	if err == nil {
		return nil, comments
	}
	return err, nil
}

// 删除
func DeleteComment(m *Comment) (err error) {
	o := orm.NewOrm()
	comment := Comment{
		Id: m.Id,
	}
	err = o.Read(&comment)
	if err == nil {
		var num int64
		if num, err = o.Delete(&comment); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		} else {
			return err
		}
	} else {
		return err
	}
	return err
}
