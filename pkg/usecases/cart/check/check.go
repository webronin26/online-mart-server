package check

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	MemberID      uint64
	MemberAddress string
}

type cartItem struct {
	ID              uint64
	Number          int64
	ProductPrice    float64
	TotalPrice      float64
	ProductImageURL string
}

func Exec(input Input) (presenter.StatusCode, error) {

	var cart entities.Cart
	cartQuery := store.DB.Model(entities.Cart{}).
		Where("member_id = ?", input.MemberID).
		Where("cart_status = ?", entities.UnCheck).
		First(&cart)
	if cartQuery.RecordNotFound() {
		return presenter.StatusRecordNotFound, errors.New("query cart record not found")
	} else if err := cartQuery.Error; err != nil {
		return presenter.StatusSQLError, errors.New("query cart error")
	}

	if cart.CompanyID == 0 || cart.ProductList == "" {
		return presenter.StatusRecordNotFound, errors.New("query cart record not found")
	}

	var currentProductList []cartItem
	if unmarshalErr := json.Unmarshal([]byte(cart.ProductList), &currentProductList); unmarshalErr != nil {
		return presenter.StatusCreatePostFailedConvert, errors.New("cart item json convert error")
	}

	tx := store.DB.Begin()

	for _, item := range currentProductList {
		var product entities.Product
		productQuery := store.DB.Model(entities.Product{}).
			Where("id = ?", item.ID).
			Where("product_status = ?", entities.Sell).
			First(&product)
		if productQuery.RecordNotFound() {
			return presenter.StatusRecordNotFound, errors.New("cart product not found")
		}
		if err := productQuery.Error; err != nil {
			return presenter.StatusSQLError, errors.New("cart product query error")
		}

		newInventoryNumber := product.InventoryNumber - item.Number
		if newInventoryNumber < 0 {
			return presenter.StatusUpdatePostFailedInventoryNotEnough, errors.New("product inventory number not enough")
		}

		updateProduct := tx.Model(entities.Product{}).
			Where("id = ?", item.ID).
			Update("item_number", newInventoryNumber)
		if err := updateProduct.Error; err != nil {
			tx.Rollback()
			return presenter.StatusUpdatePostFailed, errors.New("cart check : product inventory number update error")
		}
	}

	var order entities.Order
	order.OrderNumber = uuid.New().String()
	order.OrderAddress = input.MemberAddress
	order.ProductList = cart.ProductList
	order.TotalPrice = cart.TotalPrice
	order.OrderStatus = entities.Paid
	order.PaidTime = time.Now()
	order.MemberID = cart.MemberID
	order.CompanyID = cart.CompanyID
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return presenter.StatusCreatePostFailedCreateRecord, errors.New("create order error")
	}
	if err := store.DB.Model(&cart).Update("cart_status", entities.Check).Error; err != nil {
		tx.Rollback()
		return presenter.StatusUpdatePostFailed, errors.New("update cart status error")
	}

	tx.Commit()

	return presenter.StatusSuccess, nil
}
