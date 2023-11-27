package main

import (
    "fmt"
)

func main() {
    var a, b int
    fmt.Print("Enter the value of a: ")
    fmt.Scanf("%d", &a)
    fmt.Print("Enter the value of b: ")
    fmt.Scanf("%d", &b)

    // Swap the values using addition
    a = a + b
    b = a - b
    a = a - b

    fmt.Println("After swap:")
    fmt.Println("a =", a)
    fmt.Println("b =", b)
}
