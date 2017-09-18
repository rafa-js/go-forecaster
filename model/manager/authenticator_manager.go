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
	var count int
	manager.DB.Where("(Email = ? OR Alias = ?) AND Password = ?",
		identifier, utils.MD5(password)).Count(&count).First(&user)
	if count == 0 {
		return nil
	} else {
		return &user
	}
}

func CreateAuthenticatorManager() AuthenticatorManager {
	return AuthenticatorManager{BaseManager: Create()}
}
