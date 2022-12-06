package main

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		str  string
		want int
	}{
		{
			name: "input is mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			str:  "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 7,
		},
		{
			name: "input is bvwbjplbgvbhsrlpgdmjqwftvncz",
			str:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 5,
		},
		{
			name: "input is nppdvjthqldpwncqszvftbrmjlhg",
			str:  "nppdvjthqldpwncqszvftbrmjlhg",
			want: 6,
		},
		{
			name: "input is nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			str:  "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 10,
		},
		{
			name: "input is zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			str:  "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem1(tt.str); got != tt.want {
				t.Errorf("problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProblem2(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		str  string
		want int
	}{
		{
			name: "input is mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			str:  "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 19,
		},
		{
			name: "input is bvwbjplbgvbhsrlpgdmjqwftvncz",
			str:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 23,
		},
		{
			name: "input is nppdvjthqldpwncqszvftbrmjlhg",
			str:  "nppdvjthqldpwncqszvftbrmjlhg",
			want: 23,
		},
		{
			name: "input is nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			str:  "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 29,
		},
		{
			name: "input is zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			str:  "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := problem2(tt.str); got != tt.want {
				t.Errorf("problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}
