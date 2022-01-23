func solve(board [][]byte) {
	// 一开始觉得和岛屿数量的题思路很像
	// 但是失败了看了题解才发现思考问题的关键是：如何区别O是不是被包围
  // 考虑问题的角度很关键
  if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m := len(board), len(board[0])
	for i := 0; i < n; i++ {
		// 从边界上的O开始遍历
		dfs(board, i, 0)
		dfs(board, i, m-1)
	}
	for i := 1; i < m-1; i++ {
		dfs(board, 0, i)
		dfs(board, n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}
	n, m := len(board), len(board[0])
	if i < 0 || i >= n || j < 0 || j >= m || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'A'
	for k := 0; k < 4; k++ {
		nx := i + dx[k]
		ny := j + dy[k]
		dfs(board, nx, ny)
	}
}

