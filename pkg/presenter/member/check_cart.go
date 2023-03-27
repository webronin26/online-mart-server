package member

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/cart/check"
)

type MemberCheckCartParam struct {
	Address string `json:"address"`
}

func MemberCheckCart(ctx echo.Context) error {

	var result presenter.Result
	member := ctx.Get("member").(*entities.Member)

	var param MemberCheckCartParam
	if err := ctx.Bind(&param); err != nil {
		result.Code = presenter.StatusBindFailed
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if param.Address == "" {
		result.Code = presenter.StatusParamValidateFailed
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input check.Input
	input.MemberID = member.ID
	input.MemberAddress = param.Address

	serverStatus, checkErr := check.Exec(input)
	if checkErr != nil {
		result.Code = serverStatus
		result.Data = checkErr.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = presenter.StatusSuccess

	return ctx.JSON(http.StatusOK, result)
}
