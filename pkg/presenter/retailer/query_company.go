package retailer

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/company/query_detail"
)

func RetailerQueryCompany(ctx echo.Context) error {

	var result presenter.Result

	retailer := ctx.Get("retailer").(*entities.Retailer)

	var input query_detail.Input
	input.CompanyID = int64(retailer.CompanyID)

	output, serverStatus, err := query_detail.Exec(input)
	if err != nil {
		result.Code = serverStatus
		result.Data = err.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Data = output.Company
	result.Code = serverStatus

	return ctx.JSON(http.StatusOK, result)
}
