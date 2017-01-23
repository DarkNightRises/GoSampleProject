package main
import "fmt"

func main() {
    f()
    fmt.Println("Normaly returned from f")
whatAmI := func(i interface{}) {
    switch i.(type) {
    case bool:
        fmt.Println("Its a bool")

    case int:
        fmt.Println("Its an integer")

    default:
        fmt.Println("Dont know type")
    }
}
whatAmI(1)
whatAmI(true)
}

func f() {
    defer func() {

        if r := recover(); r!=nil {
            fmt.Println("Recovered in f",r)
        }

    }()
    fmt.Println("Calling g")
    g(0)
    fmt.Println("Return normally from g")
}

func g(i int){
    if i > 3 {
        fmt.Println("Panicking")
        panic(fmt.Sprintf("%v", i))

    }

    defer fmt.Println("Defer in g ",i)
    fmt.Println("printing in g ",i)
    g(i+1)
}


//0123Panicking