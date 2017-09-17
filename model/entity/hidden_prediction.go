package entity

import (
	"github.com/jinzhu/gorm"
)

type HiddenPrediction struct {
	gorm.Model
	FromUser           User   `json:"fromUser" groups:"readable,visible"`
	Match              Match  `json:"match" groups:"readable,visible"`
	CypheredPrediction string `json:"cypheredPrediction" groups:"readable,visible"`
}
