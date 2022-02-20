func numSquares(n int) int {
    // f[i]  表示需要f[i]个数来表示和为n的完全平方数的最少数量
    // f[i] = 1+f[i-j*j]

    // 1.
    f := make([]int, n+1)
    for i := 1; i <= n; i++{
        f[i] = math.MaxInt8
    }
    f[0] = 0

    //2.
    for i := 1; i <= n; i++{
        for j := 1; j * j <= i; j++{
            f[i] = min(f[i],1+f[i-j*j])
        }
    }
    return f[n]
}

func min(i,j int) int{
    if i < j{
        return i
    }
    return j
}
