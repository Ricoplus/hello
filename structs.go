package main

import "fmt"

type person struct{
    name string
    age int
}

func main(){
    fmt.Println(person{"Bob",20})

    fmt.Println(person{name:"rico",age:20})

    fmt.Println(person{name:"Fred"})

    fmt.Println(&person{name:"Ann",age:40})

    s := person{name:"Hope",age:21}

    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 1

    fmt.Println(s.age)

}
