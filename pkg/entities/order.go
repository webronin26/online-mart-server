package entities

import "time"

type Order struct {
	ID           uint64      `gorm:"primary_key;"`
	OrderNumber  string      `gorm:"not null"`            // 訂單編號
	OrderAddress string      `gorm:"not null"`            // 收件人地址
	ProductList  string      `gorm:"not null;type:JSON;"` // 訂貨商品+數量 的列表
	TotalPrice   float64     `gorm:"not null"`            // 訂貨總額
	OrderStatus  OrderStatus `gorm:"not null"`            // 付款狀況
	PaidTime     time.Time   `gorm:"type:datetime"`       // 付款時間
	MemberID     uint64      `gorm:"not null"`            // 用戶 ID
	CompanyID    uint64      `gorm:"not null"`            // 廠商 ID
}
