package main

import (
	"reflect"
	"testing"
)

func TestCountCals(t *testing.T) {
	t.Parallel()
	type args struct {
		x []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "input is 1,2,3,0,1,2,0,50,60,0,20,30,0",
			args: args{x: []int{1, 2, 3, 0, 1, 2, 0, 50, 60, 0, 20, 30, 0}},
			want: []int{3, 6, 50, 110},
		},
		{
			name: "input is 1000,2000,3000,0,4000,0,5000,6000,0,7000,8000,9000,0,10000,0",
			args: args{x: []int{1000, 2000, 3000, 0, 4000, 0, 5000, 6000, 0, 7000, 8000, 9000, 0, 10000, 0}},
			want: []int{4000, 6000, 10000, 11000, 24000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCals(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countCals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumTopThree(t *testing.T) {
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
			name: "input is 1,2,3,4,5",
			args: args{x: []int{1, 2, 3, 4, 5}},
			want: 12,
		},
		{
			name: "input is 4000,6000,10000,11000,24000",
			args: args{x: []int{4000, 6000, 10000, 11000, 24000}},
			want: 45000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumTopThree(tt.args.x); got != tt.want {
				t.Errorf("sumTopThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
