package general

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/index"
)

func Index(ctx echo.Context) error {

	var result presenter.Result

	output, statusCode, err := index.Exec()
	if err != nil {
		result.Code = statusCode
		result.Data = err.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = statusCode
	result.Data = output

	return ctx.JSON(http.StatusOK, result)
}
