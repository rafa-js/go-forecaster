package manager

import (
	"github.com/server-forecaster/utils"
	"github.com/server-forecaster/model/entity"
	"time"
	"fmt"
)

type AuthenticatorManager struct {
	BaseManager
}

func (manager AuthenticatorManager) Login(identifier string, password string) *entity.User {
	user := entity.User{}
	manager.DB.Where("(email = ? OR alias = ?) AND password = ?",
		identifier, identifier, utils.MD5(password)).First(&user)
	if user.Alias == "" {
		return nil
	} else {
		user.AuthToken = utils.MD5(fmt.Sprint(time.Now()))
		manager.DB.Save(user)
		return &user
	}
}

func CreateAuthenticatorManager() AuthenticatorManager {
	return AuthenticatorManager{BaseManager: Create()}
}
