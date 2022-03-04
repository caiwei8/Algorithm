func toLowerCase(s string) string {
    var ans string
    for i := 0; i < len(s); i++ {
        if s[i] <= 'Z' && s[i] >= 'A'{
            ans += string(s[i] + 32)
        }else {
            ans += string(s[i])
        }
    }
    return ans
  }

