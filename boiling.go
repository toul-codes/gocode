package main

import "fmt"

const BOILINGF = 212.0

func main() {
    var f = BOILINGF
    var c = (f-32) * 5 / 9 
    fmt.Printf("Boiling point = %g°F or %g°C\n", f, c)
}

