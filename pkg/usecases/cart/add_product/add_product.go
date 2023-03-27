package add_product

import (
	"encoding/json"
	"errors"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	MemberID  uint64
	ProductID uint64
	Number    int64
}

type Output struct {
	SimpleCart simpleCart
}

type simpleCart struct {
	ID          uint64  `json:"id"`
	ProductList string  `json:"product_list"`
	TotalPrice  float64 `json:"total_price"`
}

type cartItem struct {
	ID              uint64  `json:"id"`
	ProductName     string  `json:"name"`
	Number          int64   `json:"number"`
	ProductPrice    float64 `json:"price"`
	TotalPrice      float64 `json:"total_price"`
	ProductImageURL string  `json:"image_url"`
}

func Exec(input Input) (Output, presenter.StatusCode, error) {

	var product entities.Product
	productQuery := store.DB.Model(entities.Product{}).
		Where("id = ?", input.ProductID).
		Where("product_status = ?", entities.Sell).
		First(&product)
	if productQuery.RecordNotFound() {
		return Output{}, presenter.StatusRecordNotFound, nil
	}
	if err := productQuery.Error; err != nil {
		return Output{}, presenter.StatusSQLError, nil
	}
	if input.Number > product.MaxBuy {
		input.Number = product.MaxBuy
	}
	if input.Number > product.InventoryNumber {
		return Output{}, presenter.StatusCreatePostFailedInventoryNumber, nil
	}

	var output Output
	var currentProductList []cartItem

	var cart entities.Cart
	cartQuery := store.DB.Model(entities.Cart{}).
		Where("member_id = ?", input.MemberID).
		Where("cart_status = ?", entities.UnCheck).
		First(&cart)
	if cartQuery.RecordNotFound() {

		cartItem := cartItem{
			ID:              input.ProductID,
			ProductName:     product.ProductName,
			Number:          input.Number,
			ProductPrice:    product.ProductPrice,
			TotalPrice:      product.ProductPrice * float64(input.Number),
			ProductImageURL: product.ProductImageURL,
		}

		newProductList := append(currentProductList, cartItem)
		cartItemByteArray, _ := json.Marshal(newProductList)

		cart.MemberID = input.MemberID
		cart.CartStatus = entities.UnCheck
		cart.TotalPrice = product.ProductPrice * float64(input.Number)
		cart.CompanyID = product.CompanyID
		cart.ProductList = string(cartItemByteArray)

		if err := store.DB.Create(&cart).Error; err != nil {
			return Output{}, presenter.StatusCreatePostFailedCreateRecord, errors.New("create cart error")
		}

		output.SimpleCart.ID = cart.ID
		output.SimpleCart.ProductList = cart.ProductList
		output.SimpleCart.TotalPrice = cart.TotalPrice

		return output, presenter.StatusSuccess, nil

	} else if err := cartQuery.Error; err != nil {
		return Output{}, presenter.StatusSQLError, errors.New("query cart error")
	}

	// 目前列表沒有商品，所以加入新商品
	if cart.CompanyID == 0 {
		cartItem := cartItem{
			ID:              input.ProductID,
			ProductName:     product.ProductName,
			Number:          input.Number,
			ProductPrice:    product.ProductPrice,
			TotalPrice:      product.ProductPrice * float64(input.Number),
			ProductImageURL: product.ProductImageURL,
		}

		newProductList := append(currentProductList, cartItem)
		cartItemByteArray, _ := json.Marshal(newProductList)

		cart.TotalPrice = product.ProductPrice * float64(input.Number)
		cart.CompanyID = product.CompanyID
		cart.ProductList = string(cartItemByteArray)

		err := store.DB.Model(&cart).
			Update(entities.Cart{ProductList: cart.ProductList, TotalPrice: cart.TotalPrice, CompanyID: cart.CompanyID}).Error
		if err != nil {
			return Output{}, presenter.StatusUpdatePostFailed, errors.New("update cart error")
		}

		output.SimpleCart.ID = cart.ID
		output.SimpleCart.ProductList = cart.ProductList
		output.SimpleCart.TotalPrice = cart.TotalPrice

		return output, presenter.StatusSuccess, nil
	}

	// 檢查是否目前增加的商品的廠商，是跟購物車的廠商是同一家
	if cart.CompanyID != product.CompanyID {
		return Output{}, presenter.StatusCreatePostFailedProductCompany, nil
	}

	// 解析購物車字串
	if unmarshalErr := json.Unmarshal([]byte(cart.ProductList), &currentProductList); unmarshalErr != nil {
		return Output{}, presenter.StatusCreatePostFailedConvert, errors.New("cart item json convert error")
	}

	// 如果目前這個產品有在購物車，就累加上這個產品的數量
	// 如果目前沒有這個產品在購物車，就新加入這個產品
	var productExist bool = false
	for index, item := range currentProductList {
		if item.ID == input.ProductID {

			var totalNumber = item.Number + input.Number
			if totalNumber > product.MaxBuy {
				totalNumber = product.MaxBuy
			}

			var addNumber = totalNumber - currentProductList[index].Number
			if addNumber != 0 {
				var addPrice = float64(addNumber) * currentProductList[index].ProductPrice

				currentProductList[index].Number = totalNumber
				currentProductList[index].TotalPrice = float64(totalNumber) * currentProductList[index].ProductPrice
				cart.TotalPrice = cart.TotalPrice + addPrice
			}
			productExist = true
		}
	}

	if !productExist {
		cartItem := cartItem{
			ID:              input.ProductID,
			ProductName:     product.ProductName,
			Number:          input.Number,
			ProductPrice:    product.ProductPrice,
			TotalPrice:      product.ProductPrice * float64(input.Number),
			ProductImageURL: product.ProductImageURL,
		}

		newProductList := append(currentProductList, cartItem)
		cartItemByteArray, _ := json.Marshal(newProductList)

		cart.ProductList = string(cartItemByteArray)
		cart.TotalPrice = cart.TotalPrice + cartItem.TotalPrice
	} else {
		cartItemByteArray, _ := json.Marshal(currentProductList)
		cart.ProductList = string(cartItemByteArray)
	}

	err := store.DB.Model(&cart).Update(entities.Cart{ProductList: cart.ProductList, TotalPrice: cart.TotalPrice}).Error
	if err != nil {
		return Output{}, presenter.StatusUpdatePostFailed, errors.New("update cart error")
	}

	output.SimpleCart.ID = cart.ID
	output.SimpleCart.ProductList = cart.ProductList
	output.SimpleCart.TotalPrice = cart.TotalPrice

	return output, presenter.StatusSuccess, nil
}
