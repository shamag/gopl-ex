package main

import "fmt"

type a struct {
	f int16
	b *b
}
type b struct {
	f int16
}

func set(a *a) {
	a.f = 10
	a.b.f = 15
}
func main() {
	var b = new(b)
	b.f = 1
	var a = a{f: 10, b: b}
	set(&a)
	fmt.Printf("%#v  %#v", a, a.b)
}
