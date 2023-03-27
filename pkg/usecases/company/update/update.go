package update

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	ID uint64 // retailer 的 id

	CompanyName                    string
	ResponsiblePerson              string
	GovernmentUniformInvoiceNumber string
	RemittanceAccount              string
	OfficePhoneNumber              string
	PersonalPhoneNumber            string
	OfficeAddress                  string
	CorrespondenceAddress          string
	DeliveryFee                    float64
}

func Exec(input Input) (presenter.StatusCode, error) {

	var retailer entities.Retailer
	query := store.DB.Model(entities.Retailer{}).
		Where("user_id = ?", input.ID).
		First(&retailer)
	if query.RecordNotFound() {
		return presenter.StatusRecordNotFound, errors.New("retailer record not found error")
	}
	if query.Error != nil {
		return presenter.StatusSQLError, errors.New("query error : " + query.Error.Error())
	}
	if retailer.UserID == 0 {
		return presenter.StatusSQLError, errors.New("retailer user_id == 0")
	}

	updateParam := make(map[string]interface{})

	if input.CompanyName != "" {
		updateParam["company_name"] = input.CompanyName
	}
	if input.ResponsiblePerson != "" {
		updateParam["responsible_person"] = input.ResponsiblePerson
	}
	if input.GovernmentUniformInvoiceNumber != "" {
		updateParam["government_uniform_invoice_number"] = input.GovernmentUniformInvoiceNumber
	}
	if input.RemittanceAccount != "" {
		updateParam["remittance_account"] = input.RemittanceAccount
	}
	if input.OfficePhoneNumber != "" {
		updateParam["office_phone_number"] = input.OfficePhoneNumber
	}
	if input.PersonalPhoneNumber != "" {
		updateParam["personal_phone_number"] = input.PersonalPhoneNumber
	}
	if input.OfficeAddress != "" {
		updateParam["office_address"] = input.OfficeAddress
	}
	if input.CorrespondenceAddress != "" {
		updateParam["correspondence_address"] = input.CorrespondenceAddress
	}
	if input.DeliveryFee != 0 {
		updateParam["delivery_fee"] = input.DeliveryFee
	}

	err := store.DB.Model(entities.Company{}).
		Where("id = ?", retailer.CompanyID).
		Updates(updateParam).
		Error
	if err != nil {
		return presenter.StatusSQLError, errors.New("update error : " + err.Error())
	}

	return presenter.StatusSuccess, nil
}
