package main
import "fmt"
type person struct{
name string
age int
}
func main() {
persons := person{"Kone",22}
fmt.Println(persons)
fmt.Println(persons.name)
fmt.Println(persons.age)
}
