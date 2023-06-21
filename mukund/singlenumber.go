package main

func SingleNumberMap(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = struct{}{}
			continue
		}
		delete(m, num)

	}
	res := 0
	for k, _ := range m {
		res = k
	}
	return res
}

func SingleNumberXOR(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}

	return nums[0]
}
