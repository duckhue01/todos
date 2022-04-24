package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/duckhue01/todos/models"
)

type ConfigDB struct {
	path string
}

func NewConfigDB(path string) *ConfigDB {
	return &ConfigDB{path: path}

}

func (con *ConfigDB) GetPomoConfig() *models.PomoConfig {
	var pomoConfig *models.PomoConfig
	var setRaw, err = ioutil.ReadFile(con.path)
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}

	json.Unmarshal(setRaw, pomoConfig)

	return pomoConfig

}
