func trap(height []int) int {
	getSum := func(h []int) int {
		sum := 0
		for _, v := range h {
			sum += v
		}
		return sum
	}
	sum1 := getSum(height)

	maxi, maxv := -1, 0
	for i, v := range height {
		if v > maxv {
			maxv = v
			maxi = i
		}
	}

	for i, h := 0, 0; i < maxi; i++ {
		if height[i] > h {
			h = height[i]
		} else {
			height[i] = h
		}
	}

	for i, h := len(height)-1, 0; i > maxi; i-- {
		if height[i] > h {
			h = height[i]
		} else {
			height[i] = h
		}
	}

	return getSum(height) - sum1
}
