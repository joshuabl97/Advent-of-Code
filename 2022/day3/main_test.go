package main

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	t.Parallel()
	type args struct {
		x []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "input is ['a', 'b', 'b', 'c']",
			args: args{x: []string{"a", "b", "b", "c"}},
			want: []string{"a", "b", "c"},
		},
		{
			name: "input is ['z', 'z', 'x', 'x', 'y', 'y', 'y']",
			args: args{x: []string{"z", "z", "x", "x", "y", "y", "y"}},
			want: []string{"z", "x", "y"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharNum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "input is 'p'",
			args: "p",
			want: 16,
		},
		{
			name: "input is 'L'",
			args: "L",
			want: 38,
		},
		{
			name: "input is 'P'",
			args: "P",
			want: 42,
		},
		{
			name: "input is 't'",
			args: "t",
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := charNum(tt.args); got != tt.want {
				t.Errorf("charNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestedStringUnique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args [][]string
		want [][]string
	}{
		{
			name: "input is [['a','a','b','c','c'], ['a','b','b','c','C']]",
			args: [][]string{{"a", "a", "b", "c", "c"}, {"a", "b", "b", "c", "c", "C"}},
			want: [][]string{{"a", "b", "c"}, {"a", "b", "c", "C"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nestedStringUnique(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nestedStringUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
