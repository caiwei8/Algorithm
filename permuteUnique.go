func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}    
    used := make([]bool, len(nums))    
    var find func(path []int, depth int)    
    find = func(path []int, depth int) {        
        if len(nums) == depth {            
            tmp := make([]int, len(path))            
            copy(tmp ,path)            
            res = append(res, tmp)            
            return        
            }        
            for i := 0; i < len(nums); i++ {           
                if used[i] || (i > 0 && !used[i-1] && nums[i] == nums[i-1]){
                    continue
                    }            
                    used[i] = true;            
                    path = append(path, nums[i])            
                    find(path,depth+1)            
                    path = path[:len(path) - 1]            
                    used[i] = false;        
                    }    
                }    
                find([]int{}, 0)    
                return res
}
