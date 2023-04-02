package main

import "testing"

func TestGetSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "absent nums", args: args{nums: nil}, want: 0},
		{name: "single num", args: args{nums: []int{1}}, want: 1},
		{name: "positive and negative", args: args{nums: []int{-1, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSum(tt.args.nums...); got != tt.want {
				t.Errorf("GetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
