package main
import "fmt"
func main() {
var a[5] int
fmt.Println("emp:",a)
a[4] = 3
//b := [5]int{1,2,3,4,5}
var two2D [3][3]int

for i :=0 ; i<3; i++ {
for j :=0; j<3; j++ {
	two2D[i][j]=i+j
}
}
fmt.Println("2 d array ",two2D)
}
