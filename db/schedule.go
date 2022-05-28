package db

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/duckhue01/todos/models"
)



func (db *jsonDB) ListSchedule() (*[]models.Schedule, error) {
	res := &[]models.Schedule{}
	bytes, err := ioutil.ReadFile(filepath.Join(db.path, "schedule.json"))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, res)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (db *jsonDB) GetCurrentSchedule() (*models.Schedule, error) {
	res := &models.Schedule{}

	bytes, err := ioutil.ReadFile(filepath.Join(db.path, "schedule.json"))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, res)
	if err != nil {
		return nil, err
	}

	return res, nil

}
