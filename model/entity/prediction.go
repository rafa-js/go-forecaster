package entity

import (
	"github.com/jinzhu/gorm"
)

type Prediction struct {
	gorm.Model
	Match         Match  `json:"match" groups:"readable,visible" gorm:"index:idx_match"`
	FromUser      User   `json:"fromUser" groups:"readable,visible" gorm:"index:idx_match "`
	HomeTeamGoals int    `json:"homeTeamGoals" groups:"readable,visible"`
	AwayTeamGoals int    `json:"awayTeamGoals" groups:"readable,visible"`
}
