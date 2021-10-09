package controlls

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"stratesms/dto"
	"stratesms/servers"
)

type Exchange struct {
}

func NewExchange() *Exchange {
	return &Exchange{}
}

func (e *Exchange) Add() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var info dto.Exchange
		err := ctx.ShouldBind(&info)
		if err != nil {
			ctx.JSON(http.StatusOK, dto.Response{
				Code:    1000,
				Message: "data parse error",
				Data:    nil,
			})
			return
		}
		err = servers.ExchangeSvr.Add(&info)
		if err != nil {
			ctx.JSON(http.StatusOK, dto.Response{
				Code:    1001,
				Message: "data parse error",
				Data:    nil,
			})
			return
		}
		ctx.JSON(http.StatusOK, dto.Response{
			Code:    200,
			Message: "ok",
			Data:    nil,
		})

	}
}

func (e *Exchange) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		exchanges, err := servers.ExchangeSvr.GetAll()
		if err != nil {
			ctx.JSON(http.StatusOK, dto.Response{
				Code:    1001,
				Message: "data parse error",
				Data:    nil,
			})
			return
		}
		ctx.JSON(http.StatusOK, dto.Response{
			Code:    200,
			Message: "ok",
			Data:    exchanges,
		})
	}
}

func (e *Exchange) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var info dto.Exchange
		err := ctx.ShouldBind(&info)
		fmt.Printf("%#v", info)
		if err != nil {
			ctx.JSON(http.StatusOK, dto.Response{
				Code:    1000,
				Message: "data parse error",
				Data:    nil,
			})
			return
		}
		exchange, err := servers.ExchangeSvr.Get(&info)
		if err != nil {
			ctx.JSON(http.StatusOK, dto.Response{
				Code:    1001,
				Message: "data parse error",
				Data:    nil,
			})
			return
		}
		ctx.JSON(http.StatusOK, dto.Response{
			Code:    200,
			Message: "ok",
			Data:    exchange,
		})
	}
}

func (e *Exchange) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var info dto.Exchange
		err := ctx.ShouldBind(&info)
		if err != nil {
			ctx.JSON(200, dto.Response{Code: 1000,
				Message: "data parse error",
				Data:    nil,
			})
			return
		}
		err = servers.ExchangeSvr.Put(&info)
		if err != nil {
			ctx.JSON(http.StatusOK, dto.Response{
				Code:    1006,
				Message: "data update error",
				Data:    nil,
			})
			return
		}
		ctx.JSON(http.StatusOK, dto.Response{
			Code:    200,
			Message: "ok",
			Data:    nil,
		})
	}
}

func (e *Exchange)Delete()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		exchange,ok:=ctx.GetQuery("name")
		fmt.Println(1,exchange)
		if !ok{
			ctx.JSON(http.StatusOK, dto.Response{
				Code:    1001,
				Message: "not param delete",
				Data:    nil,
			})
			return
		}
		err:=servers.ExchangeSvr.Delete(exchange)
		if err != nil {
			ctx.JSON(http.StatusOK, dto.Response{
				Code:    1005,
				Message: "Delete error",
				Data:    nil,
			})
			return
		}
		ctx.JSON(http.StatusOK, dto.Response{
			Code:    200,
			Message: "ok",
			Data:    nil,
		})
	}
}
