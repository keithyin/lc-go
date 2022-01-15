package algos

func NumIslands(grid [][]byte) int {
	visited := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[0]))
	}

	cnt := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && visited[i][j] == 0 {
				Visit(grid, visited, i, j)
				cnt += 1
			}
		}
	}
	return cnt
}

func Visit(grid [][]byte, visited [][]int, i int, j int) {
	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || grid[i][j] == '0' || visited[i][j] == 1 {
		return
	}

	visited[i][j] = 1

	Visit(grid, visited, i, j+1)
	Visit(grid, visited, i, j-1)
	Visit(grid, visited, i+1, j)
	Visit(grid, visited, i-1, j)

}
