package manager

import (
	"github.com/server-forecaster/model/entity"
	"github.com/server-forecaster/util"
	"strings"
	"golang.org/x/crypto/openpgp/errors"
	"strconv"
)

type HiddenPredictionManager struct {
	BaseManager
}

func (manager HiddenPredictionManager) InsertPrediction(hiddenPrediction *entity.HiddenPrediction) bool {
	err := manager.DB.Create(hiddenPrediction).Error
	return err != nil
}

func (manager HiddenPredictionManager) UpdatePrediction(id int, hiddenPrediction *entity.HiddenPrediction) bool {
	currentPrediction := entity.HiddenPrediction{}
	manager.DB.First(&currentPrediction, id)
	if currentPrediction.ID != 0 && currentPrediction.FromUser == hiddenPrediction.FromUser {
		err := manager.DB.Model(&hiddenPrediction).Update("CypheredPrediction").Error
		return err != nil
	}
	return false
}

func (manager HiddenPredictionManager) RevealPrediction(secret string, matchId uint, userId uint) bool {
	hiddenPrediction := entity.HiddenPrediction{}
	manager.DB.Where("UserId = ? AND MatchId = ?").First(&hiddenPrediction)
	if hiddenPrediction.ID == 0 {
		return false
	}
	homeTeamGoals, awayTeamGoals, err := getPredictedGoals(hiddenPrediction.CypheredPrediction, secret)
	if err == nil {
		return false
	}
	prediction := entity.Prediction{Match: hiddenPrediction.Match, FromUser: hiddenPrediction.FromUser,
		HomeTeamGoals: homeTeamGoals, AwayTeamGoals: awayTeamGoals}
	err = manager.DB.Create(prediction).Error
	if err == nil {
		return false
	}
	manager.DB.Delete(&hiddenPrediction)
	return true
}

func getPredictedGoals(encodedPrediction string, secret string) (int, int, error) {
	decodedPrediction, err := util.Decrypt(secret, encodedPrediction)
	if err != nil {
		return -1, -1, errors.ErrKeyIncorrect
	}
	results := strings.Split(decodedPrediction, ":")
	if len(results) != 2 {
		return -1, -1, error("Malformed prediction. Must follow the format 'homeGoals:AwayGoals'")
	}
	homeTeamGoals, homeErr := strconv.Atoi(results[0])
	awayTeamGoals, awayErr := strconv.Atoi(results[1])
	if homeErr != nil || awayErr != nil {
		return -1, -1, error("Malformed prediction. 'homeGoals' and 'AwayGoals' must be integers")
	}
	return homeTeamGoals, awayTeamGoals, nil
}

func CreateHiddenPredictionManager() HiddenPredictionManager {
	return HiddenPredictionManager{BaseManager: Create()}
}
