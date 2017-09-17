package models

import (
	"github.com/jinzhu/gorm"
)

type Prediction struct {
	gorm.Model
	Match         Match  `json:"match" groups:"readable,visible"`
	FromUser      User   `json:"fromUser" groups:"readable,visible"`
	HomeTeamGoals int    `json:"homeTeamGoals" groups:"readable,visible"`
	AwayTeamGoals int    `json:"awayTeamGoals" groups:"readable,visible"`
}
