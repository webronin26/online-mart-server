package query_detail

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	CompanyID int64
}

type Output struct {
	Company company `json:"company"`
}

type company struct {
	ID                             uint64  `json:"id"`
	CompanyName                    string  `json:"name"`
	ResponsiblePerson              string  `json:"responsible_person"`
	GovernmentUniformInvoiceNumber string  `json:"invoice"`
	RemittanceAccount              string  `json:"remittance_account"`
	OfficePhoneNumber              string  `json:"office_phone"`
	PersonalPhoneNumber            string  `json:"personal_phone"`
	OfficeAddress                  string  `json:"office_address"`
	CorrespondenceAddress          string  `json:"correspondence_address"`
	DeliveryFee                    float64 `json:"delivery_fee"`
}

func Exec(input Input) (Output, presenter.StatusCode, error) {

	var c company
	query := store.DB.Model(entities.Company{}).
		Where("id = ?", input.CompanyID).
		Scan(&c)
	if err := query.Error; err != nil {
		return Output{}, presenter.StatusSQLError, errors.New("company query detail error")
	}

	var output Output
	output.Company = c

	return output, presenter.StatusSuccess, nil
}
