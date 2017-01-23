package main
import "fmt"
import "math"
//For struct and interfaces
func main() {
	rect := Rectangle{leftX: 0, topY: 50, width: 10, height: 15}

	//Or if we know the order
	rectnew := Rectangle{0, 50, 10, 15}
	circle := Circle{10}
	fmt.Println("Rectangle is ",rect.topY," and height is ",rectnew.height)
	fmt.Println("area is ",getArea(rect))
	fmt.Println("circle area is ",getArea(circle))
}

type Rectangle struct{
	leftX float64
	 topY float64
	 width float64
	 height float64
}

func (rect Rectangle) area()  float64{
	return rect.width * rect.height
}

type Shape interface {
	area() float64
}

type Circle struct{
	radius float64
}

func (circle Circle) area() float64{
	return math.Pi * math.Pow(circle.radius, 2)
}

func getArea(shape Shape) float64{
	return shape.area()
}