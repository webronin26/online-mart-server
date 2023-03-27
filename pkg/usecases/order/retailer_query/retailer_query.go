package retailer_query

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	CompanyID uint64
}

type Output struct {
	Count  int
	Orders []*simpleOrder
}

type simpleOrder struct {
	ID          uint64  `json:"id"`
	OrderNumber string  `json:"order_number"`
	ProductList string  `json:"product_list"`
	TotalPrice  float64 `json:"total_price"`
}

func Exec(input Input) (Output, presenter.StatusCode, error) {

	var output Output
	query := store.DB.Model(entities.Order{}).Where("company_id = ?", input.CompanyID)
	if err := query.Count(&output.Count).Error; err != nil {
		return Output{}, presenter.StatusSQLError, errors.New("query order count error")
	}

	output.Orders = make([]*simpleOrder, output.Count)
	if output.Count == 0 {
		return output, presenter.StatusSuccess, nil
	}
	if err := query.Scan(&output.Orders).Error; err != nil {
		return output, presenter.StatusSQLErrorScanFailed, errors.New("query company product error")
	}

	return output, presenter.StatusSuccess, nil
}
