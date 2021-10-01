package helper

func ContainsString(s string, r rune) int {
	for i, e := range []rune(s) {
		if e == r {
			return i
		}
	}
	return -1
}

func ContainsRune(arr []rune, r rune) int {
	for i, e := range arr {
		if e == r {
			return i
		}
	}
	return -1
}

func Sum(args ...int) int {
	r := 0
	for _, e := range args {
		r += e
	}
	return r
}

func Max(args ...int) int {
	r := MinInt32()
	for _, e := range args {
		if e > r {
			r = e
		}
	}
	return r
}

func MinInt32() int {
	return -int(^uint(0)>>1) - 1
}

func Min(args ...int) int {
	r := MaxInt32()
	for _, e := range args {
		if e < r {
			r = e
		}
	}
	return r
}

func MaxInt32() int {
	return int(^uint(0) >> 1)
}

func SelectionSort(arr []int) {
	findMinIdx := func(s int, arr []int) int {
		min := arr[s]
		idx := -1
		for i, e := range arr {
			if e < min {
				min = e
				idx = i
			}
		}
		return idx
	}

	size := len(arr)
	for i := 0; i < size; i++ {
		minIdx := findMinIdx(i, arr)
		if minIdx > -1 {
			tmp := arr[i]
			arr[i] = arr[minIdx]
			arr[minIdx] = tmp
		}
	}
}
