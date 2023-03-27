package presenter

// API 統一回傳結構
type Result struct {
	Data  interface{} `json:"data"`
	Code  StatusCode  `json:"code"` // 回應狀態碼
	Count int         `json:"count"`
}

type StatusCode uint16

// 回應狀態碼
const (
	StatusSuccess = 2001 // 成功

	StatusParamError                         = 4001  // 參數有誤
	StatusParamValidateFailed                = 4002  // 參數數值有誤
	StatusBindFailed                         = 4003  // 參數綁定失敗
	StatusRecordNotFound                     = 4004  // 找不到紀錄
	StatusCreatePostFailed                   = 4005  // Post 建立失敗
	StatusCreatePostFailedInventoryNumber    = 40051 // Post 建立失敗，庫存數量問題
	StatusCreatePostFailedCreateRecord       = 40052 // Post 建立失敗，建立新紀錄問題
	StatusCreatePostFailedProductCompany     = 40053 // Post 建立失敗，商品廠商異常
	StatusCreatePostFailedConvert            = 40054 // Post 建立失敗，資料轉型失敗
	StatusCreatePostFailedCreateRecordName   = 40055 // Post 建立失敗，建立新紀錄名稱問題
	StatusUpdatePostFailed                   = 4006  // Post 更新失敗
	StatusUpdatePostFailedInventoryNotEnough = 40061 // Post 更新失敗，庫存不足

	StatusTokenError   = 4011 // Token 有誤
	StatusAccountError = 4041 // 帳密認證有誤

	StatusServerError        = 5001  // 伺服器錯誤
	StatusSQLError           = 5002  // 資料庫錯誤
	StatusSQLErrorScanFailed = 50021 // 資料庫錯誤，轉型錯誤
)
