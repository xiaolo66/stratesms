package dto

type Exchange struct {
	ExchangeName string `form:"name" json:"name" binding:"required"`
	Website      string `form:"website"json:"website" `
	Describe     string `form:"describe"json:"describe"`
}