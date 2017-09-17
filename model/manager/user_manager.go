package manager

import "github.com/server-forecaster/model/entity"

type UserManager interface {
	AddUser(user *entity.User) bool

	GetUserByAlias(alias string) *entity.User
}

type DefaultUserManager struct {
}

func (mng DefaultUserManager) AddUser(user *entity.User) bool {
	return true
}

func (mng DefaultUserManager) GetUserByAlias(alias string) *entity.User {
	print("The alias " + alias)
	if alias == "notExistent" {
		return nil
	}
	return &entity.User{Alias: "userAlias"}
}

func Create() UserManager {
	return DefaultUserManager{}
}
