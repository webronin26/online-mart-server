package retailer

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/product/update"
)

type UpdateProductParam struct {
	ProductName     string                 `json:"name"`
	ProductImageURL string                 `json:"image_url"`
	Summary         string                 `json:"summary"`
	Information     string                 `json:"information"`
	ProductPrice    float64                `json:"price"`
	InventoryNumber int64                  `json:"number"`
	MaxBuy          int64                  `json:"max_buy"`
	ProductStatus   entities.ProductStatus `json:"product_status"`
}

func RetailerUpdateProduct(ctx echo.Context) error {

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
		result.Data = err.Error()
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var param UpdateProductParam
	if BindErr := ctx.Bind(&param); BindErr != nil {
		result.Code = presenter.StatusParamError
		result.Data = BindErr.Error()
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input update.Input
	input.ProductID = uint64(int64ID)
	input.ProductName = param.ProductName
	input.ProductImageURL = param.ProductImageURL
	input.Summary = param.Summary
	input.Information = param.Information
	input.ProductPrice = param.ProductPrice
	input.InventoryNumber = param.InventoryNumber
	input.MaxBuy = param.MaxBuy
	input.ProductStatus = param.ProductStatus

	serverStatus, updateErr := update.Exec(input)
	if updateErr != nil {
		result.Code = serverStatus
		result.Data = updateErr.Error()
		return ctx.JSON(http.StatusNotFound, result)
	}

	result.Code = serverStatus

	return ctx.JSON(http.StatusOK, result)
}
