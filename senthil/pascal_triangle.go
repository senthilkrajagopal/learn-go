package senthil

func pascal(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		sub := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				sub[j] = 1
			} else {
				sub[j] = result[i-2][j-1] + result[i-2][j]
			}
		}
		result[i-1] = sub
	}
	return result
}
