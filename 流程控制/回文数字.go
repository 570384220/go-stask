func huiwen(num int) (res bool) {
	a := strconv.Itoa(num)
	arr := a[0:]

	res = true
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] == arr[j] {
			continue
		} else {
			res = false
		}
	}

	return res
}