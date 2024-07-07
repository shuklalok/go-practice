package main

import (
	"testing"
)

func Test_circle_area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{"test", fields{radius: 5}, 78.5},
		{"test", fields{radius: 10}, 314},
		{"test", fields{radius: 20}, 1256},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := circle{
				radius: tt.fields.radius,
			}
			if got := c.area(); got != tt.want {
				t.Errorf("circle.area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_square_area(t *testing.T) {
	type fields struct {
		length float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{"test", fields{length: 5}, 25},
		{"test", fields{length: 10}, 100},
		{"test", fields{length: 20}, 400},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := square{
				length: tt.fields.length,
			}
			if got := s.area(); got != tt.want {
				t.Errorf("square.area() = %v, want %v", got, tt.want)
			}
		})
	}
}
