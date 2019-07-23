package models

import (
	_"errors"
	"github.com/astaxie/beego/orm"
	_"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"rompapi/utils"
	"fmt"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
    Token  string
}

type Pagination struct {
    Page      string
    PageSize  string
}

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id       string `json:"id"       orm:"column(id);pk;unique;size(128)"`
	Username string `json:"username" orm:"column(username);unique;size(32)"`
	Password string `json:"password" orm:"column(password);size(128)"`
	Create_time  time.Time `orm:"auto_now_add;type(datetime)"`
	Token 	 string `json:"token" orm:"column(token);size(256)"`
	Articles []*Article `orm:"reverse(many)"`
}

func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(User))
}

// 检测用户是否存在
func CheckUserId(userId string) bool {
	exist := Users().Filter("Id", userId).Exist()
	return exist
}

// 检测用户是否存在
func CheckUserName(username string) bool {
	exist := Users().Filter("Username", username).Exist()
	return exist
}

// 根据token查找用户
func GetUserByToken(token string) (err error, user *User) {
	o := orm.NewOrm()
	user = &User{Token: token}
	if err := o.QueryTable(new(User)).Filter("Token", token).RelatedSel().One(user); err == nil {
		return nil, user
	}
	return err, nil

}

// 根据用户名查找用户
func GetUserByUsername(username string) (err error, user *User) {
	o := orm.NewOrm()
	user = &User{Username: username}
	if err := o.QueryTable(user).Filter("Username", username).One(&user); err == nil {
		return nil, user
	}
	return err, nil
}

// 根据用户ID查找用户
func GetUserById(id string) (err error, user *User) {
	o := orm.NewOrm()
	user = &User{Id: id}
	if err := o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(user); err == nil {
		return nil, user
	}
	return err, nil

}

// 更新Token
func UpdateUserToken(m *User, token string) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	m.Token = token
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return err
}

func AddUser(m *User) (*User, error) {
	o := orm.NewOrm()
	Id, uuiderr := uuid.NewV4()
    if uuiderr != nil {
        fmt.Printf("Something went wrong: %s", uuiderr)
        return nil, uuiderr
    }
	hash, psderr := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
    if psderr != nil {
        fmt.Println(psderr)
    }
	encodePW := string(hash)
	token := utils.GenToken()
	user := User{
		Id: Id.String(),
		Username:m.Username,
		Password: encodePW,
		Token: token,
	}

	fmt.Println("数据是：", user)
	
	_, err := o.Insert(&user)
	fmt.Println("错误是：", err)
	if err == nil{
		return &user, err
	}

	return nil, err
}

func Login(m User) (bool, User) {
	o := orm.NewOrm()
	var user User
	// user := &User{}
	if err := o.QueryTable("user").Filter("Username", m.Username).One(&user); err == nil {
		// 验证正确密码
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m.Password))
		if err != nil {
			return false, user
		} else {
			return true, user
		}
	}
	return false, user
}