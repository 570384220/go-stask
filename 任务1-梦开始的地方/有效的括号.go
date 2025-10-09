func isValid(s string) bool {
    
	kMap := map[string]string{"{": "}", "[": "]", "(": ")"}
	arr := []string{}

	runArr := []rune(s)

	for _, v := range runArr {
		z := string(v)
		if z == "(" || z == "[" || z == "{" {
			arr = append(arr, z)
		} else {
			top := arr[len(arr) -1]
			if cap(arr) == 1 {
				arr = []string{}
			} else {
				arr = arr[0: len(arr) - 1]
			}
			if kMap[top] == z {
				continue
			} else {
				return false
			}
		}
	}
	if len(arr) != 0 {
		return false
	}

	return true
}