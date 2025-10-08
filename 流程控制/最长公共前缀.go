func longestCommonPrefix(strs []string) string {

	prefix := strs[0]
	if len(strs) <= 1 {
		return strs[0]
	}
	for i := 1; i < len(strs); i++ {
		prefix = comp(prefix, strs[i])
	}

	return prefix
}

func comp(str1, str2 string) string {
	length := 0
	if len(str1) > len(str2) {
		length = len(str2)
	} else {
		length = len(str1)
	}

	index := 0
	for i := 0; i < length; i++ {
		if str1[i] == str2[i] {
			index++
		} else {
			break
		}
	}

	return str1[:index]
}