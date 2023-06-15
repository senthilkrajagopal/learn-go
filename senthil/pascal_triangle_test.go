package senthil

import (
	"fmt"
	"testing"
)

func TestMakeTriangle(t *testing.T) {
	//expected := [][]int{{1}, {1, 1}, {1, 2, 1}}

	result := pascal(5)
	for _, sub := range result {
		fmt.Println(sub)
	}

}
