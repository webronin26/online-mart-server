package admin

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/company/query_detail"
)

func AdminQueryCompanyDetail(ctx echo.Context) error {

	var result presenter.Result

	companyID := ctx.Param("company_id")
	if companyID == "" {
		result.Code = presenter.StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	int64ID, err := strconv.ParseInt(companyID, 10, 64)
	if err != nil {
		result.Code = presenter.StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input query_detail.Input
	input.CompanyID = int64(int64ID)

	output, serverStatus, queryErr := query_detail.Exec(input)
	if queryErr != nil {
		result.Code = serverStatus
		result.Data = queryErr.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Data = output.Company
	result.Code = serverStatus

	return ctx.JSON(http.StatusOK, result)
}
