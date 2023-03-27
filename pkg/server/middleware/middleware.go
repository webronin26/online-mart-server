package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

var (
	Recover = middleware.Recover()
	CORS    = middleware.CORS()
	Logger  = middleware.Logger()
)

// 檢查 token 部分
func certifyToken(ctx echo.Context) string {
	return ctx.Request().Header.Get("Authorization")
}

func CertifyAdmin(hfc echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var result presenter.Result

		token := certifyToken(ctx)
		if token == "" {
			result.Code = presenter.StatusTokenError
			return ctx.JSON(http.StatusUnauthorized, result)
		}

		user := new(entities.User)
		query := store.DB.Model(user).
			Where("token = ?", token).
			Where("user_kind = ?", entities.AdminUser).
			First(&user)
		if query.RecordNotFound() {
			result.Code = presenter.StatusTokenError
			result.Data = "token not certify"
			return ctx.JSON(http.StatusUnauthorized, result)
		}
		if query.Error != nil {
			result.Code = presenter.StatusServerError
			result.Data = "token query error"
			return ctx.JSON(http.StatusInternalServerError, result)
		}

		ctx.Set("user", user)

		return hfc(ctx)
	}
}

func CertifyMember(hfc echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var result presenter.Result

		token := certifyToken(ctx)
		if token == "" {
			result.Code = presenter.StatusTokenError
			return ctx.JSON(http.StatusUnauthorized, result)
		}

		user := new(entities.User)
		query := store.DB.Model(user).
			Where("token = ?", token).
			Where("user_kind = ?", entities.MemberUser).
			First(&user)
		if query.RecordNotFound() {
			result.Code = presenter.StatusTokenError
			result.Data = "token not certify"
			return ctx.JSON(http.StatusUnauthorized, result)
		}
		if query.Error != nil {
			result.Code = presenter.StatusServerError
			result.Data = "token query error"
			return ctx.JSON(http.StatusInternalServerError, result)
		}

		member := new(entities.Member)
		queryMember := store.DB.Model(member).
			Where("user_id = ?", user.ID).
			First(&member)
		if queryMember.RecordNotFound() {
			result.Code = presenter.StatusTokenError
			result.Data = "query member not found"
			return ctx.JSON(http.StatusUnauthorized, result)
		}
		if queryMember.Error != nil {
			result.Code = presenter.StatusServerError
			result.Data = "query member error"
			return ctx.JSON(http.StatusInternalServerError, result)
		}

		ctx.Set("user", user)
		ctx.Set("member", member)

		return hfc(ctx)
	}
}

func CertifyRetailer(hfc echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var result presenter.Result

		token := certifyToken(ctx)
		if token == "" {
			result.Code = presenter.StatusTokenError
			return ctx.JSON(http.StatusUnauthorized, result)
		}

		user := new(entities.User)
		query := store.DB.Model(user).
			Where("token = ?", token).
			Where("user_kind = ?", entities.RetailerUser).
			First(&user)
		if query.RecordNotFound() {
			result.Code = presenter.StatusTokenError
			result.Data = "token not certify"
			return ctx.JSON(http.StatusUnauthorized, result)
		}
		if query.Error != nil {
			result.Code = presenter.StatusServerError
			result.Data = "token query error"
			return ctx.JSON(http.StatusInternalServerError, result)
		}

		retailer := new(entities.Retailer)
		queryRetailer := store.DB.Model(retailer).
			Where("user_id = ?", user.ID).
			First(&retailer)
		if queryRetailer.RecordNotFound() {
			result.Code = presenter.StatusTokenError
			result.Data = "query retailer not found"
			return ctx.JSON(http.StatusUnauthorized, result)
		}
		if queryRetailer.Error != nil {
			result.Code = presenter.StatusServerError
			result.Data = "query retailer error"
			return ctx.JSON(http.StatusInternalServerError, result)
		}

		ctx.Set("user", user)
		ctx.Set("retailer", retailer)

		return hfc(ctx)
	}
}
