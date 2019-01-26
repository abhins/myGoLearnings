package main

import "fmt"
import "math"

type geometry interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	r float64
}

type Rec struct {
	w float64
	h float64
}

func main() {

	var i geometry
	myCircle := Circle{4}
	myRec := Rec{2,3}

	i = myCircle
	fmt.Println("Area=", i.Area())
	fmt.Println("Peri=", i.Perimeter())

	i = myRec
	fmt.Println("Area=", i.Area())
	fmt.Println("Peri=", i.Perimeter())

}

func (r Circle) Area() float64 {

	return  math.Pi*(r.r)*(r.r)

}

func (r Circle) Perimeter() float64 {
	return 2*(math.Pi)*(r.r)
}

func (r Rec) Area() float64 {

	return  (r.w)*(r.h)

}

func (r Rec) Perimeter() float64 {
	return 2*(r.w)+2*(r.h)
}
