package add

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	ProductName     string
	ProductImageURL string
	Summary         string
	Information     string
	ProductPrice    float64
	InventoryNumber int64
	MaxBuy          int64
	CompanyID       uint64
}

func Exec(input Input) (presenter.StatusCode, error) {

	exist, err := isProductExist(input.ProductName)
	if err != nil {
		return presenter.StatusSQLError, errors.New("query product error")
	}
	if exist {
		return presenter.StatusCreatePostFailedCreateRecordName, errors.New("product name has been used")
	}

	product := new(entities.Product)
	product.ProductName = input.ProductName
	product.ProductImageURL = input.ProductImageURL
	product.Summary = input.Summary
	product.Information = input.Information
	product.ProductPrice = input.ProductPrice
	product.InventoryNumber = input.InventoryNumber
	product.MaxBuy = input.MaxBuy
	product.CompanyID = input.CompanyID
	product.ProductStatus = entities.Sell

	if err := store.DB.Create(product).Error; err != nil {
		return presenter.StatusCreatePostFailedCreateRecord, errors.New("create company error")
	}

	return presenter.StatusSuccess, nil
}

func isProductExist(name string) (bool, error) {

	var count int
	err := store.DB.Model(entities.Product{}).Where("product_name = ?", name).Count(&count).Error
	if err != nil {
		return false, errors.New("query product error")
	}

	return count != 0, nil
}
