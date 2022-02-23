func numIsland(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	fa := make([]int, m*n)
	for i := 0; i < m*n; i++ {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	var unionset func(x, y int)
	unionset = func(x, y int) {
		x = fa[x]
		y = fa[y]
		if x != y {
			fa[x] = y
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				if i > 0 && grid[i-1][j] == '1' {
					unionset(i*n+j, (i-1)*n+j)
				}
				if j > 0 && grid[i][j-1] == '1' {
					unionset(i*n+j, i*n+j-1)
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if fa[i*n+j] != i*n+j {
				ans--
			}
		}
	}
	return ans
}
