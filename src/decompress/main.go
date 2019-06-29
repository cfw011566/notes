package main

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/trip", tripper)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}

func tripper(w http.ResponseWriter, r *http.Request) {
	var isDeflate = false
	var isGzip = false
	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		log.Printf("Header[%q] = %q\n", k, v)
		if k == "Content-Encoding" {
			log.Println(v[0])
			isDeflate = v[0] == "deflate"
			isGzip = v[0] == "gzip"
		}
	}
	log.Printf("Host = %q\n", r.Host)
	log.Printf("RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		log.Printf("Form[%q] = %q\n", k, v)
	}

	defer r.Body.Close()

	if isDeflate {
		z, err := zlib.NewReader(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		p, err := ioutil.ReadAll(z)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("p = %v\n", p)
		log.Println(string(p))
	} else if isGzip {
		log.Println("gzip")
		z, err := gzip.NewReader(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		p, err := ioutil.ReadAll(z)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("p = %v\n", p)
		log.Println(string(p))
	} else {
		contents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("content len = %d\n", len(contents))
		log.Printf("content = %v\n", contents)
		err = ioutil.WriteFile("out.bin", contents, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	/*
		f := flate.NewReader(r.Body)
		p, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("p = %v\n", p)
		log.Println(string(p))
	*/
}
