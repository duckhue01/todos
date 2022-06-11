package db

import "github.com/duckhue01/todos/model"

type (
	jsonDB struct {
		path string
	}
	DB interface {
		ReadPomoConfig() (*model.PomoConfig, error)
		WritePomoConfig([]byte) error

		ReadMedievalMusic() string
		ReadEpicMusic() string
		ReadPianoMusic() string
		ReadChillMusic() string
	}
)

func New(path string) *jsonDB {
	return &jsonDB{path: path}
}
