package main

import "testing"

func TestNumList_Sum(t *testing.T) {
	type fields struct {
		args []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{name: "zero args", fields: fields{}, want: 0},
		{name: "single arg", fields: fields{[]int{1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NumList{
				args: tt.fields.args,
			}
			if got := n.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
