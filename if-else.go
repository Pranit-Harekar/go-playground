package main

import "fmt"

func main() {
    
    if 7%2 == 0  {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    if 1==1 {
        fmt.Println("true")
    }
    
    if num := 9 ; num > 0 {
        fmt.Println(num, "is positive")
    } else if num == 0 {
        fmt.Println(num, "is zero")
    } else {
        fmt.Println(num, "is negative")
    }
}
