
func isUgly(num int) bool {
    if num <= 0 {
        return false
    }
	for num != 1 {
		switch {
		case num%2 == 0:
			num /= 2
		case num%3 == 0:
			num /= 3
		case num%5 == 0:
			num /= 5
		default:
			return false
		}
	}
	return true
}

