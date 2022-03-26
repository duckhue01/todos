package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/duckhue01/todos/models"
)

type ScheService struct {
	db string
}

func NewScheService(db string) *ScheService {
	return &ScheService{db: db}
}

func (sche *ScheService) ListSche() {

	var scheData []models.Sche
	scheRaw, err := ioutil.ReadFile(filepath.Join(sche.db, "sche.json"))
	if err != nil {
		fmt.Println("can't read sche file")
	}

	json.Unmarshal(scheRaw, &scheData)
	fmt.Println(scheData)
}

func (sche *ScheService) CurrentSche() {
	var scheData []models.Sche
	scheRaw, err := ioutil.ReadFile(filepath.Join(sche.db, "sche.json"))
	if err != nil {
		fmt.Println("can't read sche file")
	}

	json.Unmarshal(scheRaw, &scheData)
	fmt.Println(scheData)

}
