package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}
func getPos(arr []int, index int) int {

	return index % len(arr)
}
func rotate(slice *[]int, n int) {
	tmp := make([]int, len(*slice))
	for index := 0; index < len(*slice); index++ {
		jndex := index + n
		tmp[getPos(*slice, jndex)] = (*slice)[getPos(*slice, index)]
	}
	*slice = tmp
}
func reverse(arr *[5]int) {
	for index := 0; index < len(arr)/2; index++ {
		swap(&arr[index], &arr[len(arr)-index-1])

	}
}
func main() {
	var b = [...]int{10, 1, 2, 3, 4, 5}
	slice := b[:]
	rotate(&slice, 12)
	fmt.Println(slice)
}
