package entity

import (
	"github.com/jinzhu/gorm"
)

type Prediction struct {
	gorm.Model
	Match         Match  `json:"match" groups:"readable,visible"`
	MatchID       uint   `gorm:"unique_index:idx_match"`
	FromUser      User   `json:"fromUser" groups:"readable,visible"`
	FromUserID    uint   `gorm:"unique_index:idx_match"`
	HomeTeamGoals int    `json:"homeTeamGoals" groups:"readable,visible"`
	AwayTeamGoals int    `json:"awayTeamGoals" groups:"readable,visible"`
}
