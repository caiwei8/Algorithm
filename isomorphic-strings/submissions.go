func isIsomorphic(s string, t string) bool {
    sTot := make(map[byte]byte, 0)
    tTos := make(map[byte]byte, 0)
    for i := 0; i < len(s); i++ {
        if (sTot[s[i]] > 0 && sTot[s[i]] != t[i]) || (tTos[t[i]] > 0 && tTos[t[i]] != s[i]) {
            return false
        }
        sTot[s[i]] = t[i]
        tTos[t[i]] = s[i]
    }
    return true
}
