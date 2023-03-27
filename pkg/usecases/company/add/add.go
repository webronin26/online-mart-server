package add

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
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

	exist, err := isCompanyExist(input.CompanyName)
	if err != nil {
		return presenter.StatusSQLError, errors.New("query company error")
	}
	if exist {
		return presenter.StatusCreatePostFailedCreateRecordName, errors.New("name has been used")
	}

	company := new(entities.Company)
	company.CompanyName = input.CompanyName
	company.ResponsiblePerson = input.ResponsiblePerson
	company.GovernmentUniformInvoiceNumber = input.GovernmentUniformInvoiceNumber
	company.RemittanceAccount = input.RemittanceAccount
	company.OfficePhoneNumber = input.OfficePhoneNumber
	company.PersonalPhoneNumber = input.PersonalPhoneNumber
	company.OfficeAddress = input.OfficeAddress
	company.CorrespondenceAddress = input.CorrespondenceAddress
	company.DeliveryFee = input.DeliveryFee

	if err := store.DB.Create(company).Error; err != nil {
		return presenter.StatusCreatePostFailedCreateRecord, errors.New("create company error")
	}

	return presenter.StatusSuccess, nil
}

func isCompanyExist(name string) (bool, error) {

	var count int
	err := store.DB.Model(entities.Company{}).Where("company_name = ?", name).Count(&count).Error
	if err != nil {
		return false, errors.New("query company name error")
	}

	return count != 0, nil
}
