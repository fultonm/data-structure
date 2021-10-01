package main

import (
	"container/list"
	"errors"
	"fmt"
	"main/helper"
)

func main() {
	//l := HashTable[int].New()
	// l.PushFront(1.2)
	// myString := l.Front().Value
	//fmt.Printf("max sub array: %v\n", maxSubArray([]int{7, -2, 3, -2, 4}))

	image := [][]int{{0, 0, 0}, {0, 1, 1}}
	sr, sc := 1, 1
	newColor := 1
	fmt.Printf("%v\n", floodFill(image, sr, sc, newColor))

}

// 0 0 0
// 0 1 1

// 2 3 1
// 2 2 0

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	numRow := len(image)
	if numRow == 0 {
		return image
	}
	numCol := len(image[0])
	if numCol == 0 {
		return image
	}
	if sr >= numRow || sc >= numCol || sr < 0 || sc < 0 {
		return image
	}
	ogColor := image[sr][sc]
	if PixelIsColor(Pixel{sr, sc}, image, newColor) {
		return image
	}
	s := list.New()
	Push(s, Pixel{sr, sc})
	for s.Len() > 0 {
		cur, _ := Pop(s)
		nor, est, sou, wst := Pixel{cur.r - 1, cur.c}, Pixel{cur.r, cur.c + 1}, Pixel{cur.r + 1, cur.c}, Pixel{cur.r, cur.c - 1}
		if PixelInBounds(nor, numRow, numCol) && PixelIsColor(nor, image, ogColor) {
			Push(s, nor)
		}
		if PixelInBounds(est, numRow, numCol) && PixelIsColor(est, image, ogColor) {
			Push(s, est)
		}
		if PixelInBounds(sou, numRow, numCol) && PixelIsColor(sou, image, ogColor) {
			Push(s, sou)
		}
		if PixelInBounds(wst, numRow, numCol) && PixelIsColor(wst, image, ogColor) {
			Push(s, wst)
		}
		image[cur.r][cur.c] = newColor
	}
	return image
}

type Pixel struct {
	r, c int
}

func Pop(s *list.List) (Pixel, error) {
	if s.Len() == 0 {
		return Pixel{}, errors.New("this stack is empty")
	}
	r := s.Back().Value.(Pixel)
	s.Remove(s.Back())
	return r, nil
}

func Push(s *list.List, e Pixel) {
	s.PushBack(e)
}

func PixelInBounds(p Pixel, numRow int, numCol int) bool {
	if 0 <= p.c && p.c < numCol && 0 <= p.r && p.r < numRow {
		return true
	}
	return false
}

func PixelIsColor(p Pixel, img [][]int, color int) bool {
	return img[p.r][p.c] == color
}

func QueueContainsPixel(q *list.List, p Pixel) bool {
	cur := q.Front()
	for i := 1; i < q.Len(); i++ {
		e := cur.Value.(Pixel)
		if e.r == p.r && e.c == p.c {
			return true
		}
		cur = cur.Next()
	}
	return false
}

func maxSubArray(nums []int) int {
	n := len(nums)
	// dp[i] denotes maximum subarray till the index i
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = helper.Max(dp[i-1]+nums[i], nums[i])
		}
	}

	// Find the maximum subarray sum
	max := helper.MinInt32()
	for _, e := range dp {
		if e > max {
			max = e
		}
	}

	return max
}
