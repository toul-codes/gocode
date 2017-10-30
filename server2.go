// A minimal "echo" and counter server 

package main
import (
    "fmt"
    "log"
    "net/http"
    "sync"
) 

var mu sync.Mutex
var count int 

func main() { 
    http.HandleFunc("/", handler)
    http.HandleFunc("/count", counter)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the path component of the requested URL 
func handler (w http.ResponseWriter, r *http.Request) { 
    mu.Lock()
    count++
    mu.Unlock()
    fmt.Fprint(w, "URL.Path = %q\n", r.URL.Path) 
} 

//counter echoes the numbe of calls so far 

func counter(w http.ResponseWriter, r *http.Request) {
    mu.lock()
    fmt.Fprint(w, "Count %d\n", count)
    mu.Unlock()
} 

