package services

import (
	"reflect"
	"testing"
)

func TestNewScheService(t *testing.T) {
	type args struct {
		db string
	}
	tests := []struct {
		name string
		args args
		want *ScheService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewScheService(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewScheService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheService_ListSche(t *testing.T) {
	type fields struct {
		db string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sche := &ScheService{
				db: tt.fields.db,
			}
			sche.ListSche()
		})
	}
}

func TestScheService_CurrentSche(t *testing.T) {
	type fields struct {
		db string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sche := &ScheService{
				db: tt.fields.db,
			}
			sche.CurrentSche()
		})
	}
}
