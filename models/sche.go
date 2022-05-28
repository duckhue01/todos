package models

type Schedule struct {
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
	Tasks     struct {
		Proj  []string `json:"proj"`
		Book  []string `json:"book"`
		Tech  []string `json:"tech"`
		Task  []string `json:"task"`
		Other []string `json:"other"`
	} `json:"tasks"`
}
