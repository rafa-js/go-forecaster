package entity

type Classification struct {
	Scores []ClassificationScore `json:"scores" groups:"readable,visible"`
}
