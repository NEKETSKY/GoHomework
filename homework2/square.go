package main

import "fmt"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (z Square) End() Point {
	return Point{z.start.x + int(z.a), z.start.y - int(z.a)}
}

func (z Square) Perimeter() uint {
	return z.a * 4
}
func (z Square) Area() uint {
	return z.a * z.a
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
