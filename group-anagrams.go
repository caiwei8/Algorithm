func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	ans := [][]string{}
	ansMap := make(map[string][]string)

	for _, word := range strs {
		chars := []rune(word)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		ansMap[string(chars)] = append(ansMap[string(chars)], word)
	}
	for _, v := range ansMap {
		ans = append(ans, v)
	}
	return ans
}
