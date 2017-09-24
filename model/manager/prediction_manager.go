package manager

import (
	"github.com/server-forecaster/model/entity"
)

type PredictionManager struct {
	BaseManager
}

func (manager PredictionManager) GetByMatch(matchId uint) ([]entity.Prediction, error) {
	predictions := []entity.Prediction{}
	err := manager.DB.Where("match_id = ?", matchId).Find(&predictions).Error
	return predictions, err
}

func (manager PredictionManager) UpdatePredictionResults(match entity.Match) {
	if match.Status != "FINISHED" {
		return
	}
	predictions := []entity.Prediction{}
	err := manager.DB.Where("match_id = ?", match.ID).Find(&predictions).Error
	if err != nil {
		panic(err)
	}
	for _, prediction := range predictions {
		isHit := prediction.AwayTeamGoals == match.AwayTeamGoals &&
			prediction.HomeTeamGoals == match.HomeTeamGoals
		manager.DB.Model(&prediction).Update("IsHit", isHit)
	}
}

func CreatePredictionManager() PredictionManager {
	return PredictionManager{BaseManager: Create()}
}
