func shipWithinDays(weights []int, days int) int {
    // 二分答案
    // 把这个问题转化为：
    // 把weight数组连续分成days段，days段的和的最大值要最小
    // 发现可以和例题的思路一模一样
    // 确定上下界
    var left int
    var right int
    for i :=0; i < len(weights); i++{
        left = max(left,weights[i])
        right = right+weights[i]
    }
    for left < right{
        mid := (left + right) / 2
        if validate(weights,mid,days){
            right = mid
        }else{
            left = mid + 1
        }
    }
    return right

}

func validate(weights []int,size int,days int) bool{
    count := 1
    box := 0
    for _, weight := range weights{
        if box + weight <= size{
            box += weight
        }else{
            count++
            box = weight
        }
    }
    return count <= days
}

func max(i,j int) int{
    if i >= j{
        return i
    }else{
        return j
    }
  }
