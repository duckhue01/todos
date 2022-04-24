package database

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/duckhue01/todos/models"
	"github.com/duckhue01/todos/utils"
)

type PomoDB struct {
	path string
}

func NewPomoDB(path string) *PomoDB {
	return &PomoDB{
		path: path,
	}
}

func (pomo *PomoDB) GetTodayTask() []models.Todo {

	res := map[string][]models.Todo{}
	data, err := os.ReadFile(pomo.path)
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

	if err := json.Unmarshal(data, &res); err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

	return res[utils.DateToString(time.Now())]

}
