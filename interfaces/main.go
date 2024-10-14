package main

import (
    "fmt"
    "math"
)

// lets define the interface
type Shape interface {
    area() float64
}

// define the circle class or properties
type Circle struct {
    x, y, radius float64
}

// define the rectangle class or properties
type Rectangle struct {
    width, height float64
}

// calculate on how we get the area of the circle 
func (circle Circle) area() float64{
    return math.Pi * circle.radius * circle.radius
}

// calculate on how we get the area of the rectangle 
func (rectangle Rectangle) area() float64 {
    return rectangle.width * rectangle.height

}

// function to call the get area function of the passed in shape
func getArea(shape Shape) float64{
    return shape.area()
}

func main(){
    // create our circle and the rectangle with the values provided
    circleA := Circle{x: 0, y:0, radius: 5}
    rectangleA := Rectangle{width: 10, height: 5}

    // print the areas
    fmt.Printf("Circle Area: %f\n", getArea(circleA))
    fmt.Printf("Rectangle Area: %f\n", getArea(rectangleA))
}
