func lengthOfLastWord(s string) int {
    var ans int
    n := len(s)
    j := n - 1
    for j >= 0 && s[j] == ' ' {
        j--
    }
    for j >= 0 && s[j] != ' ' {
        ans++
        j--
    }
    return ans

}
