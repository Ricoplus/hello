package main
import "fmt"


func f(from string){
    for i:=0;i<20;i++{
        fmt.Println(from,":",i)
    }
}

func main(){
    f("direct")

    go f("goroutine1")
    go f("goroutine2")
    go f("goroutine3")
    go f("goroutine4")
    go f("goroutine5")
/*
    go func(msg string){
        fmt.Println(msg)
    }("going")
*/
    var input string

    fmt.Scanln(&input)
    fmt.Println("done")
}
