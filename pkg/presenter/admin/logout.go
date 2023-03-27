package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/logout"
)

func AdminLogout(ctx echo.Context) error {

	user := ctx.Get("user").(*entities.User)

	var result presenter.Result
	var input logout.Input
	input.UserId = int64(user.ID)

	if serverStatus, err := logout.Exec(input); err != nil {
		result.Code = serverStatus
		result.Data = err.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = presenter.StatusSuccess

	return ctx.JSON(http.StatusOK, result)
}
