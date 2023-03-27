package general

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/login"
)

type RetailerLoginParam struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func RetailerLogin(ctx echo.Context) error {

	var result presenter.Result

	var param RetailerLoginParam
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
	input.UserKind = entities.RetailerUser

	output, serverStatus, loginErr := login.Exec(input)
	if loginErr != nil {
		result.Code = serverStatus
		result.Data = loginErr.Error()
		return ctx.JSON(http.StatusNotFound, result)
	}

	result.Data = output
	result.Code = presenter.StatusSuccess

	return ctx.JSON(http.StatusOK, result)
}
