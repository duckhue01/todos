package services

import (
	"reflect"
	"testing"

	"github.com/duckhue01/todos/db"
	"github.com/duckhue01/todos/models"
)

func TestNewSchedule(t *testing.T) {
	type args struct {
		db db.DB
	}
	tests := []struct {
		name string
		args args
		want *Schedule
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSchedule(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchedule_ListSchedule(t *testing.T) {
	type fields struct {
		db db.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    *[]models.Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schedule{
				db: tt.fields.db,
			}
			got, err := s.ListSchedule()
			if (err != nil) != tt.wantErr {
				t.Errorf("Schedule.ListSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schedule.ListSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchedule_GetCurrentSchedule(t *testing.T) {
	type fields struct {
		db db.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    *models.Schedule
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schedule{
				db: tt.fields.db,
			}
			got, err := s.GetCurrentSchedule()
			if (err != nil) != tt.wantErr {
				t.Errorf("Schedule.GetCurrentSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schedule.GetCurrentSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}
