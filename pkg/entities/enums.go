package entities

// AccountStatus 帳戶狀態
type AccountStatus int8

// 狀態可選種類
const (
	InvalidStatus  AccountStatus = 1 // 未啟用
	ActiveStatus   AccountStatus = 2 // 啟用
	InActiveStatus AccountStatus = 3 // 停權
	RemoveStatus   AccountStatus = 4 // 移除
)

// IsValid 是否為有效值
func (a AccountStatus) IsValid() bool {
	return InvalidStatus <= a && a <= RemoveStatus
}

// UserKind 使用者類型
type UserKind int8

// UserKind 登入方法可選種類
const (
	AdminUser    UserKind = 1 // 管理者
	MemberUser   UserKind = 2 // 用戶
	RetailerUser UserKind = 3 // 廠商
)

// IsValid 是否為有效值
func (u UserKind) IsValid() bool {
	return AdminUser <= u && u <= RetailerUser
}

// MemberLoginType 登入方法
type MemberLoginType int8

// AccountLoginType 登入方法可選種類
const (
	System   MemberLoginType = 1
	facebook MemberLoginType = 2
	google   MemberLoginType = 3
)

// IsValid 是否為有效值
func (m MemberLoginType) IsValid() bool {
	return System <= m && m <= google
}

// RetailerType 廠商使用者的類型
type RetailerType int8

// 使用者的類型可選種類
const (
	RetailerAdmin       RetailerType = 1 // 廠商管理者
	RetailerEditor      RetailerType = 2 // 廠商小編
	RetailerSales       RetailerType = 3 // 廠商業務
	RetailerAccountants RetailerType = 4 // 廠商會計
)

// IsValid 是否為有效值
func (r RetailerType) IsValid() bool {
	return RetailerAdmin <= r && r <= RetailerAccountants
}

// OrderStatus 訂單的付款狀況
type OrderStatus int8

// 訂單的付款狀況可選種類
const (
	Paid     OrderStatus = 1 // 付款
	Unpaid   OrderStatus = 2 // 未付款
	Failed   OrderStatus = 3 // 失敗
	Overtime OrderStatus = 4 // 付款超時
)

// IsValid 是否為有效值
func (o OrderStatus) IsValid() bool {
	return Paid <= o && o <= Overtime
}

// ProductStatus 商品上架狀況
type ProductStatus int8

// 商品上架狀況可選種類
const (
	Sell   ProductStatus = 1 // 販售
	UnSell ProductStatus = 2 // 未販售
	Remove ProductStatus = 3 // 下架
)

// IsValid 是否為有效值
func (p ProductStatus) IsValid() bool {
	return Sell <= p && p <= Remove
}

// ProductStatus 商品上架狀況
type CartStatus int8

// 商品上架狀況可選種類
const (
	Check   CartStatus = 1 // 已結帳
	UnCheck CartStatus = 2 // 未結帳
)

// IsValid 是否為有效值
func (c CartStatus) IsValid() bool {
	return Check <= c && c <= UnCheck
}
