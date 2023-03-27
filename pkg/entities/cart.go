package entities

type Cart struct {
	ID          uint64     `gorm:"primary_key;"`
	MemberID    uint64     `gorm:"not null"` // 用戶 ID
	CartStatus  CartStatus `gorm:"not null"` // 是否結帳了
	ProductList string     // 訂貨商品+數量 的列表
	TotalPrice  float64    // 訂貨總額
	CompanyID   uint64     // 廠商的 ID
}
