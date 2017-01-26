//timer are used to trigger an event once in a future
package main
import "fmt"
import "time"
func main() {
timer1 := time.NewTimer(time.Second*2)
<-timer1.C
fmt.Println("timer 1 expirer")
timer2 := time.NewTimer(time.Second*3)
go func() {
	<-timer2.C
	fmt.Println("timer2 expired")
}()
stop2 := timer2.Stop()
if stop2 {
fmt.Println("Timer 2 stopped")
}
}
