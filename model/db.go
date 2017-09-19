package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/server-forecaster/model/entity"
	"os"
	"fmt"
	"net/url"
	"strings"
)

var (
	db *gorm.DB
)

func GetDatabase() *gorm.DB {
	args := getConnectionParameters()
	var err error
	db, err = gorm.Open("postgres", args)
	if err != nil {
		panic(err)
	}
	configureDatabase(db)
	return db
}

func getConnectionParameters() string {
	stringUrl := os.Getenv("DATABASE_URL")
	if stringUrl == "" {
		stringUrl = "postgres://pfswnjtgetfmco:0221520ae3266aa4fb366c2828d33296aab9c2b64df6bd9146bb3d44f2d77cec@ec2-54-247-189-64.eu-west-1.compute.amazonaws.com:5432/da0t72dr0rbnp0"
	}
	dbUrl, err := url.Parse(stringUrl)
	if err != nil {
		panic(err)
	}
	hostParams := strings.Split(dbUrl.Host, ":")
	loginParams := strings.Split(dbUrl.User.String(), ":")
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		hostParams[0], hostParams[1], loginParams[0], loginParams[1], dbUrl.Path[1:])
	return args
}

func configureDatabase(db *gorm.DB) {
	db.DB().Begin()
	db.AutoMigrate(&entity.User{}, &entity.Match{}, &entity.Prediction{}, entity.HiddenPrediction{})
}
