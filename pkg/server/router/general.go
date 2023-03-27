package router

import (
	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter/general"
)

func GeneralRoute(router *echo.Echo) {
	router.GET("/index", general.Index)

	router.GET("/search", general.Search)
	router.GET("/product/:product_id", general.QueryProductDetail)
	router.GET("/company/product/:company_id", general.QueryCompanyProduct)

	router.POST("/login/member", general.MemberLogin)
	router.POST("/login/admin", general.AdminLogin)
	router.POST("/login/retailer", general.RetailerLogin)
}
