package models

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// UserData 用户信息
type UserData struct {
	ID       int64
	UserName string
	Password string
	Email    string
}

// InsertUser 插入用户信息
func InsertUser(user *UserData) (e error) {
	if user == nil {
		return errors.New("user is nil")
	}
	// 1. 有orm对象
	o := orm.NewOrm()
	// 2. 插入那个数据
	id, e := o.Insert(user)
	if e != nil {
		logs.Error("o.Insert failed, user = %v, error = %v", *user, e.Error())
		return e
	}
	user.ID = id
	logs.Debug("insert user = %v successfully", *user)
	return nil
}

// LoadUser 加载用户信息
func LoadUser(username string) (e error, user *UserData) {
	// 1. 有orm对象
	o := orm.NewOrm()
	user = &UserData{}
	user.UserName = username
	e = o.Read(user, FieldUserName)
	if e != nil {
		logs.Error("o.Read failed, error = %v", e.Error())
		return e, nil
	} else {
		logs.Debug("load user successfully, user = %v", *user)
		return nil, user
	}
}

// UpdatePassword 更新密码
func UpdatePassword(username string, newPassword string) (e error) {
	if username == "" || newPassword == "" {
		logs.Error("username or newPassword is nil, username = %v, newPassword = %v", username, newPassword)
		return errors.New("data invalid")
	}

	o := orm.NewOrm()
	user := &UserData{UserName:username}
	e = o.Read(user, FieldUserName)
	if e != nil {
		logs.Error("o.Read failed, error = %v", e.Error())
		return e
	}
	if user.Password != newPassword {
		user.Password = newPassword
		_, e = o.Update(user, FieldPassword)
		if e != nil {
			logs.Error("o.Update failed, error = %v", e.Error())
			return e
		} else {
			logs.Info("o.Update successfully")
			return nil
		}
	} else {
		// 密码一样 不需要修改
		logs.Info("new password is same to old password")
		return nil
	}
}

// DeleteUser 删除用户
func DeleteUser(username string) (e error) {
	if username == "" {
		logs.Error("username is nil")
		return errors.New("username is nil")
	}
	
	o := orm.NewOrm()
	user := &UserData{UserName: username}
	_, e = o.Delete(user)
	if e != nil {
		logs.Error("o.Delete failed, error = %v", e.Error())
		return e
	}
	return nil
}