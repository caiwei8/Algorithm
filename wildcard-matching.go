func isMatch(s string, p string) bool {
    n := len(s)
    m := len(p)

    s = " " + s
    p = " " + p

    f := make([][]bool, n + 1)
    for i := 0; i < n + 1; i++ {
        f[i] = make([]bool, m + 1)
    }

    f[0][0] = true
     for i := 1; i <= m; i++ {
        if p[i] == '*' {
            f[0][i] = true
        } else {
            break
        }
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if p[j] <= 'z' && p[j] >= 'a' && s[i] == p[j] {
                f[i][j] = f[i-1][j-1]
            }
            if p[j] == '?' {
                f[i][j] = f[i-1][j-1]
            }
            if p[j] == '*' {
                f[i][j] = f[i][j-1] || f[i-1][j]
            }
        }
    }
    return f[n][m]
  }
