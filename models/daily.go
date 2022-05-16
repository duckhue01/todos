package models


type Todo struct {
	Title  string `json:"title"`
	Tag    string `json:"tag"`
	IsDone bool   `json:"isDone"`
	Id     int    `json:"id"`
}
