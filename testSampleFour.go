package main
import "fmt"
func main() {
//Using with defer function printOne() only executes when main finishes
	defer printOne()
	defer printTwo()
//Use of recover of safety execution even on error 
	fmt.Println(safeDiv(2,0))
	fmt.Println(safeDiv(6,2))

//Use of pointer to reflect change local paramter with original paramter
		num1 := 5
		num2 := 8
		swapValue(&num1,&num2)
		fmt.Println(num1,num2)
//Catching exception

	demPanic()
	
}
func printOne(){
	fmt.Println(1)
}
func printTwo(){
	fmt.Println(2)
}
func safeDiv(num1, num2 int) int{
	defer func(){
//To throw an error
//fmt.Println(recover())
//To not throw an error and just execute
		recover()

		}()
		solution := num1/num2
		return solution
	}


func demPanic(){
	defer func(){
		fmt.Println(recover())
	}()
	panic("Feeling panice")
}

//Pointer in GO for local variable affect on global parameter

func swapValue(x,y *int){
	temp := *x
	*x = *y
	*y = temp
}
