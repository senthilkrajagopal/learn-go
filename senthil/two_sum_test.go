package senthil

import "testing"

func TestTwoSum(t *testing.T) {
	result := TwoSum([]int{1, 5, 8, 7}, 13)
	if result != [2]int{1, 2} {
		t.Fail()
	}
}
