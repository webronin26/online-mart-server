package general

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/product/query_detail"
)

func QueryProductDetail(ctx echo.Context) error {

	var result presenter.Result
	productID := ctx.Param("product_id")
	if productID == "" {
		result.Code = presenter.StatusParamError
		result.Data = "product id is null"
		return ctx.JSON(http.StatusBadRequest, result)
	}
	int64ID, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		result.Code = presenter.StatusParamError
		result.Data = err
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input query_detail.Input
	input.ProductID = uint64(int64ID)

	output, serverStatus, queryErr := query_detail.Exec(input)
	if queryErr != nil {
		result.Code = serverStatus
		result.Data = queryErr.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = serverStatus
	result.Data = output.Product

	return ctx.JSON(http.StatusOK, result)
}
