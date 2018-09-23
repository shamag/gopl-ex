package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}
func reverse(arr *[5]int) {
	for index := 0; index < len(arr)/2; index++ {
		swap(&arr[index], &arr[len(arr)-index-1])

	}
}
func main() {
	var b = [...]int{1, 2, 3, 4, 5}
	reverse(&b)
	fmt.Println(b)
}
