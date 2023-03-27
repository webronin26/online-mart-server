package query_detail

import (
	"errors"
	"time"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	MemberID uint64
	OrderID  uint64
}

type Output struct {
	Order Order `json:"order"`
}

type Order struct {
	ID           uint64               `json:"id"`
	OrderNumber  string               `json:"order_number"`
	OrderAddress string               `json:"order_address"`
	ProductList  string               `json:"product_list"`
	TotalPrice   float64              `json:"total_price"`
	OrderStatus  entities.OrderStatus `json:"order_status"`
	PaidTime     time.Time            `json:"paid_time"`
	CompanyID    uint64               `json:"company_id"`
}

func Exec(input Input) (Output, presenter.StatusCode, error) {

	var order Order
	query := store.DB.Model(entities.Order{}).
		Where("id = ?", input.OrderID).
		Where("member_id = ?", input.MemberID).
		First(&order)
	if query.RecordNotFound() {
		return Output{}, presenter.StatusRecordNotFound, errors.New("query order record not found")
	}
	if err := query.Error; err != nil {
		return Output{}, presenter.StatusSQLError, errors.New("query order error")
	}

	var output Output
	output.Order = order

	return output, presenter.StatusSuccess, nil
}
