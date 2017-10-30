// concurrency programming 
// fetches many urls concurrently so that the process will take no longer than
// the longest fetch rather than the sum of all the fetch times 

// fetches them in parallel

package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() { 
    start := time.Now()
    ch := make(chan string)
    for _, url := range os.Args[1:] { 
        go fetch(url, ch) // start a go routine
    }
    for range os.Args[1:] { 
        fmt.Println(<-ch) // receive from ch
    }
    fmt.Printf("%2fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil { 
        ch <- fmt.Sprint(err) // send to channel ch 
        return 
    }
    
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close() // don't leake resources
    if err != nil {
        ch <- fmt.Sprintf("While redaing %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}


