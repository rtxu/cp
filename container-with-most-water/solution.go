import "math"

func maxArea(height []int) int {
	max := 0
	left, right := 0, len(height)-1

	for left < right {
		area := (right - left) * int(math.Min(float64(height[left]), float64(height[right])))
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
