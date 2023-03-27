package query

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Output struct {
	Companies []*simpleCompany
	Count     int
}

type simpleCompany struct {
	ID                             uint64 `json:"id"`
	CompanyName                    string `json:"name"`
	ResponsiblePerson              string `json:"responsible_person"`
	GovernmentUniformInvoiceNumber string `json:"invoice"`
	OfficePhoneNumber              string `json:"phone"`
	OfficeAddress                  string `json:"office_address"`
}

func Exec() (Output, presenter.StatusCode, error) {

	var output Output

	query := store.DB.Model(entities.Company{}).
		Select("companies.id, companies.company_name, companies.responsible_person, companies.government_uniform_invoice_number, companies.office_phone_number, companies.office_address")
	if err := query.Count(&output.Count).Error; err != nil {
		return output, presenter.StatusSQLError, errors.New("company query count error")
	}

	output.Companies = make([]*simpleCompany, output.Count)

	if err := query.Scan(&output.Companies).Error; err != nil {
		return output, presenter.StatusSQLErrorScanFailed, errors.New("company query error")
	}

	return output, presenter.StatusSuccess, nil
}
