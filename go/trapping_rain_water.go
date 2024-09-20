// approach 1 : O(n) time, O(n) space
func trap_1(height []int) int {

	n := len(height)
	left := make([]int, n)
	right := make([]int, n)
	res := 0
	left[0] = height[0]
	right[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		if height[i] >= left[i-1] {
			left[i] = height[i]
		} else {
			left[i] = left[i-1]
		}
	}

	for i := n - 2; i >= 0; i-- {
		if height[i] >= right[i+1] {
			right[i] = height[i]
		} else {
			right[i] = right[i+1]
		}
	}

	for i := 0; i < n; i++ {
		if left[i] < right[i] {
			res += left[i] - height[i]
		} else {
			res += right[i] - height[i]
		}
	}

	return res
}
