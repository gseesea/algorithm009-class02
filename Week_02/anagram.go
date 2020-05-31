package Week_02

// 判断字符串异位
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}