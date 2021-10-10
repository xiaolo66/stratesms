package servers

import (
	"fmt"
	"stratesms/dto"
	"stratesms/model"
	"stratesms/utils"
)

type User struct {

}

func NewUser()*User{
	return &User{}
}

func(u *User)SaveUser(info dto.User)error{
	var muser model.User
	list,err :=muser.GetALLUser(info.UserName)
	if err!=nil{
		return  err
	}
	if len(list)!=0{
		return fmt.Errorf("username:%s has already exist",info.UserName)
	}
	hashSecret,err:=utils.HashAndSalt([]byte(info.PassWord))
	if err!=nil{
		return err
	}
	muser.UserName=info.UserName
	muser.Password=hashSecret
	return muser.AddUser()
}

func (u *User)CompareUser(info *dto.User)error{
	var user model.User
	list,err:=user.GetALLUser(info.UserName)
	if err!=nil{
		return  err
	}
	if len(list)==0{
		return fmt.Errorf("username unexist %s",err)
	}
	if utils.ComparePasswords(list[0].Password,[]byte(info.PassWord)){
		return nil
	}else {
		return fmt.Errorf("password wrong")
	}
}