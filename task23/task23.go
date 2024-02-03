package main

func delete_at_stable(arr []int, i int) []int {
	ret := make([]int, 0)
	ret = append(ret, arr[:i]...)
	return append(ret, arr[i+1:]...)
}

func delete_at_fast(arr []int, i int) []int {
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}
