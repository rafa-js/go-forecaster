package manager

import (
	"github.com/server-forecaster/model/entity"
)

type HiddenPredictionManager struct {
	BaseManager
}

func (manager HiddenPredictionManager) AddPrediction(hiddenPrediction *entity.HiddenPrediction) bool {
	err := manager.DB.Create(hiddenPrediction).Error
	return err != nil
}

func (manager HiddenPredictionManager) UpdatePrediction(hiddenPrediction *entity.HiddenPrediction) bool {
	err := manager.DB.Model(&hiddenPrediction).Update("CypheredPrediction").Error
	return err != nil
}

func CreateHiddenPredictionManager() HiddenPredictionManager {
	return HiddenPredictionManager{BaseManager: Create()}
}
