package query_detail

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	ProductID uint64
}

type Output struct {
	Product simpleProduct `json:"product"`
}

type simpleProduct struct {
	ID              uint64  `json:"id"`
	ProductName     string  `json:"name"`
	ProductImageURL string  `json:"image_url"`
	Summary         string  `json:"summary"`
	Information     string  `json:"information"`
	ProductPrice    float64 `json:"price"`
	InventoryNumber int64   `json:"inventory_number"`
	MaxBuy          int64   `json:"max_buy"`

	CompanyID     uint64 `json:"company_id"`
	CompanyName   string `json:"company_name"`
	OfficeAddress string `json:"company_address"`
}

func Exec(input Input) (Output, presenter.StatusCode, error) {

	var simpleProduct simpleProduct
	query := store.DB.Model(entities.Product{}).
		Select("products.id, products.product_name, products.product_image_url, products.summary, products.information, products.product_price, products.inventory_number, products.max_buy, products.company_id, companies.company_name, companies.office_address").
		Joins("JOIN companies ON products.company_id = companies.id").
		Where("products.id = ? AND products.product_status = ?", input.ProductID, entities.Sell).
		Scan(&simpleProduct)
	if query.RecordNotFound() {
		return Output{}, presenter.StatusRecordNotFound, errors.New("query product record not found")
	}
	if err := query.Error; err != nil {
		return Output{}, presenter.StatusServerError, errors.New("product query error")
	}

	var output Output
	output.Product = simpleProduct

	return output, presenter.StatusSuccess, nil
}
