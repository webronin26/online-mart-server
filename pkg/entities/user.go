package entities

type User struct {
	ID            uint64        `gorm:"primary_key;"`
	Account       string        `gorm:"unique"`   // 帳戶帳號
	Password      string        `gorm:"not null"` // 帳戶密碼
	UserName      string        // 帳戶名稱
	AccountStatus AccountStatus `gorm:"not null"` // 帳戶狀態
	UserKind      UserKind      `gorm:"not null"` // 帳戶類型
	Token         string        `gorm:"unique"`   // 帳戶 Token
}
