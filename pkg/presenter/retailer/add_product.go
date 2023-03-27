package retailer

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/product/add"
)

type AddProductParam struct {
	ProductName     string  `json:"name"`
	ProductImageURL string  `json:"image_url"`
	Summary         string  `json:"summary"`
	Information     string  `json:"information"`
	ProductPrice    float64 `json:"price"`
	InventoryNumber int64   `json:"number"`
	MaxBuy          int64   `json:"max_buy"`
}

func RetailerAddProduct(ctx echo.Context) error {

	var result presenter.Result

	retailer := ctx.Get("retailer").(*entities.Retailer)

	var param AddProductParam
	if err := ctx.Bind(&param); err != nil {
		result.Code = presenter.StatusParamError
		result.Data = "param bind error"
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input add.Input
	input.ProductName = param.ProductName
	input.ProductImageURL = param.ProductImageURL
	input.Summary = param.Summary
	input.Information = param.Information
	input.ProductPrice = param.ProductPrice
	input.InventoryNumber = param.InventoryNumber
	input.MaxBuy = param.MaxBuy
	input.CompanyID = retailer.CompanyID

	serverStatus, addErr := add.Exec(input)
	if addErr != nil {
		result.Code = serverStatus
		result.Data = addErr.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = serverStatus
	result.Code = presenter.StatusSuccess

	return ctx.JSON(http.StatusOK, result)
}
