package main

import "fmt"

func main() {

    x := 4
    y := 5

    z := add(x, y)

    fmt.Printf("Output: %d\n", z)
}

func add(a int, b int) int {

    return a + b
}
