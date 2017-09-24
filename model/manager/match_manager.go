package manager

import (
	"github.com/server-forecaster/model/entity"
)

type MatchManager struct {
	BaseManager
}

func (manager MatchManager) GetCalendar() []entity.Match {
	matches := []entity.Match{}
	manager.DB.Find(&matches).Order("Date asc")
	return matches
}

func (manager MatchManager) AddOrUpdateMatch(match *entity.Match) bool {
	currentMatch := entity.Match{}
	manager.DB.Where(entity.Match{Competition: match.Competition, MatchDay: match.MatchDay,
		HomeTeamName: match.HomeTeamName, AwayTeamName: match.AwayTeamName}).First(&currentMatch)
	isNewMatch := manager.DB.NewRecord(&currentMatch)
	if !isNewMatch {
		match.ID = currentMatch.ID
	}
	manager.DB.Save(match)
	return isNewMatch
}

func CreateMatchManager() MatchManager {
	return MatchManager{BaseManager: Create()}
}
