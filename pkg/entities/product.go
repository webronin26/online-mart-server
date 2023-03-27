package entities

type Product struct {
	ID              uint64        `gorm:"primary_key;"`
	ProductName     string        `gorm:"not null"`           // 產品名稱
	ProductImageURL string        `gorm:"not null"`           // 產品圖片
	Summary         string        `gorm:"not null;type:text"` // 產品簡介
	Information     string        `gorm:"not null;type:text"` // 產品介紹
	ProductPrice    float64       `gorm:"not null"`           // 價錢
	InventoryNumber int64         `gorm:"not null"`           // 庫存數量
	MaxBuy          int64         `gorm:"not null"`           // 最大購買量
	CompanyID       uint64        `gorm:"not null"`           // 廠商的 ID
	ProductStatus   ProductStatus // 商品狀態
}
