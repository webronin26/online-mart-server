package logout

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
)

type Input struct {
	UserId int64
}

func Exec(input Input) (presenter.StatusCode, error) {

	update := store.DB.Model(entities.User{}).
		Where("id = ?", input.UserId).
		Update("token", gorm.Expr("NULL"))
	if err := update.Error; err != nil {
		return presenter.StatusSQLError, errors.New("update error, logout error")
	}

	return presenter.StatusSuccess, nil
}
