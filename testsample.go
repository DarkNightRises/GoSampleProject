package main
import "fmt"

func main(){

fmt.Println("Hello world")
//statiscally typed language
var age int = 40
var favNum float64 = 1.618033

//randNum := 1

fmt.Println(age,favNum)
var numOne = 1.000
var numTwo = .999
fmt.Println(numOne - numTwo)
//Float are not 100% accurate
const pi float64 = 3.14159265

var myName string = "Derek Banas"
fmt.Println(myName)
fmt.Println("Length is ",len(myName))
fmt.Printf("%f \n",pi)
fmt.Printf("%T \n",pi)
//Integer print
fmt.Printf("%d \n",100)
//Binary representation
fmt.Printf("%b \n",100)
//Character representation
fmt.Printf("%c \n",100)
//Loops
for j:=5;j<10;j++ {
fmt.Println(j)
}
//Conditions Test
yourAge := 18
if yourAge>=16 {
fmt.Println("Age greater than 16")
} else if age>=18{
fmt.Println("Age is less 16")
} else {
fmt.Println("Fuck off")
}

}
