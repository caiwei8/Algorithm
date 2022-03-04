func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    n := len(strs)
    var minLenthWord string
    minWord := math.MaxInt8
    for i := 0; i < n; i++ {
        if len(strs[i]) < minWord {
            minWord = len(strs[i])
            minLenthWord = strs[i]
        }
    }
    for i := 0; i < minWord; i++ {
        for j := 0; j < n; j++ {
            if strs[j][i] != minLenthWord[i] {
                return minLenthWord[:i]
            }
        } 
    }
    return minLenthWord
  }
