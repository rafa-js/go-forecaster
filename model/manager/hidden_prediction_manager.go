package manager

import (
	"github.com/server-forecaster/model/entity"
)

type HiddenPredictionManager struct {
	BaseManager
}

func (manager HiddenPredictionManager) AddPrediction(hiddenPrediction *entity.HiddenPrediction) bool {
	err := manager.DB.Create(hiddenPrediction).Error
	if err != nil{
		panic(err)
	}
	return err != nil
}

func (manager HiddenPredictionManager) UpdatePrediction(hiddenPrediction *entity.HiddenPrediction) bool {
	err := manager.DB.Model(&hiddenPrediction).Update("CypheredPrediction").Error
	if err != nil{
		panic(err)
	}
	return err != nil
}

func CreateHiddenPredictionManager() HiddenPredictionManager {
	return HiddenPredictionManager{BaseManager: Create()}
}
