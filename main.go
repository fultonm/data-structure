package main

import (
	"fmt"
	"main/islands"
	"math/rand"
	"time"
)

func main() {
	h, w := 4000, 4000
	grid := [][]int{}
	min:= 9_000_000
	max := 0
	for n := 0; n < 100; n++ {
	grid = [][]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < h; i++ {
		grid = append(grid, []int{})
		for j := 0; j < w; j++ {
			grid[i] = append(grid[i], rand.Intn(2))
		}
	}


	result := islands.MaxAreaOfIsland(grid)
	
	if result < min { min = result }
	if result > max { max = result }
		}

	fmt.Printf("Total grid area: %v M, max Max island area: %v, Min max island area:%v\n", float64(h*w)/1_000_000.0, max, min)
}
