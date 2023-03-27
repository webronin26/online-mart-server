package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/company/add"
)

type AddCompanyParam struct {
	CompanyName                    string  `json:"company_name"`
	ResponsiblePerson              string  `json:"responsible_person"`
	GovernmentUniformInvoiceNumber string  `json:"invoice_number"`
	RemittanceAccount              string  `json:"remittance_account"`
	OfficePhoneNumber              string  `json:"office_phone_number"`
	PersonalPhoneNumber            string  `json:"personal_phone_number"`
	OfficeAddress                  string  `json:"office_address"`
	CorrespondenceAddress          string  `json:"correspondence_address"`
	DeliveryFee                    float64 `json:"delivery_fee"`
}

func AdminAddCompany(ctx echo.Context) error {

	var result presenter.Result

	var param AddCompanyParam
	if err := ctx.Bind(&param); err != nil {
		result.Code = presenter.StatusParamError
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input add.Input
	input.CompanyName = param.CompanyName
	input.ResponsiblePerson = param.ResponsiblePerson
	input.GovernmentUniformInvoiceNumber = param.GovernmentUniformInvoiceNumber
	input.RemittanceAccount = param.RemittanceAccount
	input.OfficePhoneNumber = param.OfficePhoneNumber
	input.PersonalPhoneNumber = param.PersonalPhoneNumber
	input.OfficeAddress = param.OfficeAddress
	input.CorrespondenceAddress = param.CorrespondenceAddress
	input.DeliveryFee = param.DeliveryFee

	serverStatus, addErr := add.Exec(input)
	if addErr != nil {
		result.Code = serverStatus
		result.Data = addErr.Error()
		return ctx.JSON(http.StatusInternalServerError, result)
	}

	result.Code = serverStatus

	return ctx.JSON(http.StatusOK, result)
}
