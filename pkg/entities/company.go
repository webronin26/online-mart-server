package entities

type Company struct {
	ID                             uint64  `gorm:"primary_key;"`
	CompanyName                    string  `gorm:"not null"` // 公司名稱
	ResponsiblePerson              string  // 公司負責人
	GovernmentUniformInvoiceNumber string  // 統一編號
	RemittanceAccount              string  // 分局號 + 號碼
	OfficePhoneNumber              string  // 公司電話
	PersonalPhoneNumber            string  // 聯絡人手機
	OfficeAddress                  string  // 公司地址
	CorrespondenceAddress          string  // 通訊地址
	DeliveryFee                    float64 // 運送費用
}
