package db

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/duckhue01/todos/model"
)

func (db *jsonDB) ReadPomoConfig() (*model.PomoConfig, error) {
	pomoConfig := &model.PomoConfig{}
	var setRaw, err = ioutil.ReadFile(filepath.Join(db.path, "config/pomo.json"))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(setRaw, &pomoConfig)
	if err != nil {
		return nil, err
	}

	return pomoConfig, nil
}

func (db *jsonDB) WritePomoConfig(data []byte) error {
	err := ioutil.WriteFile(filepath.Join(db.path, "config/pomo.json"), data, 0644)
	if err != nil {
		return err
	}
	return nil
}
