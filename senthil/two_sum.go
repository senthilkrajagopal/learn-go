package senthil

// {1, 4, 4, 7}, 2

func TwoSum(slice []int, target int) [2]int {
	m := make(map[int]int)
	for i := 0; i < len(slice); i++ {
		m[slice[i]] = i
	}

	for i, n := range slice {
		e := target - n
		if m[e] != i {
			return [2]int{i, m[e]}
		}
	}

	return [2]int{0, 0}
}
