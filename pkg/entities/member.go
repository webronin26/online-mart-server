package entities

type Member struct {
	ID                  uint64          `gorm:"primary_key;"`
	PhoneNumber         string          // 手機號碼
	MemberAddress       string          // 地址
	MemberAvatarURL     string          // 大頭像
	MemberEmail         string          // 帳戶 email
	PhoneCertify        bool            // 手機是否驗證過了
	MemberLoginType     MemberLoginType `gorm:"not null"` // 登入類型 1.系統 2.fb 3. google
	FacebookID          string          `gorm:"unique"`   // fb 帳號
	FacebookAccessToken string          // fb token
	FacebookEmail       string          // fb email
	FacebookAvatar      string          // fb 大頭像
	GoogleID            string          `gorm:"unique"` // google 帳號
	GoogleAccessToken   string          // google token
	GoogleEmail         string          // google email
	GoogleAvatar        string          // fb 大頭像
	GoogleName          string          // google 名稱
	UserID              uint64          `gorm:"not null"` // 對應的使用者 id
}
