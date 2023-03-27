package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/company/query"
)

func AdminQueryCompany(ctx echo.Context) error {

	var result presenter.Result

	output, serverStatus, err := query.Exec()
	if err != nil {
		result.Code = serverStatus
		result.Data = err.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Count = output.Count
	result.Data = output.Companies
	result.Code = serverStatus

	return ctx.JSON(http.StatusOK, result)
}
