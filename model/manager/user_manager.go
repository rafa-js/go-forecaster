package manager

import (
	"github.com/server-forecaster/model/entity"
	"github.com/jinzhu/gorm"
	"github.com/server-forecaster/model"
)

type UserManager interface {
	AddUser(user *entity.User) bool

	GetUserByAlias(alias string) *entity.User

	Close()
}

type DefaultUserManager struct {
	DB gorm.DB
}

func (manager DefaultUserManager) AddUser(user *entity.User) bool {
	manager.DB.Create(&user)
	return !manager.DB.NewRecord(user)
}

func (manager DefaultUserManager) GetUserByAlias(alias string) *entity.User {
	user := entity.User{}
	manager.DB.Where("Alias = ?", alias).First(&user)
	return &user
}

func (manager DefaultUserManager) Close() {
	manager.DB.Close()
}

func Create() UserManager {
	db := model.GetDatabase()
	return DefaultUserManager{DB: *db}
}
