package controlls

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"stratesms/dto"
	"stratesms/servers"
)

type User struct {
}

func NewUser() User {
	return User{}
}

//用户注册
func (u *User) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInfo dto.User
		err := c.ShouldBind(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, dto.Response{
				Code:    1000,
				Message: "data parser failed",
				Data:    nil,
			})
			return
		}
		err = servers.UserSvr.SaveUser(userInfo)
		if err != nil {
			c.JSON(http.StatusOK, dto.Response{
				Code:    2001,
				Message: fmt.Sprintf("%s", err),
				Data:    nil,
			})
			return
		}
		c.JSON(http.StatusOK, dto.Response{
			Code:    200,
			Message: "ok",
			Data:    nil,
		})

	}
}

//用户登录

func(u *User)UserLogin()gin.HandlerFunc{
	return func(c *gin.Context) {
		var userInfo dto.User
		err:=c.ShouldBind(&userInfo)
		if err!=nil{
			c.JSON(http.StatusOK,dto.Response{
				Code: 2001,
				Message: "parser data failed",
				Data: nil,
			})
			return
		}
		err=servers.UserSvr.CompareUser(&userInfo)
		if err!=nil{
			c.JSON(http.StatusOK,dto.Response{
				Code: 2002,
				Message: fmt.Sprintf("compare password failed:%s",err),
				Data: nil,
			})
			return
		}
		c.JSON(http.StatusOK, dto.Response{
			Code:    200,
			Message: "ok",
			Data:    nil,
		})
	}
}
