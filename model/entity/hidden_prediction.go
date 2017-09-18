package entity

import (
	"github.com/jinzhu/gorm"
)

type HiddenPrediction struct {
	gorm.Model
	FromUser           User   `json:"fromUser" groups:"readable,visible" gorm:"index:idx_hidden_prediction"`
	Match              Match  `json:"match" groups:"readable,visible" gorm:"index:idx_hidden_prediction"`
	CypheredPrediction string `json:"cypheredPrediction" groups:"readable,visible"`
}
