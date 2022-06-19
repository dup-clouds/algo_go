package main

import "fmt"

func main() {
	maxSum := -1
	f(0, 0, []int{20, 10, 5, 30, 35}, 5, 100, &maxSum)
	fmt.Println(maxSum)
}

func f(index int, cw int, items []int, n int, w int, maxSum *int) {
	if cw == w || index == n {
		if cw > *maxSum {
			*maxSum = cw
		}
		return
	}
	f(index+1, cw, items, n, w, maxSum)
	if items[index]+cw <= w {
		f(index+1, cw+items[index], items, n, w, maxSum)
	}
}
