package models

import (
	"errors"
	"fmt"
	"time"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/satori/go.uuid"
)

var expire = 3600

type Session struct {
	Id       string `json:"id"       orm:"column(id);pk;unique;size(128)"`
	Expire   int   `orm:"column(expire);null"`
	Data     time.Time `orm:"column(data);auto_now_add;type(datetime)"`
	User     *User `orm:"rel(one)"`
}

func init() {
	// orm.RegisterModel(new(Session))
}

// 将新会话插入数据库，并在成功时返回上次插入的ID。
func AddSession(s *Session) (sess *Session, err error) {
	o := orm.NewOrm()
	Id, uuiderr := uuid.NewV4()
	if uuiderr != nil {
        fmt.Printf("Something went wrong: %s", uuiderr)
        return nil, uuiderr
	}
	session := Session{
		Id: Id.String(),
		Expire: expire,
		User: s.User,
	}
	_, err = o.Insert(&session)
	if err == nil{
		return &session, err
	}
	return nil, err
}

// 按ID检索会话。如果ID不存在，则返回错误。
func GetSessionById(id string) (v *Session, err error) {
	o := orm.NewOrm()
	v = &Session{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// 按ID更新会话，如果要更新的记录不存在，则返回错误。
func UpdateSessionById(m *Session) (err error) {
	o := orm.NewOrm()
	v := Session{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 按ID删除会话，如果要删除的记录不存在，则返回错误。
func DeleteSession(id string) (err error) {
	o := orm.NewOrm()
	v := Session{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Session{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// 检索符合特定条件的所有会话。如果不存在记录，则返回空列表
func GetAllSession(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Session))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Session
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}