package main
import "fmt"
func main() {
fmt.Println("Sum is ",plus(2,3))
a,b := multipleReturn(5,7)
fmt.Println("Multiple return is ",a," and b is ",b)
fmt.Println("Sum is ",sum(1,2,3))
}
func plus(a, b int)int{
return a + b
}
//Multiple returns 
func multipleReturn(a,b int) (int,int){
return b,a
}
//Variadic Functions
func sum(args... int)int{
sum := 0
for i := 0;i< len(args); i++{
sum += args[i]
}
return sum
}
