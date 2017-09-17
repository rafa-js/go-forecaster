package manager

import "github.com/server-forecaster/model/entity"

type UserManager interface {
	AddUser(user *entity.User)

	GetUserByAlias(alias string) *entity.User
}

type DefaultUserManager struct {
}

func (mng DefaultUserManager) AddUser(user *entity.User) {

}

func (mng DefaultUserManager) GetUserByAlias(alias string) *entity.User {
	return &entity.User{Alias: "userAlias"}
}

func Create() UserManager {
	return DefaultUserManager{}
}
