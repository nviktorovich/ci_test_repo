package main

import "testing"

func TestGetSum(t *testing.T) {
	got := GetSum(1)
	want := 1
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
