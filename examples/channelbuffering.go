package main
import "fmt"
//default channels only can send a data when they have a receiver associated 
//But buffered channels can receive any number of data without even existence 
// of the associated receiver...
func main() {
message := make(chan string, 2)
message <- "buffered"
message <- "channel"
fmt.Println(message)
fmt.Println(message)

}
