package entity

type ClassificationScore struct {
	User        User          `json:"user" groups:"readable,visible"`
	TotalHits   int           `json:"totalHits" groups:"readable,visible"`
	Hits        []Prediction  `json:"hits" groups:"readable,visible"`
	Predictions []Prediction  `json:"predictions" groups:"readable,visible"`
}
