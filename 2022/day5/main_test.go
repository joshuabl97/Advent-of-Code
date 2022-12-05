package main

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	stacksT1, instructionsT1 := parseInput("testInput.txt")
	stacksT2, instructionsT2 := parseInput("input.txt")

	t.Parallel()

	type args struct {
		stacks       map[int]string
		instructions [][]int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input is testInput.txt",
			args: args{stacks: stacksT1, instructions: instructionsT1},
			want: "CMZ",
		},
		{
			name: "input is input.txt",
			args: args{stacks: stacksT2, instructions: instructionsT2},
			want: "CFFHVVHNC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem1(tt.args.stacks, tt.args.instructions); got != tt.want {
				t.Errorf("problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProblem2(t *testing.T) {
	stacksT1, instructionsT1 := parseInput("testInput.txt")
	stacksT2, instructionsT2 := parseInput("input.txt")

	t.Parallel()

	type args struct {
		stacks       map[int]string
		instructions [][]int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "input is testInput.txt",
			args: args{stacks: stacksT1, instructions: instructionsT1},
			want: "MCD",
		},
		{
			name: "input is input.txt",
			args: args{stacks: stacksT2, instructions: instructionsT2},
			want: "FSZWBPTBG",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem2(tt.args.stacks, tt.args.instructions); got != tt.want {
				t.Errorf("problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}
