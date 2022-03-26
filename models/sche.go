package models

type Sche struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Tasks     struct {
		Proj  []string `json:"proj"`
		Book  []string `json:"book"`
		Tech  []string `json:"tech"`
		Other []string `json:"other"`
	} `json:"tasks"`
}
