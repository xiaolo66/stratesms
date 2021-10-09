package controlls

import "github.com/gin-gonic/gin"

func RegisterRouter(engine *gin.Engine) {
	exchange := engine.Group("/home")
	{
		exchangeCtrl := NewExchange()
		exchange.POST("/exchange", exchangeCtrl.Add())
		exchange.POST("/exchanges",exchangeCtrl.Get())

		exchange.GET("/exchanges",exchangeCtrl.GetAll())

		exchange.PUT("/exchange",exchangeCtrl.Put())


		exchange.DELETE("/exchange",exchangeCtrl.Delete())
	}

}
