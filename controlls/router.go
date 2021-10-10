package controlls

import "github.com/gin-gonic/gin"

func RegisterRouter(engine *gin.Engine) {
	exchange := engine.Group("/home")
	{
		exchangeCtrl := NewExchange()
		exchange.POST("/exchange", exchangeCtrl.AddExchange())
		exchange.POST("/exchanges",exchangeCtrl.GetExchange())
		exchange.GET("/exchanges",exchangeCtrl.GetAllExchange())
		exchange.PUT("/exchange",exchangeCtrl.PutExchange())
		exchange.DELETE("/exchange",exchangeCtrl.Delete())
	}
	login:=engine.Group("/login")
	{
		userCtrl:=NewUser()
		login.POST("/register",userCtrl.Register())
		login.POST("/userlogin",userCtrl.UserLogin())
	}

}
