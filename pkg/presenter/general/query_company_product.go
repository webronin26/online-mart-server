package general

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/company/query_product"
)

func QueryCompanyProduct(ctx echo.Context) error {

	var result presenter.Result

	CompanyID := ctx.Param("company_id")
	if CompanyID == "" {
		result.Code = presenter.StatusParamError
		result.Data = "company_id is null"
		return ctx.JSON(http.StatusBadRequest, result)
	}

	int64ID, err := strconv.ParseInt(CompanyID, 10, 64)
	if err != nil {
		result.Code = presenter.StatusParamError
		result.Data = err.Error()
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input query_product.Input
	input.CompanyID = uint64(int64ID)

	output, serverStatus, queryErr := query_product.Exec(input)
	if queryErr != nil {
		result.Code = serverStatus
		result.Data = queryErr.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = serverStatus
	result.Data = output.Data
	result.Count = output.Count

	return ctx.JSON(http.StatusOK, result)
}
