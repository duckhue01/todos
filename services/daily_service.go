package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/duckhue01/todos/models"
)

type Daily struct {
	db string

}

func NewDaily() *Daily {
	return &Daily{db: "/Users/duckhue01/code/side/todos/daily.json"}
}

func (d *Daily) List(date string) []models.Todo {
	return d.getTodo(date)
}

func (d *Daily) Add(date string, todo *models.Todo) {



}

func (d *Daily) Delete()  {
	
}

func (d *Daily) Done()  {
	
}

func (d *Daily) getTodo(date string) []models.Todo {
	res := map[string][]models.Todo{}
	data, err := os.ReadFile(d.db)
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

	if err := json.Unmarshal(data, &res); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

	return res[date]
}
