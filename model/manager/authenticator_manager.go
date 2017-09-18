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
		identifier, utils.MD5(password)).First(&user)
	if user.Alias == "" {
		return nil
	} else {
		return &user
	}
}

func CreateAuthenticatorManager() AuthenticatorManager {
	return AuthenticatorManager{BaseManager: Create()}
}
