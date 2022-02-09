func climbStairs(n int) int {
    f1, f2, f3 := 0, 0, 1
    for i := 1; i <= n; i++ {
        f1 = f2
        f2 = f3
        f3 = f1 + f2
    }
    return f3
}
