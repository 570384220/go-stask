func twoSum(nums []int, target int) []int {
	maps := map[int]int{}
	for i, v := range nums {
		c := target - v
		exist, ok := maps[c]
		if ok {
			a := exist
			res := []int{a, i}
			return res
		}
		maps[v] = i
	}
	return nil
}