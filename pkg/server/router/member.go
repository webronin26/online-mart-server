package router

import (
	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter/member"
)

func MemberRoute(engine *echo.Echo, middlewares ...echo.MiddlewareFunc) {
	router := engine.Group("/member", middlewares...)

	router.DELETE("/logout", member.MemberLogout)

	router.GET("/cart", member.MemberQueryCart)
	router.POST("/cart", member.MemberAddCart)
	router.POST("/cart/check", member.MemberCheckCart)

	router.GET("/order", member.MemberQueryOrder)
	router.GET("/order/:order_id", member.MemberQueryOrderDetail)
}
