package manager

import (
	"github.com/server-forecaster/utils"
	"github.com/server-forecaster/model/entity"
)

type AuthenticatorManager struct {
	BaseManager
}

func (manager AuthenticatorManager) Login(identifier string, password string) *entity.User {
	user := entity.User{}
	manager.DB.Where("(Email = ? OR Alias = ?) AND Password = ?",
		identifier, utils.MD5(password)).FirstOrInit(&user, entity.User{Alias: "NULL"})
	if user.Alias == "NULL" {
		return nil
	} else {
		return &user
	}
}

func CreateAuthenticatorManager() AuthenticatorManager {
	return AuthenticatorManager{BaseManager: Create()}
}
