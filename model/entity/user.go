package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Alias     string `json:"alias"      groups:"readable,visible"`
	Password  string `json:"password"   groups:"readable"`
	Email     string `json:"email"      groups:"readable,visible"`
	AuthToken string `json:"authToken"  groups:"visible"`
}
