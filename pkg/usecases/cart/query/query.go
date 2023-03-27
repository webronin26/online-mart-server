package query

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	MemberID uint64
}

type Output struct {
	Cart  simpleCart `json:"cart"`
	Count int        `json:"count"`
}

type simpleCart struct {
	ID          uint64  `json:"id"`
	ProductList string  `json:"product_list"`
	TotalPrice  float64 `json:"total_price"`
	CompanyID   uint64  `json:"company_id"`
}

func Exec(input Input) (Output, presenter.StatusCode, error) {

	var cart entities.Cart
	query := store.DB.Model(entities.Cart{}).
		Where("member_id = ?", input.MemberID).
		Where("cart_status = ?", entities.UnCheck).
		First(&cart)
	if query.RecordNotFound() {
		var output Output
		output.Count = 0
		output.Cart = simpleCart{}
		return output, presenter.StatusSuccess, nil
	}
	if err := query.Error; err != nil {
		return Output{}, presenter.StatusServerError, errors.New("query retailer error")
	}

	var output Output
	output.Count = 1
	output.Cart.ID = cart.ID
	output.Cart.CompanyID = cart.CompanyID
	output.Cart.ProductList = cart.ProductList
	output.Cart.TotalPrice = cart.TotalPrice

	return output, presenter.StatusSuccess, nil
}
