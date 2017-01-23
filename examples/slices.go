package main
import "fmt"
func main() {
s := make([]string, 3)
fmt.Println("emp",s)
s[0] = "a"
s[1] = "b"
s[2] = "c"
fmt.Println("Array after fullfilment ",s)
s = append(s, "d", "e", "f")
c := make([]string,len(s))
copy(c,s)
fmt.Println("Copied array ",c)
l := s[2:5]
fmt.Println("l is sliced from s ",l)
twoD := make([][]int, 3)
for i := 0; i < 3; i++ {
innerLen := i + 1
twoD[i] = make([]int,innerLen)
for j := 0; j <innerLen ; j++ {
	twoD[i][j] = i+j 
}
}
fmt.Println("2D: ",twoD) 
}

