package router

import (
	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter/retailer"
)

func RetailerRoute(engine *echo.Echo, middlewares ...echo.MiddlewareFunc) {
	routers := engine.Group("/retailer", middlewares...)

	routers.GET("/product", retailer.RetailerQueryProduct)
	routers.GET("/product/:product_id", retailer.RetailerQueryProductDetail)
	routers.POST("/product", retailer.RetailerAddProduct)
	routers.PATCH("/product/:product_id", retailer.RetailerUpdateProduct)

	routers.GET("/order", retailer.RetailerQueryOrder)
	routers.GET("/order/:order_id", retailer.RetailerQueryOrderDetail)

	routers.GET("/company/:company_id", retailer.RetailerQueryCompany)
	routers.PATCH("/company", retailer.RetailerUpdateCompany)

	routers.DELETE("/logout", retailer.RetailerLogout)
}
