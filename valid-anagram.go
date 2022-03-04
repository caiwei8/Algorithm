func isAnagram(s string, t string) bool {
	sSlice := [26]int{}
  	tSlice := [26]int{}
	for i := 0; i < len(s); i++ {
		sSlice[s[i]- 'a']++
	}
	for i := 0; i < len(t); i++ {
		tSlice[t[i]- 'a']++
	}

	return sSlice == tSlice
}
