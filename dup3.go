// dup 3 prints the count and the text of lines that appear smore than onec in
// the named input ifles 
// reads the entire input into memory in one big gulp
// splits lines 
// processes the lines 

package main ( 
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)
func main() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] { 
        data, err := ioutil.Readfile(filename)
        if err != nil { 
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") { 
            counts[line]++
        }
    }
    for line, n := range counts { 
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}



