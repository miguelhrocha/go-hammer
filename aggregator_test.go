package main

import (
	"testing"
)

func TestP50(t *testing.T) {
	values := getValues()
	p := percentile(values, 50)
	if p != 9 {
		t.Error("p50 should be 9", p)
	}
}

func TestP90(t *testing.T) {
	values := getValues()
	p := percentile(values, 90)
	if p != 15 {
		t.Error("p90 should be 15", p)
	}
}

func TestP99(t *testing.T) {
	values := getValues()
	p := percentile(values, 99)
	if p != 15 {
		t.Error("p90 should be 15", p)
	}
}

func getValues() []int {
	return []int{
		3,
		5,
		7,
		8,
		9,
		11,
		13,
		15,
		15,
	}
}
