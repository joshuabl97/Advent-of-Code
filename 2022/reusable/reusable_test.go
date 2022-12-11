package reusable

import (
	"reflect"
	"testing"
)

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
			if got := CharNum(tt.args); got != tt.want {
				t.Errorf("CharNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDupeExist(t *testing.T) {
	t.Parallel()
	type args struct {
		arr []string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "input is ['a', 'b', 'b', 'c']",
			args: args{arr: []string{"a", "b", "b", "c"}},
			want: true,
		},
		{
			name: "input is ['a', 'b', 'c']",
			args: args{arr: []string{"a", "b", "c"}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DupeExist(tt.args.arr); got != tt.want {
				t.Errorf("DupeExist() = %v, want %v", got, tt.want)
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
			if got := NestedStringUnique(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NestedStringUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseIntSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "input is 1,2,3",
			args: []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			name: "input is 1, 4, 3, 10, 1, 1239867, 57192",
			args: []int{1, 4, 3, 10, 1, 1239867, 57192},
			want: []int{57192, 1239867, 1, 10, 3, 4, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseIntSlice(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueStrSlice(t *testing.T) {
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
			if got := UniqueStrSlice(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueStrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
