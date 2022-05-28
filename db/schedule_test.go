package db

import (
	"reflect"
	"testing"

	"github.com/duckhue01/todos/models"
)

func TestNew(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *jsonDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonDB_ListSchedule(t *testing.T) {
	type fields struct {
		path string
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
			db := &jsonDB{
				path: tt.fields.path,
			}
			got, err := db.ListSchedule()
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonDB.ListSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonDB.ListSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_jsonDB_GetCurrentSchedule(t *testing.T) {
	type fields struct {
		path string
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
			db := &jsonDB{
				path: tt.fields.path,
			}
			got, err := db.GetCurrentSchedule()
			if (err != nil) != tt.wantErr {
				t.Errorf("jsonDB.GetCurrentSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("jsonDB.GetCurrentSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}
