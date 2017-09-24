package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Match struct {
	gorm.Model
	Competition   string    `json:"competition" groups:"readable,visible" gorm:"index:idx_match"`
	Date          time.Time `json:"date" groups:"readable,visible"`
	MatchDay      int       `json:"matchDay" groups:"readable,visible" gorm:"index:idx_match"`
	Status        string    `json:"status" groups:"readable,visible"`
	HomeTeamName  string    `json:"homeTeamName" groups:"readable,visible" gorm:"index:idx_match"`
	AwayTeamName  string    `json:"awayTeamName" groups:"readable,visible" gorm:"index:idx_match"`
	HomeTeamGoals int       `json:"homeTeamGoals" groups:"readable,visible"`
	AwayTeamGoals int       `json:"awayTeamGoals" groups:"readable,visible"`
}
