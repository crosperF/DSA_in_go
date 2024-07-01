package main

import "fmt"

func main() {
	arr := []int{9, 8, 2, 5, 3, 2}
	ans := selection_sort(arr)
	fmt.Println(ans)
}

func swap(arr *[]int, x int, y int) {
	temp := (*arr)[x]
	(*arr)[x] = (*arr)[y]
	(*arr)[y] = temp
}

func selection_sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		smallest := i

		for j := i; j < len(arr); j++ {
			if arr[j] < arr[smallest] {
				smallest = j
			}
		}

		swap(&arr, i, smallest)
	}

	return arr
}
