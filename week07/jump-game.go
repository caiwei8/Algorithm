func canJump(nums []int) bool {
    // f[i]  表示f[i]能否到达i位置

    // 1
    n := len(nums)
    f := make([]bool, n)
    for i :=0; i < n; i++{
        f[i] = false
    }
    f[0] = true

    for i := 1; i < n; i++{
        for j := 1; j < i; j++{
            if f[j] && nums[j] +j >= i{
                f[i] = true
            }
        }
    }
    
    return f[n-1]
}
