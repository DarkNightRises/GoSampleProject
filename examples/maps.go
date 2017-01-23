package main
import "fmt"
func main() {
m := make(map[string]int)
m["k1"] = 2
m["k2"] = 3
fmt.Println("map: ",m)
v1 := m["k1"]
fmt.Println("V1 ",v1)
fmt.Println("V2 ",m["k2"])
fmt.Println("Len ",len(m))
delete(m, "k2") 
fmt.Println("M after deletion is ",m)
_, prs := m["k2"]
fmt.Println("prs: ",prs)
n := map[string]int{"foo":1, "bar": 2}
fmt.Println("map: ",n) 
}
