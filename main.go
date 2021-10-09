package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stratesms/controlls"
	"stratesms/model"
	"stratesms/servers"
)

func main() {
	model.InitDB()
	servers.InitServers()
	r := gin.Default()
	controlls.RegisterRouter(r)
	err := r.Run(":3001")
	if err != nil {
		fmt.Println("start server failed,err", err)
		return
	}
}
