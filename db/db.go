package db

import "github.com/duckhue01/todos/models"

type (
	jsonDB struct {
		path string
	}
	DB interface {
		ListSchedule() (*[]models.Schedule, error)
		ReadCurrentSchedule() (*models.Schedule, error)

		ReadPomoConfig() (*models.PomoConfig, error)
		WritePomoConfig([]byte) error

		ReadMedievalMusic() string
		ReadEpicMusic() string
		ReadPianoMusic() string
	}
)

func New(path string) *jsonDB {
	return &jsonDB{path: path}
}
