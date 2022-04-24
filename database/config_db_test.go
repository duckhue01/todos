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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
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
