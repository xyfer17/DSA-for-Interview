func isPalindrome(x int) bool {

	// x if negative number
	if x < 0 {
		return false
	}

	num := x
	reversed := 0

	for num != 0 {
		reversed = 10*reversed + num%10
		num /= 10
	}

	return x == reversed
}