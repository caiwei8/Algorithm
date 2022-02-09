func findNumberOfLIS(nums []int) int {
    n := len(nums)
    var ans = 0
    var maxLenth = 1
    f := make([]int,n)
    // 存储子序列的个数
    count := make([]int,n)
    for i := 0; i < n; i++{
        f[i] = 1
        count[i] = 1
        for j := 0; j < i; j++{
            if nums[j] < nums[i]{
                if f[i] < f[j] + 1{
                    f[i] = f[j] + 1
                    count[i] = count[j]
                }else if f[i] == f[j] + 1{
                    count[i] += count[j]
                }
            }
        }
        maxLenth = max(maxLenth, f[i]);
    }
    for k := 0; k < n; k++ {
            if f[k] == maxLenth{
                ans += count[k]
            }
    }
    return ans;
}

func max(i,j int)int{
    if i >j {
        return i
    }else{
        return j
    }
}



