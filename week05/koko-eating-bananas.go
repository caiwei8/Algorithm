func minEatingSpeed(piles []int, h int) int {
    // 二分答案
    // 确定上下界
    // 上界是piles中的最大值

    var left int
    var right int
    for i :=0; i < len(piles); i++{
        right = max(right,piles[i])
    }
    for left < right{
        mid := (left + right) / 2
        if 0 == mid { //如果mid == 0，则每小时吃<1数量也能吃完，则返回1，避免除0的出现  
            return 1
        }
        if validate(piles,mid,h){
            right = mid 
        }else{
            left = mid + 1
        }
    }
    return right
}

func validate(piles []int,k int,hours int) bool{
    var times int
    for _, num := range piles{
        // if num > k{
        //     count := num / k
        //     times +=count
        //     times++
        // }else{
        //     times++
        // }
       times += num / k
		if num%k > 0 {
			times++
		}
    }
    return times <= hours
}

func max(i,j int) int{
    if i >= j{
        return i
    }else{
        return j
    }
}
