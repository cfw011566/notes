**sample `httpd.con` for fastcgi**

```
prefork 1

server "default" {
    listen on * port 80
    location "/foo" {
        fastcgi socket ":9000"
    }
    location "/bar*" {
        fastcgi socket ":9001"
    }
}
```

**sample fast CGI server code in golang**
```go
package main

import (
    "fmt"
    "log"
    "net"
    "net/http"
    "net/http/fcgi"
)

func main() {
    l2, err := net.Listen("tcp", "127.0.0.1:9001")
    if err != nil {
        panic(err)
    }
    go fcgi.Serve(l2, http.HandlerFunc(barHandler))

    l, err := net.Listen("tcp", "127.0.0.1:9000")
    if err != nil {
        panic(err)
    }
    if err := fcgi.Serve(l, http.HandlerFunc(fooHandler)); err != nil {
        panic(err)
    }
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
    fmt.Fprintf(w, "Hello, Foo")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Bar Hello\n")
    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    for k, v := range r.Form {
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }
}
```
