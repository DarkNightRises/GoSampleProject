package main
import "fmt"
func main() {
listNums := []float64{1,2,3,4,5}
fmt.Println("The Sum is ", addTheNumber(listNums))
num1,num2 := multipleReturn(2)
fmt.Println(num1, num2)
//nums := []float6{1,2,3,4,5,6}
fmt.Println("Final subtracted value is ",subtract(1,2,3,4,5))



//Function inside a function
num3 := 5

doubleNum := func() int{
num3 *= 5
return num3
}
fmt.Println(doubleNum())
fmt.Println(doubleNum())
fmt.Println(factorial(5))
}
//Recursion
func factorial(num int) int {
if num == 1 {
return 1
} else {
return num*factorial(num-1) 
}
}

//Simple function demonstration
func addTheNumber(numbers []float64) float64{
sum := 0.0
for _, value := range numbers {
sum = sum+value
}

return sum
}


//Var args
func subtract(args ...int) int {
finalValue := 0
for _,value := range args {
finalValue -= value
}
return finalValue

}

//Multiple return
func multipleReturn(number int) (int, int) {
return number+1, number-1
}

