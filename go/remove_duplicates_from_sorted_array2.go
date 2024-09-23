func removeDuplicates(nums []int) int {
	n := 1
	count := 1
  
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			count = 1
			nums[n] = nums[i]
			n++
		} else {
			count++
			if count == 2 {
				nums[n] = nums[i]
				n++
			}
		}
	}
	return n
}
