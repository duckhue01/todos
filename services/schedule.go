package services

import (
	"github.com/duckhue01/todos/db"
	"github.com/duckhue01/todos/models"
)

type Schedule struct {
	db db.DB
}

func NewSchedule(db db.DB) *Schedule {
	return &Schedule{db: db}
}

func (s *Schedule) ListSchedule() (*[]models.Schedule, error) {
	return s.db.ListSchedule()
}

func (s *Schedule) GetCurrentSchedule() (*models.Schedule, error) {
	return s.db.ReadCurrentSchedule()

}
