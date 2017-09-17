package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Alias    string `json:"alias" groups:"readable,visible"`
	Password string `json:"-" groups:"readable"`
	Email    string `json:"match" groups:"readable"`
}
