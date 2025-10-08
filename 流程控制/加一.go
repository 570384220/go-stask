func plusOne(digits []int) []int {
	jin := 1
	for i := len(digits) - 1; i>=0; i-- {
		digits[i] += jin
		if(digits[i] > 9) {
			digits[i] = digits[i] % 10
			jin = 1
		} else {
			jin = 0
		}
	}
	
	if(jin > 0) {
		var newDigits = []int{}
		newDigits = append(newDigits, 1)
		newDigits = append(newDigits, digits...)
		return newDigits
	} else {
		return digits
	}
}