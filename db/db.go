package db

import "github.com/duckhue01/todos/models"

type (
	jsonDB struct {
		path string
	}
	DB interface {
		ListSchedule() (*[]models.Schedule, error)
		GetCurrentSchedule() (*models.Schedule, error)
		GetPomoConfig() (*models.PomoConfig, error)
	}
)

func New(path string) *jsonDB {
	return &jsonDB{path: path}
}
