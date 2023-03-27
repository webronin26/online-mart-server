package index

import "github.com/webronin26/online-mart-server/pkg/presenter"

type Output struct {
	HotSales         hotSales         `json:"hot_sales"`
	Latest           latest           `json:"latest"`
	DailyRecommended dailyRecommended `json:"daily_recommended"`
	MayLike          mayLike          `json:"may_like"`
}

type hotSales struct {
	FirstID    uint64 `json:"first_id"`
	FirstPic   string `json:"first_pic"`
	FirstText  string `json:"first_text"`
	FirstPrice string `json:"first_price"`

	SecondID    uint64 `json:"second_id"`
	SecondPic   string `json:"second_pic"`
	SecondText  string `json:"second_text"`
	SecondPrice string `json:"second_price"`

	ThirdID    uint64 `json:"third_id"`
	ThirdPic   string `json:"third_pic"`
	ThirdText  string `json:"third_text"`
	ThirdPrice string `json:"third_price"`

	FourID    uint64 `json:"four_id"`
	FourPic   string `json:"four_pic"`
	FourText  string `json:"four_text"`
	FourPrice string `json:"four_price"`
}

type latest struct {
	FirstID    uint64 `json:"first_id"`
	FirstPic   string `json:"first_pic"`
	FirstText  string `json:"first_text"`
	FirstPrice string `json:"first_price"`

	SecondID    uint64 `json:"second_id"`
	SecondPic   string `json:"second_pic"`
	SecondText  string `json:"second_text"`
	SecondPrice string `json:"second_price"`

	ThirdID    uint64 `json:"third_id"`
	ThirdPic   string `json:"third_pic"`
	ThirdText  string `json:"third_text"`
	ThirdPrice string `json:"third_price"`

	FourID    uint64 `json:"four_id"`
	FourPic   string `json:"four_pic"`
	FourText  string `json:"four_text"`
	FourPrice string `json:"four_price"`
}

type dailyRecommended struct {
	FirstID    uint64 `json:"first_id"`
	FirstPic   string `json:"first_pic"`
	FirstText  string `json:"first_text"`
	FirstPrice string `json:"first_price"`

	SecondID    uint64 `json:"second_id"`
	SecondPic   string `json:"second_pic"`
	SecondText  string `json:"second_text"`
	SecondPrice string `json:"second_price"`

	ThirdID    uint64 `json:"third_id"`
	ThirdPic   string `json:"third_pic"`
	ThirdText  string `json:"third_text"`
	ThirdPrice string `json:"third_price"`

	FourID    uint64 `json:"four_id"`
	FourPic   string `json:"four_pic"`
	FourText  string `json:"four_text"`
	FourPrice string `json:"four_price"`
}

type mayLike struct {
	FirstID    uint64 `json:"first_id"`
	FirstPic   string `json:"first_pic"`
	FirstText  string `json:"first_text"`
	FirstPrice string `json:"first_price"`

	SecondID    uint64 `json:"second_id"`
	SecondPic   string `json:"second_pic"`
	SecondText  string `json:"second_text"`
	SecondPrice string `json:"second_price"`

	ThirdID    uint64 `json:"third_id"`
	ThirdPic   string `json:"third_pic"`
	ThirdText  string `json:"third_text"`
	ThirdPrice string `json:"third_price"`

	FourID    uint64 `json:"four_id"`
	FourPic   string `json:"four_pic"`
	FourText  string `json:"four_text"`
	FourPrice string `json:"four_price"`
}

