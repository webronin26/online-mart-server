package router

import (
	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter/admin"
)

func AdminRoute(engine *echo.Echo, middlewares ...echo.MiddlewareFunc) {
	routers := engine.Group("/admin", middlewares...)

	routers.GET("/company", admin.AdminQueryCompany)
	routers.GET("/company/:company_id", admin.AdminQueryCompanyDetail)
	routers.POST("/company", admin.AdminAddCompany)

	routers.DELETE("/logout", admin.AdminLogout)
}
