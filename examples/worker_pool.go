//Worker to run several work concurrently on work pool
package main
import "fmt"
import "time"
func worker(id int, jobs <-chan int, results chan<- int){
for i := range jobs {
fmt.Println("worker",id," started job ",i)
time.Sleep(time.Second)
fmt.Println("worker",id,"finished job ",i)
results <- i*2 
}
}
func main() {
jobs := make(chan int,100)
results := make(chan int,100)
for w:=1 ; w<=3; w++{
go worker(w,jobs,results)
}
for i:=1;i<=5;i++{
jobs <- i
}
close(jobs)
for j:=1 ; j<=5; j++{
<-results
}
} 
