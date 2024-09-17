func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {

		var minHeight int

		if height[left] < height[right] {
			minHeight = height[left]
		} else {
			minHeight = height[right]
		}

		if maxArea < minHeight*(right-left) {
			maxArea = minHeight * (right - left)
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
