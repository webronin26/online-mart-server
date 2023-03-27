package member

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/cart/add_product"
)

type AddMemberCartParam struct {
	ProductID uint64 `json:"product_id"`
	Number    int64  `json:"number"`
}

func MemberAddCart(ctx echo.Context) error {

	var result presenter.Result

	member := ctx.Get("member").(*entities.Member)

	var param AddMemberCartParam
	if err := ctx.Bind(&param); err != nil {
		result.Code = presenter.StatusBindFailed
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if param.Number == 0 || param.ProductID == 0 || param.ProductID < 0 {
		result.Code = presenter.StatusParamValidateFailed
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input add_product.Input
	input.MemberID = member.ID
	input.ProductID = param.ProductID
	input.Number = param.Number

	output, statusCode, addErr := add_product.Exec(input)
	if addErr != nil {
		result.Code = statusCode
		result.Data = addErr.Error()
		return ctx.JSON(http.StatusBadRequest, result)
	}

	result.Code = statusCode
	result.Data = output

	return ctx.JSON(http.StatusOK, result)
}
