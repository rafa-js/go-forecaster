package manager

import (
	"github.com/server-forecaster/model/entity"
	"github.com/server-forecaster/util"
	"strings"
	"errors"
	"strconv"
)

type HiddenPredictionManager struct {
	BaseManager
}

func (manager HiddenPredictionManager) InsertPrediction(hiddenPrediction *entity.HiddenPrediction) error {
	return manager.DB.Create(hiddenPrediction).Error
}

func (manager HiddenPredictionManager) UpdatePrediction(id int, hiddenPrediction *entity.HiddenPrediction) error {
	currentPrediction := entity.HiddenPrediction{}
	manager.DB.First(&currentPrediction, id)
	if currentPrediction.ID == 0 {
		return errors.New("Invalid ID")
	} else if currentPrediction.FromUser == hiddenPrediction.FromUser {
		return errors.New("You don't have permission to update this entity")
	} else {
		currentPrediction.CypheredPrediction = hiddenPrediction.CypheredPrediction
		err := manager.DB.Save(currentPrediction).Error
		return err
	}
}

func (manager HiddenPredictionManager) RevealPrediction(secret string, matchId uint, userId uint) error {
	hiddenPrediction := entity.HiddenPrediction{}
	err := manager.DB.Where("from_user_id = ? AND match_id = ?", userId, matchId).First(&hiddenPrediction).Error
	if err != nil {
		return err
	}
	homeTeamGoals, awayTeamGoals, err := getPredictedGoals(hiddenPrediction.CypheredPrediction, secret)
	if err != nil {
		return err
	}
	prediction := entity.Prediction{Match: hiddenPrediction.Match, FromUser: hiddenPrediction.FromUser,
		HomeTeamGoals: homeTeamGoals, AwayTeamGoals: awayTeamGoals}
	err = manager.DB.Create(prediction).Error
	if err != nil {
		return err
	}
	manager.DB.Delete(&hiddenPrediction)
	return nil
}

func getPredictedGoals(encodedPrediction string, secret string) (int, int, error) {
	decodedPrediction, err := util.Decrypt(secret, encodedPrediction)
	if err != nil {
		return -1, -1, errors.New("Incorrect key")
	}
	results := strings.Split(decodedPrediction, ":")
	if len(results) != 2 {
		return -1, -1, errors.New("Malformed prediction. Must follow the format 'homeGoals:AwayGoals'")
	}
	homeTeamGoals, homeErr := strconv.Atoi(results[0])
	awayTeamGoals, awayErr := strconv.Atoi(results[1])
	if homeErr != nil || awayErr != nil {
		return -1, -1, errors.New("Malformed prediction. 'homeGoals' and 'AwayGoals' must be integers")
	}
	return homeTeamGoals, awayTeamGoals, nil
}

func CreateHiddenPredictionManager() HiddenPredictionManager {
	return HiddenPredictionManager{BaseManager: Create()}
}
