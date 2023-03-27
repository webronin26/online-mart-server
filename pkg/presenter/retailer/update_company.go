package retailer

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/usecases/company/update"
)

type UpdateCompanyParam struct {
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

func RetailerUpdateCompany(ctx echo.Context) error {

	var result presenter.Result

	user := ctx.Get("user").(*entities.User)

	var param UpdateCompanyParam
	if err := ctx.Bind(&param); err != nil {
		result.Code = presenter.StatusParamError
		result.Data = err.Error()
		return ctx.JSON(http.StatusBadRequest, result)
	}

	var input update.Input
	input.ID = uint64(user.ID)
	input.CompanyName = param.CompanyName
	input.ResponsiblePerson = param.ResponsiblePerson
	input.GovernmentUniformInvoiceNumber = param.GovernmentUniformInvoiceNumber
	input.RemittanceAccount = param.RemittanceAccount
	input.OfficePhoneNumber = param.OfficePhoneNumber
	input.PersonalPhoneNumber = param.PersonalPhoneNumber
	input.OfficeAddress = param.OfficeAddress
	input.CorrespondenceAddress = param.CorrespondenceAddress
	input.DeliveryFee = param.DeliveryFee

	serverStatus, UpdateErr := update.Exec(input)
	if UpdateErr != nil {
		result.Code = serverStatus
		result.Data = UpdateErr.Error()
		return ctx.JSON(http.StatusNotFound, result)
	}

	result.Code = serverStatus

	return ctx.JSON(http.StatusOK, result)
}
