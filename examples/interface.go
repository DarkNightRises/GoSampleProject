package main
import "fmt"
import "math"
type shape interface{
area() float64
perim() int
}
type rect struct{
width, height int
}
type circle struct{
radius float64
}

func (r rect) area() float64{
return float64(r.width*r.height)
}
func (c circle) area() float64{
	return math.Pi * math.Pow(c.radius, 2)
}
func (r rect) perim() int{
return 2*(r.width+r.height)
}
func (c circle) perim() int{
return int(2*math.Pi*c.radius)
}
func getArea(s shape) float64{
return s.area()
}
func getPerim(s shape) int{
return s.perim()
}
func main() {
r := rect{10,20}
c := circle{5}
fmt.Println(getArea(r))
fmt.Println(getArea(c))
fmt.Println(getPerim(r))
fmt.Println(getPerim(c))
}
