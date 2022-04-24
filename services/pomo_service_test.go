package services

import (
	"reflect"
	"testing"
	"time"
)

func TestNewPomo(t *testing.T) {
	type args struct {
		base string
	}
	tests := []struct {
		name string
		args args
		want *Pomo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPomo(tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPomo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPomo_StartPomoHanddler(t *testing.T) {
	type fields struct {
		base string
	}
	type args struct {
		needMusic bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pomo := &Pomo{
				base: tt.fields.base,
			}
			pomo.StartPomoHanddler(tt.args.needMusic)
		})
	}
}

func TestPomo_SetPomoHandler(t *testing.T) {
	type fields struct {
		base string
	}
	type args struct {
		key   string
		value time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pomo := &Pomo{
				base: tt.fields.base,
			}
			pomo.SetPomoHandler(tt.args.key, tt.args.value)
		})
	}
}

func TestPomo_InfoPomoHandler(t *testing.T) {
	type fields struct {
		base string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pomo := &Pomo{
				base: tt.fields.base,
			}
			pomo.InfoPomoHandler()
		})
	}
}
