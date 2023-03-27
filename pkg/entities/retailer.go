package entities

type Retailer struct {
	ID           uint64       `gorm:"primary_key;"`
	RetailerType RetailerType `gorm:"not null"` // 廠商使用者的類型
	UserID       uint64       `gorm:"not null"` // 對應的使用者 id
	CompanyID    uint64       `gorm:"not null"` // 廠商行號為何
}
