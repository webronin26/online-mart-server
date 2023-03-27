package search

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	Query string
}

type Output struct {
	Products []*simpleProduct
	Count    int
}

type simpleProduct struct {
	ID              int32  `json:"id"`
	ProductName     string `json:"name"`
	ProductPrice    int    `json:"price"`
	ProductImageURL string `json:"image_url"`
}

func Exec(input Input) (Output, presenter.StatusCode, error) {

	var output Output

	query := store.DB.Model(entities.Product{}).Where("product_name LIKE ?", "%"+input.Query+"%")
	if err := query.Count(&output.Count).Error; err != nil {
		return output, presenter.StatusServerError, errors.New("search count error")
	}

	output.Products = make([]*simpleProduct, output.Count)
	if err := query.Scan(&output.Products).Error; err != nil {
		return output, presenter.StatusSQLErrorScanFailed, errors.New("search scan error")
	}

	return output, presenter.StatusSuccess, nil
}
