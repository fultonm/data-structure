package islands

func MaxAreaOfIsland(grid [][]int) int {
	h, w := len(grid), len(grid[0])
	dir := []Point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	islandArea := func(i int, j int) int {
		grid[i][j] = -1
		area := 0
		s := []Point{{i, j}}
		for len(s) > 0 {
			cur := s[len(s)-1]
			s = s[0 : len(s)-1]
			area++
			for _, e := range dir {
				adj := Point{cur.y + e.y, cur.x + e.x}
				if adj.y >= 0 &&
					adj.y < h &&
					adj.x >= 0 &&
					adj.x < w &&
					grid[adj.y][adj.x] == 1 {
					grid[adj.y][adj.x] = -1
					s = append(s, adj)
				}
			}
		}
		return area
	}

	maxArea := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] > -1 {
				if grid[i][j] == 1 {
					grid[i][j] = -1
					area := islandArea(i, j)
					if area > maxArea {
						maxArea = area
					}
				} else {
					grid[i][j] = -1
				}
			}
		}
	}
	return maxArea

}

type Point struct {
	y int
	x int
}
