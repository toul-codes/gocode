// Echo prints its command-line arguments 
package main 

import(
    "fmt"
    "os"
)

func main() { // declares two variables of the type string  
    var s, sep string
    for i:= 1; i < len(os.Args); i++ { 
        s += sep + os.Args[i] // concatenate the old value of s with the new value of s 
        sep = " " 
    }
    fmt.Println(s)
}


