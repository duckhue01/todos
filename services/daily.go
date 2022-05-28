package services

import (
	"github.com/duckhue01/todos/models"
)

type Daily struct {
	db string
}

func NewDaily() *Daily {
	return &Daily{db: "/Users/duckhue01/code/side/todos/daily.json"}
}

func (d *Daily) List(date string) []models.Todo {

	return nil
}

func (d *Daily) Add(date string, todo *models.Todo) {

}

func (d *Daily) Delete() {

}

func (d *Daily) Done() {

}
