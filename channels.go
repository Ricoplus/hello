package main

import "fmt"

func main(){
    //message只是一个地址
    messages := make(chan string)

    fmt.Println("messages:",messages)
    go func(){messages <- "ping"}()


    msg := <-messages

    fmt.Println(msg)
}
