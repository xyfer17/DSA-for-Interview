import (
	"strconv"
)

func countAndSay(n int) string {

	if n == 1 {
		return "1"
	}

	res := helper(countAndSay(n - 1))

	return res
}

func helper(str string) string {

	n := len(str)
	cnt := 1
	var res = ""
	st := str[0]

	for i := 1; i < n; i++ {
		if st != str[i] {
			res = res + strconv.Itoa(cnt) + string(st)
			st = str[i]
			cnt = 1
		} else {
			cnt++
		}
	}

	res = res + strconv.Itoa(cnt) + string(st)

	return res
}