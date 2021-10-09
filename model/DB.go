package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	DB   *gorm.DB
	Once sync.Once
	err  error
)

func InitDB() {
	Once.Do(func() {
		DB, err = gorm.Open("mysql", "root:Monika8899174!@tcp(127.0.0.1:3306)/strategySms?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			fmt.Println("InitDb error:", err)
			return
		}
		DB.AutoMigrate(&Exchange{})
	})
}
