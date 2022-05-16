package database

import (
	"reflect"
	"testing"

	"github.com/duckhue01/todos/models"
)

func TestNewConfigDB(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *ConfigDB
	}{
		{
			"test case #1: want to receive ConfigDB struct",
			args{path: "/Users/duckhue01/code/side/todos/storage/_test/config/pomo.json"},
			&ConfigDB{
				path: "/Users/duckhue01/code/side/todos/storage/_test/config/pomo.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigDB(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigDB_GetPomoConfig(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.PomoConfig
	}{
		{
			name: "test case #1: want to receive correct config struct",
			fields: fields{
				path: "/Users/duckhue01/code/side/todos/storage/_test/config/pomo.json",
			},
			want: &models.PomoConfig{
				Pomo:     10,
				Break:    20,
				Interval: 30,
			},
		},
		{
			name: "test case #1: wrong storage path, want to receive error",
			fields: fields{
				path: "wrong/path/pomo.json",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := &ConfigDB{
				path: tt.fields.path,
			}
			if got := con.GetPomoConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConfigDB.GetPomoConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
