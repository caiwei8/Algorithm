func reverseOnlyLetters(s string) string {
	byte_s :=[]byte(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !((byte_s[left] >= 'a' && byte_s[left] <= 'z') || (byte_s[left] >= 'A' && byte_s[left] <= 'Z')) {
			left++
		}
		for left < right && !((byte_s[right] >= 'a' && byte_s[right] <= 'z') || (byte_s[right] >= 'A' && byte_s[right] <= 'Z')) {
			right--
		}
		if left < right {
			byte_s[left], byte_s[right] = byte_s[right], byte_s[left]
			left++
			right--
		}
	}
	return string(byte_s)
}
