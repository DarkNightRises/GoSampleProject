package main
import "fmt"
import "errors"

func f1(arg int) (int,error){
if arg == 42{
return -1, errors.New("cant work with 42")
}
return arg + 3,nil
}
func (e argError) Error() string{
return fmt.Sprintf("%d - %s",e.arg,e.prob)
}
type argError struct{
arg int
prob string
}
func f2(arg int) (int,error){
if arg == 42{
return -1,&argError{arg,"cant work with it"}
}
return arg+3,nil
}
func main(){

}

