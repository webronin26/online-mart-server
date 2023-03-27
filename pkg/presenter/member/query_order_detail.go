package member

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/order/query_detail"
)

func MemberQueryOrderDetail(ctx echo.Context) error {

	member := ctx.Get("member").(*entities.Member)

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

	var input query_detail.Input
	input.MemberID = member.ID
	input.OrderID = uint64(int64ID)

	output, serverStatus, queryErr := query_detail.Exec(input)
	if queryErr != nil {
		result.Code = serverStatus
		result.Data = queryErr.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = serverStatus
	result.Data = output.Order

	return ctx.JSON(http.StatusOK, result)
}
