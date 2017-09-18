package model

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/server-forecaster/model/entity"
	"os"
	"fmt"
	"net/url"
)

func GetDatabase() *gorm.DB {
	args := getConnectionParameters()
	db, err := gorm.Open("mysql", args[1])
	if err != nil {
		panic(err)
	}
	configureDatabase(db)
	return db
}

func getConnectionParameters() string {
	stringUrl := os.Getenv("JAWSDB_URL")
	if stringUrl == "" {
		stringUrl = "mysql://ztqr51phul7ksdlw:qanww8l60vd3n8sw@cvktne7b4wbj4ks1.chr7pe7iynqr.eu-west-1.rds.amazonaws.com:3306/llrfhejo56g49g35"
	}
	dbUrl, err := url.Parse(stringUrl)
	if err != nil {
		panic(err)
	}
	args := fmt.Sprintf("%v@tcp(%v)%v?charset=utf8&parseTime=True&loc=Local",
		dbUrl.User.String(), dbUrl.Host, dbUrl.Path)
	return args
}

func configureDatabase(db *gorm.DB) {
	db.AutoMigrate(&entity.User{}, &entity.Match{}, &entity.Prediction{}, entity.HiddenPrediction{},
		entity.ClassificationScore{}, entity.Classification{})
}
