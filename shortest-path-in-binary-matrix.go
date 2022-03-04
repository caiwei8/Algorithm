func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
    // 特殊情况的判断
    if grid[0][0] == 1 || grid[n-1][n-1] == 1{
        return -1
    }
    if n == 1 {
        return 1
    }

	var getNum func(i, j int) int
	getNum = func(i, j int) int {
		return i*n + j
	}
	direction := []int{-1, 0, 1}
	q := []int{}
	depth := make(map[int]int, 0)
	depth[0] = 1
	q = append(q, getNum(0, 0))
	for len(q) != 0 {
		curr := q[0]
		currX := q[0] / n
		currY := q[0] % n
		q = q[1:len(q)]
		for _, i := range direction {
			for _, j := range direction {
				if i == 0 && j == 0 {
					continue
				}
				nx := currX + i
				ny := currY + j
                if depth[getNum(nx, ny)] != 0 {
                    continue
                }
				if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == 0 {
					depth[getNum(nx, ny)] = depth[curr] + 1
					q = append(q, getNum(nx, ny))
				}
				if nx == n-1 && ny == n-1 {
					return depth[getNum(nx, ny)]
				}
			}
		}
	}
	return -1
}
