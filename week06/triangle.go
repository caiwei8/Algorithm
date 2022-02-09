func minimumTotal(triangle [][]int) int {
    // 构建一个二维数组f[i][j]表示第i行第j个节点时的最小路径
    // f[0][0] = triangle[0][0]
    // 状态转移方程 f[i][j] = min(f[i-1][j],f[i-1][j-1])+triangle[i][j]

    m := len(triangle)
    n := len(triangle[m-1])
    var f = make([][]int,m)
    for i := 0; i < m; i++{
        f[i] = make([]int,n)
    }
    f[0][0] = triangle[0][0]
    for i := 1; i < m; i++{
        f[i][0] = f[i-1][0] + triangle[i][0] // 处理边界的特殊情况
        for j := 1; j < i;j++{
            f[i][j] = min(f[i - 1][j], f[i - 1][j - 1]) + triangle[i][j]
        }
        f[i][i] = f[i - 1][i - 1] + triangle[i][i] // 处理边界的特殊情况
    }
    
    ans := math.MaxInt32
    for i := 0; i < n; i++ {
        ans = min(ans, f[n-1][i])
    }
    return ans
}

func min(i, j int) int {
    if i < j {
        return i
    }
    return j
}
