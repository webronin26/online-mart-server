package general

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/login"
)

type MemberLoginParam struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func MemberLogin(ctx echo.Context) error {

	var result presenter.Result

	var param MemberLoginParam
	if err := ctx.Bind(&param); err != nil {
		result.Code = presenter.StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if param.Account == "" || param.Password == "" {
		result.Code = presenter.StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input login.Input
	input.Account = param.Account
	input.Password = param.Password
	input.UserKind = entities.MemberUser

	output, serverStatus, loginErr := login.Exec(input)
	if loginErr != nil {
		result.Code = serverStatus
		result.Data = loginErr.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Data = output
	result.Code = serverStatus

	return ctx.JSON(http.StatusOK, result)
}
