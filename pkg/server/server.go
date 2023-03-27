package server

import (
	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/config"
	"github.com/webronin26/online-mart-server/pkg/server/middleware"
	"github.com/webronin26/online-mart-server/pkg/server/router"
)

func Init() {

	e := echo.New()

	// 註冊中間件
	e.Use(
		middleware.Recover,
		middleware.CORS,
		middleware.Logger,
	)

	// 註冊各個 route
	// general (通用通道)：不用驗證即可通過
	router.GeneralRoute(e)

	// admin (管理者通道)：需要管理者驗證
	router.AdminRoute(e, middleware.CertifyAdmin)

	// member (會員通道)：需要會員驗證
	router.MemberRoute(e, middleware.CertifyMember)

	// retailer (店家通道)：需要店家驗證
	router.RetailerRoute(e, middleware.CertifyRetailer)

	e.Logger.Fatal(e.Start(":" + config.GetSystemConfig().Address))
}
