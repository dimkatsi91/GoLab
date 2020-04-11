package main

import (
    "fmt"
)

// Variadic args example
//
func print(names ... string) {
    for _, val := range names {
        fmt.Println(val)
    }
}


func main() {
    names := []string{"Dimos", "Mathieu", "Jiulia", "Roberto"}
    
    fmt.Println("*****************************************")
    fmt.Printf("Type(names) = %T\n", names)
    fmt.Printf("Type(print) = %T\n", print)
    fmt.Println("*****************************************")
    print(names...)
    fmt.Println("*****************************************")
}
