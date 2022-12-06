package main

import (
	"reflect"
	"testing"
)

func TestProblem1(t *testing.T) {
	testInput := parseInput("testInput.txt")
	input := parseInput("input.txt")

	t.Parallel()

	type args struct {
		x [][]string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input is testInput.txt",
			args: args{x: testInput},
			want: 157,
		},
		{
			name: "input is input.txt",
			args: args{x: input},
			want: 8401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem1(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProblem2(t *testing.T) {
	testInput := parseInput("testInput.txt")
	input := parseInput("input.txt")

	t.Parallel()

	type args struct {
		x [][]string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input is testInput.txt",
			args: args{x: testInput},
			want: 70,
		},
		{
			name: "input is input.txt",
			args: args{x: input},
			want: 2641,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem2(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}
