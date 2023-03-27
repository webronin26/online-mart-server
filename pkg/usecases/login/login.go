package login

import (
	"errors"
	"fmt"

	"github.com/webronin26/online-mart-server/pkg/entities"
	"github.com/webronin26/online-mart-server/pkg/presenter"
	"github.com/webronin26/online-mart-server/pkg/store"
	"github.com/webronin26/online-mart-server/pkg/util"
)

type Input struct {
	Account  string
	Password string
	UserKind entities.UserKind
}

type OutPut struct {
	Token string `json:"token"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
}

func Exec(input Input) (OutPut, presenter.StatusCode, error) {

	var output OutPut
	var user entities.User

	query := store.DB.Model(entities.User{}).
		Where("account = ? AND password = ? And user_kind = ?", input.Account, input.Password, input.UserKind).
		Limit(1).
		Scan(&user)
	if queryErr := query.Error; queryErr != nil {
		return output, presenter.StatusSQLError, queryErr
	}
	if user.AccountStatus != entities.ActiveStatus {
		return output, presenter.StatusSQLError, errors.New("account not active")
	}

	token, tokenErr := util.CreateToken(input.Account, input.Password)
	if tokenErr != nil {
		return output, presenter.StatusTokenError, tokenErr
	}

	update := store.DB.Model(entities.User{}).
		Where("account = ? AND password = ?", input.Account, input.Password).
		Update("token", token)
	if updateErr := update.Error; updateErr != nil {
		fmt.Println(updateErr.Error())
		return output, presenter.StatusSQLError, errors.New("update token error")
	}

	output.ID = int64(user.ID)
	output.Name = user.UserName
	output.Token = token

	return output, presenter.StatusSuccess, nil
}
