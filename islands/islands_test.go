package islands

import (
	"math/rand"
	"testing"
)

func Test1(t *testing.T) {
	h, w := 15000, 20000
	grid := [][]int{}
	for i := 0; i < h; i++ {
		grid = append(grid, []int{})
		for j := 0; j < w; j++ {
			grid[i] = append(grid[i], rand.Intn(2))
		}
	}

	t.Logf("max island: %v\n", MaxAreaOfIsland(grid[:][:]))
}

func Test2(t *testing.T) {
	grid := [][]int{{1, 1}, {1, 1}}
	result := MaxAreaOfIsland(grid)
	if result != 4 {
		t.Fatalf("Area of island was not 4 but %v", result)
	}
}

func Test3(t *testing.T) {
	grid := [][]int{
		{1, 0, 1},
		{1, 1, 1},
		{0, 0, 1}}

	result := MaxAreaOfIsland(grid)

	if result != 6 {
		t.Fatalf("Area of island should be 6, got %v", result)
	}

}
