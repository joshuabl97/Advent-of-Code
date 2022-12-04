package main

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	testInputG1, testInputG2 := parseInput("testInput.txt")
	inputG1, inputG2 := parseInput("input.txt")

	t.Parallel()

	type args struct {
		group1 [][]int
		group2 [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input is testInput.txt",
			args: args{group1: testInputG1, group2: testInputG2},
			want: 2,
		},
		{
			name: "input is input.txt",
			args: args{group1: inputG1, group2: inputG2},
			want: 547,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem1(tt.args.group1, tt.args.group2); got != tt.want {
				t.Errorf("problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProblem2(t *testing.T) {
	testInputG1, testInputG2 := parseInput("testInput.txt")
	inputG1, inputG2 := parseInput("input.txt")

	t.Parallel()

	type args struct {
		group1 [][]int
		group2 [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input is testInput.txt",
			args: args{group1: testInputG1, group2: testInputG2},
			want: 4,
		},
		{
			name: "input is input.txt",
			args: args{group1: inputG1, group2: inputG2},
			want: 843,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem2(tt.args.group1, tt.args.group2); got != tt.want {
				t.Errorf("problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}
