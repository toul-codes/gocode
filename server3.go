// handler echoes the HTTP request
func handler(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprint(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    
    handler := func(w http.ResponseWriter, r *http.Request) {
        lissajous(w)
    } 
    http.HandleFunc("/", handler)
    for k, v := range r.Header { 
        fmt.Fprint(w, "Header[%q] = %q\n", k, v) 
    }

    fmt.Fprintf(w, "Host = %q\n", r.Host)
    fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
    if err := r.ParseForm(); err != nil { 
        log.Print(err)
    } 

    for k, v := range r.Form{
        fmt.Fprint(w, "Form[%q] = %q\n", k, v) 
    }
}



