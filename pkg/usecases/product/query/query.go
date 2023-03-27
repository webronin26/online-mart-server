package query

import (
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	CompanyID int64
}

type Output struct {
	Data  Data
	Count int
}

type Data struct {
	ID            uint64           `json:"company_id"`
	CompanyName   string           `json:"company_name"`
	OfficeAddress string           `json:"company_address"`
	Products      []*simpleProduct `json:"products"`
}

type simpleProduct struct {
	ID              int32  `json:"id"`
	ProductName     string `json:"name"`
	ProductPrice    int    `json:"price"`
	InventoryNumber int64  `json:"inventory_number"`
}

func Exec(input Input) (Output, error) {

	var company entities.Company
	query := store.DB.Model(entities.Company{}).
		Where("id = ?", input.CompanyID).
		Scan(&company)
	if err := query.Error; err != nil {
		return Output{}, errors.New("company query error")
	}

	var output Output

	output.Data.ID = company.ID
	output.Data.CompanyName = company.CompanyName
	output.Data.OfficeAddress = company.OfficeAddress

	productQuery := store.DB.Model(entities.Product{}).Where("company_id = ?", input.CompanyID)
	if err := productQuery.Count(&output.Count).Error; err != nil {
		return Output{}, errors.New("query company product count error")
	}
	if output.Count == 0 {
		return output, nil
	}

	output.Data.Products = make([]*simpleProduct, output.Count)

	if err := productQuery.Scan(&output.Data.Products).Error; err != nil {
		return output, errors.New("query company product error")
	}

	return output, nil
}
