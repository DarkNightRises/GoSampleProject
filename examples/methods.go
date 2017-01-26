package main
import "fmt"
type rect struct {
 width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int{
	return (r.width + r.height)*2
}
func main() {
r := rect{5,10}
fmt.Println(r.width)
fmt.Println(r.area())
fmt.Println(r.perim())
}
