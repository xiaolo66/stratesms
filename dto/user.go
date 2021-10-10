package dto

type User struct {
	UserName string `form:"username",json:"username",binding:"required"`
	PassWord string `form:"password",json:"password",binding:"required"`
}
