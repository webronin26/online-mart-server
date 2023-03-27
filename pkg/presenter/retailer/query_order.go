package retailer

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/order/retailer_query"
)

func RetailerQueryOrder(ctx echo.Context) error {

	retailer := ctx.Get("retailer").(*entities.Retailer)

	var result presenter.Result

	var input retailer_query.Input
	input.CompanyID = retailer.CompanyID

	output, serverStatus, err := retailer_query.Exec(input)
	if err != nil {
		result.Code = serverStatus
		result.Data = err.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Count = output.Count
	result.Data = output.Orders
	result.Code = serverStatus

	return ctx.JSON(http.StatusOK, result)
}
