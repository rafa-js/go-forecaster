package models

type Classification struct {
	Scores []ClassificationScore `json:"scores" groups:"readable,visible"`
}
