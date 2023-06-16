package senthil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	result := TwoSum([]int{1, 5, 8, 7}, 13)
	expected := [2]int{1, 2}
	assert.Equal(t, result, expected)
}
