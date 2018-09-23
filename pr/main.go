package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Circle struct {
	Point  Point `json:"point"`
	Radius int   `json:"radius"`
}

type Wheel struct {
	Circle Circle `json:"circle"`
	Spokes int    `json:"spokes"`
}

func main() {
	var w, w2 Wheel
	var cir = Circle{
		Point:  Point{X: 8, Y: 8},
		Radius: 5,
	}
	w = Wheel{
		Circle: cir,
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
	file, err := os.Open("file.go") // For read access.
	if err != nil {
		fmt.Println(err)
		file, _ = os.Create("file.go")
	}
	enc := json.NewEncoder(file)
	enc.Encode(w)

	dec := json.NewDecoder(file)
	dec.Decode(&w2)
	fmt.Printf("%#v\n", w2)
	//маршалинг в файл, затем из файла
}
