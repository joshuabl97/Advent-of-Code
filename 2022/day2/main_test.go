package main

import (
	"testing"
)

func TestPlayRPS(t *testing.T) {
	t.Parallel()
	type args struct {
		x []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input is 1,2",
			args: args{x: []int{1, 2}},
			want: 8,
		},
		{
			name: "input is 2,1",
			args: args{x: []int{2, 1}},
			want: 1,
		},
		{
			name: "input is 3,3",
			args: args{x: []int{3, 3}},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playRPS(tt.args.x); got != tt.want {
				t.Errorf("playRPS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProblem1(t *testing.T) {
	t.Parallel()
	type args struct {
		x [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input is [][]int{{1, 2}, {2, 1}, {3, 3}}",
			args: args{x: [][]int{{1, 2}, {2, 1}, {3, 3}}},
			want: 15,
		},
		{
			name: "input is [][]int{{1, 2}, {2, 1}, {3, 3},{1, 2}, {2, 1}, {3, 3},{3, 2},{2, 3}}",
			args: args{x: [][]int{{1, 2}, {2, 1}, {3, 3}, {1, 2}, {2, 1}, {3, 3}, {3, 2}, {2, 3}}},
			want: 41,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem1(tt.args.x); got != tt.want {
				t.Errorf("problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProblem2(t *testing.T) {
	t.Parallel()
	type args struct {
		x [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "input is [][]int{{1, 2}, {2, 1}, {3, 3}",
			args: args{x: [][]int{{1, 2}, {2, 1}, {3, 3}}},
			want: 12,
		},
		{
			name: "input is [][]int{{1, 2}, {2, 1}, {3, 1}, {2, 2}, {1, 2}}",
			args: args{x: [][]int{{1, 2}, {2, 1}, {3, 1}, {2, 2}, {1, 2}}},
			want: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem2(tt.args.x); got != tt.want {
				t.Errorf("problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}
