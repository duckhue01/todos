package database

import (
	"reflect"
	"testing"

	"github.com/duckhue01/todos/models"
)

func TestNewPomoDB(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *PomoDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPomoDB(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPomoDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPomoDB_GetTodayTask(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		want   []models.Todo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pomo := &PomoDB{
				path: tt.fields.path,
			}
			if got := pomo.GetTodayTask(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PomoDB.GetTodayTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
