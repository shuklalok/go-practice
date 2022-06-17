package main

import (
	"reflect"
	"testing"
)

func TestHuman_Laugh(t *testing.T) {
	type fields struct {
		eyes  int
		hands int
		tail  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Human{
				eyes:  tt.fields.eyes,
				hands: tt.fields.hands,
				tail:  tt.fields.tail,
			}
			if got := h.Laugh(); got != tt.want {
				t.Errorf("Human.Laugh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHuman(t *testing.T) {
	tests := []struct {
		name string
		want *Human
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHuman(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}
