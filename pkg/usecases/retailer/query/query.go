package query

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	UserId uint64
}

type Output struct {
	Retailer simpleRetailer
}

type simpleRetailer struct {
	RetailerType entities.RetailerType `json:"retailer_type"`
	CompanyID    uint64                `json:"company_id"`
}

func Exec(input Input) (Output, error) {

	var retailer simpleRetailer
	query := store.DB.Model(entities.User{}).
		Joins("JOIN retailers ON retailers.user_id = users.id").
		Where("id = ?", input.UserId).
		First(&retailer)
	if query.RecordNotFound() {
		return Output{}, errors.New("query retailer record not found")
	}
	if err := query.Error; err != nil {
		return Output{}, errors.New("query retailer error")
	}

	var output Output
	output.Retailer.CompanyID = retailer.CompanyID
	output.Retailer.RetailerType = retailer.RetailerType

	return output, nil
}
