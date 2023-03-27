package member

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/order/query"
)

func MemberQueryOrder(ctx echo.Context) error {

	member := ctx.Get("member").(*entities.Member)

	var result presenter.Result

	var input query.Input
	input.MemberID = member.ID
	output, serverStatus, err := query.Exec(input)
	if err != nil {
		result.Code = serverStatus
		result.Data = err.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Count = output.Count
	result.Code = serverStatus
	result.Data = output.Orders

	return ctx.JSON(http.StatusOK, result)
}
