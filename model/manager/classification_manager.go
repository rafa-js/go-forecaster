package manager

import (
	"github.com/server-forecaster/model/entity"
)

type ClassificationManager struct {
	BaseManager
}

func (manager ClassificationManager) GetClassification() *entity.Classification {
	classification := entity.Classification{Scores: []entity.ClassificationScore{}}
	users := []entity.User{}
	err := manager.DB.Find(&users).Error
	if err != nil {
		panic(err)
	}
	for _, user := range users {
		hits := []entity.Prediction{}
		err := manager.DB.Where("is_hit = true AND from_user_id = ?", user.ID).Find(&hits).Error
		if err != nil {
			panic(err)
		}
		manager.DB.Preload("Match").Find(&hits)
		classificationScore := entity.ClassificationScore{User: user, Hits: hits, TotalHits: len(hits)}
		classification.Scores = append(classification.Scores, classificationScore)
	}
	return &classification
}

func CreateClassificationManager() ClassificationManager {
	return ClassificationManager{BaseManager: Create()}
}
