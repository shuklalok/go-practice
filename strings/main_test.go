package main

import (
	"testing"
)

func Test_countVowels(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test", args{str: "hello"}, 2},
		{"test", args{str: "hola"}, 2},
		{"test", args{str: "mundo"}, 2},
		{"test", args{str: "hello world"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowels(tt.args.str); got != tt.want {
				t.Errorf("countVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_countNumerics(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test", args{str: "hello"}, 0},
		{"test", args{str: "hola1"}, 1},
		{"test", args{str: "mundo123"}, 3},
		{"test", args{str: "hello world 1234"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumerics(tt.args.str); got != tt.want {
				t.Errorf("countNumerics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countConsonants(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test", args{str: "hello"}, 3},
		{"test", args{str: "hola"}, 2},
		{"test", args{str: "mundo"}, 3},
		{"test", args{str: "hello world"}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countConsonants(tt.args.str); got != tt.want {
				t.Errorf("countConsonants() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSpaces(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test", args{str: "hello"}, 0},
		{"test", args{str: "hello world"}, 1},
		{"test", args{str: "hello world 1"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSpaces(tt.args.str); got != tt.want {
				t.Errorf("countSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
