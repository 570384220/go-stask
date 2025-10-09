
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else {
			return false
		}
	})

	arr := [][]int{}
	arr = append(arr, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := arr[len(arr)-1]
		fast1 := last[1]

		if fast1 >= intervals[i][0] {
			if fast1 <= intervals[i][1] {
				last[1] = intervals[i][1]
			}
		} else {
			arr = append(arr, intervals[i])
		}
	}
	return arr
}
