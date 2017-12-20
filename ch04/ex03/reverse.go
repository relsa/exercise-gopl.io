package main

func reverse(arr *[8]int) {
	for i, j := 0, 7; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
