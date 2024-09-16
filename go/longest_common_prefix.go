func longestCommonPrefix(strs []string) string {

	var prefix bytes.Buffer
	str := strs[0]

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(str) {
			str = strs[i]
		}
	}

	for i := 0; i < len(str); i++ {
		to_add := true
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != str[i] {
				to_add = false
			}
		}
		if to_add {
			prefix.WriteByte(str[i])
		} else {
			break
		}
	}

	return prefix.String()
}