// 返回主頁需要的資訊
func Exec() (Output, presenter.StatusCode, error) {

	var output Output

	output.HotSales.FirstID = 1
	output.HotSales.FirstPic = "KQ0ldVK.jpg"
	output.HotSales.FirstText = "Jordan Flight MVP 男款 Statement 外套"
	output.HotSales.FirstPrice = "3500"

	output.HotSales.SecondID = 2
	output.HotSales.SecondPic = "80ksliW.jpg"
	output.HotSales.SecondText = "Acer VG280K 28型IPS 4K高解析HDR電競電腦螢幕"
	output.HotSales.SecondPrice = "6888"

	output.HotSales.ThirdID = 3
	output.HotSales.ThirdPic = "d9mkOeo.jpg"
	output.HotSales.ThirdText = "韓版秋冬內搭短洋裝收腰長袖洋裝顯瘦泡泡袖連身裙"
	output.HotSales.ThirdPrice = "580"

	output.HotSales.FourID = 4
	output.HotSales.FourPic = "cku2lUj.jpg"
	output.HotSales.FourText = "鈦鋼鍊條/白鋼鍊條/防過敏/不生鏽/不生銹/西德鋼手鍊"
	output.HotSales.FourPrice = "799"

	output.Latest.FirstID = 3
	output.Latest.FirstPic = "d9mkOeo.jpg"
	output.Latest.FirstText = "韓版秋冬內搭短洋裝收腰長袖洋裝顯瘦泡泡袖連身裙"
	output.Latest.FirstPrice = "580"

	output.Latest.SecondID = 4
	output.Latest.SecondPic = "cku2lUj.jpg"
	output.Latest.SecondText = "鈦鋼鍊條/白鋼鍊條/防過敏/不生鏽/不生銹/西德鋼手鍊"
	output.Latest.SecondPrice = "799"

	output.Latest.ThirdID = 5
	output.Latest.ThirdPic = "ThWonHF.jpg"
	output.Latest.ThirdText = "canon 4K數位相機 4800萬高清像素"
	output.Latest.ThirdPrice = "12800"

	output.Latest.FourID = 1
	output.Latest.FourPic = "KQ0ldVK.jpg"
	output.Latest.FourText = "Jordan Flight MVP 男款 Statement 外套"
	output.Latest.FourPrice = "3500"

	output.DailyRecommended.FirstID = 5
	output.DailyRecommended.FirstPic = ""
	output.DailyRecommended.FirstText = "canon 4K數位相機 4800萬高清像素"
	output.DailyRecommended.FirstPrice = "12800"

	output.DailyRecommended.SecondID = 1
	output.DailyRecommended.SecondPic = "KQ0ldVK.jpg"
	output.DailyRecommended.SecondText = "Jordan Flight MVP 男款 Statement 外套"
	output.DailyRecommended.SecondPrice = "3500"

	output.DailyRecommended.ThirdID = 2
	output.DailyRecommended.ThirdPic = "80ksliW.jpg"
	output.DailyRecommended.ThirdText = "Acer VG280K 28型IPS 4K高解析HDR電競電腦螢幕"
	output.DailyRecommended.ThirdPrice = "6888"

	output.DailyRecommended.FourID = 3
	output.DailyRecommended.FourPic = "d9mkOeo.jpg"
	output.DailyRecommended.FourText = "韓版秋冬內搭短洋裝收腰長袖洋裝顯瘦泡泡袖連身裙"
	output.DailyRecommended.FourPrice = "580"

	output.MayLike.FirstID = 2
	output.MayLike.FirstPic = "80ksliW.jpg"
	output.MayLike.FirstText = "Acer VG280K 28型IPS 4K高解析HDR電競電腦螢幕"
	output.MayLike.FirstPrice = "6888"

	output.MayLike.SecondID = 3
	output.MayLike.SecondPic = "d9mkOeo.jpg"
	output.MayLike.SecondText = "韓版秋冬內搭短洋裝收腰長袖洋裝顯瘦泡泡袖連身裙"
	output.MayLike.SecondPrice = "580"

	output.MayLike.ThirdID = 4
	output.MayLike.ThirdPic = "cku2lUj.jpg"
	output.MayLike.ThirdText = "鈦鋼鍊條/白鋼鍊條/防過敏/不生鏽/不生銹/西德鋼手鍊"
	output.MayLike.ThirdPrice = "799"

	output.MayLike.FourID = 5
	output.MayLike.FourPic = "ThWonHF.jpg"
	output.MayLike.FourText = "canon 4K數位相機 4800萬高清像素"
	output.MayLike.FourPrice = "12800"

	return output, presenter.StatusSuccess, nil
}
