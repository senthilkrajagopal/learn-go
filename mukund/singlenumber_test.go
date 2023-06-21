package main

import (
	"testing"
)

func TestSingleNumberMap(t *testing.T) {
	nums := []int{2, 2, 1}
	result := 1

	if SingleNumberMap(nums) != result {
		t.Fail()
	}
}

func TestSingleNumberXOR(t *testing.T) {
	nums := []int{2, 2, 1}
	result := 1

	if SingleNumberXOR(nums) != result {
		t.Fail()
	}
}
