func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	//1.MakeSet
	fa := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		fa[i] = i
	}

	//2.find
	var find func(x int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}


	// 3.UnionSet
	var unionSet func(x, y int) bool
	unionSet = func(x, y int) bool {
		x = find(x)
		y = find(y)
        // 不属于同一个并查集，则合并且返回false
		if x != y {
			fa[x] = y
            return false
		}
        // 属于同一个并查集，返回true
        return true
	}

	for _, edge := range edges {
        // 如果edge[0]和edge[1]不属于同一个并查集，则让他们合并
        // 如果属于，说明他们之间已经合并，新的edge产生了环
		if unionSet(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}



