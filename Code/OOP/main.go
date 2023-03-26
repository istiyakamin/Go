package main

import (
	"fmt"
	"math"
)

type GeomatryFormula interface {
	area() float64
}

type Ranctangle struct {
	height float64
	width float64
}

type Circle struct {
	radious float64
}

type Triangle struct {
	base float64
	height float64
}

// Receving Function
func (R Ranctangle) area() float64 {
	return R.height*R.width
}

func (T Triangle) area() float64 {
	return (T.height*T.base)/2
}

func (C Circle) area() float64 {
	return math.Pi*C.radious*C.radious
}

func geomatry_area(g GeomatryFormula) float64 {
	return g.area()
}

func main(){

	ract := Ranctangle{
		height: 10,
		width: 4,
	}
	ract_result := geomatry_area(ract)

	tri := Triangle{
		base: 10,
		height: 4,
	}
	triangle_result := geomatry_area(tri)
	
	circle := Circle{
		radious: 10,
	}
	cicrcle_result := geomatry_area(circle)

	fmt.Println("Ractangle Area: ->", ract_result)
	fmt.Println("Triangle Area: ->", triangle_result)
	fmt.Println("Cicle Area: ->", cicrcle_result)

}