package manager

import (
	"github.com/jinzhu/gorm"
	"github.com/server-forecaster/model"
)

type BaseManager struct {
	DB *gorm.DB
}

func (manager BaseManager) Close() {
	manager.DB.Close()
}

func Create() BaseManager {
	return BaseManager{DB: model.DB}
}
