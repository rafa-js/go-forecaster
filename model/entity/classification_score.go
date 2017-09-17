package entity

type ClassificationScore struct {
	User      User       `json:"user" groups:"readable,visible"`
	TotalHits int        `json:"totalHits" groups:"readable,visible"`
	Hits      []Match    `json:"hits" groups:"readable,visible"`
}
