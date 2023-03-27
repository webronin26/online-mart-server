package general

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/product/search"
)

type SearchParam struct {
	Query string `query:"q"`
}

// 搜尋某產品關鍵字
func Search(ctx echo.Context) error {

	var result presenter.Result
	var param SearchParam

	if err := ctx.Bind(&param); err != nil {
		result.Code = presenter.StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input search.Input
	input.Query = param.Query

	output, serverStatus, searchErr := search.Exec(input)
	if searchErr != nil {
		result.Code = serverStatus
		result.Data = searchErr.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = serverStatus
	result.Count = output.Count
	result.Data = output.Products

	return ctx.JSON(http.StatusOK, result)
}
