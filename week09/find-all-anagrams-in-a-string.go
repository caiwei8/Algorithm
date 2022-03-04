

func findAnagrams(s, p string) (ans []int) {
	m, n := len(s), len(p)
	if m < n {
		return
	}

	var s_arry, P_arry [26]int
	for i, ch := range p {
		s_arry[s[i]-'a']++
		P_arry[ch-'a']++
	}
	if s_arry == P_arry {
		ans = append(ans, 0)
	}

	for i, ch := range s[:m-n] {
		s_arry[ch-'a']--
		s_arry[s[i+n]-'a']++
		if s_arry == P_arry {
			ans = append(ans, i+1)
		}
	}
	return
}



