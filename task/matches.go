package task

import (
	"time"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/server-forecaster/model/entity"
	"github.com/server-forecaster/model/manager"
)

const SERVICE_ENDPOINT = "http://api.football-data.org/v1/teams/86/fixtures"

type JsonMatchResult struct {
	HomeTeamGoals int `json:"goalsHomeTeam"`
	AwayTeamGoals int `json:"goalsAwayTeam"`
}

type Links struct {
	Competition string `json:"competition"`
	HomeTeam    string `json:"homeTeam"`
	AwayTeam    string `json:"awayTeam"`
}

type JsonMatch struct {
	Date         time.Time `json:"date"`
	Status       string `json:"status"`
	Competition  string `json:"competition"`
	MatchDay     int `json:"matchDay"`
	HomeTeamName string `json:"homeTeamName"`
	AwayTeamName string `json:"awayTeamName"`
	Result       JsonMatchResult `json:"result"`
	Links        Links `json:"_links"`
}

type ApiResponse struct {
	Fixtures []JsonMatch `json:"fixtures"`
}

func UpdateMatches() []entity.Match {
	apiResponse := getApiMatches()
	matchManager := manager.CreateMatchManager()
	updatedMatches := []entity.Match{}
	for _, jsonMatch := range apiResponse.Fixtures {
		match := createMatchFromJson(&jsonMatch)
		if matchManager.AddOrUpdateMatch(match) {
			updatedMatches = append(updatedMatches, *match)
		}
	}
	return updatedMatches
}

func getApiMatches() ApiResponse {
	jsonData := getApiBody()
	response := ApiResponse{}
	jsonErr := json.Unmarshal(jsonData, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return response
}

func getApiBody() []byte {
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, SERVICE_ENDPOINT, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func createMatchFromJson(jsonMatch *JsonMatch) *entity.Match {
	match := entity.Match{}
	match.Date = jsonMatch.Date
	match.MatchDay = jsonMatch.MatchDay
	match.Competition = jsonMatch.Links.Competition
	match.Status = jsonMatch.Status
	match.HomeTeamName = jsonMatch.HomeTeamName
	match.AwayTeamName = jsonMatch.AwayTeamName
	match.HomeTeamGoals = jsonMatch.Result.HomeTeamGoals
	match.AwayTeamGoals = jsonMatch.Result.AwayTeamGoals
	return &match
}
