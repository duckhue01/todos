package db

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/duckhue01/todos/models"
)

func (db *jsonDB) GetPomoConfig() (*models.PomoConfig, error) {
	pomoConfig := &models.PomoConfig{}
	var setRaw, err = ioutil.ReadFile(filepath.Join(db.path, "pomo.json"))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(setRaw, &pomoConfig)
	if err != nil {
		return nil, err
	}

	return pomoConfig, nil
}
