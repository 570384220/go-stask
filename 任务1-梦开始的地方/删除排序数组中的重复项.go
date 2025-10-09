func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[j-1] != nums[j] {
			nums[i] = nums[j]
			i++
		} else {
			continue
		}
	}
	return i
}