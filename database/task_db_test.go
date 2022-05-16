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
		{
			name: "test case #1: want to receive new  PomoDB struct",
			args: args{
				path: "/Users/duckhue01/code/side/todos/storage/_test/daily.json",
			},
			want: &PomoDB{
				path: "/Users/duckhue01/code/side/todos/storage/_test/daily.json",
			},
		},
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
		{
			name: "test case #1: want to receive correct today task",
			fields: fields{
				path: "/Users/duckhue01/code/side/todos/storage/_test/daily.json",
			},
			want: []models.Todo{
				{
					Title:  "title",
					Tag:    "des",
					IsDone: false,
					Id:     1,
				},
				{
					Title:  "title",
					Tag:    "des",
					IsDone: false,
					Id:     2,
				},
				{
					Title:  "title",
					Tag:    "des",
					IsDone: false,
					Id:     3,
				},
			},
		},
		{
			name: "test case #1: invalid path, want to ",
			fields: fields{
				path: "/Users/duckhue01/code/side/todos/storage/_test/daily.json",
			},
			want: []models.Todo{
				{
					Title:  "title",
					Tag:    "des",
					IsDone: false,
					Id:     1,
				},
				{
					Title:  "title",
					Tag:    "des",
					IsDone: false,
					Id:     2,
				},
				{
					Title:  "title",
					Tag:    "des",
					IsDone: false,
					Id:     3,
				},
			},
		},
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
