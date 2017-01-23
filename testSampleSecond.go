package main
import "fmt"

func main(){
var favNum[10] float64
favNum[0] = 1
favNum[1] = 10
fmt.Println(favNum[0])
favNums := [5]float64{1,2,3,4,5}
for _,value := range favNums{
fmt.Println(value)
}
numSlice := []int {5,4,3,2,1}
numSlice2 := numSlice[3:5]
fmt.Println("NumSLice is ",numSlice2[:2])
//Create an array of int default initialise first five values and set the size t//o 10
numsSlice := make([]int, 5, 10)
 copy(numsSlice,numSlice)
fmt.Println("Array is ",numsSlice[:5])
numsSlice = append(numsSlice,3,2,2)
fmt.Println("ArraYs are ",numsSlice[:7])


//MAPS or key value pairs in GO
 presentAge := make(map[string] int)
presentAge["Kone"]= 1
fmt.Println(presentAge["Kone"])
//Print lenght of Map
fmt.Println(len(presentAge))
//Deleting the value associated with a particular key in Map
delete(presentAge,"Kone")
fmt.Println("Length after deletion is ",len(presentAge))
}

