package views

import (
	"github.com/server-forecaster/model/entity"
	"github.com/server-forecaster/model/manager"
)

func GetUserByToken(token string) *entity.User {
	return manager.CreateUserManager().GetUserByToken(token)
}
