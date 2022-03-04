func numJewelsInStones(jewels string, stones string) int {
	var ans int
	jewelsMap := make(map[byte]int, 0)
	for i := 0; i < len(jewels); i++ {
		jewelsMap[jewels[i]]++
	}
	for j := 0; j < len(stones); j++ {
		if _, ok := jewelsMap[stones[j]]; ok {
			ans++
		}
	}
	return ans
}
