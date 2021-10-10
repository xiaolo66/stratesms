package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string
	Password string
}

func(u *User)GetALLUser(name string)(list []User,err error){
	err=DB.Debug().Where("user_name=?",name).Find(&list).Error
	return list,err
}

func(u *User)AddUser()error{
	err:=DB.Debug().Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}
