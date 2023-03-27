package retailer

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/order/retailer_query_detail"
)

func RetailerQueryOrderDetail(ctx echo.Context) error {

	retailer := ctx.Get("retailer").(*entities.Retailer)

	var result presenter.Result

	orderID := ctx.Param("order_id")
	if orderID == "" {
		result.Code = presenter.StatusParamError
		result.Data = "order id is null"
		return ctx.JSON(http.StatusBadRequest, result)
	}

	int64ID, err := strconv.ParseInt(orderID, 10, 64)
	if err != nil {
		result.Code = presenter.StatusParamError
		result.Data = err.Error()
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input retailer_query_detail.Input
	input.CompanyID = retailer.CompanyID
	input.OrderID = uint64(int64ID)

	output, serverStatus, queryErr := retailer_query_detail.Exec(input)
	if queryErr != nil {
		result.Code = serverStatus
		result.Data = queryErr.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = serverStatus
	result.Data = output.Order

	return ctx.JSON(http.StatusOK, result)
}
