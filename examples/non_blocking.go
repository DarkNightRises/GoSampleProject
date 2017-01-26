package main
import "fmt"
func main() {
messages := make(chan string)
signal := make(chan string)
go func() {
messages <- "hellobro"
}()
select {
case  <- messages:
fmt.Println("Printing messa")
default:
fmt.Println("No message yet")
}
msg := "hi"
select {
case messages <- msg:
fmt.Println("No activity still",msg)
default:
fmt.Println("Still blank")
}
select {
case msg := <- messages:
fmt.Println("received messages",msg)
case sig := <-signal:
fmt.Println("Received signal",sig)
default:
fmt.Println("No activity at all")
}
}
