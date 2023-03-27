package update

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	ProductID       uint64
	ProductName     string
	ProductImageURL string
	Summary         string
	Information     string
	ProductPrice    float64
	InventoryNumber int64
	MaxBuy          int64
	CompanyID       uint64
	ProductStatus   entities.ProductStatus
}

func Exec(input Input) (presenter.StatusCode, error) {

	updateParam := make(map[string]interface{})

	if input.ProductName != "" {
		updateParam["product_name"] = input.ProductName
	}
	if input.ProductImageURL != "" {
		updateParam["product_image_url"] = input.ProductImageURL
	}
	if input.Summary != "" {
		updateParam["summary"] = input.Summary
	}
	if input.Information != "" {
		updateParam["information"] = input.Information
	}
	if input.ProductPrice != 0 {
		updateParam["product_price"] = input.ProductPrice
	}
	if input.InventoryNumber != 0 {
		updateParam["inventory_number"] = input.InventoryNumber
	}
	if input.MaxBuy != 0 {
		updateParam["max_buy"] = input.MaxBuy
	}
	if input.ProductStatus != 0 {
		updateParam["product_status"] = input.ProductStatus
	}

	err := store.DB.Model(entities.Product{}).
		Where("id = ?", input.ProductID).
		Updates(updateParam).
		Error
	if err != nil {
		return presenter.StatusSQLError, errors.New("update error : " + err.Error())
	}

	return presenter.StatusSuccess, nil
}
